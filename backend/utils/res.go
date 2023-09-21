package utils

import "github.com/gin-gonic/gin"

func Response(status string, message string, data interface{}, others gin.H) gin.H {
	res := gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	}
	for k, v := range others {
		res[k] = v
	}

	return res
}

func ResponseMsg(message string) gin.H {
	return Response("error", message, nil, nil)
}
