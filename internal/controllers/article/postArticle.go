package article

import (
	"database/sql"
	"fmt"
	"go/basic/g/modules"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostArticle(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var article modules.Article
		var ctx = c.Request.Context()

		if e:=c.ShouldBindJSON(&article); e!=nil{
			c.JSON(http.StatusBadRequest,"error")
			return
		}
		

		var row = db.QueryRowContext(ctx, "SELECT MAX(Id) FROM `Article`" )  

		if e := row.Scan(&article.Id); e != nil{
			if e==sql.ErrNoRows{
				c.JSON(http.StatusNotFound,"ERROR")
				return
			}
			c.JSON(http.StatusInternalServerError,"ERRORRR")
		}
		c.JSON(http.StatusOK, article)
 
		if _, e := db.ExecContext(ctx, fmt.Sprintf("INSERT INTO `Article` VALUES(%d,'%s','%s','%s','%s');",article.Id+1,article.Name,article.Anons,article.FullText,article.Image)); e != nil {
			c.JSON(http.StatusInternalServerError,"internal error")
			return
		}
		
			

		c.Writer.Header().Add("Location",fmt.Sprintf("/articles/%d",article.Id+1))
		c.JSON(http.StatusCreated,"succesfully created")
	}	
}