CREATE TABLE IF NOT EXISTS chats(
    id serial PRIMARY KEY,
    sender_id uuid NOT NULL,
    receiver_id uuid NOT NULL,
    message varchar(255) NOT NULL,
    created_at timestamp NOT NULL
);
