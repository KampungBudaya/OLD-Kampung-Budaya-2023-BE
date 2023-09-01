CREATE TABLE IF NOT EXISTS `payment_photos`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT,
    `user_id`     INT UNIQUE,
    `link_photo`  VARCHAR(255),

    `created_at`  DATETIME     NOT NULL DEFAULT (NOW()),
    `updated_at`  DATETIME     NOT NULL DEFAULT (NOW())
);

ALTER TABLE `payment_photos` ADD CONSTRAINT `payment_photos_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);
