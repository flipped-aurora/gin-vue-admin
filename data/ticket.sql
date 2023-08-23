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

 Date: 23/08/2023 07:56:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for order_info
-- ----------------------------
DROP TABLE IF EXISTS `order_info`;
CREATE TABLE `order_info` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `order_id` varchar(32) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '订单ID',
  `ticket_number` varchar(12) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '订单取票号',
  `contact_phone` varchar(11) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '订单联系电话',
  `is_transfer` tinyint(1) DEFAULT NULL COMMENT '是否换乘.0:false',
  `is_occupy_seat` bigint(20) DEFAULT NULL COMMENT '是否占座.0:false',
  `complete_status` bigint(20) DEFAULT NULL COMMENT '完成状态.0:未完成',
  `fail_reason` varchar(255) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '出票失败原因',
  `machine_name` varchar(128) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '设备名称',
  `deleted_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_order_info_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT='订单统计表';

-- ----------------------------
-- Records of order_info
-- ----------------------------
BEGIN;
INSERT INTO `order_info` VALUES (1, 'p-order-id-01', 'p-number-01', '13510480000', 0, 0, 1, '', 'p-mac', NULL, '2023-08-23 03:11:22.000', '2023-08-23 03:10:00.000');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
