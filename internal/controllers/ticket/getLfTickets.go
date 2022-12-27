package ticket

import (
	"database/sql"
	"fmt"
	"strings"

	// "fmt"
	// "fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"go/basic/g/modules"
)

func GetLfTickets(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var tick modules.Tick
		// var ctx = c.Request.Context()

		var rows *sql.Rows
		var e error

		if e:=c.ShouldBindJSON(&tick); e!=nil{
			c.JSON(http.StatusBadRequest,"error")
			return
		}

		if rows, e = db.Query("SELECT * FROM `tickets`"); e != nil {
			c.JSON(http.StatusInternalServerError,e)
			return
		}
		defer rows.Close()

		var tickets []modules.Ticket
		var ticks []modules.Ticket

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
		// for i := 0; i < len(tickets); i++ {
		// 	fmt.Println(tickets[i].DepartureDate)
		// }
		// c.JSON(http.StatusOK, tickets)
		for i := 0; i < len(tickets); i++ {
			// fmt.Println(tickets[i].FromWhere)
			// fmt.Println(tick.ToWhere)
			if(strings.TrimSpace(tickets[i].FromWhere)==tick.FromWhere && strings.TrimSpace(tickets[i].ToWhere)==tick.ToWhere && strings.TrimSpace(tickets[i].DepartureDate)==tick.DepartureDate){
				ticks = append(ticks, tickets[i])
				fmt.Println("findddd")
			}

			
		
		}

		
		c.JSON(http.StatusOK, ticks)
		// c.JSON(http.StatusOK, tick)
	}
}