package customer

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func PutCustomer(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var putResCustomer modules.PutCustomer

		if e := c.ShouldBindJSON(&putResCustomer); e!=nil{
			c.JSON(http.StatusBadRequest,"please try again or enter the valid data")
			return
		}

		if _,e := db.ExecContext(ctx,fmt.Sprintf("UPDATE `Customer` SET `name`='%s',`surname`='%s',`age`=%d,`sex`='%s',`card_number`='%s' WHERE `id`=%d;",putResCustomer.Name,putResCustomer.Surname,putResCustomer.Age,putResCustomer.Sex,putResCustomer.CardNumber,id)); e!=nil {
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}
		// чтобы проверить все успешно обновили или нет
		// var res modules.Customer
		// var row = db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM `Customer` WHERE `Id`=%d",id))
		// if e := row.Scan(&res.Id, &res.Name, &res.Surname, &res.Age, &res.Sex, &res.CardNumber); e != nil{
		// 	if e==sql.ErrNoRows{
		// 		c.JSON(http.StatusNotFound,"there is now row with such id")
		// 		return
		// 	}

		// 	c.JSON(http.StatusInternalServerError,"error")
		// 	return
		// }
		// c.JSON(http.StatusOK, res)

		c.JSON(http.StatusOK,"succesfully updated!!!")

	}
}