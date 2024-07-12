package ast

import (
	"github.com/pkg/errors"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"os"
)

type Api struct {
	Path        string // api路径
	Group       string // api分组
	Method      string // api方法
	Description string // api描述
}

type PluginInitializeApi struct {
	Type     Type           // 类型
	Path     string         // 文件路径
	Apis     []Api          // api列表
	ApiMap   map[string]Api // api列表
	ApiIndex map[string]int // api索引
}

func NewPluginInitializeApi(astType Type, path string, apis []Api) *PluginInitializeApi {
	entity := &PluginInitializeApi{Type: astType, Path: path, Apis: apis}
	entity.ApiMap = make(map[string]Api, len(apis))
	entity.ApiIndex = make(map[string]int, len(apis))
	for i := 0; i < len(entity.Apis); i++ {
		key := apis[i].Path + apis[i].Method
		entity.ApiMap[key] = apis[i]
		entity.ApiIndex[key] = i
	}
	return entity
}

func (a *PluginInitializeApi) Rollback() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name != "Api" {
				continue
			}
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					if len(v2.Lhs) > 1 {
						v3, o3 := v2.Lhs[0].(*ast.Ident)
						if o3 {
							if v3.Name != "entities" {
								break
							} // 判断局部变量名是否为entities
						}
					}
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CompositeLit)
						if o3 {
							remove := make([]int, 0, 5)
							length := len(v3.Elts)
							for l := 0; l < length; l++ {
								v4, o4 := v3.Elts[l].(*ast.CompositeLit)
								if o4 {
									var entity Api
									for m := 0; m < len(v4.Elts); m++ {
										v5, o5 := v4.Elts[m].(*ast.KeyValueExpr)
										if o5 {
											v6, o6 := v5.Key.(*ast.Ident)
											v7, o7 := v5.Value.(*ast.BasicLit)
											if o6 && o7 {
												switch v6.Name {
												case "Path":
													entity.Path = v7.Value
												case "Method":
													entity.Method = v7.Value
												case "ApiGroup":
													entity.Group = v7.Value
												case "Description":
													entity.Description = v7.Value
												}
											}
										}
									}
									key := entity.Path + entity.Method
									_, ok := a.ApiMap[key]
									if ok {
										remove = append(remove, l)
									}
								}
							}
							for l := len(remove) - 1; l >= 0; l-- {
								index := remove[l]
								v3.Elts = append(v3.Elts[:index], v3.Elts[index+1:]...)
								if l == 0 {
									if remove[len(remove)-1] >= length {
										continue
									} // 说明回滚的是最新生成的代码, 所以无需处理空行
									// TODO 如果回滚中间的代码, 则需要处理空行
								}
							}
						}
					}
				}
			}
		}
	}
	create, err := os.Create(a.Path)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	defer func() {
		_ = create.Close()
	}()
	err = format.Node(create, fileSet, file)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]注入失败!", a.Path)
	}
	return nil
}

func (a *PluginInitializeApi) Injection() error {
	fileSet := token.NewFileSet()
	file, err := parser.ParseFile(fileSet, a.Path, nil, parser.ParseComments)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	for i := 0; i < len(file.Decls); i++ {
		v1, o1 := file.Decls[i].(*ast.FuncDecl)
		if o1 {
			if v1.Name.Name != "Api" {
				continue
			}
			for j := 0; j < len(v1.Body.List); j++ {
				v2, o2 := v1.Body.List[j].(*ast.AssignStmt)
				if o2 {
					if len(v2.Lhs) > 1 {
						v3, o3 := v2.Lhs[0].(*ast.Ident)
						if o3 {
							if v3.Name != "entities" {
								break
							} // 判断局部变量名是否为entities
						}
					}
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CompositeLit)
						if o3 {
							for l := 0; l < len(v3.Elts); l++ {
								v4, o4 := v3.Elts[l].(*ast.CompositeLit)
								if o4 {
									var entity Api
									for m := 0; m < len(v4.Elts); m++ {
										v5, o5 := v4.Elts[m].(*ast.KeyValueExpr)
										if o5 {
											v6, o6 := v5.Key.(*ast.Ident)
											v7, o7 := v5.Value.(*ast.BasicLit)
											if o6 && o7 {
												switch v6.Name {
												case "Path":
													entity.Path = v7.Value
												case "Method":
													entity.Method = v7.Value
												case "ApiGroup":
													entity.Group = v7.Value
												case "Description":
													entity.Description = v7.Value
												}
											}
										}
									}
									key := entity.Path + entity.Method
									_, ok := a.ApiMap[key]
									if ok {
										index := a.ApiIndex[key]
										a.Apis = append(a.Apis[:index], a.Apis[index+1:]...)
									} // 删除已存在的api
								}
							}
						}
					}
					length := len(a.Apis)
					composites := make([]ast.Expr, 0, length)
					for k := 0; k < len(a.Apis); k++ {
						composite := &ast.CompositeLit{
							Elts: []ast.Expr{
								&ast.KeyValueExpr{
									Key: &ast.Ident{Name: "Path"},
									Value: &ast.BasicLit{
										Kind:  token.STRING,
										Value: a.Apis[k].Path,
									},
								},
								&ast.KeyValueExpr{
									Key: &ast.Ident{Name: "Method"},
									Value: &ast.BasicLit{
										Kind:  token.STRING,
										Value: a.Apis[k].Method,
									},
								},
								&ast.KeyValueExpr{
									Key: &ast.Ident{Name: "ApiGroup"},
									Value: &ast.BasicLit{
										Kind:  token.STRING,
										Value: a.Apis[k].Group,
									},
								},
								&ast.KeyValueExpr{
									Key: &ast.Ident{Name: "Description"},
									Value: &ast.BasicLit{
										Kind:  token.STRING,
										Value: a.Apis[k].Description,
									},
								},
							},
						}
						composites = append(composites, composite)
					}
					for k := 0; k < len(v2.Rhs); k++ {
						v3, o3 := v2.Rhs[k].(*ast.CompositeLit)
						if o3 {
							length = len(v3.Elts)
							v3.Elts = append(v3.Elts, composites...)
							for l := length; l < len(v3.Elts); l++ {
								v4, o4 := v3.Elts[l].(*ast.CompositeLit)
								if o4 {
									v5, o5 := v3.Elts[l-1].(*ast.CompositeLit)
									if o5 {
										tokenLbracePosition := fileSet.Position(v5.Lbrace)
										tokenFileSet := fileSet.File(v5.Lbrace)
										tokenNextLbracePos := tokenFileSet.LineStart(tokenLbracePosition.Line + 1)
										tokenNextLbracePosition := fileSet.Position(tokenNextLbracePos)
										tokenNextLbracePosition.Column = tokenLbracePosition.Column
										v4.Lbrace = tokenFileSet.Pos(tokenNextLbracePosition.Offset)
										if l == len(v3.Elts)-1 {
											v3.Elts = append(v3.Elts, &ast.BasicLit{
												Kind:  token.STRING,
												Value: "\n",
											})
										} // 最后一个元素l == len(v3.Elts)-1 添加回车
									}
								}
							}
						} // 注入Apis
					}
				}
			}
		}
	}
	create, err := os.Create(a.Path)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]打开文件失败!", a.Path)
	}
	defer func() {
		_ = create.Close()
	}()
	err = format.Node(create, fileSet, file)
	if err != nil {
		return errors.Wrapf(err, "[filepath:%s]注入失败!", a.Path)
	}
	return nil
}
