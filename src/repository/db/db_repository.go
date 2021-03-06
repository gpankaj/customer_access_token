package db

import (
	"github.com/gocql/gocql"
	"github.com/gpankaj/customer_access_token/src/clients/cassandra"
	"log"

	"github.com/gpankaj/customer_access_token/src/domain/access_token"

	"github.com/gpankaj/go-utils/rest_errors_package"

)

const (
	queryAccessToken = "SELECT access_token, customer_id, client_id, expires FROM access_tokens WHERE access_token=?;"
	queryInsertToken = "INSERT into access_tokens(access_token, customer_id, client_id, expires) VALUES (?,?,?,?);"
	queryUpdateToken = "UPDATE access_tokens SET expires=? WHERE access_token=?;"

)
func NewRepository() DbRepository {
	return &dbRepository{}
}

type DbRepository interface {

	GetById(string) (*access_token.AccessToken, *rest_errors_package.RestErr)
	Create(access_token.AccessToken) (*rest_errors_package.RestErr)
	UpdateExpirationTime(access_token.AccessToken)(*rest_errors_package.RestErr)
}

type dbRepository struct {

}

func (db *dbRepository) GetById(id string) (*access_token.AccessToken, *rest_errors_package.RestErr) {


	var result access_token.AccessToken

	if err := cassandra.GetSession().Query(queryAccessToken,id).Scan(
					&result.Access_token,
					&result.Customer_id,
					&result.Client_id,
					&result.Expires); err!= nil{

		if err == gocql.ErrNotFound {
			return nil, rest_errors_package.NewNotFoundError("no access token found with the given id")
		}
		return nil,rest_errors_package.NewInternalServerError("GetSession failed",err)
	}
	return &result,nil
}

func (db *dbRepository) Create(token access_token.AccessToken) (*rest_errors_package.RestErr){

	//queryInsertToken
	log.Println("Adding Cassandra entry ", token)
	if err := cassandra.GetSession().Query(queryInsertToken,token.Access_token,token.Customer_id, token.Client_id, token.Expires).Exec(); err!= nil{
		return rest_errors_package.NewInternalServerError("Failed to connect to cassendra", err)
	}

	return nil
}

func (db *dbRepository) UpdateExpirationTime(token access_token.AccessToken) (*rest_errors_package.RestErr){

	//queryInsertToken
	if err := cassandra.GetSession().Query(queryUpdateToken,token.Expires,token.Access_token).Exec(); err!= nil{
		return rest_errors_package.NewInternalServerError("Failed to get Session",err)
	}
	return nil
}