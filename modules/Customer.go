package modules
type Customer struct{
	Id int `json:"id" binding:"required"`
	Name string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Age int `json:"age" binding:"required,gt=0"`
	Sex string `json:"sex" binding:"required,oneof=Male Female"`
	CardNumber string `json:"card_number" binding:"required"`
}

type PutCustomer struct{
	Name string `json:"name" binding:"required"`
	Surname string `json:"surname" binding:"required"`
	Age int `json:"age" binding:"required,gt=0"`
	Sex string `json:"sex" binding:"required,oneof=Male Female"`
	CardNumber string `json:"card_number" binding:"required"`
} 

