
CREATE TABLE
    IF NOT EXISTS `m_markdown` (
                              `id` int NOT NULL PRIMARY KEY  AUTO_INCREMENT COMMENT 'id',
                              `card_id` int NOT NULL COMMENT '卡片id',
                              `content` text  NOT NULL COMMENT '内容',
                              `created_at` timestamp NULL DEFAULT NULL COMMENT '创建时间',
                              `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                              `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除支持字段',
                              INDEX `idx_card_id` (`card_id`)
) ENGINE=INNODB
    DEFAULT CHARSET = utf8mb4 COMMENT 'markdown信息表';


