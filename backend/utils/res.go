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

func ResponseMsg(status string, message string) gin.H {
	return Response(status, message, nil, nil)
}
