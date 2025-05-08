package utils

func Response(status int, message string, data interface{}) (int, interface{}) {
	return status, map[string]interface{}{
		"data":    data,
		"message": message,
	}
}

func ResponseError(status int, message string) (int, interface{}) {
	return status, map[string]interface{}{
		"message": message,
	}
}
