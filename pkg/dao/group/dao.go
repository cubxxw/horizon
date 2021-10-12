package group

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"g.hz.netease.com/horizon/pkg/dao/common"
	"g.hz.netease.com/horizon/pkg/lib/orm"
	"g.hz.netease.com/horizon/pkg/lib/q"

	"gorm.io/gorm"
)

var (
	ErrPathConflict = errors.New("path conflict")
	ErrNameConflict = errors.New("name conflict")
)

type DAO interface {
	// CheckNameUnique check whether the name is unique
	CheckNameUnique(ctx context.Context, group *Group) error
	// CheckPathUnique check whether the path is unique
	CheckPathUnique(ctx context.Context, group *Group) error
	// Create a group
	Create(ctx context.Context, group *Group) (uint, error)
	// Delete a group by id
	Delete(ctx context.Context, id uint) (int64, error)
	// GetByID get a group by id
	GetByID(ctx context.Context, id uint) (*Group, error)
	// GetByNameFuzzily get groups that fuzzily matching the given name
	GetByNameFuzzily(ctx context.Context, name string) ([]*Group, error)
	// GetByIDs get groups by ids
	GetByIDs(ctx context.Context, ids []uint) ([]*Group, error)
	// GetByPaths get groups by paths
	GetByPaths(ctx context.Context, paths []string) ([]*Group, error)
	// CountByParentID get the count of the records matching the given parentID
	CountByParentID(ctx context.Context, parentID uint) (int64, error)
	// UpdateBasic update basic info of a group
	UpdateBasic(ctx context.Context, group *Group) error
	// ListWithoutPage query groups without paging
	ListWithoutPage(ctx context.Context, query *q.Query) ([]*Group, error)
	// List query groups with paging
	List(ctx context.Context, query *q.Query) ([]*Group, int64, error)
	// Transfer move a group under another parent group
	Transfer(ctx context.Context, id, newParentID uint) error
}

// newDAO returns an instance of the default DAO
func newDAO() DAO {
	return &dao{}
}

type dao struct{}

func (d *dao) Transfer(ctx context.Context, id, newParentID uint) error {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return err
	}

	// check records exist
	group, err := d.GetByID(ctx, id)
	if err != nil {
		return err
	}
	pGroup, err := d.GetByID(ctx, newParentID)
	if err != nil {
		return err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// change parentID
		if err := tx.Exec(common.GroupUpdateParentID, newParentID, id).Error; err != nil {
			return err
		}

		// update traversalIDs
		oldTIDs := group.TraversalIDs
		newTIDs := fmt.Sprintf("%s,%d", pGroup.TraversalIDs, group.ID)
		if err := tx.Exec(common.GroupUpdateTraversalIDsPrefix, oldTIDs, newTIDs, oldTIDs+"%").Error; err != nil {
			return err
		}

		// commit when return nil
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (d *dao) CountByParentID(ctx context.Context, parentID uint) (int64, error) {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return 0, err
	}

	var count int64
	result := db.Raw(common.GroupCountByParentID, parentID).Scan(&count)

	return count, result.Error
}

func (d *dao) GetByPaths(ctx context.Context, paths []string) ([]*Group, error) {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	var groups []*Group
	result := db.Raw(common.GroupQueryByPaths, paths).Scan(&groups)

	return groups, result.Error
}

func (d *dao) GetByIDs(ctx context.Context, ids []uint) ([]*Group, error) {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	var groups []*Group
	result := db.Raw(common.GroupQueryByIDs, ids).Scan(&groups)

	return groups, result.Error
}

// CheckPathUnique todo check application table too
func (d *dao) CheckPathUnique(ctx context.Context, group *Group) error {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return err
	}

	queryResult := Group{}
	result := db.Raw(common.GroupQueryByParentIDAndPath, group.ParentID, group.Path).First(&queryResult)

	// update group conflict, has another record with the same parentID & path
	if group.ID > 0 && queryResult.ID > 0 && queryResult.ID != group.ID {
		return ErrPathConflict
	}

	// create group conflict
	if group.ID == 0 && result.RowsAffected > 0 {
		return ErrPathConflict
	}

	return nil
}

