// Package usd defines a decimal type for representing US dollars with 1/1000 cent ($0.00001) precision, as
// well as a few associated functions.
package usd

import "fmt"

// USD is a type representing US dollars as an int64 with 1/1000 cent precision.
// This means that dollar values in the range
//     [-$92233720368547.75808, $92233720368547.75807]
// (approx. ±$92 trillion) in increments of $0.00001 are represented precisely.
type USD int64

const (
	Cent   USD = 1000
	Dollar USD = 100 * Cent
)

// String formats d with a dollar sign ($) and rounded to the nearest cent.
func (d USD) String() string {
	s := ""
	if d < 0 {
		s = "-"
		d = -d
	}
	dollars := d / Dollar
	r := d % Dollar
	cents := r / Cent
	if r%Cent > (Cent / 2) {
		cents++
	}
	return fmt.Sprintf("%s$%d.%02d", s, dollars, cents)
}
