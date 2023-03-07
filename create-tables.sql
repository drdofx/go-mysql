DROP TABLE IF EXISTS albums;

CREATE TABLE albums (
    id INT NOT NULL AUTO_INCREMENT,
    title VARCHAR(255) NOT NULL,
    artist VARCHAR(255) NOT NULL,
    price DECIMAL(5,2) NOT NULL,
    PRIMARY KEY (id)
);

INSERT INTO albums (title, artist, price) VALUES
('The Dark Side of the Moon', 'Pink Floyd', 9.99),
('The Wall', 'Pink Floyd', 9.99),
('Wish You Were Here', 'Pink Floyd', 9.99),
('Animals', 'Pink Floyd', 9.99),
('The Piper at the Gates of Dawn', 'Pink Floyd', 9.99),
('A Saucerful of Secrets', 'Pink Floyd', 9.99);