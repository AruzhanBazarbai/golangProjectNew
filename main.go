package main

import (
	"go/basic/g/internal/controllers"
	"go/basic/g/internal/controllers/article"
	"go/basic/g/internal/controllers/customer"
	"go/basic/g/internal/controllers/myTickets"
	"go/basic/g/internal/controllers/ticket"
	"go/basic/g/internal/controllers/user"

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

	auth:=router.Group("/auth")
	{
		// authorization endpoints
		auth.POST("/sign-up",controllers.SignUp(db))
		auth.POST("/sign-in",controllers.SignIn(db))
	}
	
	api:=router.Group("/")
	{
		// customer endpoints
		customers:=api.Group("/customers",controllers.UserIdentity)
		{
			customers.GET("/",customer.GetCustomers(db))
			customers.GET("/:id",customer.GetCustomer(db))
			customers.DELETE("/:id",customer.DeleteCustomer(db))
			customers.POST("/",customer.PostCustomer(db))
			customers.PUT("/:id",customer.PutCustomer(db))
		}
		// article endpoints
		articles:=api.Group("/articles",controllers.UserIdentity)
		{
			articles.GET("/",article.GetArticles(db))
			articles.GET("/:id",article.GetArticle(db))
			articles.DELETE("/:id",article.DeleteArticle(db))
			articles.POST("/",article.PostArticle(db))
			articles.PUT("/:id",article.PutArticle(db))
		}
		// user endpoints
		users:=api.Group("/users",controllers.UserIdentity)
		{
			users.GET("/",user.GetUsers(db))
			users.GET("/:id",user.GetUser(db))
			users.DELETE("/:id",user.DeleteUser(db))
			users.POST("/",user.PostUser(db))
			users.PUT("/:id",user.PutUser(db))
		}
		// ticket endpoints
		tickets:=api.Group("/tickets",controllers.UserIdentity)
		{
			tickets.GET("/",ticket.GetTickets(db))
			tickets.GET("/:id",ticket.GetTicket(db))
			tickets.DELETE("/:id",ticket.DeleteTicket(db))
			tickets.POST("/",ticket.PostTicket(db))
			tickets.PUT("/:id",ticket.PutTicket(db))
			tickets.POST("/get",ticket.GetLfTickets(db))
			tickets.GET("/asc",ticket.GetTicketsByPriceAsc(db))
			tickets.GET("/desc",ticket.GetTicketsByPriceDesc(db))
		}
		// myticket endpoints
		myTicket:=api.Group("/myTickets",controllers.UserIdentity)
		{
			myTicket.GET("/",myTickets.GetMyTickets(db))
			myTicket.POST("/",myTickets.PostMyTicket(db))
		}
		
	}
	log.Fatalln(router.Run(address))
}