package manage

import (
	"context"
	"devinggo/manage/api/manage"
	"devinggo/manage/model/req"
	"devinggo/manage/model/res"
	sManage "devinggo/manage/service/manage"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/myerror"

	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"

	"devinggo/modules/system/pkg/excel"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/request"

	"devinggo/manage/model/do"

	manageDao "devinggo/manage/dao"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/util/gconv"
)

var (
	ManageKnowledgeController = manageKnowledgeController{}
)

type manageKnowledgeController struct {
	base.BaseController
}

func (c *manageKnowledgeController) Index(ctx context.Context, in *manage.IndexManageKnowledgeReq) (out *manage.IndexManageKnowledgeRes, err error) {
	out = &manage.IndexManageKnowledgeRes{}
	items, totalCount, err := sManage.ManageKnowledge().GetPageList(ctx, &in.PageListReq, &in.ManageKnowledgeSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.ManageKnowledge, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *manageKnowledgeController) Recycle(ctx context.Context, in *manage.RecycleManageKnowledgeReq) (out *manage.RecycleManageKnowledgeRes, err error) {
	out = &manage.RecycleManageKnowledgeRes{}
	pageListReq := &in.PageListReq
	pageListReq.Recycle = true
	items, totalCount, err := sManage.ManageKnowledge().GetPageList(ctx, pageListReq, &in.ManageKnowledgeSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.ManageKnowledge, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *manageKnowledgeController) List(ctx context.Context, in *manage.ListManageKnowledgeReq) (out *manage.ListManageKnowledgeRes, err error) {
	out = &manage.ListManageKnowledgeRes{}
	rs, err := sManage.ManageKnowledge().GetList(ctx, &in.ListReq, &req.ManageKnowledgeSearch{})
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.ManageKnowledge, 0)
	}
	return
}

