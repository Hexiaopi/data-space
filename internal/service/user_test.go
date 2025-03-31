package service

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"

	"github.com/hexiaopi/data-space/internal/global"
	"github.com/hexiaopi/data-space/internal/model"
	"github.com/hexiaopi/data-space/internal/store"
	"github.com/hexiaopi/data-space/internal/store/mysql"
	"github.com/hexiaopi/data-space/pkg/jwt"
	"github.com/hexiaopi/data-space/pkg/logger"
)

func TestUserService_LoginSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStore := store.NewMockUserStore(ctrl)
	userStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&model.User{ID: 1, Name: "admin", Password: "$2a$10$qTuvYPjR8os3HltI.cwn1O.pGF2CXtFdMvK2uWfu9pqr/cXaz5V/G"}, nil)

	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(userStore).AnyTimes()

	userSrv := NewUserService(mockFactory, mysql.NewOption(), jwt.NewJWT(time.Hour, "data-space", "hexiaopi"), logger.Std)
	resp, err := userSrv.Login(context.Background(), &LoginRequest{"admin", "123456"})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestUserService_LoginPasswordWrong(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStore := store.NewMockUserStore(ctrl)
	// fake password
	userStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(&model.User{ID: 1, Name: "admin", Password: "xxx"}, nil)

	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(userStore).AnyTimes()

	userSrv := NewUserService(mockFactory, mysql.NewOption(), jwt.NewJWT(time.Hour, "data-space", "hexiaopi"), logger.Std)
	_, err := userSrv.Login(context.Background(), &LoginRequest{"admin", "123456"})
	t.Log(err)
	if !errors.Is(err, global.UserPassWordWrong) {
		t.Fatal("expected error: UserPassWordWrong,but got:", err)
	}
}

func TestUserService_LoginUserNotExist(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userStore := store.NewMockUserStore(ctrl)
	// no exist user
	userStore.EXPECT().Get(gomock.Any(), gomock.Any()).Return(nil, nil)

	mockFactory := store.NewMockFactory(ctrl)
	mockFactory.EXPECT().Users().Return(userStore).AnyTimes()

	userSrv := NewUserService(mockFactory, mysql.NewOption(), jwt.NewJWT(time.Hour, "data-space", "hexiaopi"), logger.Std)
	_, err := userSrv.Login(context.Background(), &LoginRequest{"admin", "123456"})
	t.Log(err)
	if !errors.Is(err, global.UserNotExist) {
		t.Fatal("expected error: UserNotExist,but got:", err)
	}
}
