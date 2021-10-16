package migrations

import "github.com/joaomarcuslf/sucellus/definitions"

var list = []definitions.Migration{
	{
		Name:           "start_Database",
		Implementation: StartDatabase,
	},
}

func GetList() []definitions.Migration {
	return list
}
