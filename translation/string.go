package translation

import (
	"database/sql"
	"database/sql/driver"

	"github.com/lib/pq/hstore"
	"github.com/spacetab-io/i18n-go/translation"
)

type String struct {
	translation.String
}

// Implementation of interface for Hstore helpers
func (o *String) Hstore() hstore.Hstore {
	mapping := make(map[string]sql.NullString)

	for lang, str := range o.Translate {
		mapping[lang] = sql.NullString{
			String: str,
			Valid:  true,
		}
	}

	return hstore.Hstore{
		Map: mapping,
	}
}

// Implementation of interface for Hstore helpers
func (o *String) SetHstore(h hstore.Hstore) {
	o.Translate = make(map[string]string)

	for lang, v := range h.Map {
		o.Translate[lang] = v.String
	}
}

func (o *String) Scan(value interface{}) error {
	h := hstore.Hstore{}
	err := h.Scan(value)
	if err != nil {
		return err
	}

	o.SetHstore(h)

	return nil
}

func (o String) Value() (driver.Value, error) {
	return o.Hstore().Value()
}
