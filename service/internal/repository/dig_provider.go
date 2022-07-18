package repository

import "go.uber.org/dig"

func NewRepository(in repositoryIn) repositoryOut {
	self := &repository{
		in: in,

		repositoryOut: repositoryOut{
			ExampleDao: newExampleDao(in),
		},
	}

	return self.repositoryOut
}

type repositoryIn struct {
	dig.In
}

type repository struct {
	in repositoryIn

	repositoryOut
}

type repositoryOut struct {
	dig.Out

	ExampleDao IExampleDao
}
