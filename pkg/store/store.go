package store

import (
	"os"
	"path"
)

type KvFsStore struct {
	Root string
}

func (s *KvFsStore) Put(key string, value string) {
	os.WriteFile(path.Join(s.Root, key), []byte(value), os.ModeAppend)
}

func (s *KvFsStore) Get(key string) string {
	b, err := os.ReadFile(path.Join(s.Root, key))
	if err != nil {
		return ""
	}

	return string(b)
}

func (s *KvFsStore) Delete(key string) {
	os.Remove(path.Join(s.Root, key))
}
