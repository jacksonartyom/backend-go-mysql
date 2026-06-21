package response

type UserResponse struct {
	UserId       string  `json:"userId"`
	Email        string  `json:"email"`
	FirstName    string  `json:"firstName"`
	MidName      *string `json:"midName"`
	LastName     string  `json:"lastName"`
	PhoneNo      string  `json:"phoneNo"`
	Token        string  `json:"token"`
	ImageProfile *string `json:"imageProfile"`
}
