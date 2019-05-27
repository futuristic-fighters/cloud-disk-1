package cfg

type response struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
	Msg  string      `json:"msg"`
}

func Response() *response {
	return &response{}
}

func NewErrResponse(c int) *response {
	return &response{Code: c, Msg: Lang[c]}
}

func NewResponse(c int, d interface{}) *response {
	return &response{Code: c, Msg: Lang[c], Data: d}
}

func (r *response) SetData(d interface{}) *response {
	r.Data = d
	return r
}

func (r *response) SetCode(c int) *response {
	r.Code = c
	return r
}

func (r *response) SetMsg(m string) *response {
	r.Msg = m
	return r
}
