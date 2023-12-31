package server

import (
	"mu_previous_papers_be/model"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Store interface {
	HealthCheck() error
	GetTitles(subject, code, year string) []model.QpapersInfo
	GetSubjects() []string
	GetSubjectCodes() []string
}

type Server struct {
	router *gin.Engine
	store  Store
}

func NewServer(store Store) *Server {
	return &Server{
		router: gin.Default(),
		store:  store,
	}
}

func (s *Server) Run() {
	s.router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	s.router.GET("/health", s.healthCheck())
	s.router.GET("/papersData", s.getPaperTitles())
	s.router.GET("/subjects", s.getSubjects())
	s.router.Run(":8080")
}

func (s *Server) healthCheck() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"status": "online"})
	}
}

func (s *Server) getSubjects() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		res := s.store.GetSubjects()
		resCodes := s.store.GetSubjectCodes()
		ctx.JSON(http.StatusOK, gin.H{
			"subjects":     res,
			"subjectCodes": resCodes,
		})
	}
}

func (s *Server) getPaperTitles() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		year := ctx.Query("year")
		subject := ctx.Query("subject")
		code := ctx.Query("code")
		res := s.store.GetTitles(subject, code, year)
		ctx.JSON(http.StatusOK, res)
	}
}
