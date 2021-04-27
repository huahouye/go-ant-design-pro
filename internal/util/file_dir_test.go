package util

import (
	"log"
	"testing"
)

// go test h-trader/internal/util
func TestCreateFileWithDir(t *testing.T) {

	str1 := "./log/abc/ddd"
	str2 := "./log/xyz/ddd.txt"
	str3 := "/Users/houyehua/Desktop/.tmp/xyz/ddd.txt"
	log.Println(str1)
	CreateFileWithDirs(str1)
	CreateFileWithDirs(str2)
	CreateFileWithDirs(str3)
}
