package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func PutUser(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var user modules.User

		if e := c.ShouldBindJSON(&user); e!=nil{
			c.JSON(http.StatusBadRequest,"please try again or enter the valid data")
			return
		}

		if _,e := db.ExecContext(ctx,fmt.Sprintf("UPDATE `User` SET `name`='%s',`password`='%s' WHERE `id`=%d;",user.Name,user.Password,id)); e!=nil {
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		c.JSON(http.StatusOK,"succesfully updated!!!")

	}
}