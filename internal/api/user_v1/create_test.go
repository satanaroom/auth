package user_v1

import (
	"context"
	"errors"
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/golang/mock/gomock"
	"github.com/satanaroom/auth/internal/model"
	userMocks "github.com/satanaroom/auth/internal/repository/user/mocks"
	userService "github.com/satanaroom/auth/internal/service/user"
	"github.com/satanaroom/auth/internal/utils"
	desc "github.com/satanaroom/auth/pkg/user_v1"
	"github.com/stretchr/testify/require"
)

func TestCreate(t *testing.T) {
	var (
		ctx      = context.Background()
		mockCtrl = gomock.NewController(t)

		id = gofakeit.Int64()

		fakePassword = gofakeit.Password(true, true, true, true, false, 8)

		username        = gofakeit.Name()
		email           = gofakeit.Email()
		password        = fakePassword
		passwordConfirm = fakePassword

		req = &desc.CreateRequest{
			Info: &desc.UserInfo{
				Username:   username,
				Email:      email,
				Password:   password,
				Role:       1,
				Department: nil,
			},
			PasswordConfirm: passwordConfirm,
		}

		repoReq = model.UserInfo{
			User: model.User{
				Username:   username,
				Email:      email,
				Password:   password,
				Role:       1,
				Department: nil,
			},
			PasswordConfirm: passwordConfirm,
		}

		repoErrText = gofakeit.Phrase()

		repoErr = errors.New(repoErrText)

		validResp = &desc.CreateResponse{
			Id: id,
		}
	)

	passwordHash, hashErr := utils.GeneratePasswordHash(req.Info.Password)
	require.Nil(t, hashErr)
	repoReq.User.Password = passwordHash

	userMock := userMocks.NewMockRepository(mockCtrl)
	gomock.InOrder(
		userMock.EXPECT().Create(ctx, repoReq).Return(id, nil),
		userMock.EXPECT().Create(ctx, repoReq).Return(int64(0), repoErr),
	)
	api := newMockUserV1(Implementation{
		userService: userService.NewMockService(userMock),
	})

	t.Run("success case", func(t *testing.T) {
		resp, err := api.Create(ctx, req)
		require.Nil(t, err)
		require.Equal(t, validResp, resp)
	})

	t.Run("user repo err", func(t *testing.T) {
		_, err := api.Create(ctx, req)
		require.NotNil(t, err)
		require.Equal(t, repoErrText, err.Error())
	})
}
