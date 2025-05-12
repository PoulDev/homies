-- !! SCONSIGLIATO !!
-- DROP TABLE IF EXISTS users;
-- DROP TABLE IF EXISTS houses;
-- DROP TABLE IF EXISTS avatars;
-- se scommentati eliminano COMPLETAMENTE il database :3
-- da usare in setup iniziale durante testing


CREATE TABLE users (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    house INT UNSIGNED,
    pwd_hash CHAR(64) NOT NULL,
    pwd_salt CHAR(32) NOT NULL,
    avatar INT UNSIGNED NOT NULL
);

CREATE TABLE houses (
    id INT UNSIGNED NOT NULL PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(255) NOT NULL,
    owner_id INT UNSIGNED
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

-- `ON DELETE CASCADE` se l'utente owner elimina l'account anche la casa viene eliminata
ALTER TABLE houses
    ADD CONSTRAINT link_owner_user FOREIGN KEY (owner_i) REFERENCES users(id)
    ON DELETE CASCADE
    ON UPDATE CASCADE;


