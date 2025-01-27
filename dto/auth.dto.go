package dto

type LoginRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponseDTO struct {
	AccessToken    string `json:"accessToken"`
	RefreshToken   string `json:"refreshToken"`
	WarehouseId    string `json:"warehouseId"`
	OrganizationId string `json:"organizationId"`
	Type           string `json:"type"`
}
