package weserver

import (
	"fmt"

	"github.com/qiniu/api.v7/auth/qbox"
	"github.com/qiniu/api.v7/storage"
)

var (
	accessKey = "1KuZx_CTOWU3HNIu_KNuTebM2kVRwp-BOXThqFgq"
	secretKey = "q61BzgBvSOtLiXWdpCJwBnTKYI54e9rZH8Lext-p"
	bucket    = "mainsite"
)

func getUpToken(keyToOverwrite string) string {
	fmt.Println(keyToOverwrite)
	putPolicy := storage.PutPolicy{
		// Scope: fmt.Sprintf("%s:%s", bucket, keyToOverwrite),
		Scope: bucket,
	}
	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)
	return upToken
}
