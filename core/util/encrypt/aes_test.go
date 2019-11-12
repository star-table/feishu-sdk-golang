package encrypt

import (
	"gotest.tools/assert"
	"testing"
)

func TestAesDecrypt(t *testing.T) {

	key := "test key"
	encrypt := "P37w+VZImNgPEO1RBhJ6RtKl7n6zymIbEG1pReEzghk="

	d, e := AesDecrypt(key, encrypt)
	t.Log(e)
	t.Log(d)

	assert.Equal(t, d, "hello world")
}