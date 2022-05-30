package internal

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path"
	"strings"
)

type Table struct {
	TableName          string
	StructName         string
	TableComment       string
	FunctionName       string
	FunctionReturnType string
}

// GetTablesComment ast解析注释得到表的注释信息
// Author [SliverHorn](https://github.com/SliverHorn)
func GetTablesComment() (map[string]Table, error) {
	byteMap := make(map[string][]byte, 100)
	dirs, err := os.ReadDir(global.GVA_CONFIG.AutoCode.Model())
	if err != nil {
		return nil, err
	}
	length := len(dirs)
	for i := 0; i < length; i++ {
		if dirs[i].IsDir() {
			var secondDirs []os.DirEntry
			dirPath := path.Join(global.GVA_CONFIG.AutoCode.Model(), dirs[i].Name())
			secondDirs, err = os.ReadDir(dirPath)
			for j := 0; j < len(secondDirs); j++ {
				if !secondDirs[j].IsDir() {
					var bytes []byte
					secondDirPath := path.Join("model", dirs[i].Name(), secondDirs[j].Name())
					bytes, err = os.ReadFile(secondDirPath)
					if err != nil {
						return nil, err
					}
					byteMap[secondDirPath] = bytes
				}
			}
			if err != nil {
				return nil, err
			}
		}
	}
	length = len(byteMap)
	tableMap := make(map[string]Table, length)
	for key, values := range byteMap {
		fileSet := token.NewFileSet()
		var parseFile *ast.File
		parseFile, err = parser.ParseFile(fileSet, key, values, parser.ParseComments)
		if err != nil {
			return nil, err
		}

		var table Table
		for i := 0; i < len(parseFile.Decls); i++ {
			if v1, o1 := parseFile.Decls[i].(*ast.FuncDecl); o1 {
				if v1.Name.Name != "TableName" {
					continue
				}
				table.FunctionName = v1.Name.Name
				for j := 0; j < len(v1.Recv.List); j++ {
					if v2, o2 := v1.Recv.List[j].Type.(*ast.Ident); o2 {
						table.StructName = v2.Name
						continue
					}
					if v3, o3 := v1.Recv.List[j].Type.(*ast.StarExpr); o3 {
						if v4, o4 := v3.X.(*ast.Ident); o4 {
							table.StructName = v4.Name
						}
					}
				}
				for j := 0; j < len(v1.Type.Results.List); j++ {
					table.FunctionReturnType = v1.Type.Results.List[j].Type.(*ast.Ident).Name
					continue
				}
				for j := 0; j < len(v1.Body.List); j++ {
					if v2, o2 := v1.Body.List[j].(*ast.ReturnStmt); o2 {
						for k := 0; k < len(v2.Results); k++ {
							if v3, o3 := v2.Results[k].(*ast.BasicLit); o3 {
								name := strings.Replace(v3.Value, "\"", "", -1)
								table.TableName = name
								continue
							}
						}
					}
				}

			}
		}
		for i := 0; i < len(parseFile.Comments); i++ {
			for j := 0; j < len(parseFile.Comments[i].List); j++ {
				if strings.Contains(parseFile.Comments[i].List[j].Text, "TableName") {
					texts := strings.Split(parseFile.Comments[i].List[j].Text, " ")
					if len(texts) == 3 {
						comment := strings.Replace(texts[2], "\"", "", -1)
						table.TableComment = comment
					}
				}
			}
		}
		if table.StructName != "" && table.TableComment != "" {
			tableMap[table.StructName] = table
		}
	}
	return tableMap, nil
}
