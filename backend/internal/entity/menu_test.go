package entity

import (
	"testing"

	"github.com/hexiaopi/data-space/internal/model"
)

func TestGenerateMenuTree(t *testing.T) {
	menus := []model.Menu{
		{ID: 1, ParentId: 0, Name: "菜单1", Order: 1},
		{ID: 2, ParentId: 0, Name: "菜单2", Order: 2},
		{ID: 3, ParentId: 1, Name: "菜单11", Order: 1},
		{ID: 4, ParentId: 1, Name: "菜单12", Order: 2},
		{ID: 5, ParentId: 2, Name: "菜单21", Order: 1},
		{ID: 6, ParentId: 2, Name: "菜单22", Order: 2},
		{ID: 7, ParentId: 6, Name: "菜单221", Order: 2},
		{ID: 8, ParentId: 6, Name: "菜单222", Order: 1},
	}
	tree := GenerateMenuTree(menus)
	for _, v := range tree {
		v.Print(0)
	}
}
