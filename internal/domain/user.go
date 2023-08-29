package domain

type User struct {
	ID   string   `json:"id" gorm:"primary_key"`
	Data []string `json:"data"`
}

type Segment struct {
	Name    string `json:"name" gorm:"primary_key"`
	Percent int    `json:"percent"`
}
