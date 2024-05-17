package addons

/*
AddOnName represents the name of an add-on.
It is used for debugging and identification purposes.
*/
type AddOnName string

const (
	GinServerName  AddOnName = "ginServer"
	PgDatabaseName AddOnName = "postgres"
)

/*
AddOnPrefix represents the prefix used to track an add-on.
Prefixes are formatted in dash-case for consistency.
*/
type AddOnPrefix string

const (
	GinServerPrefix  AddOnPrefix = "gin"
	PgDatabasePrefix AddOnPrefix = "pg"
)
