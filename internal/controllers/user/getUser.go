package user

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func GetUser(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var row = db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM `User` WHERE `Id`=%d",id))
		var user modules.User  

		if e := row.Scan(&user.Id, &user.Name, &user.Password); e != nil{
			if e==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"there is now row with such id")
				return
			}

			c.JSON(http.StatusInternalServerError,"error")
			
			return
		}
		fmt.Println(id)
		c.JSON(http.StatusOK,user)

	}
}