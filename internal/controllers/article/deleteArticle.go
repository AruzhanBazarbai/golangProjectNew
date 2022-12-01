package article

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)

func DeleteArticle(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()

		var result sql.Result
		var e error
		if result, e = db.ExecContext(ctx, fmt.Sprintf("DELETE FROM `Article` WHERE `id`=%d",id)); e!=nil{
			c.JSON(http.StatusInternalServerError,"internal error")
			return
		}

		if nProducts,_ := result.RowsAffected(); nProducts==0{
			c.JSON(http.StatusNotFound, "there is no such row")
			return

		}

		c.JSON(http.StatusNoContent,nil)

	}
}