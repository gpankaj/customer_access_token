package access_token

import (
	"fmt"
	"github.com/gpankaj/storage_partners_api/utils/crypto_utils"
	"github.com/gpankaj/go-utils/rest_errors_package"

	"time"
)



const (
	expirationTime = 24
)
type AccessToken struct {
	Access_token	string
	User_id			int64
	Client_id		int64 //web or android or something else..third party so that we can define expiration of client

	Expires 		int64 //timestamp when access token actually expires.
}

type AccessTokenRequest struct {
	Email_id string
	Password string
}


func (at *AccessToken) Validate() *rest_errors_package.RestErr {
	if at.Access_token == "" {
		return rest_errors_package.NewBadRequestError(fmt.Sprintf("Invalid access token id %s" , at.Access_token))
	}

	if len(at.Access_token) == 0 {
		return rest_errors_package.NewBadRequestError(fmt.Sprintf("Invalid access token id %s" , at.Access_token))
	}

	if at.User_id <= 0 {
		return rest_errors_package.NewBadRequestError(fmt.Sprintf("Invalid User_id %s" , at.User_id))
	}


	if at.Client_id <= 0 {
		return rest_errors_package.NewBadRequestError(fmt.Sprintf("Invalid clientid %s" , at.Client_id))
	}

	if  at.Expires <= 0 {
		return rest_errors_package.NewBadRequestError(fmt.Sprintf("Invalid Expires token id %s" , at.Expires))
	}

	return nil
}

func GetNewAccessToken(id int64) AccessToken {
	return AccessToken{
		User_id: id,
		Expires: time.Now().UTC().Add(expirationTime*time.Hour).Unix(),
	}
}

func (at *AccessToken) Generate() {
	at.Access_token = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.User_id, at.Expires))
}
func (at AccessToken) IsExpired()bool{
	now := time.Now().UTC()
	expirationTime := time.Unix(at.Expires, 0)
	fmt.Println(expirationTime)

	return expirationTime.Before(now)
}