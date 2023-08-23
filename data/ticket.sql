/*
 Navicat Premium Data Transfer

 Source Server         : mac-leeprince
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : localhost:3306
 Source Schema         : ticket

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 24/08/2023 01:21:31
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order_info
-- ----------------------------
DROP TABLE IF EXISTS `order_info`;
CREATE TABLE `order_info` (
  `id` bigint(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_id` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '订单ID',
  `ticket_number` varchar(12) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '订单取票号',
  `contact_phone` varchar(11) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '订单联系电话',
  `is_transfer` bigint(20) DEFAULT NULL COMMENT '是否换乘.0:false',
  `is_occupy_seat` bigint(20) DEFAULT NULL COMMENT '是否占座.0:false',
  `complete_status` bigint(20) DEFAULT NULL COMMENT '完成状态.0:未完成',
  `fail_reason` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '出票失败原因',
  `machine_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备名称',
  `deleted_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order_info_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单统计表';

-- ----------------------------
-- Records of order_info
-- ----------------------------
BEGIN;
INSERT INTO `order_info` VALUES (1, 'order-id-123', 'no-aqaq', '13510480000', 0, 0, 1, NULL, 'p-mac-10', NULL, '2023-08-23 23:08:42.000', '2023-08-21 23:39:57.000');
INSERT INTO `order_info` VALUES (2, 'order-id-234', 'no-nlnl', '13510580000', 0, 0, 1, NULL, 'p-mac-11', NULL, '2023-08-23 23:08:42.000', '2023-08-22 23:40:05.000');
INSERT INTO `order_info` VALUES (3, 'order-id-asd', 'no-oioi', '13510680000', 0, 0, 1, NULL, 'p-mac-21', NULL, '2023-08-23 23:08:42.000', '2023-08-23 23:08:45.000');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
