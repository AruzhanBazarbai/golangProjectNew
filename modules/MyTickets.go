package modules

type MyTickets struct{
	Id 				int 		`json:"id"`
	UserId 			int			`json:"userId"`
	TicketId 		int  		`json:"ticketId" binding:"required"`
}