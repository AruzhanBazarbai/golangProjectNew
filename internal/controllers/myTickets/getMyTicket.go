package myTickets

import (
	"database/sql"
	"fmt"
	"net/http"
	// "strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func GetMyTicket(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		id:= c.GetHeader("userId")
		fmt.Println(id)
		// var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var row = db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM `myTickets` WHERE `Id`=%d",id))
		var myTickets modules.MyTickets

		if e := row.Scan(&myTickets.Id, &myTickets.UserId, &myTickets.TicketId); e != nil{
			if e==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"there is now row with such id")
				return
			}

			c.JSON(http.StatusInternalServerError,"error")
			
			return
		}
		fmt.Println(id)
		c.JSON(http.StatusOK, myTickets)

	}
}