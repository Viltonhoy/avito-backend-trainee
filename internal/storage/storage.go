package storage

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/log/zapadapter"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type Storage struct {
	logger *zap.Logger
	db     *pgxpool.Pool
}

func NewStorage(ctx context.Context, logger *zap.Logger) (*Storage, error) {
	if logger == nil {
		return nil, errors.New("no logger provided")
	}

	config, _ := pgxpool.ParseConfig("")
	config.ConnConfig.Logger = zapadapter.NewLogger(logger)
	config.ConnConfig.LogLevel = pgx.LogLevelError

	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		logger.Error("error database connection", zap.Error(err))
	}

	err = pool.Ping(ctx)
	if err != nil {
		logger.Error("connection was not established", zap.Error(err))
		return &Storage{}, err
	}

	return &Storage{
		logger: logger,
		db:     pool,
	}, err
}

func (s *Storage) Close() {
	s.logger.Info("clothing Storage connection")
	s.db.Close()
}

func (s *Storage) AddNewAd(ctx context.Context, name string, description string, links []string, price decimal.Decimal) (int64, error) {
	logger := s.logger.With()
	logger.Debug("add new ad to table")

	var id int64

	err := s.db.QueryRow(
		ctx,
		firstIsertExec,
		name,
		description,
		links,
		price,
	).Scan(id)
	if err != nil {
		//
		return 0, err
	}
	return id, err
}

func (s *Storage) ReadAllAds(ctx context.Context, order *OrdBy, offset int64, gradation PaginationFields) ([]GetAdResponse, error) {
	logger := s.logger.With()
	logger.Debug("")

	var newQuery string
	var newRows pgx.Rows

	if order != nil {
		switch gradation {
		case Ascending:
			newQuery = readTableOrder
		case Descending:
			newQuery = readTableOrderDesc
		}
		rows, err := s.db.Query(
			ctx,
			newQuery,
			*order,
			Limit,
			offset,
		)
		if err != nil {
			//
			return nil, err
		}
		newRows = rows
	} else {
		newQuery = readTable
		rows, err := s.db.Query(
			ctx,
			newQuery,
			Limit,
			offset,
		)
		if err != nil {
			//
			return nil, err
		}
		newRows = rows
	}

	var dd []GetAdResponse
	for newRows.Next() {
		var d GetAdResponse
		err := newRows.Scan(&d.Name, &d.Links, &d.Price)
		if err != nil {
			//
			return nil, err
		}
		dd = append(dd, d)
	}
	return dd, nil
}

func (s *Storage) ReadAd(ctx context.Context, id int64, fields ...AdditionalFields) (GetAdResponse, error) {
	logger := s.logger.With()
	logger.Debug("")

	var userAd GetAdResponse

	var readTable string

	var total AdditionalFields

	for _, field := range fields {
		total = total | field
	}

	switch total {
	case total & 0:
		readTable = readNoFields
	case total & Description:
		readTable = readNoLinks
	case total & Links:
		readTable = readNoDescription
	case total & (Description | Links):
		readTable = readAllFields
	}

	err := s.db.QueryRow(
		ctx,
		readTable,
		id,
	).Scan(&userAd.Name, &userAd.Description, &userAd.Links, &userAd.Price)
	if err != nil {
		//
		return GetAdResponse{}, err
	}

	return userAd, nil
}
