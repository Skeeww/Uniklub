CREATE TABLE clubs_users(
    user_id TEXT REFERENCES users(email) ON DELETE CASCADE,
    club_id INTEGER REFERENCES clubs(id) ON DELETE CASCADE,
    manage BOOL NOT NULL DEFAULT false,
    PRIMARY KEY (user_id, club_id)
);