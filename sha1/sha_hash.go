package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"crypto/rand"
	"math/big"
	"os"
    "bufio"
)

const (
	saltMinLength = 10
	saltMaxLength = 20
)

func generateSalt(minLength, maxLength int64) (string, error) {
	randNumber, err := rand.Int(rand.Reader,big.NewInt(maxLength - minLength))
	if err != nil {
		return "", err
	}
	saltLength := randNumber.Int64() + minLength
	salt := make([]byte, saltLength)
	_, err = rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

func getPassword(password string) string {
	salt, _ := generateSalt(saltMinLength, saltMaxLength)
	passwordHash := SHA1(password + salt)
	return fmt.Sprintf("%s|%s", passwordHash, salt)
}

func SHA1(str string) string{
	c:=sha1.New()
	c.Write([]byte(str))
	return hex.EncodeToString(c.Sum(nil))
}

func writeLines(pws []string, path string) error {
    file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
	    panic(err)
	}

	defer file.Close()
    w := bufio.NewWriter(file)
	for _, line := range pws {
    	fmt.Fprintln(w, line)
	}	
   
    return w.Flush()
}

func main() {
	nAcc := 40000
	var pws []string
	for i:=1; i<=nAcc; i++ {
		pass := getPassword("123456")
		pws = append(pws, pass)
	}
	_ = writeLines(pws, "./password.txt")
}