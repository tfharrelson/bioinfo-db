package app

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"net/http"
)

// Constants for use in database validation
type AllowedDatabase string

const (
	KeggDB   = AllowedDatabase("kegg")
	EntrezDB = AllowedDatabase("entrez")
)

type WebAppRoute struct {
	route   string
	handler http.Handler
}

func NewRoute(route string, handler http.Handler) WebAppRoute {
	// TODO: add some middleware. it looks like it may be useful as I
	// add some boilerplate middleware to every route
	return WebAppRoute{route: route, handler: handler}
}

// define the custom route structs
func CreateRoutes(router *mux.Router, driver *neo4j.DriverWithContext) *mux.Router {

	router.HandleFunc("/health-check", http.HandlerFunc(HealthCheckHandler))

	router.HandleFunc(
		"/api/data",
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { DataQueryHandler(w, r, driver) }),
	)

	return router
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func DataQueryHandler(w http.ResponseWriter, r *http.Request, driver *neo4j.DriverWithContext) {
	// decode the request
	dataRequest := &GetReactionNetworkRequest{}
	err := schema.NewDecoder().Decode(dataRequest, r.URL.Query())

	// return if we done goofed in the request
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("encountered an error while parsing the request: " + err.Error()))
	}

	// query the database for reactions
	reactions, err := FindReactions(dataRequest, driver)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong while fetching reaction data. Error message: " + err.Error()))
	}

	// write the response
	response := GetReactionNetworkResponse{Reactions: reactions}
	responseString, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Something went wrong with the response from the DB..."))
	}
	w.Write(responseString)
}

// let's hard code a reaction for now
// TODO: hook up a database to allow for dynamically querying different reactions based on the input requests
func FindReactions(req *GetReactionNetworkRequest, driver *neo4j.DriverWithContext) ([]Reaction, error) {

	// define the substrates
	pyr := Metabolite{ID: "C00032", SmilesString: "CC(=O)C(=O)[O-]"}
	thi := Metabolite{ID: "C00068", SmilesString: "CC1=C(SC=[N+]1CC2=CN=C(N=C2N)C)CCOP(=O)(O)OP(=O)(O)O"}

	// define the products
	hoetThia := Metabolite{ID: "C05125", SmilesString: "CC1=C(SC(=[N+]1CC2=CN=C(N=C2N)C)C(C)O)CCOP(=O)(O)OP(=O)(O)O"}
	co2 := Metabolite{ID: "C00011", SmilesString: "C(=O)=O"}

	// define the enzymes
	pyrDehy := Enzyme{ID: "1.2.4.1", Name: "pyruvate dehydrogenase"}
	aceSynth := Enzyme{ID: "2.2.1.6", Name: "acetolactate synthase"}
	pyrDecarb := Enzyme{ID: "4.1.1.1", Name: "pyruvate decarboxylase"}

	rxn := Reaction{
		Substrates: []Metabolite{pyr, thi},
		Products:   []Metabolite{hoetThia, co2},
		ID:         "R00014",
		Enzymes:    []Enzyme{pyrDehy, aceSynth, pyrDecarb},
	}
	return []Reaction{rxn}, nil
}
