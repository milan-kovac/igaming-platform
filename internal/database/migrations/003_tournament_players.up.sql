CREATE TABLE IF NOT EXISTS tournament_players (
    tournament_id BIGINT NOT NULL,
    player_id BIGINT NOT NULL,
    PRIMARY KEY (tournament_id, player_id),
    FOREIGN KEY (tournament_id) REFERENCES tournaments(id) ON DELETE CASCADE,
    FOREIGN KEY (player_id) REFERENCES players(id) ON DELETE CASCADE
);