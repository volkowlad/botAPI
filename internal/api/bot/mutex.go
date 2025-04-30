package bot

import (
	"sync"
	"time"
)

var (
	chatMutexes  = make(map[int64]*sync.Mutex)
	lastActivity = make(map[int64]time.Time)
	globalMutex  sync.Mutex // Для защиты мапы chatMutexes
)

func getChatMutex(chatID int64) *sync.Mutex {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	if _, exists := chatMutexes[chatID]; !exists {
		chatMutexes[chatID] = &sync.Mutex{}
	}

	lastActivity[chatID] = time.Now()

	return chatMutexes[chatID]
}

func cleanupMutexes() {
	globalMutex.Lock()
	defer globalMutex.Unlock()

	for chatID := range chatMutexes {
		// Удаляем мьютексы чатов, где не было активности N минут
		if time.Since(lastActivity[chatID]) > 1*time.Hour {
			delete(chatMutexes, chatID)
			delete(lastActivity, chatID)
		}
	}
}

func StartCleanupRoutine(interval time.Duration) {
	go func() {
		ticker := time.NewTicker(interval)
		defer ticker.Stop()

		for range ticker.C {
			cleanupMutexes()
		}
	}()
}
