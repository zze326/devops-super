package tests

import (
	"devops-super/utility/util"
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"strings"
	"testing"
)

func Test1_1(t *testing.T) {
	t.Log(util.EncryptPassword("admin123"))
	//a := "$2a$10$yCU6QndNxyrLGjEMSn.YIOXph2LfaedbnrGKkZJ8vpcpaDhKMSb5K"
	b := "$2a$10$lmPQ9aQ287342NpIdZctH.m9vP1gljNGJExhT1TVhbisVia.0cNRC"
	t.Log(util.ComparePassword(b, "admin123"))
}

func Test1(t *testing.T) {
	t.Log(gtime.Now().Unix())
	t.Log(gtime.Now().UnixMilli())
	t.Log(gtime.Now().UnixMicro())
	t.Log(gtime.Now().UnixNano())
}

func Test1_3(t *testing.T) {
	var iv int = 3
	switch iv {
	case 1 | 2:
		fmt.Println("into 1")
	case 4:
		fmt.Println("into 4")
	}
	fmt.Println("testSwitch run end")
}

func Test1_4(t *testing.T) {
	route := "get:/test/:id"
	t.Log(strings.SplitN(route, ":", 2))
}
