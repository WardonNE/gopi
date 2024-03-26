package builder

import (
	"fmt"

	queryclause "github.com/wardonne/gopi/database/query/clause"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Where add where clause
//
//	builder.Where("id", 1) // id = 1
//	builder.Where(id, nil) // id IS NULL
//	builder.Where(id, []int{1, 2, 3}) // id IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.Where("user_id", query) // id = (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	query := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.Where("user_id", db) // id = (SELECT id FROM users WHERE id = 1)
func (builder *Builder) Where(column any, value any) *Builder {

	switch v := value.(type) {
	case *gorm.DB, *Builder, Callback:
		var query = fmt.Sprintf("%s = (?)", builder.QuoteField(column))
		builder.db = builder.db.Where(query, builder.FormatValue(v))
	default:
		builder.db = builder.db.Where(clause.Eq{
			Column: clause.Expr{
				SQL: builder.QuoteField(column),
			},
			Value: builder.FormatValue(value),
		})
	}
	return builder
}

// WhereNot add where not clause
//
//	builder.WhereNot("id", 1) // id <> 1
//	builder.WhereNot("id") // id IS NOT NULL
//	builder.WhereNot("id", []int{1, 2, 3}) // id NOT IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.WhereNot("user_id", query) // id <> (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	query := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.WhereNot("user_id", db) // id <> (SELECT id FROM users WHERE id = 1)
func (builder *Builder) WhereNot(column any, value any) *Builder {

	switch v := value.(type) {
	case *gorm.DB, *Builder, Callback:
		var query = fmt.Sprintf("%s <> (?)", builder.QuoteField(column))
		builder.db = builder.db.Where(query, builder.FormatValue(v))
	default:
		builder.db = builder.db.Where(clause.Neq{
			Column: clause.Expr{
				SQL: builder.QuoteField(column),
			},
			Value: builder.FormatValue(value),
		})
	}
	return builder
}

// OrWhere add where or clause
//
//	builder.OrWhere("id", 1) // OR id = 1
//	builder.OrWhere("id") // OR id IS NULL
//	builder.OrWhere("id", []int{1, 2, 3}) // OR id IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrWhere("user_id", query) // OR id = (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	query := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrWhere("user_id", db) // OR id = (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrWhere(column any, value any) *Builder {

	switch v := value.(type) {
	case *gorm.DB, *Builder, Callback:
		var query = fmt.Sprintf("%s = (?)", builder.QuoteField(column))
		builder.db = builder.db.Or(query, builder.FormatValue(v))
	default:
		builder.db = builder.db.Or(clause.Eq{
			Column: clause.Expr{
				SQL: builder.QuoteField(column),
			},
			Value: builder.FormatValue(value),
		})
	}
	return builder
}

// OrWhereNot add where or not clause
//
//	builder.OrWhereNot("id", 1) // OR id <> 1
//	builder.OrWhereNot("id") // OR id IS NOT NULL
//	builder.OrWhereNot("id", []int{1, 2, 3}) // OR id IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrWhereNot("user_id", query) // OR id <> (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	query := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrWhereNot("user_id", db) // OR id <> (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrWhereNot(column any, value any) *Builder {

	switch v := value.(type) {
	case *gorm.DB, *Builder, Callback:
		var query = fmt.Sprintf("%s <> (?)", builder.QuoteField(column))
		builder.db = builder.db.Or(query, builder.FormatValue(v))
	default:
		builder.db = builder.db.Or(clause.Neq{
			Column: clause.Expr{
				SQL: builder.QuoteField(column),
			},
			Value: builder.FormatValue(value),
		})
	}
	return builder
}

// WhereNull add where null clause
//
//	builder.WhereNull("id") // id IS NULL
func (builder *Builder) WhereNull(column any) *Builder {

	builder.db = builder.db.Where(clause.Eq{Column: clause.Expr{
		SQL: builder.QuoteField(column),
	}})
	return builder
}

// OrWhereNull add or where null clause
//
//	builder.OrWhereNull("id") // OR id IS NULL
func (builder *Builder) OrWhereNull(column any) *Builder {

	builder.db = builder.db.Or(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
	})
	return builder
}

// WhereNotNull add where not null clause
//
//	builder.WhereNotNull("id") // id IS NOT NULL
func (builder *Builder) WhereNotNull(column any) *Builder {

	builder.db = builder.db.Where(clause.Not(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
	}))
	return builder
}

// OrWhereNotNull add or where not null clause
//
//	builder.OrWhereNotNull("id") // OR id IS NOT NULL
func (builder *Builder) OrWhereNotNull(column any) *Builder {

	builder.db = builder.db.Or(clause.Not(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
	}))
	return builder
}

