package drivers

import (
	_ "github.com/lib/pq"
	"fmt"
	"database/sql"
	"github.com/Liar233/Scheduler/config"
	"github.com/Liar233/Scheduler/model"
)

type StorageDriver struct {
	config     config.StorageConfig
	connection *sql.DB
}

func (sd *StorageDriver) Connect() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s/%s?sslmode=disable",
		sd.config.Options["username"],
		sd.config.Options["password"],
		sd.config.Options["host"],
		sd.config.Options["database"],
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return fmt.Errorf("fail to connect to postgres with error: %s", err)
	}

	err = db.Ping()
	if err != nil {
		return fmt.Errorf("fail to ping database with error: %s", err)
	}

	sd.connection = db

	return nil
}

func (sd *StorageDriver) Close() error {
	return sd.connection.Close()
}

func (sd *StorageDriver) Create(event model.Event) (interface{}, error) {
	return nil, nil
}

func (sd *StorageDriver) Delete(id interface{}) error {
	return nil
}

func (sd *StorageDriver) Query(params map[string]interface{}) {

}

func NewStorageDriver(conf config.StorageConfig) *StorageDriver {
	sd := &StorageDriver{
		config: conf,
	}

	return sd
}
