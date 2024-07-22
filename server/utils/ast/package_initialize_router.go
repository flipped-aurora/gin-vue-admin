package ast

import (
	"fmt"
	"go/ast"
	"go/token"
	"io"
)

// PackageInitializeRouter 包初始化路由
// ModuleName := PackageName.AppName.GroupName
// ModuleName.FunctionName(RouterGroupName)
type PackageInitializeRouter struct {
	Base
	Type                 Type   // 类型
	Path                 string // 文件路径
	ImportPath           string // 导包路径
	RelativePath         string // 相对路径
	AppName              string // 应用名称
	GroupName            string // 分组名称
	ModuleName           string // 模块名称
	PackageName          string // 包名
	FunctionName         string // 函数名
	RouterGroupName      string // 路由分组名称
	LeftRouterGroupName  string // 左路由分组名称
	RightRouterGroupName string // 右路由分组名称
}

func (a *PackageInitializeRouter) Parse(filename string, writer io.Writer) (file *ast.File, err error) {
	if filename == "" {
		if a.RelativePath == "" {
			filename = a.Path
			a.RelativePath = a.Base.RelativePath(a.Path)
			return a.Base.Parse(filename, writer)
		}
		a.Path = a.Base.AbsolutePath(a.RelativePath)
		filename = a.Path
	}
	return a.Base.Parse(filename, writer)
}

func (a *PackageInitializeRouter) Rollback(file *ast.File) error {
	funcDecl := FindFunction(file, "initBizRouter")
	exprNum := 0
	for i := range funcDecl.Body.List {
		if IsBlockStmt(funcDecl.Body.List[i]) {
			if VariableExistsInBlock(funcDecl.Body.List[i].(*ast.BlockStmt), a.ModuleName) {
				for ii, stmt := range funcDecl.Body.List[i].(*ast.BlockStmt).List {
					// 检查语句是否为 *ast.ExprStmt
					exprStmt, ok := stmt.(*ast.ExprStmt)
					if !ok {
						continue
					}
					// 检查表达式是否为 *ast.CallExpr
					callExpr, ok := exprStmt.X.(*ast.CallExpr)
					if !ok {
						continue
					}
					// 检查是否调用了我们正在寻找的函数
					selExpr, ok := callExpr.Fun.(*ast.SelectorExpr)
					if !ok {
						continue
					}
					// 检查调用的函数是否为 systemRouter.InitApiRouter
					ident, ok := selExpr.X.(*ast.Ident)
					//只要存在调用则+1
					if ok && ident.Name == a.ModuleName {
						exprNum++
					}
					//判断是否为目标结构
					if !ok || ident.Name != a.ModuleName || selExpr.Sel.Name != a.FunctionName {
						continue
					}
					exprNum--
					// 从语句列表中移除。
					funcDecl.Body.List[i].(*ast.BlockStmt).List = append(funcDecl.Body.List[i].(*ast.BlockStmt).List[:ii], funcDecl.Body.List[i].(*ast.BlockStmt).List[ii+1:]...)
					// 如果不再存在任何调用，则删除导入和变量。
					if exprNum == 0 {
						funcDecl.Body.List = append(funcDecl.Body.List[:i], funcDecl.Body.List[i+1:]...)
					}
					break
				}
				break
			}
		}
	}

	return nil
}

func (a *PackageInitializeRouter) Injection(file *ast.File) error {
	funcDecl := FindFunction(file, "initBizRouter")
	hasRouter := false
	var varBlock *ast.BlockStmt
	for i := range funcDecl.Body.List {
		if IsBlockStmt(funcDecl.Body.List[i]) {
			if VariableExistsInBlock(funcDecl.Body.List[i].(*ast.BlockStmt), a.ModuleName) {
				hasRouter = true
				varBlock = funcDecl.Body.List[i].(*ast.BlockStmt)
				break
			}
		}
	}
	if !hasRouter {
		stmt := a.CreateAssignStmt()
		varBlock = &ast.BlockStmt{
			List: []ast.Stmt{
				stmt,
			},
		}
	}
	routerStmt := CreateStmt(fmt.Sprintf("%s.%s(%s,%s)", a.ModuleName, a.FunctionName, a.LeftRouterGroupName, a.RightRouterGroupName))
	varBlock.List = append(varBlock.List, routerStmt)
	if !hasRouter {
		funcDecl.Body.List = append(funcDecl.Body.List, varBlock)
	}
	return nil
}

func (a *PackageInitializeRouter) Format(filename string, writer io.Writer, file *ast.File) error {
	if filename == "" {
		filename = a.Path
	}
	return a.Base.Format(filename, writer, file)
}

func (a *PackageInitializeRouter) CreateAssignStmt() *ast.AssignStmt {
	//创建左侧变量
	ident := &ast.Ident{
		Name: a.ModuleName,
	}

	//创建右侧的赋值语句
	selector := &ast.SelectorExpr{
		X: &ast.SelectorExpr{
			X:   &ast.Ident{Name: a.PackageName},
			Sel: &ast.Ident{Name: a.AppName},
		},
		Sel: &ast.Ident{Name: a.GroupName},
	}

	// 创建一个组合的赋值语句
	stmt := &ast.AssignStmt{
		Lhs: []ast.Expr{ident},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{selector},
	}

	return stmt
}
