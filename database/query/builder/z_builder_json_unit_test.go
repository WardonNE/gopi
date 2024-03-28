package builder

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_WhereJSONContains(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE JSON_CONTAINS (`tags`, JSON_ARRAY(?))").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).WhereJSONContains("tags", 1).Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereJSONNotContains(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE NOT JSON_CONTAINS (`tags`, JSON_ARRAY(?))").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).WhereJSONNotContains("tags", 1).Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereJSONContains(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR JSON_CONTAINS (`tags`, JSON_ARRAY(?))").WithArgs(1, 1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Where("status", 1).OrWhereJSONContains("tags", 1).Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereJSONNotContains(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR NOT JSON_CONTAINS (`tags`, JSON_ARRAY(?))").WithArgs(1, 1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Where("status", 1).OrWhereJSONNotContains("tags", 1).Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereJSONContainsPath(t *testing.T) {
	t.Run("Builder.WhereJSONContainsPath_All", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs("all", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).WhereJSONContainsPath("meta", true, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.WhereJSONContainsPath_One", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs("one", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).WhereJSONContainsPath("meta", false, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereJSONNotContainsPath(t *testing.T) {
	t.Run("Builder.WhereJSONNotContainsPath_All", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE NOT JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs("all", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).WhereJSONNotContainsPath("meta", true, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.WhereJSONNotContainsPath_One", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE NOT JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs("one", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).WhereJSONNotContainsPath("meta", false, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrWhereJSONContainsPath(t *testing.T) {
	t.Run("Builder.OrWhereJSONContainsPath_All", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs(1, "all", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Where("status", 1).OrWhereJSONContainsPath("meta", true, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrWhereJSONContainsPath_One", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs(1, "one", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Where("status", 1).OrWhereJSONContainsPath("meta", false, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})
}
func TestBuilder_OrWhereJSONNotContainsPath(t *testing.T) {
	t.Run("Builder.OrWhereJSONNotContainsPath_All", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR NOT JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs(1, "all", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Where("status", 1).OrWhereJSONNotContainsPath("meta", true, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrWhereJSONNotContainsPath_One", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR NOT JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs(1, "one", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Where("status", 1).OrWhereJSONNotContainsPath("meta", false, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_WhereJSONOverlaps(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE JSON_OVERLAPS(`tags`,?)").WithArgs("[1,2,3]").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).WhereJSONOverlaps("tags", "[1,2,3]").Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereJSONNotOverlaps(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE NOT JSON_OVERLAPS(`tags`,?)").WithArgs("[1,2,3]").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).WhereJSONNotOverlaps("tags", "[1,2,3]").Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereJSONOverlaps(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR JSON_OVERLAPS(`tags`,?)").WithArgs(1, "[1,2,3]").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Where("status", 1).OrWhereJSONOverlaps("tags", "[1,2,3]").Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereJSONNotOverlaps(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR NOT JSON_OVERLAPS(`tags`,?)").WithArgs(1, "[1,2,3]").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Where("status", 1).OrWhereJSONNotOverlaps("tags", "[1,2,3]").Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereJSONHasKey(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE JSON_EXTRACT(`meta`,?) IS NOT NULL").WithArgs("$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").WhereJSONHasKey("meta", "department", "status").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_WhereJSONNotHasKey(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE NOT JSON_EXTRACT(`meta`,?) IS NOT NULL").WithArgs("$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").WhereJSONNotHasKey("meta", "department", "status").Find(&dest)
	assert.Nil(t, err)
}
func TestBuilder_OrWhereJSONHasKey(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR JSON_EXTRACT(`meta`,?) IS NOT NULL").WithArgs(1, "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereJSONHasKey("meta", "department", "status").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrWhereJSONNotHasKey(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` WHERE `status` = ? OR NOT JSON_EXTRACT(`meta`,?) IS NOT NULL").WithArgs(1, "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Where("status", 1).OrWhereJSONNotHasKey("meta", "department", "status").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingJSONContains(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING JSON_CONTAINS (`tags`, JSON_ARRAY(?))").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).HavingJSONContains("tags", 1).Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingJSONNotContains(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING NOT JSON_CONTAINS (`tags`, JSON_ARRAY(?))").WithArgs(1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).HavingJSONNotContains("tags", 1).Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingJSONContains(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR JSON_CONTAINS (`tags`, JSON_ARRAY(?))").WithArgs(1, 1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Having("status", 1).OrHavingJSONContains("tags", 1).Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingJSONNotContains(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR NOT JSON_CONTAINS (`tags`, JSON_ARRAY(?))").WithArgs(1, 1).WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Having("status", 1).OrHavingJSONNotContains("tags", 1).Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingJSONContainsPath(t *testing.T) {
	t.Run("Builder.HavingJSONContainsPath_All", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs("all", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).HavingJSONContainsPath("meta", true, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.HAVINGJSONContainsPath_One", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs("one", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).HavingJSONContainsPath("meta", false, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingJSONNotContainsPath(t *testing.T) {
	t.Run("Builder.HavingJSONNotContainsPath_All", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING NOT JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs("all", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).HavingJSONNotContainsPath("meta", true, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.HavingJSONNotContainsPath_One", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING NOT JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs("one", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).HavingJSONNotContainsPath("meta", false, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_OrHavingJSONContainsPath(t *testing.T) {
	t.Run("Builder.OrHavingJSONContainsPath_All", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs(1, "all", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Having("status", 1).OrHavingJSONContainsPath("meta", true, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrHavingJSONContainsPath_One", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs(1, "one", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Having("status", 1).OrHavingJSONContainsPath("meta", false, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})
}
func TestBuilder_OrHavingJSONNotContainsPath(t *testing.T) {
	t.Run("Builder.OrHavingJSONNotContainsPath_All", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR NOT JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs(1, "all", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Having("status", 1).OrHavingJSONNotContainsPath("meta", true, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})

	t.Run("Builder.OrHavingJSONNotContainsPath_One", func(t *testing.T) {
		mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR NOT JSON_CONTAINS_PATH (`meta`, ?, ?)").WithArgs(1, "one", "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
		var dest = make([]map[string]any, 0)
		err := NewBuilder(mockDB).Having("status", 1).OrHavingJSONNotContainsPath("meta", false, "$.department.status").Table("users").Find(&dest)
		assert.Nil(t, err)
	})
}

func TestBuilder_HavingJSONOverlaps(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING JSON_OVERLAPS(`tags`,?)").WithArgs("[1,2,3]").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).HavingJSONOverlaps("tags", "[1,2,3]").Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingJSONNotOverlaps(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING NOT JSON_OVERLAPS(`tags`,?)").WithArgs("[1,2,3]").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).HavingJSONNotOverlaps("tags", "[1,2,3]").Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingJSONOverlaps(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR JSON_OVERLAPS(`tags`,?)").WithArgs(1, "[1,2,3]").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Having("status", 1).OrHavingJSONOverlaps("tags", "[1,2,3]").Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingJSONNotOverlaps(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR NOT JSON_OVERLAPS(`tags`,?)").WithArgs(1, "[1,2,3]").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Having("status", 1).OrHavingJSONNotOverlaps("tags", "[1,2,3]").Table("users").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingJSONHasKey(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING JSON_EXTRACT(`meta`,?) IS NOT NULL").WithArgs("$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").HavingJSONHasKey("meta", "department", "status").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_HavingJSONNotHasKey(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING NOT JSON_EXTRACT(`meta`,?) IS NOT NULL").WithArgs("$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").HavingJSONNotHasKey("meta", "department", "status").Find(&dest)
	assert.Nil(t, err)
}
func TestBuilder_OrHavingJSONHasKey(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR JSON_EXTRACT(`meta`,?) IS NOT NULL").WithArgs(1, "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingJSONHasKey("meta", "department", "status").Find(&dest)
	assert.Nil(t, err)
}

func TestBuilder_OrHavingJSONNotHasKey(t *testing.T) {
	mock.ExpectQuery("SELECT * FROM `users` HAVING `status` = ? OR NOT JSON_EXTRACT(`meta`,?) IS NOT NULL").WithArgs(1, "$.department.status").WillReturnRows(sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "wardonne"))
	var dest = make([]map[string]any, 0)
	err := NewBuilder(mockDB).Table("users").Having("status", 1).OrHavingJSONNotHasKey("meta", "department", "status").Find(&dest)
	assert.Nil(t, err)
}
