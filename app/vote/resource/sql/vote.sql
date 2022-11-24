
CREATE TABLE
    IF NOT EXISTS `m_vote` (
                          `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
                          `card_id` int NOT NULL COMMENT '卡片id',
                          `title` varchar(255)  NOT NULL COMMENT '投票标题',
                          `created_at` timestamp NOT NULL COMMENT '创建时间',
                          `updated_at` timestamp NULL COMMENT '更新时间',
                          `deleted_at` timestamp NULL COMMENT '删除时间，软删除支持字段',
                          INDEX `m_vote_card_id_IDX` (`card_id`)
)ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4 COMMENT '投票信息表';

-- ljxsteam.m_vote_item definition

CREATE TABLE
    IF NOT EXISTS `m_vote_item` (
                               `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
                               `vote_id` int NOT NULL COMMENT '投票id',
                               `content` varchar(255)  NOT NULL COMMENT '投票项',
                               `created_at` timestamp NOT NULL COMMENT '创建时间',
                               `updated_at` timestamp NULL  COMMENT '更新时间',
                               `deleted_at` timestamp NULL  COMMENT '删除时间，软删除支持字段',
                                INDEX `m_vote_item_vote_id_IDX` (`vote_id`)
) ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4 COMMENT '投票项信息表';

-- ljxsteam.m_vote_commit definition

CREATE TABLE
    IF NOT EXISTS `m_vote_commit` (
                                 `id` int NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
                                 `vote_item_id` int NOT NULL COMMENT '投票项id',
                                 `user_id` int NOT NULL COMMENT '用户id',
                                 `created_at` timestamp NOT NULL COMMENT '创建时间',
                                 `updated_at` timestamp NULL  COMMENT '更新时间',
                                 `deleted_at` timestamp NULL COMMENT '删除时间,软删除支持字段',
                                    INDEX `m_vote_commit_vote_item_id_IDX` (`vote_item_id`),
                                    INDEX `m_vote_commit_user_id_IDX` (`user_id`)
)  ENGINE = InnoDB
    DEFAULT CHARSET = utf8mb4 COMMENT '投票结果信息表';
