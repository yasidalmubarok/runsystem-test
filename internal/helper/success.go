package helper

type SuccessResponse struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Code	int    `json:"code"`
	Message string `json:"message"`
	Status  string `json:"status"`
}

func ApiResponse(code int, message string, status string, data interface{}) SuccessResponse {
	meta := Meta{
		Code:    code,
		Message: message,
		Status:  status,
	}

	jsonResponse := SuccessResponse{
		Meta: meta,
		Data: data,
	}

	return jsonResponse
}
