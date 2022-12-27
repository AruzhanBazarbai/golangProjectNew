package modules

type Tick struct{
	FromWhere 		string		`json:"fromWhere" binding:"required"`
	ToWhere 		string  	`json:"toWhere" binding:"required"`
	DepartureDate 	string		`json:"departureDate" binding:"required"`
}