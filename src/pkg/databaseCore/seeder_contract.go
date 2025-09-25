package databaseCore

type Seeder interface {
	GetName() string
	Run() error
}
