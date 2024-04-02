CREATE TABLE `notifications` (
    `id` VARCHAR(255) NOT NULL,
    `title` TEXT NOT NULL,
    `body` TEXT NOT NULL,
    `status` INT NOT NULL,
    `is_read` INT NOT NULL ,
    `account_id` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

ALTER TABLE `notifications` ADD CONSTRAINT `notifications_account_id_fkey` FOREIGN KEY (`account_id`) REFERENCES `accounts`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;