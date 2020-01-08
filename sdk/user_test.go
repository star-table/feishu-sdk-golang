package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"testing"
)

func TestUser_SearchUser(t *testing.T) {

	user := User{
		UserAccessToken: "u-cSD8NdMRKX6hlDelxWqxSg",
	}

	resp, err := user.SearchUser("l", 0, "")
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestSearch(t *testing.T){
	body, e := http.Get("https://open.feishu.cn/api/v3/app/cli_9d5e49aae9ae9101/developer/search?name=l", nil)
	t.Log(e)
	t.Log(body)
}