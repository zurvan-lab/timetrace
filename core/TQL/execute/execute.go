package execute

import "github.com/zurvan-lab/TimeTrace/core/database"

type Executor func(database.IDataBase, []string) string

type ExecutorMap map[string]Executor

var Executors ExecutorMap = ExecutorMap{
	"SET":   database.IDataBase.AddSet,
	"SSET":  database.IDataBase.AddSubSet,
	"PUSH":  database.IDataBase.PushElement,
	"DRPS":  database.IDataBase.DropSet,
	"DRPSS": database.IDataBase.DropSubSet,
	"CLN":   database.IDataBase.CleanSets,
	"CLNS":  database.IDataBase.CleanSet,
	"CLNSS": database.IDataBase.CleanSubSet,
}

func Execute(q database.Query, db database.IDataBase) string {
	execute, ok := Executors[q.Command]
	if !ok {
		return "INVALID"
	}

	result := execute(db, q.Args)

	return result
}
