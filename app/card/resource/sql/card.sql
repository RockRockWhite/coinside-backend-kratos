CREATE TABLE
    IF NOT EXISTS `c_card`
(
    `id`         int          NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `team_id`    int          NOT NULL COMMENT '团队id',
    `title`      varchar(255) NOT NULL COMMENT '卡片标题',
    `content`    text         NOT NULL COMMENT '卡片详细内容，以标记语言存储',
    `status`     tinyint      NOT NULL DEFAULT '0' COMMENT '卡片状态, 0：进行中，1：已完成',
    `created_at` timestamp    NOT NULL COMMENT '创建时间',
    `updated_at` timestamp    NULL COMMENT '更新时间',
    `deleted_at` timestamp    NULL COMMENT '删除时间，软删除支持字段',
    INDEX `idx_team_id` (`team_id`),
    INDEX `idx_delete_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '卡片信息表';


CREATE TABLE
    IF NOT EXISTS `c_member`
(
    `id`         int       NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `card_id`    int       NULL COMMENT '所属卡片id',
    `user_id`    int       NOT NULL COMMENT '用户id',
    `is_admin`   tinyint   NOT NULL DEFAULT '0' COMMENT '管理员：0：非管理员 1：管理员',
    `created_at` timestamp NOT NULL COMMENT '创建时间',
    `updated_at` timestamp NULL COMMENT '更新时间',
    `deleted_at` timestamp NULL COMMENT '删除时间,软删除支持字段',
    INDEX `idx_card_id` (`card_id`),
    INDEX `idx_user_id` (`user_id`),
    INDEX `idx_delete_at` (`deleted_at`),
    UNIQUE `idx_card_user` (`card_id`, `user_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '卡片成员表';

CREATE TABLE
    IF NOT EXISTS `c_tag`
(
    `id`         int         NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `content`    varchar(32) NOT NULL UNIQUE COMMENT '标签内容',
    `created_at` timestamp   NULL COMMENT '创建时间',
    `updated_at` timestamp   NULL COMMENT '更新时间',
    `deleted_at` timestamp   NULL COMMENT '删除时间，软删除支持字段',
    INDEX `idx_delete_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '标签信息表';

CREATE TABLE
    IF NOT EXISTS `c_card_tag`
(
    `id`         int       NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `card_id`    int       NULL COMMENT '卡片id',
    `tag_id`     int       NOT NULL COMMENT '标签id',
    INDEX `idx_card_id` (`card_id`),
    INDEX `idx_tag_id` (`tag_id`),
    UNIQUE `idx_card_tag` (`card_id`, `tag_id`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '卡片-标签表';
