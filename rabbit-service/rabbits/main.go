//go:generate stringer -type=Rabbit

package rabbits

import "strconv"

type Rabbit int

const (
	Alaska Rabbit = iota
	Altex
	American
	Angora
	Astrex
	Bauscat
	Brazilian
	Beveren
	Pygmy
	Cottontail
)

func StringToRabbit(s string) (Rabbit,error) {
	i, err := strconv.Atoi(s)
	return Rabbit(i), err
}
