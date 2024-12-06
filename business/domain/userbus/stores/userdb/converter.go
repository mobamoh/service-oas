package userdb

import (
	"fmt"
	"github.com/mobamoh/service-oas/business/domain/common/db/gen"
	"github.com/mobamoh/service-oas/business/domain/userbus"
	"github.com/mobamoh/service-oas/business/types/name"
	"github.com/mobamoh/service-oas/business/types/role"
	"net/mail"
	"time"
)

// ======================== Bus Converters ========================
func toBusUser(db gen.User) (userbus.User, error) {

	nme, err := name.Parse(db.Name)
	if err != nil {
		return userbus.User{}, fmt.Errorf("parse name: %w", err)
	}
	addr := mail.Address{
		Address: db.Email,
	}
	roles, err := role.ParseMany(db.Roles)
	if err != nil {
		return userbus.User{}, fmt.Errorf("parse: %w", err)
	}

	return userbus.User{
		ID:           db.UserID,
		Name:         nme,
		Email:        addr,
		Roles:        roles,
		PasswordHash: db.PasswordHash,
		Enabled:      db.Enabled,
		DateCreated:  db.DateCreated.In(time.Local),
		DateUpdated:  db.DateUpdated.In(time.Local),
	}, nil
}

// ======================== DB Converters ========================

func toDBUser(bus userbus.User) gen.CreateUserParams {
	return gen.CreateUserParams{
		UserID:       bus.ID,
		Name:         bus.Name.String(),
		Email:        bus.Email.Address,
		PasswordHash: bus.PasswordHash,
		Roles:        role.ParseToString(bus.Roles),
		Enabled:      bus.Enabled,
		DateUpdated:  bus.DateUpdated,
		DateCreated:  bus.DateCreated,
	}
}
