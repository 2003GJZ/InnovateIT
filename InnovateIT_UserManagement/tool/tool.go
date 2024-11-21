package tool

//工具类
import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

func CompareMD5(password string, passwordMd5 string) bool { //判断密码是否正确

	hashString := GetMd5(password)
	return hashString == passwordMd5
}

func GetMd5(password string) string {
	// 计算MD5哈希值
	hasher := md5.New()
	hasher.Write([]byte(password))
	hashBytes := hasher.Sum(nil)
	hashString := hex.EncodeToString(hashBytes)
	return hashString
}

// 字符串分割，节点任务拆分
func SplitString(ags string, partition string) (string, string, error) { //返回截取字符串（去除分割符），和截取完的字符串
	index := strings.Index(ags, partition)
	if index == -1 {
		return ags, "", nil
	}
	return ags[:index], ags[index+len(partition):], nil
}
