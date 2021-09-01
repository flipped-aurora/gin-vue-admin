package response

import "github.com/flipped-aurora/gin-vue-admin/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
