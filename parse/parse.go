package parse

import (
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"os"
	"text/template"
)

func Parse(filePath string) {

	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	println(dir)

	content, err := os.ReadFile("tmpl/model.tmpl")
	if err != nil {
		panic(err)
	}

	sql, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	stmt, err := sqlparser.Parse(string(sql))
	if err != nil {
		panic(err)
	}

	tmpl, err := template.New("test").Parse(string(content))
	if err != nil {
		panic(err)
	}

	f, err := os.OpenFile("model.py", os.O_CREATE|os.O_WRONLY, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = tmpl.Execute(f, stmt.(*sqlparser.CreateTable))
	if err != nil {
		panic(err)
	}
}
