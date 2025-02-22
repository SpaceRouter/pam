package pam

//#include <sys/types.h>
//#include <stdlib.h>
//#include <shadow.h>
//#include <unistd.h>
//#include <stdio.h>
//#include <errno.h>
//#include <string.h>
//#cgo CFLAGS: -Wall -std=c99
//#cgo LDFLAGS: -lpam
import "C"
import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"time"
	"unsafe"
)

type ShadowPassword struct {
	Login      string    /* Login name */
	Password   string    /* Encrypted password */
	LastChange time.Time /* Date of last change
	   (measured in days since
	   1970-01-01 00:00:00 +0000 (UTC)) */
	MinDaysBtwChanges int32 /* Min # of days between changes */
	MaxDaysBtwChanges int32 /* Max # of days between changes */
	DaysBeforeExpires int32 /* # of days before password expires
	o warn user to change it */
	DaysAfterExpires int32 /* # of days after password expires
	ntil account is disabled */
	ExpirationDate time.Time /* Date when account expires
	   (measured in days since
	   1970-01-01 00:00:00 +0000 (UTC)) */
	Flag uint32 /* Reserved */
}

func GetPasswordInfo(userId string) (*ShadowPassword, error) {
	user := C.CString(userId)
	defer C.free(unsafe.Pointer(user))

	spnam := C.getspnam(user)
	log.Println(spnam)

	if spnam == nil {
		return nil, fmt.Errorf("unable to reach user info")
	}
	infos := ShadowPassword{
		C.GoString(spnam.sp_namp),
		C.GoString(spnam.sp_pwdp),
		time.Unix(0, int64(C.long(spnam.sp_lstchg))*int64(time.Hour)*24),
		int32(C.long(spnam.sp_min)),
		int32(C.long(spnam.sp_max)),
		int32(C.long(spnam.sp_warn)),
		int32(C.long(spnam.sp_inact)),
		time.Unix(0, int64(C.long(spnam.sp_lstchg))*int64(time.Hour)*24),
		uint32(C.ulong(spnam.sp_flag)),
	}
	return &infos, nil
}

func ChangePassword(userId string, newPassword string) error {

	cmd := exec.Command("chpasswd")
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return err
	}
	stderr, err := cmd.StderrPipe()
	if err != nil {
		return err
	}

	err = cmd.Start()
	if err != nil {
		return err
	}

	_, err = io.WriteString(stdin, userId+":"+newPassword)
	if err != nil {
		return err
	}

	err = stdin.Close()
	if err != nil {
		return err
	}

	err = cmd.Wait()
	if err != nil {
		var bufStdout []byte
		_, errStdout := stdout.Read(bufStdout)

		var bufStderr []byte
		_, errStderr := stderr.Read(bufStderr)

		errorMessage := "Stdout error"
		if errStdout == nil {
			errorMessage = string(bufStdout)
		}

		if errStderr == nil {
			errorMessage += " " + string(bufStderr)
		} else {
			errorMessage += " Stderr error"
		}

		return fmt.Errorf("error %s \noutput: %s", err.Error(), errorMessage)
	}
	return nil
}
