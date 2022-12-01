package article

import (
	"net/http"
	"database/sql"

	"github.com/gin-gonic/gin"

	"go/basic/g/modules"
)

func GetArticles(db *sql.DB) gin.HandlerFunc{
	return func(c *gin.Context){
		var rows *sql.Rows
		var e error

		if rows, e = db.Query("SELECT * FROM `Article`"); e != nil {
			c.JSON(http.StatusInternalServerError,e)
			return
		}
		defer rows.Close()

		var articles []modules.Article

		for rows.Next(){
			var article modules.Article
			
			if e := rows.Scan(&article.Id, &article.Name, &article.Anons, &article.FullText, &article.Image); e!=nil{
				c.JSON(http.StatusInternalServerError, e)
				return
			}

			articles = append(articles, article)
		}

		if len(articles)==0{
			c.JSON(http.StatusNotFound, sql.ErrNoRows)
			return
		}
		c.JSON(http.StatusOK,articles)

	}
}