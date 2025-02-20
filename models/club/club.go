package club

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"noan.dev/uniklub/constants"
)

const table = "clubs"

type Club struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
type ClubPrimaryKey struct {
	Id int
}
type ClubCreationFields struct {
	Name string
}
type ClubUpdateFields struct {
	Name string
}

func Find(ctx context.Context, pk ClubPrimaryKey) *Club {
	driver := ctx.Value(constants.DatabaseCtx).(*pgx.Conn)

	row := driver.QueryRow(ctx, fmt.Sprintf("SELECT * FROM %s WHERE id = $1", table), pk.Id)

	var club Club
	if err := row.Scan(&club.Id, &club.Name); err != nil {
		fmt.Println("warn:", err.Error())
		return nil
	}

	return &club
}

func FindAll(ctx context.Context) []*Club {
	clubs := make([]*Club, 0)
	driver := ctx.Value(constants.DatabaseCtx).(*pgx.Conn)

	rows, err := driver.Query(ctx, fmt.Sprintf("SELECT id, name FROM %s", table))
	if err != nil {
		fmt.Println("warn:", err.Error())
		return clubs
	}
	defer rows.Close()

	for rows.Next() {
		var club Club
		if err := rows.Scan(&club.Id, &club.Name); err != nil {
			fmt.Println("warn:", err.Error())
			continue
		}
		clubs = append(clubs, &club)
	}

	return clubs
}

func Create(ctx context.Context, fields ClubCreationFields) *Club {
	driver := ctx.Value(constants.DatabaseCtx).(*pgx.Conn)
	row := driver.QueryRow(ctx, fmt.Sprintf("INSERT INTO %s(name) VALUES ($1) RETURNING *", table), fields.Name)

	var club Club
	if err := row.Scan(&club.Id, &club.Name); err != nil {
		fmt.Println("warn:", err.Error())
		return nil
	}

	return &club
}

func Update(ctx context.Context, pk ClubPrimaryKey, fields ClubUpdateFields) *Club {
	driver := ctx.Value(constants.DatabaseCtx).(*pgx.Conn)
	row := driver.QueryRow(ctx, fmt.Sprintf("UPDATE %s SET name = $1 WHERE id = $2 RETURNING *", table), fields.Name, pk.Id)

	var club Club
	if err := row.Scan(&club.Id, &club.Name); err != nil {
		fmt.Println("warn:", err.Error())
		return nil
	}

	return &club
}

func Delete(ctx context.Context, pk ClubPrimaryKey) error {
	driver := ctx.Value(constants.DatabaseCtx).(*pgx.Conn)

	rows, err := driver.Query(ctx, fmt.Sprintf("DELETE FROM %s WHERE id = $1", table), pk.Id)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
