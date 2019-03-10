package library
type Resp struct {
	Errno int `json:"errno""`
	Errmsg string `json:"errmsg"`
	Data interface{} `json:"data""`
}

func Respmsg(errno int,errmsg string ,data interface{}) interface{}  {
	resp :=Resp{}
	resp.Errno =errno
	resp.Errmsg=errmsg
	resp.Data=data
	return resp
}

