package commons

// Row é a estrutura que guardará os valores das linhas na representação do banco.
type Row struct {
	Row []string `json:"row"`
}

// Table é a estrutura que guardará as informações na representação do banco.
type Table struct {
	Name     string   `json:"name"`
	Columns  []string `json:"columns"`
	Rows     []Row    `json:"rows"`
	Optional string   `json:"optional"`
}

// NewTable retorna a instânci de uma estrutura Table.
func (api Api) NewTable() *Table {
	return new(Table)
}
