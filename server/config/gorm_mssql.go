package config

type Mssql struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

// Dsn "sqlserver://gorm:LoremIpsum86@localhost:9930?database=gorm"
func (m *Mssql) Dsn() string {
	return "sqlserver://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "?database=" + m.Dbname + "&encrypt=disable"
}
