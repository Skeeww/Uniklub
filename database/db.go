package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type ConnectionInformation struct {
	Username string
	Password string
	Address  string
	Port     int
	Database string
}

func Init(ctx context.Context, connInfo ConnectionInformation) (*pgx.Conn, error) {
	driver, err := connect(ctx, connInfo)
	if err != nil {
		return nil, err
	}

	for _, item := range sqlCreateEnums {
		fmt.Println("creating type", item.name)
		_, err := driver.Exec(ctx, item.query)
		if err != nil {
			fmt.Println("warn:", err.Error())
		}
	}

	for _, item := range sqlCreateTables {
		fmt.Println("creating table", item.name)
		_, err := driver.Exec(ctx, item.query)
		if err != nil {
			fmt.Println("warn:", err.Error())
		}
	}

	for _, item := range sqlCreateIndexes {
		fmt.Println("creating index", item.name)
		_, err := driver.Exec(ctx, item.query)
		if err != nil {
			fmt.Println("warn:", err.Error())
		}
	}

	return driver, nil
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
