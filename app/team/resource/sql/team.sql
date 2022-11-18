CREATE TABLE
    IF NOT EXISTS `t_member`
(
    `id`         INT         NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `team_id`    INT       NULL  COMMENT '所属团队id',
    `user_id`    INT         NOT NULL COMMENT '用户id',
    `is_admin`   TINYINT     NOT NULL DEFAULT '0' COMMENT '是否为管理员，1是管理员',
    `created_at` TIMESTAMP   NULL COMMENT '创建时间',
    `updated_at` TIMESTAMP   NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at` TIMESTAMP   NULL DEFAULT NULL COMMENT '删除时间，软删除支持字段',
    INDEX `idx_team_id` (`team_id`) ,
    INDEX `idx_user_id` (`user_id`),
    UNIQUE `idx_team_user` (`team_id`, `user_id`)
) ENGINE=INNODB
DEFAULT CHARSET = utf8mb4 COMMENT '团队成员信息表';

CREATE TABLE
    IF NOT EXISTS `t_team`
(
    `id`           int           NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '团队id',
    `name`         varchar(255)  NOT NULL COMMENT '团队名',
    `description`  varchar(255)  COMMENT '团队描述',
    `website`      varchar(255)  COMMENT '团队官网',
    `avatar`       varchar(255)  COMMENT '头像链接',
    `email`        varchar(255)  COMMENT '邮箱',
    `created_at`   timestamp     NOT NULL COMMENT '创建时间',
    `updated_at`   timestamp     NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at`   timestamp     NULL DEFAULT NULL COMMENT '删除时间，软删除支持字段',
    INDEX `idx_name` (`name`)
)  ENGINE=INNODB
   DEFAULT CHARSET = utf8mb4 COMMENT '团队信息表';
