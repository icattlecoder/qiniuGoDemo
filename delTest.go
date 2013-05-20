package main

import (
	"fmt"
	. "github.com/qiniu/api/conf"
	"github.com/qiniu/api/rs"
	"strconv"
)

func main() {
	ACCESS_KEY = "gPhMyVzzbQ_LOjboaVsy7dbCB4JHgyVPonmhT3Dp"
	SECRET_KEY = "OjY7IMysXu1erRRuWe7gkaiHcD6-JMJ4hXeRPZ1B"
	bucketName := "icattlecoder"
	rsClient := rs.New()
	for i:= 1;i<10;i++{
		rsClient.Delete(nil,bucketName,"test.go"+strconv.Itoa(i))
	}
	rsClient.Delete(nil, bucketName, "fileKey")
	fmt.Print("end!\n")
}
