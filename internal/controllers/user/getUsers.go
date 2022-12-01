package user

import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"

	"go/basic/g/modules"
)

func GetUsers(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var rows *sql.Rows
		var e error

		if rows, e = db.Query("SELECT * FROM `User`"); e != nil {
			c.JSON(http.StatusInternalServerError,e)
			return
		}
		defer rows.Close()

		var users []modules.User

		for rows.Next(){
			var user modules.User
			
			if e := rows.Scan(&user.Id, &user.Name, &user.Password); e!=nil{
				c.JSON(http.StatusInternalServerError, e)
				return
			}

			users = append(users, user)
		}

		if len(users)==0{
			c.JSON(http.StatusNotFound, sql.ErrNoRows)
			return
		}
		c.JSON(http.StatusOK,users)

	}
}