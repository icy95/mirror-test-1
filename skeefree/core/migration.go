package core

import (
	"fmt"
	"strings"
)

type Migration struct {
	Schema string
	Table string
	PRStatement string
	Type string
}

func NewMigration(schema, prStatement string) *Migration {
	return &Migration{
		Schema: schema,
		PRStatement: prStatement,
	}
}

func (migration *Migration) Evaluate() error {
	migration.Type = string(migration.GetMigrationType())

	migrationEvalMap := map[MigrationType](func() error){
		CreateTableMigrationType:   migration.evalCreateTable,
		DropTableMigrationType:     migration.evalDropTable,
		AlterTableMigrationType:    migration.evalAlterTable,
	}
	if f, ok := migrationEvalMap[migration.GetMigrationType()]; ok {
		return f()
	}
	return fmt.Errorf("Evaluation unsupported for %v", migration.GetMigrationType())
}

func GetMigrations(info *SkeemaDiffInfo) []Migration {
	var migrations []Migration
	for _ ,s := range info.Statements {
		migrations = append(migrations, Migration{Schema: info.SchemaName, PRStatement: s})
	}
	return migrations
}

func (migration *Migration) evalCreateTable() error {
	tableName, err := dissectCreateTableStatement(migration.PRStatement)
	if err != nil {
		return err
	}
	migration.Table = tableName
	return nil
}
//TODO RENAME AND DROP
func (migration *Migration) evalDropTable() error {
	tableName, err := dissectDropTableStatement(migration.PRStatement)
	if err != nil {
		return err
	}
	migration.Table = tableName
	return nil
}

func (migration *Migration) evalAlterTable() error {
	tableName, alter, err := dissectAlterTableStatement(migration.PRStatement)
	if err != nil {
		return err
	}
	alter = strings.TrimRight(alter, ";")
	migration.Table = tableName
	return nil
}