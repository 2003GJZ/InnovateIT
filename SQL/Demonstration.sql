-- 订单数据库
CREATE DATABASE order_db
CHARACTER SET = 'utf8mb4'
COLLATE = 'utf8mb4_general_ci';


-- 订单表
CREATE TABLE order_record (
                              order_id BIGINT AUTO_INCREMENT PRIMARY KEY,    -- 订单ID
                              user_id BIGINT NOT NULL,                       -- 用户ID
                              product_id BIGINT NOT NULL,                    -- 商品ID
                              order_date DATETIME NOT NULL,                  -- 订单日期
                              status TINYINT(1) NOT NULL DEFAULT 0,          -- 订单状态（0：待支付，1：已支付，2：已发货）
                              amount DECIMAL(10, 2) NOT NULL,                -- 订单金额
                              created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间
                              updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP, -- 更新时间
                              INDEX idx_user_order (user_id, order_date)     -- 联合索引：user_id和order_date
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- 商品信息表
CREATE TABLE product (
                         product_id BIGINT AUTO_INCREMENT PRIMARY KEY,  -- 商品ID
                         product_name VARCHAR(255) NOT NULL,            -- 商品名称
                         price DECIMAL(10, 2) NOT NULL,                 -- 商品价格
                         stock INT NOT NULL,                            -- 库存数量
                         created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP, -- 创建时间
                         updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP -- 更新时间
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;
