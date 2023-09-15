package micro

import (
	"github.com/metadiv-io/base"
	"github.com/metadiv-io/ginger"
	"github.com/metadiv-io/sql"
	"gorm.io/gorm"
)

type BaseWorkspaceRepository[T any] struct {
	base.Repository[T]
}

func (r *BaseWorkspaceRepository[T]) FindByWorkspaceAndUUID(tx *gorm.DB, workspace string, uuid string) (*T, error) {
	return r.FindOne(tx, workspace, sql.Eq("uuid", uuid))
}

func (r *BaseWorkspaceRepository[T]) FindOne(tx *gorm.DB, workspace string, clause *sql.Clause, tableName ...string) (*T, error) {
	fieldName := "workspace"
	if len(tableName) > 0 {
		fieldName = tableName[0] + ".workspace"
	}
	return sql.FindOne[T](tx, sql.And(
		sql.Eq(fieldName, workspace),
		clause,
	))
}

func (r *BaseWorkspaceRepository[T]) FindAll(tx *gorm.DB, workspace string, clause *sql.Clause, tableName ...string) ([]T, error) {
	fieldName := "workspace"
	if len(tableName) > 0 {
		fieldName = tableName[0] + ".workspace"
	}
	return sql.FindAll[T](tx, sql.And(
		sql.Eq(fieldName, workspace),
		clause,
	))
}

func (r *BaseWorkspaceRepository[T]) Count(tx *gorm.DB, workspace string, clause *sql.Clause, tableName ...string) (int64, error) {
	fieldName := "workspace"
	if len(tableName) > 0 {
		fieldName = tableName[0] + ".workspace"
	}
	return sql.Count[T](tx, sql.And(
		sql.Eq(fieldName, workspace),
		clause,
	))
}

func (r *BaseWorkspaceRepository[T]) FindAllComplex(
	tx *gorm.DB, workspace string, clause *sql.Clause, sort *sql.Sort, page *sql.Pagination, tableName ...string) ([]T, *sql.Pagination, error) {
	fieldName := "workspace"
	if len(tableName) > 0 {
		fieldName = tableName[0] + ".workspace"
	}
	return sql.FindAllComplex[T](tx, sql.And(
		sql.Eq(fieldName, workspace),
		clause,
	), sort, page)
}

func (r *BaseWorkspaceRepository[T]) DeleteBy(tx *gorm.DB, workspace string, clause *sql.Clause, tableName ...string) error {
	fieldName := "workspace"
	if len(tableName) > 0 {
		fieldName = tableName[0] + ".workspace"
	}
	return sql.DeleteAllByClause[T](tx, sql.And(
		sql.Eq(fieldName, workspace),
		clause,
	))
}

type BaseWorkspaceModel struct {
	gorm.Model
	Workspace string `gorm:"not null;"`
	UUID      string `gorm:"not null;"`
}

func (m *BaseWorkspaceModel) SetWorkspace(workspace string) {
	m.Workspace = workspace
}

func (m *BaseWorkspaceModel) SetUUID(prefix ...string) {
	m.UUID = ginger.NewNanoID(prefix...)
}

type GeneralListRequest struct {
	Keyword string `form:"keyword"`
}

func (r *GeneralListRequest) BuildSimilarClause(fields ...string) *sql.Clause {
	var clauses = make([]*sql.Clause, 0)
	for _, field := range fields {
		clauses = append(clauses, sql.Similar(field, r.Keyword))
	}
	if len(clauses) == 0 {
		return nil
	}
	return sql.Or(clauses...)
}

type GeneralGetRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type GeneralUpdateRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}

type GeneralDeleteRequest struct {
	UUID string `uri:"uuid" binding:"required"`
}
