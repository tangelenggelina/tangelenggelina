package answerremover

type commentvariation int

const (
	doubleslash commentvariation = iota
	dashdash
	hash
	slashasterisk
	curlyslashasterisk
)

type endanswer struct {
	annotation *string
	variation  *commentvariation
}
