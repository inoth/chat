package toybox

type Component interface {
	Init() error
	// Close() error
}
