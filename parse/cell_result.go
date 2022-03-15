package parse

type CellResult string

const (
	Correct   CellResult = "correct"
	Misplaced CellResult = "Misplaced"
	Wrong     CellResult = "wrong"
	Empty     CellResult = ""
)
