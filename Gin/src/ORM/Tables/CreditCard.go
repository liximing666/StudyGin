package Tables

import "github.com/jinzhu/gorm"

type CreditCard struct {
	gorm.Model
	CreditNum int
}
