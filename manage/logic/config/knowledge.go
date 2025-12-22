package config

import (
	"context"
	"devinggo/manage/dao"
	"devinggo/manage/model/do"
	"devinggo/manage/model/entity"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	"devinggo/manage/service/manage"
	"devinggo/modules/system/logic/base"
	"devinggo/modules/system/model"
	"devinggo/modules/system/pkg/handler"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
)

type sManageKnowledge struct {
	base.BaseService
}

func init() {
	manage.RegisterManageKnowledge(NewManageKnowledge())
}

func NewManageKnowledge() *sManageKnowledge {
	return &sManageKnowledge{}
}

func (s *sManageKnowledge) Model(ctx context.Context) *gdb.Model {
	return dao.ManageKnowledge.Ctx(ctx).Hook(hook.Bind()).Cache(orm.SetCacheOption(ctx))
}

func (s *sManageKnowledge) handleSearch(ctx context.Context, in *req.ManageKnowledgeSearch) (m *gdb.Model) {
	m = s.Model(ctx)

	if !g.IsEmpty(in.Name) {
		m = m.Where("name", in.Name)
	}

	if !g.IsEmpty(in.Code) {
		m = m.Where("code", in.Code)
	}

	if !g.IsEmpty(in.KnowledgeType) {
		m = m.Where("knowledge_type", in.KnowledgeType)
	}

	return
}

func (s *sManageKnowledge) GetList(ctx context.Context, inReq *model.ListReq, in *req.ManageKnowledgeSearch) (out []*res.ManageKnowledge, err error) {
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	m = orm.GetList(m, inReq)
	err = m.Scan(&out)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sManageKnowledge) GetPageList(ctx context.Context, req *model.PageListReq, in *req.ManageKnowledgeSearch) (rs []*res.ManageKnowledge, total int, err error) {
	m := s.handleSearch(ctx, in).Handler(handler.FilterAuth)
	var entity []*entity.ManageKnowledge
	err = orm.GetPageList(m, req).ScanAndCount(&entity, &total, false)
	if utils.IsError(err) {
		return nil, 0, err
	}
	rs = make([]*res.ManageKnowledge, 0)
	if !g.IsEmpty(entity) {
		if err = gconv.Structs(entity, &rs); err != nil {
			return nil, 0, err
		}
	}
	return
}

func (s *sManageKnowledge) Save(ctx context.Context, in *req.ManageKnowledgeSave) (id uint64, err error) {
	var saveData *do.ManageKnowledge
	if err = gconv.Struct(in, &saveData); err != nil {
		return
	}
	rs, err := s.Model(ctx).OmitEmptyData().Data(saveData).Save()
	if utils.IsError(err) {
		return
	}
	tmpId, err := rs.LastInsertId()
	if err != nil {
		return
	}
	id = gconv.Uint64(tmpId)
	return
}

func (s *sManageKnowledge) GetById(ctx context.Context, id uint64) (res *res.ManageKnowledge, err error) {
	err = s.Model(ctx).Where("id", id).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sManageKnowledge) Update(ctx context.Context, in *req.ManageKnowledgeUpdate) (err error) {
	var updateData *do.ManageKnowledge
	if err = gconv.Struct(in, &updateData); err != nil {
		return
	}
	_, err = s.Model(ctx).OmitEmptyData().Data(updateData).Where("id", in.Id).Update()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sManageKnowledge) Delete(ctx context.Context, ids []uint64) (err error) {
	_, err = s.Model(ctx).WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sManageKnowledge) RealDelete(ctx context.Context, ids []uint64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Delete()
	if utils.IsError(err) {
		return
	}
	return
}

func (s *sManageKnowledge) Recovery(ctx context.Context, ids []uint64) (err error) {
	_, err = s.Model(ctx).Unscoped().WhereIn("id", ids).Update(g.Map{"deleted_at": nil})
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sManageKnowledge) ChangeStatus(ctx context.Context, id uint64, status int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{"status": status}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sManageKnowledge) NumberOperation(ctx context.Context, id uint64, numberName string, numberValue int) (err error) {
	_, err = s.Model(ctx).Data(g.Map{numberName: numberValue}).Where("id", id).Update()
	if utils.IsError(err) {
		return err
	}
	return
}

func (s *sManageKnowledge) GetExportList(ctx context.Context, req *model.ListReq, in *req.ManageKnowledgeSearch) (res []*res.ManageKnowledgeExcel, err error) {
	m := s.handleSearch(ctx, in)
	err = orm.GetList(m, req).Scan(&res)
	if utils.IsError(err) {
		return
	}
	return
}
