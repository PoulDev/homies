-- !! SCONSIGLIATO !!
ALTER TABLE users DROP FOREIGN KEY link_user_house;
DROP TABLE IF EXISTS houses;
DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS avatars;
-- se scommentati eliminano COMPLETAMENTE il database :3
-- da usare in setup iniziale durante testing


CREATE TABLE users (
    id BINARY(16) NOT NULL PRIMARY KEY,
    name VARCHAR(32) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    house INT UNSIGNED,
    pwd_hash BINARY(64) NOT NULL,
    pwd_salt BINARY(32) NOT NULL,
    avatar INT UNSIGNED NOT NULL,
    is_owner BOOLEAN DEFAULT 0
);

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

ALTER TABLE 
    users ADD UNIQUE INDEX name_index (name);

ALTER TABLE users
    ADD CONSTRAINT link_user_avatar FOREIGN KEY (avatar) REFERENCES avatars(id)
    ON DELETE RESTRICT
    ON UPDATE CASCADE;

ALTER TABLE users
    ADD CONSTRAINT link_user_house FOREIGN KEY (house) REFERENCES houses(id)
    ON DELETE SET NULL
    ON UPDATE CASCADE;


