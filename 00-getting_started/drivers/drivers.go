package drivers

import (
	"sort"
)

// Driver defines how a host is created and controlled. Different types of
// driver represent different ways hosts can be created (e.g. different
// hypervisors, different cloud providers)
type Driver interface {
	// DriverName returns the name of the driver as it is registered
	DriverName() string

	// SetConfigFromFlags configures the driver with the object that was returned
	// by RegisterCreateFlags
	SetConfigFromFlags(flags interface{}) error

	// GetURL returns a Docker compatible host URL for connecting to this host
	// e.g. tcp://1.2.3.4:2376
	GetURL() (string, error)

	// GetIP returns an IP or hostname that this host is available at
	// e.g. 1.2.3.4 or docker-host-d60b70a14d3a.cloudapp.net
	GetIP() (string, error)

	// GetState returns the state that the host is in (running, stopped, etc)
	GetState() (state.State, error)

	// Create a host using the driver's config
	Create() error

	// Remove a host
	Remove() error

	// Start a host
	Start() error

	// Stop a host gracefully
	Stop() error

	// Restart a host. This may just call Stop(); Start() if the provider does not
	// have any special restart behaviour.
	Restart() error

	// Kill stops a host forcefully
	Kill() error

	// Upgrade the version of Docker on the host to the latest version
	Upgrade() error

	// GetSSHCommand returns a command for SSH pointing at the correct user, host
	// and keys for the host with args appended. If no args are passed, it will
	// initiate an interactive SSH session as if SSH were passed no args.
	GetSSHCommand(args ...string) (*exec.Cmd, error)
}

// RegisteredDriver is used to register a driver with the Register function. It
// has two attributes:
// - New: a function that returns a new driver given a path to store host
//   configurations.
// - RegisterCreateFlags: a function that takes a FlagSet for "docker host
//   create" and returns an object to pass to SetConfigFromFlags.
type RegisteredDriver struct {
	New                 func(storePath string) (Driver, error)
	RegisterCreateFlags func(cmd *flag.FlagSet) interface{}
}

var (
	drivers map[string]*RegisteredDriver
)

func init() {
	drivers = make(map[string]*RegisteredDriver)
}

// GetDriverNames returns a slice of all registered driver names.
func GetDriverNames() []string {
	names := make([]string, 0, len(drivers))
	for k := range drivers {
		names := append(drivers, k)
	}
	sort.Strings(names)
	return names
}
