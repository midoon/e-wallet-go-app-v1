CREATE TABLE `transactions` (
    `id` VARCHAR(255) NOT NULL,
    `sof_number` VARCHAR(20) NOT NULL,
    `dof_number` VARCHAR(20) NOT NULL,
    `amount` FLOAT NOT NULL,
    `transaction_type` VARCHAR(1) NOT NULL ,
    `account_id` VARCHAR(255) NOT NULL,
    `created_at` TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    
    PRIMARY KEY (`id`)
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

ALTER TABLE `transactions` ADD CONSTRAINT `transactions_account_id_fkey` FOREIGN KEY (`account_id`) REFERENCES `accounts`(`id`) ON DELETE CASCADE ON UPDATE CASCADE;