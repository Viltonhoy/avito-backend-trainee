package storage

import (
	"encoding/json"
	"errors"
)

var ErrBadOrderType = errors.New("wrong value of ordBy type")

func (j *OrdBy) UnmarshalJSON(v []byte) error {
	var s string

	if err := json.Unmarshal(v, &s); err != nil {
		return err
	}

	switch s {
	case string(OrderByPrice):
		*j = OrderByPrice
	case string(OrderByDate):
		*j = OrderByDate
	default:
		return ErrBadOrderType
	}

	return nil
}
