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

func QuoteToCreateResp(q model.Quote) dto.CreateQuoteResp {
	return dto.CreateQuoteResp{
		ID:      q.ID.String(),
		Author:  q.Author,
		Quote:   q.Quote,
		Created: q.CreatedAt.String(),
	}
}
