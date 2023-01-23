package servicetest

import (
	"errors"
	"testing"

	"github.com/fxmbx/go_practice_pr/dto/userdto"
	mockdb "github.com/fxmbx/go_practice_pr/test/mock"
	"github.com/fxmbx/go_practice_pr/utils"
	"github.com/golang/mock/gomock"
)

func TestCreateUser(t *testing.T) {
	testCases := []struct {
		name       string
		buildStubs func(store mockdb.MockUserService)
	}{
		{
			name: "CreateUser",
			buildStubs: func(store mockdb.MockUserService) {
				arg := userdto.RegisterRequest{
					FirstName: utils.RandomString(6),
					LastName:  utils.RandomString(6),
					Email:     utils.RandomString(6),
					Password:  utils.RandomString(6),
					UserName:  utils.RandomString(6),
				}

				store.
					EXPECT().
					CreateUser(arg).
					AnyTimes().Return(&userdto.UserResponse{}, nil)
			},
		},
		{
			name: "InvalidRequest",
			buildStubs: func(store mockdb.MockUserService) {
				arg := userdto.RegisterRequest{
					FirstName: utils.RandomString(6),
				}

				store.
					EXPECT().
					CreateUser(arg).
					AnyTimes().Return(&userdto.UserResponse{}, errors.New(""))
			},
		},
		{
			name: "UserExist",
			buildStubs: func(store mockdb.MockUserService) {
				arg := userdto.RegisterRequest{
					FirstName: utils.RandomString(6),
					LastName:  utils.RandomString(6),
					Email:     utils.RandomString(6),
					Password:  utils.RandomString(6),
					UserName:  utils.RandomString(6),
				}

				store.
					EXPECT().
					CreateUser(arg).
					AnyTimes().Return(&userdto.UserResponse{}, errors.New(""))
			},
		},
	}
	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			store := mockdb.NewMockUserService(ctrl)

			tc.buildStubs(*store)

		})

	}
}
