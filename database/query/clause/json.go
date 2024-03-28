package clause

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type JSONContainsPath struct {
	Column string
	Pathes []string
	All    bool
}

func (json JSONContainsPath) Build(builder clause.Builder) {
	if stmt, ok := builder.(*gorm.Statement); ok {
		switch stmt.Dialector.Name() {
		case "mysql":
			builder.WriteString("JSON_CONTAINS_PATH (")
			builder.WriteString(stmt.Quote(json.Column))
			builder.WriteString(", ")
			if json.All {
				builder.AddVar(builder, "all")
			} else {
				builder.AddVar(builder, "one")
			}
			builder.WriteString(", ")
			for _, path := range json.Pathes {
				builder.AddVar(builder, path)
			}
			builder.WriteByte(')')
		}
	}
}
