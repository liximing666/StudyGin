package Tables

type Info struct {
	Id uint `gorm:"primaryKey"`
	Desc string
}
