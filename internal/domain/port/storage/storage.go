package storage

import "github.com/Sabyradinov/golang-hex-layout/internal/domain/port/repository"

type IDB interface {
	InitRepo() *repository.Manager
	RegisterMetrics(dbname string) (err error)
	Close() (err error)
}
