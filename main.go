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
	// customer endpoints
	router.GET("/customers",customer.GetCustomers(db))
	router.GET("/customers/:id",customer.GetCustomer(db))
	router.DELETE("/customers/:id",customer.DeleteCustomer(db))
	router.POST("/customers",customer.PostCustomer(db))
	router.PUT("/customers/:id",customer.PutCustomer(db))
	// article endpoints
	router.GET("/articles",article.GetArticles(db))
	router.GET("/articles/:id",article.GetArticle(db))
	router.DELETE("/articles/:id",article.DeleteArticle(db))
	router.POST("/articles",article.PostArticle(db))
	router.PUT("/articles/:id",article.PutArticle(db))
	// user endpoints
	router.GET("/users",user.GetUsers(db))
	router.GET("/users/:id",user.GetUser(db))
	router.DELETE("/users/:id",user.DeleteUser(db))
	router.POST("/users",user.PostUser(db))
	router.PUT("/users/:id",user.PutUser(db))
	// ticket endpoints
	router.GET("/tickets",ticket.GetTickets(db))
	router.GET("/tickets/:id",ticket.GetTicket(db))
	router.DELETE("/tickets/:id",ticket.DeleteTicket(db))
	router.POST("/tickets",ticket.PostTicket(db))
	router.PUT("/tickets/:id",ticket.PutTicket(db))
	// myticket endpoints
	router.GET("/myTickets",myTickets.GetMyTickets(db))
	router.GET("/myTickets/:id",myTickets.GetMyTickets(db))
	// router.DELETE("/myTickets/:id",myTickets.DeleteMyTicket(db))
	router.POST("/myTickets",myTickets.PostMyTicket(db))
	// router.PUT("/myTickets/:id",myTickets.PutMyTicket(db))

	log.Fatalln(router.Run(address))

}