func (c *manageKnowledgeController) Save(ctx context.Context, in *manage.SaveManageKnowledgeReq) (out *manage.SaveManageKnowledgeRes, err error) {
	out = &manage.SaveManageKnowledgeRes{}
	id, err := sManage.ManageKnowledge().Save(ctx, &in.ManageKnowledgeSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *manageKnowledgeController) Read(ctx context.Context, in *manage.ReadManageKnowledgeReq) (out *manage.ReadManageKnowledgeRes, err error) {
	out = &manage.ReadManageKnowledgeRes{}
	info, err := sManage.ManageKnowledge().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *manageKnowledgeController) Update(ctx context.Context, in *manage.UpdateManageKnowledgeReq) (out *manage.UpdateManageKnowledgeRes, err error) {
	out = &manage.UpdateManageKnowledgeRes{}
	err = sManage.ManageKnowledge().Update(ctx, &in.ManageKnowledgeUpdate)
	if err != nil {
		return
	}
	return
}

func (c *manageKnowledgeController) Delete(ctx context.Context, in *manage.DeleteManageKnowledgeReq) (out *manage.DeleteManageKnowledgeRes, err error) {
	out = &manage.DeleteManageKnowledgeRes{}
	err = sManage.ManageKnowledge().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}
func (c *manageKnowledgeController) RealDelete(ctx context.Context, in *manage.RealDeleteManageKnowledgeReq) (out *manage.RealDeleteManageKnowledgeRes, err error) {
	out = &manage.RealDeleteManageKnowledgeRes{}
	err = sManage.ManageKnowledge().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *manageKnowledgeController) Recovery(ctx context.Context, in *manage.RecoveryManageKnowledgeReq) (out *manage.RecoveryManageKnowledgeRes, err error) {
	out = &manage.RecoveryManageKnowledgeRes{}
	err = sManage.ManageKnowledge().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *manageKnowledgeController) ChangeStatus(ctx context.Context, in *manage.ChangeStatusManageKnowledgeReq) (out *manage.ChangeStatusManageKnowledgeRes, err error) {
	out = &manage.ChangeStatusManageKnowledgeRes{}
	err = sManage.ManageKnowledge().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *manageKnowledgeController) NumberOperation(ctx context.Context, in *manage.NumberOperationManageKnowledgeReq) (out *manage.NumberOperationManageKnowledgeRes, err error) {
	out = &manage.NumberOperationManageKnowledgeRes{}
	err = sManage.ManageKnowledge().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}

func (c *manageKnowledgeController) Export(ctx context.Context, in *manage.ExportManageKnowledgeReq) (out *manage.ExportManageKnowledgeRes, err error) {
	var (
		fileName  = "manageKnowledge"
		sheetName = "Sheet1"
	)
	exports, err := sManage.ManageKnowledge().GetExportList(ctx, &in.ListReq, &in.ManageKnowledgeSearch)
	if err != nil {
		return
	}
	//创建导出对象
	export := excel.NewExcelExport(sheetName, res.ManageKnowledgeExcel{})
	//销毁对象
	defer export.Close()
	newExports := []res.ManageKnowledgeExcel{}
	if !g.IsEmpty(exports) {
		for _, item := range exports {
			newExports = append(newExports, *item)
		}
	}
	err = export.ExportSmallExcelByStruct(newExports).Download(ctx, fileName).Error()
	if err != nil {
		return
	}
	return
}

func (c *manageKnowledgeController) Import(ctx context.Context, in *manage.ImportManageKnowledgeReq) (out *manage.ImportManageKnowledgeRes, err error) {
	tmpPath := utils.GetTmpDir()
	fileName, err := in.File.Save(tmpPath, true)
	if err != nil {
		return nil, err
	}
	localPath := tmpPath + "/" + fileName
	var result []res.ManageKnowledgeExcel
	importFile := excel.NewExcelImportFile(localPath, res.ManageKnowledgeExcel{})
	defer importFile.Close()

	err = importFile.ImportDataToStruct(&result).Error()
	if err != nil {
		return nil, err
	} else {
		if !g.IsEmpty(result) {
			err = manageDao.ManageKnowledge.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
				for _, item := range result {
					var saveData *do.ManageKnowledge
					if err = gconv.Struct(item, &saveData); err != nil {
						return
					}
					_, err = sManage.ManageKnowledge().Model(ctx).OmitEmptyData().Data(saveData).Save()
					if err != nil {
						return err
					}
				}
				return
			})
			if err != nil {
				return
			}
		} else {
			err = myerror.ValidationFailed(ctx, "没有数据!")
		}
	}
	return
}

func (c *manageKnowledgeController) DownloadTemplate(ctx context.Context, in *manage.DownloadTemplateManageKnowledgeReq) (out *manage.DownloadTemplateManageKnowledgeRes, err error) {
	var (
		fileName  = "manageKnowledge_template"
		sheetName = "Sheet1"
		exports   = make([]res.ManageKnowledgeExcel, 0)
	)
	export := excel.NewExcelExport(sheetName, res.ManageKnowledgeExcel{})
	defer export.Close()
	err = export.ExportSmallExcelByStruct(exports).Download(ctx, fileName).Error()
	if err != nil {
		return
	}
	return
}

func (c *manageKnowledgeController) Remote(ctx context.Context, in *manage.RemoteManageKnowledgeReq) (out *manage.RemoteManageKnowledgeRes, err error) {
	out = &manage.RemoteManageKnowledgeRes{}
	r := request.GetHttpRequest(ctx)
	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := sManage.ManageKnowledge().Model(ctx)
	var rs res.ManageKnowledge
	remote := orm.NewRemote(m, rs)
	openPage := params.GetVar("openPage")
	items, totalCount, err := remote.GetRemote(ctx, params)
	if err != nil {
		return
	}
	if !g.IsEmpty(openPage) && openPage.Bool() {
		if !g.IsEmpty(items) {
			for _, item := range items {
				out.Items = append(out.Items, item)
			}
		} else {
			out.Items = make([]res.ManageKnowledge, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.ManageKnowledge, 0)
		}
	}
	return
}
