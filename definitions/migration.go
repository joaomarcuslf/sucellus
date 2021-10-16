package definitions

type Migration struct {
	Name           string
	Implementation func(DatabaseClient)
}
