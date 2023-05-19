package user

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/satanaroom/auth/internal/model"
	"github.com/satanaroom/auth/pkg/logger"
	desc "github.com/satanaroom/auth/pkg/user_v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToInfo(info *desc.UserInfo, passwordConfirm string) *model.UserInfo {
	var (
		deptType model.Dept
		deptData []byte
		err      error
	)
	switch info.GetDepartment().(type) {
	case *desc.UserInfo_Development:
		deptType = model.DevelopmentDept
		deptData, err = json.Marshal(model.Development{
			Grade:    info.GetDevelopment().GetGrade(),
			Language: info.GetDevelopment().GetLanguage(),
			Rate:     model.Rate(info.GetDevelopment().GetRate()),
		})
		if err != nil {
			logger.Errorf("marshal: %s", err.Error())
			return nil
		}
	case *desc.UserInfo_Analytics:
		deptType = model.AnalyticsDept
		deptData, err = json.Marshal(model.Analytics{
			Specialization: info.GetAnalytics().GetSpecialization(),
			Rate:           model.Rate(info.GetAnalytics().GetRate()),
		})
		if err != nil {
			logger.Errorf("marshal: %s", err.Error())
			return nil
		}
	}

	return &model.UserInfo{
		User: model.User{
			Username:       info.GetUsername(),
			Email:          info.GetEmail(),
			Password:       info.GetPassword(),
			Role:           model.Role(info.GetRole().Number()),
			Department:     deptData,
			DepartmentType: deptType,
		},
		PasswordConfirm: passwordConfirm,
	}
}
func ToUsername(username string) model.Username {
	return model.Username(username)
}

func ToUpdateUser(info *desc.UpdateUser) *model.UserRepo {
	var usernameValid, emailValid, passwordValid, roleValid, departmentValid bool
	if info.Username.ProtoReflect().IsValid() {
		usernameValid = true
	}
	if info.Email.ProtoReflect().IsValid() {
		emailValid = true
	}
	if info.Password.ProtoReflect().IsValid() {
		passwordValid = true
	}
	if info.Role.ProtoReflect().IsValid() {
		roleValid = true
	}

	var (
		deptType model.Dept
		deptData []byte
		err      error
	)
	switch info.GetDepartment().(type) {
	case *desc.UpdateUser_Development:
		deptType = model.DevelopmentDept
		deptData, err = json.Marshal(model.Development{
			Grade:    info.GetDevelopment().GetGrade(),
			Language: info.GetDevelopment().GetLanguage(),
			Rate:     model.Rate(info.GetDevelopment().GetRate()),
		})
		if err != nil {
			logger.Errorf("marshal: %s", err.Error())
			return nil
		}
	case *desc.UpdateUser_Analytics:
		deptType = model.AnalyticsDept
		deptData, err = json.Marshal(model.Analytics{
			Specialization: info.GetAnalytics().GetSpecialization(),
			Rate:           model.Rate(info.GetAnalytics().GetRate()),
		})
		if err != nil {
			logger.Errorf("marshal: %s", err.Error())
			return nil
		}
	}

	if deptType != 0 {
		departmentValid = true
	}

	return &model.UserRepo{
		Username: sql.NullString{
			String: info.Username.GetValue(),
			Valid:  usernameValid,
		},
		Email: sql.NullString{
			String: info.Email.GetValue(),
			Valid:  emailValid,
		},
		Password: sql.NullString{
			String: info.Password.GetValue(),
			Valid:  passwordValid,
		},
		Role: sql.NullInt32{
			Int32: info.Role.GetValue(),
			Valid: roleValid,
		},
		Department: deptData,
		DepartmentType: sql.NullInt32{
			Int32: int32(deptType),
			Valid: departmentValid,
		},
	}
}

func ToGetService(user *model.GetUser) *model.UserService {
	var (
		username, email, password string
		role                      model.Role
		createdAt                 time.Time
		updatedAt                 time.Time
		departmentType            model.Dept
	)

	if user.User.Username.Valid {
		username = user.User.Username.String
	}
	if user.User.Email.Valid {
		email = user.User.Email.String
	}
	if user.User.Password.Valid {
		password = user.User.Password.String
	}
	if user.User.Role.Valid {
		role = model.Role(user.User.Role.Int32)
	}
	if user.CreatedAt.Valid {
		createdAt = user.CreatedAt.Time
	}
	if user.UpdatedAt.Valid {
		updatedAt = user.UpdatedAt.Time
	}
	if user.User.DepartmentType.Valid {
		departmentType = model.Dept(user.User.DepartmentType.Int32)
	}
	return &model.UserService{
		User: model.User{
			Username:       username,
			Email:          email,
			Password:       password,
			Role:           role,
			Department:     user.User.Department,
			DepartmentType: departmentType,
		},
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func ToGetDesc(user *model.UserService) *desc.GetResponse {
	resp := &desc.GetResponse{
		Info: &desc.UserInfo{
			Username: user.User.Username,
			Email:    user.User.Email,
			Password: user.User.Password,
			Role:     desc.Role(user.User.Role),
		},
		CreatedAt: timestamppb.New(user.CreatedAt),
		UpdatedAt: timestamppb.New(user.UpdatedAt),
	}
	switch user.User.DepartmentType {
	case model.DevelopmentDept:
		var dev model.Development
		if err := json.Unmarshal(user.User.Department, &dev); err != nil {
			logger.Errorf("marshal: %s", err.Error())
			return nil
		}
		resp.Info.Department = &desc.UserInfo_Development{
			Development: &desc.Development{
				Grade:    dev.Grade,
				Language: dev.Language,
				Rate:     desc.Rate(dev.Rate),
			},
		}
	case model.AnalyticsDept:
		var analytics model.Analytics
		if err := json.Unmarshal(user.User.Department, &analytics); err != nil {
			logger.Errorf("marshal: %s", err.Error())
			return nil
		}
		resp.Info.Department = &desc.UserInfo_Analytics{
			Analytics: &desc.Analytics{
				Specialization: analytics.Specialization,
				Rate:           desc.Rate(analytics.Rate),
			},
		}
	}
	return resp
}
