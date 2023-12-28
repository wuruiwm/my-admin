package config

type AdminConfig struct {
	Notice     Notice     `json:"notice"`
	Ssl        Ssl        `json:"ssl"`
	K8s        K8s        `json:"k8s"`
	Script     Script     `json:"script"`
	Cloudflare Cloudflare `json:"cloudflare"`
	Nas        Nas        `json:"nas"`
	Pay        Pay        `json:"pay"`
}

type Notice struct {
	Type              string `json:"type"`
	EmailServerHost   string `json:"email_server_host"`
	EmailUsername     string `json:"email_username"`
	EmailPassword     string `json:"email_password"`
	EmailIsEncrypt    int    `json:"email_is_encrypt,string"`
	EmailPort         int    `json:"email_port,string"`
	EmailReceiveUser  string `json:"email_receive_user"`
	GotifyServerUrl   string `json:"gotify_server_url"`
	GotifyServerToken string `json:"gotify_server_token"`
}

type Ssl struct {
	Path string `json:"path"`
}

type K8s struct {
	Host          string `json:"host"`
	Port          int    `json:"port,string"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	AdminUsername string `json:"admin_username"`
}

type Script struct {
	TwLolLuckDrawSk string `json:"tw_lol_luck_draw_sk"`
	YoutubeSaveDir  string `json:"youtube_save_dir"`
	M3u8SaveDir     string `json:"m3u8_save_dir"`
}

type Cloudflare struct {
	ZoneId string `json:"zone_id"`
	Key    string `json:"key"`
	Email  string `json:"email"`
	Dns    string `json:"dns"`
}

type Nas struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Pay struct {
	Config string `json:"config"`
}
