package http

import (
	"github.com/galaxy-book/feishu-sdk-golang/core/model/vo"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"gotest.tools/assert"
	"testing"
)

func TestConvertToQueryParams(t *testing.T) {
	str := ConvertToQueryParams(map[string]interface{}{
		"a": "aa",
		"c": nil,
	})
	assert.Equal(t, str, "?a=aa")
}

func TestConvertToQueryParams2(t *testing.T) {

	t.Log(json.ToJsonIgnoreError(vo.UpdateChatReqVo{
		ChatId:      "111",
		OwnerOpenId: "111",
	}))
}
