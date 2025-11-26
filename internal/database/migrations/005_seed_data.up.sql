INSERT INTO players (id, name, email, account_balance)
VALUES
    (1, 'Alice', 'alice@example.com', 100),
    (2, 'Bob', 'bob@example.com', 100),
    (3, 'Charlie', 'charlie@example.com', 100),
    (4, 'Diana', 'diana@example.com', 100),
    (5, 'Eve', 'eve@example.com', 100),
    (6, 'Frank', 'frank@example.com', 100),
    (7, 'Grace', 'grace@example.com', 100),
    (8, 'Hank', 'hank@example.com', 100),
    (9, 'Ivy', 'ivy@example.com', 100),
    (10, 'Jack', 'jack@example.com', 100);

INSERT INTO tournaments (id, name, prize_pool, start_date, end_date)
VALUES
    (1, 'Tournament 1', 10000, '2025-11-01', '2025-11-05'),
    (2, 'Tournament 2', 11000, '2025-11-06', '2025-11-10'),
    (3, 'Tournament 3', 12000, '2025-11-11', '2025-11-15'),
    (4, 'Tournament 4', 13000, '2025-11-16', '2025-11-20'),
    (5, 'Tournament 5', 14000, '2025-11-21', '2025-11-25'),
    (6, 'Tournament 6', 15000, '2025-11-26', '2025-11-30'),
    (7, 'Tournament 7', 16000, '2025-12-01', '2025-12-05'),
    (8, 'Tournament 8', 17000, '2025-12-06', '2025-12-10'),
    (9, 'Tournament 9', 18000, '2025-12-11', '2025-12-15'),
    (10, 'Tournament 10', 20000, '2025-12-16', '2025-12-20');

INSERT INTO tournament_bets (tournament_id, player_id, bet_amount)
VALUES
    (1, 1, 50), (1, 3, 100), (1, 5, 50), (1, 7, 200),
    (2, 2, 100), (2, 4, 200), (2, 6, 50),
    (3, 1, 100), (3, 2, 50), (3, 3, 50), (3, 4, 200),
    (4, 5, 100), (4, 6, 50), (4, 7, 50), (4, 8, 100),
    (5, 1, 50), (5, 2, 100), (5, 9, 200),
    (6, 3, 100), (6, 4, 200), (6, 10, 100),
    (7, 1, 200), (7, 5, 100), (7, 6, 50), (7, 7, 200), (7, 8, 100),
    (8, 2, 50), (8, 3, 100), (8, 4, 200), (8, 9, 50),
    (9, 5, 200), (9, 6, 50), (9, 7, 100), (9, 10, 200),
    (10, 1, 200), (10, 2, 100), (10, 3, 50), (10, 4, 200),
    (10, 5, 100), (10, 6, 50), (10, 7, 100);