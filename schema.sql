-- !! SCONSIGLIATO !!
-- ALTER TABLE users DROP FOREIGN KEY link_user_house;
-- DROP TABLE IF EXISTS houses;
-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS avatars;
-- se scommentati eliminano COMPLETAMENTE il database :3
-- da usare in setup iniziale durante testing

CREATE TABLE houses (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE avatars (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    bg_color CHAR(6) NOT NULL,
    face_color CHAR(6) NOT NULL,
    face_x FLOAT NOT NULL,
    face_y FLOAT NOT NULL,
    left_eye_x FLOAT NOT NULL,
    left_eye_y FLOAT NOT NULL,
    right_eye_x FLOAT NOT NULL,
    right_eye_y FLOAT NOT NULL,
    bezier CHAR(11) NOT NULL
);

CREATE TABLE users (
    id BINARY(16) NOT NULL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    house INT UNSIGNED,
    pwd_hash BINARY(64) NOT NULL,
    pwd_salt BINARY(32) NOT NULL,
    avatar INT UNSIGNED NOT NULL,
    is_owner BOOLEAN DEFAULT 0,
    UNIQUE INDEX name_index (name),
    INDEX house_index (house),
    CONSTRAINT link_user_avatar FOREIGN KEY (avatar) REFERENCES avatars(id)
        ON DELETE RESTRICT
        ON UPDATE CASCADE,
    CONSTRAINT link_user_house FOREIGN KEY (house) REFERENCES houses(id)
        ON DELETE SET NULL
        ON UPDATE CASCADE
);

CREATE TABLE lists (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    user_id BINARY(16) NOT NULL,
    name VARCHAR(256) NOT NULL,
    items INT UNSIGNED NOT NULL DEFAULT 0,

    INDEX user_id_index (user_id),
    CONSTRAINT link_list_user FOREIGN KEY (user_id) REFERENCES users(id)
        ON DELETE CASCADE
);

CREATE TABLE todos (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    list_id INT UNSIGNED NOT NULL,
    completed BOOLEAN DEFAULT FALSE NOT NULL,
    text VARCHAR(512) NOT NULL,
    author BINARY(16) NOT NULL,
    INDEX list_id_index (list_id),
    CONSTRAINT link_todo_lists FOREIGN KEY (list_id) REFERENCES lists(id)
        ON DELETE CASCADE
);
