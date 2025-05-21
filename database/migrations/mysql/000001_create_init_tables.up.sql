-- -----------------------------------------------------
-- Table users
-- -----------------------------------------------------
CREATE TABLE users (
                       id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                       email VARCHAR (255) NOT NULL UNIQUE,
                       password VARCHAR (255) NOT NULL,
                       fullname VARCHAR (255) NULL,
                       phone VARCHAR(20) NULL,
                       token VARCHAR (100) NULL,
                       status ENUM('pending', 'active', 'blocked') NOT NULL DEFAULT 'pending',
                       avatar VARCHAR (255) NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NULL,
                       verified_at TIMESTAMP NULL,
                       blocked_at TIMESTAMP NULL,
                       deleted_at TIMESTAMP NULL,
                       last_access_at TIMESTAMP NULL
);

-- Add indexes
CREATE INDEX active_users ON users (id);
CREATE UNIQUE INDEX email_users ON users (email ASC);


-- -----------------------------------------------------
-- Table roles
-- -----------------------------------------------------
CREATE TABLE roles (
                       id INT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                       name VARCHAR(100) NOT NULL,
                       slug VARCHAR(100) NOT NULL,
                       created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NULL
);

-- Add indexes
CREATE INDEX active_roles ON roles (id);

-- Insert data
INSERT INTO roles (name, slug) VALUES('Admin', 'admin');
INSERT INTO roles (name, slug) VALUES('Moderator', 'moderator');
INSERT INTO roles (name, slug) VALUES('Member', 'member');
INSERT INTO roles (name, slug) VALUES('Guest', 'guest');

-- -----------------------------------------------------
-- Table user_roles
-- -----------------------------------------------------
CREATE TABLE user_roles (
                            id BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
                            role_id INT,
                            user_id BIGINT,
                            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                            CONSTRAINT fk_user_to_role
                                FOREIGN KEY (role_id)
                                    REFERENCES roles (id)
                                    ON DELETE CASCADE,
                            CONSTRAINT fk_role_to_user
                                FOREIGN KEY (user_id)
                                    REFERENCES users (id)
                                    ON DELETE CASCADE
);

-- -----------------------------------------------------
-- Table 'address'
-- -----------------------------------------------------
CREATE TABLE address (
                         id SERIAL PRIMARY KEY,
                         user_id BIGINT,
                         type ENUM ('address', 'billing', 'shipping') NOT NULL DEFAULT 'address',
                         is_default BOOL DEFAULT TRUE,
                         address_line1 VARCHAR(150) NOT NULL,
                         address_line2 VARCHAR(150) NULL,
                         ward VARCHAR(100) NULL,
                         district VARCHAR(100) NULL,
                         city VARCHAR(100) NULL,
                         state VARCHAR(100) NULL,
                         country VARCHAR(100) NULL,
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                         updated_at TIMESTAMP NULL,
                         deleted_at TIMESTAMP NULL,
                         CONSTRAINT fk_address_users
                             FOREIGN KEY (user_id)
                                 REFERENCES users (id)
                                 ON DELETE CASCADE
);

-- --------------------------------------------------------------------------------------
-- ------------------------------------ Initial data ------------------------------------
-- --------------------------------------------------------------------------------------
-- P@seWor9  ===>  $2a$04$9QD944312deeQjnxF.zNauGx7NQ0GtS.xJhLy.zWqWxOE8B/XCN9i
INSERT INTO users (email, password, fullname, phone, token, status, avatar, created_at, updated_at)
VALUES ('admin@gfly.dev', '$2a$04$9QD944312deeQjnxF.zNauGx7NQ0GtS.xJhLy.zWqWxOE8B/XCN9i', 'Admin', '0989831911', null, 'active', 'https://www.dancefitvn.com/assets/avatar.png', '2024-05-15 13:07:48.888668 +07:00', '2024-05-15 13:07:48.888668 +07:00');

insert into user_roles (role_id, user_id, created_at)
values (1, 1, '2024-05-15 13:07:48.888668 +07:00');
