CREATE TABLE IF NOT EXISTS `post_contents` (
  `content_id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `conent_url` VARCHAR(500) NOT NULL,
  `post_id` INT UNSIGNED NOT NULL,
  `created_at` DATETIME NULL,
  `updated_at` DATETIME NULL,
  `deleted_at` DATETIME NULL,
  PRIMARY KEY (`content_id`),
  INDEX `post_id_fk_idx` (`post_id` ASC) VISIBLE,
  CONSTRAINT `post_id_fk`
    FOREIGN KEY (`post_id`)
    REFERENCES `mydb`.`posts` (`post_id`)
    ON DELETE NO ACTION
    ON UPDATE NO ACTION)
ENGINE = InnoDB DEFAULT CHARSET = utf8mb4;


 