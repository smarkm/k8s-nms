package cmd

//Result API result wap
type Result struct {
	Code int
	Msg  string
	Data interface{}
}

//OK rs
func OK(data interface{}) *Result {
	return &Result{Code: 0, Msg: "success", Data: data}
}

//Error rs
func Error(code int, msg string, data interface{}) *Result {
	return &Result{Code: code, Msg: msg, Data: data}
}
