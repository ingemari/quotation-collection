package mapper

import (
	"quotation-collection/internal/handler/dto"
	"quotation-collection/internal/model"
)

func CreateReqToQuote(req dto.CreateQuoteReq) model.Quote {
	return model.Quote{
		Author: req.Author,
		Quote:  req.Quote,
	}
}

func QuoteToQuoteResp(q model.Quote) dto.QuoteResp {
	return dto.QuoteResp{
		ID:      q.ID.String(),
		Author:  q.Author,
		Quote:   q.Quote,
		Created: q.CreatedAt.String(),
	}
}

func QuotesToListResp(quotes []model.Quote) []dto.QuoteResp {
	var resp []dto.QuoteResp
	for _, q := range quotes {
		resp = append(resp, QuoteToQuoteResp(q))
	}
	return resp
}
