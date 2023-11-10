package main

import "fmt"

func main() {
	sha := &SHA{}
	md5 := &MD5{}

	protector := NewPasswordProtector("juank", "mypassword", sha)
	protector.Hash()
	protector.SetHashAlgorithm(md5)
	protector.Hash()
}

func NewPasswordProtector(user string, password string, algorithm Algorithm) PasswordProtector {
	return PasswordProtector{
		user:      user,
		password:  password,
		algorithm: algorithm,
	}
}

type PasswordProtector struct {
	user      string
	password  string
	algorithm Algorithm
}

func (pp *PasswordProtector) Hash() {
	pp.algorithm.Hash(pp.password)
}

func (pp *PasswordProtector) SetHashAlgorithm(algorithm Algorithm) {
	pp.algorithm = algorithm
}

type Algorithm interface {
	Hash(string)
}

type MD5 struct{}

func (md5 *MD5) Hash(text string) {
	fmt.Printf("Hashing with MD5 text:%s\n", text)
}

type SHA struct{}

func (sha *SHA) Hash(text string) {
	fmt.Printf("Hashing with SHA text:%s\n", text)
}
