package builder

import (
	"gorm.io/datatypes"
	"gorm.io/gorm/clause"

	queryclause "github.com/wardonne/gopi/database/query/clause"
)

// WhereJSONContains where JSON_CONTAINS
//
//	builder.WhereJSONContains("tags", 1)
func (builder *Builder) WhereJSONContains(column any, value any) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Where(datatypes.JSONArrayQuery(builder.QuoteField(column)).Contains(builder.FormatValue(value)))
	return builder
}

// WhereJSONNotContains where NOT JSON_CONTAINS
//
//	builder.WhereJSONNOTContains("tags", 1)
func (builder *Builder) WhereJSONNotContains(column any, value any) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Not(datatypes.JSONArrayQuery(builder.QuoteField(column)).Contains(builder.FormatValue(value)))
	return builder
}

// OrWhereJSONContains where OR JSON_CONTAINS
//
//	builder.OrWhereJSONContains("tags", 1)
func (builder *Builder) OrWhereJSONContains(column any, value any) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Or(datatypes.JSONArrayQuery(builder.QuoteField(column)).Contains(builder.FormatValue(value)))
	return builder
}

// OrWhereJSONNotContains where OR NOT JSON_CONTAINS
//
//	builder.OrWhereJSONNotContains("tags", 1)
func (builder *Builder) OrWhereJSONNotContains(column any, value any) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Or(clause.Not(datatypes.JSONArrayQuery(builder.QuoteField(column)).Contains(builder.FormatValue(value))))
	return builder
}

// WhereJSONContainsPath where JSON_CONTAINS_PATH
//
//	builder.WhereJSONContainsPath("meta", true, "$.department.status", "$.department.name")
//	builder.WhereJSONContainsPath("meta", false, "$.department.status", "$.department.name")
func (builder *Builder) WhereJSONContainsPath(column any, all bool, pathes ...string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Where(queryclause.JSONContainsPath{
		Column: builder.QuoteField(column),
		Pathes: pathes,
		All:    all,
	})
	return builder
}

// WhereJSONNotContainsPath where NOT JSON_CONTAINS_PATH
//
//	builder.WhereJSONNotContainsPath("meta", true, "$.department.status", "$.department.name")
//	builder.WhereJSONNotContainsPath("meta", false, "$.department.status", "$.department.name")
func (builder *Builder) WhereJSONNotContainsPath(column any, all bool, pathes ...string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Not(queryclause.JSONContainsPath{
		Column: builder.QuoteField(column),
		Pathes: pathes,
		All:    all,
	})
	return builder
}

// OrWhereJSONContainsPath where OR JSON_CONTAINS_PATH
//
//	builder.OrWhereJSONContainsPath("meta", true, "$.department.status", "$.department.name")
//	builder.OrWhereJSONContainsPath("meta", false, "$.department.status", "$.department.name")
func (builder *Builder) OrWhereJSONContainsPath(column any, all bool, pathes ...string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Or(queryclause.JSONContainsPath{
		Column: builder.QuoteField(column),
		Pathes: pathes,
		All:    all,
	})
	return builder
}

// OrWhereJSONNotContainsPath where OR NOT JSON_CONTAINS_PATH
//
//	builder.OrWhereJSONNotContainsPath("meta", true, "$.department.status", "$.department.name")
//	builder.OrWhereJSONNotContainsPath("meta", false, "$.department.status", "$.department.name")
func (builder *Builder) OrWhereJSONNotContainsPath(column any, all bool, pathes ...string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Or(clause.Not(queryclause.JSONContainsPath{
		Column: builder.QuoteField(column),
		Pathes: pathes,
		All:    all,
	}))
	return builder
}

// WhereJSONOverlaps where JSON_OVERLAPS
//
//	builder.WhereJSONOverlaps("tags", "[1,2,3]")
func (builder *Builder) WhereJSONOverlaps(column any, value string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Where(datatypes.JSONOverlaps(clause.Expr{
		SQL: builder.QuoteField(column),
	}, value))
	return builder
}

// WhereJSONNotOverlaps where NOT JSON_OVERLAPS
//
//	builder.WhereJSONOverlaps("tags", "[1,2,3]")
func (builder *Builder) WhereJSONNotOverlaps(column any, value string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Not(datatypes.JSONOverlaps(clause.Expr{
		SQL: builder.QuoteField(column),
	}, value))
	return builder
}

