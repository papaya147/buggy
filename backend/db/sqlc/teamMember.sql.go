// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: teamMember.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const createTeamMember = `-- name: CreateTeamMember :one
INSERT INTO teamMember (team, profile, admin)
VALUES ($1, $2, $3)
RETURNING team, profile, admin, createdat, updatedat
`

type CreateTeamMemberParams struct {
	Team    uuid.UUID `json:"team"`
	Profile uuid.UUID `json:"profile"`
	Admin   bool      `json:"admin"`
}

func (q *Queries) CreateTeamMember(ctx context.Context, arg CreateTeamMemberParams) (Teammember, error) {
	row := q.db.QueryRow(ctx, createTeamMember, arg.Team, arg.Profile, arg.Admin)
	var i Teammember
	err := row.Scan(
		&i.Team,
		&i.Profile,
		&i.Admin,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}

const getAllTeamMembers = `-- name: GetAllTeamMembers :many
SELECT p.id, p.tokenid, p.name, p.email, p.password, p.verified, p.createdat, p.updatedat
FROM teamMember tm
    INNER JOIN profile p ON tm.profile = p.id
WHERE tm.team = $1
`

func (q *Queries) GetAllTeamMembers(ctx context.Context, team uuid.UUID) ([]Profile, error) {
	rows, err := q.db.Query(ctx, getAllTeamMembers, team)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Profile{}
	for rows.Next() {
		var i Profile
		if err := rows.Scan(
			&i.ID,
			&i.Tokenid,
			&i.Name,
			&i.Email,
			&i.Password,
			&i.Verified,
			&i.Createdat,
			&i.Updatedat,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getTeamMember = `-- name: GetTeamMember :one
SELECT team, profile, admin, createdat, updatedat
FROM teamMember
WHERE team = $1
    AND profile = $2
LIMIT 1
`

type GetTeamMemberParams struct {
	Team    uuid.UUID `json:"team"`
	Profile uuid.UUID `json:"profile"`
}

func (q *Queries) GetTeamMember(ctx context.Context, arg GetTeamMemberParams) (Teammember, error) {
	row := q.db.QueryRow(ctx, getTeamMember, arg.Team, arg.Profile)
	var i Teammember
	err := row.Scan(
		&i.Team,
		&i.Profile,
		&i.Admin,
		&i.Createdat,
		&i.Updatedat,
	)
	return i, err
}
