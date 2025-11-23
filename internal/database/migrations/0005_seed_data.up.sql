INSERT INTO players (id, name, email, account_balance)
VALUES
    (1, 'Alice', 'alice@example.com', 1000),
    (2, 'Bob', 'bob@example.com', 1200),
    (3, 'Charlie', 'charlie@example.com', 800),
    (4, 'Diana', 'diana@example.com', 1500),
    (5, 'Eve', 'eve@example.com', 1100),
    (6, 'Frank', 'frank@example.com', 900),
    (7, 'Grace', 'grace@example.com', 1300),
    (8, 'Hank', 'hank@example.com', 950),
    (9, 'Ivy', 'ivy@example.com', 1250),
    (10, 'Jack', 'jack@example.com', 1050);


INSERT INTO tournaments (id, name, prize_pool, start_date, end_date)
VALUES
    (1, 'Tournament 1', 1000, '2025-11-01', '2025-11-05'),
    (2, 'Tournament 2', 1200, '2025-11-06', '2025-11-10'),
    (3, 'Tournament 3', 1500, '2025-11-11', '2025-11-15'),
    (4, 'Tournament 4', 1100, '2025-11-16', '2025-11-20'),
    (5, 'Tournament 5', 1300, '2025-11-21', '2025-11-25'),
    (6, 'Tournament 6', 1400, '2025-11-26', '2025-11-30'),
    (7, 'Tournament 7', 1600, '2025-12-01', '2025-12-05'),
    (8, 'Tournament 8', 1700, '2025-12-06', '2025-12-10'),
    (9, 'Tournament 9', 1800, '2025-12-11', '2025-12-15'),
    (10, 'Tournament 10', 2000, '2025-12-16', '2025-12-20');


INSERT INTO tournament_players (tournament_id, player_id)
VALUES
    (1, 1), (1, 3), (1, 5), (1, 7),
    (2, 2), (2, 4), (2, 6),
    (3, 1), (3, 2), (3, 3), (3, 4),
    (4, 5), (4, 6), (4, 7), (4, 8),
    (5, 1), (5, 2), (5, 9),
    (6, 3), (6, 4), (6, 10),
    (7, 1), (7, 5), (7, 6), (7, 7), (7, 8),
    (8, 2), (8, 3), (8, 4), (8, 9),
    (9, 5), (9, 6), (9, 7), (9, 10),
    (10, 1), (10, 2), (10, 3), (10, 4), (10, 5), (10, 6), (10, 7);