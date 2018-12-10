// Copyright 2018 cloudy 272685110@qq.com.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package all_fileds

import "time"

type RoleMenu struct {
	ID        int       `json:"id" gorm:"column:id" comment:"主键ID"`
	CreatedAt time.Time `json:"created_at,omitempty" gorm:"column:created_at" comment:"记录创建时间"`
	UpdatedAt time.Time `json:"updated_at,omitempty" gorm:"column:updated_at" comment:"记录更新时间"`
	Role      *Role     `json:"role" yaml:"role" comment:"角色"`
	RoleID    int       `json:"role_id" yaml:"role_id" comment:"角色ID"`
	Menu      *Menu     `json:"menu" yaml:"menu" comment:"菜单"`
	MenuID    int       `json:"menu_id" yaml:"menu_id" comment:"菜单ID"`
}

func (mh *RoleMenu) TableName() string {
	return "role_menu"
}
