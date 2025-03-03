package app

type MoleculeID string

type Molecule struct {
	// TODO: find the best id format... what does neo4j recommend
	ID    MoleculeID
	Name  string
	Atoms []Atom
	Bonds []Bond
}

type ChiralityVariant int

const (
	L ChiralityVariant = iota + 1
	R
)

type Atom struct {
	Symbol    string
	Chirality ChiralityVariant
}

type BondVariant int

const (
	Single BondVariant = iota + 1
	Double
	Triple
	Aromatic
)

type Bond struct {
	FirstIndex  int
	SecondIndex int
	Order       BondVariant
}

// higher level models (aka coarse grained concepts of molecules)
// like proteins, DNA, etc.
// at this level we don't care about their specific classification systems
// e.g. no domain models for DNA specifically.
type SuperMoleculeID string

type SuperMolecule struct {
	ID        SuperMoleculeID
	Name      string
	Molecules []MoleculeID
	Bonds     []Bond
}
