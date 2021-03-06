package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func getGravatarHash(email string) string {
	email = strings.TrimSpace(email)
	email = strings.ToLower(email)

	h := md5.New()
	io.WriteString(h, email)
	finalBytes := h.Sum(nil)
	finalString := hex.EncodeToString(finalBytes)
	return finalString
}

func main() {
	gravatarHash := getGravatarHash("master@hackers-lab.org")
	fmt.Fprint(os.Stderr, "Enter your name: ")
	var name string
	fmt.Scanln(&name)

	
	fmt.Println(`<!DOCTYPE html>
<html>
<head>
</head>
<body>
<img src="http://www.gravatar.com/avatar/` + gravatarHash  + `?d=identicon">
</body>
</html>`)
}
