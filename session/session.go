package session

import (
	"context"
	"log"
	"strconv"
	"sync"

	"github.com/bougou/go-ipmi"
)

type SessionManager struct {
	mu     sync.RWMutex
	client *ipmi.Client
}

var (
	manager *SessionManager
	once    sync.Once
)

func GetInstance() *SessionManager {
	once.Do(
		func() {
			manager = &SessionManager{}
		})
	return manager
}

func (s *SessionManager) UpdateClient(c *ipmi.Client) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.client = c
}

func (s *SessionManager) GetClient() *ipmi.Client {
	s.mu.Lock()
	defer s.mu.RUnlock()
	return s.client
}

func (s *SessionManager) DeleteClient() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.client = nil
}

func (s *SessionManager) Login(host, port, user, pass string) error {
	p, _ := strconv.Atoi(port)
	client, err := ipmi.NewClient(host, p, user, pass)
	client.WithInterface(ipmi.InterfaceLanplus)
	if err != nil {
		return err
	}

	ctx := context.Background()
	if err := client.Connect(ctx); err != nil {
		return err
	}
	guid, err := client.GetSystemGUID(ctx)
	if err != nil {
		return err
	}
	log.Println(guid)
	s.UpdateClient(client)
	return nil
}

func (s *SessionManager) Logout() error {
	ctx := context.Background()
	err := s.client.Close(ctx)
	s.DeleteClient()
	return err
}

