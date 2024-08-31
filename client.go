package flair

import "golang.org/x/oauth2"

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=auth/config.yaml auth/api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=bridges/config.yaml bridges/api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=pucks/config.yaml pucks/api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=rooms/config.yaml rooms/api.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=structures/config.yaml structures/api.yaml

const APIURL = "https://api.flair.co"

var Endpoint = oauth2.Endpoint{
	AuthURL:  APIURL + "/oauth2/authorize",
	TokenURL: APIURL + "/oauth2/token",
}

const (
	// User Scope - view user names and settings

	UserViewScope  = "users.view"
	UsersEditScope = "users.edit"

	// Structures Scope - view the user's Flair home setup and configuration (includes rooms and hvac-units)

	StructuresViewScope = "structures.view"
	StructuresEditScope = "structures.edit"

	// Pucks Scope - view Pucks associated with this user

	PucksViewScope = "pucks.view"
	PucksEditScope = "pucks.edit"

	// Vents Scope - view Vents associated with this user

	VentsViewScope = "vents.view"
	VentsEditScope = "vents.edit"

	// Thermostats Scope - view smart thermostats associated with this user

	ThermostatsViewScope = "thermostats.view"
	ThermostatsEditScope = "thermostats.edit"
)
