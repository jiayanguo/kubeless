package kubeless

import (
	"github.com/kubeless/kubeless/pkg/functions"
)

// Foo sample function
func Handler(event functions.Event, context functions.Context) (string, error) {
	return "Hello world!", nil
}