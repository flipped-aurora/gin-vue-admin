package response

import "gin-vue-admin/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
