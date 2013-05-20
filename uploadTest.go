package main

import (
	"fmt"
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/io"
	"github.com/qiniu/api/rs"
	"os"
//	"strconv"
)

func main() {
	ACCESS_KEY = "gPhMyVzzbQ_LOjboaVsy7dbCB4JHgyVPonmhT3Dp"
	SECRET_KEY = "OjY7IMysXu1erRRuWe7gkaiHcD6-JMJ4hXeRPZ1B"
	bucketName := "icattlecoder"
	policy := rs.PutPolicy{

		// [string] 必须, 指定授权上传的bucket
		Scope: bucketName,

		// [int] 表示有效时间为3600秒, 即一个小时
		Expires: 3600,

		// [string] 用于设置文件上传成功后，七牛云存储服务端要回调客户方的业务服务器地址
		//CallbackUrl: "http://example.com",

		// [string] 用于设置文件上传成功后，七牛云存储服务端向客户方的业务服务器发送回调请求的 `Content-Type`
		CallbackBodyType: "application/x-www-form-urlencoded",

		// [string] 客户方终端用户（End User）的ID，该字段可以用来标示一个文件的属主，这在一些特殊场景下（比如给终端用户上传的图片打上名字水印）非常有用。
		Customer: "",

		// [string] 用于设置文件上传成功后，执行指定的预转指令。
		// 参考 http://docs.qiniutek.com/v3/api/io/#uploadToken-asyncOps
		AsyncOps: "",

		// [uint16] 可选值 0 或者 1，缺省为 0 。值为 1 表示 callback 传递的自定义数据中允许存在转义符号 `$(VarExpression)
		// 参考 http://docs.qiniutek.com/v3/api/words/#VarExpression
		Escape: 0,

		// [uint16] 可选值 0 或者 1, 缺省为 0. 值为 1 表示在服务端自动识别上传文件类型.
		DetectMime: 0,
	}

	// 生成 uploadToken, string类型
	token := policy.Token()
	filePath := "./Monkey 1.bmp"
	f, _ := os.Open(filePath)
	defer f.Close()

	// 声明 PutExtra
	extra := &io.PutExtra{
		// [string] 必选, 指定上传的目标仓库
		Bucket: bucketName,

		// [string] 可选。在 uptoken 没有指定 DetectMime 时，用户客户端可自己指定 MimeType

		// [string] 可选。用户自定义 Meta，不能超过 256 字节
		CustomMeta: "",

		// [string] 当 uptoken 指定了 CallbackUrl，则 CallbackParams 必须非空
		//CallbackParams: "",
	}
	key := "Monkey1.bmp"
	var ret io.PutRet
/*
	for i := 0; i < 10; i++ {
		tmpkey := key + strconv.Itoa(i)
		f, _ := os.Open(filePath)
		err := io.Put(nil, &ret, token, tmpkey, f, extra)
		if err == nil {
			hash := ret.Hash
			fmt.Print(hash)
			f.Close()
		}
	}
*/
	err := io.Put(nil, &ret, token, key, f, extra) // PutRet, error
	if err != nil {
		// 上传失败
		fmt.Print("error!\n")
		fmt.Print(err)
		return
	}
	// 当上传成功后, 得到的hash值
	hash := ret.Hash
	fmt.Print(hash)
	fmt.Print("end!\n")

}
