// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package common

import (
	"github.com/itcloudy/base-framework/pkg/models"
	"github.com/jinzhu/gorm"
)

type MenuRepository struct {
	*gorm.DB
}

func (repo *MenuRepository) FindMenuByRoles(roleSlice []string) (menus []models.Menu, err error) {
	return
}