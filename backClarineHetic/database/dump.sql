CREATE TABLE IF NOT EXISTS users (
    uuid UUID PRIMARY KEY,
    username TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS conversations (
    uuid UUID PRIMARY KEY,
    user_id_a UUID NOT NULL REFERENCES users(uuid),
    user_id_b UUID NOT NULL REFERENCES users(uuid)
);

CREATE TABLE IF NOT EXISTS events (
    uuid UUID PRIMARY KEY,
    name TEXT NOT NULL,
    longitude TEXT NOT NULL,
    latitude TEXT NOT NULL,
    adress TEXT NOT NULL,
    city TEXT NOT NULL,
    start_date TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(uuid)
);

CREATE TABLE IF NOT EXISTS event_users (
    uuid UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(uuid),
    event_id UUID NOT NULL REFERENCES events(uuid)
);

CREATE TABLE IF NOT EXISTS instruments (
    uuid UUID PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS instrument_users (
    uuid UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES users(uuid),
    instrument_id UUID NOT NULL REFERENCES instruments(uuid)
);

CREATE TABLE IF NOT EXISTS messages (
    uuid UUID PRIMARY KEY,
    content TEXT NOT NULL,
    insert_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL REFERENCES users(uuid),
    conversation_id UUID NOT NULL REFERENCES conversations(uuid)
);