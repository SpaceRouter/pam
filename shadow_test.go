package pam

import (
	"encoding/json"
	"log"
	"os/user"
	"runtime"
	"testing"
)

func TestGetPasswordInfo(t *testing.T) {
	u, _ := user.Current()
	if u.Uid != "0" {
		t.Skip("run this test as root")
	}
	infos, err := GetPasswordInfo("test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	encoded, err := json.Marshal(infos)

	log.Println(infos)
	log.Println(err)
	log.Println(string(encoded))
	runtime.GC()
}

func TestChangePassword(t *testing.T) {
	u, _ := user.Current()
	if u.Uid != "0" {
		t.Skip("run this test as root")
	}
	err := ChangePassword("test", "blip")
	if err != nil {
		t.Fatalf(err.Error())
	}
	runtime.GC()
}
