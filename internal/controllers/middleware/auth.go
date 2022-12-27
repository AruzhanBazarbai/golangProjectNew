package middle

import (
	"errors"
	"fmt"
	"strings"

	// "gin/gonic"
	"net/http"

	// "go/token"
	"github.com/gin-gonic/gin"
	"github.com/techschool/simplebank/token"
)

const (
	authorizationHeaderKey = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

func authMiddleware(tokenMaker token.Maker) gin.HandlerFunc{
	return func(ctx *gin.Context){
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)
		if len(authorizationHeaderKey) == 0 {
			err := errors.New("Authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.Error(err))
			return
		}
		fields:=strings.Fields(authorizationHeader)
		if len(fields)<2{
			err := errors.New("Invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.Error(err))
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if(authorizationType!=authorizationTypeBearer){
			err := fmt.Errorf("unsupported authorization type %s",authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.Error(err))
			return
		}
		accessToken := fields[1]
		payload, err :=tokenMaker.VerifyToken(accessToken)
		if err!=nil{
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, error.Error(err))
			return
		}

		ctx.Set(authorizationPayloadKey,payload)
		ctx.Next()

	}
}