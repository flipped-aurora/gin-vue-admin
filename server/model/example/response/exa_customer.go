package response

import "github.com/gzpz/golf-sales-system/server/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
