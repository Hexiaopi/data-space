package mysql

import (
	"context"
	"testing"
	"time"

	"gorm.io/gorm"

	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/pkg/mysql"
)

func setup() (db *gorm.DB, err error) {
	return mysql.New(&mysql.Config{
		Host:                  "127.0.0.1",
		Port:                  "3306",
		UserName:              "root",
		PassWord:              "930619",
		DataBase:              "data-space",
		Charset:               "utf8mb4",
		MaxIdleConnections:    10,
		MaxOpenConnections:    20,
		MaxConnectionLifeTime: time.Hour,
	})
}

func TestComGet(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	dao := ComDao{db: db}
	if err != nil {
		t.Fatal(err)
	}
	options := []Option{
		//WithTable("sys_user"),
		WithId(1),
	}
	user := model.User{}
	if err := dao.Get(context.Background(), "sys_user", &user, options...); err != nil {
		t.Fatal(err)
	}
	t.Log(user)
}

func TestComCreate(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	dao := ComDao{db: db}
	if err != nil {
		t.Fatal(err)
	}
	user := model.User{
		Name:       "张三",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
	}
	if err := dao.Create(context.Background(), &user); err != nil {
		t.Fatal(err)
	}
}

func TestComUpdate(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	dao := ComDao{db: db}
	if err != nil {
		t.Fatal(err)
	}
	user := model.User{
		ID:         6,
		Name:       "张三",
		CreateTime: time.Now(),
		UpdateTime: time.Now(),
		State:      1,
	}
	if err := dao.Update(context.Background(), &user); err != nil {
		t.Fatal(err)
	}
}

func TestComDelete(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	dao := ComDao{db: db}
	if err != nil {
		t.Fatal(err)
	}
	user := model.User{
		ID: 6,
	}
	if err := dao.Delete(context.Background(), &user); err != nil {
		t.Fatal(err)
	}
}

func TestComList(t *testing.T) {
	db, err := setup()
	if err != nil {
		t.Fatal(err)
	}
	dao := ComDao{db: db}
	if err != nil {
		t.Fatal(err)
	}
	options := []Option{
		//WithTable("sys_user"),
		WithPage(10, 1),
	}
	users := make([]model.User, 0)
	if err = dao.List(context.Background(), "sys_user", &users, options...); err != nil {
		t.Fatal(err)
	}
	t.Log(users)
}
