package kvstore

import (
	"bufio"
	"encoding/json"
	"errors"
	"log"
	"os"
	"sync"
)

type LogEntry struct {
	Operation string `json:"operation"`
	Key       string `json:"key"`
	Value     string `json:"value,omitempty"`
}

type KVStore interface {
	Set(k string, v string)
	Get(k string) (string, bool)
	Del(k string) error
}

type Store struct {
	data    map[string]string
	mu      sync.RWMutex
	aofFile *os.File
	aofPath string
}

func NewKVStore(persistence bool, aofPath string) *Store {
	if persistence {
		return NewPersistentKVStore(aofPath)
	}
	return &Store{
		data: make(map[string]string),
		mu:   sync.RWMutex{},
	}
}

func NewPersistentKVStore(aofPath string) *Store {
	s := &Store{
		data:    make(map[string]string),
		mu:      sync.RWMutex{},
		aofPath: aofPath,
	}

	var err error
	s.aofFile, err = os.OpenFile(aofPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatalf("Error opening AOF file: %v\n", err)
	}

	s.loadFromAOF()

	return s
}

func (s *Store) Set(k string, v string) {
	s.mu.Lock() // we are writing so we use Lock (write lock)
	defer s.mu.Unlock()

	s.data[k] = v

	// append to AOF if persistence is enabled (aofFile exists)
	if s.aofFile != nil {
		s.appendLog("SET", k, v)
	}
}

func (s *Store) Get(k string) (string, bool) {
	s.mu.RLock() // we aren't writing so we use RLock (read lock)
	defer s.mu.RUnlock()

	v, ok := s.data[k]
	return v, ok
}

func (s *Store) Del(k string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.data[k]; !ok {
		return errors.New("key not found")
	}

	delete(s.data, k)

	// append to AOF if persistence is enabled (aofFile exists)
	if s.aofFile != nil {
		s.appendLog("DEL", k, "")
	}

	return nil
}

func (s *Store) appendLog(operation, key, value string) {
	entry := LogEntry{
		Operation: operation,
		Key:       key,
		Value:     value,
	}

	data, err := json.Marshal(entry)
	if err != nil {
		log.Printf("Error marshaling AOF entry: %v\n", err)
		return
	}

	_, err = s.aofFile.WriteString(string(data) + "\n")
	if err != nil {
		log.Printf("Error writing to AOF file: %v\n", err)
		return
	}

	s.aofFile.Sync()
}

func (s *Store) loadFromAOF() {
	file, err := os.Open(s.aofPath)
	if err != nil {
		if os.IsNotExist(err) {
			log.Println("AOF file doesn't exist, starting with empty database")
			return
		}
		log.Printf("Error opening AOF file: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lineCount := 0

	for scanner.Scan() {
		lineCount++
		line := scanner.Text()
		if line == "" {
			continue
		}

		var entry LogEntry
		err := json.Unmarshal([]byte(line), &entry)
		if err != nil {
			log.Printf("Error unmarshaling AOF entry at line %d: %v\n", lineCount, err)
			continue
		}

		// apply the write operations again to the in-memory store to reconstruct the state essentially "tricking" persistence
		switch entry.Operation {
		case "SET":
			s.data[entry.Key] = entry.Value
		case "DEL":
			delete(s.data, entry.Key)
		default:
			log.Printf("Unknown operation in AOF: %s\n", entry.Operation)
		}
	}

	if lineCount == 0 {
		log.Println("AOF file is empty, starting with empty database")
		return
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading AOF file: %v\n", err)
		return
	}

	log.Printf("Loaded %d operations from AOF file\n", lineCount)
}

func (s *Store) Close() error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.aofFile != nil {
		return s.aofFile.Close()
	}
	return nil
}
