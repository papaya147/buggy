package bug

import (
	"time"

	"github.com/google/uuid"
	db "github.com/papaya147/buggy/backend/db/sqlc"
)

type createBugInput struct {
	Name         string         `json:"name" validate:"required,max=32" example:"Improper Input Validation"`
	Description  string         `json:"description" validate:"required,max=255" example:"Input validation is not working"`
	Priority     db.Bugpriority `json:"priority" validate:"required,oneof=URGENT HIGH LOW" example:"URGENT"`
	AssignedTeam uuid.UUID      `json:"assigned_team" validate:"required" example:"00000000-0000-0000-0000-000000000000"`
	AssigneeTeam uuid.UUID      `json:"assignee_team" validate:"required" example:"00000000-0000-0000-0000-000000000000"`
}

type bugOutput struct {
	Id                uuid.UUID      `json:"id" example:"00000000-0000-0000-0000-000000000000"`
	Name              string         `json:"name" example:"Improper Input Validation"`
	Description       string         `json:"description" example:"Input validation is not working"`
	Status            db.Bugstatus   `json:"status" example:"PENDING"`
	Priority          db.Bugpriority `json:"priority" example:"URGENT"`
	Assignedto        uuid.UUID      `json:"assigned_to" example:"00000000-0000-0000-0000-000000000000"`
	Assignedbyprofile uuid.UUID      `json:"assigned_by_profile" example:"00000000-0000-0000-0000-000000000000"`
	Assignedbyteam    uuid.UUID      `json:"assigned_by_team" example:"00000000-0000-0000-0000-000000000000"`
	Completed         bool           `json:"completed" example:"false"`
	Createdat         time.Time      `json:"created_at" example:"1710579130"`
	Updatedat         time.Time      `json:"updated_at" example:"1710579130"`
	Closedby          *uuid.UUID     `json:"closed_by" example:"00000000-0000-0000-0000-000000000000"`
	Remarks           *string        `json:"remarks" example:"None"`
	Closedat          *time.Time     `json:"closed_at" example:"1710579130"`
}

type id struct {
	Id string `json:"id" validate:"required,uuid" example:"00000000-0000-0000-0000-000000000000"`
}

type updateBugInput struct {
	Id          uuid.UUID      `json:"id" validate:"required" example:"00000000-0000-0000-0000-000000000000"`
	Name        string         `json:"name" validate:"max=32" example:"Improper Input Validation"`
	Description string         `json:"description" validate:"max=255" example:"Input validation is not working"`
	Status      db.Bugstatus   `json:"status" validate:"oneof=PENDING PROCESSING" example:"PENDING"`
	Priority    db.Bugpriority `json:"priority" validate:"oneof=URGENT HIGH LOW" example:"URGENT"`
}

type organisationOutput struct {
	ID          uuid.UUID `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Name        string    `json:"name" example:"buggy org"`
	Description string    `json:"description" example:"The organisation for bugs!"`
	Owner       uuid.UUID `json:"owner" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	CreatedAt   int64     `json:"created_at" example:"1710579130"`
	UpdatedAt   int64     `json:"updated_at" example:"1710579130"`
}

type teamOutput struct {
	Id          uuid.UUID `json:"id" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Name        string    `json:"name" example:"testing team"`
	Description string    `json:"description" example:"The Buggy testing team!"`
	CreatedAt   int64     `json:"created_at" example:"1710579130"`
	UpdatedAt   int64     `json:"updated_at" example:"1710579130"`
}

type closeBugInput struct {
	Id      uuid.UUID `json:"id" validate:"required" example:"6ba7b810-9dad-11d1-80b4-00c04fd430c8"`
	Remarks string    `json:"remarks" validate:"required,max=200" example:"This bug is done."`
}
