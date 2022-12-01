package modules

type Article struct{
	Id int `json:"id"`
	Name string `json:"name" binding:"required"`
	Anons string `json:"anons" binding:"required"`
	FullText string `json:"full_text" binding:"required"`
	Image string `json:"image" binding:"required"`
}
