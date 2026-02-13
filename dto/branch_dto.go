package dto

type CreateBranchRequest struct {
	Name    string `json:"name" binding:"required"`
	IFSC    string `json:"ifsc" binding:"required"`
	Address string `json:"address"`
	BankID  uint   `json:"bank_id" binding:"required"`
}
