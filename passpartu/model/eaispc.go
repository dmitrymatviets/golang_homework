package model

type EAISPC struct {
	Budkas map[string]*Budka
}

func NewEAISPC() *EAISPC {
	r := new(EAISPC)
	r.Budkas = make(map[string]*Budka)
	return r
}

func (e *EAISPC) RegisterBudka(b *Budka, name string) {
	e.Budkas[name] = b
}
