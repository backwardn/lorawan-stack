// Copyright © 2018 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package store

import (
	"context"

	"github.com/gogo/protobuf/types"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
)

// ApplicationStore interface for storing Applications.
//
// All functions assume the input and fieldMask to be validated, and assume
// sufficient rights to perform the action.
type ApplicationStore interface {
	CreateApplication(ctx context.Context, app *ttnpb.Application) (*ttnpb.Application, error)
	FindApplications(ctx context.Context, ids []*ttnpb.ApplicationIdentifiers, fieldMask *types.FieldMask) ([]*ttnpb.Application, error)
	GetApplication(ctx context.Context, id *ttnpb.ApplicationIdentifiers, fieldMask *types.FieldMask) (*ttnpb.Application, error)
	UpdateApplication(ctx context.Context, app *ttnpb.Application, fieldMask *types.FieldMask) (*ttnpb.Application, error)
	DeleteApplication(ctx context.Context, id *ttnpb.ApplicationIdentifiers) error
}

// ClientStore interface for storing Clients.
//
// All functions assume the input and fieldMask to be validated, and assume
// sufficient rights to perform the action.
type ClientStore interface {
	CreateClient(ctx context.Context, cli *ttnpb.Client) (*ttnpb.Client, error)
	FindClients(ctx context.Context, ids []*ttnpb.ClientIdentifiers, fieldMask *types.FieldMask) ([]*ttnpb.Client, error)
	GetClient(ctx context.Context, id *ttnpb.ClientIdentifiers, fieldMask *types.FieldMask) (*ttnpb.Client, error)
	UpdateClient(ctx context.Context, cli *ttnpb.Client, fieldMask *types.FieldMask) (*ttnpb.Client, error)
	DeleteClient(ctx context.Context, id *ttnpb.ClientIdentifiers) error
}

// GatewayStore interface for storing Gateways.
//
// All functions assume the input and fieldMask to be validated, and assume
// sufficient rights to perform the action.
type GatewayStore interface {
	CreateGateway(ctx context.Context, gtw *ttnpb.Gateway) (*ttnpb.Gateway, error)
	FindGateways(ctx context.Context, ids []*ttnpb.GatewayIdentifiers, fieldMask *types.FieldMask) ([]*ttnpb.Gateway, error)
	GetGateway(ctx context.Context, id *ttnpb.GatewayIdentifiers, fieldMask *types.FieldMask) (*ttnpb.Gateway, error)
	UpdateGateway(ctx context.Context, gtw *ttnpb.Gateway, fieldMask *types.FieldMask) (*ttnpb.Gateway, error)
	DeleteGateway(ctx context.Context, id *ttnpb.GatewayIdentifiers) error
}

// OrganizationStore interface for storing Organizations.
//
// All functions assume the input and fieldMask to be validated, and assume
// sufficient rights to perform the action.
type OrganizationStore interface {
	CreateOrganization(ctx context.Context, org *ttnpb.Organization) (*ttnpb.Organization, error)
	FindOrganizations(ctx context.Context, ids []*ttnpb.OrganizationIdentifiers, fieldMask *types.FieldMask) ([]*ttnpb.Organization, error)
	GetOrganization(ctx context.Context, id *ttnpb.OrganizationIdentifiers, fieldMask *types.FieldMask) (*ttnpb.Organization, error)
	UpdateOrganization(ctx context.Context, org *ttnpb.Organization, fieldMask *types.FieldMask) (*ttnpb.Organization, error)
	DeleteOrganization(ctx context.Context, id *ttnpb.OrganizationIdentifiers) error
}

// UserStore interface for storing Users.
//
// All functions assume the input and fieldMask to be validated, and assume
// sufficient rights to perform the action.
type UserStore interface {
	CreateUser(ctx context.Context, usr *ttnpb.User) (*ttnpb.User, error)
	FindUsers(ctx context.Context, ids []*ttnpb.UserIdentifiers, fieldMask *types.FieldMask) ([]*ttnpb.User, error)
	GetUser(ctx context.Context, id *ttnpb.UserIdentifiers, fieldMask *types.FieldMask) (*ttnpb.User, error)
	UpdateUser(ctx context.Context, usr *ttnpb.User, fieldMask *types.FieldMask) (*ttnpb.User, error)
	DeleteUser(ctx context.Context, id *ttnpb.UserIdentifiers) error
}

// UserSessionStore interface for storing User sessions.
//
// For internal use (by the OAuth server) only.
type UserSessionStore interface {
	CreateSession(ctx context.Context, sess *ttnpb.UserSession) (*ttnpb.UserSession, error)
	FindSessions(ctx context.Context, userIDs *ttnpb.UserIdentifiers) ([]*ttnpb.UserSession, error)
	GetSession(ctx context.Context, userIDs *ttnpb.UserIdentifiers, sessionID string) (*ttnpb.UserSession, error)
	UpdateSession(ctx context.Context, sess *ttnpb.UserSession) (*ttnpb.UserSession, error)
	DeleteSession(ctx context.Context, userIDs *ttnpb.UserIdentifiers, sessionID string) error
}

// MembershipStore interface for storing membership (collaboration) relations
// between accounts (users or organizations) and entities (applications, clients,
// gateways or organizations).
//
// As the operations in this store may be quite expensive, the results of FindXXX
// operations should typically be cached. The recommended cache behavior is:
//
// - The results of FindMembers are cached by (entity_type,entity_id)
// - The results of FindMemberRights are cached by (account_id,[entityType])
// - The results of FindAllMemberRights are cached by (account_id,[entityType])
// - The results of FindMemberRightsOn are cached by (account_id,entity_type,entity_id)
// - Any (successful or unsuccessful) call to SetMember expires the caches
//   for (account_id,entity_type,entity_id), (account_id,[entityType]) and (entity_type,entity_id).
type MembershipStore interface {
	// Find (direct) member of the given entity.
	FindMembers(ctx context.Context, entityID *ttnpb.EntityIdentifiers) (map[*ttnpb.OrganizationOrUserIdentifiers]*ttnpb.Rights, error)
	// Find direct member rights of the given organization or user. The entityType may be omitted.
	FindMemberRights(ctx context.Context, id *ttnpb.OrganizationOrUserIdentifiers, entityType string) (map[*ttnpb.EntityIdentifiers]*ttnpb.Rights, error)
	// Find (recursive) member rights of the given organization or user. The entityType may be omitted.
	FindAllMemberRights(ctx context.Context, id *ttnpb.OrganizationOrUserIdentifiers, entityType string) (map[*ttnpb.EntityIdentifiers]*ttnpb.Rights, error)
	// Recursively find member rights of the given organization or user on the given entity.
	FindMemberRightsOn(ctx context.Context, id *ttnpb.OrganizationOrUserIdentifiers, entityID *ttnpb.EntityIdentifiers) (*ttnpb.Rights, error)
	// Set member rights on an entity. Rights can be deleted by not passing any rights.
	SetMember(ctx context.Context, id *ttnpb.OrganizationOrUserIdentifiers, entityID *ttnpb.EntityIdentifiers, rights *ttnpb.Rights) error
}