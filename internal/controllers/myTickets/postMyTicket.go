package myTickets

import (
	"database/sql"
	"fmt"
	"go/basic/g/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostMyTicket(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var myTickets modules.MyTickets
		var ctx = c.Request.Context()

		if e:=c.ShouldBindJSON(&myTickets); e!=nil{
			c.JSON(http.StatusBadRequest,"error")
			return
		}

		// var row = db.QueryRowContext(ctx, "SELECT MAX(Id) FROM `tickets`")  
		var row = db.QueryRowContext(ctx, "SELECT MAX(Id) FROM `myTickets`")  

		if e := row.Scan(&myTickets.Id); e != nil{
			if e==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"ERROR")
				return
			}
			c.JSON(http.StatusInternalServerError, row.Scan(&myTickets.Id))
		}
		c.JSON(http.StatusOK, myTickets)
 
		if _, e := db.ExecContext(ctx, fmt.Sprintf("INSERT INTO `myTickets` VALUES(%d, %d, %d);", myTickets.Id + 1, myTickets.UserId, myTickets.TicketId)); e != nil {
			c.JSON(http.StatusInternalServerError,"internal error")
			return
		}


		c.Writer.Header().Add("Location",fmt.Sprintf("/myTickets/%d",myTickets.Id+1))
		c.JSON(http.StatusCreated,"succesfully created")
	}
	
}