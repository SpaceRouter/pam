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
	infos := GetUserInfos("louis")
	encoded, _ := json.Marshal(infos)
	log.Println(string(encoded))
	runtime.GC()
}
