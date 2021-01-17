package rest

import (
	"encoding/json"
	"errors"
	"github.com/gpankaj/storage_access_tokken/src/domain/partners"
	"github.com/gpankaj/go-utils/rest_errors_package"
	"github.com/mercadolibre/golang-restclient/rest"
	"log"
	"time"
)

type RestPartnerRepostiryInterface interface {

	LoginPartner(string, string)(*partners.Partner, *rest_errors_package.RestErr)
}

type restPartnerRepostiry struct {

}

var (
	partnersRestClient = rest.RequestBuilder{
		BaseURL: "http://localhost:8080",
		Timeout: 200 * time.Millisecond,
	}
)
func NewRepository() RestPartnerRepostiryInterface {
	return &restPartnerRepostiry{}
}

func (r *restPartnerRepostiry) LoginPartner(email string, password string) (*partners.Partner, *rest_errors_package.RestErr){
	request:=partners.PartnerLoginRequest{Email_id: email, Password: password}
	log.Println("Received email id ", request.Email_id)
	log.Println("Received Password ", request.Password)

	response := partnersRestClient.Post("/partners/login", request)
	log.Println("Response ", response.Response.Status)

	if response == nil || response.Response == nil  { //Timeout situation.
		return nil, rest_errors_package.NewInternalServerError("invalid restClientRequest when trying to login partner", errors.New(""))
	}

	if response.StatusCode > 299 { //Means we have an error situation.
		var restError rest_errors_package.RestErr
		err := json.Unmarshal(response.Bytes(), &restError)
		if err!= nil {
			return nil, rest_errors_package.NewInternalServerError("invalid error interface, when trying to login user", err)
		}
		return nil, &restError
	}
	var partner partners.Partner
	if err := json.Unmarshal(response.Bytes(), &partner); err!=nil {
		return nil, rest_errors_package.NewInternalServerError("Mismatch in signature of partner data", err)
	}
	return &partner, nil
}
