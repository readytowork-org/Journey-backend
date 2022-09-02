CREATE TABLE IF NOT EXISTS `comment_likes`
(
    `user_id`    VARCHAR(50) NOT NULL,
    `comment_id` INT UNSIGNED    NOT NULL,
    `created_at` DATETIME     NOT NULL,
    PRIMARY KEY ( `user_id`,`comment_id`),
    INDEX `cl_user_id_fk_idx` (`user_id` ASC) VISIBLE,
    INDEX `cl_comment_id_fk_idx` (`comment_id` ASC) VISIBLE,
    
    CONSTRAINT `cl_comment_id_fk`
        FOREIGN KEY (`comment_id`)
            REFERENCES `comments` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT `likes_cl_user_id_fk`
        FOREIGN KEY (`user_id`)
            REFERENCES `users` (`id`)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;