package database

type schemaItem struct {
	name  string
	query string
}

var sqlCreateEnums = []schemaItem{
	{
		name:  "item_state",
		query: "CREATE TYPE item_state AS ENUM ('unavailable', 'available', 'non_loanable')",
	},
}

var sqlCreateTables = []schemaItem{
	{
		name: "users",
		query: `
			CREATE TABLE IF NOT EXISTS users(
				email TEXT PRIMARY KEY,
				name VARCHAR(50) NOT NULL,
				surname VARCHAR(50) NOT NULL,
				password TEXT NOT NULL
			)
		`,
	},
	{
		name: "clubs",
		query: `
			CREATE TABLE IF NOT EXISTS clubs(
				id SERIAL PRIMARY KEY,
				name VARCHAR(80) NOT NULL
			)
		`,
	},
	{
		name: "clubs_members",
		query: `
			CREATE TABLE IF NOT EXISTS clubs_members(
				user_id TEXT REFERENCES users(email) ON DELETE CASCADE,
				club_id INTEGER REFERENCES clubs(id) ON DELETE CASCADE,
				manage BOOL NOT NULL DEFAULT false,
				PRIMARY KEY (user_id, club_id)
			)
		`,
	},
	{
		name: "items",
		query: `
			CREATE TABLE IF NOT EXISTS items(
				name VARCHAR(80) NOT NULL,
				picture BYTEA,
				description TEXT,
				quantity INT NOT NULL CHECK(quantity > -1) DEFAULT 1,
				state item_state NOT NULL DEFAULT 'available',
				club_id INTEGER REFERENCES clubs(id) ON DELETE CASCADE,
				PRIMARY KEY (name, club_id)
			)
		`,
	},
	{
		name: "categories",
		query: `
			CREATE TABLE IF NOT EXISTS categories(
				name VARCHAR(80) NOT NULL,
				is_public BOOL NOT NULL DEFAULT false,
				club_id INTEGER REFERENCES clubs(id),
				PRIMARY KEY (name, club_id)
			)
		`,
	},
}

var sqlCreateIndexes = []schemaItem{
	{
		name:  "items_club",
		query: "CREATE INDEX IF NOT EXISTS idx_items_club ON items(club_id)",
	},
}
