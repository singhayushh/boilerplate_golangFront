package handlers

func (server *Server) initRoutes() {
	server.Router.GET("/", RenderHome)
}
