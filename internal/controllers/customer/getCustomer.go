package customer

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func GetCustomer(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var row = db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM `Customer` WHERE `Id`=%d",id))
		var customer modules.Customer  

		if e := row.Scan(&customer.Id, &customer.Name, &customer.Surname, &customer.Age, &customer.Sex, &customer.CardNumber); e != nil{
			if e==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"there is now row with such id")
				return
			}

			c.JSON(http.StatusInternalServerError,"error")
			
			return
		}
		fmt.Println(id)
		c.JSON(http.StatusOK,customer)

	}
}