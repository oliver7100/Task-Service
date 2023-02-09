package database

type Image struct {
	Path   string
	TaskID int
}

type Task struct {
	ID          int `gorm:"primaryKey"`
	Email       string
	Images      []Image
	Description string
}
