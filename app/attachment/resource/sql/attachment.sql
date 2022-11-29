
CREATE TABLE
    IF NOT EXISTS`m_attachment` (
                                `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
                                `card_id` int NOT NULL COMMENT '卡片id',
                                `link` varchar(255)  DEFAULT NULL COMMENT '资源链接',
                                `download_count` int NOT NULL DEFAULT '0' COMMENT '下载量',
                                `created_at` timestamp NOT NULL COMMENT '创建时间',
                                `updated_at` timestamp NULL DEFAULT NULL COMMENT '更新时间',
                                `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间，软删除支持字段',
                                INDEX `m_attachment_card_id_IDX` (`card_id`)
) ENGINE=INNODB
    DEFAULT CHARSET = utf8mb4 COMMENT 'attachment';


