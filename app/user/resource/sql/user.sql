CREATE TABLE
    IF NOT EXISTS `u_user`
(
    `id`             int          NOT NULL PRIMARY KEY AUTO_INCREMENT COMMENT '用户id',
    `passwd_hash`    varchar(255) NOT NULL COMMENT '密码加盐哈希',
    `passwd_salt`    varchar(255) NOT NULL COMMENT '密码盐值',
    `nickname`       varchar(255) NOT NULL UNIQUE COMMENT '用户昵称',
    `fullname`       varchar(255) COMMENT '用户全名',
    `avatar`         varchar(255) COMMENT '头像链接',
    `email`          varchar(255) UNIQUE COMMENT '邮箱',
    `email_verified` tinyint(1)   NOT NULL DEFAULT 0 COMMENT '邮箱验证状态',
    `mobile`         varchar(50) UNIQUE COMMENT '手机号',
    `config`         varchar(255) COMMENT '用户信息配置文件',
    `logined_at`     timestamp    NULL COMMENT '最后一次登录时间',
    `created_at`     timestamp    NOT NULL COMMENT '创建时间',
    `updated_at`      timestamp    NULL COMMENT '更新时间',
    `deleted_at`     timestamp    NULL COMMENT '删除时间，软删除支持字段',
    INDEX `idx_delete_at` (`deleted_at`)
) ENGINE = InnoDB
  DEFAULT CHARSET = utf8mb4 COMMENT '用户信息表';