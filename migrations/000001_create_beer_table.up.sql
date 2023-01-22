CREATE TABLE IF NOT EXISTS beer (
    id INT auto_increment PRIMARY KEY,
    name VARCHAR(255),
    created_at DATETIME,
    last_updated DATETIME,
    version INT
)