package support

import "fmt"

type Stringable interface {
	fmt.Stringer
}
