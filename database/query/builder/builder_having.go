package builder

import (
	"fmt"

	queryclause "github.com/wardonne/gopi/database/query/clause"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Having add having clause
//
//	builder.Having("id", 1) // id = 1
//	builder.Having(id) // id IS NULL
//	builder.Having(id, []int{1, 2, 3}) // id IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.Having("user_id", query) // id = (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	query := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.Having("user_id", db) // id = (SELECT id FROM users WHERE id = 1)
func (builder *Builder) Having(column any, value any) *Builder {
	builder = builder.instance()
	switch v := value.(type) {
	case *gorm.DB, *Builder, Callback:
		var query = fmt.Sprintf("%s = (?)", builder.QuoteField(column))
		builder.having.Add(clause.Expr{SQL: query, Vars: []any{builder.FormatValue(v)}})
	default:
		builder.having.Add(clause.Eq{
			Column: clause.Expr{
				SQL: builder.QuoteField(column),
			},
			Value: builder.FormatValue(value),
		})
	}
	return builder
}

// HavingNot add having not clause
//
//	builder.HavingNot("id", 1) // id <> 1
//	builder.HavingNot(id) // id IS NOT NULL
//	builder.HavingNot(id, []int{1, 2, 3}) // id NOT IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.HavingNot("user_id", query) // id <> (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	query := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.HavingNot("user_id", db) // id <> (SELECT id FROM users WHERE id = 1)
func (builder *Builder) HavingNot(column any, value any) *Builder {
	builder = builder.instance()
	switch v := value.(type) {
	case *gorm.DB, *Builder, Callback:
		var query = fmt.Sprintf("%s <> (?)", builder.QuoteField(column))
		builder.having.Add(clause.Expr{SQL: query, Vars: []any{builder.FormatValue(v)}})
	default:
		builder.having.Add(clause.Neq{
			Column: clause.Expr{
				SQL: builder.QuoteField(column),
			},
			Value: builder.FormatValue(value),
		})
	}
	return builder
}

// OrHaving add or having not clause
//
//	builder.OrHaving("id", 1) // OR id = 1
//	builder.OrHaving(id) // OR id IS NULL
//	builder.OrHaving(id, []int{1, 2, 3}) // OR id IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrHaving("user_id", query) // OR id = (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	query := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrHaving("user_id", db) // OR id = (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrHaving(column any, value any) *Builder {
	builder = builder.instance()
	switch v := value.(type) {
	case *gorm.DB, *Builder, Callback:
		var query = fmt.Sprintf("%s = (?)", builder.QuoteField(column))
		fmt.Println(query)
		builder.having.Add(clause.Or(clause.Expr{SQL: query, Vars: []any{builder.FormatValue(v)}}))
	default:
		builder.having.Add(clause.Or(clause.Eq{
			Column: clause.Expr{
				SQL: builder.QuoteField(column),
			},
			Value: builder.FormatValue(value),
		}))
	}
	return builder
}

// OrHavingNot add or having not clause
//
//	builder.OrHavingNot("id", 1) // OR id <> 1
//	builder.OrHavingNot(id) // OR id IS NULL
//	builder.OrHavingNot(id, []int{1, 2, 3}) // OR id NOT IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrHavingNot("user_id", query) // OR id <> (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	query := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrHavingNot("user_id", db) // OR id <> (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrHavingNot(column any, value any) *Builder {
	builder = builder.instance()
	switch v := value.(type) {
	case *gorm.DB, *Builder, Callback:
		var query = fmt.Sprintf("%s <> (?)", builder.QuoteField(column))
		builder.having.Add(clause.Or(clause.Expr{SQL: query, Vars: []any{builder.FormatValue(v)}}))
	default:
		builder.having.Add(clause.Or(clause.Neq{
			Column: clause.Expr{
				SQL: builder.QuoteField(column),
			},
			Value: builder.FormatValue(value),
		}))
	}
	return builder
}

// HavingNull add having null clause
//
//	builder.HavingNull("id") // id IS NULL
func (builder *Builder) HavingNull(column any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Eq{Column: clause.Expr{
		SQL: builder.QuoteField(column),
	}})
	return builder
}

