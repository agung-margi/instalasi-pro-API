package user

import "time"

type RegisterFormatter struct {
	Email string `json:"email"`
	Role  string `json:"role"`
}

func FormatRegister(user User) RegisterFormatter {
	formatter := RegisterFormatter{
		Email: user.Email,
		Role:  user.Role,
	}
	return formatter
}

type UserFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
		Address:   user.Address,
		Phone:     user.Phone,
		UpdatedAt: user.UpdatedAt,
	}
	return formatter
}

type UsersFormatter struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Role      string    `json:"role"`
	Address   string    `json:"address"`
	Phone     string    `json:"phone"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FormatUsers(users []User) []UsersFormatter {
	if len(users) == 0 {
		return []UsersFormatter{}
	}
	usersFormatter := []UsersFormatter{}
	for _, user := range users {
		userFormatter := UsersFormatter{
			ID:        user.ID,
			Name:      user.Name,
			Email:     user.Email,
			Role:      user.Role,
			Address:   user.Address,
			Phone:     user.Phone,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		}
		usersFormatter = append(usersFormatter, userFormatter)
	}
	return usersFormatter
}
