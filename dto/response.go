package dto

func ResponseFailed(msg string, code int) map[string]interface{} {
	return map[string]interface{}{
		"status":      "error",
		"status_code": code,
		"message":     msg,
	}
}

func ResponseSuccesNoData(msg string) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
	}
}

func ResponseSuccesWithData(msg string, data interface{}) map[string]interface{} {
	return map[string]interface{}{
		"status":  "success",
		"message": msg,
		"data":    data,
	}
}

type PayloadID struct {
	Id int `json:"id"`
}

type ResponseMsg struct {
	Messageresp string `json:"message"`
}
