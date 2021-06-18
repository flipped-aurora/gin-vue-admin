package response

import "github.com/eyotang/game-proxy/server/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
