DROP PROCEDURE IF EXISTS distribute_prizes;


CREATE PROCEDURE distribute_prizes(IN tournamentId BIGINT)
BEGIN
    DECLARE p1, p2, p3 BIGINT;
    DECLARE total_players INT;

    SELECT COUNT(*) INTO total_players 
    FROM tournament_players 
    WHERE tournament_id = tournamentId;

    IF total_players < 3 THEN
        SIGNAL SQLSTATE '45000' 
        SET MESSAGE_TEXT = 'Not enough participants to distribute prizes';
    ELSE
        START TRANSACTION;

        SELECT player_id INTO p1 
        FROM tournament_players tp
        JOIN players p ON tp.player_id = p.id
        WHERE tp.tournament_id = tournamentId
        ORDER BY p.account_balance DESC
        LIMIT 1 OFFSET 0
        FOR UPDATE;

        SELECT player_id INTO p2 
        FROM tournament_players tp
        JOIN players p ON tp.player_id = p.id
        WHERE tp.tournament_id = tournamentId
        ORDER BY p.account_balance DESC
        LIMIT 1 OFFSET 1
        FOR UPDATE;

        SELECT player_id INTO p3 
        FROM tournament_players tp
        JOIN players p ON tp.player_id = p.id
        WHERE tp.tournament_id = tournamentId
        ORDER BY p.account_balance DESC
        LIMIT 1 OFFSET 2
        FOR UPDATE;

      
        UPDATE players 
        SET account_balance = account_balance + (SELECT prize_pool*0.5 FROM tournaments WHERE id = tournamentId) 
        WHERE id = p1;

        UPDATE players 
        SET account_balance = account_balance + (SELECT prize_pool*0.3 FROM tournaments WHERE id = tournamentId) 
        WHERE id = p2;

        UPDATE players 
        SET account_balance = account_balance + (SELECT prize_pool*0.2 FROM tournaments WHERE id = tournamentId) 
        WHERE id = p3;

        COMMIT;
    END IF;
END
