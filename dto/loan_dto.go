package dto

type CreateLoanRequest struct {
	UserID    uint    `json:"user_id" binding:"required"`
	BranchID  uint    `json:"branch_id" binding:"required"`
	Principal float64 `json:"principal" binding:"required,gt=0"`
}

type RepayLoanRequest struct {
	Amount float64 `json:"amount" binding:"required,gt=0"`
}
