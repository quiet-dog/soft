package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/do"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"fmt"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sArea struct {
	base.BaseService
}

func init() {
	manage.RegisterManageArea(NewManageArea())
}

func NewManageArea() *sArea {
	return &sArea{}
}

func (s *sArea) Model(ctx context.Context) *gdb.Model {
	return dao.ManageArea.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx)).OnConflict("id")
}

func (s *sArea) GetPageListForSearch(ctx context.Context, req *model.PageListReq, in *req.ManageAreaSearch) (res []*res.AreaTableRow, total int, err error) {
	m := s.handleAreaSearch(ctx, in)
	err = orm.GetPageList(m, req).ScanAndCount(&res, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	return
}

func (s *sArea) Save(ctx context.Context, in *req.ManageAreaSave) (id int64, err error) {
	if in.ParentId != 0 {
		isExitArea, err := s.IsExitAreaById(ctx, in.ParentId)
		if err != nil {
			return 0, err
		}
		if !isExitArea {
			return 0, fmt.Errorf("上级区域不存在")
		}
	}

	var area *do.ManageArea
	if err = gconv.Struct(in, &area); err != nil {
		return
	}

	rs, err := s.Model(ctx).Data(area).Insert()
	if utils.IsError(err) {
		return 0, err
	}
	id, err = rs.LastInsertId()
	return
}

func (s *sArea) Delete(ctx context.Context, ids []int64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	return
}

func (s *sArea) Tree(ctx context.Context, in *req.ManageAreaSearch) (items []*res.AreaTree, err error) {
	query := s.handleAreaSearch(ctx, in)
	var rows []*res.AreaTableRow
	err = query.Fields(dao.ManageArea.Columns().Id, dao.ManageArea.Columns().Name, dao.ManageArea.Columns().ParentId).OrderAsc(dao.ManageArea.Columns().Sort).Scan(&rows)
	if utils.IsError(err) {
		return nil, err
	}
	for _, row := range rows {
		isHave, err := s.isHaveChild(ctx, row.Id)
		if utils.IsError(err) {
			return nil, err
		}
		items = append(items, &res.AreaTree{
			Label:  row.Name,
			Value:  row.Id,
			IsLeaf: isHave,
		})
	}
	return
}

func (s *sArea) IsExitAreaById(ctx context.Context, id int64) (isExit bool, err error) {
	if id <= 0 {
		return false, fmt.Errorf("区域不存在")
	}

	var count int
	count, err = s.Model(ctx).WherePri(id).Count()
	if utils.IsError(err) {
		return false, err
	}
	if count > 0 {
		isExit = true
	}
	return
}

func (s *sArea) isHaveChild(ctx context.Context, id int64) (isHave bool, err error) {

	var count int
	count, err = s.Model(ctx).Where(dao.ManageArea.Columns().ParentId, id).Count()
	if utils.IsError(err) {
		return false, err
	}
	if count > 0 {
		isHave = true
	}
	return
}

func (s *sArea) handleAreaSearch(ctx context.Context, in *req.ManageAreaSearch) (query *gdb.Model) {
	query = s.Model(ctx)
	if !g.IsEmpty(in.Name) {
		query = query.WhereLike(dao.ManageArea.Table()+".name", "%"+in.Name+"%")
	}

	if !g.IsEmpty(in.Remark) {
		query = query.WhereLike(dao.ManageArea.Table()+".remark", "%"+in.Remark+"%")
	}

	if !g.IsNil(in.Ids) && len(in.Ids) > 0 {
		query = query.WhereIn(dao.ManageArea.Table()+".id", in.Ids)
	}

	// if in.ParentId > 0 {
	query = query.Where(dao.ManageArea.Table()+".parent_id", in.ParentId)
	// }
	return
}
