package response

import "devops-manage/model/example"

type ExaCustomerResponse struct {
	Customer example.ExaCustomer `json:"customer"`
}
