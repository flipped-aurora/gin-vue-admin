package config

type Oracle struct {
	GeneralDB `yaml:",inline" mapstructure:",squash"`
}

func (m *Oracle) Dsn() string {
	return "oracle://" + m.Username + ":" + m.Password + "@" + m.Path + ":" + m.Port + "/" + m.Dbname + "?" + m.Config

}

func (m *Oracle) GetLogMode() string {
	return m.LogMode
}
