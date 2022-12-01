package ticket

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func PutTicket(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var ticket modules.Ticket

		if e := c.ShouldBindJSON(&ticket); e!=nil{
			c.JSON(http.StatusBadRequest,"please try again or enter the valid data")
			return
		}

		if _,e := db.ExecContext(ctx,fmt.Sprintf("UPDATE `tickets` SET `fromWhere`='%s',`toWhere`='%s',`departureDate`='%s',`departureTime`='%s', `arrivalTime`='%s', `duration`='%s', `price`=%d  WHERE `id`=%d;", ticket.FromWhere, ticket.ToWhere, ticket.DepartureDate, ticket.DepartureTime,ticket.ArrivalTime, ticket.Duration, ticket.Price, id)); e!=nil {
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		c.JSON(http.StatusOK,"succesfully updated!!!")

	}
}