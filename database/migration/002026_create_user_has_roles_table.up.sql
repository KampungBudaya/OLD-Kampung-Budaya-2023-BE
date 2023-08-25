CREATE TABLE IF NOT EXISTS `user_has_roles`
(
    `id`      INT PRIMARY KEY AUTO_INCREMENT,
    `user_id` INT,
    `role_id` INT
);

ALTER TABLE `user_has_roles` ADD CONSTRAINT `user_has_roles_fk_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

ALTER TABLE `user_has_roles` ADD CONSTRAINT `user_has_roles_fk_role_id` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
