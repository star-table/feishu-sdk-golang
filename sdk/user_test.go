package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"testing"
)

func TestUser_SearchUser(t *testing.T) {

	user := User{
		UserAccessToken: "u-JJSpC2QT03VM3c5hV3I7ic",
	}

	resp, err := user.SearchUser("1", 0, "")
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))

}