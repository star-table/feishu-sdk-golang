package http

import (
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
