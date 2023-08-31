CREATE TABLE IF NOT EXISTS `photos`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT,
    `link_photo`  VARCHAR(255),
    `user_id`     INT UNIQUE,

    `created_at`  DATETIME     NOT NULL DEFAULT (NOW()),
    `updated_at`  DATETIME     NOT NULL DEFAULT (NOW()),

    FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE
);
