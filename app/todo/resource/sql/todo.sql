CREATE TABLE
    IF NOT EXISTS`m_todo`
(
    `id`          INT          NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `card_id`     INT          NOT NULL COMMENT '卡片id',
    `title`       VARCHAR(255) NOT NULL COMMENT 'todolist标题',
    `created_at`  TIMESTAMP    NOT NULL COMMENT '创建时间',
    `updated_at`  TIMESTAMP    NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at`  TIMESTAMP    NULL DEFAULT NULL COMMENT '删除时间，软删除支持字段',
    INDEX `idx_card_id` (`card_id`)
    ) ENGINE=INNODB
    DEFAULT CHARSET = utf8mb4 COMMENT '待办组件信息表';



CREATE TABLE
    IF NOT EXISTS`m_todo_item`
(
    `id`               INT          NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT 'id',
    `todo_id`          INT          COMMENT 'todolist id',
    `content`          VARCHAR(255) NOT NULL COMMENT '待办项',
    `is_finished`      TINYINT      NOT NULL DEFAULT '0' COMMENT '是否完成，0未完成，1已完成',
    `finished_user_id` INT          DEFAULT NULL COMMENT '完成用户id',
    `created_at`       TIMESTAMP    NOT NULL COMMENT '创建时间',
    `updated_at`       TIMESTAMP    NULL DEFAULT NULL COMMENT '更新时间',
    `deleted_at`       TIMESTAMP    NULL DEFAULT NULL COMMENT '删除时间，软删除支持字段',
    INDEX `idx_todo_id` (`todo_id`),
    INDEX `idx_finished_user_id` (`finished_user_id`)
    ) ENGINE=INNODB
    DEFAULT CHARSET = utf8mb4 COMMENT '待办项信息表';

