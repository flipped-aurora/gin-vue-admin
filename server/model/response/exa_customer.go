package response

import "github.com/eyotang/game-api-admin/server/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
