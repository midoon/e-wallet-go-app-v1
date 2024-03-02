CREATE TABLE `accounts` (
    `id` VARCHAR(255) NOT NULL,
    `account_number` VARCHAR(20) NOT NULL,
    `pin`  INT NOT NULL,
    `balance` FLOAT NOT NULL,
    `user_id` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    `updated_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    UNIQUE INDEX `accounts_user_id_key`(`user_id`),
    UNIQUE INDEX `accounts_account_number_key`(`account_number`),
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

ALTER TABLE `accounts` ADD CONSTRAINT `accounts_user_id_fkey` FOREIGN KEY (`user_id`) REFERENCES `users`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;