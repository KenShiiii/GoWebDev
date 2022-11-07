## create table in db vinylshop
```postgresql
CREATE TABLE albums (
    ID SERIAL PRIMARY KEY NOT NULL,    
    title  TEXT,
    artist TEXT,
    price REAL
  );
``` 

## insert initial albums
```postgresql
INSERT INTO albums (title, artist, price) VALUES ('Blue Train', 'John Coltrane', 56.99),('Jeru', 'Gerry Mulligan', 17.99),('Sarah Vaughan and Clifford Brown', 'Sarah Vaughan', 39.99);
```