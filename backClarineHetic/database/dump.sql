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




INSERT INTO users (uuid, username, email, password) VALUES
    ('e2a3f4d0-7a3b-4f93-9d7a-12d8c1a4c7f2', 'Alice', 'alice@example.com', '$2a$10$kLvycA8dYZ.JnkIJsO1PJOSTIrxPh4CJ2pv0iWoQxXbmLlbaODtFW'),
    ('b7d4f98c-1a52-432a-a6b7-5cae4f82e9a1', 'Bob', 'bob@example.com', '$2a$10$kLvycA8dYZ.JnkIJsO1PJOSTIrxPh4CJ2pv0iWoQxXbmLlbaODtFW'),
    ('c8de7a23-75f1-4d19-b97e-8bb8196d1a59', 'Charlie', 'charlie@example.com', '$2a$10$kLvycA8dYZ.JnkIJsO1PJOSTIrxPh4CJ2pv0iWoQxXbmLlbaODtFW'),
    ('3357a9b4-9532-44ff-aa60-ddd0296e0272', 'Adrien', 'adrien@mail.com', '$2a$10$kLvycA8dYZ.JnkIJsO1PJOSTIrxPh4CJ2pv0iWoQxXbmLlbaODtFW');

INSERT INTO conversations (uuid, user_id_a, user_id_b) VALUES
    ('a117f4cd-8e3c-4b9d-9dae-5f0b9f0b3f2e', 'e2a3f4d0-7a3b-4f93-9d7a-12d8c1a4c7f2', 'b7d4f98c-1a52-432a-a6b7-5cae4f82e9a1');

INSERT INTO conversations (uuid, user_id_a, user_id_b) VALUES
    ('e1b0f5e9-6f3d-41a4-9c3d-1b2a2b3c4d5e', 'b7d4f98c-1a52-432a-a6b7-5cae4f82e9a1', 'c8de7a23-75f1-4d19-b97e-8bb8196d1a59');

INSERT INTO events (uuid, name, longitude, latitude, adress, city, start_date, user_id) VALUES
    ('f3d7c4e1-51b9-4d73-9286-3a5d6f7e8b9a', 'Concert', '2.3522', '48.8566', '123 Rue de la Musique', 'Paris', '2023-12-01 20:00:00', 'e2a3f4d0-7a3b-4f93-9d7a-12d8c1a4c7f2');

INSERT INTO events (uuid, name, longitude, latitude, adress, city, start_date, user_id) VALUES
    ('d4c8e2f3-2e7b-4a9d-827d-6f5c8e3a1b2d', 'Exposition', '4.8357', '45.7640', '45 Avenue des Arts', 'Lyon', '2023-12-05 18:00:00', 'b7d4f98c-1a52-432a-a6b7-5cae4f82e9a1');

INSERT INTO event_users (uuid, user_id, event_id) VALUES
    ('c1b2a3d4-5e6f-7a89-b0c1-d2e3f4a5b6c7', 'b7d4f98c-1a52-432a-a6b7-5cae4f82e9a1', 'f3d7c4e1-51b9-4d73-9286-3a5d6f7e8b9a');

INSERT INTO event_users (uuid, user_id, event_id) VALUES
    ('d8c7b6a5-4e3f-2d1c-0b9a-8f7e6d5c4b3a', 'c8de7a23-75f1-4d19-b97e-8bb8196d1a59', 'd4c8e2f3-2e7b-4a9d-827d-6f5c8e3a1b2d');

INSERT INTO instruments (uuid, name) VALUES
    ('a0b1c2d3-e4f5-6789-abcd-ef0123456789', 'Guitare'),
    ('b1c2d3e4-f5a6-7890-bcde-f0123456789a', 'Piano');

INSERT INTO instrument_users (uuid, user_id, instrument_id) VALUES
    ('c2d3e4f5-a6b7-8901-cdef-0123456789ab', 'e2a3f4d0-7a3b-4f93-9d7a-12d8c1a4c7f2', 'a0b1c2d3-e4f5-6789-abcd-ef0123456789');

INSERT INTO instrument_users (uuid, user_id, instrument_id) VALUES
    ('d3e4f5a6-b7c8-9012-def0-123456789abc', 'b7d4f98c-1a52-432a-a6b7-5cae4f82e9a1', 'b1c2d3e4-f5a6-7890-bcde-f0123456789a');

INSERT INTO messages (uuid, content, insert_at, user_id, conversation_id) VALUES
    ('e4f5a6b7-c8d9-0123-ef45-6789abcdef01', 'Salut Bob, comment ça va ?', '2023-11-30 10:00:00', 'e2a3f4d0-7a3b-4f93-9d7a-12d8c1a4c7f2', 'a117f4cd-8e3c-4b9d-9dae-5f0b9f0b3f2e'),
    ('f5a6b7c8-d9e0-1234-f567-89abcdef0123', 'Ça va bien, merci Alice !', '2023-11-30 10:05:00', 'b7d4f98c-1a52-432a-a6b7-5cae4f82e9a1', 'a117f4cd-8e3c-4b9d-9dae-5f0b9f0b3f2e');
