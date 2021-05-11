package response

import "github.com/eyotang/gin-vue-admin/server/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
