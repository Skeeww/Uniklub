package database

import (
	"context"
	"fmt"
	"sync"

	"github.com/jackc/pgx/v5"
)

var databaseInstance *pgx.Conn = nil
var databaseInstanceMutex *sync.Mutex = &sync.Mutex{}

type ConnectionInformation struct {
	Username string
	Password string
	Address  string
	Port     int
	Database string
}

func CreateConnection(ctx context.Context, connInfo ConnectionInformation) error {
	databaseInstanceMutex.Lock()
	defer databaseInstanceMutex.Unlock()

	if databaseInstance != nil {
		return fmt.Errorf("an instance has already been created")
	}

	driver, err := connect(ctx, connInfo)
	if err != nil {
		return err
	}

	databaseInstance = driver
	return nil
}

func GetConnection() *pgx.Conn {
	databaseInstanceMutex.Lock()
	defer databaseInstanceMutex.Unlock()

	return databaseInstance
}

func connect(ctx context.Context, connInfo ConnectionInformation) (*pgx.Conn, error) {
	driver, err := pgx.Connect(ctx, fmt.Sprintf("postgres://%s:%s@%s:%d/%s", connInfo.Username, connInfo.Password, connInfo.Address, connInfo.Port, connInfo.Database))
	if err != nil {
		return nil, err
	}

	if err := driver.Ping(ctx); err != nil {
		return nil, err
	}

	return driver, nil
}
