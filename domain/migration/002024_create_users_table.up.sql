CREATE TABLE IF NOT EXISTS `users`
(
    `id`          INT PRIMARY KEY AUTO_INCREMENT,
    `provider`    VARCHAR(100),
    `provider_id` VARCHAR(255) UNIQUE,

    `username`    VARCHAR(255) NOT NULL,
    `email`       VARCHAR(255) NOT NULL UNIQUE,
    `password`    VARCHAR(255),

    `created_at`  DATETIME     NOT NULL DEFAULT (NOW()),
    `updated_at`  DATETIME     NOT NULL DEFAULT (NOW())
);
