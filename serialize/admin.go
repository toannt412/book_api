package serialize

type Admin struct {
	Id       string `json:"id"`
	FullName string `json:"fullname,omitempty"`
	Phone    string `json:"phone,omitempty"`
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}
