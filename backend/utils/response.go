package utils

type Body struct {
	Status  string                 `json:"status"`
	Message string                 `json:"message"`
	Data    interface{}            `json:"data"`
	Others  map[string]interface{} `json:"others"`
}

func NewBody(body Body) map[string]interface{} {
	res := make(map[string]interface{})

	if res["status"] = body.Status; len(body.Status) == 0 {
		res["status"] = "success"
	}
	res["message"] = body.Message
	res["data"] = body.Data
	for k, v := range body.Others {
		res[k] = v
	}
	return res
}
