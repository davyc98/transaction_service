package request

type TransactionRequest struct {
	Page        int    `json:"page"  form:"page"   query:"page"   validate:"omitempty,number"  example:"1"`  // This is page, if load all data set value to 0
	Limit       int    `json:"limit" form:"limit"  query:"limit"  validate:"omitempty,number"  example:"10"` // This is limit, if load all data set value to -1
	SelectMonth string `json:"select_month"  form:"select_month"   query:"select_month" example:"2022-11"`   // Month 1 - 12
}
