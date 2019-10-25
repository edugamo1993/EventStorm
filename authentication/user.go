package authentication

import "encoding/json"

/*User Type for user data*/
type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Birthday    string `json:"birthday"`
	Country     string `json:"country"`
	CP          string `json:"cp"`
	PhoneNumber string `json:"phone"`
	Device      string `json:"device"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Verified    bool
}

//NewUser initializates var type User
func (u *User) NewUser(jsonData []byte) {
	json.Unmarshal(jsonData, &u)
}

//Verificate function set user as verified
func (u *User) Verificate() {
	u.Verified = true
}
