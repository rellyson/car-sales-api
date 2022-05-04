package entities

type Seller struct {
	BaseEntity
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
