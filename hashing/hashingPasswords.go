package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	hash1, _ := bcrypt.GenerateFromPassword([]byte("deuterium"), 10)
	hash2, _ := bcrypt.GenerateFromPassword([]byte("ForTheEmperor!"), 10)
	hash3, _ := bcrypt.GenerateFromPassword([]byte("nuke"), 10)
	fmt.Printf("Hash for 'deuterium': '%v'\n", string(hash1))
	fmt.Printf("Hash for 'ForTheEmperor!': '%v'\n", string(hash2))
	fmt.Printf("Hash for 'nuke': '%v'\n", string(hash3))
}
