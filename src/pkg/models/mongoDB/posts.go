package mongoDB

import (

	// "src/pkg/models"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/BurntSushi/toml"
	"github.com/dgrijalva/jwt-go"

	"src/pkg/models"
)

//PostModel is a struct
type PostModel struct{}

//Config is struct
// type Config struct {
// 	Serial string
// }

//Insert is a function
func (m *PostModel) Insert(p *models.Post) error {
	// id := "1"

	content, _ := ioutil.ReadFile("/secrets.toml")
	var config Config
	if _, err := toml.Decode(string(content), &config); err != nil {
		return err
	}

	var SERIAL = []byte(config.Serial)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"type":    p.Type,
		"title":   p.Title,
		"user_id": p.UserId,
		"content": p.Content,
		"tags":    [...]int{},
		// "tags": p.Tags,
	})
	// fmt.Println(token)

	tokenString, err := token.SignedString(SERIAL)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "http://api:8000/new_post", bytes.NewBuffer([]byte(tokenString)))
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
	body, _ := ioutil.ReadAll(resp.Body)
	// fmt.Println("response Body:", string(body))

	var date map[string]interface{}
	json.Unmarshal(body, &date)

	if resp.StatusCode == 400 || resp.StatusCode == 500 {
		// fmt.Println("is error")
		return models.ErrBadRequest

	}

	// collection := m.DB.Database("searchandfind").Collection("posts")
	// insertResult, err := collection.InsertOne(context.TODO(), p)
	// if err != nil {
	// 	return "", err
	// }
	// id := insertResult.InsertedID.(primitive.ObjectID).Hex()

	return nil
}

//Get is a function
func (m *PostModel) Get(f string) (*models.Post, error) {
	var post *models.Post = new(models.Post)

	content, _ := ioutil.ReadFile("/secrets.toml")
	var config Config
	if _, err := toml.Decode(string(content), &config); err != nil {
		return nil, err
	}

	var SERIAL = []byte(config.Serial)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	// fmt.Println(token)

	tokenString, err := token.SignedString(SERIAL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "http://api:8000/posts/"+f, bytes.NewBuffer([]byte(tokenString)))
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

	// fmt.Println(date["posts"].([]interface{}))
	v := date["posts"].([]interface{})
	// fmt.Println("start")
	vv := v[0].(map[string]interface{})
	// fmt.Println(vv["content"])

	post.ID = int(vv["id"].(float64))
	userID := vv["user_id"].(float64)
	userIDstr := strconv.FormatFloat(userID, 'f', -1, 64)
	post.UserId = userIDstr
	post.Type = vv["type"].(string)
	post.Title = vv["title"].(string)
	post.Content = vv["content"].(string)

	// fmt.Println(post)

	// fmt.Println(v["c"].(string))

	// filter := bson.D{{"_id", "ObjectId(\"" + f + "\")"}}
	// fmt.Println("ObjectId(\"" + f + "\")")
	// objID, _ := primitive.ObjectIDFromHex(f)
	// collection := m.DB.Database("searchandfind").Collection("posts")

	// err := collection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&post)
	// // fmt.Println(post.ID, "ok")
	// if err != nil {
	// 	return nil, err
	// }
	return post, nil
}

//Latest is a function
func (m *PostModel) Latest() ([]*models.Post, error) {
	var posts []*models.Post

	content, _ := ioutil.ReadFile("/secrets.toml")
	var config Config
	if _, err := toml.Decode(string(content), &config); err != nil {
		return nil, err
	}

	var SERIAL = []byte(config.Serial)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{})
	// fmt.Println(token)

	tokenString, err := token.SignedString(SERIAL)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", "http://api:8000/posts", bytes.NewBuffer([]byte(tokenString)))
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

	// fmt.Println(date["posts"])

	// fmt.Println(len(date["posts"].([]interface{})))

	for _, v := range date["posts"].([]interface{}) {
		var post *models.Post = new(models.Post)
		// fmt.Println("1")
		vv := v.(map[string]interface{})
		// fmt.Println("2")

		// post.ID = int(vv["id"].(float64))
		// fmt.Println(vv["type"].(string))
		// var typeStr string
		// typeStr = vv["type"].(string)
		userID := vv["user_id"].(float64)
		userIDstr := strconv.FormatFloat(userID, 'f', -1, 64)

		// fmt.Println(typeStr)
		// typeStr.Decode(&post.Type)

		// post.Type = string(typeStr)
		// fmt.Println("1")
		// i, err: = strconv.ParseInt ("- 42", 10, 64)
		post.ID = int(vv["id"].(float64))
		post.Type = string(vv["type"].(string))
		post.UserId = string(userIDstr)
		post.Title = string(vv["title"].(string))
		post.Content = string(vv["content"].(string))
		// post.Tags = vv["tags"].(float64)
		// fmt.Println("3")

		// err := v.Decode(&post)
		// fmt.Println(v)
		// fmt.Println("start")
		// fmt.Println(v.(map[string]interface{})["content"])
		// json := `{
		// 	"type":"` + typeStr + `",
		// 	"userId":"` + userIDstr + `",
		// 	"title":"` + userIDstr + `",
		// }`
		// fmt.Println(json)
		// post.ID = 12

		// json.(&post)
		// json.Unmarshal([]byte(json), post)
		posts = append(posts, post)
		// fmt.Println(post)
		// fmt.Println(posts)
	}

	// collection := m.DB.Database("searchandfind").Collection("posts")

	// cur, err := collection.Find(context.TODO(), bson.D{})
	// if err != nil {
	// 	return nil, err
	// }

	// for cur.Next(context.TODO()) {
	// 	var post *models.Post

	// 	err := cur.Decode(&post)
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	posts = append(posts, post)
	// }

	// if err := cur.Err(); err != nil {
	// 	return nil, err
	// }

	// // cur.Close(context.TODO())
	// fmt.Println("here")
	// fmt.Println(posts)
	return posts, nil
}

