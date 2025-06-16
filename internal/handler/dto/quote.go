package dto

type CreateQuoteReq struct {
	Author string `json:"author"`
	Quote  string `json:"quote"`
}

type CreateQuoteResp struct {
	ID      string `json:"id"`
	Author  string `json:"author"`
	Quote   string `json:"quote"`
	Created string `json:"created_at"`
}
