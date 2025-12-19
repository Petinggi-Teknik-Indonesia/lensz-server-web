package service

import (
	"sync"
	"time"
)

type ScannerHeartbeatStore struct {
	mu        sync.Mutex
	lastSeen  map[string]time.Time
	timeout   time.Duration
}

func NewScannerHeartbeatStore(timeout time.Duration) *ScannerHeartbeatStore {
	return &ScannerHeartbeatStore{
		lastSeen: make(map[string]time.Time),
		timeout:  timeout,
	}
}

// Called by hardware heartbeat
func (s *ScannerHeartbeatStore) Beat(deviceName string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.lastSeen[deviceName] = time.Now()
}

// Used by API to determine status
func (s *ScannerHeartbeatStore) IsOnline(deviceName string) (bool, *time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()

	t, ok := s.lastSeen[deviceName]
	if !ok {
		return false, nil
	}

	online := time.Since(t) <= s.timeout
	return online, &t
}
