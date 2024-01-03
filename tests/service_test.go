package tests

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/tfharrelson/bioinfo-db/internal/app"
)

var a app.App

func TestMain(m *testing.M) {
	a.Initialize()
	code := m.Run()
	os.Exit(code)
}

func TestFindReactions(t *testing.T) {
	getRequest := app.GetReactionNetworkRequest{
		NetworkID: "1",
	}
	rxns, err := app.FindReactions(&getRequest)
	if err != nil {
		t.Fatalf(`Error encountered when finding reaction. This should not happen. Error = %q`, err)
	}
	if rxns[0].ID != "R00014" {
		t.Fatalf(`Unsuspected reaction found! found %q when "R00014" is desired`, rxns[0].ID)
	}
}

// Test that we can spin up a server without crashing
func TestHealthCheck(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/health-check", nil)
	response := executeRequest(request)

    // make sure we get a 200 ok
    checkResponseCode(t, http.StatusOK, response.Code)
}

func TestDataRequest(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "/api/data", nil)
	response := executeRequest(request)

    // make sure we get a 200 ok
    checkResponseCode(t, http.StatusOK, response.Code)

    // make sure that the response body contains what we think it should
    // first convert the json into a map
    var rxnStruct app.GetReactionNetworkResponse
    json.Unmarshal(response.Body.Bytes(), &rxnStruct)

    // now let's see if we have a list of one reaction
    rxns := rxnStruct.Reactions
    if len(rxns) != 1 {
        t.Errorf("Found too many reactions, expected 1 and found %d", len(rxns))
    }
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	response := httptest.NewRecorder()

	a.Router.ServeHTTP(response, req)

	return response
}

func checkResponseCode(t *testing.T, expected, actual int) {
    if expected != actual {
        t.Errorf("Expected response code %d. Got %d\n", expected, actual)
    }
}
