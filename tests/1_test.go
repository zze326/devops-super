package tests

import (
	"devops-super/utility/util"
	"fmt"
	"github.com/logrusorgru/aurora"
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

	fmt.Println(aurora.BgRed(aurora.White("aaa")))
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