func (d *dao) GetByNameFuzzily(ctx context.Context, name string) ([]*Group, error) {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	var groups []*Group
	result := db.Raw(common.GroupQueryByNameFuzzily, fmt.Sprintf("%%%s%%", name)).Scan(&groups)

	return groups, result.Error
}

// CheckNameUnique todo check application table too
func (d *dao) CheckNameUnique(ctx context.Context, group *Group) error {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return err
	}

	queryResult := Group{}
	result := db.Raw(common.GroupQueryByParentIDAndName, group.ParentID, group.Name).First(&queryResult)

	// update group conflict, has another record with the same parentID & name
	if group.ID > 0 && queryResult.ID > 0 && queryResult.ID != group.ID {
		return ErrNameConflict
	}

	// create group conflict
	if group.ID == 0 && result.RowsAffected > 0 {
		return ErrNameConflict
	}

	return nil
}

func (d *dao) Create(ctx context.Context, group *Group) (uint, error) {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return 0, err
	}

	var pGroup *Group
	// check if parent exists
	if group.ParentID > 0 {
		pGroup, err = d.GetByID(ctx, group.ParentID)
		if err != nil {
			return 0, err
		}
	}

	// check if there's a record with the same parentID and name
	err = d.CheckNameUnique(ctx, group)
	if err != nil {
		return 0, err
	}
	// check if there's a record with the same parentID and path
	err = d.CheckPathUnique(ctx, group)
	if err != nil {
		return 0, err
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		// create, get id returned by the database
		if err := tx.Create(group).Error; err != nil {
			// rollback when error
			return err
		}

		// update traversalIDs
		id := group.ID
		var traversalIDs string
		if pGroup == nil {
			traversalIDs = strconv.Itoa(int(id))
		} else {
			traversalIDs = fmt.Sprintf("%s,%d", pGroup.TraversalIDs, id)
		}

		if err := tx.Exec(common.GroupUpdateTraversalIDs, traversalIDs, id).Error; err != nil {
			// rollback when error
			return err
		}

		// commit when return nil
		return nil
	})

	if err != nil {
		return 0, err
	}

	return group.ID, nil
}

// Delete can only delete a group that doesn't have any children
func (d *dao) Delete(ctx context.Context, id uint) (int64, error) {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return 0, err
	}

	result := db.Exec(common.GroupDelete, id)

	return result.RowsAffected, result.Error
}

func (d *dao) GetByID(ctx context.Context, id uint) (*Group, error) {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	var group Group
	result := db.Raw(common.GroupQueryByID, id).First(&group)

	return &group, result.Error
}

func (d *dao) ListWithoutPage(ctx context.Context, query *q.Query) ([]*Group, error) {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return nil, err
	}

	var groups []*Group

	sort := orm.FormatSortExp(query)
	result := db.Order(sort).Where(query.Keywords).Find(&groups)

	return groups, result.Error
}

func (d *dao) List(ctx context.Context, query *q.Query) ([]*Group, int64, error) {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return nil, 0, err
	}

	var groups []*Group

	sort := orm.FormatSortExp(query)
	offset := (query.PageNumber - 1) * query.PageSize
	var count int64
	result := db.Order(sort).Where(query.Keywords).Offset(offset).Limit(query.PageSize).Find(&groups).
		Offset(-1).Count(&count)
	return groups, count, result.Error
}

// UpdateBasic just update base info, not contains transfer logic
func (d *dao) UpdateBasic(ctx context.Context, group *Group) error {
	db, err := orm.FromContext(ctx)
	if err != nil {
		return err
	}

	result := db.Exec(common.GroupUpdateBasic, group.Name, group.Path, group.Description, group.VisibilityLevel, group.ID)

	return result.Error
}