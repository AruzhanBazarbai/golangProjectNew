package main

import (
	"go/basic/g/internal/controllers/customer"
	"go/basic/g/internal/controllers/article"
	"go/basic/g/internal/controllers/user"
	"go/basic/g/internal/controllers/ticket"
	"go/basic/g/internal/controllers/myTickets"
	"log"
	// "fmt"

	"github.com/gin-gonic/gin"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	// "go/basic/g/modules"
)

// var customers=[]Customer{}
func main(){
	var router = gin.Default()
	var address = ":8000"

	db,err:=sql.Open("mysql","root:@tcp(127.0.0.1:3306)/golangProject")
	if err!=nil{
		panic(err)
	}
	defer db.Close()
	var accounts = map[string]string{
		"john":"doe",
		"foo":"bar",
	}
	var authMiddleware = gin.BasicAuth(accounts)
	// customer endpoints
	router.GET("/customers",authMiddleware,customer.GetCustomers(db))
	router.GET("/customers/:id",authMiddleware,customer.GetCustomer(db))
	router.DELETE("/customers/:id",authMiddleware,customer.DeleteCustomer(db))
	router.POST("/customers",authMiddleware,customer.PostCustomer(db))
	router.PUT("/customers/:id",authMiddleware,customer.PutCustomer(db))
	// article endpoints
	router.GET("/articles",authMiddleware,article.GetArticles(db))
	router.GET("/articles/:id",authMiddleware,article.GetArticle(db))
	router.DELETE("/articles/:id",authMiddleware,article.DeleteArticle(db))
	router.POST("/articles",authMiddleware,article.PostArticle(db))
	router.PUT("/articles/:id",authMiddleware,article.PutArticle(db))
	// user endpoints
	router.GET("/users",authMiddleware,user.GetUsers(db))
	router.GET("/users/:id",authMiddleware,user.GetUser(db))
	router.DELETE("/users/:id",authMiddleware,user.DeleteUser(db))
	router.POST("/users",authMiddleware,user.PostUser(db))
	router.PUT("/users/:id",authMiddleware,user.PutUser(db))
	// ticket endpoints
	router.GET("/tickets",authMiddleware,ticket.GetTickets(db))
	router.GET("/tickets/:id",authMiddleware,ticket.GetTicket(db))
	router.DELETE("/tickets/:id",authMiddleware,ticket.DeleteTicket(db))
	router.POST("/tickets",authMiddleware,ticket.PostTicket(db))
	router.PUT("/tickets/:id",authMiddleware,ticket.PutTicket(db))
	// myticket endpoints
	router.GET("/myTickets",authMiddleware,myTickets.GetMyTickets(db))
	router.GET("/myTickets/:id",authMiddleware,myTickets.GetMyTickets(db))
	// router.DELETE("/myTickets/:id",myTickets.DeleteMyTicket(db))
	router.POST("/myTickets",authMiddleware,myTickets.PostMyTicket(db))
	// router.PUT("/myTickets/:id",myTickets.PutMyTicket(db))

	log.Fatalln(router.Run(address))

}