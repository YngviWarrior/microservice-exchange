CREATE DATABASE IF NOT EXISTS `exchange`;

CREATE TABLE `coin`(
    `coin` INT(11) auto_increment,
    `symbol` VARCHAR(255) NOT NULL UNIQUE,
    `active` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`coin`)
);

INSERT INTO `coin`(`symbol`, `active`) VALUES("USDT", 1),("BRL", 1),("BTC", 1);

CREATE TABLE `parity`(
    `parity` INT(11) auto_increment,
    `symbol` VARCHAR(255) NOT NULL UNIQUE,
    `active` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`parity`)
);

INSERT INTO `parity`(`symbol`, `active`) VALUES("BTCUSDT", 1),("USDTBRL", 1),("BTCBRL", 1);

CREATE TABLE `exchange` (
    `exchange` INT(11) NOT NULL auto_increment,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY(`exchange`)
);

INSERT INTO `exchange` (`name`) 
VALUES ("BINANCE"), ("BYBIT"), ("TESTER");

CREATE TABLE `wallet` (
    `wallet` BIGINT(11) auto_increment,
    `exchange` BIGINT(11) NOT NULL,
    `coin` INT(11) NOT NULL,
    `amount` DECIMAL(60,8) NOT NULL DEFAULT 0,
    PRIMARY KEY(`wallet`),
    FOREIGN KEY(`exchange`) REFERENCES `exchange`(`exchange`),
    FOREIGN KEY(`coin`) REFERENCES `coin`(`coin`)
);

ALTER TABLE `wallet` ADD COLUMN `exchange` INT(11) NOT NULL AFTER `wallet`;
ALTER TABLE `wallet` ADD CONSTRAINT `fk_exchange_id` FOREIGN KEY (`exchange`) REFERENCES `exchange`(`exchange`);

INSERT INTO `wallet`(`user`, `exchange`, `coin`, `amount`) VALUES (1, 2, 4, 0), (1, 2, 1, 2), (1, 2, 1, 3);

CREATE TABLE `transaction_type` (
    `transaction_type` INT(11) auto_increment,
    `description` VARCHAR(255) NOT NULL,
    `active` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`transaction_type`)
);

INSERT INTO `transaction_type`(`description`, `active`) VALUES("Buy", 1),("Sell", 1),("Deposit", 1),("Withdraw", 1);

CREATE TABLE `transaction` (
    `transaction` BIGINT(11) auto_increment,
    `transaction_type` INT(11) NOT NULL,
    `wallet_sent` BIGINT(11) NOT NULL,
    `wallet_received` BIGINT(11) NOT NULL,
    `amount` DECIMAL(60,8) NOT NULL DEFAULT 0,
    PRIMARY KEY(`transaction`),
    FOREIGN KEY(`transaction_type`) REFERENCES `transaction_type`(`transaction_type`),
    FOREIGN KEY(`wallet_sent`) REFERENCES `wallet`(`wallet`),
    FOREIGN KEY(`wallet_received`) REFERENCES `wallet`(`wallet`)
);

DROP TABLE `candle`;
CREATE TABLE `candle`(
    `parity` INT(11) NOT NULL,
    `exchange` INT(11) NOT NULL,
    `mts` BIGINT(20) NOT NULL,
    `open` DECIMAL(60,8) NOT NULL,
    `close` DECIMAL(60,8) NOT NULL,
    `high` DECIMAL(60,8) NOT NULL,
    `low` DECIMAL(60,8) NOT NULL,
    `volume` DECIMAL(60,8) NOT NULL,
    PRIMARY KEY (`parity`, `exchange`, `mts`)
);

ALTER TABLE `candle` ADD COLUMN `exchange` INT(11) NOT NULL AFTER `parity`;
ALTER TABLE `candle` ADD CONSTRAINT `fk_candle_exchange` FOREIGN KEY (`exchange`) REFERENCES `exchange`(`exchange`);

DROP TABLE IF EXISTS `operation`;
CREATE TABLE `operation`(
    `operation` BIGINT(20) auto_increment,
    `user` BIGINT(11) NOT NULL,
    `parity` INT(11) NOT NULL,
    `exchange` INT(11) NOT NULL,
    `strategy` INT(11) NOT NULL,
    `strategy_variant` INT(11) NOT NULL,
    `mts_start` BIGINT(20) NOT NULL,
    `mts_finish` BIGINT(20) NOT NULL DEFAULT 0,
    `open_price` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `close_price` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `invested_amount` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `profit_amount` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `profit` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `closed` TINYINT(1) NOT NULL DEFAULT 0,
    `audit` TINYINT(1) NOT NULL DEFAULT 0,
    `enabled` TINYINT(1) NOT NULL DEFAULT 0,
    `times_canceled` INT(11) NOT NULL DEFAULT 0,
    PRIMARY KEY(`operation`),
    FOREIGN KEY (`parity`) REFERENCES `parity`(`parity`),
    FOREIGN KEY (`exchange`) REFERENCES `exchange`(`exchange`)
);


