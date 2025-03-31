package mysql

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/hexiaopi/data-space/internal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestUserDao_Get(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open(mysql.New(mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	dao := NewUserDao(gdb)

	{ //根据name查询用户
		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "admin")
		mock.ExpectQuery("SELECT * FROM `sys_user` WHERE name = ? ORDER BY `sys_user`.`id` LIMIT ?").WithArgs("admin", 1).WillReturnRows(rows)
		user, err := dao.Get(context.Background(), NewOption().WithName("admin"))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(user)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}
	{ //根据id查询用户
		rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "admin")
		mock.ExpectQuery("SELECT * FROM `sys_user` WHERE id = ? ORDER BY `sys_user`.`id` LIMIT ?").WithArgs(1, 1).WillReturnRows(rows)
		user, err := dao.Get(context.Background(), NewOption().WithId(1))
		if err != nil {
			t.Fatal(err)
		}
		t.Log(user)
		if err := mock.ExpectationsWereMet(); err != nil {
			t.Errorf("there were unfulfilled expectations: %s", err)
		}
	}
}

func TestUserDao_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	gdb, err := gorm.Open(mysql.New(mysql.Config{DriverName: "mysql", Conn: db, SkipInitializeWithVersion: true}), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	dao := NewUserDao(gdb)
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `sys_user`").
		WithArgs("admin", "", "", sqlmock.AnyArg(), sqlmock.AnyArg(), 0).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()
	err = dao.Create(context.Background(), &model.User{Name: "admin"})
	if err != nil {
		t.Fatal(err)
	}
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
