package controllers

import (
	"database/sql"
	"fmt"
	"go/basic/g/modules"
	"strings"

	// "strconv"
	"net/http"
	"time"
	"errors"

	"github.com/golang-jwt/jwt/v4"

	// "fmt"
	// "net/http"
	"go/basic/g/internal/controllers/user"
	// "strconv"

	// "go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

const(
	signingKey = "hcvjasedljsbdev23"
	authorizationHeader= "Authorization"
	userCtx="userId"
)
var UserId=0;
type tokenClaims struct{
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func SignUp(db *sql.DB) gin.HandlerFunc{
	return user.PostUser(db);
}

func SignIn(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var userSignedIn modules.User
		var ctx = c.Request.Context()

		if e:=c.ShouldBindJSON(&userSignedIn); e!=nil{
			c.JSON(http.StatusBadRequest,"error")
			return
		}

		
		var row = db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM `User` WHERE `Name`='%s' AND `Password`='%s' LIMIT 1;",userSignedIn.Name,userSignedIn.Password))

		var user modules.User
		// fmt.Println(userSignedIn.Id, userSignedIn.Name, userSignedIn.Password);
		err := row.Scan(&user.Id, &user.Name, &user.Password)

		if err != nil{
			if err==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"there is no such registered user")
				return
			}

			c.JSON(http.StatusInternalServerError,"error")
			
			return
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
			jwt.StandardClaims{
				ExpiresAt: time.Now().Add(12*time.Hour).Unix(),
				IssuedAt: time.Now().Unix(),
			},
			user.Id,

		})

		t,err:=token.SignedString([]byte(signingKey))

		c.Set(authorizationHeader, fmt.Sprintf("Bearer %s",t))
		// c.Set(userId,)
		a,_:=c.Get(authorizationHeader)
		fmt.Println(a, "ok")

		c.JSON(http.StatusOK,map[string]interface{}{
			// "user":user,
			// "err": err,
			"token": token,
			"token_payload": t,
		})

	}
	

	
}

func UserIdentity(c *gin.Context){
	header:=c.GetHeader(authorizationHeader)
	// fmt.Println(header)

	if header == ""{
		c.JSON(http.StatusUnauthorized,"empty auth header")
		return
	}
	headerParts:=strings.Split(header, " ")
	if len(headerParts)!=2{
		c.JSON(http.StatusUnauthorized,"invalid auth header")
		return
	}
	accessToken:=headerParts[1]
	// fmt.Println(accessToken)
	// fmt.Println("----------------")
	
	// parse token
	token,err:=jwt.ParseWithClaims(accessToken, &tokenClaims{},func(token *jwt.Token)(interface{},error){
		if _,ok:=token.Method.(*jwt.SigningMethodHMAC);!ok{
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey),nil
	})
	if err!=nil{
		c.JSON(http.StatusUnauthorized,"invalid signing method ddd")
		return
	}
	claims,ok:=token.Claims.(*tokenClaims)
	if !ok{
		c.JSON(http.StatusUnauthorized,"error")
		return
	}
	// c.Get(userCtx)
	c.Set(userCtx,claims.UserId) // в контекст пишем юзер айди на переменную userCtx="userId"
	c.Set("token",accessToken)
	id,_:= c.Get(userCtx)
	UserId=claims.UserId
	// fmt.Println(id)
	

	c.JSON(http.StatusOK,map[string]interface{}{
		"message": "Successful",
		"user id": claims.UserId,
		"id":id,
	})
}


// var jsonData = JSON.parse(responseBody);
// postman.setEnvironmentVariable("token", jsonData.token_payload);

