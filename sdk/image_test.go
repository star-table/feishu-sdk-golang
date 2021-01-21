package sdk

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/galaxy-book/feishu-sdk-golang/core/consts"
	"github.com/galaxy-book/feishu-sdk-golang/core/util/json"
	uuid "github.com/satori/go.uuid"
	"os"
	"strconv"
	"strings"
	"sync"
	"testing"
)

var fileTypeMap sync.Map

func TestTenant_NewFileUploadRequest(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2c29657237ce175d")
	t.Log(e)

	//t.Log(NewFileUploadRequest("https://open.feishu.cn/open-apis/image/v4/put/",
	//		data, "image", "C:\\Users\\admin\\Desktop\\11111.png"))

	data := make(map[string]string)
	resp := tenant.NewFileUploadRequest("https://open.feishu.cn/open-apis/image/v4/put/", data, "image", "C:\\Users\\admin\\Desktop\\12345.png")
	t.Log(json.ToJsonIgnoreError(resp))
}

func TestTenant_GetImage(t *testing.T) {
	app, e := BuildApp(consts.TestAppId, consts.TestAppSecret, consts.TestTicket)
	t.Log(e)
	t.Log(json.ToJsonIgnoreError(app))
	tenant, e := BuildTenant(app.AppAccessToken, "2ed263bf32cf1651")
	t.Log(e)

	//resp, err := tenant.GetImage("img_b2092abb-5a13-4e68-bd04-fb70da58227g", false)
	resp, err := tenant.GetImage("img_ead2e693-606e-40a0-a979-80243a19ee8g", false)
	//resp, err := tenant.GetImage("img_f9713733-cb19-4a4b-88f8-7d5a3937b5ag", false)
	t.Log(err)
	t.Log(json.ToJsonIgnoreError(resp))

	// 创建OSSClient实例。
	client, err := oss.New("xxxxxx", "xxxxx", "xxxxxxx")
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

	//var nextPos int64 = 0
	//objName := "test/test28.png"
	//defer resp.Close()
	////分片逐步写入
	//buf := make([]byte, 40960)
	//for {
	//	n, err := resp.Read(buf)
	//	if err != nil {
	//		break
	//	}
	//
	//	nextPos, err = bucket.AppendObject(objName, bytes.NewReader(buf[:n]), nextPos)
	//	if err != nil {
	//		fmt.Println("Error:", err)
	//		os.Exit(-1)
	//	}
	//}

	var nextPos int64 = 0
	objName := "test/" + NewUuid()
	fileType := ""
	defer resp.Close()
	//分片逐步写入
	buf := make([]byte, 1024)
	i := 0
	for {
		n, err := resp.Read(buf)
		if err != nil {
			break
		}
		if i == 0 {
			fileType = GetFileType(buf[:n])
			fmt.Println("fileType:", buf[:n])
			if fileType != "" {
				objName = objName + "." + fileType
			}
		}
		i++
		fmt.Println(i)

		nextPos, err = bucket.AppendObject(objName, bytes.NewReader(buf[:n]), nextPos)
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(-1)
		}
	}
	t.Log(objName)

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

func NewUuid() string {
	id := uuid.NewV4()
	return strings.ReplaceAll(id.String(), "-", "")
}

//4749463839612c013002f400007a7a
//47494638396137023702f644

