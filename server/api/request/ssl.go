package request

type Ssl struct {
	Domain string `form:"domain" binding:"required"`
}
