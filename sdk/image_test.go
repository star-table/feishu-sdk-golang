package sdk

import (
	"bytes"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	"os"
	"strings"
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

	////上传文件流。
	//err = bucket.PutObject("test/test6.png", resp)
	//if err != nil {
	//	fmt.Println("Error4:", err)
	//	os.Exit(-1)
	//}

	var nextPos int64 = 0
	objName := "test/test8.png"
	defer resp.Close()
	//分片逐步写入
	buf := make([]byte, 4096)
	for {
		n, err := resp.Read(buf)
		if err != nil {
			break
		}

		nextPos, err = bucket.AppendObject(objName, bytes.NewReader(buf[:n]), nextPos)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
	}

	//// 第一次追加上传的位置是0，返回值为下一次追加的位置。后续追加的位置是追加前文件的长度。
	//// <yourObjectName>填写不包含Bucket名称在内的Object的完整路径，例如example/test.txt。
	//nextPos, err = bucket.AppendObject(objName, strings.NewReader("YourObjectAppendValue1"), nextPos)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}
	//
	//// 第二次追加上传。
	//nextPos, err = bucket.AppendObject(objName, strings.NewReader("YourObjectAppendValue2"), nextPos)
	//if err != nil {
	//	fmt.Println("Error:", err)
	//	os.Exit(-1)
	//}

}

func TestBuildInternalApp(t *testing.T) {
	// 创建OSSClient实例。
	client, err := oss.New("<yourEndpoint>", "<yourAccessKeyId>", "<yourAccessKeySecret>")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	bucketName := "<yourBucketName>"

	// 获取存储空间。
	bucket, err := client.Bucket(bucketName)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
	var nextPos int64 = 0
	// 第一次追加上传的位置是0，返回值为下一次追加的位置。后续追加的位置是追加前文件的长度。
	// <yourObjectName>填写不包含Bucket名称在内的Object的完整路径，例如example/test.txt。
	nextPos, err = bucket.AppendObject("<yourObjectName>", strings.NewReader("YourObjectAppendValue1"), nextPos)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}

	// 如果不是第一次追加上传，可以通过bucket.GetObjectDetailedMeta方法或上次追加返回值的X-Oss-Next-Append-Position的属性，获取追加位置。
	//props, err := bucket.GetObjectDetailedMeta("<yourObjectName>")
	//if err != nil {
	//    fmt.Println("Error:", err)
	//    os.Exit(-1)
	//}
	//nextPos, err = strconv.ParseInt(props.Get("X-Oss-Next-Append-Position"), 10, 64)
	//if err != nil {
	//    fmt.Println("Error:", err)
	//    os.Exit(-1)
	//}

	// 第二次追加上传。
	nextPos, err = bucket.AppendObject("<yourObjectName>", strings.NewReader("YourObjectAppendValue2"), nextPos)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(-1)
	}
}
