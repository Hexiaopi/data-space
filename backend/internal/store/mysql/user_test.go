package mysql

import (
	"context"
	"testing"
	"time"

	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/pkg/mysql"
)

func TestUserList(t *testing.T) {
	db, err := mysql.New(&mysql.Config{
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
	if err != nil {
		t.Fatal(err)
	}
	dao := UserDao{db: db}
	var option = NewOption()
	options := []store.Option{option.WithUserName("admin"), option.WithState(1), option.WithPage(10, 1)}
	users, err := dao.List(context.Background(), options...)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range users {
		t.Log(v)
	}
}
