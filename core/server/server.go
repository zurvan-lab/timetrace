package server

import (
	"fmt"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/zurvan-lab/TimeTrace/config"
	"github.com/zurvan-lab/TimeTrace/core/TQL/execute"
	parser "github.com/zurvan-lab/TimeTrace/core/TQL/parser"
	"github.com/zurvan-lab/TimeTrace/core/database"
	ttlogger "github.com/zurvan-lab/TimeTrace/log"
)

type Server struct {
	ListenAddress     string
	Listener          net.Listener
	QuitChannel       chan struct{}
	Wg                sync.WaitGroup
	ActiveConnections map[net.Conn]struct{}
	ActiveConnsMux    sync.Mutex
	Config            *config.Config

	db *database.Database
}

func NewServer(cfg *config.Config, db *database.Database) *Server {
	lna := fmt.Sprintf("%v:%v", cfg.Server.IP, cfg.Server.Port)

	return &Server{
		ListenAddress:     lna,
		QuitChannel:       make(chan struct{}),
		ActiveConnections: make(map[net.Conn]struct{}),
		db:                db,
		Config:            cfg,
	}
}

func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.ListenAddress)
	if err != nil {
		return err
	}
	defer listener.Close()
	s.Listener = listener

	ttlogger.Info("server started", "address", s.ListenAddress, "db-name", s.Config.Name)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-signalChan
		ttlogger.Info("Received signal, shutting down...", "signal", sig, "db-name", s.Config.Name)

		close(s.QuitChannel)

		s.CloseAllConnections()
	}()

	s.Wg.Add(1)
	go s.AcceptConnections()

	<-s.QuitChannel

	return nil
}

func (s *Server) AcceptConnections() {
	defer s.Wg.Done()

	for {
		select {
		case <-s.QuitChannel:
			return
		default:
		}

		conn, err := s.Listener.Accept()
		if err != nil {
			ttlogger.Error("error accepting connection", "error", err, "db-name", s.Config.Name)

			continue
		}

		buffer := make([]byte, 1024)

		n, err := conn.Read(buffer)
		if err != nil {
			ttlogger.Error("error reading connection", "error", err, "db-name", s.Config.Name)

			_ = conn.Close()
		}

		query := parser.ParseQuery(string(buffer[:n]))

		result := execute.Execute(query, s.db)
		if result != "DONE" {
			ttlogger.Warn("invalid user try to connect", "db-name", s.Config.Name)

			_ = conn.Close()
		}

		s.ActiveConnsMux.Lock()
		s.ActiveConnections[conn] = struct{}{}
		s.ActiveConnsMux.Unlock()

		ttlogger.Info("new connection", "remote address", conn.RemoteAddr().String(), "db-name", s.Config.Name)

		s.Wg.Add(1)
		go s.ReadConneciton(conn)
	}
}

func (s *Server) ReadConneciton(conn net.Conn) {
	defer conn.Close()
	defer s.Wg.Done()

	buffer := make([]byte, 2050)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			ttlogger.Error("Connection closed", "remote address", conn.RemoteAddr().String(), "db-name", s.Config.Name)

			s.ActiveConnsMux.Lock()
			delete(s.ActiveConnections, conn)
			s.ActiveConnsMux.Unlock()

			return
		}

		query := parser.ParseQuery(string(buffer[:n]))
		result := execute.Execute(query, s.db)

		_, err = conn.Write([]byte(result))
		if err != nil {
			ttlogger.Error("Can't write on TCP connection", "error", err, "db-name", s.Config.Name)
		}
	}
}

func (s *Server) CloseAllConnections() {
	s.ActiveConnsMux.Lock()
	defer s.ActiveConnsMux.Unlock()

	for conn := range s.ActiveConnections {
		conn.Close()
		delete(s.ActiveConnections, conn)
	}
}
