package mongoDB

import (
	"bytes"
	"io/ioutil"

	// "context"

	// "fmt"
	"net/http"
	"src/pkg/models"
	// "json"
	"encoding/json"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"
	// "golang.org/x/crypto/bcrypt"
	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
)

var apiHost = "api:8000"

//UserModel is struct
type UserModel struct{}

//M is type of data
// type M map[string]interface{}

//Config is struct
type Config struct {
	Serial string
}

//Insert data to our api
func (m *UserModel) Insert(u *models.User, password string) error {

	content, _ := ioutil.ReadFile("/secrets.toml")
	var config Config
	if _, err := toml.Decode(string(content), &config); err != nil {
		return err
	}

	var SERIAL = []byte(config.Serial)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		// "name":     u.Name,
		// "surname":  u.Surname,
		"nick":     u.Nick,
		"email":    u.Email,
		"password": password,
	})
	// fmt.Println(token)

	tokenString, err := token.SignedString(SERIAL)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://api:8000/new_user", bytes.NewBuffer([]byte(tokenString)))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	// body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	return nil

}

//Authenticate is a function
func (m *UserModel) Authenticate(email, password string) (interface{}, error) {

	content, _ := ioutil.ReadFile("/secrets.toml")
	var config Config
	if _, err := toml.Decode(string(content), &config); err != nil {
		return nil, err
	}

	var SERIAL = []byte(config.Serial)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":    email,
		"password": password,
	})
	// fmt.Println(token)

	tokenString, err := token.SignedString(SERIAL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "http://api:8000/login", bytes.NewBuffer([]byte(tokenString)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	var date map[string]interface{}
	json.Unmarshal(body, &date)

	if resp.StatusCode == 400 {

		return nil, models.ErrInvalidCredentials
	}
	// var id string = string(date["id"].(string))
	// id := fmt.Sprintf("%f", date["id"].(float64))
	str := date["id"].(float64)
	id := strconv.FormatFloat(str, 'f', -1, 64)
	// id := date["id"].(float64)
	// fmt.Println(str)

	return id, nil
}

//Get is a function
func (m *UserModel) Get(u string) (*models.User, error) {
	s := &models.User{}

	content, _ := ioutil.ReadFile("/secrets.toml")
	var config Config
	if _, err := toml.Decode(string(content), &config); err != nil {
		return nil, err
	}

	var SERIAL = []byte(config.Serial)
	// fmt.Println(passwordHashed)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	// fmt.Println("Logowanie hashed")
	// fmt.Println(email)
	// fmt.Println(password)

	tokenString, err := token.SignedString(SERIAL)
	if err != nil {
		return nil, err
	}

	// fmt.Println(u)
	var url string = "http://api:8000/users/" + u
	// var url string = "http://api:8000/users/1004"
	// fmt.Println(url)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(tokenString)))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// fmt.Println("response Status:", resp.Status)
	// fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))
	// fmt.Println(string(body[1]))

	date := make(map[string]interface{})

	json.Unmarshal(body, &date)

	if resp.StatusCode == 400 {

		return nil, models.ErrNoRecord
	}

	str := date["users"].([]interface{})
	str1 := str[0].(map[string]interface{})

	str4 := str1["id"].(float64)

	s.ID = int(str4)

	if str1["name"] != nil {
		s.Name = str1["name"].(string)
	}
	if str1["name"] != nil {
		s.Surname = str1["surname"].(string)
	}
	s.Nick = str1["nick"].(string)
	// fmt.Println(str1["nick"])

	s.Email = str1["email"].(string)

	return s, nil
}

//Exists is a function
func (m *UserModel) Exists(email string) error {
	// filter := bson.D{{"email", email}}
	// var user models.User
	// // fmt.Println("OK")
	// // fmt.Println(filter)

	// collection := m.DB.Database("searchandfind").Collection("users")
	// err := collection.FindOne(context.TODO(), filter).Decode(&user)
	// // fmt.Println(err)
	// // fmt.Println(user.Name)
	// if err != nil {
	// 	return nil
	// }
	// return errors.New("Is not empty")
	return nil

}
