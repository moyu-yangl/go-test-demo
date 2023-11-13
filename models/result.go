package models

func ResultSuccessDefault() *Result {
	var result = Result{Code: 200, Message: "success"}
	return &result
}

func ResultSuccess(data interface{}) *Result {
	var result = Result{Code: 200, Message: "success", Data: data}
	return &result
}

func ResultError(code int, message string) *Result {
	var result = Result{Code: code, Message: message}
	return &result
}
