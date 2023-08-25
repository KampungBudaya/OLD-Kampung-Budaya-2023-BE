CREATE TABLE IF NOT EXISTS `roles`
(
    `id`         INT PRIMARY KEY AUTO_INCREMENT,

    `name`       VARCHAR(100) UNIQUE,

    `created_at` DATETIME DEFAULT (NOW()),
    `updated_at` DATETIME DEFAULT (NOW())
);
