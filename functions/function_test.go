package functions

import (
	"testing"

	"github.com/delineateio/mimas/routes"
)

func TestFunction(t *testing.T) {
	items := routes.AddDefaultRoutes(nil)
	NewFunction(items)
}
