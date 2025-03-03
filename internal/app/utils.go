package app

import (
	"fmt"
	"strings"
)

func FromSMILES(smiles string) Molecule {
  atoms := []Atom {}
  bonds := []Bond {}
  var prevAtom *Atom
  var prevBond *Bond
  atomIndex := 0
	index := 0
	for index < len(smiles) {
    // TODO: refactor to use runes for speed
    smilesChar := smiles[index: index + 1]
		if smilesChar == "(" {
      // TODO: convert all the panics to proper error handling
      if prevAtom == nil {
        panic("Cannot create a branch of no previous atom exists")
      }
			// find inner molecule within branch
			lastParenIndex := strings.LastIndex(smiles, ")")
      index = index + lastParenIndex + 1
			branchMolecule := FromSMILES(smiles[index : lastParenIndex+1])
      atoms = append(atoms, branchMolecule.Atoms...)
      bonds = append(bonds, Bond{FirstIndex: atomIndex, SecondIndex: atomIndex + 1, Order: Single})
      atomIndex = atomIndex + len(branchMolecule.Atoms)
    // TODO: add constructs for Cl and Br
    } else if strings.Contains("BCNOPSFI", smilesChar) {
      if prevBond == nil {
        bonds = append(bonds, Bond{FirstIndex: atomIndex, SecondIndex: atomIndex + 1, Order: Single})
      } else {
        bonds = append(bonds, *prevBond)
        prevBond = nil
      }
      atoms = append(atoms, Atom{Symbol: smilesChar, Chirality: R})
      atomIndex = atomIndex + 1
      index = index + 1
    // TODO: are there other bond types?
    } else if smilesChar == "=" {
      // we've found a new double bond
      prevBond = &Bond{FirstIndex: atomIndex, SecondIndex: atomIndex + 1, Order: Double}
      index = index + 1
    } else if smilesChar == "#" {
      // we've found a new double bond
      prevBond = &Bond{FirstIndex: atomIndex, SecondIndex: atomIndex + 1, Order: Triple}
      index = index + 1
    } else {
      panic("Unknown smiles character, " + smilesChar + ", encountered!")
    }
	}
  // TODO: generate new id's from incrementing index of database
  //  or separate the database ID from the molecule structure
  return Molecule{ID: "1", Name: "", Atoms: atoms, Bonds: bonds}
}
