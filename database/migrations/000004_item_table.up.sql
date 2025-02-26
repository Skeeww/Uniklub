CREATE TYPE item_state AS ENUM ('unavailable', 'available', 'non_loanable');
CREATE TABLE items(
    name VARCHAR(80) NOT NULL,
    picture BYTEA,
    description TEXT,
    quantity INT NOT NULL CHECK(quantity > -1) DEFAULT 1,
    state item_state NOT NULL DEFAULT 'available',
    club_id INTEGER REFERENCES clubs(id) ON DELETE CASCADE,
    PRIMARY KEY (name, club_id)
);
CREATE INDEX idx_items_club ON items(club_id);