// OrHavingNull add or where null clause
//
//	builder.OrHavingNull("id") // OR id IS NULL
func (builder *Builder) OrHavingNull(column any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
	}))
	return builder
}

// HavingNotNull add where not null clause
//
//	builder.HavingNotNull("id") // id IS NOT NULL
func (builder *Builder) HavingNotNull(column any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Not(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
	}))
	return builder
}

// OrHavingNotNull add or where not null clause
//
//	builder.OrHavingNotNull("id") // OR id IS NOT NULL
func (builder *Builder) OrHavingNotNull(column any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Not(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
	})))
	return builder
}

// HavingEq add having equals to clause
//
//	builder.HavingEq("id", 1)
func (builder *Builder) HavingEq(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrHavingEq add or having equals to clause
//
//	builder.OrHavingEq("id", 1)
func (builder *Builder) OrHavingEq(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Eq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	}))
	return builder
}

// HavingNeq add having not equals to clause
//
//	builder.HavingNeq("id", 1)
func (builder *Builder) HavingNeq(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Neq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrHavingNeq add or having not equals to clause
//
//	builder.OrHavingNeq("id", 1)
func (builder *Builder) OrHavingNeq(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Neq{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	}))
	return builder
}

// HavingGt add having greater than clause
//
//	builder.HavingGt("id", 1) // id > 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.HavingGt("user_id", query) // user_id > (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.HavingGt("user_id", tx) // user_id > (SELECT id FROM users WHERE id = 1)
func (builder *Builder) HavingGt(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Gt{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrHavingGt add or having greater than clause
//
//	builder.OrHavingGt("id", 1) // OR id > 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrHavingGt("user_id", query) // OR user_id > (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrHavingGt("user_id", tx) // OR user_id > (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrHavingGt(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Gt{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	}))
	return builder
}

// HavingGte add having greater than or equals to clause
//
//	builder.HavingGte("id", 1) // id >= 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.HavingGte("user_id", query) // user_id >= (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.HavingGte("user_id", tx) // user_id >= (SELECT id FROM users WHERE id = 1)
func (builder *Builder) HavingGte(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Gte{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrHavingGte add having greater than or equals to clause
//
//	builder.OrHavingGte("id", 1) // OR id >= 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrHavingGte("user_id", query) // OR user_id >= (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrHavingGte("user_id", tx) // OR user_id >= (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrHavingGte(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Gte{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	}))
	return builder
}

// HavingLt add having less than clause
//
//	builder.HavingLt("id", 1) // id < 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.HavingLt("user_id", query) // user_id < (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.HavingLt("user_id", tx) // user_id < (SELECT id FROM users WHERE id = 1)
func (builder *Builder) HavingLt(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Lt{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrHavingLt add or having less than clause
//
//	builder.OrHavingLt("id", 1) // OR id < 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrHavingLt("user_id", query) // OR user_id < (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrHavingLt("user_id", tx) // OR user_id < (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrHavingLt(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Lt{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	}))
	return builder
}

// HavingLte add having less than clause
//
//	builder.HavingLte("id", 1) // id <= 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.HavingLte("user_id", query) // user_id <= (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.HavingLte("user_id", tx) // user_id <= (SELECT id FROM users WHERE id = 1)
func (builder *Builder) HavingLte(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Lte{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	})
	return builder
}

// OrHavingLte add or having less than or equals to clause
//
//	builder.OrHavingLte("id", 1) // OR id <= 1
//
//	query := NewBuilder1(db).Select("id").From("users").Where("id", 1)
//	builder.OrHavingLte("user_id", query) // OR user_id <= (SELECT id FROM users WHERE id = 1)
//
//	var db *gorm.DB
//	tx := db.Table("users").Where("id = ?", 1).Select("id")
//	builder.OrHavingLte("user_id", tx) // OR user_id <= (SELECT id FROM users WHERE id = 1)
func (builder *Builder) OrHavingLte(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Lte{
		Column: clause.Expr{
			SQL: builder.QuoteField(column),
		},
		Value: builder.FormatValue(value),
	}))
	return builder
}

