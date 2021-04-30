package common_test

import (
	"testing"

	"github.com/dannov91/multitype/common"
	"github.com/dannov91/multitype/models"
)

func Test_UserName(t *testing.T) {
	tcases := []struct {
		UserType interface{}
		HasErr   bool

		Expected string
	}{
		{
			UserType: models.RegularUser{Name: "Seth"},
			HasErr:   false,
			Expected: "Seth",
		},
		{
			UserType: models.PrivilegedUser{Name: "Goth"},
			HasErr:   false,
			Expected: "Goth",
		},
		{
			UserType: models.PhoneLine{},
			HasErr:   true,
		},
	}

	for _, tc := range tcases {
		name, err := common.UserName(tc.UserType)

		if tc.HasErr {
			if err == nil {
				t.Error("should have error")
			}

			continue
		}

		if err != nil {
			t.Errorf("Unexpected error: %v", err.Error())
		}

		if name != tc.Expected {
			t.Errorf("Name is not equal, expected: %v - got: %v", tc.Expected, name)
		}
	}
}

func Test_SetUserName(t *testing.T) {
	tcases := []struct {
		User   interface{}
		Name   string
		HasErr bool
	}{
		{
			User:   models.RegularUser{Name: "Default"},
			Name:   "Daniel",
			HasErr: false,
		},
		{
			User:   models.RegularUser{Name: "Myne"},
			Name:   "Ralph",
			HasErr: false,
		},
		{
			User:   models.PhoneLine{},
			Name:   "Ralph",
			HasErr: true,
		},
	}

	for _, tc := range tcases {
		if tc.HasErr {
			err := common.SetUserName(&tc.User, tc.Name)
			if err == nil {
				t.Error("should contain error")
			}

			continue
		}

		var name string

		switch tc.User.(type) {
		case models.RegularUser:
			user := tc.User.(models.RegularUser)
			common.SetUserName(&user, tc.Name)
			name = user.Name

		case models.PrivilegedUser:
			user := tc.User.(models.PrivilegedUser)
			common.SetUserName(&user, tc.Name)
			name = user.Name
		}

		if name != tc.Name {
			t.Errorf("Name is not equal, expected: %v - got: %v", tc.Name, name)
		}
	}
}
