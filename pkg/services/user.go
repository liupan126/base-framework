// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/interfaces/repositories"
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/itcloudy/base-framework/tools"
	"github.com/jinzhu/gorm"
)

const salt = "cloudy"

type UserService struct {
	DB *gorm.DB
	repositories.IUserRepository
}

func (service *UserService) ServiceGetSelf(id int) (result models.UserDetail, err error) {
	return service.FindUserByID(service.DB, id)
}
func (service *UserService) ServiceGetUserByID(id int) (result models.UserDetail, err error) {
	return service.FindUserByID(service.DB, id)
}
func (service *UserService) ServiceGetUserByUserName(username string) (result models.UserDetail, err error) {
	return service.FindUserByUserName(service.DB, username)
}
func (service *UserService) ServiceUserCreate(model models.UserCreate) (result models.UserDetail, err error) {
	model.ID = 0
	model.Pwd = tools.SHA256(tools.StringsJoin(model.Password, salt))
	return service.InsertUser(service.DB, model)
}
func (service *UserService) ServiceUserDelete(ids []int) (err error) {
	return service.DeleteUser(service.DB, ids)
}
func (service *UserService) ServiceCheckUser(username, pwd string) (result models.UserDetail, err error) {
	loginPwd := tools.SHA256(tools.StringsJoin(pwd, salt))
	result, err = service.FindUserByUserNameAndPwd(service.DB, username, loginPwd)
	return
}
func (service *UserService) ServiceGetAllUser(page, size int, order string, query string, queryArgs ...interface{}) (results []*models.UserList, pagination conf.Pagination, err error) {
	pagination.Current = page
	pagination.Size = size
	results, pagination.Total, err = service.FindAllUser(service.DB, page, size, order, query, queryArgs)
	return
}
