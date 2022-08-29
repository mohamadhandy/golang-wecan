package helper

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  string `json:"status"`
}

func ResponseAPI(data interface{}, status string, code int, message string) Response {
	meta := Meta{
		Message: message,
		Status:  status,
		Code:    code,
	}
	return Response{
		Meta: meta,
		Data: data,
	}
}
