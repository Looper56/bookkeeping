-- Adminer 4.8.1 MySQL 5.7.26 dump
SET NAMES utf8;
SET
time_zone = '+08:00';
SET
foreign_key_checks = 0;
SET
sql_mode = 'NO_AUTO_VALUE_ON_ZERO';

SET NAMES utf8mb4;

DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`
(
    `id`                  int unsigned NOT NULL AUTO_INCREMENT COMMENT 'id',
    `uid`                 varchar(64) NOT NULL COMMENT 'uid',
    `official_account_id` varchar(128)                      DEFAULT NULL COMMENT 'oa openid',
    `mp_open_id`          varchar(128)                      DEFAULT NULL COMMENT 'mp openid',
    `union_id`            varchar(128)                      DEFAULT NULL COMMENT 'open union id',
    `channel`             tinyint unsigned NOT NULL DEFAULT '1' COMMENT 'channel:wechat',
    `nick_name`           varchar(64) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT 'nick name',
    `avatar_url`          varchar(255)                      DEFAULT NULL COMMENT 'profile url',
    `gender`              tinyint unsigned NOT NULL DEFAULT '0' COMMENT '0 is unknown 1 is man 2 is female',
    `mobile`              varchar(64)                       DEFAULT NULL COMMENT 'mobile number',
    `province`            varchar(64) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT 'province',
    `country`             varchar(64) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT 'country',
    `city`                varchar(64) CHARACTER SET utf8mb4 DEFAULT NULL COMMENT 'city',
    `is_white`            tinyint unsigned NOT NULL DEFAULT '0' COMMENT 'is white',
    `is_cancel`           tinyint unsigned NOT NULL DEFAULT '0' COMMENT 'is cancel',
    `is_sub_oa`           tinyint unsigned NOT NULL DEFAULT '0' COMMENT 'is subscribe official account',
    `last_active_time`    datetime                          DEFAULT NULL COMMENT 'last active time',
    `sub_oa_time`         datetime                          DEFAULT NULL COMMENT 'subscribe official account time',
    `unfollow_oa_time`    datetime                          DEFAULT NULL COMMENT 'unfollow official account time',
    `created_at`          datetime    NOT NULL              DEFAULT CURRENT_TIMESTAMP COMMENT 'created time',
    `updated_at`          datetime    NOT NULL              DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT 'updated time',
    PRIMARY KEY (`id`),
    UNIQUE KEY `rid` (`uid`),
    KEY                   `idx_official_account_id` (`official_account_id`),
    KEY                   `idx_mp_open_id` (`mp_open_id`),
    KEY                   `idx_union_id` (`union_id`),
    KEY                   `idx_mobile` (`mobile`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='user';
