package response

import "github.com/flipped-aurora/gin-vue-admin/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
