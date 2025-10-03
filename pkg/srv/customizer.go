package srv

// ServerCustomizer a server customizer
type ServerCustomizer interface {
	// Register registers route
	Register(server *Server)
}
