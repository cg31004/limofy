package web

import "github.com/gin-gonic/gin"

func (s *restService) setApiRouters(parentRouteGroup *gin.RouterGroup) {
	privateRouteGroup := parentRouteGroup.Group("/v1")

	s.setMovieAPIRoutes(privateRouteGroup)
	s.setLineAPIRoutes(privateRouteGroup)
}
