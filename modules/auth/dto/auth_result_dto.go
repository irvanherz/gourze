package dto

type AuthResultDto struct {
	AccessToken          string `json:"accessToken"`
	RefreshToken         string `json:"refreshToken"`
	AccessTokenExpiredAt int64  `json:"accessTokenExpiredAt"`
}
