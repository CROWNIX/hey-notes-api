CREATE TABLE chattings (
    `id` BIGINT UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    `group_id` BIGINT UNSIGNED NOT NULL,
    `from` BIGINT UNSIGNED NOT NULL,
    `message` LONGTEXT NOT NULL,
    `reply_id` BIGINT UNSIGNED DEFAULT NULL,
);
