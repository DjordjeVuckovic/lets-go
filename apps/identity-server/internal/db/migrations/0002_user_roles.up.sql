BEGIN;
CREATE TABLE IF NOT EXISTS identity.users
(
    id                 uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    username           VARCHAR(50) NOT NULL UNIQUE,
    password           bytea       NOT NULL,
    email              VARCHAR(50) NOT NULL UNIQUE,
    is_email_confirmed BOOLEAN          DEFAULT FALSE,
    security_stamp     uuid             DEFAULT uuid_generate_v4(),
    lock_version       uuid             DEFAULT uuid_generate_v4(),
    created_at         TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at         TIMESTAMP        DEFAULT NULL,
    deleted_at         TIMESTAMP        DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS identity.roles
(
    id         uuid PRIMARY KEY,
    name       VARCHAR(50) NOT NULL UNIQUE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS identity.user_roles
(
    id      uuid PRIMARY KEY,
    user_id uuid NOT NULL REFERENCES identity.users (id),
    role_id uuid NOT NULL REFERENCES identity.roles (id)
);

COMMIT;
END;