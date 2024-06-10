package entity

import (
	"fmt"
	"strings"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
)

type Menu struct {
	ID         int64   `json:"id"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	Path       string  `json:"path"`
	Icon       string  `json:"icon"`
	Component  string  `json:"component"`
	Type       string  `json:"type"`
	Order      int64   `json:"order"`
	ParentId   int64   `json:"-"`
	CreateTime string  `json:"create_time"`
	UpdateTime string  `json:"update_time"`
	Children   []*Menu `json:"children"`
}

func (m Menu) Print(i int) {
	fmt.Printf("%s%s\n", strings.Repeat("—", i), m.Name)
	for _, child := range m.Children {
		child.Print(i + 1)
	}
}

func ToMenu(menu model.Menu) *Menu {
	return &Menu{
		ID:         menu.ID,
		Name:       menu.Name,
		Desc:       menu.Desc,
		Path:       menu.Path,
		Icon:       menu.Icon,
		Component:  menu.Component,
		Type:       menu.Type,
		Order:      menu.Order,
		ParentId:   menu.ParentId,
		CreateTime: menu.CreateTime.Format(global.DateTimeFormat),
		UpdateTime: menu.UpdateTime.Format(global.DateTimeFormat),
	}
}

func GenerateMenuTree(menus []model.Menu) []*Menu {
	// 构建菜单树
	menuMap := make(map[int64]*Menu)
	menuTree := make([]*Menu, 0)
	for _, menu := range menus {
		menuMap[menu.ID] = ToMenu(menu)
		if menu.ParentId == 0 {
			menuTree = append(menuTree, menuMap[menu.ID])
		}
	}
	for _, menu := range menuMap {
		if parentMenu, ok := menuMap[menu.ParentId]; ok {
			parentMenu.Children = append(parentMenu.Children, menu)
		}
	}
	return menuTree
}
