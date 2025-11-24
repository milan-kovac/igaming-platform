DROP PROCEDURE IF EXISTS distribute_prizes;

CREATE PROCEDURE distribute_prizes(IN tournamentId BIGINT)
BEGIN
    DECLARE total_players INT;
    DECLARE already_distributed BOOLEAN;
    DECLARE tournament_prize_pool DECIMAL(10,2);
    DECLARE p1, p2, p3 BIGINT;

    DECLARE EXIT HANDLER FOR SQLEXCEPTION
    BEGIN
        ROLLBACK;
        RESIGNAL;
    END;

    START TRANSACTION;


    SELECT prizes_distributed, prize_pool 
    INTO already_distributed, tournament_prize_pool
    FROM tournaments
    WHERE id = tournamentId
    FOR UPDATE;

    IF already_distributed THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Prizes already distributed for this tournament';
    END IF;

   
    SELECT COUNT(DISTINCT player_id) INTO total_players
    FROM tournament_bets
    WHERE tournament_id = tournamentId;

    IF total_players < 3 THEN
        SIGNAL SQLSTATE '45000'
        SET MESSAGE_TEXT = 'Not enough participants to distribute prizes';
    END IF;

 
    SELECT 
        MAX(CASE WHEN rn = 1 THEN player_id END),
        MAX(CASE WHEN rn = 2 THEN player_id END),
        MAX(CASE WHEN rn = 3 THEN player_id END)
    INTO p1, p2, p3
    FROM (
        SELECT 
            player_id,
            ROW_NUMBER() OVER (ORDER BY SUM(bet_amount) DESC) AS rn
        FROM tournament_bets
        WHERE tournament_id = tournamentId
        GROUP BY player_id
    ) ranked
    WHERE rn <= 3;


    UPDATE players
    SET account_balance = account_balance + 
        CASE 
            WHEN id = p1 THEN tournament_prize_pool * 0.5
            WHEN id = p2 THEN tournament_prize_pool * 0.3  
            WHEN id = p3 THEN tournament_prize_pool * 0.2
            ELSE 0
        END
    WHERE id IN (p1, p2, p3);


    UPDATE tournaments
    SET prizes_distributed = TRUE
    WHERE id = tournamentId;

    COMMIT;

END;