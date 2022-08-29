package user

type UserDTO struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Occupation string `json:"occupation"`
	Token      string `json:"token"`
	ImageURL   string `json:"image_url"`
}

func FormatUser(token string, user User) UserDTO {
	return UserDTO{
		Name:       user.Name,
		Email:      user.Email,
		ID:         user.ID,
		Occupation: user.Occupation,
		Token:      token,
		ImageURL:   user.AvatarFileName,
	}
}
