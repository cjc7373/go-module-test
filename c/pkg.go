package c

import (
	"fmt"

	"github.com/cjc7373/go-module-test/d"
)

func For() string {
	return fmt.Sprint(d.What())
}
