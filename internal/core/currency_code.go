package core

type CurrencyCode struct{ s string }

func (g CurrencyCode) FromTo(outer CurrencyCode) string { return g.s + "->" + outer.s }
func (g CurrencyCode) EqualsTo(outer CurrencyCode) bool { return g == outer }
func (g CurrencyCode) String() string                   { return string(g.s) }
