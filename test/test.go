package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("123456"))
	cipherStr := h.Sum(nil)
	fmt.Println(cipherStr)
	fmt.Println(hex.EncodeToString(cipherStr))
}
