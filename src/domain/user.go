package domain

type User struct {
	ID   string   `json:"id" gorm:"primary_key"`
	Data []string `json:"data"`
}
