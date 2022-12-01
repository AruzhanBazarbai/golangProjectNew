package article

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"go/basic/g/modules"

	"github.com/gin-gonic/gin"
)

func PutArticle(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var id, _ = strconv.Atoi(c.Param("id"))
		var ctx = c.Request.Context()
		var article modules.Article

		if e := c.ShouldBindJSON(&article); e!=nil{
			c.JSON(http.StatusBadRequest,"please try again or enter the valid data")
			return
		}

		if _,e := db.ExecContext(ctx,fmt.Sprintf("UPDATE `Article` SET `name`='%s',`anons`='%s',`full_text`='%s',`image`='%s' WHERE `id`=%d;",article.Name,article.Anons,article.FullText,article.Image,id)); e!=nil {
			c.JSON(http.StatusInternalServerError, "internal server error")
			return
		}

		c.JSON(http.StatusOK,"succesfully updated!!!")

	}
}