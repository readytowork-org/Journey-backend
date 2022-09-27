CREATE TABLE IF NOT EXISTS post_contents
(
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    content_url VARCHAR(500) NOT NULL,
    content_type VARCHAR(500) ,
    thumbnail VARCHAR(500) ,
    post_id     INT UNSIGNED NOT NULL,
    PRIMARY KEY (id),
    INDEX post_id_fk_idx (post_id ASC) VISIBLE,
    CONSTRAINT contents_post_id_fk
        FOREIGN KEY (post_id)
            REFERENCES posts (id)
            ON DELETE NO ACTION
            ON UPDATE NO ACTION
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;