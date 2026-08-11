package autocode

import (
	"go/format"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"text/template"

	systemReq "github.com/flipped-aurora/gin-vue-admin/server/model/system/request"
)

func renderReservedColumnService(t *testing.T, templateKind string, isAdd bool) string {
	t.Helper()

	_, currentFile, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller() failed")
	}
	repoRoot := filepath.Clean(filepath.Join(filepath.Dir(currentFile), "..", "..", ".."))
	templatePath := filepath.Join(repoRoot, "server", "resource", templateKind, "server", "service", "service.go.tpl")

	serviceTemplate, err := template.New(filepath.Base(templatePath)).Funcs(GetTemplateFuncMap()).ParseFiles(templatePath)
	if err != nil {
		t.Fatalf("ParseFiles(%q) error = %v", templatePath, err)
	}

	primaryField := &systemReq.AutoCodeField{
		FieldName:  "ID",
		FieldJson:  "id",
		ColumnName: "id",
	}
	reservedField := &systemReq.AutoCodeField{
		FieldName:       "Group",
		FieldJson:       "group",
		ColumnName:      "group",
		FieldType:       "string",
		FieldSearchType: "LIKE",
		Sort:            true,
	}
	info := systemReq.AutoCode{
		Package:         "demo",
		PackageName:     "record",
		StructName:      "Record",
		Description:     "record",
		Abbreviation:    "record",
		HumpPackageName: "record",
		Module:          "example.com/project/server",
		PrimaryField:    primaryField,
		Fields:          []*systemReq.AutoCodeField{reservedField},
		NeedSort:        true,
		GenerateServer:  true,
		IsAdd:           isAdd,
	}

	var rendered strings.Builder
	if err = serviceTemplate.Execute(&rendered, info); err != nil {
		t.Fatalf("Execute(%q) error = %v", templatePath, err)
	}
	return rendered.String()
}

func TestGeneratedServicesQuoteReservedSearchAndSortColumns(t *testing.T) {
	for _, templateKind := range []string{"package", "plugin"} {
		t.Run(templateKind, func(t *testing.T) {
			rendered := renderReservedColumnService(t, templateKind, false)

			if !strings.Contains(rendered, `clause.Column{Name: "group"}`) {
				t.Fatalf("generated service leaves the reserved search column unquoted:\n%s", rendered)
			}
			if !strings.Contains(rendered, "db.Order(clause.OrderByColumn{") {
				t.Fatalf("generated service leaves the reserved sort column unquoted:\n%s", rendered)
			}
			if !strings.Contains(rendered, "Column: clause.Column{Name: info.Sort}") {
				t.Fatalf("generated service does not bind the requested sort field as a quoted column:\n%s", rendered)
			}
			if _, err := format.Source([]byte(rendered)); err != nil {
				t.Fatalf("generated service is not valid Go source: %v\n%s", err, rendered)
			}

			additions := renderReservedColumnService(t, templateKind, true)
			if !strings.Contains(additions, `import "gorm.io/gorm/clause"`) {
				t.Fatalf("incremental service snippet omits its clause import:\n%s", additions)
			}
		})
	}
}

func TestGenerateSearchConditionsQuotesColumnsForEveryOperatorBranch(t *testing.T) {
	tests := []struct {
		name      string
		fieldType string
		operator  string
		wantSQL   string
	}{
		{name: "enum like", fieldType: "enum", operator: "LIKE", wantSQL: `SQL: "? LIKE ?"`},
		{name: "enum comparison", fieldType: "enum", operator: "=", wantSQL: `SQL: "? = ?"`},
		{name: "time range", fieldType: "time.Time", operator: "BETWEEN", wantSQL: `SQL: "? BETWEEN ? AND ?"`},
		{name: "numeric range", fieldType: "int", operator: "NOT BETWEEN", wantSQL: `SQL: "? NOT BETWEEN ? AND ?"`},
		{name: "string like", fieldType: "string", operator: "LIKE", wantSQL: `SQL: "? LIKE ?"`},
		{name: "numeric comparison", fieldType: "int", operator: ">=", wantSQL: `SQL: "? >= ?"`},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			field := &systemReq.AutoCodeField{
				FieldName:       "Group",
				ColumnName:      "group",
				FieldType:       tt.fieldType,
				FieldSearchType: tt.operator,
			}
			generated := GenerateSearchConditions([]*systemReq.AutoCodeField{field})

			if !strings.Contains(generated, `clause.Column{Name: "group"}`) {
				t.Fatalf("generated condition leaves the reserved column unquoted:\n%s", generated)
			}
			if !strings.Contains(generated, tt.wantSQL) {
				t.Fatalf("generated condition SQL = %q, want fragment %q", generated, tt.wantSQL)
			}
			source := "package generated\n\nimport \"gorm.io/gorm/clause\"\n\nfunc query() {\n" + generated + "\n}\n"
			if _, err := format.Source([]byte(source)); err != nil {
				t.Fatalf("generated condition is not valid Go source: %v\n%s", err, generated)
			}
		})
	}
}
