package controllers

// initializeRoutes
func (server *Server) initializeRoutes() {
	// Home Route
	server.Router.HandleFunc("/", SetMiddlewareJSON(server.Home)).Methods("GET")

	// Login Route
	server.Router.HandleFunc("/login", SetMiddlewareJSON(server.Login)).Methods("POST")

	// User Route
	server.Router.HandleFunc("/users", SetMiddlewareAuthentificaton(AuthCheckRole(server.CreateUser))).Methods("POST")
	server.Router.HandleFunc("/users", SetMiddlewareJSON(server.GetUsers)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", SetMiddlewareJSON(server.GetUser)).Methods("GET")
	server.Router.HandleFunc("/users/{id}", SetMiddlewareJSON(SetMiddlewareAuthentificaton(server.UpdateUser))).Methods("PUT")
	server.Router.HandleFunc("/users/{id}", SetMiddlewareAuthentificaton(server.DeleteUser)).Methods("DELETE")

	//Posts Route
	server.Router.HandleFunc("/posts", SetMiddlewareJSON(server.CreatePost)).Methods("POST")
	server.Router.HandleFunc("/posts", SetMiddlewareJSON(server.GetPosts)).Methods("GET")
	server.Router.HandleFunc("/posts/{id}", SetMiddlewareJSON(server.GetPost)).Methods("GET")
	server.Router.HandleFunc("/posts/{id}", SetMiddlewareJSON(SetMiddlewareAuthentificaton(server.UpdatePost))).Methods("PUT")
	server.Router.HandleFunc("/posts/{id}", SetMiddlewareAuthentificaton(server.DeletePost)).Methods("DELETE")

	// Role Route
	server.Router.HandleFunc("/role", SetMiddlewareJSON(CreateRoleCasbin)).Methods("POST")

}
