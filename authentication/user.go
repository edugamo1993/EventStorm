package authentication

import (
	"EventStorm/config"
	"encoding/json"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
)

/*User Type for user data*/
type User struct {
	ID          string `json:"_id" bson:"_id"`
	Username    string `json:"username" bson:"username"`
	Password    string `json:"password" bson:"password"`
	Birthday    string `json:"birthday" bson:"birthday"`
	Country     string `json:"country" bson:"country"`
	City        string `json:"city" bson:"city"`
	CP          string `json:"cp,omitempty" bson:"cp"`
	PhoneNumber string `json:"phone,omitempty" bson:"phone"`
	Device      string `json:"device,omitempty" bson:"device"`
	FirstName   string `json:"firstName,omitempty" bson:"firstName"`
	LastName    string `json:"lastName,omitempty" bson:"lastName"`
	Private     bool   `json:"private" bson:"private"`
	Verified    bool   `bson:"private"`
}

//NewUser initializates var type User
func NewUser(c *config.Config, jsonData []byte) (u *User, err error) {
	var id interface{}
	bjsonData := bson.D{}
	var marshaldata []byte
	json.Unmarshal(jsonData, &u)
	u.ID = u.Username
	if err != nil {
		return nil, err
	}
	marshaldata, err = bson.Marshal(u)
	fmt.Println(string(marshaldata))
	err = bson.Unmarshal(marshaldata, &bjsonData)
	id, err = c.Mongo.InsertData(c, "users", bjsonData)
	if err != nil {
		return nil, err
	}
	fmt.Println(id)
	return u, err
}

//GetUserByID returns type User with the user info.
func (u *User) GetUserByID(c *config.Config, username string) (jsonResult []byte, err error) {
	s := fmt.Sprintf("{_id: %s}", username)
	query := []byte(s)
	result, err := c.Mongo.FindOne(c, query, "users")
	jsonResult, err = json.Marshal(result)
	if err != nil {
		return nil, err
	}
	// err = json.Unmarshal(jsonResult, u)
	return
}

//Verificate function set user as verified
func (u *User) Verificate() {
	u.Verified = true
}
