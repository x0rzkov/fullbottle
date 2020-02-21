package middleware

import (
	"github.com/micro/go-micro/v2/errors"
	pbauth "github.com/vegchic/fullbottle/auth/proto/auth"
	"github.com/vegchic/fullbottle/common"
	"github.com/vegchic/fullbottle/models"
	pbuser "github.com/vegchic/fullbottle/user/proto/user"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"net/http"
)

func LoginRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		// validate jwt token
		authClient := common.GetAuthSrvClient()
		authorization := c.GetHeader("authorization")
		if !strings.HasPrefix(authorization, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"msg": "Not authorized",
			})
			return
		}

		token := authorization[7:]
		authResp, err := authClient.ParseJwtToken(c, &pbauth.ParseJwtTokenRequest{Token: token})
		if err != nil {
			e := errors.Parse(err.Error())
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"msg": e.Detail,
			})
			return
		}

		// get user info
		client := common.GetUserSrvClient()
		userResp, err := client.GetUserInfo(c, &pbuser.GetUserRequest{Id: authResp.UserId})
		if err != nil {
			e := errors.Parse(err.Error())
			if e.Code == common.UserNotFound {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"msg": e.Detail,
				})
			} else {
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"msg": e.Detail,
				})
			}
			return
		}

		ct := time.Unix(userResp.CreateTime, 0)
		ut := time.Unix(userResp.UpdateTime, 0)
		dt := time.Unix(userResp.DeleteTime, 0)
		c.Set("CurrentUser", &models.User{
			BasicModel: models.BasicModel{
				ID:         userResp.Id,
				Status:     userResp.Status,
				CreateTime: &ct,
				UpdateTime: &ut,
				DeleteTime: &dt,
			},
			Username:  userResp.Username,
			Password:  userResp.Password,
			Email:     userResp.Email,
			Role:      userResp.Role,
			AvatarUrl: userResp.AvatarUrl,
		})
	}
}
