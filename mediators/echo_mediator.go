package mediators

import "fmt"

type Response struct {
	Message string
}

func NewEchoMediator(bodyDict map[string]interface{}) *Response {
	name := bodyDict["name"]
	responseMsg := fmt.Sprintf("Hello %s!", name)

	return &Response{
		Message: responseMsg,
	}
}