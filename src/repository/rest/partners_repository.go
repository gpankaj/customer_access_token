package rest

import (
	"encoding/json"
	"errors"
	"github.com/gpankaj/customer_access_token/src/domain/customers"
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/mercadolibre/golang-restclient/rest"
	"log"
	"strings"
	"time"
)

type RestPartnerRepostiryInterface interface {

	LoginCustomer(string, string)(*customers.Customer, *rest_errors_package.RestErr)
}

type restCustomerRepostiry struct {

}

var (
	partnersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 200 * time.Millisecond,
	}
)
func NewRepository() RestPartnerRepostiryInterface {
	return &restCustomerRepostiry{}
}

func (r *restCustomerRepostiry) LoginCustomer(Customer_email_id string, Customer_password string) (*customers.Customer, *rest_errors_package.RestErr){
	request:=customers.CustomerLoginRequest{Customer_email_id: strings.TrimSpace(Customer_email_id),
		Customer_password: strings.TrimSpace(Customer_password)}

	log.Println("Received email id", request.Customer_email_id)
	log.Println("Received Password", request.Customer_password)

	response := partnersRestClient.Post("/customers/login", request)


	if response == nil || response.Response == nil  { //Timeout situation.
		log.Println(response.Err.Error())
		return nil, rest_errors_package.NewInternalServerError("invalid restClientRequest when trying to login customer",
			errors.New(""))
	}

	if response.StatusCode > 299 { //Means we have an error situation.
		var restError rest_errors_package.RestErr
		err := json.Unmarshal(response.Bytes(), &restError)
		if err!= nil {
			return nil, rest_errors_package.NewInternalServerError("invalid error interface, when trying to login customer", err)
		}
		return nil, &restError
	}
	var customer customers.Customer
	if err := json.Unmarshal(response.Bytes(), &customer); err!=nil {
		return nil, rest_errors_package.NewInternalServerError("Mismatch in signature of customer data", err)
	}
	return &customer, nil
}
