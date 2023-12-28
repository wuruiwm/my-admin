package response

type Ssl struct {
	Key        string   `json:"key"`
	Pem        string   `json:"pem"`
	Domain     string   `json:"domain"`
	List       []string `json:"list"`
	ExpireTime string   `json:"expire_time"`
}
