package solution

type RowItem struct {
	Left  string
	Right string
}

type ProblemData struct {
	Path string

	Rows map[string]*RowItem
}

func NewProblemData() *ProblemData {
	return &ProblemData{Rows: make(map[string]*RowItem)}
}
