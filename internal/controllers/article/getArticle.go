package article

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func GetArticle(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var row = db.QueryRowContext(ctx, fmt.Sprintf("SELECT * FROM `Article` WHERE `Id`=%d",id))
		var article modules.Article  

		if e := row.Scan(&article.Id, &article.Name, &article.Anons, &article.FullText, &article.Image); e != nil{
			if e==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"there is now row with such id")
				return
			}

			c.JSON(http.StatusInternalServerError,"error")
			fmt.Println(id)
			return
		}
		c.JSON(http.StatusOK,article)

	}
}