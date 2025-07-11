DROP DATABASE IF EXISTS `exchange`;
CREATE DATABASE IF NOT EXISTS `exchange`;
USE `exchange`;

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

INSERT INTO `parity`(`symbol`, `active`) VALUES("BTCUSDT", 1),("ETHUSDT", 1);

CREATE TABLE `exchange` (
    `exchange` INT(11) NOT NULL auto_increment,
    `name` VARCHAR(255) NOT NULL,
    PRIMARY KEY(`exchange`)
);

INSERT INTO `exchange` (`name`) 
VALUES ("BINANCE"), ("BYBIT"), ("TESTER");

CREATE TABLE `wallet` (
    `wallet` BIGINT(11) auto_increment,
    `exchange` INT(11) NOT NULL,
    `coin` INT(11) NOT NULL,
    `amount` VARCHAR(32) NOT NULL DEFAULT "0",
    PRIMARY KEY(`wallet`),
    FOREIGN KEY(`exchange`) REFERENCES `exchange`(`exchange`),
    FOREIGN KEY(`coin`) REFERENCES `coin`(`coin`)
);

-- INSERT INTO `wallet`(`exchange`, `coin`, `amount`) VALUES (2, 4, 0), (2, 1, 2), (2, 1, 3);

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
    `amount` VARCHAR(32) NOT NULL DEFAULT "0",
    PRIMARY KEY(`transaction`),
    FOREIGN KEY(`transaction_type`) REFERENCES `transaction_type`(`transaction_type`),
    FOREIGN KEY(`wallet_sent`) REFERENCES `wallet`(`wallet`),
    FOREIGN KEY(`wallet_received`) REFERENCES `wallet`(`wallet`)
);

CREATE TABLE `candle`(
    `parity` INT(11) NOT NULL,
    `exchange` INT(11) NOT NULL,
    `mts` BIGINT(20) NOT NULL,
    `open` VARCHAR(32) NOT NULL,
    `close` VARCHAR(32) NOT NULL,
    `high` VARCHAR(32) NOT NULL,
    `low` VARCHAR(32) NOT NULL,
    `volume` VARCHAR(32) NOT NULL,
    PRIMARY KEY (`parity`, `exchange`, `mts`),
    FOREIGN KEY(`exchange`) REFERENCES `exchange`(`exchange`)
);

-- ALTER TABLE `candle` ADD COLUMN `exchange` INT(11) NOT NULL AFTER `parity`;
-- ALTER TABLE `candle` ADD CONSTRAINT `fk_candle_exchange` FOREIGN KEY (`exchange`) REFERENCES `exchange`(`exchange`);
DROP TABLE IF EXISTS `modality`;
CREATE TABLE `modality` (
    `modality` INT(11) NOT NULL auto_increment,
    `name` VARCHAR(200) NOT NULL,
    `enabled` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`modality`)
);

INSERT INTO `modality`(`name`) VALUES ('Spot');

DROP TABLE IF EXISTS `strategy`;
CREATE TABLE `strategy` (
    `strategy` INT(11) NOT NULL auto_increment,
    `modality` INT(11) NOT NULL,
    `name` VARCHAR(200) NOT NULL,
    `enabled` TINYINT(1) NOT NULL DEFAULT 0,
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
    `enabled` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`strategy_variant`),
    FOREIGN KEY (`strategy`) REFERENCES `strategy`(`strategy`)
);

INSERT INTO `strategy_variant` (`strategy`, `name`) VALUES 
(1, 'Default Price'),
(1, 'Day Average Price'),
(1, 'Week Average Price'),
(1, 'Month Average Price'),
(1, 'SMA 200');


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
    `open_price` VARCHAR(32) NOT NULL DEFAULT 0,
    `close_price` VARCHAR(32) NOT NULL DEFAULT 0,
    `invested_amount` VARCHAR(32) NOT NULL DEFAULT 0,
    `profit_amount` VARCHAR(32) NOT NULL DEFAULT 0,
    `profit` VARCHAR(32) NOT NULL DEFAULT 0,
    `closed` TINYINT(1) NOT NULL DEFAULT 0,
    `audit` TINYINT(1) NOT NULL DEFAULT 0,
    `enabled` TINYINT(1) NOT NULL DEFAULT 0,
    `in_transaction` TINYINT(1) NOT NULL DEFAULT 0,
    `times_canceled` INT(11) NOT NULL DEFAULT 0,
    PRIMARY KEY(`operation`),
    FOREIGN KEY (`parity`) REFERENCES `parity`(`parity`),
    FOREIGN KEY (`exchange`) REFERENCES `exchange`(`exchange`),
    FOREIGN KEY(`strategy`) REFERENCES `strategy`(`strategy`),
    FOREIGN KEY(`strategy_variant`) REFERENCES `strategy_variant`(`strategy_variant`)
);

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
    `coin_price` VARCHAR(32) NOT NULL DEFAULT 0,
    `coin_quantity` VARCHAR(32) NOT NULL DEFAULT 0,
    `stable_price` VARCHAR(32) NOT NULL DEFAULT 0,
    `stable_quantity` VARCHAR(32) NOT NULL DEFAULT 0,
    `fee` VARCHAR(32) NOT NULL DEFAULT 0,
    `operation_exchange_id` VARCHAR(200) NOT NULL,
    `operation_exchange_status` INT(11) NOT NULL DEFAULT 1,
    PRIMARY KEY(`operation_history`),
    FOREIGN KEY (`operation`) REFERENCES `operation`(`operation`),
    FOREIGN KEY (`transaction_type`) REFERENCES `transaction_type`(`transaction_type`),
    FOREIGN KEY (`operation_exchange_status`) REFERENCES `operation_exchange_status`(`operation_exchange_status`)
);

