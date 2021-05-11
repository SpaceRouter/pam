package pam

import (
	"encoding/json"
	"log"
	"os/user"
	"runtime"
	"testing"
)

func TestGetUserInfos001(t *testing.T) {
	u, _ := user.Current()
	if u.Uid != "0" {
		t.Skip("run this test as root")
	}
	infos, err := GetUserInfos("test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	encoded, _ := json.Marshal(infos)
	log.Println(string(encoded))
	if infos.Username == "" {

	}
	runtime.GC()
}
