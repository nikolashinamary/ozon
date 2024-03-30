package Storages

import "github.com/sirupsen/logrus"

type Storage interface {
	GetByUrl(url string, log1 *logrus.Logger) string
	GetByHash(hash string, log1 *logrus.Logger) string
	WriteByUrl(hash string, url string, log1 *logrus.Logger) error
	ContainsByUrl(url string, log1 *logrus.Logger) bool
	ContainsByHash(hash string, log1 *logrus.Logger) bool
	DeleteByUrl(url string, log1 *logrus.Logger)
}