// HavingIn add having in clause
//
//	builder.HavingIn("id", 1, 2, 3) // id IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users")
//	builder.HavingIn("id", query) // id IN (SELECT id FROM users)
//
//	var db *gorm.DB
//	tx := db.Table("users").Select("id")
//	builder.HavingIn("id", tx) // id IN (SELECT id FROM users)
func (builder *Builder) HavingIn(column any, values ...any) *Builder {
	builder = builder.instance()
	if len(values) == 1 {
		query := fmt.Sprintf("%s IN (?)", builder.QuoteField(column))
		switch v := values[0].(type) {
		case *Builder, *gorm.DB, Callback:
			builder.having.Add(clause.Expr{SQL: query, Vars: []any{builder.FormatValue(v)}})
		default:
			builder.having.Add(clause.Expr{SQL: query, Vars: []any{v}})
		}
	} else {
		builder.having.Add(clause.IN{
			Column: clause.Expr{SQL: builder.QuoteField(column)},
			Values: builder.FormatValues(values...),
		})
	}
	return builder
}

// OrHavingIn add or having in clause
//
//	builder.OrHavingIn("id", 1, 2, 3) // OR id IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users")
//	builder.OrHavingIn("id", query) // OR id IN (SELECT id FROM users)
//
//	var db *gorm.DB
//	tx := db.Table("users").Select("id")
//	builder.OrHavingIn("id", tx) // OR id IN (SELECT id FROM users)
func (builder *Builder) OrHavingIn(column any, values ...any) *Builder {
	builder = builder.instance()
	if len(values) == 1 {
		query := fmt.Sprintf("%s IN (?)", builder.QuoteField(column))
		switch v := values[0].(type) {
		case *Builder, *gorm.DB, Callback:
			builder.having.Add(clause.Or(clause.Expr{SQL: query, Vars: []any{builder.FormatValue(v)}}))
		default:
			builder.having.Add(clause.Or(clause.Expr{SQL: query, Vars: []any{v}}))
		}
	} else {
		builder.having.Add(clause.Or(clause.IN{
			Column: clause.Expr{SQL: builder.QuoteField(column)},
			Values: builder.FormatValues(values...),
		}))
	}
	return builder
}

// HavingNotIn add having not in clause
//
//	builder.HavingNotIn("id", 1, 2, 3) // id NOT IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users")
//	builder.HavingNotIn("id", query) // id NOT IN (SELECT id FROM users)
//
//	var db *gorm.DB
//	tx := db.Table("users").Select("id")
//	builder.HavingNotIn("id", tx) // id NOT IN (SELECT id FROM users)
func (builder *Builder) HavingNotIn(column any, values ...any) *Builder {
	builder = builder.instance()
	if len(values) == 1 {
		query := fmt.Sprintf("%s NOT IN (?)", builder.QuoteField(column))
		switch v := values[0].(type) {
		case *Builder, *gorm.DB, Callback:
			builder.having.Add(clause.Expr{SQL: query, Vars: []any{builder.FormatValue(v)}})
		default:
			builder.having.Add(clause.Expr{SQL: query, Vars: []any{v}})
		}
	} else {
		builder.having.Add(clause.Not(clause.IN{
			Column: clause.Expr{SQL: builder.QuoteField(column)},
			Values: builder.FormatValues(values...),
		}))
	}
	return builder
}

// OrHavingNotIn add or having not in clause
//
//	builder.OrHavingNotIn("id", 1, 2, 3) // OR id NOT IN (1, 2, 3)
//
//	query := NewBuilder1(db).Select("id").From("users")
//	builder.OrHavingNotIn("id", query) // OR id NOT IN (SELECT id FROM users)
//
//	var db *gorm.DB
//	tx := db.Table("users").Select("id")
//	builder.OrHavingNotIn("id", tx) // OR id NOT IN (SELECT id FROM users)
func (builder *Builder) OrHavingNotIn(column any, values ...any) *Builder {
	builder = builder.instance()
	if len(values) == 1 {
		query := fmt.Sprintf("%s NOT IN (?)", builder.QuoteField(column))
		switch v := values[0].(type) {
		case *Builder, *gorm.DB, Callback:
			builder.having.Add(clause.Or(clause.Expr{SQL: query, Vars: []any{builder.FormatValue(v)}}))
		default:
			builder.having.Add(clause.Or(clause.Expr{SQL: query, Vars: []any{v}}))
		}
	} else {
		builder.having.Add(clause.Or(clause.Not(clause.IN{
			Column: clause.Expr{SQL: builder.QuoteField(column)},
			Values: builder.FormatValues(values...),
		})))
	}
	return builder
}

