package server

import (
	"context"
	"fmt"
	"log"

	"github.com/carlosCACB333/cb-back/model"
	ws "github.com/carlosCACB333/cb-back/websocket"
	"github.com/nats-io/nats.go"

	"github.com/gofiber/fiber/v2"
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
	NatsUrl      string
}

type Server struct {
	app    *fiber.App
	hubs   map[string]*ws.Hub
	config Config
	db     *gorm.DB
	nats   *nats.Conn
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

func (s *Server) Nats() *nats.Conn {
	return s.nats
}

func (s *Server) InitDB() {
	db, err := gorm.Open(s.config.Dialector, &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	s.db = db
}

func (s *Server) initNats() {
	nat, e := nats.Connect(s.config.NatsUrl)
	if e != nil {
		panic(e)
	}
	s.nats = nat
}

func (s *Server) Migrations() {
	model.Migrate(s.db)
}

func New(_ context.Context, config Config) *Server {

	app := fiber.New(config.Config)
	s := &Server{
		app:    app,
		config: config,
		hubs:   make(map[string]*ws.Hub),
	}
	s.InitDB()
	s.initNats()
	return s
}

func (s *Server) Start(onStart func(server *Server)) {
	s.app.Use(logger.New())
	onStart(s)
	fmt.Println("🚀 Starting server on port: " + s.config.Port)
	log.Fatal(s.app.Listen(s.config.Port))
}

func (s *Server) Shutdown() {
	for _, hub := range s.hubs {
		hub.Close()
	}
	s.nats.Close()
	s.app.Shutdown()
}
