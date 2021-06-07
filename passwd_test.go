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
	defer runtime.GC()

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
}

func TestChangeUserName001(t *testing.T) {
	defer runtime.GC()

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

}

func TestChangeUserEmail001(t *testing.T) {
	defer runtime.GC()

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

}

func TestListUsers001(t *testing.T) {
	defer runtime.GC()

	u, _ := user.Current()
	if u.Uid != "0" {
		t.Skip("run this test as root")
	}

	_, err := ListUsers()
	if err != nil {
		t.Fatalf(err.Error())
	}

}

func TestAddGroup001(t *testing.T) {
	defer runtime.GC()

	u, _ := user.Current()
	if u.Uid != "0" {
		t.Skip("run this test as root")
	}

	err := AddGroup("test", "sudo")
	if err != nil {
		t.Fatalf(err.Error())
	}

	userLookup, err := user.Lookup("test")
	if err != nil {
		t.Fatalf(err.Error())
	}

	group, err := user.LookupGroup("sudo")
	if err != nil {
		t.Fatalf(err.Error())
	}

	userGroups, err := userLookup.GroupIds()
	if err != nil {
		t.Fatalf(err.Error())
	}

	for _, userGroup := range userGroups {
		if userGroup == group.Gid {
			return
		}
	}

	t.Fail()

}

func TestChangeGroup001(t *testing.T) {
	defer runtime.GC()

	u, _ := user.Current()
	if u.Uid != "0" {
		t.Skip("run this test as root")
	}

	err := ChangeGroups("test", []string{
		"test",
		"louis",
	})
	if err != nil {
		t.Fatalf("changegroup %s", err.Error())
	}

	userLookup, err := user.Lookup("test")
	if err != nil {
		t.Fatalf("userLookup %s", err.Error())
	}

	group, err := user.LookupGroup("louis")
	if err != nil {
		t.Fatalf("groupLoockup %s", err.Error())
	}

	userGroups, err := userLookup.GroupIds()
	if err != nil {
		t.Fatalf("userLookup.GroupIds() : %s", err.Error())
	}

	for _, userGroup := range userGroups {
		if userGroup == group.Gid {
			return
		}
	}

	t.Fatalf("Group did not change")

}
