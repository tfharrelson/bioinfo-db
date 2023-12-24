package app

type GetReactionNetworkRequest struct {
	NetworkID   string   `schema,required:"network"`
	Metabolites []string `schema:"metabolites"`
	Enzymes     []string `schema:"enzymes"`
	ReactionIDs []string `schema:"rxnIDs"`
}

type GetReactionNetworkResponse struct {
	Reactions []Reaction `json:"rxns"`
}

// Nested structures that support the main get/post requests/responses 
type Reaction struct {
	Substrates []Metabolite `json:"substrates"`
	Products   []Metabolite `json:"products"`
	ID         string       `json:"rxn_id"`
	Enzymes    []Enzyme     `json:"enzymes"`
}

type Metabolite struct {
	ID           string `json:"metabolite_id"`
	SmilesString string `json:"smiles"`
}

type Enzyme struct {
    ID string `json:"enzyme_id"`
    Name string `json:"name"`
    Sequence string `json:"sequence"`
}
