package server

import (
	"cb/model"
	ws "cb/websocket"
	"context"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"gorm.io/gorm"
)

type Config struct {
	Port         string
	Config       fiber.Config
	Dialector    gorm.Dialector
	ApiKey       string
	ApiKeyPublic string
	JWTSecret    string
}

type Server struct {
	app    *fiber.App
	hubs   map[string]*ws.Hub
	config Config
	db     *gorm.DB
}

func (s *Server) Config() Config {
	return s.config
}
func (s *Server) Hub(key string) *ws.Hub {
	hub, ok := s.hubs[key]
	if ok {
		return hub
	}
	h := ws.NewHub()
	s.hubs[key] = h
	return h
}

func (s *Server) App() *fiber.App {
	return s.app
}

func (s *Server) DB() *gorm.DB {
	return s.db
}
func (s *Server) initDB() {
	db, err := gorm.Open(s.config.Dialector, &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	s.db = db
}
func (s *Server) Migrations() {
	model.Migrate(s.db)
}

func New(_ context.Context, config Config) *Server {
	app := fiber.New(config.Config)
	return &Server{
		app:    app,
		config: config,
		hubs:   make(map[string]*ws.Hub),
	}
}

func (s *Server) Start(setRouters func(server *Server)) {
	s.app.Use(logger.New())
	s.app.Static("/public", "./public")
	s.app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000,https://carloscb.com,https://www.carloscb.com",
		AllowHeaders: "*",
	}))

	s.initDB()

	setRouters(s)
	fmt.Println("🚀 Starting server on port: " + s.config.Port)
	log.Fatal(s.app.Listen(s.config.Port))
}