ALTER TABLE `operation` ADD COLUMN `times_canceled` INT(11) NOT NULL DEFAULT 0;
ALTER TABLE `operation` ADD COLUMN `audit` TINYINT(1) NOT NULL DEFAULT 0 after `closed`;
ALTER TABLE `operation` ADD COLUMN `enabled` TINYINT(1) NOT NULL DEFAULT 0 after `audit`;
ALTER TABLE `operation` ADD COLUMN `invested_amount` DECIMAL(60,8) NOT NULL DEFAULT 0 AFTER `close_price`;
ALTER TABLE `operation` ADD COLUMN `profit_amount` DECIMAL(60,8) NOT NULL DEFAULT 0 AFTER `invested_amount`;
ALTER TABLE `operation` ADD COLUMN `exchange` INT(11) NOT NULL AFTER `parity`;
ALTER TABLE `operation` ADD CONSTRAINT `fk_exchange_operation` FOREIGN KEY (`exchange`) REFERENCES `exchange`(`exchange`);
ALTER TABLE `operation` ADD COLUMN `strategy` INT(11) NOT NULL DEFAULT 2 AFTER `exchange`;
ALTER TABLE `operation` ADD CONSTRAINT `fk_strategy_operation` FOREIGN KEY (`strategy`) REFERENCES `strategy`(`strategy`);
ALTER TABLE `operation` ADD COLUMN `strategy_variant` INT(11) NOT NULL DEFAULT 1 AFTER `strategy`;
ALTER TABLE `operation` ADD CONSTRAINT `fk_strategy_variant_operation` FOREIGN KEY (`strategy_variant`) REFERENCES `strategy_variant`(`strategy_variant`);

DROP TABLE IF EXISTS `operation_exchange_status`;
CREATE TABLE `operation_exchange_status`(
    `operation_exchange_status` INT(11) NOT NULL auto_increment,
    `exchange` INT(11) NOT NULL,
    `description` VARCHAR(200) NOT NULL,
    PRIMARY KEY (`operation_exchange_status`),
    FOREIGN KEY (`exchange`) REFERENCES `exchange`(`exchange`)
);

INSERT INTO `operation_exchange_status`(`description`, `exchange`) VALUES
('NEW', 2), ('PARTIALLY_FILLED', 2), ('FILLED', 2), ('CANCELED', 2), ('PENDING_CANCEL', 2), ('PENDING_NEW', 2), ('REJECTED', 2), 
('ORDER_NEW', 2), ('ORDER_CANCELED', 2), ('ORDER_FILLED', 2), ('ORDER_FAILED', 2);

DROP TABLE IF EXISTS `operation_history`;
CREATE TABLE `operation_history`(
    `operation_history` BIGINT(20) auto_increment,
    `operation` BIGINT(20) NOT NULL,
    `transaction_type` INT(11) NOT NULL,
    `coin_price` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `coin_quantity` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `stable_price` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `stable_quantity` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `fee` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `operation_exchange_id` VARCHAR(200) NOT NULL,
    `operation_exchange_status` INT(11) NOT NULL DEFAULT 0,
    PRIMARY KEY(`operation_history`),
    FOREIGN KEY (`operation`) REFERENCES `operation`(`operation`),
    FOREIGN KEY (`transaction_type`) REFERENCES `transaction_type`(`transaction_type`)
);

ALTER TABLE `operation_history` ADD COLUMN `order_id` VARCHAR(200) NOT NULL AFTER `fee`;
ALTER TABLE `operation_history` RENAME COLUMN `order_id` TO `operation_exchange_id`;
ALTER TABLE `operation_history` ADD COLUMN `operation_exchange_status` INT(11) NOT NULL DEFAULT 1;
ALTER TABLE `operation_history` ADD CONSTRAINT `fk_operation_exchange_status_id` FOREIGN KEY (`operation_exchange_status`) REFERENCES `operation_exchange_status`(`operation_exchange_status`);

