package core

import (
	"strings"
)

type MigrationType string

const (
	CreateTableMigrationType   MigrationType = "CreateTableMigrationType"
	DropTableMigrationType     MigrationType = "DropTableMigrationType"
	AlterTableMigrationType    MigrationType = "AlterTableMigrationType"
	AlterDatabaseMigrationType MigrationType = "AlterDatabaseMigrationType"
	UnsupportedMigrationType   MigrationType = "UnsupportedMigrationType"
)

func (migration *Migration) GetMigrationType() MigrationType {
	if strings.HasPrefix(migration.PRStatement, "CREATE TABLE") {
		return CreateTableMigrationType
	}
	if strings.HasPrefix(migration.PRStatement, "DROP TABLE") {
		return DropTableMigrationType
	}
	if strings.HasPrefix(migration.PRStatement, "RENAME TABLE") && strings.HasSuffix(migration.PRStatement, "_DRP") {
		return DropTableMigrationType
	}
	if strings.HasPrefix(migration.PRStatement, "ALTER TABLE") {
		return AlterTableMigrationType
	}
	if strings.HasPrefix(migration.PRStatement, "ALTER DATABASE") {
		return AlterDatabaseMigrationType
	}
	return UnsupportedMigrationType
}