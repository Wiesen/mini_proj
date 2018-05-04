package controllers

type PostLikeReq struct {
	EmotionID		int		`json:"emotion_id"`
}

type PostEmotionReq struct {
	Content		string		`json:"content"`
	LabelID		int			`json:"label_id"`
	Strong		int8		`json:"strong"`
	Visiable	int8		`json:"visiable"`
}

type CommonRsp struct{
	RetCode		int				`json:"ret_code"`
	Message		string			`json:"message"`
	Data		[]interface{}	`json:"data"`
}

