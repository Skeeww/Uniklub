CREATE TABLE categories(
    name VARCHAR(80) NOT NULL,
    is_public BOOL NOT NULL DEFAULT false,
    club_id INTEGER REFERENCES clubs(id),
    PRIMARY KEY (name, club_id)
);