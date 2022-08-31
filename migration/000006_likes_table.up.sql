CREATE TABLE IF NOT EXISTS `post_likes`
(
    `post_id`    INT UNSIGNED NOT NULL,
    `user_id`    INT UNSIGNED NOT NULL,
    `created_at` DATETIME     NOT NULL,
    PRIMARY KEY (`post_id`, `user_id`),
    INDEX `post_id_fk_idx` (`post_id` ASC) VISIBLE,
    INDEX `user_id_fk_idx` (`user_id` ASC) VISIBLE,
    CONSTRAINT `post_id_fk`
        FOREIGN KEY (`post_id`)
            REFERENCES `posts` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `likes_user_id_fk`
        FOREIGN KEY (`user_id`)
            REFERENCES `users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;