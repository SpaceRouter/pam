package pam

//#include <sys/types.h>
//#include <stdlib.h>
//#include <pwd.h>
//extern struct passwd *getpwent (void);
//extern void endpwent (void);
//#cgo CFLAGS: -Wall -std=c99 -ansi
//#cgo LDFLAGS: -lpam
import "C"
import (
	"fmt"
	"os/exec"
	"strings"
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

	pwnam, err := C.getpwnam(user)
	if err != nil {
		return nil, err
	}

	if pwnam == nil {
		return nil, fmt.Errorf("unable to reach user info")
	}
	return passwdToUserInfo(pwnam), nil
}

func ChangeUserName(userId string, userName string) error {
	cmd := exec.Command("chfn", userId, "-f", userName)
	return execCommand(cmd)
}

func ChangeUserEmail(userId string, email string) error {
	cmd := exec.Command("chfn", userId, "-o", email)
	return execCommand(cmd)
}

func AddGroup(username string, group string) error {
	cmd := exec.Command("usermod", "-aG", group, username)
	return execCommand(cmd)
}

func ChangeGroups(username string, groups []string) error {

	args := []string{
		"-G",
	}
	args = append(args, strings.Join(groups, ","))
	args = append(args, username)

	cmd := exec.Command("usermod", args...)
	return execCommand(cmd)
}

func ListUsers() ([]UserInfo, error) {
	var users []UserInfo
	defer C.endpwent()

	for {
		pwnam := C.getpwent()
		if pwnam == nil {
			return users, nil
		}
		users = append(users, *passwdToUserInfo(pwnam))
	}
}

func passwdToUserInfo(passwd *C.struct_passwd) *UserInfo {
	infos := UserInfo{
		C.GoString(passwd.pw_name),
		uint(C.uint(passwd.pw_uid)),
		uint(C.uint(passwd.pw_gid)),
		C.GoString(passwd.pw_gecos),
		C.GoString(passwd.pw_dir),
		C.GoString(passwd.pw_shell),
	}
	return &infos
}
