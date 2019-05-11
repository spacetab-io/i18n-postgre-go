# i18n-go-postgre
Expending i18-go structures to bind them for PostgreSQL libs (gorm, lib/pq)

# Usage
## Translation

```go
package pack

import (
	"database/sql"

	"github.com/jinzhu/gorm"
	"github.com/microparts/i18n-go/translation"
	translation_postgre "github.com/microparts/i18n-go-postgre/translation"
	
)

// A model definition

type Record struct {
	Id   int                `json:"id"`                  
	Name translation.String `json:"name"`
}

type RecordGorm struct {
	Id   int                        `json:"id"`                  
	Name translation_postgre.String `json:"name"`
}

func FetchRec(rows *sql.Result) (*Record, error) {
	record := &Record{}
	
	return record, rows.Scan(
		&record.Id,
		&translation_postgre.Bind{
		    V: &record.Name,
	    },
	)
}

func Get(db *gorm.DB, id int) *RecordGorm {
	record := &RecordGorm{}

	db.First(record, id)
	
	return record
}
```