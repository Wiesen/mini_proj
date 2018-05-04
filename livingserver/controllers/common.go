package controllers

type PostLikeReq struct {
	EmotionID		int		`json:"emotion_id"`
}

type CommonRsp struct{
	RetCode		int				`json:"ret_code"`
	Message		string			`json:"message"`
	Data		[]interface{}	`json:"data"`
}

