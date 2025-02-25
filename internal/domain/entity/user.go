package entity

import (
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name      string
	Surname   string
	FullName  string
	BirthDate time.Time
}