// OrWhereJSONOverlaps where OR JSON_OVERLAPS
//
//	builder.OrWhereJSONOverlaps("tags", "[1,2,3]")
func (builder *Builder) OrWhereJSONOverlaps(column any, value string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Or(datatypes.JSONOverlaps(clause.Expr{
		SQL: builder.QuoteField(column),
	}, value))
	return builder
}

// OrWhereJSONNotOverlaps where OR NOT JSON_OVERLAPS
//
//	builder.OrWhereJSONNotOverlaps("tags", "[1,2,3]")
func (builder *Builder) OrWhereJSONNotOverlaps(column any, value string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Or(clause.Not(datatypes.JSONOverlaps(clause.Expr{
		SQL: builder.QuoteField(column),
	}, value)))
	return builder
}

// WhereJSONHasKey Where JSON_EXTRACT IS NOT NULL
//
//	builder.WhereJSONHasKey("meta", "department", "status")
func (builder *Builder) WhereJSONHasKey(column any, keys ...string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Where(datatypes.JSONQuery(builder.QuoteField(column)).HasKey(keys...))
	return builder
}

// WhereJSONNotHasKey Where JSON_EXTRACT IS NULL
//
//	builder.WhereJSONNotHasKey("meta", "department", "status")
func (builder *Builder) WhereJSONNotHasKey(column any, keys ...string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Not(datatypes.JSONQuery(builder.QuoteField(column)).HasKey(keys...))
	return builder
}

// OrWhereJSONHasKey Where OR JSON_EXTRACT IS NOT NULL
//
//	builder.OrWhereJSONHasKey("meta", "department", "status")
func (builder *Builder) OrWhereJSONHasKey(column any, keys ...string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Or(datatypes.JSONQuery(builder.QuoteField(column)).HasKey(keys...))
	return builder
}

// OrWhereJSONNotHasKey Where OR JSON_EXTRACT IS NULL
//
//	builder.OrWhereJSONNotHasKey("meta", "department", "status")
func (builder *Builder) OrWhereJSONNotHasKey(column any, keys ...string) *Builder {
	builder = builder.instance()
	builder.db = builder.db.Or(clause.Not(datatypes.JSONQuery(builder.QuoteField(column)).HasKey(keys...)))
	return builder
}

// HavingJSONContains Having JSON_CONTAINS
//
//	builder.HavingJSONContains("tags", 1)
func (builder *Builder) HavingJSONContains(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(datatypes.JSONArrayQuery(builder.QuoteField(column)).Contains(builder.FormatValue(value)))
	return builder
}

// HavingJSONNotContains Having NOT JSON_CONTAINS
//
//	builder.HavingJSONNotContains("tags", 1)
func (builder *Builder) HavingJSONNotContains(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Not(datatypes.JSONArrayQuery(builder.QuoteField(column)).Contains(builder.FormatValue(value))))
	return builder
}

// OrHavingJSONContains Having OR JSON_CONTAINS
//
//	builder.OrHavingJSONContains("tags", 1)
func (builder *Builder) OrHavingJSONContains(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(datatypes.JSONArrayQuery(builder.QuoteField(column)).Contains(builder.FormatValue(value))))
	return builder
}

// OrHavingJSONNotContains Having OR NOT JSON_CONTAINS
//
//	builder.OrHavingJSONNotContains("tags", 1)
func (builder *Builder) OrHavingJSONNotContains(column any, value any) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Not(datatypes.JSONArrayQuery(builder.QuoteField(column)).Contains(builder.FormatValue(value)))))
	return builder
}

// HavingJSONContainsPath Having JSON_CONTAINS_PATH
//
//	builder.HavingJSONContainsPath("meta", true, "$.department.status", "$.department.name")
//	builder.HavingJSONContainsPath("meta", false, "$.department.status", "$.department.name")
func (builder *Builder) HavingJSONContainsPath(column any, all bool, pathes ...string) *Builder {
	builder = builder.instance()
	builder.having.Add(queryclause.JSONContainsPath{
		Column: builder.QuoteField(column),
		Pathes: pathes,
		All:    all,
	})
	return builder
}

// HavingJSONNotContainsPath Having NOT JSON_CONTAINS_PATH
//
//	builder.HavingJSONNotContainsPath("meta", true, "$.department.status", "$.department.name")
//	builder.HavingJSONNotContainsPath("meta", false, "$.department.status", "$.department.name")
func (builder *Builder) HavingJSONNotContainsPath(column any, all bool, pathes ...string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Not(queryclause.JSONContainsPath{
		Column: builder.QuoteField(column),
		Pathes: pathes,
		All:    all,
	}))
	return builder
}

