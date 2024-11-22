package entity

import (
	"fmt"
	"sort"
	"strings"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
)

// 菜单tree
type MenuTree []*Menu

// 实现sort.Interface接口进行排序
func (a MenuTree) Len() int           { return len(a) }
func (a MenuTree) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a MenuTree) Less(i, j int) bool { return a[i].Order < a[j].Order }

type Menu struct {
	ID         int64    `json:"id"`          //id
	Name       string   `json:"name"`        //名称
	Desc       string   `json:"desc"`        //描述
	Path       string   `json:"path"`        //路径
	Icon       string   `json:"icon"`        //图标
	Component  string   `json:"component"`   //组件
	Type       string   `json:"type"`        //类型
	Order      int64    `json:"order"`       //排序
	ParentId   int64    `json:"parent_id"`   //父级id
	CreateTime string   `json:"create_time"` //创建时间
	UpdateTime string   `json:"update_time"` //修改时间
	Hidden     bool     `json:"hidden"`      //是否隐藏
	Children   MenuTree `json:"children"`    //子级菜单
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
		Hidden:     menu.Hidden,
	}
}

func GenerateMenuTree(menus []model.Menu) MenuTree {
	// 构建菜单树
	menuMap := make(map[int64]*Menu)
	menuTree := make(MenuTree, 0)
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
	SortMenuTree(menuTree)
	return menuTree
}

// 递归排序
func SortMenuTree(tree MenuTree) {
	sort.Sort(tree)
	for _, menu := range tree {
		SortMenuTree(menu.Children)
	}
}
