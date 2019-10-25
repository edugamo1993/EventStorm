package authentication

import "encoding/json"

/*User Type for user data*/
type User struct {
	ID          string `json:"_id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Birthday    string `json:"birthday"`
	Country     string `json:"country"`
	City        string `json:"city"`
	CP          string `json:"cp"`
	PhoneNumber string `json:"phone,omitempty"`
	Device      string `json:"device,omitempty"`
	FirstName   string `json:"firstName,omitempty"`
	LastName    string `json:"lastName,omitempty"`
	Verified    bool
}

//NewUser initializates var type User
func (u *User) NewUser(jsonData []byte) {
	json.Unmarshal(jsonData, &u)
	u.ID = u.Username

}

//Verificate function set user as verified
func (u *User) Verificate() {
	u.Verified = true
}
