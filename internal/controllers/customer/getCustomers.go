package customer

import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"

	"go/basic/g/modules"
)

func GetCustomers(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var rows *sql.Rows
		var e error

		if rows, e = db.Query("SELECT * FROM `Customer`"); e != nil {
			c.JSON(http.StatusInternalServerError,e)
			return
		}
		defer rows.Close()

		var customers []modules.Customer

		for rows.Next(){
			var customer modules.Customer
			
			if e := rows.Scan(&customer.Id, &customer.Name, &customer.Surname, &customer.Age, &customer.Sex, &customer.CardNumber); e!=nil{
				c.JSON(http.StatusInternalServerError, e)
				return
			}

			customers = append(customers, customer)
		}

		if len(customers)==0{
			c.JSON(http.StatusNotFound, sql.ErrNoRows)
			return
		}
		c.JSON(http.StatusOK,customers)

	}
}