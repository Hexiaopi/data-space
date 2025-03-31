package entity

import (
	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
)

type Department struct {
	ID         int64  `json:"id"`          //id
	Name       string `json:"name"`        //名称
	Desc       string `json:"desc"`        //描述
	State      uint8  `json:"state"`       //状态
	CreateTime string `json:"create_time"` //创建时间
	UpdateTime string `json:"update_time"` //更新时间
}

func ToDepartment(department model.Department) Department {
	return Department{
		ID:         department.ID,
		Name:       department.Name,
		Desc:       department.Desc,
		State:      department.State,
		CreateTime: department.CreateTime.Format(global.DateTimeFormat),
		UpdateTime: department.UpdateTime.Format(global.DateTimeFormat),
	}
}

func ToDepartments(departments []model.Department) []Department {
	result := make([]Department, 0, len(departments))
	for _, department := range departments {
		result = append(result, ToDepartment(department))
	}
	return result
}
