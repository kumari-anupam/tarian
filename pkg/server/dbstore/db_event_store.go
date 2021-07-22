package dbstore

import (
	"context"
	"encoding/json"
	"time"

	"github.com/devopstoday11/tarian/pkg/tarianpb"
	"github.com/driftprogramming/pgxpoolmock"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/jackc/pgx/v4/pgxpool"
)

// DbEventStore implements store.EventStore
type DbEventStore struct {
	// pool holds the connection pool. It uses a compatible interface with the real PgxPool.
	// This is to make it mockable.
	pool pgxpoolmock.PgxPool
}

func NewDbEventStore(dsn string) (*DbEventStore, error) {
	// TODO: pass context from param?
	dbpool, err := pgxpool.Connect(context.Background(), dsn)

	if err != nil {
		return nil, err
	}

	store := &DbEventStore{pool: dbpool}

	return store, nil
}

// eventRow represents a row of database table events
// Fields are exported because to work around type
// being a reserved name.
type eventRow struct {
	Id              int
	Type            string
	ServerTimestamp time.Time
	ClientTimestamp time.Time
	Targets         string
}

func (e *eventRow) toEvent() *tarianpb.Event {
	event := &tarianpb.Event{}
	event.Type = e.Type
	event.ServerTimestamp = timestamppb.New(e.ServerTimestamp)
	event.ClientTimestamp = timestamppb.New(e.ClientTimestamp)
	json.Unmarshal([]byte(e.Targets), &event.Targets)

	return event
}

func (d *DbEventStore) GetAll() ([]*tarianpb.Event, error) {
	rows, err := d.pool.Query(context.Background(), "SELECT * FROM events")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	allEvents := []*tarianpb.Event{}

	for rows.Next() {
		e := eventRow{}

		err := rows.Scan(&e.Id, &e.Type, &e.ServerTimestamp, &e.ClientTimestamp, &e.Targets)
		if err != nil {
			// TODO: logger.Errorw()

			continue
		}

		allEvents = append(allEvents, e.toEvent())
	}

	return allEvents, nil
}

func (d *DbEventStore) Add(event *tarianpb.Event) error {
	var id int
	targetsJson, err := json.Marshal(event.GetTargets())
	if err != nil {
		return err
	}

	err = d.pool.
		QueryRow(
			context.Background(),
			"INSERT INTO events(type, server_timestamp, client_timestamp, targets) VALUES($1, $2, $3, $4) RETURNING id",
			event.GetType(), event.GetServerTimestamp().AsTime(), event.GetClientTimestamp().AsTime(), targetsJson).
		Scan(&id)
	if err != nil {
		return err
	}

	return nil
}