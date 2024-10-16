package customer

type Customer struct {
	ID        string `json:"id" gorm:"primaryKey int"`
	Name      string `json:"name" gorm:"varchar(100) not null"`
	Email     string `json:"email" gorm:"varchar(100) not null email"`
	Address   string `json:"address" gorm:"varchar(150) not null"`
	Phone     string `json:"phone" gorm:"varchar(15) not null"`
	CreatedAt string `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	CreatedBy string `json:"created_by" gorm:"default:CURRENT_USER"`
	UpdatedAt string `json:"updated_at" gorm:"default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP"`
	UpdatedBy string `json:"updated_by" gorm:"default:CURRENT_USER"`
}
