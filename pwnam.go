package pam

//#include <sys/types.h>
//#include <stdlib.h>
//#include <pwd.h>
//#cgo CFLAGS: -Wall -std=c99
//#cgo LDFLAGS: -lpam
//int getpwnam_r(const char *name, struct passwd *pwd,
//            char *buf, size_t buflen, struct passwd **result);
import "C"
import (
	"log"
)

type UserInfo struct {
	Username        string /* username */
	Password        string /* user password */
	UserId          uint   /* user ID */
	GroupId         uint   /* group ID */
	UserInformation string /* user information */
	HomeDirectory   string /* home directory */
	ShellProgram    string /* shell program */
}

func GetUserInfos(userId string) UserInfo {
	user := C.CString(userId)

	pwnam := C.getpwnam(user)
	log.Println(pwnam)
	infos := UserInfo{
		C.GoString(pwnam.pw_name),
		C.GoString(pwnam.pw_passwd),
		uint(C.uint(pwnam.pw_uid)),
		uint(C.uint(pwnam.pw_gid)),
		C.GoString(pwnam.pw_gecos),
		C.GoString(pwnam.pw_dir),
		C.GoString(pwnam.pw_shell),
	}
	return infos
}
