package storage

import (
	"github.com/Liar233/Scheduler/drivers"
	"github.com/Liar233/Scheduler/config"
)

type EventStorage struct {
	driver *drivers.StorageDriver
}

func NewEventStorage(conf *config.AppConfig) (*EventStorage, error) {
	d := drivers.NewStorageDriver(conf.Storage)

	return &EventStorage{driver: d}, nil
}
