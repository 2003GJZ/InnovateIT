-- 创建用户管理库，选择utf8编码
CREATE DATABASE user_management_db
CHARACTER SET = 'utf8'
COLLATE = 'utf8_general_ci';

-- 隔离级别
SET GLOBAL TRANSACTION ISOLATION LEVEL REPEATABLE READ;



CREATE TABLE user_info (
                           id BIGINT AUTO_INCREMENT PRIMARY KEY,         -- 主键，自增
                           username VARCHAR(255) NOT NULL,               -- 用户名
                           wechat_id VARCHAR(255),                       -- 绑定微信号
                           email VARCHAR(255),                           -- 绑定邮箱
                           phone VARCHAR(20) NOT NULL,                   -- 绑定电话号码，不能为空
                           created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间
                           updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间
                           UNIQUE KEY uq_email (email),                  -- 唯一约束：绑定邮箱
                           UNIQUE KEY uq_phone (phone),                  -- 唯一约束：绑定电话
                           UNIQUE KEY uq_wechat (wechat_id)              -- 唯一约束：绑定微信号
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

CREATE TABLE user_login (
                            phone VARCHAR(20) PRIMARY KEY,                -- 主键：绑定电话（唯一）
                            username VARCHAR(255) NOT NULL,               -- 用户名
                            password VARCHAR(255) NOT NULL,               -- 密码
                            user_info_id BIGINT,                          -- 个人信息表中的外键
                            login_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 登录时间
                            FOREIGN KEY (user_info_id) REFERENCES user_info(id) -- 外键关联个人信息表
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;
