package response

import "asyncClient/model"

type ExaCustomerResponse struct {
	Customer model.ExaCustomer `json:"customer"`
}
