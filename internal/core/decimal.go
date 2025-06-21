package core

import (
	"fmt"
	"strconv"
)

type Decimal struct{ val float64 }

func (d Decimal) String() string              { return strconv.FormatFloat(d.val, 'g', -1, 64) }
func (d Decimal) IsZero() bool                { return d == Decimal{} }
func (d Decimal) EqualsTo(outer Decimal) bool { return d == outer }
func (d Decimal) Value() float64              { return d.val }

func NewDecimal(val float64) (Decimal, error) {
	if val < 0 {
		return Decimal{}, fmt.Errorf("invalid value: must be a non-negative number")
	}

	return Decimal{val: val}, nil
}

type DecimalPrecision struct{ val uint32 }

func (d DecimalPrecision) Value() uint32                        { return d.val }
func (d DecimalPrecision) EqualsTo(outer DecimalPrecision) bool { return d.val == outer.val }

func NewDecimalPrecision(val uint32) (DecimalPrecision, error) {
	if val == 0 {
		return DecimalPrecision{}, fmt.Errorf("invalid precision: must be a number greater than zero")
	}

	return DecimalPrecision{val: val}, nil
}
