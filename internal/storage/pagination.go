package storage

import (
	"encoding/json"
	"errors"
)

var ErrBadPaginationFields = errors.New("wrong value of PaginationFields type")

func (j *PaginationFields) UnmarshalJSON(v []byte) error {
	var s string

	if err := json.Unmarshal(v, &s); err != nil {
		return err
	}

	switch s {
	case string(Ascending):
		*j = Ascending
	case string(Descending):
		*j = Descending
	default:
		return ErrBadOrderType
	}

	return nil
}
