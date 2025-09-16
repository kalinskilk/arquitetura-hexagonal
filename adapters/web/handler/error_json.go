package handler

import "encoding/json"


func jsonError(msg string)[]byte{
	errr:= struct {
				Message string `json:"message"`
	}{
	Message:msg,
	}

	r, err:= json.Marshal(errr)
	if err !=nil{
		return []byte(`{"message":"internal error"}`)
	}
	return r
}