-- ALTER TABLE `operation_history` ADD COLUMN `order_id` VARCHAR(200) NOT NULL AFTER `fee`;
-- ALTER TABLE `operation_history` RENAME COLUMN `order_id` TO `operation_exchange_id`;
-- ALTER TABLE `operation_history` ADD COLUMN `operation_exchange_status` INT(11) NOT NULL DEFAULT 1;
-- ALTER TABLE `operation_history` ADD CONSTRAINT `fk_operation_exchange_status_id` FOREIGN KEY (`operation_exchange_status`) REFERENCES `operation_exchange_status`(`operation_exchange_status`);

-- --------------------------------------------------------------

DROP TABLE IF EXISTS `average_price`;
CREATE TABLE `average_price` (
    `parity` INT(11) NOT NULL,
    `exchange` INT(11) NOT NULL,
    `day` VARCHAR(32) NOT NULL DEFAULT 0,
    `day_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0,
    `week` VARCHAR(32) NOT NULL DEFAULT 0,
    `week_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0,
    `month` VARCHAR(32) NOT NULL DEFAULT 0,
    `month_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0,
    `sma` VARCHAR(32) NOT NULL DEFAULT 0,
    `sma_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0,
    PRIMARY KEY (`parity`, `exchange`),
    FOREIGN KEY (`parity`) REFERENCES `parity`(`parity`),
    FOREIGN KEY (`exchange`) REFERENCES `exchange`(`exchange`)
);

-- ALTER TABLE `average_price` ADD COLUMN `sma` VARCHAR(32) NOT NULL DEFAULT 0;
-- ALTER TABLE `average_price` ADD COLUMN `sma_update_timestamp` BIGINT(20) NOT NULL DEFAULT 0;
ALTER TABLE `average_price` ADD COLUMN `day_roc` VARCHAR(32) NOT NULL DEFAULT 0 AFTER `day`;
ALTER TABLE `average_price` ADD COLUMN `week_roc` VARCHAR(32) NOT NULL DEFAULT 0 AFTER `week`;
ALTER TABLE `average_price` ADD COLUMN `month_roc` VARCHAR(32) NOT NULL DEFAULT 0 AFTER `month`;
ALTER TABLE `average_price` ADD COLUMN `sma_roc` VARCHAR(32) NOT NULL DEFAULT 0 AFTER `sma`;
-- ALTER TABLE `average_price` DROP COLUMN `roc`;

INSERT INTO `average_price`(`parity`, `exchange`)
VALUES (1,2),(3,2);

DROP TABLE IF EXISTS `operation_meta_data_fast_trade`;
CREATE TABLE `operation_meta_data_fast_trade` (
    `operation_meta_data_fast_trade` BIGINT(20) NOT NULL auto_increment,
    `operation` BIGINT(20) NOT NULL,
    `minumum_price` VARCHAR(32) NOT NULL DEFAULT 0,
    `maximum_price` VARCHAR(32) NOT NULL DEFAULT 0,
    `is_receding` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY (`operation_meta_data_fast_trade`),
    FOREIGN KEY (`operation`) REFERENCES `operation`(`operation`)
);

ALTER TABLE `operation_meta_data_fast_trade` ADD COLUMN `last_price` VARCHAR(32) NOT NULL DEFAULT 0 AFTER `maximum_price`;
ALTER TABLE `operation_meta_data_fast_trade` RENAME COLUMN `minumum_price` TO `minimum_price`;

DROP TABLE IF EXISTS `trade_config`;
CREATE TABLE `trade_config` (
    `trade_config` BIGINT(20) NOT NULL UNIQUE auto_increment,
    `user` BIGINT(20) NOT NULL,
    `strategy` INT(11) NOT NULL,
    `strategy_variant` INT(11) NOT NULL,
    `modality` INT(11) NOT NULL,
    `parity` INT(11) NOT NULL,
    `exchange` INT(11) NOT NULL,
    `operation_quantity` INT(11) NOT NULL DEFAULT 0,
    `operation_amount` VARCHAR(32) NOT NULL DEFAULT 0,
    `default_profit_percentage` VARCHAR(32) NOT NULL DEFAULT 0,
    `wallet_value_limit` VARCHAR(32) NOT NULL DEFAULT 0,
    `enabled` TINYINT(1) NOT NULL DEFAULT 0,
    PRIMARY KEY(`user`, `modality`, `strategy`, `strategy_variant`, `parity`, `exchange`)
);
