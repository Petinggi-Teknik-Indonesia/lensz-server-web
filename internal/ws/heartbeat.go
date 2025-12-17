package ws

import (
	"sync"
	"time"
)

type ScannerHeartbeat struct {
	ScannerID uint
	LastSeen  time.Time
}

var (
	store = make(map[uint]ScannerHeartbeat)
	mu    sync.RWMutex
)
