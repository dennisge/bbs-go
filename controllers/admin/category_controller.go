package admin

import (
	"strconv"

	"github.com/kataras/iris"
	"github.com/mlogclub/simple"

	"github.com/mlogclub/mlog/model"
	"github.com/mlogclub/mlog/services"
)

type CategoryController struct {
	Ctx             iris.Context
	CategoryService *services.CategoryService
}

func (this *CategoryController) GetBy(id int64) *simple.JsonResult {
	t := this.CategoryService.Get(id)
	if t == nil {
		return simple.ErrorMsg("Not found, id=" + strconv.FormatInt(id, 10))
	}
	return simple.JsonData(t)
}

func (this *CategoryController) AnyList() *simple.JsonResult {
	list, paging := this.CategoryService.Query(simple.NewParamQueries(this.Ctx).
		EqAuto("status").
		PageAuto().Desc("id"))
	return simple.JsonData(&simple.PageResult{Results: list, Page: paging})
}

func (this *CategoryController) PostCreate() *simple.JsonResult {
	t := &model.Category{}
	this.Ctx.ReadForm(t)

	if len(t.Name) == 0 {
		return simple.ErrorMsg("name is required")
	}

	if this.CategoryService.FindByName(t.Name) != nil {
		return simple.ErrorMsg("分类「" + t.Name + "」已存在")
	}

	t.Status = model.CategoryStatusOk
	t.CreateTime = simple.NowTimestamp()
	t.UpdateTime = simple.NowTimestamp()

	err := this.CategoryService.Create(t)
	if err != nil {
		return simple.ErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

func (this *CategoryController) PostUpdate() *simple.JsonResult {
	id := this.Ctx.PostValueInt64Default("id", 0)
	if id <= 0 {
		return simple.ErrorMsg("id is required")
	}
	t := this.CategoryService.Get(id)
	if t == nil {
		return simple.ErrorMsg("entity not found")
	}

	this.Ctx.ReadForm(t)

	if len(t.Name) == 0 {
		return simple.ErrorMsg("name is required")
	}

	if tmp := this.CategoryService.FindByName(t.Name); tmp != nil && tmp.Id != id {
		return simple.ErrorMsg("分类「" + t.Name + "」已存在")
	}

	t.UpdateTime = simple.NowTimestamp()

	err := this.CategoryService.Update(t)
	if err != nil {
		return simple.ErrorMsg(err.Error())
	}
	return simple.JsonData(t)
}

// options选项
func (this *CategoryController) AnyOptions() *simple.JsonResult {
	categories, err := this.CategoryService.GetCategories()
	if err != nil {
		return simple.JsonData([]interface{}{})
	}

	var results []map[string]interface{}
	for _, cat := range categories {
		option := make(map[string]interface{})
		option["value"] = cat.Id
		option["label"] = cat.Name

		results = append(results, option)
	}
	return simple.JsonData(results)
}
