package storage

import "sync"

type Locator struct {
	mu       sync.RWMutex
	services map[string]interface{}
}

var globalLocator = &Locator{services: make(map[string]interface{})}

func Register(key string, service interface{}) {
	globalLocator.mu.Lock()
	defer globalLocator.mu.Unlock()
	globalLocator.services[key] = service
}

func Get(key string) interface{} {
	globalLocator.mu.RLock()
	defer globalLocator.mu.RUnlock()
	return globalLocator.services[key]
}

func GetDBManager() *DBManager {
	if service, ok := Get("DBManager").(*DBManager); ok {
		return service
	}
	return nil
}
