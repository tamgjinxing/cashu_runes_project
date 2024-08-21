package main

import (
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
)

// ApiKeyStore 用于存储和管理 API Key
type ApiKeyStore struct {
	keys map[string]bool
	mu   sync.Mutex
}

func NewApiKeyStore() *ApiKeyStore {
	return &ApiKeyStore{
		keys: make(map[string]bool),
	}
}

func (s *ApiKeyStore) GenerateApiKey() string {
	// 生成随机的 API Key
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	apiKey := hex.EncodeToString(bytes)

	// 存储 API Key
	s.mu.Lock()
	defer s.mu.Unlock()
	s.keys[apiKey] = true

	return apiKey
}

func (s *ApiKeyStore) InitApiKeyTo(){
	
}

func (s *ApiKeyStore) ValidateApiKey(apiKey string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.keys[apiKey]
}

var store = NewApiKeyStore()

// 获取 API Key 的处理程序
func GetApiKeyHandler(c *gin.Context) {
	callerId := c.Query("callerId")
	apiKey := store.GenerateApiKey()
	SaveApiKey(callerId, apiKey)

	c.JSON(http.StatusOK, gin.H{"api_key": apiKey})
}

// API Key 验证中间件
func ApiKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-Api-Key")
		if !store.ValidateApiKey(apiKey) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing API Key"})
			c.Abort()
			return
		}
		c.Next()
	}
}
