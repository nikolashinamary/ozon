package Storages

type Model struct {
	ShortUrl string `gorm:"primaryKey"`
	Url      string
}
