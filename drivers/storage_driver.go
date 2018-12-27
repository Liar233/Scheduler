package drivers

import (
	_ "github.com/lib/pq"
	"fmt"
	"database/sql"
	"github.com/Liar233/Scheduler/config"
	"github.com/Liar233/Scheduler/model"
	"time"
	"strconv"
	"strings"
)

const insertQuery = "INSERT INTO events (channel, firetime, payload) VALUES ($1, $2, $3) RETURNING id;"
const deleteQuery = "DELETE FROM events WHERE id = $1;"
const selectQuery = "SELECT * FROM events %s;"
const getQuery = "SELECT * FROM events WHERE id=$1 LIMIT 1;"

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

func (sd *StorageDriver) Create(event model.Event) (string, error) {
	var id int

	err := sd.connection.QueryRow(insertQuery, event.Channel, event.FireTime, event.Payload).Scan(&id)

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%v", id), nil
}

func (sd *StorageDriver) Delete(id string) error {
	_, err := sd.connection.Exec(deleteQuery, id)

	return err
}

func (sd *StorageDriver) Query(params map[string]map[string]interface{}) ([]model.Event, error) {
	results := make([]model.Event, 0)

	var args []interface{}
	var conditions []string

	if len(params) != 0 {
		var i int = 1

		for field, operators := range params {
			for operator, value := range operators {
				condition := fmt.Sprintf(" %s %s $%d ", field, operator, i)

				conditions = append(conditions, condition)
				args = append(args, value)

				i++
			}
		}
	} else {
		conditions = append(conditions, "")
	}

	t := fmt.Sprintf(" WHERE %s", strings.Join(conditions," AND "))
	query := fmt.Sprintf(selectQuery, t);

	rows, err := sd.connection.Query(query, args...)

	if err != nil {
		println(err)

		return results, err
	}

	defer rows.Close()
	for rows.Next() {
		var (
			id, channel string
			firetime    time.Time
			payload     string
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

func (sd *StorageDriver) Get(id string) (model.Event, error) {
	event := model.Event{}

	intId, err := strconv.ParseInt(id, 10, 32)

	row := sd.connection.QueryRow(getQuery, intId)

	var (
		eventId, channel string
		firetime         time.Time
		payload          string
	)

	err = row.Scan(&eventId, &channel, &firetime, &payload)

	if err != nil {
		return event, err
	}

	event.ID = fmt.Sprintf("%v", eventId)
	event.Channel = channel
	event.FireTime = firetime
	event.Payload = payload

	return event, nil
}

func NewStorageDriver(conf config.StorageConfig) *StorageDriver {
	sd := &StorageDriver{
		config: conf,
	}

	return sd
}
