CREATE TABLE IF NOT EXISTS users(
   user_id serial PRIMARY KEY,
   data VARCHAR (50) UNIQUE NOT NULL
);