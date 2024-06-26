package utils

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// GetMd5String 生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// UniqueId 生成Guid字串
//func UniqueId() string {
//	b := make([]byte, 48)
//	if _, err := io.ReadFull(rand.Reader, b); err != nil {
//		return ""
//	}
//	return GetMd5String(base64.URLEncoding.EncodeToString(b))
//}

// 生成随机字符的函数
func generateRandomDigits(n int) string {
	rand.Seed(time.Now().UnixNano())
	var sb strings.Builder
	for i := 0; i < n; i++ {
		sb.WriteString(strconv.Itoa(rand.Intn(10))) // 生成0到9的随机整数
	}
	return sb.String()
}

// 生成带日期时间前缀的序列号
func UniqueId() string {
	now := time.Now()
	datePrefix := now.Format("20060102150405")                 // 格式化日期时间为"yyyyMMddHHmmss"格式
	randomDigits := generateRandomDigits(16 - len(datePrefix)) // 剩下的位数用随机整数字符填充
	serialNumber := datePrefix + randomDigits
	return serialNumber
}
