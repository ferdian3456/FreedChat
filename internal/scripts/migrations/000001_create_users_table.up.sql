CREATE TABLE IF NOT EXISTS users(
    id uuid  PRIMARY KEY,
    username  varchar(50) unique NOT NULL,
    password  varchar(60) NOT NULL,
    profile_picture varchar(255),
    created_at timestamp NOT NULL,
    updated_at timestamp
);
