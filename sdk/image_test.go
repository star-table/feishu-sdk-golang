package sdk

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"os"
	"testing"
)

func TestTenant_GetImage(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	resp, err := tenant.GetImage("img_f9713733-cb19-4a4b-88f8-7d5a3937b5ag", false)
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))

	fp, err := os.Create("a.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer fp.Close()

	//buf := new(bytes.Buffer)
	//binary.Write(buf, binary.LittleEndian, resp)
	//fp.Write(buf.Bytes())
	//
	// 创建OSSClient实例。
	client, err := oss.New("https://oss-cn-shanghai.aliyuncs.com", "LTAIrqZIswji9vEL", "fAlOQ3cp0idrhBYHg0hkIL8QLZ4Owt")
	if err != nil {
		fmt.Println("Error1:", err)
		os.Exit(-1)
	}

	// 获取存储空间。
	bucket, err := client.Bucket("polaris-hd2")
	if err != nil {
		fmt.Println("Error2:", err)
		os.Exit(-1)
	}

	//// 读取本地文件。
	//fd, err := os.Open("C:\\Users\\admin\\Desktop\\11111.png")
	//if err != nil {
	//	fmt.Println("Error3:", err)
	//	os.Exit(-1)
	//}
	//defer fd.Close()

	//上传文件流。
	err = bucket.PutObject("test/test5.png", bytes.NewReader(resp))
	if err != nil {
		fmt.Println("Error4:", err)
		os.Exit(-1)
	}
}
