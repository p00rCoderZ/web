package main

import (
	"io/ioutil"
	"net/http"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"bytes"
	"github.com/BurntSushi/toml"
)

type Config struct {
	Serial string
}

func main() {
	
	content, _ := ioutil.ReadFile("/secrets.toml")
	var config Config
	if _, err := toml.Decode(string(content), &config); err != nil {
		fmt.Println("error")
		fmt.Println(err)
	}
	fmt.Println(config.Serial)
	var SERIAL = []byte(config.Serial)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"testing": "if",
		"everything": "works",
		"properly": true,
	})

	tokenString, err := token.SignedString(SERIAL)
	if err != nil {
		fmt.Println("err")
		fmt.Println(err)
	}

	fmt.Println(tokenString)

	req, err := http.NewRequest("POST", "http://api:8000/jwt", bytes.NewBuffer([]byte(tokenString)))
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))
}