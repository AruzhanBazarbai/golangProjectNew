package customer

import (
	"database/sql"
	"fmt"
	"go/basic/g/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostCustomer(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var customer modules.Customer
		var ctx = c.Request.Context()

		if e:=c.ShouldBindJSON(&customer); e!=nil{
			c.JSON(http.StatusBadRequest,"error")
			return
		}
 
		if _, e := db.ExecContext(ctx, fmt.Sprintf("INSERT INTO `Customer` VALUES(%d,'%s','%s',%d,'%s','%s');",customer.Id,customer.Name,customer.Surname,customer.Age,customer.Sex,customer.CardNumber)); e != nil {
			c.JSON(http.StatusInternalServerError,"internal error")
			return
		}
		// чтобы проверить существуют в базе данных сохранились или нет введенные данные
		// var res modules.Customer
		// var row = db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM `Customer` WHERE `Id`=%d",customer.Id))
		// if e := row.Scan(&res.Id, &res.Name, &res.Surname, &res.Age, &res.Sex, &res.CardNumber); e != nil{
		// 	if e==sql.ErrNoRows{
		// 		c.JSON(http.StatusNotFound,"there is now row with such id")
		// 		return
		// 	}

		// 	c.JSON(http.StatusInternalServerError,"error")
		// 	return
		// }
		// c.JSON(http.StatusOK, res)

		c.Writer.Header().Add("Location",fmt.Sprintf("/customers/%d",customer.Id))
		c.JSON(http.StatusCreated,"succesfully created")
	}
	
}