package parse

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"os"
	"testing"
	"text/template"
)

func Test_sql2go(t *testing.T) {

	content, err := os.ReadFile("model.tmpl")
	if err != nil {
		t.Fatal(err)
	}

	sql, err := os.ReadFile("result.sql")
	if err != nil {
		t.Fatal(err)
	}

	stmt, err := sqlparser.Parse(string(sql))
	if err != nil {
		t.Fatal(err)
	}
	//t.Logf("%+v", stmt)
	//for _, column := range stmt.(*sqlparser.CreateTable).Columns {
	//	t.Logf("%+v", column.Name)
	//}

	tmpl, err := template.New("test").Parse(string(content))
	if err != nil {
		t.Fatal(err)
	}

	f, err := os.OpenFile("model.py", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, stmt.(*sqlparser.CreateTable))
	if err != nil {
		t.Fatal(err)
	}

}
