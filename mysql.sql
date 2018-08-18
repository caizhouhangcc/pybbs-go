-- 创建数据库
CREATE DATABASE `pybbs-go` default CHARACTER SET utf8mb4 collate utf8mb4_unicode_ci;

-- User增加昵称
ALTER TABLE `user` ADD COLUMN `nickname` varchar(255) NOT NULL  DEFAULT ''