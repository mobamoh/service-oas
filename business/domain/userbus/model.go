package userbus

import (
	"github.com/google/uuid"
	"github.com/mobamoh/service-oas/business/types/name"
	"github.com/mobamoh/service-oas/business/types/role"
	"net/mail"
	"time"
)

// User represents information about an individual user.
type User struct {
	ID           uuid.UUID
	Name         name.Name
	Email        mail.Address
	Roles        []role.Role
	PasswordHash []byte
	Department   name.Null
	Enabled      bool
	DateCreated  time.Time
	DateUpdated  time.Time
}
