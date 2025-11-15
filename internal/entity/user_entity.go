package entity

type UserEntity struct {
	ID      int `gorm:"primaryKey;autoIncrement"`
	Name    string `gorm:"type:varchar(100);not null;"`
	Hobbies string `gorm:"type:varchar(255);not null;"`
	CreatedAt string `gorm:"type:timestamp;default:current_timestamp"`
}