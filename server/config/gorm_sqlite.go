package config

type Sqlite struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (s *Sqlite) Dsn() string {
	return s.Path + "\\" + s.Dbname + ".db"
}

func (s *Sqlite) GetLogMode() string {
	return s.LogMode
}
