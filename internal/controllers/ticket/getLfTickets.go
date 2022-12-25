package ticket

import (
	"database/sql"
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
			var res *sql.Rows
			var e error

			var t modules.Ticket

			if res, e = db.Query("SELECT FROM `tickets` WHERE `tick.FromWhere` = &tickets[i].FromWhere AND `tick.ToWhere` = `tickets[i].ToWhere` AND `tick.DepartureDate` = `tickets[i].DepartureDate`"); e != nil {
				c.JSON(http.StatusInternalServerError,"internal error")
				return
			}
			defer res.Close()

			if e := res.Scan(&t.Id, &t.FromWhere, &t.ToWhere, &t.DepartureDate, &t.DepartureTime, &t.ArrivalTime, &t.Duration, &t.Price); e!=nil{
				c.JSON(http.StatusInternalServerError, e)
				return
			}

			ticks = append(ticks, t)

			c.JSON(http.StatusOK, ticks)
		
		}

		


	}
}