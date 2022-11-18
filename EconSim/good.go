package sim

import "fmt"

// Good
type Good struct {
	name  string
	value float64
}

func (g Good) String() string {
	return fmt.Sprintf("%s | %f", g.name, g.value)
}

var goods []Good
