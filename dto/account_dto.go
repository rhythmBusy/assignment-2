package dto

type CreateAccountRequest struct {
	AccountNumber string `json:"account_number" binding:"required"`
	UserIDs       []uint `json:"user_ids" binding:"required"`
	BranchID      uint   `json:"branch_id" binding:"required"`
}

type AmountRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}
