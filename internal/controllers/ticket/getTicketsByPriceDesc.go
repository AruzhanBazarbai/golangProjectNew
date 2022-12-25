package ticket

import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"

	"go/basic/g/modules"
)

func GetTicketsByPriceDesc(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var rows *sql.Rows
		var e error

		if rows, e = db.Query("SELECT * FROM `tickets` ORDER BY `price` desc"); e != nil {
			c.JSON(http.StatusInternalServerError,e)
			return
		}
		defer rows.Close()

		var tickets []modules.Ticket

		for rows.Next(){
			var ticket modules.Ticket
			
			if e := rows.Scan(&ticket.Id, &ticket.FromWhere, &ticket.ToWhere, &ticket.DepartureDate, &ticket.DepartureTime, &ticket.ArrivalTime, &ticket.Duration, &ticket.Price); e!=nil{
				c.JSON(http.StatusInternalServerError, e)
				return
			}

			tickets = append(tickets, ticket)
		}

		if len(tickets)==0{
			c.JSON(http.StatusNotFound, sql.ErrNoRows)
			return
		}
		c.JSON(http.StatusOK, tickets)

	}
}