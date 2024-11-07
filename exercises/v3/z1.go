package v3

type Oblik interface {
	Obim()
}

type Trougao struct{
	A int
	B int
	C int
}

type Pravougaonik struct{
	A int
	B int
}

func (p *Pravougaonik) Obim() int {
	return p.A * 2 + p.B * 2
}

func (t *Trougao) Obim() int {
	return t.A + t.B + t.C
}

