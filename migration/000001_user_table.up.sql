CREATE TABLE IF NOT EXISTS users
(
    id          INT UNSIGNED NOT NULL AUTO_INCREMENT,
    email       VARCHAR(45)  NOT NULL,
    full_name   VARCHAR(45)  NOT NULL,
    profile_url VARCHAR(255) NULL,
    bio         VARCHAR(500) NULL,
    cover_url   VARCHAR(255) NULL,
    is_creator  TINYINT      NOT NULL,
    create_at   DATETIME     NOT NULL,
    updated_at  DATETIME     NULL,
    deleted_at  DATETIME     NULL,
    PRIMARY KEY (id),
    UNIQUE INDEX id_unique (id ASC) VISIBLE,
    UNIQUE INDEX email_unique (email ASC) VISIBLE
)
    ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4;



