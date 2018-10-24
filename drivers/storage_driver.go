package drivers

import "github.com/Liar233/Scheduler/model"

type StorageDriver interface {
	Connect(conf map[string]interface{}) error
	Close() error
	Creator
	Remover
	Querier
}

type Creator interface {
	Create(e model.Event) (interface{}, error)
}

type Remover interface {
	Remove(id interface{}) error
}

type Querier interface {
	Get(id interface{}) (model.Event, error)
}
