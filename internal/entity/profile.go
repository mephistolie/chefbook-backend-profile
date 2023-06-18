package entity

import (
	"time"
)

type ProfileMinInfo struct {
	Id          string
	VisibleName *string
	Avatar      *string
}

type Profile struct {
	// Auth Service
	Id                    *string
	Nickname              *string
	Email                 *string
	Role                  *string
	OAuth                 *OAuth
	IsBlocked             bool
	RegistrationTimestamp *time.Time

	// User Service
	FirstName   *string
	LastName    *string
	Description *string
	AvatarLink  *string
}

type OAuth struct {
	GoogleId *string
	VkId     *int64
}
