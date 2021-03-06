// Modify from https://github.com/bsm/redislock/blob/master/redislock.go

package kv

import (
	"errors"
	"github.com/go-redis/redis/v7"
	"github.com/vegchic/fullbottle/common"
	"github.com/vegchic/fullbottle/util"
	"strconv"
	"time"
)

var (
	luaRefresh = redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("pexpire", KEYS[1], ARGV[2]) else return 0 end`)
	luaRelease = redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("del", KEYS[1]) else return 0 end`)
	luaPTTL    = redis.NewScript(`if redis.call("get", KEYS[1]) == ARGV[1] then return redis.call("pttl", KEYS[1]) else return -3 end`)
)

const DefaultTTL = 400 * time.Millisecond

type Lock struct {
	client *redis.Client
	key    string
	value  string
}

func (l *Lock) Key() string {
	return l.key
}

func (l *Lock) Token() string {
	return l.value
}

func (l *Lock) TTL() (time.Duration, error) {
	res, err := luaPTTL.Run(l.client, []string{l.key}, l.value).Result()
	if err == redis.Nil {
		return 0, nil
	} else if err != nil {
		return 0, err
	}

	if num := res.(int64); num > 0 {
		return time.Duration(num) * time.Millisecond, nil
	}
	return 0, nil
}

func (l *Lock) Refresh(ttl time.Duration) error {
	ttlVal := strconv.FormatInt(int64(ttl/time.Millisecond), 10)
	res, err := luaRefresh.Run(l.client, []string{l.key}, l.value, ttlVal).Result()
	if err != nil {
		return err
	} else if res == 1 {
		return nil
	}

	return common.NewRedisError(errors.New("cannot refresh lock"))
}

func (l *Lock) Release() error {
	res, err := luaRelease.Run(l.client, []string{l.key}, l.value).Result()
	if err == redis.Nil {
		return common.NewRedisError(errors.New("cannot release lock"))
	} else if err != nil {
		return err
	}

	if i, ok := res.(int64); !ok || i != 1 {
		return common.NewRedisError(errors.New("cannot release lock"))
	}
	return nil
}

func Obtain(key string, ttl time.Duration) (*Lock, error) {
	token := util.GenToken(10)

	var timer *time.Timer
	for ddl := time.Now().Add(ttl); time.Now().Before(ddl); {
		ok, err := client.SetNX(key, token, ttl).Result()
		if err != nil {
			return nil, err
		} else if ok {
			return &Lock{client: client, key: key, value: token}, nil
		}

		if timer == nil {
			timer = time.NewTimer(time.Millisecond)
			defer timer.Stop()
		} else {
			timer.Reset(time.Millisecond)
		}

		select {
		case <-timer.C:
		}
	}

	return nil, common.NewRedisError(errors.New("cannot obtain lock"))
}
