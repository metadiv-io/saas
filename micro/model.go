package micro

import (
	"time"

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
	cls := make([]*sql.Clause, 0)
	cls = append(cls, sql.Eq(fieldName, workspace))
	if clause != nil {
		cls = append(cls, clause)
	}
	return sql.FindOne[T](tx, sql.And(cls...))
}

func (r *BaseWorkspaceRepository[T]) FindAll(tx *gorm.DB, workspace string, clause *sql.Clause, tableName ...string) ([]T, error) {
	fieldName := "workspace"
	if len(tableName) > 0 {
		fieldName = tableName[0] + ".workspace"
	}
	cls := make([]*sql.Clause, 0)
	cls = append(cls, sql.Eq(fieldName, workspace))
	if clause != nil {
		cls = append(cls, clause)
	}
	return sql.FindAll[T](tx, sql.And(cls...))
}

func (r *BaseWorkspaceRepository[T]) Count(tx *gorm.DB, workspace string, clause *sql.Clause, tableName ...string) (int64, error) {
	fieldName := "workspace"
	if len(tableName) > 0 {
		fieldName = tableName[0] + ".workspace"
	}
	cls := make([]*sql.Clause, 0)
	cls = append(cls, sql.Eq(fieldName, workspace))
	if clause != nil {
		cls = append(cls, clause)
	}
	return sql.Count[T](tx, sql.And(cls...))
}

func (r *BaseWorkspaceRepository[T]) FindAllComplex(
	tx *gorm.DB, workspace string, clause *sql.Clause, sort *sql.Sort, page *sql.Pagination, tableName ...string) ([]T, *sql.Pagination, error) {
	fieldName := "workspace"
	if len(tableName) > 0 {
		fieldName = tableName[0] + ".workspace"
	}
	cls := make([]*sql.Clause, 0)
	cls = append(cls, sql.Eq(fieldName, workspace))
	if clause != nil {
		cls = append(cls, clause)
	}
	return sql.FindAllComplex[T](tx, sql.And(cls...), sort, page)
}

func (r *BaseWorkspaceRepository[T]) DeleteBy(tx *gorm.DB, workspace string, clause *sql.Clause, tableName ...string) error {
	fieldName := "workspace"
	if len(tableName) > 0 {
		fieldName = tableName[0] + ".workspace"
	}
	cls := make([]*sql.Clause, 0)
	cls = append(cls, sql.Eq(fieldName, workspace))
	if clause != nil {
		cls = append(cls, clause)
	}
	return sql.DeleteAllByClause[T](tx, sql.And(cls...))
}

type Model struct {
	ID        uint           `gorm:"primarykey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type BaseWorkspaceModel struct {
	Model
	Workspace string `gorm:"not null;" json:"workspace"`
	UUID      string `gorm:"not null;" json:"uuid"`
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
