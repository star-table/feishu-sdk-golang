package sdk

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/http"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/log"
	"gotest.tools/assert"
	"testing"
)

func TestUser_SearchUser(t *testing.T) {

	user := User{
		UserAccessToken: "u-ahufHbuFHqXVOM0HqlXCxd",
	}

	resp, err := user.SearchUser("abcde", 0, "")
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestSearch(t *testing.T) {
	body, e := http.Get("https://open.feishu.cn/api/v3/app/cli_9d5e49aae9ae9101/developer/search?name=l", nil)
	t.Log(e)
	t.Log(body)
}

func TestTenant_BatchGetId(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "12243f37560ed740")
	t.Log(e)

	resp, err := tenant.BatchGetId([]string{"xxxxxx@qq.com"}, []string{"+86xxxxxxx", "xxxxx"})
	log.Info(json.ToJsonIgnoreError(resp), err)
	assert.Equal(t, err, nil)
	assert.Equal(t, resp.Code, 0)
}
