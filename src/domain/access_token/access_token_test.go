package access_token

import "testing"

func TestGetNewAccessToken(t *testing.T) {
	at := GetNewAccessToken()
	if at.IsExpired() {
		t.Error("Brand new access token should not be expired.")
	}
	if at.Access_token != "" {
		t.Error("New Access Token should not have defined Access_token field value inside AccessToken")
	}

	if at.User_id !=0 {
		t.Error("New Access Token should not have associated user_id")
	}
}

func TestAccessToken_IsExpired(t *testing.T) {
	at := GetNewAccessToken()

	if at.IsExpired() {
		t.Error("empty access token should be expired")
	}
}