// OrHavingJSONContainsPath Having OR JSON_CONTAINS_PATH
//
//	builder.OrHavingJSONContainsPath("meta", true, "$.department.status", "$.department.name")
//	builder.OrHavingJSONContainsPath("meta", false, "$.department.status", "$.department.name")
func (builder *Builder) OrHavingJSONContainsPath(column any, all bool, pathes ...string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(queryclause.JSONContainsPath{
		Column: builder.QuoteField(column),
		Pathes: pathes,
		All:    all,
	}))
	return builder
}

// OrHavingJSONNotContainsPath Having OR NOT JSON_CONTAINS_PATH
//
//	builder.OrHavingJSONNotContainsPath("meta", true, "$.department.status", "$.department.name")
//	builder.OrHavingJSONNotContainsPath("meta", false, "$.department.status", "$.department.name")
func (builder *Builder) OrHavingJSONNotContainsPath(column any, all bool, pathes ...string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Not(queryclause.JSONContainsPath{
		Column: builder.QuoteField(column),
		Pathes: pathes,
		All:    all,
	})))
	return builder
}

// HavingJSONOverlaps Having JSON_OVERLAPS
//
//	builder.HavingJSONOverlaps("tags", "[1,2,3]")
func (builder *Builder) HavingJSONOverlaps(column any, value string) *Builder {
	builder = builder.instance()
	builder.having.Add(datatypes.JSONOverlaps(clause.Expr{
		SQL: builder.QuoteField(column),
	}, value))
	return builder
}

// HavingJSONNotOverlaps Having NOT JSON_OVERLAPS
//
//	builder.HavingJSONNotOverlaps("tags", "[1,2,3]")
func (builder *Builder) HavingJSONNotOverlaps(column any, value string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Not(datatypes.JSONOverlaps(clause.Expr{
		SQL: builder.QuoteField(column),
	}, value)))
	return builder
}

// OrHavingJSONOverlaps Having OR JSON_OVERLAPS
//
//	builder.OrHavingJSONOverlaps("tags", "[1,2,3]")
func (builder *Builder) OrHavingJSONOverlaps(column any, value string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(datatypes.JSONOverlaps(clause.Expr{
		SQL: builder.QuoteField(column),
	}, value)))
	return builder
}

// OrHavingJSONNotOverlaps Having OR NOT JSON_OVERLAPS
//
//	builder.OrHavingJSONNotOverlaps("tags", "[1,2,3]")
func (builder *Builder) OrHavingJSONNotOverlaps(column any, value string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Not(datatypes.JSONOverlaps(clause.Expr{
		SQL: builder.QuoteField(column),
	}, value))))
	return builder
}

// HavingJSONHasKey Having JSON_EXTRACT IS NOT NULL
//
//	builder.HavingJSONHasKey("meta", "department", "status")
func (builder *Builder) HavingJSONHasKey(column any, keys ...string) *Builder {
	builder = builder.instance()
	builder.having.Add(datatypes.JSONQuery(builder.QuoteField(column)).HasKey(keys...))
	return builder
}

// HavingJSONNotHasKey Having JSON_EXTRACT IS NOT NULL
//
//	builder.HavingJSONNotHasKey("meta", "department", "status")
func (builder *Builder) HavingJSONNotHasKey(column any, keys ...string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Not(datatypes.JSONQuery(builder.QuoteField(column)).HasKey(keys...)))
	return builder
}

// OrHavingJSONHasKey Having JSON_EXTRACT IS NOT NULL
//
//	builder.OrHavingJSONHasKey("meta", "department", "status")
func (builder *Builder) OrHavingJSONHasKey(column any, keys ...string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(datatypes.JSONQuery(builder.QuoteField(column)).HasKey(keys...)))
	return builder
}

// OrHavingJSONNotHasKey Having JSON_EXTRACT IS NOT NULL
//
//	builder.OrHavingJSONNotHasKey("meta", "department", "status")
func (builder *Builder) OrHavingJSONNotHasKey(column any, keys ...string) *Builder {
	builder = builder.instance()
	builder.having.Add(clause.Or(clause.Not(datatypes.JSONQuery(builder.QuoteField(column)).HasKey(keys...))))
	return builder
}
