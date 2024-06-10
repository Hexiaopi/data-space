package entity

import (
	"testing"

	"github.com/hexiaopi/data-space/internal/model"
)

func TestGenerateMenuTree(t *testing.T) {
	menus := []model.Menu{
		{ID: 1, ParentId: 0, Name: "菜单1"},
		{ID: 2, ParentId: 0, Name: "菜单2"},
		{ID: 3, ParentId: 1, Name: "菜单11"},
		{ID: 4, ParentId: 1, Name: "菜单12"},
		{ID: 5, ParentId: 2, Name: "菜单21"},
		{ID: 6, ParentId: 2, Name: "菜单22"},
		{ID: 7, ParentId: 6, Name: "菜单221"},
		{ID: 8, ParentId: 6, Name: "菜单222"},
	}
	tree := GenerateMenuTree(menus)
	for _, v := range tree {
		v.Print(0)
	}
}
