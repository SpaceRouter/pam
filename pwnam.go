package pam

//#include <sys/types.h>
//#include <stdlib.h>
//#include <pwd.h>
//#cgo CFLAGS: -Wall -std=c99
//#cgo LDFLAGS: -lpam
import "C"
import (
	"fmt"
	"log"
	"unsafe"
)

type UserInfo struct {
	Username        string /* username */
	UserId          uint   /* user ID */
	GroupId         uint   /* group ID */
	UserInformation string /* user information */
	HomeDirectory   string /* home directory */
	ShellProgram    string /* shell program */
}

func GetUserInfos(userId string) (*UserInfo, error) {
	user := C.CString(userId)
	defer C.free(unsafe.Pointer(user))

	pwnam := C.getpwnam(user)
	log.Println(pwnam)

	if pwnam == nil {
		return nil, fmt.Errorf("unable to reach user info")
	}
	infos := UserInfo{
		C.GoString(pwnam.pw_name),
		uint(C.uint(pwnam.pw_uid)),
		uint(C.uint(pwnam.pw_gid)),
		C.GoString(pwnam.pw_gecos),
		C.GoString(pwnam.pw_dir),
		C.GoString(pwnam.pw_shell),
	}
	return &infos, nil
}
