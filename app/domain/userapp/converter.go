package userapp

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/mobamoh/service-oas/app/oas"
	"github.com/mobamoh/service-oas/business/domain/userbus"
	"github.com/mobamoh/service-oas/business/types/name"
	"github.com/mobamoh/service-oas/business/types/role"
	"net/mail"
	"time"
)

func toAppUser(bus userbus.User) *oas.User {
	return &oas.User{
		ID:          bus.ID.String(),
		Name:        bus.Name.String(),
		Email:       bus.Email.String(),
		Department:  oas.OptString{Value: bus.Department.String()},
		Enabled:     oas.NewOptBool(bus.Enabled),
		DateCreated: bus.DateCreated.Format(time.RFC3339),
		DateUpdated: bus.DateUpdated.Format(time.RFC3339),
	}
}

func toBusNewUser(app oas.UserCommand) (userbus.NewUser, error) {

	var appRoles []string
	for _, rl := range app.GetRoles() {
		appRoles = append(appRoles, string(rl))
	}

	roles, err := role.ParseMany(appRoles)
	if err != nil {
		return userbus.NewUser{}, fmt.Errorf("parse: %w", err)
	}

	nme, err := name.Parse(app.GetName())
	if err != nil {
		return userbus.NewUser{}, fmt.Errorf("parse: %w,%+v", err, app)
	}

	addr, err := mail.ParseAddress(app.GetEmail())
	if err != nil {
		return userbus.NewUser{}, fmt.Errorf("parse: %w", err)
	}

	department, err := name.ParseNull(app.GetDepartment().Value)
	if err != nil {
		return userbus.NewUser{}, fmt.Errorf("parse: %w", err)
	}

	note := userbus.NewUser{
		Name:       nme,
		Email:      *addr,
		Department: department,
		Roles:      roles,
		Password:   app.Password,
	}
	return note, nil
}

func toBusUserID(id string) (uuid.UUID, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("parse: %w", err)
	}
	return uid, nil
}
