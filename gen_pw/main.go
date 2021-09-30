package main

import (
	"fmt"
	"crypto/rand"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
	"math/big"
	"os"
    "bufio"
    "sync"
)

const (
	saltMinLength = 10
	saltMaxLength = 20
)

func generatePasswordHashSalt(password string) (string, error) {
	randNumber, err := rand.Int(rand.Reader,big.NewInt(saltMaxLength - saltMinLength))
	if err != nil {
		panic(err)
	}
	saltLength := randNumber.Int64() + saltMinLength
	salt := make([]byte, saltLength)
	_, err = rand.Read(salt)
	if err != nil {
		panic(err)
	}
	passwordHash, err := bcrypt.GenerateFromPassword(append([]byte(password), salt...), 14)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s|%s", hex.EncodeToString(passwordHash), hex.EncodeToString(salt)), nil
}

// writeLines writes the lines to the given file.
func writeLines(pw2d [][]string, path string) error {
    file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
	    panic(err)
	}

	defer file.Close()
    w := bufio.NewWriter(file)
    for _, lines := range pw2d {
    	for _, line := range lines {
        	fmt.Fprintln(w, line)
    	}	
    }
   
    return w.Flush()
}

func main() {
	// n_accounts := 10000000 + 100
	// n_accounts := 10
	nRoutines := 20
	nAcc := 10
	nTimes := 200
	var wg sync.WaitGroup
	pw2d := make([][]string, nRoutines)
	for t:=0; t<nTimes; t++ {
		for i:=0; i<nRoutines; i++ {
			wg.Add(1)
			go func(i int) {
				pw2d[i] = make([]string, nAcc)
				for j:=0; j<nAcc; j++ {
					s, _ := generatePasswordHashSalt("123456")
					pw2d[i][j] = s
				}
			    wg.Done()
			    fmt.Printf("go routine %d done\n", i)
			}(i)
		}

		wg.Wait()

		fmt.Printf("OK time %d\n", t)
		if err := writeLines(pw2d, "./password.txt"); err != nil {
	        fmt.Printf("writeLines: %d", err)
		}
	}

	fmt.Println("Done")
}