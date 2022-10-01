package server

import (
	"avito-backend-trainee/internal/storage"
	"context"

	"github.com/shopspring/decimal"
)

type Storager interface {
	AddNewAd(ctx context.Context, name string, description string, links []string, price decimal.Decimal) (int64, error)
	ReadAllAds(ctx context.Context, order *storage.OrdBy, offset int64, gradation storage.PaginationFields) ([]storage.GetAdResponse, error)
	ReadAd(ctx context.Context, id int64, fields ...storage.AdditionalFields) (storage.GetAdResponse, error)
}
