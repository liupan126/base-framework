// Copyright 2018 cloudy itcloudy@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package services

import (
	"github.com/itcloudy/base-framework/pkg/conf"
	"github.com/itcloudy/base-framework/pkg/models"
)

type ISystemAPIService interface {
	ServiceGetSystemAPIByID(id int) (model models.SystemApiDetail, err error)
	ServiceSystemAPICreate(model models.SystemApiCreate) (result models.SystemApiDetail, err error)
	ServiceSystemAPIUpdate(update models.SystemApiUpdate) (result models.SystemApiDetail, err error)
	ServiceSystemAPIDelete(ids []int) (err error)
	ServiceGetAllSystemAPI(page, size int, order string, query string, queryArgs ...interface{}) (systemApis []models.SystemApiList, pagination conf.Pagination, err error)
}
