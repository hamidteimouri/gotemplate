package common

type GetCoinsRequest struct {
	Symbol string `json:"symbol" query:"symbol"`
}