// 用文件前面几个字节来判断
// fSrc: 文件字节流（就用前面几个字节）
func GetFileType(fSrc []byte) string {
	//fileTypeMap.Store("ffd8ffe000104a464946", "jpg")  //JPEG (jpg)
	fileTypeMap.Store("ffd8", "jpg") //JPEG (jpg)
	//fileTypeMap.Store("89504e470d0a1a0a0000", "png")  //PNG (png)
	fileTypeMap.Store("89504e", "png")               //PNG (png)
	fileTypeMap.Store("47494638396126026f01", "gif") //GIF (gif)
	//fileTypeMap.Store("49492a00227105008037", "tif")  //TIFF (tif)
	fileTypeMap.Store("49492a", "tif")                //TIFF (tif)
	fileTypeMap.Store("424d228c010000000000", "bmp")  //16色位图(bmp)
	fileTypeMap.Store("424d8240090000000000", "bmp")  //24位位图(bmp)
	fileTypeMap.Store("424d8e1b030000000000", "bmp")  //256色位图(bmp)
	fileTypeMap.Store("41433130313500000000", "dwg")  //CAD (dwg)
	fileTypeMap.Store("3c21444f435459504520", "html") //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c68746d6c3e0", "html")        //HTML (html)   3c68746d6c3e0  3c68746d6c3e0
	fileTypeMap.Store("3c21646f637479706520", "htm")  //HTM (htm)
	fileTypeMap.Store("48544d4c207b0d0a0942", "css")  //css
	fileTypeMap.Store("696b2e71623d696b2e71", "js")   //js
	fileTypeMap.Store("7b5c727466315c616e73", "rtf")  //Rich Text Format (rtf)
	//fileTypeMap.Store("38425053000100000000", "psd")  //Photoshop (psd)
	fileTypeMap.Store("384250", "psd")               //Photoshop (psd)
	fileTypeMap.Store("46726f6d3a203d3f6762", "eml") //Email [Outlook Express 6] (eml)
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "doc") //MS Excel 注意：word、msi 和 excel的文件头一样
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "vsd") //Visio 绘图
	fileTypeMap.Store("5374616E64617264204A", "mdb") //MS Access (mdb)
	fileTypeMap.Store("252150532D41646F6265", "ps")
	fileTypeMap.Store("255044", "pdf") //Adobe Acrobat (pdf)
	//fileTypeMap.Store("255044462d312e350d0a", "pdf")  //Adobe Acrobat (pdf)
	fileTypeMap.Store("2e524d46000000120001", "rmvb") //rmvb/rm相同
	fileTypeMap.Store("464c5601050000000900", "flv")  //flv与f4v相同
	fileTypeMap.Store("00000020667479706d70", "mp4")
	fileTypeMap.Store("49443303000000002176", "mp3")
	fileTypeMap.Store("000001ba210001000180", "mpg") //
	fileTypeMap.Store("3026b2758e66cf11a6d9", "wmv") //wmv与asf相同
	fileTypeMap.Store("52494646e27807005741", "wav") //Wave (wav)
	fileTypeMap.Store("52494646d07d60074156", "avi")
	fileTypeMap.Store("4d546864000000060001", "mid") //MIDI (mid)
	fileTypeMap.Store("504b0304140000000800", "zip")
	fileTypeMap.Store("526172211a0700cf9073", "rar")
	fileTypeMap.Store("235468697320636f6e66", "ini")
	fileTypeMap.Store("504b03040a0000000000", "jar")
	fileTypeMap.Store("4d5a9000030000000400", "exe")        //可执行文件
	fileTypeMap.Store("3c25402070616765206c", "jsp")        //jsp文件
	fileTypeMap.Store("4d616e69666573742d56", "mf")         //MF文件
	fileTypeMap.Store("3c3f786d6c2076657273", "xml")        //xml文件
	fileTypeMap.Store("494e5345525420494e54", "sql")        //xml文件
	fileTypeMap.Store("7061636b616765207765", "java")       //java文件
	fileTypeMap.Store("406563686f206f66660d", "bat")        //bat文件
	fileTypeMap.Store("1f8b0800000000000000", "gz")         //gz文件
	fileTypeMap.Store("6c6f67346a2e726f6f74", "properties") //bat文件
	fileTypeMap.Store("cafebabe0000002e0041", "class")      //bat文件
	fileTypeMap.Store("49545346030000006000", "chm")        //bat文件
	fileTypeMap.Store("04000000010000001300", "mxp")        //bat文件
	fileTypeMap.Store("504b0304140006000800", "docx")       //docx文件
	fileTypeMap.Store("d0cf11e0a1b11ae10000", "wps")        //WPS文字wps、表格et、演示dps都是一样的
	fileTypeMap.Store("6431303a637265617465", "torrent")
	fileTypeMap.Store("6D6F6F76", "mov")         //Quicktime (mov)
	fileTypeMap.Store("FF575043", "wpd")         //WordPerfect (wpd)
	fileTypeMap.Store("CFAD12FEC5FD746F", "dbx") //Outlook Express (dbx)
	fileTypeMap.Store("2142444E", "pst")         //Outlook (pst)
	fileTypeMap.Store("AC9EBD8F", "qdf")         //Quicken (qdf)
	fileTypeMap.Store("E3828596", "pwl")         //Windows Password (pwl)
	fileTypeMap.Store("2E7261FD", "ram")         //Real Audio (ram)

	var fileType string
	fileCode := bytesToHexString(fSrc)
	fmt.Println(fileCode)

	fileTypeMap.Range(func(key, value interface{}) bool {
		k := key.(string)
		v := value.(string)
		if strings.HasPrefix(fileCode, strings.ToLower(k)) ||
			strings.HasPrefix(k, strings.ToLower(fileCode)) {
			fileType = v
			return false
		}
		return true
	})
	return fileType
}

// 获取前面结果字节的二进制
func bytesToHexString(src []byte) string {
	res := bytes.Buffer{}
	if src == nil || len(src) <= 0 {
		return ""
	}
	temp := make([]byte, 0)
	for _, v := range src {
		sub := v & 0xFF
		hv := hex.EncodeToString(append(temp, sub))
		if len(hv) < 2 {
			res.WriteString(strconv.FormatInt(int64(0), 10))
		}
		res.WriteString(hv)
	}
	return res.String()
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
