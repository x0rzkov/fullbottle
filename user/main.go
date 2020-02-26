package main

import (
	"github.com/micro/go-micro/v2"

	"github.com/micro/go-micro/v2/client"
	gclient "github.com/micro/go-micro/v2/client/grpc"

	gserver "github.com/micro/go-micro/v2/server/grpc"

	"github.com/vegchic/fullbottle/common"
	"github.com/vegchic/fullbottle/common/log"
	"github.com/vegchic/fullbottle/config"
	"github.com/vegchic/fullbottle/user/handler"

	user "github.com/vegchic/fullbottle/user/proto/user"
)

func options() []micro.Option {
	return []micro.Option{
		micro.Server(gserver.NewServer( // need to pass first
			gserver.MaxMsgSize(config.MaxMsgSize),
		)),

		micro.Name(config.UserSrvName),
		micro.Version("latest"),

		micro.Client(gclient.NewClient(
			client.WrapCall(common.ClientLogWrapper),
			gclient.MaxSendMsgSize(config.MaxMsgSendSize),
			gclient.MaxRecvMsgSize(config.MaxMsgRecvSize))),

		//micro.WrapHandler(common.ServiceErrorRecovery),
		micro.WrapHandler(common.ServiceLogWrapper),
	}
}

func main() {
	service := micro.NewService(options()...)

	service.Init()

	common.SetClient(service.Client())

	if err := user.RegisterUserServiceHandler(service.Server(), new(handler.UserHandler)); err != nil {
		log.WithError(err).Fatalf("RegisterUserServiceHandler failed")
	}

	if err := service.Run(); err != nil {
		log.WithError(err).Fatalf("Service running failed")
	}
}
