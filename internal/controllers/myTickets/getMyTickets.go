package myTickets

import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"

	"go/basic/g/modules"
)

func GetMyTickets(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var rows *sql.Rows
		var e error

		if rows, e = db.Query("SELECT * FROM `myTickets`"); e != nil {
			c.JSON(http.StatusInternalServerError,e)
			return
		}
		defer rows.Close()

		var myTickets []modules.MyTickets

		for rows.Next(){
			var myTicket modules.MyTickets
			
			if e := rows.Scan(&myTicket.Id, &myTicket.UserId, &myTicket.TicketId); e!=nil{
				c.JSON(http.StatusInternalServerError, e)
				return
			}

			myTickets = append(myTickets, myTicket)
		}

		if len(myTickets)==0{
			c.JSON(http.StatusNotFound, sql.ErrNoRows)
			return
		}
		c.JSON(http.StatusOK, myTickets)

	}
}