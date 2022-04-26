package entities

type UserEntity struct {
	Id      uint32 `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Phone   string `json:"phone"`
}
