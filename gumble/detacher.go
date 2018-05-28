package gumble // import "github.com/triwav/gumble/gumble"

// Detacher is an interface that event listeners implement. After the Detach
// method is called, the listener will no longer receive events.
type Detacher interface {
	Detach()
}