// HavingBetween add having between clause
//
//	builder.HavingBetween("id", 1, 100) // id BETWEEN 1 AND 100
func (builder *Builder) HavingBetween(column any, start, end any) *Builder {
	builder = builder.instance()
	builder.having.Add(queryclause.Between{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Values: [2]any{
			builder.FormatValue(start),
			builder.FormatValue(end),
		},
	})
	return builder
}

// OrHavingBetween add or having between clause
//
//	builder.OrHavingBetween("id", 1, 100) // id BETWEEN 1 AND 100
func (builder *Builder) OrHavingBetween(column any, start, end any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(queryclause.Between{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Values: [2]any{
			builder.FormatValue(start),
			builder.FormatValue(end),
		},
	}))
	return builder
}

// HavingNotBetween add having not between clause
//
//	builder.HavingNotBetween("id", 1, 100) // id NOT BETWEEN 1 AND 100
func (builder *Builder) HavingNotBetween(column any, start, end any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Not(queryclause.Between{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Values: [2]any{
			builder.FormatValue(start),
			builder.FormatValue(end),
		},
	}))
	return builder
}

// OrHavingNotBetween add or having not between clause
//
//	builder.OrHavingNotBetween("id", 1, 100) // OR id NOT BETWEEN 1 AND 100
func (builder *Builder) OrHavingNotBetween(column string, start, end any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Not(queryclause.Between{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Values: [2]any{
			builder.FormatValue(start),
			builder.FormatValue(end),
		},
	})))
	return builder
}

// HavingLike add where like clause
//
//	builder.HavingLike("name", "%war%")
func (builder *Builder) HavingLike(column string, value string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Like{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Value:  value,
	})
	return builder
}

// HavingNotLike add where not like clause
//
//	builder.HavingNotLike("name", "%war%")
func (builder *Builder) HavingNotLike(column string, value string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Not(clause.Like{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Value:  value,
	}))
	return builder
}

// OrHavingLike add or where like clause
//
//	builder.OrHavingLike("id", "%war%")
func (builder *Builder) OrHavingLike(column string, value string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Like{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Value:  value,
	}))
	return builder
}

// OrHavingNotLike add or where not like clause
//
//	builder.OrHavingNotLike("id", "%war%")
func (builder *Builder) OrHavingNotLike(column string, value string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Not(clause.Like{
		Column: clause.Expr{SQL: builder.QuoteField(column)},
		Value:  value,
	})))
	return builder
}

// HavingBuilder merge having from another [Builder] with AND
func (builder *Builder) HavingBuilder(query *Builder) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.And(query.having.ToArray()...))
	return builder
}

// HavingNotBuilder merge having from another [Builder] with AND NOT
func (builder *Builder) HavingNotBuilder(query *Builder) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Not(query.having.ToArray()...))
	return builder
}

// OrHavingBuilder merge having from another [Builder] with OR
func (builder *Builder) OrHavingBuilder(query *Builder) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.And(query.having.ToArray()...)))
	return builder
}

// OrHavingNotBuilder merge having from another [Builder] with OR NOT
func (builder *Builder) OrHavingNotBuilder(query *Builder) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Not(clause.And(query.having.ToArray()...))))
	return builder
}

// HavingCallback having callback
func (builder *Builder) HavingCallback(callback Clause) *Builder {
	query := NewBuilder(builder.conn)
	return builder.HavingBuilder(callback(query))
}

// HavingNotCallback having NOT callback
func (builder *Builder) HavingNotCallback(callback Clause) *Builder {
	query := NewBuilder(builder.conn)
	return builder.HavingNotBuilder(callback(query))
}

// OrHavingCallback having OR callback
func (builder *Builder) OrHavingCallback(callback Clause) *Builder {
	query := NewBuilder(builder.conn)
	return builder.OrHavingBuilder(callback(query))
}

// OrHavingNotCallback having OR NOT callback
func (builder *Builder) OrHavingNotCallback(callback Clause) *Builder {
	query := NewBuilder(builder.conn)
	return builder.OrHavingNotBuilder(callback(query))
}
