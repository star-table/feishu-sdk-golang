package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"testing"
)

func TestUser_SearchUser(t *testing.T) {

	user := User{
		UserAccessToken: "u-X7iHiq3dWFWavEv7W1kQza",
	}

	resp, err := user.SearchUser("1", 0, "")
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))

}