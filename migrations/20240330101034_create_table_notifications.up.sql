CREATE TABLE `notifications` (
    `id` VARCHAR(255) NOT NULL,
    `title` TEXT NOT NULL,
    `body` TEXT NOT NULL,
    `status` INT NOT NULL,
    `is_read` INT NOT NULL ,
    `user_id` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

ALTER TABLE `notifications` ADD CONSTRAINT `notifications_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;