// WhereEq add where equals to clause
//
//	builder.WhereEq("id", 1)
func (builder *Builder) WhereEq(column any, value any) *Builder {

	builder.db = builder.db.Where(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// WhereNeq add where not equals to clause
//
//	builder.WhereNeq("id", 1)
func (builder *Builder) WhereNeq(column any, value any) *Builder {

	builder.db = builder.db.Where(clause.Neq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrWhereEq add or where equals to clause
//
//	builder.OrWhereEq("id", 1)
func (builder *Builder) OrWhereEq(column any, value any) *Builder {

	builder.db = builder.db.Or(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrWhereNeq add or where not equals to clause
//
//	builder.OrWhereNeq("id", 1)
func (builder *Builder) OrWhereNeq(column any, value any) *Builder {

	builder.db = builder.db.Or(clause.Neq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// WhereGt add where greater than clause
//
//	builder.WhereGt("id", 1) // id > 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.WhereGt("user_id", query) // user_id > (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.WhereGt("user_id", tx) // user_id > (SELECT id FROM users WHERE id = 1)
func (builder *Builder) WhereGt(column any, value any) *Builder {

	builder.db = builder.db.Where(clause.Gt{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrWhereGt add or where greater than clause
//
//	builder.OrWhereGt("id", 1) // OR id > 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrWhereGt("user_id", query) // OR user_id > (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrWhereGt("user_id", tx) // OR user_id > (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrWhereGt(column any, value any) *Builder {

	builder.db = builder.db.Or(clause.Gt{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// WhereGte add where greater than or equals to clause
//
//	builder.WhereGte("id", 1) // id >= 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.WhereGte("user_id", query) // user_id >= (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.WhereGte("user_id", tx) // user_id >= (SELECT id FROM users WHERE id = 1)
func (builder *Builder) WhereGte(column any, value any) *Builder {

	builder.db = builder.db.Where(clause.Gte{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrWhereGte add where greater than or equals to clause
//
//	builder.OrWhereGte("id", 1) // OR id >= 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrWhereGte("user_id", query) // OR user_id >= (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrWhereGte("user_id", tx) // OR user_id >= (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrWhereGte(column any, value any) *Builder {

	builder.db = builder.db.Or(clause.Gte{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// WhereLt add where less than clause
//
//	builder.WhereLt("id", 1) // id < 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.WhereLt("user_id", query) // user_id < (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.WhereLt("user_id", tx) // user_id < (SELECT id FROM users WHERE id = 1)
func (builder *Builder) WhereLt(column any, value any) *Builder {

	builder.db = builder.db.Where(clause.Lt{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrWhereLt add or where less than clause
//
//	builder.OrWhereLt("id", 1) // OR id < 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrWhereLt("user_id", query) // OR user_id < (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrWhereLt("user_id", tx) // OR user_id < (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrWhereLt(column any, value any) *Builder {

	builder.db = builder.db.Or(clause.Lt{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// WhereLte add where less than clause
//
//	builder.WhereLte("id", 1) // id <= 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.WhereLte("user_id", query) // user_id <= (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.WhereLte("user_id", tx) // user_id <= (SELECT id FROM users WHERE id = 1)
func (builder *Builder) WhereLte(column any, value any) *Builder {

	builder.db = builder.db.Where(clause.Lte{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrWhereLte add or where less than or equals to clause
//
//	builder.OrWhereLte("id", 1) // OR id <= 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrWhereLte("user_id", query) // OR user_id <= (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrWhereLte("user_id", tx) // OR user_id <= (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrWhereLte(column any, value any) *Builder {

	builder.db = builder.db.Or(clause.Lte{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// WhereIn add where in clause
//
//	builder.WhereIn("id", 1, 2, 3) // id IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users")
//	builder.WhereIn("id", query) // id IN (SELECT id FROM users)
//
//	var db *gorm.DB
//	tx := db.Table("users").Select("id")
//	builder.WhereIn("id", tx) // id IN (SELECT id FROM users)
func (builder *Builder) WhereIn(column any, values ...any) *Builder {

	if len(values) == 1 {
		query := fmt.Sprintf("%s IN (?)", builder.QuoteField(column))
		switch v := values[0].(type) {
		case *Builder, *gorm.DB, Callback:
			builder.db = builder.db.Where(query, builder.FormatValue(v))
		default:
			builder.db = builder.db.Where(query, v)
		}
	} else {
		builder.db = builder.db.Where(clause.IN{
			Column: clause.Expr{SQL: builder.QuoteField(column)},
			Values: builder.FormatValues(values...),
		})
	}
	return builder
}

// OrWhereIn add or where in clause
//
//	builder.OrWhereIn("id", 1, 2, 3) // OR id IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users")
//	builder.OrWhereIn("id", query) // OR id IN (SELECT id FROM users)
//
//	var db *gorm.DB
//	tx := db.Table("users").Select("id")
//	builder.OrWhereIn("id", tx) // OR id IN (SELECT id FROM users)
func (builder *Builder) OrWhereIn(column any, values ...any) *Builder {

	if len(values) == 1 {
		query := fmt.Sprintf("%s IN (?)", builder.QuoteField(column))
		switch v := values[0].(type) {
		case *Builder, *gorm.DB, Callback:
			builder.db = builder.db.Or(query, builder.FormatValue(v))
		default:
			builder.db = builder.db.Or(query, v)
		}
	} else {
		builder.db = builder.db.Or(clause.IN{
			Column: clause.Expr{SQL: builder.QuoteField(column)},
			Values: builder.FormatValues(values...),
		})
	}
	return builder
}

// WhereNotIn add where not in clause
//
//	builder.WhereNotIn("id", 1, 2, 3) // id NOT IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users")
//	builder.WhereNotIn("id", query) // id NOT IN (SELECT id FROM users)
//
//	var db *gorm.DB
//	tx := db.Table("users").Select("id")
//	builder.WhereNotIn("id", tx) // id NOT IN (SELECT id FROM users)
func (builder *Builder) WhereNotIn(column any, values ...any) *Builder {

	if len(values) == 1 {
		query := fmt.Sprintf("%s NOT IN (?)", builder.QuoteField(column))
		switch v := values[0].(type) {
		case *Builder, *gorm.DB, Callback:
			builder.db = builder.db.Where(query, builder.FormatValue(v))
		default:
			builder.db = builder.db.Where(query, v)
		}
	} else {
		builder.db = builder.db.Where(clause.Not(clause.IN{
			Column: clause.Expr{SQL: builder.QuoteField(column)},
			Values: builder.FormatValues(values...),
		}))
	}
	return builder
}

// OrWhereNotIn add or where not in clause
//
//	builder.OrWhereNotIn("id", 1, 2, 3) // OR id NOT IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users")
//	builder.OrWhereNotIn("id", query) // OR id NOT IN (SELECT id FROM users)
//
//	var db *gorm.DB
//	tx := db.Table("users").Select("id")
//	builder.OrWhereNotIn("id", tx) // OR id NOT IN (SELECT id FROM users)
func (builder *Builder) OrWhereNotIn(column any, values ...any) *Builder {

	if len(values) == 1 {
		query := fmt.Sprintf("%s NOT IN (?)", builder.QuoteField(column))
		switch v := values[0].(type) {
		case *Builder, *gorm.DB, Callback:
			builder.db = builder.db.Or(query, builder.FormatValue(v))
		default:
			builder.db = builder.db.Or(query, v)
		}
	} else {
		builder.db = builder.db.Or(clause.Not(clause.IN{
			Column: clause.Expr{SQL: builder.QuoteField(column)},
			Values: builder.FormatValues(values...),
		}))
	}
	return builder
}

// WhereBetween add where between clause
//
//	builder.WhereBetween("id", 1, 100) // id BETWEEN 1 AND 100
func (builder *Builder) WhereBetween(column any, start, end any) *Builder {

	builder.db = builder.db.Where(queryclause.Between{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Values: [2]any{
			builder.FormatValue(start),
			builder.FormatValue(end),
		},
	})
	return builder
}

// OrWhereBetween add or where between clause
//
//	builder.OrWhereBetween("id", 1, 100) // id BETWEEN 1 AND 100
func (builder *Builder) OrWhereBetween(column string, start, end any) *Builder {

	builder.db = builder.db.Or(queryclause.Between{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Values: [2]any{
			builder.FormatValue(start),
			builder.FormatValue(end),
		},
	})
	return builder
}

// WhereNotBetween add where not between clause
//
//	builder.WhereNotBetween("id", 1, 100) // id NOT BETWEEN 1 AND 100
func (builder *Builder) WhereNotBetween(column string, start, end any) *Builder {

	builder.db = builder.db.Where(clause.Not(queryclause.Between{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Values: [2]any{
			builder.FormatValue(start),
			builder.FormatValue(end),
		},
	}))
	return builder
}

// OrWhereNotBetween add or where not between clause
//
//	builder.OrWhereNotBetween("id", 1, 100) // OR id NOT BETWEEN 1 AND 100
func (builder *Builder) OrWhereNotBetween(column string, start, end any) *Builder {

	builder.db = builder.db.Or(clause.Not(queryclause.Between{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Values: [2]any{
			builder.FormatValue(start),
			builder.FormatValue(end),
		},
	}))
	return builder
}

// WhereLike add where like clause
//
//	builder.WhereLike("name", "%war%")
func (builder *Builder) WhereLike(column any, value string) *Builder {

	builder.db = builder.db.Where(clause.Like{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Value:  value,
	})
	return builder
}

// WhereNotLike add where not like clause
//
//	builder.WhereNotLike("name", "%war%")
func (builder *Builder) WhereNotLike(column any, value string) *Builder {

	builder.db = builder.db.Where(clause.Not(clause.Like{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Value:  value,
	}))
	return builder
}

// OrWhereLike add or where like clause
//
//	builder.OrWhereLike("id", "%war%")
func (builder *Builder) OrWhereLike(column any, value string) *Builder {

	builder.db = builder.db.Or(clause.Like{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Value:  value,
	})
	return builder
}

// OrWhereNotLike add or where not like clause
//
//	builder.OrWhereNotLike("id", "%war%")
func (builder *Builder) OrWhereNotLike(column any, value string) *Builder {

	builder.db = builder.db.Or(clause.Not(clause.Like{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Value:  value,
	}))
	return builder
}
