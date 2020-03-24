package autoCodeModel

//开发中功能，若您发现这块代码可以研究，可以无视
type AutoCodeStruct struct {
	StructName string      `json:"structName"`
	StructType []string    `json:"structType"`
	Components []Component `json:"components"`
}

type Component struct {
	ComponentName       string       `json:"componentName"`
	ComponentType       string       `json:"componentType"`
	Ismultiple          bool         `json:"isMultiple"`
	ComponentShowType   string       `json:"componentShowType"`
	NideDictionary      bool         `json:"nideDictionary"`
	DictionaryName      string       `json:"dictionaryName"`
	ComponentDictionary []Dictionary `json:"dictionary"`
}

type Dictionary struct {
	Label string `json:"label"`
	Value string `json:"value"`
}
