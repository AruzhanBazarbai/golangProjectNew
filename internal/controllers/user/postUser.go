package user

import (
	"database/sql"
	"fmt"
	"go/basic/g/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostUser(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var user modules.User
		var ctx = c.Request.Context()

		if e:=c.ShouldBindJSON(&user); e!=nil{
			c.JSON(http.StatusBadRequest,"error")
			return
		}
		

		var row = db.QueryRowContext(ctx, "SELECT MAX(Id) FROM `User`" )  

		if e := row.Scan(&user.Id); e != nil{
			if e==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"ERROR")
				return
			}
			c.JSON(http.StatusInternalServerError,"ERRORRR")
		}
		c.JSON(http.StatusOK, user)
 
		if _, e := db.ExecContext(ctx, fmt.Sprintf("INSERT INTO `User` VALUES(%d,'%s','%s');",user.Id+1,user.Name,user.Password)); e != nil {
			c.JSON(http.StatusInternalServerError,"internal error")
			return
		}
		
			

		c.Writer.Header().Add("Location",fmt.Sprintf("/users/%d",user.Id+1))
		c.JSON(http.StatusCreated,"succesfully created")
	}	
}