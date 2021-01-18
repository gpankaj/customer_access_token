package access_token_service

import (
	"errors"
	"github.com/gpankaj/customer_access_token/src/domain/access_token"
	"github.com/gpankaj/customer_access_token/src/repository/db"
	"github.com/gpankaj/customer_access_token/src/repository/rest"
	rest_errors_package	"github.com/gpankaj/go-utils/rest_errors_package"
	"log"

	"strings"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *rest_errors_package.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors_package.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *rest_errors_package.RestErr
}

type service struct {
	restUsersRepo rest.RestPartnerRepostiryInterface
	dbRepo        db.DbRepository
}

func NewService(usersRepo rest.RestPartnerRepostiryInterface, dbRepo db.DbRepository) Service {
	return &service{
		restUsersRepo: usersRepo,
		dbRepo:        dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *rest_errors_package.RestErr) {
	accessTokenId = strings.TrimSpace(accessTokenId)
	if len(accessTokenId) == 0 {
		return nil, rest_errors_package.NewBadRequestError("invalid access token id")
	}
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *rest_errors_package.RestErr) {


	//TODO: Support both grant types: client_credentials and password

	// Authenticate the user against the Users API:
	customer, err := s.restUsersRepo.LoginCustomer(request.Customer_email_id, request.Customer_password)
	if err != nil {
		return nil, rest_errors_package.NewInternalServerError("Failed to login", errors.New("Error"))
	}

	// Generate a new access token:
	log.Println("Customer details -- Id ", customer.Customer_id)

	at := access_token.GetNewAccessToken(customer.Customer_id)
	at.Generate()

	log.Println("Access token has user id ", at.Customer_id);

	// Save the new access token in Cassandra:
	if err := s.dbRepo.Create(at); err != nil {
		return nil, err
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *rest_errors_package.RestErr {
	if err := at.Validate(); err != nil {
		return err
	}

	return s.dbRepo.UpdateExpirationTime(at)
}