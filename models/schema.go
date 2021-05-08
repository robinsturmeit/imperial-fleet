package model

type Spacecraft struct {
	Id       int        `db:"id", json:"id"`
	Name     string     `db:"name", json:"name"`
	Class    *string    `db:"class", json:"class,omitempty"`
	Armament []Armament `json:"armament,omitempty"`
	Crew     *int       `db:"crew", json:"crew,omitempty"`
	Image    *string    `db:"image", json:"image,omitempty"`
	Value    *float32   `db:"value", json:"value,omitempty"`
	Status   string     `db:"status", json:"status"`
}

type Armament struct {
	Name     string `db:"name",json:"title"`
	Quantity int    `db:"quantity",json:"quantity"`
}

type Weapon struct {
	Id   int    `db:"id"`
	Name string `db:"name"`
}