-- DROP TABLE IF EXISTS `config`;
-- CREATE TABLE `config` (
--     `max_trade_transactions` INT(11) NOT NULL DEFAULT 0,
--     `max_wallet_limit_value` DECIMAL(60,8) NOT NULL DEFAULT 0,
--     `min_profit_to_sell` DECIMAL(60,8) NOT NULL DEFAULT 0,
--     `buy_value` DECIMAL(60,8) NOT NULL DEFAULT 0
-- );

-- INSERT INTO `config`(`max_trade_transactions`, `max_wallet_limit_value`, `min_profit_to_sell`, `buy_value`) VALUES (4, 65, 200, 15);

----------------------------------------------------------------
DROP TABLE IF EXISTS `modality`;
CREATE TABLE `modality` (
    `modality` INT(11) NOT NULL auto_increment,
    `name` VARCHAR(200) NOT NULL,
    `enable` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`modality`)a
);

INSERT INTO `modality`(`name`) VALUES ('Spot');

DROP TABLE IF EXISTS `strategy`;
CREATE TABLE `strategy` (
    `strategy` INT(11) NOT NULL auto_increment,
    `modality` INT(11) NOT NULL,
    `name` VARCHAR(200) NOT NULL,
    `enable` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`strategy`),
    FOREIGN KEY (`modality`) REFERENCES `modality`(`modality`)
);

INSERT INTO `strategy` (`modality`, `name`) VALUES (1, 'Average Price'), (1, 'Close/Open');
INSERT INTO `strategy` (`modality`, `name`) VALUES (1, 'Fast Trade');

DROP TABLE IF EXISTS `strategy_variant`;
CREATE TABLE `strategy_variant` (
    `strategy_variant` INT(11) NOT NULL auto_increment,
    `strategy` INT(11) NOT NULL,
    `name` VARCHAR(200) NOT NULL,
    `enable` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`strategy_variant`),
    FOREIGN KEY (`strategy`) REFERENCES `strategy`(`strategy`)
);

INSERT INTO `strategy_variant` (`strategy`, `name`) VALUES 
(1, 'Default Price'),
(1, 'Day Average Price'),
(1, 'Week Average Price'),
(1, 'Month Average Price'),
(1, 'SMA 200');

DROP TABLE IF EXISTS `average_price`;
CREATE TABLE `average_price` (
    `parity` INT(11) NOT NULL,
    `exchange` INT(11) NOT NULL,
    `day` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `day_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0,
    `week` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `week_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0,
    `month` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `month_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0,
    `sma` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `sma_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (`parity`, `exchange`),
    FOREIGN KEY (`parity`) REFERENCES `parity`(`parity`),
    FOREIGN KEY (`exchange`) REFERENCES `exchange`(`exchange`)
);

ALTER TABLE `average_price` ADD COLUMN `sma` DECIMAL(60,8) NOT NULL DEFAULT 0;
ALTER TABLE `average_price` ADD COLUMN `sma_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0;
ALTER TABLE `average_price` ADD COLUMN `day_roc` DECIMAL(60,8) NOT NULL DEFAULT 0 AFTER `day`;
ALTER TABLE `average_price` ADD COLUMN `week_roc` DECIMAL(60,8) NOT NULL DEFAULT 0 AFTER `week`;
ALTER TABLE `average_price` ADD COLUMN `month_roc` DECIMAL(60,8) NOT NULL DEFAULT 0 AFTER `month`;
ALTER TABLE `average_price` ADD COLUMN `sma_roc` DECIMAL(60,8) NOT NULL DEFAULT 0 AFTER `sma`;
ALTER TABLE `average_price` DROP COLUMN `roc`;

INSERT INTO `average_price`(`parity`)
VALUES (1),(4);

DROP TABLE IF EXISTS `operation_meta_data_fast_trade`;
CREATE TABLE `operation_meta_data_fast_trade` (
    `operation_meta_data_fast_trade` BIGINT(20) NOT NULL auto_increment,
    `operation` BIGINT(20) NOT NULL,
    `minumum_price` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `maximum_price` DECIMAL(60,8) NOT NULL DEFAULT 0,
    `is_receding` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`operation_meta_data_fast_trade`),
    FOREIGN KEY (`operation`) REFERENCES `operation`(`operation`)
);

ALTER TABLE `operation_meta_data_fast_trade` ADD COLUMN `last_price` DECIMAL(60,8) NOT NULL DEFAULT 0 AFTER `maximum_price`;
ALTER TABLE `operation_meta_data_fast_trade` RENAME COLUMN `minumum_price` TO `minimum_price`;
