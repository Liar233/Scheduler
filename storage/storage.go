package storage

import (
	"github.com/Liar233/Scheduler/drivers"
	"github.com/Liar233/Scheduler/config"
	"github.com/Liar233/Scheduler/model"
	"fmt"
	"time"
)

type EventStorage struct {
	driver *drivers.StorageDriver
	freezeTimeout uint
}

func (es *EventStorage) Create(event model.Event) (model.Event, error) {
	id, err := es.driver.Create(event)

	if err != nil {
		return event, err
	}

	event.ID = fmt.Sprintf("%v", id)

	return event, err
}

func (es *EventStorage) Delete(event model.Event) error {
	return es.driver.Delete(event.ID)
}

func (es *EventStorage) Query(params map[string]map[string]interface{}) ([]model.Event, error) {
	return es.driver.Query(params)
}

func (es *EventStorage) Get(id string) (model.Event, error) {
	return es.driver.Get(id)
}

func (es *EventStorage) GetFreezedEvents() ([]model.Event, error) {
	now := time.Now()

	params := make(map[string]map[string]interface{})
	params["firetime"][">="] = now
	params["firetime"]["<="] = now.Add(time.Minute * time.Duration(es.freezeTimeout))

	return es.driver.Query(params)
}

func NewEventStorage(conf *config.AppConfig) (*EventStorage, error) {
	d := drivers.NewStorageDriver(conf.Storage)

	err := d.Connect()

	if err != nil {
		return nil, err
	}

	return &EventStorage{
		driver: d,
		freezeTimeout: conf.FreezeTimeout,
	}, nil
}
