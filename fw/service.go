package fw

import (
	"fmt"
)

type Service struct {
	name   string
	server Server
	logger Logger
}

func (s Service) Start(port int) {
	defer s.logger.Info(fmt.Sprintf("%s service started", s.name))

	go func() {
		err := s.server.ListenAndServe(port)

		if err != nil {
			s.logger.Error(err)
		}
	}()
}

func (s Service) Stop() {
	defer s.logger.Info(fmt.Sprintf("%s stopped", s.name))

	err := s.server.Shutdown()
	if err != nil {
		s.logger.Error(err)
	}
}

func NewService(name string, server Server, logger Logger) Service {
	return Service{
		name:   name,
		server: server,
		logger: logger,
	}
}