package dto

import "github.com/irvanherz/gourze/modules/user"

type AuthResultDto struct {
	AccessToken          string   `json:"accessToken"`
	RefreshToken         string   `json:"refreshToken"`
	AccessTokenExpiredAt int64    `json:"accessTokenExpiredAt"`
	User                 AuthUser `json:"user"`
}

type AuthUser struct {
	ID       uint          `json:"id"`
	Username string        `json:"username"`
	Email    string        `json:"email"`
	Role     user.UserRole `json:"role"`
	FullName string        `json:"fullName"`
}
