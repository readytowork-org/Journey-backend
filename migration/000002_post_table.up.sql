CREATE TABLE IF NOT EXISTS posts
(
    id         INT UNSIGNED                            NOT NULL AUTO_INCREMENT,
    title      VARCHAR(128)                            NOT NULL,
    caption    TEXT                                    NULL,
    user_id    VARCHAR(50)                            NOT NULL,
    likes      INT                                     NOT NULL,
    audience   ENUM ('private', 'public', 'followers') NOT NULL DEFAULT 'public',
    created_at DATETIME                                NOT NULL,
    updated_at DATETIME                                NULL,
    deleted_at DATETIME                                NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX post_id_unique (id ASC) VISIBLE,
    INDEX user_id_fk_idx (user_id ASC) VISIBLE,
    CONSTRAINT user_id_fk
        FOREIGN KEY (user_id)
            REFERENCES users (id)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;