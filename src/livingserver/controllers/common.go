package controllers

type CommonRsp struct{
	RetCode		int				`json:"ret_code"`
	Message		string			`json:"message"`
	Data		[]interface{}	`json:"data"`
}