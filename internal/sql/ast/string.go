package sqlc

type String struct {
	Str string
}

func (n *String) Pos() int {
	return 0
}
