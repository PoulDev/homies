CREATE TABLE `houses` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(255) NOT NULL,
  `invite` VARCHAR(5) NOT NULL
);

CREATE TABLE `users` (
  `id` BINARY(16) PRIMARY KEY NOT NULL,
  `name` VARCHAR(32) UNIQUE NOT NULL,
  `house` BIGINT,
  `pwd_hash` BINARY(64) NOT NULL,
  `pwd_salt` BINARY(32) NOT NULL,
  `is_owner` BOOLEAN DEFAULT 0,
  `bg_color` CHAR(6) NOT NULL,
  `face_color` CHAR(6) NOT NULL,
  `face_x` FLOAT NOT NULL,
  `face_y` FLOAT NOT NULL,
  `left_eye_x` FLOAT NOT NULL,
  `left_eye_y` FLOAT NOT NULL,
  `right_eye_x` FLOAT NOT NULL,
  `right_eye_y` FLOAT NOT NULL,
  `bezier` CHAR(11) NOT NULL
);

CREATE TABLE `lists` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `house_id` BIGINT NOT NULL,
  `name` VARCHAR(256) NOT NULL,
  `items` INT NOT NULL DEFAULT 0
);

CREATE TABLE `todos` (
  `id` BIGINT PRIMARY KEY NOT NULL AUTO_INCREMENT,
  `list_id` BIGINT NOT NULL,
  `completed` BOOLEAN NOT NULL DEFAULT false,
  `text` VARCHAR(512) NOT NULL,
  `author` BINARY(16) NOT NULL
);

CREATE UNIQUE INDEX `invite_index` ON `houses` (`invite`);

CREATE UNIQUE INDEX `name_index` ON `users` (`name`);

CREATE INDEX `house_index` ON `users` (`house`);

CREATE INDEX `house_id_index` ON `lists` (`house_id`);

CREATE INDEX `list_id_index` ON `todos` (`list_id`);

ALTER TABLE `users` ADD CONSTRAINT `link_user_house` FOREIGN KEY (`house`) REFERENCES `houses` (`id`) ON DELETE SET NULL ON UPDATE CASCADE;

ALTER TABLE `lists` ADD CONSTRAINT `link_list_house` FOREIGN KEY (`house_id`) REFERENCES `houses` (`id`) ON DELETE CASCADE;

ALTER TABLE `todos` ADD CONSTRAINT `link_todo_lists` FOREIGN KEY (`list_id`) REFERENCES `lists` (`id`) ON DELETE CASCADE;
