// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	CompleteOrganisationTransfer(ctx context.Context, arg CompleteOrganisationTransferParams) (Organisationtransfer, error)
	CreateOrganisation(ctx context.Context, arg CreateOrganisationParams) (Organisation, error)
	CreateOrganisationTransfer(ctx context.Context, arg CreateOrganisationTransferParams) (Organisationtransfer, error)
	CreateProfile(ctx context.Context, arg CreateProfileParams) (Profile, error)
	CreateTeam(ctx context.Context, arg CreateTeamParams) (Team, error)
	CreateTeamMember(ctx context.Context, arg CreateTeamMemberParams) (Teammember, error)
	DeleteOrganisationTransfer(ctx context.Context, id uuid.UUID) (Organisationtransfer, error)
	GetActiveOrganisationTransfer(ctx context.Context, organisation uuid.UUID) (Organisationtransfer, error)
	GetAllTeamMembers(ctx context.Context, team uuid.UUID) ([]Profile, error)
	GetIncomingOrganisationTransfers(ctx context.Context, toprofile uuid.UUID) ([]GetIncomingOrganisationTransfersRow, error)
	GetOrganisation(ctx context.Context, owner uuid.UUID) (Organisation, error)
	GetOrganisationTeams(ctx context.Context, organisation uuid.UUID) ([]Team, error)
	GetOutgoingOrganisationTransfers(ctx context.Context, fromprofile uuid.UUID) ([]Organisationtransfer, error)
	GetProfile(ctx context.Context, id uuid.UUID) (Profile, error)
	GetProfileByEmail(ctx context.Context, email string) (Profile, error)
	GetTeamMember(ctx context.Context, arg GetTeamMemberParams) (Teammember, error)
	GetTeamOrganisation(ctx context.Context, id uuid.UUID) (uuid.UUID, error)
	UpdateOrganisation(ctx context.Context, arg UpdateOrganisationParams) (Organisation, error)
	UpdateOrganisationOwner(ctx context.Context, arg UpdateOrganisationOwnerParams) (Organisation, error)
	UpdatePassword(ctx context.Context, arg UpdatePasswordParams) (Profile, error)
	UpdateProfile(ctx context.Context, arg UpdateProfileParams) (Profile, error)
	UpdateTeam(ctx context.Context, arg UpdateTeamParams) (Team, error)
	UpdateTokenId(ctx context.Context, arg UpdateTokenIdParams) (Profile, error)
	UpdateTokenIdByEmail(ctx context.Context, arg UpdateTokenIdByEmailParams) (Profile, error)
	VerifyProfile(ctx context.Context, id uuid.UUID) (Profile, error)
}

var _ Querier = (*Queries)(nil)
