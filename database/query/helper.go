package query

import "gorm.io/gorm/clause"

// Raw creates clause.Expr
func Raw(sql string, vars ...any) clause.Expr {
	return clause.Expr{
		SQL:  sql,
		Vars: vars,
	}
}
