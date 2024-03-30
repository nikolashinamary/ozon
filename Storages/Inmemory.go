package Storages

import "github.com/sirupsen/logrus"

type InMemoryStorage struct {
	StorageHashByUrl map[string]string
	StorageUrlByHash map[string]string
}

func InMemoryStorageConstr() *InMemoryStorage {
	return &(InMemoryStorage{StorageHashByUrl: make(map[string]string), StorageUrlByHash: make(map[string]string)})
}

func (inMemoryStorage *InMemoryStorage) GetByHash(hash string, log1 *logrus.Logger) string {
	return inMemoryStorage.StorageUrlByHash[hash]
}

func (inMemoryStorage *InMemoryStorage) GetByUrl(url string, log1 *logrus.Logger) string {
	return inMemoryStorage.StorageHashByUrl[url]
}

func (inMemoryStorage *InMemoryStorage) WriteByUrl(hash string, url string, log1 *logrus.Logger) error {
	inMemoryStorage.StorageHashByUrl[url] = hash
	inMemoryStorage.StorageUrlByHash[hash] = url
	return nil

}

func (inMemoryStorage *InMemoryStorage) ContainsByUrl(url string, log1 *logrus.Logger) bool {
	_, flag := inMemoryStorage.StorageHashByUrl[url]
	if flag {
		return true
	}
	return false
}

func (inMemoryStorage *InMemoryStorage) ContainsByHash(hash string, log1 *logrus.Logger) bool {
	_, flag := inMemoryStorage.StorageUrlByHash[hash]
	if flag {
		return true
	}
	return false
}

func (inMemoryStorage *InMemoryStorage) DeleteByUrl(url string, log1 *logrus.Logger) {
	hash := inMemoryStorage.GetByUrl(url, log1)
	delete(inMemoryStorage.StorageUrlByHash, hash)
	delete(inMemoryStorage.StorageHashByUrl, url)
}
