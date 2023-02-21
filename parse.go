package main

import (
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"os"
	"sql2python/tmpl"
	"strings"
	"text/template"
)

const (
	TmplProperties = "properties"
	BaseFileName   = "_base"
	ToolFileName   = "_tool"
	TFileName      = "t_"
	MFileName      = "m_"
)

const ModelDirName = "model/"

var TmplMap = map[string]string{
	TmplProperties: tmpl.PropertiesString,
	BaseFileName:   tmpl.BaseString,
	ToolFileName:   tmpl.ToolString,
	TFileName:      tmpl.TString,
	MFileName:      tmpl.MString,
}

var FuncMap = map[string]interface{}{
	"i2t": IdentFirstToUpper,
	"s2t": StrFirstToUpper,
}

func HasDir(path string) (bool, error) {
	_, _err := os.Stat(path)
	if _err == nil {
		return true, nil
	}
	if os.IsNotExist(_err) {
		return false, nil
	}
	return false, _err
}

func CreateDir(dirName string) error {
	_exist, _ := HasDir(dirName)
	if !_exist {
		return os.Mkdir(dirName, 0777)
	}
	return nil
}

func constructDriName(tableName string) string {
	return fmt.Sprintf("%s%s/", ModelDirName, tableName)
}

func IdentFirstToUpper(idt sqlparser.TableIdent) string {
	str := idt.String()
	return StrFirstToUpper(str)
}

func StrFirstToUpper(str string) string {
	temp := strings.Split(str, "_")
	var upperStr string
	for y := 0; y < len(temp); y++ {
		vv := []rune(temp[y])
		for i := 0; i < len(vv); i++ {
			if i == 0 {
				vv[i] -= 32
				upperStr += string(vv[i]) // + string(vv[i+1])
			} else {
				upperStr += string(vv[i])
			}
		}
	}
	return upperStr
}

func Parse(filePath string) error {

	sql, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	stmt, err := sqlparser.Parse(string(sql))
	if err != nil {
		return err
	}

	tableName := stmt.(*sqlparser.CreateTable).DDL.NewName.Name.String()
	dirName := constructDriName(tableName)
	err = CreateDir(dirName)
	if err != nil {
		return err
	}
	err = CreateFile(tableName, TmplProperties, dirName, stmt)
	if err != nil {
		return err
	}
	err = CreateFile(tableName, BaseFileName, ModelDirName, stmt)
	if err != nil {
		return err
	}
	err = CreateFile(tableName, ToolFileName, ModelDirName, stmt)
	if err != nil {
		return err
	}
	err = CreateFile(tableName, TFileName, ModelDirName, stmt)
	if err != nil {
		return err
	}
	err = CreateFile(tableName, MFileName, ModelDirName, stmt)
	if err != nil {
		return err
	}

	return nil
}

func CreateTmpl(tmplName string) (*template.Template, error) {
	t, err := template.New(tmplName).Funcs(FuncMap).Parse(TmplMap[tmplName])
	if err != nil {
		return nil, err
	}
	return t, nil
}

func CreateFile(tableName string, tmplName string, dirName string, stmt sqlparser.Statement) error {

	t, err := CreateTmpl(tmplName)
	if err != nil {
		return err
	}
	fileName := ""

	switch tmplName {
	case TFileName:
		fileName = fmt.Sprintf("%s/%s%s.py", dirName, tmplName, tableName)
	case MFileName:
		fileName = fmt.Sprintf("%s/%s%s.py", dirName, tmplName, tableName)
		exist, err := HasDir(fileName)
		if err != nil {
			return err
		}
		if exist {
			fmt.Println("m_ 文件已经存在")
			return nil
		}
	default:
		fileName = fmt.Sprintf("%s/%s.py", dirName, tmplName)
	}

	return createFile(fileName, stmt, t)
}

func createFile(fileName string, stmt sqlparser.Statement, t *template.Template) error {

	f, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0777)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	err = t.Execute(f, stmt.(*sqlparser.CreateTable))
	if err != nil {
		return err
	}

	return nil
}
