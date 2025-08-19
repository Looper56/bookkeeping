package scope

import (
	"fmt"

	"gorm.io/gorm"
)

// UpdateScope update condition
type UpdateScope struct {
	Condition string
	Value     interface{}
}

// Update update scope
func Update(updateScope ...*UpdateScope) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		for _, scope := range updateScope {
			query := fmt.Sprintf("%s = ?", scope.Condition)
			db = db.Where(query, scope.Value)
		}
		return db
	}
}
