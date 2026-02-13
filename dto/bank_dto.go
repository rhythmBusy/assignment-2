package dto

type CreateBankRequest struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}

type UpdateBankRequest struct {
	Name string `json:"name"`
	Code string `json:"code"`
}
