package server

import (
	"fold/internal"
	"fold/internal/model"
	"fold/internal/service"
	"net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/gin-gonic/gin"
)

type Server struct {
	ProjectService internal.Service
	Engine         *gin.Engine
}

func New(esClient *elasticsearch.Client) (*Server, error) {
	r := gin.Default()

	projectService, err := service.NewProject(esClient)
	if err != nil {
		return nil, err
	}
	s := &Server{
		ProjectService: projectService,
		Engine:         r,
	}
	r.GET("/projects/created-by/:user_id", s.searchProjectsByUser)
	r.POST("/projects/search-with-hashtags", s.searchProjectsByHashtags)
	r.GET("/projects/search", s.fullTextSearchProjects)
	return s, nil
}

func (s *Server) Start() {
	s.Engine.Run(":8080")
}

func (s *Server) searchProjectsByUser(c *gin.Context) {
	userID := strings.TrimSpace(c.Param("user_id"))
	if userID == "" {
		c.JSON(400, gin.H{
			"error": "user_id is mandatory",
		})
		return
	}

	projects, err := s.ProjectService.SearchProjectsByUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "An internal server error occurred. Please try again later.",
		})
		return
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})
}

func (s *Server) searchProjectsByHashtags(c *gin.Context) {

	var req model.HashTagsInput
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if len(req.Hashtags) == 0 {
		c.JSON(400, gin.H{
			"error": "hashtags are mandatory",
		})
		return
	}

	projects, err := s.ProjectService.SearchProjectsByHashtags(c.Request.Context(), req.Hashtags)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "An internal server error occurred. Please try again later.",
		})
		return
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})
}

func (s *Server) fullTextSearchProjects(c *gin.Context) {
	searchQuery := strings.TrimSpace(c.Query("q"))
	if searchQuery == "" {
		c.JSON(400, gin.H{
			"error": "Query param q is mandatory",
		})
		return
	}

	projects, err := s.ProjectService.FullTextSearchProjects(c.Request.Context(), searchQuery)
	if err != nil {
		c.JSON(500, gin.H{
			"error": "An internal server error occurred. Please try again later.",
		})
		return
	}

	c.JSON(200, gin.H{
		"projects": projects,
	})
}
