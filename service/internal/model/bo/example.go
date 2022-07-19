package bo

import "time"

type Example struct {
	ID        string
	Name      string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type ExampleGet struct {
	ID string
}
