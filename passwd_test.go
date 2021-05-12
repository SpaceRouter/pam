package pam

import (
	"encoding/json"
	"log"
	"os/user"
	"runtime"
	"strings"
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

func TestChangeUserName001(t *testing.T) {
	u, _ := user.Current()
	if u.Uid != "0" {
		t.Skip("run this test as root")
	}
	err := ChangeUserName("test", "Test qdqzdqz")
	if err != nil {
		t.Fatalf(err.Error())
	}

	infos, err := GetUserInfos("test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !strings.Contains(infos.UserInformation, "Test qdqzdqz") {
		t.Fatalf("Can't change user name")
	}

	runtime.GC()
}

func TestChangeUserEmail001(t *testing.T) {
	u, _ := user.Current()
	if u.Uid != "0" {
		t.Skip("run this test as root")
	}
	err := ChangeUserEmail("test", "test@test.test")
	if err != nil {
		t.Fatalf(err.Error())
	}

	infos, err := GetUserInfos("test")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if !strings.Contains(infos.UserInformation, "test@test.test") {
		t.Fatalf("Can't change user name")
	}

	runtime.GC()
}
