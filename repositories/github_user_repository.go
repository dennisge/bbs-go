package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/mlogclub/simple"

	"github.com/mlogclub/mlog/model"
)

type GithubUserRepository struct {
}

func NewGithubUserRepository() *GithubUserRepository {
	return &GithubUserRepository{}
}

func (this *GithubUserRepository) Get(db *gorm.DB, id int64) *model.GithubUser {
	ret := &model.GithubUser{}
	if err := db.First(ret, "id = ?", id).Error; err != nil {
		return nil
	}
	return ret
}

func (this *GithubUserRepository) Take(db *gorm.DB, where ...interface{}) *model.GithubUser {
	ret := &model.GithubUser{}
	if err := db.Take(ret, where...).Error; err != nil {
		return nil
	}
	return ret
}

func (this *GithubUserRepository) QueryCnd(db *gorm.DB, cnd *simple.QueryCnd) (list []model.GithubUser, err error) {
	err = cnd.DoQuery(db).Find(&list).Error
	return
}

func (this *GithubUserRepository) Query(db *gorm.DB, queries *simple.ParamQueries) (list []model.GithubUser, paging *simple.Paging) {
	queries.StartQuery(db).Find(&list)
	queries.StartCount(db).Model(&model.GithubUser{}).Count(&queries.Paging.Total)
	paging = queries.Paging
	return
}

func (this *GithubUserRepository) Create(db *gorm.DB, t *model.GithubUser) (err error) {
	err = db.Create(t).Error
	return
}

func (this *GithubUserRepository) Update(db *gorm.DB, t *model.GithubUser) (err error) {
	err = db.Save(t).Error
	return
}

func (this *GithubUserRepository) Updates(db *gorm.DB, id int64, columns map[string]interface{}) (err error) {
	err = db.Model(&model.GithubUser{}).Where("id = ?", id).Updates(columns).Error
	return
}

func (this *GithubUserRepository) UpdateColumn(db *gorm.DB, id int64, name string, value interface{}) (err error) {
	err = db.Model(&model.GithubUser{}).Where("id = ?", id).UpdateColumn(name, value).Error
	return
}

func (this *GithubUserRepository) Delete(db *gorm.DB, id int64) {
	db.Model(&model.GithubUser{}).Delete("id", id)
}
