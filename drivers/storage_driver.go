package drivers

import (
	_ "github.com/lib/pq"
	"fmt"
	"database/sql"
	"github.com/Liar233/Scheduler/config"
	"github.com/Liar233/Scheduler/model"
	"time"
)

const insertQuery = "INSERT INTO events (channel, firetime, payload) VALUES ($1, $2, $3) RETURNING id;"
const deleteQuery = "DELETE FROM events WHERE id = $1;"
const selectQuery = "SELECT * FROM events $1;"

type StorageDriver struct {
	config     config.StorageConfig
	connection *sql.DB
}

func (sd *StorageDriver) Connect() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disable",
		sd.config.Options["username"],
		sd.config.Options["password"],
		sd.config.Options["host"],
		sd.config.Options["port"],
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
	var id int

	err := sd.connection.QueryRow(insertQuery, event.Channel, event.FireTime, event.Payload).Scan(&id)

	if err != nil {
		return nil, err
	}

	return id, nil
}

func (sd *StorageDriver) Delete(id interface{}) error {
	_, err := sd.connection.Exec(deleteQuery, id)

	return err
}

func (sd *StorageDriver) Query(params map[string]interface{}) ([]model.Event, error) {
	//var conditions = ""

	results := make([]model.Event, 0)

	if len(params) != 0 {

	}

	rows, err := sd.connection.Query("SELECT * FROM events;")
	defer rows.Close()

	if err != nil {
		return results, err
	}

	for rows.Next() {
		var (
			id, channel string
			firetime    time.Time
			payload     []byte
		)

		err := rows.Scan(&id, &channel, &firetime, &payload)

		if err != nil {
			return results, err
		}

		event := model.Event{
			ID:       fmt.Sprintf("%v", id),
			Channel:  channel,
			FireTime: firetime,
			Payload:  payload,
		}

		results = append(results, event)
	}

	return results, nil
}

func NewStorageDriver(conf config.StorageConfig) *StorageDriver {
	sd := &StorageDriver{
		config: conf,
	}

	return sd
}
