DROP PROCEDURE IF EXISTS distribute_prizes;

CREATE PROCEDURE distribute_prizes(IN tournamentId BIGINT)
BEGIN
    DECLARE total_players INT;
    DECLARE p1, p2, p3 BIGINT;

    SELECT player_id INTO p1 FROM players ORDER BY account_balance DESC LIMIT 1 OFFSET 0;
    SELECT player_id INTO p2 FROM players ORDER BY account_balance DESC LIMIT 1 OFFSET 1;
    SELECT player_id INTO p3 FROM players ORDER BY account_balance DESC LIMIT 1 OFFSET 2;

    UPDATE players SET account_balance = account_balance + (SELECT prize_pool*0.5 FROM tournaments WHERE id=tournamentId) WHERE id = p1;
    UPDATE players SET account_balance = account_balance + (SELECT prize_pool*0.3 FROM tournaments WHERE id=tournamentId) WHERE id = p2;
    UPDATE players SET account_balance = account_balance + (SELECT prize_pool*0.2 FROM tournaments WHERE id=tournamentId) WHERE id = p3;
END;
