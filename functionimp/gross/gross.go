package gross

import (
	b "functionimp/basic"
)

func GrossSalary() int {
	a, t := b.Calculation()
	return ((b.Basic + a) - t)
}