package models

import uuid "github.com/satori/go.uuid"

type Atom struct {
	ID         string  `json:"id"`
	AtomNumber uint    `json:"number"`
	Mass       float64 `json:"mass"`
	Name       string  `json:"name"`
	Symbol     string  `json:"symbol"`
}

var Atoms = []Atom{
	{
		ID:         uuid.NewV4().String(),
		AtomNumber: 1,
		Mass:       1.01,
		Name:       "Hydrogen",
		Symbol:     "H",
	},
	{
		ID:         uuid.NewV4().String(),
		AtomNumber: 2,
		Mass:       2.02,
		Name:       "Deuterium",
		Symbol:     "D",
	},
	{
		ID:         uuid.NewV4().String(),
		AtomNumber: 3,
		Mass:       4.04,
		Name:       "Hellium",
		Symbol:     "He",
	},
	{
		ID:         uuid.NewV4().String(),
		AtomNumber: 4,
		Mass:       6.06,
		Name:       "Lithium",
		Symbol:     "Li",
	},
}
