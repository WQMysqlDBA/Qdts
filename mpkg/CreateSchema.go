package mpkg

import (
	"Qdts/rwfile"
	"database/sql"
	"fmt"
	"strings"
)

func CreateSchema(dbconn *sql.DB, dir string) {
	//cmd1 := "ls " + dir + " | grep 'schema-create.sql'"
	cmd2 := "ls " + dir + "  | grep 'schema.sql'"
	//out1 := string(Cmd(cmd1, true, false))
	out2 := string(Cmd(cmd2, true, false))
	//dbschemafile := strings.Split(out1, "\n")
	tableschemafile := strings.Split(out2, "\n")
	//for _, v := range dbschemafile {
	//	if v != "" {
	//		schemaFile := dir + "/" + v
	//		fmt.Println(schemaFile)
	//		sql := rwfile.ReadFileAll(schemaFile)
	//		fmt.Println(sql)
	//	}
	//}
	for _, v := range tableschemafile {
		if v != "" {
			schemaFile := dir + "/" + v
			fmt.Println(schemaFile)
			sql := rwfile.ReadFileAll(schemaFile)
			fmt.Println(sql)
		}
	}
}
