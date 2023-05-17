package toybox

type ServerOption func(sv Server)

type Server interface {
	Init() error
	Start() error
}
