// +build windows

package windows

import (
	"fmt"

	"github.com/hyperhq/hypercli/daemon/execdriver"
)

// Stats implements the exec driver Driver interface.
func (d *Driver) Stats(id string) (*execdriver.ResourceStats, error) {
	return nil, fmt.Errorf("Windows: Stats not implemented")
}
