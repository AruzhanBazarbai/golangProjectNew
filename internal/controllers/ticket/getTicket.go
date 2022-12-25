package ticket

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func GetTicket(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var row = db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM `tickets` WHERE `Id`=%d",id))
		var ticket modules.Ticket  

		if e := row.Scan(&ticket.Id, &ticket.FromWhere, &ticket.ToWhere, &ticket.DepartureDate, &ticket.DepartureTime, &ticket.ArrivalTime, &ticket.Duration, &ticket.Price); e != nil{
			if e==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"there is now row with such id")
				return
			}

			c.JSON(http.StatusInternalServerError,"error")
			
			return
		}
		fmt.Println(id)
		c.JSON(http.StatusOK, ticket)

	}
}
