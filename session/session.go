package session

import (
	"context"
	"log"
	"strconv"
	"sync"
	"sync/atomic"
	"time"

	"github.com/bougou/go-ipmi"
)

type SessionManager struct {
	mu          sync.RWMutex
	client      *ipmi.Client
	sdrSnapshot atomic.Value
}

var (
	manager       *SessionManager
	once          sync.Once
	cancelMonitor context.CancelFunc
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
	// init 
	//! this code is very ugly, must be rewrite
	newList, err := s.client.GetSDRs(ctx)
	if err != nil {
		log.Println(err)	
	}
	s.UpdateSDRs(newList)

	//Start auto refresh sdr
	ctx, cancelMonitor = context.WithCancel(context.Background())
	s.StartAutoRefresh(ctx)
	return nil
}

func (s *SessionManager) Logout() error {
	if cancelMonitor != nil {
		// Stop auto refresh sdr
		cancelMonitor()
	}
	ctx := context.Background()
	err := s.client.Close(ctx)
	s.DeleteClient()
	return err
}

func (s *SessionManager) GetSDRs() []*ipmi.SDR {
	value := s.sdrSnapshot.Load()
	if value == nil {
		return nil
	}
	return value.([]*ipmi.SDR)
}

func (s *SessionManager) UpdateSDRs(newSDRs []*ipmi.SDR) {
	s.sdrSnapshot.Store(newSDRs)
}

func (s *SessionManager) StartAutoRefresh(ctx context.Context) {
	go func() {
		ticker := time.NewTicker(5 * time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				s.doRefresh(ctx)
			}
		}
	}()
}

func (s *SessionManager) doRefresh(ctx context.Context) {
	newList, err := s.client.GetSDRs(ctx)
	if err != nil {
		log.Println(err)
		return
	}
	s.UpdateSDRs(newList)
}
