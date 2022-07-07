package dao

import (
	"encoding/json"
	"pool-event/models"

	"github.com/Viva-con-Agua/vcago"
)

func UserGetRequest() (payload []models.User, err error) {
	uRL := "/admin/users"
	response := new(vcago.Response)
	if response, err = AdminRequest.Get(uRL); err != nil {
		return nil, vcago.NewErrorLog(err, "ERROR", "request")
	}
	payload = []models.User{}
	if response.Payload != nil {
		bytes, _ := json.Marshal(&response.Payload)
		_ = json.Unmarshal(bytes, &payload)
	}
	if payload == nil {
		return nil, vcago.NewBadRequest("message", "no user found with recipient group")
	}
	return
}