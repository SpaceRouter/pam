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

	p := "blip"

	err := ChangePassword("test", p)
	if err != nil {
		t.Fatalf(err.Error())
	}

	tx, err := StartFunc("", "test", func(s Style, msg string) (string, error) {
		return p, nil
	})
	if err != nil {
		t.Fatalf("start #error: %v", err)
	}
	err = tx.Authenticate(0)
	if err != nil {
		t.Fatalf("authenticate #error: %v", err)
	}

	p = "secret"

	err = ChangePassword("test", p)
	if err != nil {
		t.Fatalf(err.Error())
	}

	tx, err = StartFunc("", "test", func(s Style, msg string) (string, error) {
		return p, nil
	})
	if err != nil {
		t.Fatalf("start #error: %v", err)
	}
	err = tx.Authenticate(0)
	if err != nil {
		t.Fatalf("authenticate #error: %v", err)
	}
	runtime.GC()
}
