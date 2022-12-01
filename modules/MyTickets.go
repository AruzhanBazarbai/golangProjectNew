package modules

type MyTickets struct{
	Id 				int 		`json:"id"`
	UserId 			int			`json:"userId" binding:"required"`
	TicketId 		int  		`json:"ticketId" binding:"required"`
}