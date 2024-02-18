package api

import (
	"os"
	"testing"

	"github.com/papaya147/buggy/backend/api/organisation"
	"github.com/papaya147/buggy/backend/api/profile"
	"github.com/papaya147/buggy/backend/api/team"
)

var testApp *server

func TestMain(m *testing.M) {
	testApp = &server{}

	testApp.profileHandler = profile.NewHandler(nil, nil, nil)
	testApp.organisationHandler = organisation.NewHandler(nil, nil, nil)
	testApp.teamHandler = team.NewHandler(nil, nil, nil)
	testApp.router = testApp.routes()

	os.Exit(m.Run())
}
