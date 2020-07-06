package translation

import (
	"database/sql/driver"
	"github.com/lib/pq/hstore"
	"github.com/spacetab-io/i18n-go/translation"
)

const (
	ErrBindNullValue ErrBind = "binding value is nil"
)

type ErrBind string

func (o ErrBind) Error() string {
	return string(o)
}

type Bind struct {
	V *translation.String
}

// Scan implements the Scanner interface.
func (o *Bind) Scan(value interface{}) error {
	if o.V == nil {
		return ErrBindNullValue
	}

	h := hstore.Hstore{}
	err := h.Scan(value)
	if err != nil {
		return err
	}

	v := String{}
	v.SetHstore(h)

	o.V.Translate = v.Translate

	return nil
}

// Value implements the driver Valuer interface.
func (o Bind) Value() (driver.Value, error) {
	if o.V == nil {
		return nil, ErrBindNullValue
	}

	return (&String{*o.V}).Hstore().Value()
}
