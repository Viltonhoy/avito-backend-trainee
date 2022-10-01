package generated

import (
	"avito-backend-trainee/internal/storage"

	"github.com/shopspring/decimal"
)

type NewAdRequest struct {
	Name        string
	Description string
	Links       []string
	Price       decimal.Decimal
}

type NewAdResponse struct {
	Id   int64
	Code int64 ``
}

type GetAdRequest struct {
	Id     int64
	Fields []string
}

type GetAdResponse = storage.GetAdResponse

type AdListRequest struct {
	Order      *storage.OrdBy
	Offset     int64
	Pagination *string
}

type AdListResponse []GetAdResponse
