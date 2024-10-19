package user

import "time"

type User struct {
	ID        int       `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"varchar(100) null"`
	Email     string    `json:"email" gorm:"varchar(100) not null;unique"`
	Password  string    `json:"password" gorm:"varchar(255) not null"`
	Role      string    `json:"role" gorm:"varchar(10)"`
	Address   string    `json:"address" gorm:"varchar(150) null"`
	Phone     string    `json:"phone" gorm:"varchar(15) null"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	CreatedBy string    `json:"created_by" gorm:"default:system"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
	UpdatedBy string    `json:"updated_by" gorm:"default:system"`
}
