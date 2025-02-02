package main

import (
	"database/sql"
	"fmt"
	"web-article/config"
	"web-article/controller"
	"web-article/middleware"
	"web-article/repository"
	"web-article/usecase"
	"web-article/utils/service"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	// Instean usecase
	commentUC  usecase.CommentUsecase
	articleUC  usecase.ArticleUsecase
	userUC     usecase.UserUseCase
	authUC     usecase.AuthenticationUseCase
	jwtService service.JWTService
	engine     *gin.Engine
	host       string
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.Database)

	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic(err)
	}

	jwtService := service.NewJWTService(cfg.TokenConfig)

	// Instean repository
	userRepo := repository.NewUserRepository(db)
	articleRepo := repository.NewArticleRepository(db)
	commentRepo := repository.NewCommentRepository(db)
	// Instean usecase
	userUseCase := usecase.NewUserUseCase(userRepo)
	authUseCase := usecase.NewAuthenticationUsecase(userUseCase, jwtService)
	articleUseCase := usecase.NewArticleUseCase(userUseCase, articleRepo)
	commentUseCase := usecase.NewCommentUseCase(commentRepo, articleUseCase, userUseCase)

	engine := gin.Default()

	host := fmt.Sprintf(":%s", cfg.ApiPort)

	return &Server{
		// Instean usecase
		commentUC:  commentUseCase,
		articleUC:  articleUseCase,
		userUC:     userUseCase,
		authUC:     authUseCase,
		jwtService: jwtService,
		engine:     engine,
		host:       host,
	}
}

func (s *Server) initRoute() {
	rg := s.engine.Group("/api")

	authMiddleware := middleware.NewAuthMiddleware(s.jwtService)

	// Instean controller
	controller.NewCommentController(s.commentUC, rg, authMiddleware).Route() // route dengan middleware
	controller.NewArticleController(s.articleUC, rg, authMiddleware).Route() // route dengan middleware
	controller.NewUserController(s.userUC, rg, authMiddleware).Route()       // route dengan middleware
	controller.NewAuthController(s.authUC, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()

	err := s.engine.Run(s.host)
	if err != nil {
		panic(fmt.Errorf("failed to start server on host %s: %v", s.host, err))
	}
}
