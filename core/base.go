package core

type DatabaseTable struct {
	Name    string
	Columns []DatabaseTableColumn
}
type DatabaseTableColumn struct {
	Name string
	Type string
}
type DatabaseCredentials struct {
	Name   string
	Path   string
	Driver string
}
type Field struct {
	Name string `yaml:"name"`
	Type string `yaml:"type"`
}
type Entity struct {
	Name   string  `yaml:"name"`
	Fields []Field `yaml:"fields"`
}
type Manifest struct {
	Entities []Entity `yaml:"entities"`
}
