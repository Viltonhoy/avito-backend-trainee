package storage

import "github.com/shopspring/decimal"

const (
	firstIsertExec = "INSERT INTO adtable (ad_name, ad_description, photo_links, price) VALUES ($1, $2, $3, $4) returning id;"

	readAllFields     = "SELECT ad_name, ad_description, photo_links, price FROM adtable WHERE id = $1;"
	readNoDescription = "SELECT ad_name, NULL, photo_links, price FROM adtable WHERE id = $1;"
	readNoLinks       = "SELECT ad_name, ad_description, photo_links[1], price FROM adtable WHERE id = $1;"
	readNoFields      = "SELECT ad_name, NULL, photo_links[1], price FROM adtable WHERE id = $1;"

	readTable          = "SELECT ad_name, photo_links[1], price FROM adtable LIMIT $1 OFFSET $2;"
	readTableOrder     = "SELECT ad_name, photo_links[1], price FROM adtable ORDER BY $1 LIMIT $2 OFFSET $3;"
	readTableOrderDesc = "SELECT ad_name, photo_links[1], price FROM adtable ORDER BY $1 LIMIT $2 OFFSET $3 DESC;"
)

type GetAdResponse struct {
	Name        string
	Description *string
	Price       decimal.Decimal
	Links       []string
}

const Limit int64 = 10

type AdditionalFields uint8

const (
	Description AdditionalFields = 1 << iota
	Links
)

type PaginationFields string

const (
	Ascending  PaginationFields = "ascending"
	Descending PaginationFields = "descending"
)

type OrdBy string

const (
	OrderByPrice OrdBy = "price"
	OrderByDate  OrdBy = "date"
)
