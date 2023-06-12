package history

import (
	"github.com/metadiv-io/base"
	"github.com/metadiv-io/saas/types"
	"gorm.io/gorm"
)

type BaseRepository[T any] struct {
	base.Repository[T]
}

func (r *BaseRepository[T]) CreateWithHistory(tx *gorm.DB, claims *types.Jwt, entity *T, history IHistory) (*T, error) {
	history.SetHistory(claims, ACTION_CREATE)
	err := tx.Create(history).Error
	if err != nil {
		return nil, err
	}
	return r.Save(tx, entity)
}

func (r *BaseRepository[T]) UpdateWithHistory(tx *gorm.DB, claims *types.Jwt, entity *T, history IHistory) (*T, error) {
	history.SetHistory(claims, ACTION_UPDATE)
	err := tx.Create(history).Error
	if err != nil {
		return nil, err
	}
	return r.Save(tx, entity)
}

func (r *BaseRepository[T]) DeleteWithHistory(tx *gorm.DB, claims *types.Jwt, entity *T, history IHistory) error {
	history.SetHistory(claims, ACTION_DELETE)
	err := tx.Create(history).Error
	if err != nil {
		return err
	}
	return r.Delete(tx, entity)
}
