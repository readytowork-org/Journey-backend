CREATE TABLE IF NOT EXISTS comments
(
    id           INT UNSIGNED NOT NULL AUTO_INCREMENT,
    comment      VARCHAR(500) NOT NULL,
    post_id      INT UNSIGNED NOT NULL,
    likes        INT          NULL DEFAULT 0,
    parent_id_fk INT UNSIGNED NULL,
    user_id      INT UNSIGNED NOT NULL,
    create_at    DATETIME     NOT NULL,
    updated_at   DATETIME     NULL,
    deleted_at   DATETIME     NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX comment_id_unique (id ASC) VISIBLE,
    INDEX comment_id_fk_idx (parent_id_fk ASC) VISIBLE,
    INDEX post_id_fk_idx (post_id ASC) VISIBLE,
    INDEX user_id_fk_idx (user_id ASC) VISIBLE,
    CONSTRAINT comment_id_fk
        FOREIGN KEY (parent_id_fk)
            REFERENCES comments (id)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT post_id_fk
        FOREIGN KEY (post_id)
            REFERENCES posts (post_id)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION,
    CONSTRAINT user_id_fk
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;