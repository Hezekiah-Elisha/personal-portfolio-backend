-- Create users table with unique constraints on username and email
CREATE TABLE users (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR2(50) NOT NULL UNIQUE,
    email VARCHAR2(100) NOT NULL UNIQUE,
    password_hash VARCHAR2(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- CREATE UNIQUE INDEX "idx_users_username" ON "users" ("username");
CREATE UNIQUE INDEX "idx_users_username" ON "users" ("username");

-- CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
CREATE UNIQUE INDEX "idx_users_email" ON "users" ("email");
