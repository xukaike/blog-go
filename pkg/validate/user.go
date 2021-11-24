package validate

type Register struct {
	Name string `json:"name" binding:"required,max=255"`
	Email string `json:"email" binding:"email"`
	Password string `json:"password" binding:"gte=6"`
	Desc string `json:"desc"`
	Avatar string `json:"avatar"`
}

type Login struct {
	Name string `json:"name" binding:"required,max=255"`
	Password string `json:"password" binding:"gte=6"`
}