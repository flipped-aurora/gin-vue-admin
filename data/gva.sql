/*
 Navicat Premium Data Transfer

 Source Server         : mac-leeprince
 Source Server Type    : MySQL
 Source Server Version : 80013
 Source Host           : localhost:3306
 Source Schema         : gva

 Target Server Type    : MySQL
 Target Server Version : 80013
 File Encoding         : 65001

 Date: 24/08/2023 01:21:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v0` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v1` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v2` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v3` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v4` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `v5` varchar(100) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=404 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
BEGIN;
INSERT INTO `casbin_rule` VALUES (280, 'p', '888', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (281, 'p', '888', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (286, 'p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (284, 'p', '888', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (285, 'p', '888', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (283, 'p', '888', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (282, 'p', '888', '/api/updateApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (287, 'p', '888', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (288, 'p', '888', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (289, 'p', '888', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (291, 'p', '888', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (292, 'p', '888', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (290, 'p', '888', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (356, 'p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (355, 'p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (354, 'p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (328, 'p', '888', '/autoCode/createPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (325, 'p', '888', '/autoCode/createPlug', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (322, 'p', '888', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (330, 'p', '888', '/autoCode/delPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (334, 'p', '888', '/autoCode/delSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (324, 'p', '888', '/autoCode/getColumn', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (320, 'p', '888', '/autoCode/getDB', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (331, 'p', '888', '/autoCode/getMeta', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (329, 'p', '888', '/autoCode/getPackage', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (333, 'p', '888', '/autoCode/getSysHistory', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (321, 'p', '888', '/autoCode/getTables', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (326, 'p', '888', '/autoCode/installPlugin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (323, 'p', '888', '/autoCode/preview', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (327, 'p', '888', '/autoCode/pubPlug', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (332, 'p', '888', '/autoCode/rollback', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (294, 'p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (293, 'p', '888', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (358, 'p', '888', '/chatGpt/createSK', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (360, 'p', '888', '/chatGpt/deleteSK', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (359, 'p', '888', '/chatGpt/getSK', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (357, 'p', '888', '/chatGpt/getTable', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (317, 'p', '888', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (318, 'p', '888', '/customer/customer', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (316, 'p', '888', '/customer/customer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (315, 'p', '888', '/customer/customer', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (319, 'p', '888', '/customer/customerList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (353, 'p', '888', '/email/emailTest', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (305, 'p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (306, 'p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (309, 'p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (310, 'p', '888', '/fileUploadAndDownload/editFileName', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (304, 'p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (311, 'p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (307, 'p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (308, 'p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (269, 'p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (295, 'p', '888', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (303, 'p', '888', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (297, 'p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (299, 'p', '888', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (301, 'p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (296, 'p', '888', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (302, 'p', '888', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (300, 'p', '888', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (298, 'p', '888', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (361, 'p', '888', '/orderInfo/findOrderInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (362, 'p', '888', '/orderInfo/getOrderInfoList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (351, 'p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (352, 'p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (350, 'p', '888', '/simpleUploader/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (340, 'p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (341, 'p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (343, 'p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (344, 'p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (342, 'p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (336, 'p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (337, 'p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (338, 'p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (339, 'p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (335, 'p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (345, 'p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (348, 'p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (349, 'p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (346, 'p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (347, 'p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (312, 'p', '888', '/system/getServerInfo', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (313, 'p', '888', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (314, 'p', '888', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (271, 'p', '888', '/user/admin_register', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (277, 'p', '888', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (270, 'p', '888', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES (275, 'p', '888', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (272, 'p', '888', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (279, 'p', '888', '/user/resetPassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (274, 'p', '888', '/user/setSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (276, 'p', '888', '/user/setUserAuthorities', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (278, 'p', '888', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (273, 'p', '888', '/user/setUserInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (395, 'p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (400, 'p', '9528', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (402, 'p', '9528', '/orderInfo/findOrderInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (403, 'p', '9528', '/orderInfo/getOrderInfoList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (401, 'p', '9528', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (398, 'p', '9528', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES (397, 'p', '9528', '/user/getUserInfo', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES (396, 'p', '9528', '/user/setSelfInfo', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES (399, 'p', '9528', '/user/setUserAuthority', 'POST', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint(20) unsigned DEFAULT NULL,
  `file_chunk_number` bigint(20) DEFAULT NULL,
  `file_chunk_path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of exa_file_chunks
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
BEGIN;
INSERT INTO `exa_file_upload_and_downloads` VALUES (1, '2023-08-23 22:44:11.566', '2023-08-23 22:44:11.566', NULL, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (2, '2023-08-23 22:44:11.566', '2023-08-23 22:44:11.566', NULL, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (3, '2023-08-23 22:55:13.512', '2023-08-23 22:55:13.512', NULL, 'wallpaper1677037330538.jpg', 'uploads/file/27fbd46f8a577320c40eac4a07dffbce_20230823225513.jpg', 'jpg', '27fbd46f8a577320c40eac4a07dffbce_20230823225513.jpg');
INSERT INTO `exa_file_upload_and_downloads` VALUES (4, '2023-08-24 00:22:17.987', '2023-08-24 00:22:17.987', NULL, 'Snipaste_2023-02-15_13-56-06.png', 'uploads/file/d97ad29a9f5ea5d419518ee56c70e596_20230824002217.png', 'png', 'd97ad29a9f5ea5d419518ee56c70e596_20230824002217.png');
COMMIT;

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_md5` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `file_path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `chunk_total` bigint(20) DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of exa_files
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text COLLATE utf8mb4_general_ci COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
BEGIN;
INSERT INTO `jwt_blacklists` VALUES (1, '2023-08-24 00:19:01.808', '2023-08-24 00:19:01.808', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYTk5NDc0NzctNDkwNy00NWRhLWFhZmEtMGE4NTNiNTJlNWQxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6InByaW5jZS5sZWUiLCJBdXRob3JpdHlJZCI6ODg4LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoicW1QbHVzIiwiYXVkIjpbIkdWQSJdLCJleHAiOjE2OTM0MDg3MTEsIm5iZiI6MTY5MjgwMzkxMX0.072bR8HFtAGE9M2zqmjagdTr9sJbBNKgEQI1jnAQoIE');
INSERT INTO `jwt_blacklists` VALUES (2, '2023-08-24 00:20:12.514', '2023-08-24 00:20:12.514', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiYTk5NDc0NzctNDkwNy00NWRhLWFhZmEtMGE4NTNiNTJlNWQxIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6InByaW5jZS5sZWUiLCJBdXRob3JpdHlJZCI6ODg4LCJCdWZmZXJUaW1lIjo4NjQwMCwiaXNzIjoicW1QbHVzIiwiYXVkIjpbIkdWQSJdLCJleHAiOjE2OTM0MTIzOTEsIm5iZiI6MTY5MjgwNzU5MX0.i6hrjFgYvtgapkQ2EnjbZyvjLz1Dx3dD3JHjFfsx3ak');
INSERT INTO `jwt_blacklists` VALUES (3, '2023-08-24 00:21:21.852', '2023-08-24 00:21:21.852', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNGY1ODNjYzMtOTJmMy00MTg1LWJkNmMtMTM5YzJiYThiMWI2IiwiSUQiOjMsIlVzZXJuYW1lIjoibGlodWFuZ3ppIiwiTmlja05hbWUiOiLnrqHnkIblkZgiLCJBdXRob3JpdHlJZCI6OTUyOCwiQnVmZmVyVGltZSI6ODY0MDAsImlzcyI6InFtUGx1cyIsImF1ZCI6WyJHVkEiXSwiZXhwIjoxNjkzNDEyNDI3LCJuYmYiOjE2OTI4MDc2Mjd9.tOhAZoqnynRxKFRX2u_2E0PIaXEUFEct5bApsOosLZw');
COMMIT;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=100 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
BEGIN;
INSERT INTO `sys_apis` VALUES (1, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE');
INSERT INTO `sys_apis` VALUES (3, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/admin_register', '用户注册', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (4, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (6, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (7, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET');
INSERT INTO `sys_apis` VALUES (8, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (9, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (13, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES (19, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (20, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (21, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (22, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT');
INSERT INTO `sys_apis` VALUES (23, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (24, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (26, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (27, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (28, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (29, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (32, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (36, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'GET');
INSERT INTO `sys_apis` VALUES (37, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (38, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (39, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (40, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/fileUploadAndDownload/upload', '文件上传示例', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (41, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (42, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (43, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (44, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (45, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (46, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (47, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/customer/customer', '更新客户', '客户', 'PUT');
INSERT INTO `sys_apis` VALUES (48, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/customer/customer', '创建客户', '客户', 'POST');
INSERT INTO `sys_apis` VALUES (49, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/customer/customer', '删除客户', '客户', 'DELETE');
INSERT INTO `sys_apis` VALUES (50, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/customer/customer', '获取单一客户', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (51, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (52, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (53, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (54, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (55, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (56, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (57, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/createPlug', '自动创建插件包', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (58, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/installPlugin', '安装插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (59, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/pubPlug', '打包插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (60, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/createPackage', '生成包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES (61, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/getPackage', '获取所有包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES (62, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/delPackage', '删除包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES (63, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (64, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (65, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (66, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (67, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT');
INSERT INTO `sys_apis` VALUES (68, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST');
INSERT INTO `sys_apis` VALUES (69, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE');
INSERT INTO `sys_apis` VALUES (70, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (71, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (72, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST');
INSERT INTO `sys_apis` VALUES (73, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE');
INSERT INTO `sys_apis` VALUES (74, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT');
INSERT INTO `sys_apis` VALUES (75, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (76, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (77, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST');
INSERT INTO `sys_apis` VALUES (78, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (79, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (80, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (81, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (82, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST');
INSERT INTO `sys_apis` VALUES (83, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (84, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (85, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (86, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/email/emailSend', '发送邮件示例', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (87, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (88, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (89, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (90, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/chatGpt/getTable', '通过gpt获取内容', '万用表格', 'POST');
INSERT INTO `sys_apis` VALUES (91, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/chatGpt/createSK', '录入sk', '万用表格', 'POST');
INSERT INTO `sys_apis` VALUES (92, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/chatGpt/getSK', '获取sk', '万用表格', 'GET');
INSERT INTO `sys_apis` VALUES (93, '2023-08-23 22:44:11.277', '2023-08-23 22:44:11.277', NULL, '/chatGpt/deleteSK', '删除sk', '万用表格', 'DELETE');
INSERT INTO `sys_apis` VALUES (94, '2023-08-23 23:27:59.000', '2023-08-23 23:27:59.000', NULL, '/orderInfo/createOrderInfo', '新增orderInfo表', 'orderInfo表', 'POST');
INSERT INTO `sys_apis` VALUES (95, '2023-08-23 23:27:59.000', '2023-08-23 23:27:59.000', NULL, '/orderInfo/deleteOrderInfo', '删除orderInfo表', 'orderInfo表', 'DELETE');
INSERT INTO `sys_apis` VALUES (96, '2023-08-23 23:27:59.000', '2023-08-23 23:27:59.000', NULL, '/orderInfo/deleteOrderInfoByIds', '批量删除orderInfo表', 'orderInfo表', 'DELETE');
INSERT INTO `sys_apis` VALUES (97, '2023-08-23 23:27:59.000', '2023-08-23 23:27:59.000', NULL, '/orderInfo/updateOrderInfo', '更新orderInfo表', 'orderInfo表', 'PUT');
INSERT INTO `sys_apis` VALUES (98, '2023-08-23 23:27:59.000', '2023-08-23 23:27:59.000', NULL, '/orderInfo/findOrderInfo', '根据ID获取orderInfo表', 'orderInfo表', 'GET');
INSERT INTO `sys_apis` VALUES (99, '2023-08-23 23:27:59.000', '2023-08-23 23:27:59.000', NULL, '/orderInfo/getOrderInfoList', '获取orderInfo表列表', 'orderInfo表', 'GET');
COMMIT;

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9529 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
BEGIN;
INSERT INTO `sys_authorities` VALUES ('2023-08-23 22:44:11.288', '2023-08-23 22:50:33.205', NULL, 888, '超级管理员', 0, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2023-08-23 22:44:11.288', '2023-08-24 00:24:43.981', NULL, 9528, '管理员', 0, 'order-list');
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint(20) unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_authority_btns
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_authority_menus` VALUES (1, 888);
INSERT INTO `sys_authority_menus` VALUES (2, 888);
INSERT INTO `sys_authority_menus` VALUES (3, 888);
INSERT INTO `sys_authority_menus` VALUES (4, 888);
INSERT INTO `sys_authority_menus` VALUES (5, 888);
INSERT INTO `sys_authority_menus` VALUES (6, 888);
INSERT INTO `sys_authority_menus` VALUES (7, 888);
INSERT INTO `sys_authority_menus` VALUES (8, 888);
INSERT INTO `sys_authority_menus` VALUES (9, 888);
INSERT INTO `sys_authority_menus` VALUES (10, 888);
INSERT INTO `sys_authority_menus` VALUES (11, 888);
INSERT INTO `sys_authority_menus` VALUES (11, 9528);
INSERT INTO `sys_authority_menus` VALUES (12, 888);
INSERT INTO `sys_authority_menus` VALUES (13, 888);
INSERT INTO `sys_authority_menus` VALUES (14, 888);
INSERT INTO `sys_authority_menus` VALUES (15, 888);
INSERT INTO `sys_authority_menus` VALUES (16, 888);
INSERT INTO `sys_authority_menus` VALUES (17, 888);
INSERT INTO `sys_authority_menus` VALUES (18, 888);
INSERT INTO `sys_authority_menus` VALUES (19, 888);
INSERT INTO `sys_authority_menus` VALUES (20, 888);
INSERT INTO `sys_authority_menus` VALUES (21, 888);
INSERT INTO `sys_authority_menus` VALUES (22, 888);
INSERT INTO `sys_authority_menus` VALUES (23, 888);
INSERT INTO `sys_authority_menus` VALUES (24, 888);
INSERT INTO `sys_authority_menus` VALUES (25, 888);
INSERT INTO `sys_authority_menus` VALUES (26, 888);
INSERT INTO `sys_authority_menus` VALUES (27, 888);
INSERT INTO `sys_authority_menus` VALUES (28, 888);
INSERT INTO `sys_authority_menus` VALUES (29, 888);
INSERT INTO `sys_authority_menus` VALUES (30, 888);
INSERT INTO `sys_authority_menus` VALUES (31, 888);
INSERT INTO `sys_authority_menus` VALUES (32, 888);
INSERT INTO `sys_authority_menus` VALUES (32, 9528);
INSERT INTO `sys_authority_menus` VALUES (33, 888);
INSERT INTO `sys_authority_menus` VALUES (33, 9528);
COMMIT;

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `business_db` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `table_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `request_meta` text COLLATE utf8mb4_general_ci,
  `auto_code_path` text COLLATE utf8mb4_general_ci,
  `injection_meta` text COLLATE utf8mb4_general_ci,
  `struct_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `struct_cn_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `api_ids` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `flag` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------
BEGIN;
INSERT INTO `sys_auto_code_histories` VALUES (1, '2023-08-23 01:16:17.000', '2023-08-23 02:35:57.000', '2023-08-23 23:18:58.273', 'orderinfo', '', 'order_info', '{\"structName\":\"OrderInfo\",\"tableName\":\"order_info\",\"packageName\":\"orderInfo\",\"humpPackageName\":\"order_info\",\"abbreviation\":\"orderInfo\",\"description\":\"orderInfo表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"OrderId\",\"fieldDesc\":\"订单ID\",\"fieldType\":\"string\",\"fieldJson\":\"orderId\",\"dataTypeLong\":\"32\",\"comment\":\"订单ID\",\"columnName\":\"order_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"TicketNumber\",\"fieldDesc\":\"订单取票号\",\"fieldType\":\"string\",\"fieldJson\":\"ticketNumber\",\"dataTypeLong\":\"12\",\"comment\":\"订单取票号\",\"columnName\":\"ticket_number\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContactPhone\",\"fieldDesc\":\"订单联系电话\",\"fieldType\":\"string\",\"fieldJson\":\"contactPhone\",\"dataTypeLong\":\"11\",\"comment\":\"订单联系电话\",\"columnName\":\"contact_phone\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsTransfer\",\"fieldDesc\":\"是否换乘.0:false;1:true\",\"fieldType\":\"bool\",\"fieldJson\":\"isTransfer\",\"dataTypeLong\":\"\",\"comment\":\"是否换乘.0:false;1:true\",\"columnName\":\"is_transfer\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsOccupySeat\",\"fieldDesc\":\"是否占座.0:false;1:true\",\"fieldType\":\"int\",\"fieldJson\":\"isOccupySeat\",\"dataTypeLong\":\"\",\"comment\":\"是否占座.0:false;1:true\",\"columnName\":\"is_occupy_seat\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompleteStatus\",\"fieldDesc\":\"完成状态.0:未完成;1:出票成功;2:出票失败\",\"fieldType\":\"int\",\"fieldJson\":\"completeStatus\",\"dataTypeLong\":\"\",\"comment\":\"完成状态.0:未完成;1:出票成功;2:出票失败\",\"columnName\":\"complete_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"FailReason\",\"fieldDesc\":\"出票失败原因\",\"fieldType\":\"string\",\"fieldJson\":\"failReason\",\"dataTypeLong\":\"255\",\"comment\":\"出票失败原因\",\"columnName\":\"fail_reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MachineName\",\"fieldDesc\":\"设备名称\",\"fieldType\":\"string\",\"fieldJson\":\"machineName\",\"dataTypeLong\":\"128\",\"comment\":\"设备名称\",\"columnName\":\"machine_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"orderinfo\"}', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/request/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/web/src/api/orderInfo.js;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderInfo/orderInfoForm.vue;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderInfo/orderInfo.vue;', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/enter.go@ApiGroup@OrderInfoApi;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/enter.go@RouterGroup@OrderInfoRouter;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/enter.go@ServiceGroup@OrderInfoService;', 'OrderInfo', 'orderInfo表', '', 1);
INSERT INTO `sys_auto_code_histories` VALUES (2, '2023-08-23 02:37:02.000', '2023-08-23 02:39:10.000', '2023-08-23 23:18:56.039', 'orderinfo', '', 'order_info', '{\"structName\":\"OrderInfo\",\"tableName\":\"order_info\",\"packageName\":\"orderInfo\",\"humpPackageName\":\"order_info\",\"abbreviation\":\"orderInfo\",\"description\":\"orderInfo表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"OrderId\",\"fieldDesc\":\"订单ID\",\"fieldType\":\"string\",\"fieldJson\":\"orderId\",\"dataTypeLong\":\"32\",\"comment\":\"订单ID\",\"columnName\":\"order_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"TicketNumber\",\"fieldDesc\":\"订单取票号\",\"fieldType\":\"string\",\"fieldJson\":\"ticketNumber\",\"dataTypeLong\":\"12\",\"comment\":\"订单取票号\",\"columnName\":\"ticket_number\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContactPhone\",\"fieldDesc\":\"订单联系电话\",\"fieldType\":\"string\",\"fieldJson\":\"contactPhone\",\"dataTypeLong\":\"11\",\"comment\":\"订单联系电话\",\"columnName\":\"contact_phone\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsTransfer\",\"fieldDesc\":\"是否换乘\",\"fieldType\":\"bool\",\"fieldJson\":\"isTransfer\",\"dataTypeLong\":\"\",\"comment\":\"是否换乘.0:false;1:true\",\"columnName\":\"is_transfer\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsOccupySeat\",\"fieldDesc\":\"是否占座\",\"fieldType\":\"int\",\"fieldJson\":\"isOccupySeat\",\"dataTypeLong\":\"\",\"comment\":\"是否占座.0:false;1:true\",\"columnName\":\"is_occupy_seat\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompleteStatus\",\"fieldDesc\":\"完成状态\",\"fieldType\":\"int\",\"fieldJson\":\"completeStatus\",\"dataTypeLong\":\"\",\"comment\":\"完成状态.0:未完成;1:出票成功;2:出票失败\",\"columnName\":\"complete_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"FailReason\",\"fieldDesc\":\"出票失败原因\",\"fieldType\":\"string\",\"fieldJson\":\"failReason\",\"dataTypeLong\":\"255\",\"comment\":\"出票失败原因\",\"columnName\":\"fail_reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MachineName\",\"fieldDesc\":\"设备名称\",\"fieldType\":\"string\",\"fieldJson\":\"machineName\",\"dataTypeLong\":\"128\",\"comment\":\"设备名称\",\"columnName\":\"machine_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"orderinfo\"}', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/request/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/order_info.go;/Users/leeprince/票务处项目/ticket_office/admin/web/src/api/orderInfo.js;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderInfo/orderInfoForm.vue;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderInfo/orderInfo.vue;', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/enter.go@ApiGroup@OrderInfoApi;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/enter.go@RouterGroup@OrderInfoRouter;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/enter.go@ServiceGroup@OrderInfoService;', 'OrderInfo', 'orderInfo表', '', 1);
INSERT INTO `sys_auto_code_histories` VALUES (3, '2023-08-23 02:40:41.000', '2023-08-23 23:19:23.551', '2023-08-23 23:23:06.530', 'orderinfo', '', 'order_info', '{\"structName\":\"OrderInfo\",\"tableName\":\"order_info\",\"packageName\":\"orderinfo\",\"humpPackageName\":\"orderinfo\",\"abbreviation\":\"orderInfo\",\"description\":\"订单表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"票务服务\",\"fields\":[{\"fieldName\":\"OrderId\",\"fieldDesc\":\"订单ID\",\"fieldType\":\"string\",\"fieldJson\":\"orderId\",\"dataTypeLong\":\"32\",\"comment\":\"订单ID\",\"columnName\":\"order_id\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"TicketNumber\",\"fieldDesc\":\"订单取票号\",\"fieldType\":\"string\",\"fieldJson\":\"ticketNumber\",\"dataTypeLong\":\"12\",\"comment\":\"订单取票号\",\"columnName\":\"ticket_number\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContactPhone\",\"fieldDesc\":\"订单联系电话\",\"fieldType\":\"string\",\"fieldJson\":\"contactPhone\",\"dataTypeLong\":\"11\",\"comment\":\"订单联系电话\",\"columnName\":\"contact_phone\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsTransfer\",\"fieldDesc\":\"是否换乘\",\"fieldType\":\"bool\",\"fieldJson\":\"isTransfer\",\"dataTypeLong\":\"\",\"comment\":\"是否换乘.0:false;1:true\",\"columnName\":\"is_transfer\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsOccupySeat\",\"fieldDesc\":\"是否占座\",\"fieldType\":\"int\",\"fieldJson\":\"isOccupySeat\",\"dataTypeLong\":\"\",\"comment\":\"是否占座.0:false;1:true\",\"columnName\":\"is_occupy_seat\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompleteStatus\",\"fieldDesc\":\"完成状态\",\"fieldType\":\"int\",\"fieldJson\":\"completeStatus\",\"dataTypeLong\":\"\",\"comment\":\"完成状态.0:未完成;1:出票成功;2:出票失败\",\"columnName\":\"complete_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"FailReason\",\"fieldDesc\":\"出票失败原因\",\"fieldType\":\"string\",\"fieldJson\":\"failReason\",\"dataTypeLong\":\"255\",\"comment\":\"出票失败原因\",\"columnName\":\"fail_reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MachineName\",\"fieldDesc\":\"设备名称\",\"fieldType\":\"string\",\"fieldJson\":\"machineName\",\"dataTypeLong\":\"128\",\"comment\":\"设备名称\",\"columnName\":\"machine_name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"orderinfo\"}', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/request/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/web/src/api/orderinfo.js;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfoForm.vue;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfo.vue;', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/enter.go@ApiGroup@OrderInfoApi;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/enter.go@RouterGroup@OrderInfoRouter;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/enter.go@ServiceGroup@OrderInfoService;', 'OrderInfo', '订单表', '', 1);
INSERT INTO `sys_auto_code_histories` VALUES (4, '2023-08-23 23:21:06.805', '2023-08-23 23:23:00.605', NULL, 'orderinfo', '', 'order_info', '{\"structName\":\"OrderInfo\",\"tableName\":\"order_info\",\"packageName\":\"orderinfo\",\"humpPackageName\":\"orderinfo\",\"abbreviation\":\"orderInfo\",\"description\":\"订单表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"票务服务\",\"fields\":[{\"fieldName\":\"OrderId\",\"fieldDesc\":\"订单ID\",\"fieldType\":\"string\",\"fieldJson\":\"orderId\",\"dataTypeLong\":\"32\",\"comment\":\"订单ID\",\"columnName\":\"order_id\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"TicketNumber\",\"fieldDesc\":\"订单取票号\",\"fieldType\":\"string\",\"fieldJson\":\"ticketNumber\",\"dataTypeLong\":\"12\",\"comment\":\"订单取票号\",\"columnName\":\"ticket_number\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContactPhone\",\"fieldDesc\":\"订单联系电话\",\"fieldType\":\"string\",\"fieldJson\":\"contactPhone\",\"dataTypeLong\":\"11\",\"comment\":\"订单联系电话\",\"columnName\":\"contact_phone\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsTransfer\",\"fieldDesc\":\"是否换乘\",\"fieldType\":\"int\",\"fieldJson\":\"isTransfer\",\"dataTypeLong\":\"\",\"comment\":\"是否换乘.0:false;1:true\",\"columnName\":\"is_transfer\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsOccupySeat\",\"fieldDesc\":\"是否占座\",\"fieldType\":\"int\",\"fieldJson\":\"isOccupySeat\",\"dataTypeLong\":\"\",\"comment\":\"是否占座.0:false;1:true\",\"columnName\":\"is_occupy_seat\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompleteStatus\",\"fieldDesc\":\"完成状态\",\"fieldType\":\"int\",\"fieldJson\":\"completeStatus\",\"dataTypeLong\":\"\",\"comment\":\"完成状态.0:未完成;1:出票成功;2:出票失败\",\"columnName\":\"complete_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"FailReason\",\"fieldDesc\":\"出票失败原因\",\"fieldType\":\"string\",\"fieldJson\":\"failReason\",\"dataTypeLong\":\"255\",\"comment\":\"出票失败原因\",\"columnName\":\"fail_reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MachineName\",\"fieldDesc\":\"设备名称\",\"fieldType\":\"string\",\"fieldJson\":\"machineName\",\"dataTypeLong\":\"128\",\"comment\":\"设备名称\",\"columnName\":\"machine_name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"orderinfo\"}', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/request/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/web/src/api/orderinfo.js;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfoForm.vue;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfo.vue;', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/enter.go@ApiGroup@OrderInfoApi;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/enter.go@RouterGroup@OrderInfoRouter;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/enter.go@ServiceGroup@OrderInfoService;', 'OrderInfo', '订单表', '', 1);
INSERT INTO `sys_auto_code_histories` VALUES (5, '2023-08-23 23:24:23.590', '2023-08-23 23:42:45.775', NULL, 'orderinfo', '', 'order_info', '{\"structName\":\"OrderInfo\",\"tableName\":\"order_info\",\"packageName\":\"orderinfo\",\"humpPackageName\":\"orderinfo\",\"abbreviation\":\"orderInfo\",\"description\":\"订单表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"票务服务\",\"fields\":[{\"fieldName\":\"OrderId\",\"fieldDesc\":\"订单ID\",\"fieldType\":\"string\",\"fieldJson\":\"orderId\",\"dataTypeLong\":\"32\",\"comment\":\"订单ID\",\"columnName\":\"order_id\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"TicketNumber\",\"fieldDesc\":\"订单取票号\",\"fieldType\":\"string\",\"fieldJson\":\"ticketNumber\",\"dataTypeLong\":\"12\",\"comment\":\"订单取票号\",\"columnName\":\"ticket_number\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContactPhone\",\"fieldDesc\":\"订单联系电话\",\"fieldType\":\"string\",\"fieldJson\":\"contactPhone\",\"dataTypeLong\":\"11\",\"comment\":\"订单联系电话\",\"columnName\":\"contact_phone\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsTransfer\",\"fieldDesc\":\"是否换乘\",\"fieldType\":\"int\",\"fieldJson\":\"isTransfer\",\"dataTypeLong\":\"\",\"comment\":\"是否换乘.0:false;1:true\",\"columnName\":\"is_transfer\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsOccupySeat\",\"fieldDesc\":\"是否占座\",\"fieldType\":\"int\",\"fieldJson\":\"isOccupySeat\",\"dataTypeLong\":\"\",\"comment\":\"是否占座.0:false;1:true\",\"columnName\":\"is_occupy_seat\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompleteStatus\",\"fieldDesc\":\"完成状态\",\"fieldType\":\"int\",\"fieldJson\":\"completeStatus\",\"dataTypeLong\":\"\",\"comment\":\"完成状态.0:未完成;1:出票成功;2:出票失败\",\"columnName\":\"complete_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"FailReason\",\"fieldDesc\":\"出票失败原因\",\"fieldType\":\"string\",\"fieldJson\":\"failReason\",\"dataTypeLong\":\"255\",\"comment\":\"出票失败原因\",\"columnName\":\"fail_reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MachineName\",\"fieldDesc\":\"设备名称\",\"fieldType\":\"string\",\"fieldJson\":\"machineName\",\"dataTypeLong\":\"128\",\"comment\":\"设备名称\",\"columnName\":\"machine_name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"orderinfo\"}', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/request/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/web/src/api/orderinfo.js;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfoForm.vue;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfo.vue;', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/enter.go@ApiGroup@OrderInfoApi;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/enter.go@RouterGroup@OrderInfoRouter;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/enter.go@ServiceGroup@OrderInfoService;', 'OrderInfo', '订单表', '', 1);
INSERT INTO `sys_auto_code_histories` VALUES (6, '2023-08-23 23:43:24.202', '2023-08-23 23:43:59.470', NULL, 'orderinfo', '', 'order_info', '{\"structName\":\"OrderInfo\",\"tableName\":\"order_info\",\"packageName\":\"orderinfo\",\"humpPackageName\":\"orderinfo\",\"abbreviation\":\"orderInfo\",\"description\":\"订单表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"票务服务\",\"fields\":[{\"fieldName\":\"OrderId\",\"fieldDesc\":\"订单ID\",\"fieldType\":\"string\",\"fieldJson\":\"orderId\",\"dataTypeLong\":\"32\",\"comment\":\"订单ID\",\"columnName\":\"order_id\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"TicketNumber\",\"fieldDesc\":\"订单取票号\",\"fieldType\":\"string\",\"fieldJson\":\"ticketNumber\",\"dataTypeLong\":\"12\",\"comment\":\"订单取票号\",\"columnName\":\"ticket_number\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContactPhone\",\"fieldDesc\":\"订单联系电话\",\"fieldType\":\"string\",\"fieldJson\":\"contactPhone\",\"dataTypeLong\":\"11\",\"comment\":\"订单联系电话\",\"columnName\":\"contact_phone\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsTransfer\",\"fieldDesc\":\"是否换乘\",\"fieldType\":\"int\",\"fieldJson\":\"isTransfer\",\"dataTypeLong\":\"\",\"comment\":\"是否换乘.0:false;1:true\",\"columnName\":\"is_transfer\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsOccupySeat\",\"fieldDesc\":\"是否占座\",\"fieldType\":\"int\",\"fieldJson\":\"isOccupySeat\",\"dataTypeLong\":\"\",\"comment\":\"是否占座.0:false;1:true\",\"columnName\":\"is_occupy_seat\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompleteStatus\",\"fieldDesc\":\"完成状态\",\"fieldType\":\"int\",\"fieldJson\":\"completeStatus\",\"dataTypeLong\":\"\",\"comment\":\"完成状态.0:未完成;1:出票成功;2:出票失败\",\"columnName\":\"complete_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"FailReason\",\"fieldDesc\":\"出票失败原因\",\"fieldType\":\"string\",\"fieldJson\":\"failReason\",\"dataTypeLong\":\"255\",\"comment\":\"出票失败原因\",\"columnName\":\"fail_reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MachineName\",\"fieldDesc\":\"设备名称\",\"fieldType\":\"string\",\"fieldJson\":\"machineName\",\"dataTypeLong\":\"128\",\"comment\":\"设备名称\",\"columnName\":\"machine_name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"orderinfo\"}', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/request/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/web/src/api/orderinfo.js;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfoForm.vue;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfo.vue;', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/enter.go@ApiGroup@OrderInfoApi;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/enter.go@RouterGroup@OrderInfoRouter;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/enter.go@ServiceGroup@OrderInfoService;', 'OrderInfo', '订单表', '', 1);
INSERT INTO `sys_auto_code_histories` VALUES (7, '2023-08-23 23:44:36.082', '2023-08-23 23:47:36.858', NULL, 'orderinfo', '', 'order_info', '{\"structName\":\"OrderInfo\",\"tableName\":\"order_info\",\"packageName\":\"orderinfo\",\"humpPackageName\":\"orderinfo\",\"abbreviation\":\"orderInfo\",\"description\":\"订单表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"票务服务\",\"fields\":[{\"fieldName\":\"OrderId\",\"fieldDesc\":\"订单ID\",\"fieldType\":\"string\",\"fieldJson\":\"orderId\",\"dataTypeLong\":\"32\",\"comment\":\"订单ID\",\"columnName\":\"order_id\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"TicketNumber\",\"fieldDesc\":\"订单取票号\",\"fieldType\":\"string\",\"fieldJson\":\"ticketNumber\",\"dataTypeLong\":\"12\",\"comment\":\"订单取票号\",\"columnName\":\"ticket_number\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContactPhone\",\"fieldDesc\":\"订单联系电话\",\"fieldType\":\"string\",\"fieldJson\":\"contactPhone\",\"dataTypeLong\":\"11\",\"comment\":\"订单联系电话\",\"columnName\":\"contact_phone\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":true},{\"fieldName\":\"IsTransfer\",\"fieldDesc\":\"是否换乘\",\"fieldType\":\"int\",\"fieldJson\":\"isTransfer\",\"dataTypeLong\":\"\",\"comment\":\"是否换乘.0:false;1:true\",\"columnName\":\"is_transfer\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsOccupySeat\",\"fieldDesc\":\"是否占座\",\"fieldType\":\"int\",\"fieldJson\":\"isOccupySeat\",\"dataTypeLong\":\"\",\"comment\":\"是否占座.0:false;1:true\",\"columnName\":\"is_occupy_seat\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompleteStatus\",\"fieldDesc\":\"完成状态\",\"fieldType\":\"int\",\"fieldJson\":\"completeStatus\",\"dataTypeLong\":\"\",\"comment\":\"完成状态.0:未完成;1:出票成功;2:出票失败\",\"columnName\":\"complete_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"FailReason\",\"fieldDesc\":\"出票失败原因\",\"fieldType\":\"string\",\"fieldJson\":\"failReason\",\"dataTypeLong\":\"255\",\"comment\":\"出票失败原因\",\"columnName\":\"fail_reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MachineName\",\"fieldDesc\":\"设备名称\",\"fieldType\":\"string\",\"fieldJson\":\"machineName\",\"dataTypeLong\":\"128\",\"comment\":\"设备名称\",\"columnName\":\"machine_name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"orderinfo\"}', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/request/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/web/src/api/orderinfo.js;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfoForm.vue;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfo.vue;', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/enter.go@ApiGroup@OrderInfoApi;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/enter.go@RouterGroup@OrderInfoRouter;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/enter.go@ServiceGroup@OrderInfoService;', 'OrderInfo', '订单表', '', 1);
INSERT INTO `sys_auto_code_histories` VALUES (8, '2023-08-23 23:48:13.623', '2023-08-23 23:48:13.623', NULL, 'orderinfo', '', 'order_info', '{\"structName\":\"OrderInfo\",\"tableName\":\"order_info\",\"packageName\":\"orderinfo\",\"humpPackageName\":\"orderinfo\",\"abbreviation\":\"orderInfo\",\"description\":\"订单表\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"票务服务\",\"fields\":[{\"fieldName\":\"OrderId\",\"fieldDesc\":\"订单ID\",\"fieldType\":\"string\",\"fieldJson\":\"orderId\",\"dataTypeLong\":\"32\",\"comment\":\"订单ID\",\"columnName\":\"order_id\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"TicketNumber\",\"fieldDesc\":\"订单取票号\",\"fieldType\":\"string\",\"fieldJson\":\"ticketNumber\",\"dataTypeLong\":\"12\",\"comment\":\"订单取票号\",\"columnName\":\"ticket_number\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ContactPhone\",\"fieldDesc\":\"订单联系电话\",\"fieldType\":\"string\",\"fieldJson\":\"contactPhone\",\"dataTypeLong\":\"11\",\"comment\":\"订单联系电话\",\"columnName\":\"contact_phone\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsTransfer\",\"fieldDesc\":\"是否换乘\",\"fieldType\":\"int\",\"fieldJson\":\"isTransfer\",\"dataTypeLong\":\"\",\"comment\":\"是否换乘.0:false;1:true\",\"columnName\":\"is_transfer\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"IsOccupySeat\",\"fieldDesc\":\"是否占座\",\"fieldType\":\"int\",\"fieldJson\":\"isOccupySeat\",\"dataTypeLong\":\"\",\"comment\":\"是否占座.0:false;1:true\",\"columnName\":\"is_occupy_seat\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompleteStatus\",\"fieldDesc\":\"完成状态\",\"fieldType\":\"int\",\"fieldJson\":\"completeStatus\",\"dataTypeLong\":\"\",\"comment\":\"完成状态.0:未完成;1:出票成功;2:出票失败\",\"columnName\":\"complete_status\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"FailReason\",\"fieldDesc\":\"出票失败原因\",\"fieldType\":\"string\",\"fieldJson\":\"failReason\",\"dataTypeLong\":\"255\",\"comment\":\"出票失败原因\",\"columnName\":\"fail_reason\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MachineName\",\"fieldDesc\":\"设备名称\",\"fieldType\":\"string\",\"fieldJson\":\"machineName\",\"dataTypeLong\":\"128\",\"comment\":\"设备名称\",\"columnName\":\"machine_name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"orderinfo\"}', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/model/orderinfo/request/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/orderinfo.go;/Users/leeprince/票务处项目/ticket_office/admin/web/src/api/orderinfo.js;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfoForm.vue;/Users/leeprince/票务处项目/ticket_office/admin/web/src/view/orderinfo/orderinfo.vue;', '/Users/leeprince/票务处项目/ticket_office/admin/server/api/v1/orderinfo/enter.go@ApiGroup@OrderInfoApi;/Users/leeprince/票务处项目/ticket_office/admin/server/router/orderinfo/enter.go@RouterGroup@OrderInfoRouter;/Users/leeprince/票务处项目/ticket_office/admin/server/service/orderinfo/enter.go@ServiceGroup@OrderInfoService;', 'OrderInfo', '订单表', '', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_auto_codes
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_codes`;
CREATE TABLE `sys_auto_codes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '包名',
  `label` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示名',
  `desc` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_codes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_auto_codes
-- ----------------------------
BEGIN;
INSERT INTO `sys_auto_codes` VALUES (1, '2023-08-23 23:19:49.000', '2023-08-23 23:19:49.000', NULL, 'orderinfo', '统计订单', '统计订单信息');
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_base_menu_btns
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menu_btns` VALUES (1, '2023-08-24 00:01:53.459', '2023-08-24 00:01:53.459', NULL, 'orderListUpdate', '变更', 33);
INSERT INTO `sys_base_menu_btns` VALUES (2, '2023-08-24 00:01:53.459', '2023-08-24 00:01:53.459', NULL, 'orderListDelete', '删除', 33);
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=34 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
BEGIN;
INSERT INTO `sys_base_menus` VALUES (1, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '', 0, 0, '仪表盘', 'odometer', 0);
INSERT INTO `sys_base_menus` VALUES (2, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '0', 'about', 'about', 0, 'view/about/index.vue', 9, '', 0, 0, '关于我们', 'info-filled', 0);
INSERT INTO `sys_base_menus` VALUES (3, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '', 0, 0, '超级管理员', 'user', 0);
INSERT INTO `sys_base_menus` VALUES (4, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, '', 0, 0, '角色管理', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (5, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, '', 1, 0, '菜单管理', 'tickets', 0);
INSERT INTO `sys_base_menus` VALUES (6, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, '', 1, 0, 'api管理', 'platform', 0);
INSERT INTO `sys_base_menus` VALUES (7, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, '', 0, 0, '用户管理', 'coordinate', 0);
INSERT INTO `sys_base_menus` VALUES (8, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '', 0, 0, '字典管理', 'notebook', 0);
INSERT INTO `sys_base_menus` VALUES (9, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 'dictionary', 0, 0, '字典详情-${id}', 'list', 0);
INSERT INTO `sys_base_menus` VALUES (10, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '', 0, 0, '操作历史', 'pie-chart', 0);
INSERT INTO `sys_base_menus` VALUES (11, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 4, '', 0, 0, '个人信息', 'message', 0);
INSERT INTO `sys_base_menus` VALUES (12, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 7, '', 0, 0, '示例文件', 'management', 0);
INSERT INTO `sys_base_menus` VALUES (13, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '12', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, '', 0, 0, '媒体库（上传下载）', 'upload', 0);
INSERT INTO `sys_base_menus` VALUES (14, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '12', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, '', 0, 0, '断点续传', 'upload-filled', 0);
INSERT INTO `sys_base_menus` VALUES (15, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '12', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, '', 0, 0, '客户列表（资源示例）', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (16, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, '', 0, 0, '系统工具', 'tools', 0);
INSERT INTO `sys_base_menus` VALUES (17, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '16', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '', 1, 0, '代码生成器', 'cpu', 0);
INSERT INTO `sys_base_menus` VALUES (18, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '16', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, '', 1, 0, '表单生成器', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (19, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '16', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, '', 0, 0, '系统配置', 'operation', 0);
INSERT INTO `sys_base_menus` VALUES (20, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '16', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, '', 0, 0, '自动化代码管理', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (21, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '16', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, '', 0, 0, '自动化代码-${id}', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (22, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '16', 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, '', 0, 0, '自动化package', 'folder', 0);
INSERT INTO `sys_base_menus` VALUES (23, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '0', 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', 0, '', 0, 0, '官方网站', 'home-filled', 0);
INSERT INTO `sys_base_menus` VALUES (24, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '0', 'state', 'state', 0, 'view/system/state.vue', 8, '', 0, 0, '服务器状态', 'cloudy', 0);
INSERT INTO `sys_base_menus` VALUES (25, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '0', 'plugin', 'plugin', 0, 'view/routerHolder.vue', 6, '', 0, 0, '插件系统', 'cherry', 0);
INSERT INTO `sys_base_menus` VALUES (26, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '25', 'https://plugin.gin-vue-admin.com/', 'https://plugin.gin-vue-admin.com/', 0, 'https://plugin.gin-vue-admin.com/', 0, '', 0, 0, '插件市场', 'shop', 0);
INSERT INTO `sys_base_menus` VALUES (27, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '25', 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, '', 0, 0, '插件安装', 'box', 0);
INSERT INTO `sys_base_menus` VALUES (28, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '25', 'autoPlug', 'autoPlug', 0, 'view/systemTools/autoPlug/autoPlug.vue', 2, '', 0, 0, '插件模板', 'folder', 0);
INSERT INTO `sys_base_menus` VALUES (29, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '25', 'pubPlug', 'pubPlug', 0, 'view/systemTools/pubPlug/pubPlug.vue', 3, '', 0, 0, '打包插件', 'files', 0);
INSERT INTO `sys_base_menus` VALUES (30, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '25', 'plugin-email', 'plugin-email', 0, 'plugin/email/view/index.vue', 4, '', 0, 0, '邮件插件', 'message', 0);
INSERT INTO `sys_base_menus` VALUES (31, '2023-08-23 22:44:11.311', '2023-08-23 22:44:11.311', NULL, 0, '16', 'chatTable', 'chatTable', 0, 'view/chatgpt/chatTable.vue', 6, '', 0, 0, '万用表格', 'chat-dot-square', 0);
INSERT INTO `sys_base_menus` VALUES (32, '2023-08-23 01:59:24.000', '2023-08-23 02:45:02.000', NULL, 0, '0', 'order-info', 'order-info', 0, 'view/routerHolder.vue', 0, '', 0, 0, '订单信息', 'coin', 0);
INSERT INTO `sys_base_menus` VALUES (33, '2023-08-23 02:22:51.000', '2023-08-24 00:07:12.911', NULL, 0, '32', 'order-list', 'order-list', 0, 'view/orderInfo/orderinfo.vue', 0, '', 0, 0, '订单列表', 'list', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_chat_gpt_options
-- ----------------------------
DROP TABLE IF EXISTS `sys_chat_gpt_options`;
CREATE TABLE `sys_chat_gpt_options` (
  `sk` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_chat_gpt_options
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
BEGIN;
INSERT INTO `sys_data_authority_id` VALUES (888, 888);
INSERT INTO `sys_data_authority_id` VALUES (888, 8881);
INSERT INTO `sys_data_authority_id` VALUES (888, 9528);
INSERT INTO `sys_data_authority_id` VALUES (9528, 8881);
INSERT INTO `sys_data_authority_id` VALUES (9528, 9528);
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
BEGIN;
INSERT INTO `sys_dictionaries` VALUES (1, '2023-08-23 23:04:58.000', '2023-08-23 22:44:11.297', NULL, '性别', 'gender', 1, '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '2023-08-23 22:44:11.294', '2023-08-23 22:44:11.299', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '2023-08-23 22:44:11.294', '2023-08-23 22:44:11.301', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2023-08-23 22:44:11.294', '2023-08-23 22:44:11.304', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '2023-08-23 22:44:11.294', '2023-08-23 22:44:11.306', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '2023-08-23 22:44:11.294', '2023-08-23 22:44:11.309', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');
COMMIT;

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '展示值',
  `value` bigint(20) DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
BEGIN;
INSERT INTO `sys_dictionary_details` VALUES (1, '2023-08-23 22:44:11.297', '2023-08-23 22:44:11.297', NULL, '男', 1, 1, 1, 1);
INSERT INTO `sys_dictionary_details` VALUES (2, '2023-08-23 22:44:11.297', '2023-08-23 22:44:11.297', NULL, '女', 2, 1, 2, 1);
INSERT INTO `sys_dictionary_details` VALUES (3, '2023-08-23 22:44:11.299', '2023-08-23 22:44:11.299', NULL, 'smallint', 1, 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '2023-08-23 22:44:11.299', '2023-08-23 22:44:11.299', NULL, 'mediumint', 2, 1, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '2023-08-23 22:44:11.299', '2023-08-23 22:44:11.299', NULL, 'int', 3, 1, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (6, '2023-08-23 22:44:11.299', '2023-08-23 22:44:11.299', NULL, 'bigint', 4, 1, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (7, '2023-08-23 22:44:11.301', '2023-08-23 22:44:11.301', NULL, 'date', 0, 1, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (8, '2023-08-23 22:44:11.301', '2023-08-23 22:44:11.301', NULL, 'time', 1, 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (9, '2023-08-23 22:44:11.301', '2023-08-23 22:44:11.301', NULL, 'year', 2, 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (10, '2023-08-23 22:44:11.301', '2023-08-23 22:44:11.301', NULL, 'datetime', 3, 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (11, '2023-08-23 22:44:11.301', '2023-08-23 22:44:11.301', NULL, 'timestamp', 5, 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (12, '2023-08-23 22:44:11.304', '2023-08-23 22:44:11.304', NULL, 'float', 0, 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (13, '2023-08-23 22:44:11.304', '2023-08-23 22:44:11.304', NULL, 'double', 1, 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (14, '2023-08-23 22:44:11.304', '2023-08-23 22:44:11.304', NULL, 'decimal', 2, 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (15, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'char', 0, 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (16, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'varchar', 1, 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (17, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'tinyblob', 2, 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (18, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'tinytext', 3, 1, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (19, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'text', 4, 1, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (20, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'blob', 5, 1, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (21, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'mediumblob', 6, 1, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (22, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'mediumtext', 7, 1, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'longblob', 8, 1, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (24, '2023-08-23 22:44:11.306', '2023-08-23 22:44:11.306', NULL, 'longtext', 9, 1, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (25, '2023-08-23 22:44:11.309', '2023-08-23 23:04:25.000', NULL, 'tinyint', 0, 1, 0, 6);
COMMIT;

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '错误信息',
  `body` text COLLATE utf8mb4_general_ci COMMENT '请求Body',
  `resp` text COLLATE utf8mb4_general_ci COMMENT '响应Body',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=48 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
BEGIN;
INSERT INTO `sys_operation_records` VALUES (1, '2023-08-23 22:49:24.192', '2023-08-23 22:49:24.192', NULL, '127.0.0.1', 'POST', '/authority/deleteAuthority', 200, 3289831, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":8881}', '{\"code\":7,\"data\":{},\"msg\":\"删除失败此角色有用户正在使用禁止删除\"}', 1);
INSERT INTO `sys_operation_records` VALUES (2, '2023-08-23 22:49:36.105', '2023-08-23 22:49:36.105', NULL, '127.0.0.1', 'POST', '/authority/deleteAuthority', 200, 1800488, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":8881}', '{\"code\":7,\"data\":{},\"msg\":\"删除失败此角色有用户正在使用禁止删除\"}', 1);
INSERT INTO `sys_operation_records` VALUES (3, '2023-08-23 22:49:45.442', '2023-08-23 22:49:45.442', NULL, '127.0.0.1', 'POST', '/user/setUserAuthorities', 200, 49726025, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":1,\"authorityIds\":[888,9528]}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (4, '2023-08-23 22:49:55.486', '2023-08-23 22:49:55.486', NULL, '127.0.0.1', 'POST', '/authority/deleteAuthority', 200, 34748658, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":8881}', '{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (5, '2023-08-23 22:50:01.954', '2023-08-23 22:50:01.954', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 105546645, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"menus\":[{\"ID\":23,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"https://www.gin-vue-admin.com\",\"name\":\"https://www.gin-vue-admin.com\",\"hidden\":false,\"component\":\"/\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"官方网站\",\"icon\":\"home-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":32,\"CreatedAt\":\"2023-08-23T01:59:24+08:00\",\"UpdatedAt\":\"2023-08-23T02:45:02+08:00\",\"parentId\":\"0\",\"path\":\"order-info\",\"name\":\"order-info\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单信息\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-23T02:45:30+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-23T02:45:30+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":1,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"dashboard\",\"name\":\"dashboard\",\"hidden\":false,\"component\":\"view/dashboard/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"仪表盘\",\"icon\":\"odometer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":3,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"超级管理员\",\"icon\":\"user\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":4,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":9,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"activeName\":\"dictionary\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情-${id}\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":10,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":9,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"activeName\":\"dictionary\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情-${id}\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":10,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"systemTools\",\"name\":\"systemTools\",\"hidden\":false,\"component\":\"view/systemTools/index.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统工具\",\"icon\":\"tools\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":21,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码-${id}\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":22,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"autoPkg\",\"name\":\"autoPkg\",\"hidden\":false,\"component\":\"view/systemTools/autoPkg/autoPkg.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化package\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":31,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"chatTable\",\"name\":\"chatTable\",\"hidden\":false,\"component\":\"view/chatgpt/chatTable.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"万用表格\",\"icon\":\"chat-dot-square\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":21,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码-${id}\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":22,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"autoPkg\",\"name\":\"autoPkg\",\"hidden\":false,\"component\":\"view/systemTools/autoPkg/autoPkg.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化package\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":31,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"16\",\"path\":\"chatTable\",\"name\":\"chatTable\",\"hidden\":false,\"component\":\"view/chatgpt/chatTable.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"万用表格\",\"icon\":\"chat-dot-square\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":25,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"plugin\",\"name\":\"plugin\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件系统\",\"icon\":\"cherry\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":26,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"https://plugin.gin-vue-admin.com/\",\"name\":\"https://plugin.gin-vue-admin.com/\",\"hidden\":false,\"component\":\"https://plugin.gin-vue-admin.com/\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件市场\",\"icon\":\"shop\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":27,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"installPlugin\",\"name\":\"installPlugin\",\"hidden\":false,\"component\":\"view/systemTools/installPlugin/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件安装\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":28,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"autoPlug\",\"name\":\"autoPlug\",\"hidden\":false,\"component\":\"view/systemTools/autoPlug/autoPlug.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件模板\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":29,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"pubPlug\",\"name\":\"pubPlug\",\"hidden\":false,\"component\":\"view/systemTools/pubPlug/pubPlug.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"打包插件\",\"icon\":\"files\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":30,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"plugin-email\",\"name\":\"plugin-email\",\"hidden\":false,\"component\":\"plugin/email/view/index.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"邮件插件\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":26,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"https://plugin.gin-vue-admin.com/\",\"name\":\"https://plugin.gin-vue-admin.com/\",\"hidden\":false,\"component\":\"https://plugin.gin-vue-admin.com/\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件市场\",\"icon\":\"shop\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":27,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"installPlugin\",\"name\":\"installPlugin\",\"hidden\":false,\"component\":\"view/systemTools/installPlugin/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件安装\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":28,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"autoPlug\",\"name\":\"autoPlug\",\"hidden\":false,\"component\":\"view/systemTools/autoPlug/autoPlug.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件模板\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":29,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"pubPlug\",\"name\":\"pubPlug\",\"hidden\":false,\"component\":\"view/systemTools/pubPlug/pubPlug.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"打包插件\",\"icon\":\"files\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":30,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"25\",\"path\":\"plugin-email\",\"name\":\"plugin-email\",\"hidden\":false,\"component\":\"plugin/email/view/index.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"邮件插件\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"example\",\"name\":\"example\",\"hidden\":false,\"component\":\"view/example/index.vue\",\"sort\":7,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"示例文件\",\"icon\":\"management\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":13,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"12\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":14,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"12\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"12\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"12\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":14,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"12\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"12\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"state\",\"name\":\"state\",\"hidden\":false,\"component\":\"view/system/state.vue\",\"sort\":8,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器状态\",\"icon\":\"cloudy\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":2,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"about\",\"name\":\"about\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":9,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":888}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (6, '2023-08-23 22:50:05.274', '2023-08-23 22:50:05.274', NULL, '127.0.0.1', 'POST', '/casbin/updateCasbin', 200, 75348486, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":888,\"casbinInfos\":[{\"path\":\"/jwt/jsonInBlacklist\",\"method\":\"POST\"},{\"path\":\"/user/deleteUser\",\"method\":\"DELETE\"},{\"path\":\"/user/admin_register\",\"method\":\"POST\"},{\"path\":\"/user/getUserList\",\"method\":\"POST\"},{\"path\":\"/user/setUserInfo\",\"method\":\"PUT\"},{\"path\":\"/user/setSelfInfo\",\"method\":\"PUT\"},{\"path\":\"/user/getUserInfo\",\"method\":\"GET\"},{\"path\":\"/user/setUserAuthorities\",\"method\":\"POST\"},{\"path\":\"/user/changePassword\",\"method\":\"POST\"},{\"path\":\"/user/setUserAuthority\",\"method\":\"POST\"},{\"path\":\"/user/resetPassword\",\"method\":\"POST\"},{\"path\":\"/api/createApi\",\"method\":\"POST\"},{\"path\":\"/api/deleteApi\",\"method\":\"POST\"},{\"path\":\"/api/updateApi\",\"method\":\"POST\"},{\"path\":\"/api/getApiList\",\"method\":\"POST\"},{\"path\":\"/api/getAllApis\",\"method\":\"POST\"},{\"path\":\"/api/getApiById\",\"method\":\"POST\"},{\"path\":\"/api/deleteApisByIds\",\"method\":\"DELETE\"},{\"path\":\"/authority/copyAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/createAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/deleteAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/updateAuthority\",\"method\":\"PUT\"},{\"path\":\"/authority/getAuthorityList\",\"method\":\"POST\"},{\"path\":\"/authority/setDataAuthority\",\"method\":\"POST\"},{\"path\":\"/casbin/updateCasbin\",\"method\":\"POST\"},{\"path\":\"/casbin/getPolicyPathByAuthorityId\",\"method\":\"POST\"},{\"path\":\"/menu/addBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getMenu\",\"method\":\"POST\"},{\"path\":\"/menu/deleteBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/updateBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuById\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuList\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuTree\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/menu/addMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/findFile\",\"method\":\"GET\"},{\"path\":\"/fileUploadAndDownload/breakpointContinue\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/breakpointContinueFinish\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/removeChunk\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/upload\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/deleteFile\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/editFileName\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/getFileList\",\"method\":\"POST\"},{\"path\":\"/system/getServerInfo\",\"method\":\"POST\"},{\"path\":\"/system/getSystemConfig\",\"method\":\"POST\"},{\"path\":\"/system/setSystemConfig\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"PUT\"},{\"path\":\"/customer/customer\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"DELETE\"},{\"path\":\"/customer/customer\",\"method\":\"GET\"},{\"path\":\"/customer/customerList\",\"method\":\"GET\"},{\"path\":\"/autoCode/getDB\",\"method\":\"GET\"},{\"path\":\"/autoCode/getTables\",\"method\":\"GET\"},{\"path\":\"/autoCode/createTemp\",\"method\":\"POST\"},{\"path\":\"/autoCode/preview\",\"method\":\"POST\"},{\"path\":\"/autoCode/getColumn\",\"method\":\"GET\"},{\"path\":\"/autoCode/createPlug\",\"method\":\"POST\"},{\"path\":\"/autoCode/installPlugin\",\"method\":\"POST\"},{\"path\":\"/autoCode/pubPlug\",\"method\":\"POST\"},{\"path\":\"/autoCode/createPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/getPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/delPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/getMeta\",\"method\":\"POST\"},{\"path\":\"/autoCode/rollback\",\"method\":\"POST\"},{\"path\":\"/autoCode/getSysHistory\",\"method\":\"POST\"},{\"path\":\"/autoCode/delSysHistory\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/updateSysDictionaryDetail\",\"method\":\"PUT\"},{\"path\":\"/sysDictionaryDetail/createSysDictionaryDetail\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/deleteSysDictionaryDetail\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionaryDetail/findSysDictionaryDetail\",\"method\":\"GET\"},{\"path\":\"/sysDictionaryDetail/getSysDictionaryDetailList\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/createSysDictionary\",\"method\":\"POST\"},{\"path\":\"/sysDictionary/deleteSysDictionary\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionary/updateSysDictionary\",\"method\":\"PUT\"},{\"path\":\"/sysDictionary/findSysDictionary\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/getSysDictionaryList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/createSysOperationRecord\",\"method\":\"POST\"},{\"path\":\"/sysOperationRecord/findSysOperationRecord\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/getSysOperationRecordList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecord\",\"method\":\"DELETE\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecordByIds\",\"method\":\"DELETE\"},{\"path\":\"/simpleUploader/upload\",\"method\":\"POST\"},{\"path\":\"/simpleUploader/checkFileMd5\",\"method\":\"GET\"},{\"path\":\"/simpleUploader/mergeFileMd5\",\"method\":\"GET\"},{\"path\":\"/email/emailTest\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/setAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/getAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/canRemoveAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/chatGpt/getTable\",\"method\":\"POST\"},{\"path\":\"/chatGpt/createSK\",\"method\":\"POST\"},{\"path\":\"/chatGpt/getSK\",\"method\":\"GET\"},{\"path\":\"/chatGpt/deleteSK\",\"method\":\"DELETE\"},{\"path\":\"/orderInfo/createOrderInfo\",\"method\":\"POST\"},{\"path\":\"/orderInfo/deleteOrderInfo\",\"method\":\"DELETE\"},{\"path\":\"/orderInfo/deleteOrderInfoByIds\",\"method\":\"DELETE\"},{\"path\":\"/orderInfo/updateOrderInfo\",\"method\":\"PUT\"},{\"path\":\"/orderInfo/findOrderInfo\",\"method\":\"GET\"},{\"path\":\"/orderInfo/getOrderInfoList\",\"method\":\"GET\"}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (7, '2023-08-23 22:50:33.264', '2023-08-23 22:50:33.264', NULL, '127.0.0.1', 'PUT', '/authority/updateAuthority', 200, 59330640, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":888,\"authorityName\":\"超级管理员\",\"parentId\":0}', '{\"code\":0,\"data\":{\"authority\":{\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"authorityId\":888,\"authorityName\":\"超级管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"\"}},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (8, '2023-08-23 22:50:50.713', '2023-08-23 22:50:50.713', NULL, '127.0.0.1', 'PUT', '/authority/updateAuthority', 200, 95627106, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0}', '{\"code\":0,\"data\":{\"authority\":{\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"\"}},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (9, '2023-08-23 22:53:47.069', '2023-08-23 22:53:47.069', NULL, '127.0.0.1', 'POST', '/user/setUserAuthorities', 200, 12864482, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":1,\"authorityIds\":[888]}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (10, '2023-08-23 22:53:50.727', '2023-08-23 22:53:50.727', NULL, '127.0.0.1', 'POST', '/user/setUserAuthorities', 200, 23247948, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":2,\"authorityIds\":[]}', '{\"code\":7,\"data\":{},\"msg\":\"修改失败\"}', 1);
INSERT INTO `sys_operation_records` VALUES (11, '2023-08-23 22:53:54.890', '2023-08-23 22:53:54.890', NULL, '127.0.0.1', 'POST', '/user/setUserAuthorities', 200, 60858016, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":2,\"authorityIds\":[9528]}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (12, '2023-08-23 22:53:55.884', '2023-08-23 22:53:55.884', NULL, '127.0.0.1', 'POST', '/user/setUserAuthorities', 200, 33089551, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":2,\"authorityIds\":[9528]}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (13, '2023-08-23 22:54:16.514', '2023-08-23 22:54:16.514', NULL, '127.0.0.1', 'PUT', '/user/setUserInfo', 200, 88326318, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":1,\"CreatedAt\":\"2023-08-23T22:44:11.539+08:00\",\"UpdatedAt\":\"2023-08-23T22:49:45.394+08:00\",\"uuid\":\"a9947477-4907-45da-aafa-0a853b52e5d1\",\"userName\":\"admin\",\"nickName\":\"prince.lee\",\"sideMode\":\"dark\",\"headerImg\":\"https://qmplusimg.henrongyi.top/gva_header.jpg\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"authorityId\":888,\"authority\":{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:33.205+08:00\",\"DeletedAt\":null,\"authorityId\":888,\"authorityName\":\"超级管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"},\"authorities\":[{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:33.205+08:00\",\"DeletedAt\":null,\"authorityId\":888,\"authorityName\":\"超级管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"},{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:50.618+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"}],\"phone\":\"13510480000\",\"email\":\"333333333@qq.com\",\"enable\":1,\"authorityIds\":[888]}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (14, '2023-08-23 22:54:34.176', '2023-08-23 22:54:34.176', NULL, '127.0.0.1', 'PUT', '/user/setUserInfo', 200, 80687500, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":2,\"CreatedAt\":\"2023-08-23T22:44:11.539+08:00\",\"UpdatedAt\":\"2023-08-23T22:53:55.852+08:00\",\"uuid\":\"ee5745dc-27b4-4c1e-a598-2d75aa51f1c7\",\"userName\":\"a303176530\",\"nickName\":\"管理员\",\"sideMode\":\"dark\",\"headerImg\":\"https:///qmplusimg.henrongyi.top/1572075907logo.png\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"authorityId\":9528,\"authority\":{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:50.618+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"},\"authorities\":[{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:50.618+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"}],\"phone\":\"17611111111\",\"email\":\"333333333@qq.com\",\"enable\":1,\"authorityIds\":[9528]}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (15, '2023-08-23 22:55:34.493', '2023-08-23 22:55:34.493', NULL, '127.0.0.1', 'PUT', '/user/setUserInfo', 200, 15632876, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":1,\"CreatedAt\":\"2023-08-23T22:44:11.539+08:00\",\"UpdatedAt\":\"2023-08-23T22:54:16.512+08:00\",\"uuid\":\"a9947477-4907-45da-aafa-0a853b52e5d1\",\"userName\":\"admin\",\"nickName\":\"prince.lee\",\"sideMode\":\"dark\",\"headerImg\":\"uploads/file/27fbd46f8a577320c40eac4a07dffbce_20230823225513.jpg\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"authorityId\":888,\"authority\":{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:33.205+08:00\",\"DeletedAt\":null,\"authorityId\":888,\"authorityName\":\"超级管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"},\"authorities\":[{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:33.205+08:00\",\"DeletedAt\":null,\"authorityId\":888,\"authorityName\":\"超级管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"}],\"phone\":\"13510480000\",\"email\":\"333333333@qq.com\",\"enable\":1,\"authorityIds\":[888]}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (16, '2023-08-23 23:53:28.875', '2023-08-23 23:53:28.875', NULL, '127.0.0.1', 'POST', '/casbin/updateCasbin', 200, 61324541, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":888,\"casbinInfos\":[{\"path\":\"/jwt/jsonInBlacklist\",\"method\":\"POST\"},{\"path\":\"/user/deleteUser\",\"method\":\"DELETE\"},{\"path\":\"/user/admin_register\",\"method\":\"POST\"},{\"path\":\"/user/getUserList\",\"method\":\"POST\"},{\"path\":\"/user/setUserInfo\",\"method\":\"PUT\"},{\"path\":\"/user/setSelfInfo\",\"method\":\"PUT\"},{\"path\":\"/user/getUserInfo\",\"method\":\"GET\"},{\"path\":\"/user/setUserAuthorities\",\"method\":\"POST\"},{\"path\":\"/user/changePassword\",\"method\":\"POST\"},{\"path\":\"/user/setUserAuthority\",\"method\":\"POST\"},{\"path\":\"/user/resetPassword\",\"method\":\"POST\"},{\"path\":\"/api/createApi\",\"method\":\"POST\"},{\"path\":\"/api/deleteApi\",\"method\":\"POST\"},{\"path\":\"/api/updateApi\",\"method\":\"POST\"},{\"path\":\"/api/getApiList\",\"method\":\"POST\"},{\"path\":\"/api/getAllApis\",\"method\":\"POST\"},{\"path\":\"/api/getApiById\",\"method\":\"POST\"},{\"path\":\"/api/deleteApisByIds\",\"method\":\"DELETE\"},{\"path\":\"/authority/copyAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/createAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/deleteAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/updateAuthority\",\"method\":\"PUT\"},{\"path\":\"/authority/getAuthorityList\",\"method\":\"POST\"},{\"path\":\"/authority/setDataAuthority\",\"method\":\"POST\"},{\"path\":\"/casbin/updateCasbin\",\"method\":\"POST\"},{\"path\":\"/casbin/getPolicyPathByAuthorityId\",\"method\":\"POST\"},{\"path\":\"/menu/addBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getMenu\",\"method\":\"POST\"},{\"path\":\"/menu/deleteBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/updateBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuById\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuList\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuTree\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/menu/addMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/findFile\",\"method\":\"GET\"},{\"path\":\"/fileUploadAndDownload/breakpointContinue\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/breakpointContinueFinish\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/removeChunk\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/upload\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/deleteFile\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/editFileName\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/getFileList\",\"method\":\"POST\"},{\"path\":\"/system/getServerInfo\",\"method\":\"POST\"},{\"path\":\"/system/getSystemConfig\",\"method\":\"POST\"},{\"path\":\"/system/setSystemConfig\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"PUT\"},{\"path\":\"/customer/customer\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"DELETE\"},{\"path\":\"/customer/customer\",\"method\":\"GET\"},{\"path\":\"/customer/customerList\",\"method\":\"GET\"},{\"path\":\"/autoCode/getDB\",\"method\":\"GET\"},{\"path\":\"/autoCode/getTables\",\"method\":\"GET\"},{\"path\":\"/autoCode/createTemp\",\"method\":\"POST\"},{\"path\":\"/autoCode/preview\",\"method\":\"POST\"},{\"path\":\"/autoCode/getColumn\",\"method\":\"GET\"},{\"path\":\"/autoCode/createPlug\",\"method\":\"POST\"},{\"path\":\"/autoCode/installPlugin\",\"method\":\"POST\"},{\"path\":\"/autoCode/pubPlug\",\"method\":\"POST\"},{\"path\":\"/autoCode/createPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/getPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/delPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/getMeta\",\"method\":\"POST\"},{\"path\":\"/autoCode/rollback\",\"method\":\"POST\"},{\"path\":\"/autoCode/getSysHistory\",\"method\":\"POST\"},{\"path\":\"/autoCode/delSysHistory\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/updateSysDictionaryDetail\",\"method\":\"PUT\"},{\"path\":\"/sysDictionaryDetail/createSysDictionaryDetail\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/deleteSysDictionaryDetail\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionaryDetail/findSysDictionaryDetail\",\"method\":\"GET\"},{\"path\":\"/sysDictionaryDetail/getSysDictionaryDetailList\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/createSysDictionary\",\"method\":\"POST\"},{\"path\":\"/sysDictionary/deleteSysDictionary\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionary/updateSysDictionary\",\"method\":\"PUT\"},{\"path\":\"/sysDictionary/findSysDictionary\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/getSysDictionaryList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/createSysOperationRecord\",\"method\":\"POST\"},{\"path\":\"/sysOperationRecord/findSysOperationRecord\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/getSysOperationRecordList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecord\",\"method\":\"DELETE\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecordByIds\",\"method\":\"DELETE\"},{\"path\":\"/simpleUploader/upload\",\"method\":\"POST\"},{\"path\":\"/simpleUploader/checkFileMd5\",\"method\":\"GET\"},{\"path\":\"/simpleUploader/mergeFileMd5\",\"method\":\"GET\"},{\"path\":\"/email/emailTest\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/setAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/getAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/canRemoveAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/chatGpt/getTable\",\"method\":\"POST\"},{\"path\":\"/chatGpt/createSK\",\"method\":\"POST\"},{\"path\":\"/chatGpt/getSK\",\"method\":\"GET\"},{\"path\":\"/chatGpt/deleteSK\",\"method\":\"DELETE\"},{\"path\":\"/orderInfo/findOrderInfo\",\"method\":\"GET\"},{\"path\":\"/orderInfo/getOrderInfoList\",\"method\":\"GET\"}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (17, '2023-08-24 00:01:53.562', '2023-08-24 00:01:53.562', NULL, '127.0.0.1', 'POST', '/menu/updateBaseMenu', 200, 104606157, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-23T02:45:30+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"name\":\"order-list#update\",\"desc\":\"变更\"},{\"name\":\"order-list#delete\",\"desc\":\"删除\"}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (18, '2023-08-24 00:06:27.262', '2023-08-24 00:06:27.262', NULL, '127.0.0.1', 'POST', '/menu/updateBaseMenu', 200, 32577658, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"order-list-update\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"order-list-delete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (19, '2023-08-24 00:07:12.994', '2023-08-24 00:07:12.994', NULL, '127.0.0.1', 'POST', '/menu/updateBaseMenu', 200, 84602981, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:06:27.232+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListUpdate\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListDelete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (20, '2023-08-24 00:14:21.570', '2023-08-24 00:14:21.570', NULL, '127.0.0.1', 'DELETE', '/user/deleteUser', 200, 39796425, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"id\":2}', '{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (21, '2023-08-24 00:15:04.080', '2023-08-24 00:15:04.080', NULL, '127.0.0.1', 'POST', '/user/admin_register', 200, 137371603, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":2,\"CreatedAt\":\"2023-08-23T22:44:11.539+08:00\",\"UpdatedAt\":\"2023-08-23T22:54:34.174+08:00\",\"uuid\":\"ee5745dc-27b4-4c1e-a598-2d75aa51f1c7\",\"userName\":\"lihuangzi\",\"nickName\":\"管理员\",\"sideMode\":\"dark\",\"headerImg\":\"\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"authorityId\":9528,\"authority\":{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:50.618+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"},\"authorities\":[{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:50.618+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"}],\"phone\":\"17611111111\",\"email\":\"333333333@qq.com\",\"enable\":1,\"authorityIds\":[9528],\"password\":\"123456\"}', '{\"code\":0,\"data\":{\"user\":{\"ID\":3,\"CreatedAt\":\"2023-08-24T00:15:04.034+08:00\",\"UpdatedAt\":\"2023-08-24T00:15:04.034+08:00\",\"uuid\":\"4f583cc3-92f3-4185-bd6c-139c2ba8b1b6\",\"userName\":\"lihuangzi\",\"nickName\":\"管理员\",\"sideMode\":\"dark\",\"headerImg\":\"https://qmplusimg.henrongyi.top/gva_header.jpg\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"authorityId\":9528,\"authority\":{\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"authorityId\":0,\"authorityName\":\"\",\"parentId\":null,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"\"},\"authorities\":[{\"CreatedAt\":\"2023-08-24T00:15:04.034+08:00\",\"UpdatedAt\":\"2023-08-24T00:15:04.034+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"\",\"parentId\":null,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"}],\"phone\":\"17611111111\",\"email\":\"333333333@qq.com\",\"enable\":1}},\"msg\":\"注册成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (22, '2023-08-24 00:15:21.059', '2023-08-24 00:15:21.059', NULL, '127.0.0.1', 'PUT', '/user/setUserInfo', 200, 13172913, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":3,\"CreatedAt\":\"2023-08-24T00:15:04.034+08:00\",\"UpdatedAt\":\"2023-08-24T00:15:04.034+08:00\",\"uuid\":\"4f583cc3-92f3-4185-bd6c-139c2ba8b1b6\",\"userName\":\"lihuangzi\",\"nickName\":\"管理员\",\"sideMode\":\"dark\",\"headerImg\":\"https://qmplusimg.henrongyi.top/gva_header.jpg\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"authorityId\":9528,\"authority\":{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:50.618+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"},\"authorities\":[{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:50.618+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"}],\"phone\":\"13510480001\",\"email\":\"333333333@qq.com\",\"enable\":1,\"authorityIds\":[9528]}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (23, '2023-08-24 00:15:43.791', '2023-08-24 00:15:43.791', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 12213246, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"menus\":[{\"ID\":32,\"CreatedAt\":\"2023-08-23T01:59:24+08:00\",\"UpdatedAt\":\"2023-08-23T02:45:02+08:00\",\"parentId\":\"0\",\"path\":\"order-info\",\"name\":\"order-info\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单信息\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:07:12.911+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListUpdate\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListDelete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:07:12.911+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListUpdate\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListDelete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]}],\"authorityId\":9528}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (24, '2023-08-24 00:18:37.685', '2023-08-24 00:18:37.685', NULL, '127.0.0.1', 'POST', '/casbin/updateCasbin', 200, 60005302, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":9528,\"casbinInfos\":[{\"path\":\"/jwt/jsonInBlacklist\",\"method\":\"POST\"},{\"path\":\"/user/setSelfInfo\",\"method\":\"PUT\"},{\"path\":\"/user/getUserInfo\",\"method\":\"GET\"},{\"path\":\"/user/changePassword\",\"method\":\"POST\"},{\"path\":\"/menu/getMenu\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/upload\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/deleteFile\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/editFileName\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/getFileList\",\"method\":\"POST\"},{\"path\":\"/system/getSystemConfig\",\"method\":\"POST\"},{\"path\":\"/system/setSystemConfig\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"PUT\"},{\"path\":\"/customer/customer\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"DELETE\"},{\"path\":\"/customer/customer\",\"method\":\"GET\"},{\"path\":\"/customer/customerList\",\"method\":\"GET\"},{\"path\":\"/autoCode/createTemp\",\"method\":\"POST\"},{\"path\":\"/orderInfo/createOrderInfo\",\"method\":\"POST\"},{\"path\":\"/orderInfo/deleteOrderInfo\",\"method\":\"DELETE\"},{\"path\":\"/orderInfo/deleteOrderInfoByIds\",\"method\":\"DELETE\"},{\"path\":\"/orderInfo/updateOrderInfo\",\"method\":\"PUT\"},{\"path\":\"/orderInfo/findOrderInfo\",\"method\":\"GET\"},{\"path\":\"/orderInfo/getOrderInfoList\",\"method\":\"GET\"}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (25, '2023-08-24 00:18:45.630', '2023-08-24 00:18:45.630', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 38169160, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"menus\":[{\"ID\":32,\"CreatedAt\":\"2023-08-23T01:59:24+08:00\",\"UpdatedAt\":\"2023-08-23T02:45:02+08:00\",\"parentId\":\"0\",\"path\":\"order-info\",\"name\":\"order-info\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单信息\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:07:12.911+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListUpdate\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListDelete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:07:12.911+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListUpdate\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListDelete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]}],\"authorityId\":9528}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (26, '2023-08-24 00:20:00.833', '2023-08-24 00:20:00.833', NULL, '127.0.0.1', 'PUT', '/authority/updateAuthority', 200, 74824070, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":9528,\"AuthorityName\":\"管理员\",\"parentId\":0,\"defaultRouter\":\"order-info\"}', '{\"code\":0,\"data\":{\"authority\":{\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"order-info\"}},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (27, '2023-08-24 00:20:06.523', '2023-08-24 00:20:06.523', NULL, '127.0.0.1', 'PUT', '/authority/updateAuthority', 200, 48734398, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":9528,\"AuthorityName\":\"管理员\",\"parentId\":0,\"defaultRouter\":\"order-list\"}', '{\"code\":0,\"data\":{\"authority\":{\"CreatedAt\":\"0001-01-01T00:00:00Z\",\"UpdatedAt\":\"0001-01-01T00:00:00Z\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"order-list\"}},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (28, '2023-08-24 00:20:09.235', '2023-08-24 00:20:09.235', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 30007650, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"menus\":[{\"ID\":32,\"CreatedAt\":\"2023-08-23T01:59:24+08:00\",\"UpdatedAt\":\"2023-08-23T02:45:02+08:00\",\"parentId\":\"0\",\"path\":\"order-info\",\"name\":\"order-info\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单信息\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:07:12.911+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListUpdate\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListDelete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:07:12.911+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListUpdate\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListDelete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]}],\"authorityId\":9528}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (29, '2023-08-24 00:22:30.076', '2023-08-24 00:22:30.076', NULL, '127.0.0.1', 'PUT', '/user/setUserInfo', 200, 90873551, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":1,\"CreatedAt\":\"2023-08-23T22:44:11.539+08:00\",\"UpdatedAt\":\"2023-08-23T22:55:34.491+08:00\",\"uuid\":\"a9947477-4907-45da-aafa-0a853b52e5d1\",\"userName\":\"admin\",\"nickName\":\"prince.lee\",\"sideMode\":\"dark\",\"headerImg\":\"uploads/file/d97ad29a9f5ea5d419518ee56c70e596_20230824002217.png\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"authorityId\":888,\"authority\":{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:33.205+08:00\",\"DeletedAt\":null,\"authorityId\":888,\"authorityName\":\"超级管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"},\"authorities\":[{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-23T22:50:33.205+08:00\",\"DeletedAt\":null,\"authorityId\":888,\"authorityName\":\"超级管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"dashboard\"}],\"phone\":\"13510480000\",\"email\":\"333333333@qq.com\",\"enable\":1,\"authorityIds\":[888]}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (30, '2023-08-24 00:22:49.135', '2023-08-24 00:22:49.135', NULL, '127.0.0.1', 'PUT', '/user/setUserInfo', 200, 12276485, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":3,\"CreatedAt\":\"2023-08-24T00:15:04.034+08:00\",\"UpdatedAt\":\"2023-08-24T00:15:21.057+08:00\",\"uuid\":\"4f583cc3-92f3-4185-bd6c-139c2ba8b1b6\",\"userName\":\"lihuangzi\",\"nickName\":\"管理员\",\"sideMode\":\"dark\",\"headerImg\":\"uploads/file/27fbd46f8a577320c40eac4a07dffbce_20230823225513.jpg\",\"baseColor\":\"#fff\",\"activeColor\":\"#1890ff\",\"authorityId\":9528,\"authority\":{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-24T00:20:09.206+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"order-list\"},\"authorities\":[{\"CreatedAt\":\"2023-08-23T22:44:11.288+08:00\",\"UpdatedAt\":\"2023-08-24T00:20:09.206+08:00\",\"DeletedAt\":null,\"authorityId\":9528,\"authorityName\":\"管理员\",\"parentId\":0,\"dataAuthorityId\":null,\"children\":null,\"menus\":null,\"defaultRouter\":\"order-list\"}],\"phone\":\"13510480001\",\"email\":\"333333333@qq.com\",\"enable\":1,\"authorityIds\":[9528]}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (31, '2023-08-24 00:24:24.237', '2023-08-24 00:24:24.237', NULL, '127.0.0.1', 'POST', '/casbin/updateCasbin', 200, 74766041, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":9528,\"casbinInfos\":[{\"path\":\"/jwt/jsonInBlacklist\",\"method\":\"POST\"},{\"path\":\"/user/setSelfInfo\",\"method\":\"PUT\"},{\"path\":\"/user/getUserInfo\",\"method\":\"GET\"},{\"path\":\"/user/changePassword\",\"method\":\"POST\"},{\"path\":\"/user/setUserAuthority\",\"method\":\"POST\"},{\"path\":\"/menu/getMenu\",\"method\":\"POST\"},{\"path\":\"/sysOperationRecord/getSysOperationRecordList\",\"method\":\"GET\"},{\"path\":\"/orderInfo/findOrderInfo\",\"method\":\"GET\"},{\"path\":\"/orderInfo/getOrderInfoList\",\"method\":\"GET\"}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (32, '2023-08-24 00:24:43.989', '2023-08-24 00:24:43.989', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 9741872, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"menus\":[{\"ID\":32,\"CreatedAt\":\"2023-08-23T01:59:24+08:00\",\"UpdatedAt\":\"2023-08-23T02:45:02+08:00\",\"parentId\":\"0\",\"path\":\"order-info\",\"name\":\"order-info\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单信息\",\"icon\":\"coin\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:07:12.911+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListUpdate\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListDelete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2023-08-23T02:22:51+08:00\",\"UpdatedAt\":\"2023-08-24T00:07:12.911+08:00\",\"parentId\":\"32\",\"path\":\"order-list\",\"name\":\"order-list\",\"hidden\":false,\"component\":\"view/orderInfo/orderinfo.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"订单列表\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[{\"ID\":1,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListUpdate\",\"desc\":\"变更\",\"sysBaseMenuID\":33},{\"ID\":2,\"CreatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"UpdatedAt\":\"2023-08-24T00:01:53.459+08:00\",\"name\":\"orderListDelete\",\"desc\":\"删除\",\"sysBaseMenuID\":33}]},{\"ID\":11,\"CreatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"UpdatedAt\":\"2023-08-23T22:44:11.311+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":9528}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (33, '2023-08-24 00:25:02.553', '2023-08-24 00:25:02.553', NULL, '127.0.0.1', 'POST', '/casbin/updateCasbin', 200, 41466745, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":9528,\"casbinInfos\":[{\"path\":\"/jwt/jsonInBlacklist\",\"method\":\"POST\"},{\"path\":\"/user/setSelfInfo\",\"method\":\"PUT\"},{\"path\":\"/user/getUserInfo\",\"method\":\"GET\"},{\"path\":\"/user/changePassword\",\"method\":\"POST\"},{\"path\":\"/user/setUserAuthority\",\"method\":\"POST\"},{\"path\":\"/menu/getMenu\",\"method\":\"POST\"},{\"path\":\"/sysOperationRecord/getSysOperationRecordList\",\"method\":\"GET\"},{\"path\":\"/orderInfo/findOrderInfo\",\"method\":\"GET\"},{\"path\":\"/orderInfo/getOrderInfoList\",\"method\":\"GET\"}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (34, '2023-08-24 00:25:25.857', '2023-08-24 00:25:25.857', NULL, '127.0.0.1', 'POST', '/user/setUserAuthorities', 200, 89294115, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"ID\":1,\"authorityIds\":[888,9528]}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (35, '2023-08-24 00:25:36.256', '2023-08-24 00:25:36.256', NULL, '127.0.0.1', 'POST', '/user/setUserAuthority', 200, 101134408, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":9528}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (36, '2023-08-24 00:26:13.372', '2023-08-24 00:26:13.372', NULL, '127.0.0.1', 'PUT', '/user/setSelfInfo', 200, 61267568, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"nickName\":\"prince.lee\"}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (37, '2023-08-24 00:28:35.541', '2023-08-24 00:28:35.541', NULL, '127.0.0.1', 'POST', '/user/setUserAuthority', 200, 44194973, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":888}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (38, '2023-08-24 00:29:45.179', '2023-08-24 00:29:45.179', NULL, '127.0.0.1', 'POST', '/system/getServerInfo', 200, 200584508, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '', '{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"darwin\",\"numCpu\":12,\"compiler\":\"gc\",\"goVersion\":\"go1.19.8\",\"numGoroutine\":14},\"cpu\":{\"cpus\":[63.157894733617304,5.263157899574044,52.63157894978702,9.523809519189868,50,5.00000000509317,34.999999999272404,9.999999995634425,60.00000000145519,4.999999990541255,40.000000005820766,23.809523804904153],\"cores\":6},\"ram\":{\"usedMb\":10701,\"totalMb\":16384,\"usedPercent\":65},\"disk\":{\"usedMb\":229878,\"usedGb\":224,\"totalMb\":239072,\"totalGb\":233,\"usedPercent\":96}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (39, '2023-08-24 00:29:55.570', '2023-08-24 00:29:55.570', NULL, '127.0.0.1', 'POST', '/system/getServerInfo', 200, 201288802, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '', '{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"darwin\",\"numCpu\":12,\"compiler\":\"gc\",\"goVersion\":\"go1.19.8\",\"numGoroutine\":14},\"cpu\":{\"cpus\":[4.999999990541255,0,34.99999999199645,0,0,0,4.761904759594934,0,0,0,0,0],\"cores\":6},\"ram\":{\"usedMb\":10697,\"totalMb\":16384,\"usedPercent\":65},\"disk\":{\"usedMb\":229878,\"usedGb\":224,\"totalMb\":239072,\"totalGb\":233,\"usedPercent\":96}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (40, '2023-08-24 00:30:06.011', '2023-08-24 00:30:06.011', NULL, '127.0.0.1', 'POST', '/system/getServerInfo', 200, 201051771, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '', '{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"darwin\",\"numCpu\":12,\"compiler\":\"gc\",\"goVersion\":\"go1.19.8\",\"numGoroutine\":14},\"cpu\":{\"cpus\":[47.619047616737795,0,45.00000000291038,0,29.999999994179234,0,35.00000000145519,0,28.57142857142857,0,24.99999999636202,0],\"cores\":6},\"ram\":{\"usedMb\":10718,\"totalMb\":16384,\"usedPercent\":65},\"disk\":{\"usedMb\":229875,\"usedGb\":224,\"totalMb\":239072,\"totalGb\":233,\"usedPercent\":96}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (41, '2023-08-24 00:30:33.672', '2023-08-24 00:30:33.672', NULL, '127.0.0.1', 'POST', '/user/setUserAuthority', 200, 43782124, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":9528}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (42, '2023-08-24 00:34:49.705', '2023-08-24 00:34:49.705', NULL, '127.0.0.1', 'PUT', '/user/setSelfInfo', 200, 11060335, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"sideMode\":\"dark\"}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (43, '2023-08-24 00:34:59.740', '2023-08-24 00:34:59.740', NULL, '127.0.0.1', 'PUT', '/user/setSelfInfo', 200, 80427968, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"sideMode\":\"light\"}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (44, '2023-08-24 00:35:06.132', '2023-08-24 00:35:06.132', NULL, '127.0.0.1', 'PUT', '/user/setSelfInfo', 200, 55739735, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"sideMode\":\"dark\"}', '{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (45, '2023-08-24 01:09:35.501', '2023-08-24 01:09:35.501', NULL, '127.0.0.1', 'POST', '/user/setUserAuthority', 200, 83611403, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{\"authorityId\":888}', '{\"code\":0,\"data\":{},\"msg\":\"修改成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (46, '2023-08-24 01:09:47.253', '2023-08-24 01:09:47.253', NULL, '127.0.0.1', 'GET', '/chatGpt/getSK', 200, 1860381, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '{}', '{\"code\":0,\"data\":{\"ok\":false},\"msg\":\"无sk或获取失败\"}', 1);
INSERT INTO `sys_operation_records` VALUES (47, '2023-08-24 01:09:49.775', '2023-08-24 01:09:49.775', NULL, '127.0.0.1', 'POST', '/system/getSystemConfig', 200, 996902, 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:109.0) Gecko/20100101 Firefox/116.0', '', '', '{\"code\":0,\"data\":{\"config\":{\"jwt\":{\"signing-key\":\"252959d4-0e4f-436c-ab71-4fe3e1809bef\",\"expires-time\":\"7d\",\"buffer-time\":\"1d\",\"issuer\":\"qmPlus\"},\"zap\":{\"level\":\"info\",\"prefix\":\"[github.com/flipped-aurora/gin-vue-admin/server]\",\"format\":\"console\",\"director\":\"log\",\"encode-level\":\"LowercaseColorLevelEncoder\",\"stacktrace-key\":\"stacktrace\",\"max-age\":0,\"show-line\":true,\"log-in-console\":true},\"redis\":{\"addr\":\"127.0.0.1:6379\",\"password\":\"\",\"db\":0},\"email\":{\"to\":\"xxx@qq.com\",\"from\":\"xxx@163.com\",\"host\":\"smtp.163.com\",\"secret\":\"xxx\",\"nickname\":\"test\",\"port\":465,\"is-ssl\":true},\"system\":{\"env\":\"public\",\"db-type\":\"mysql\",\"oss-type\":\"local\",\"router-prefix\":\"\",\"addr\":8888,\"iplimit-count\":15000,\"iplimit-time\":3600,\"use-multipoint\":false,\"use-redis\":false},\"captcha\":{\"key-long\":6,\"img-width\":240,\"img-height\":80,\"open-captcha\":0,\"open-captcha-timeout\":3600},\"autocode\":{\"server-model\":\"/model/%s\",\"server-router\":\"/router/%s\",\"server\":\"/server\",\"server-api\":\"/api/v1/%s\",\"server-plug\":\"/plugin/%s\",\"server-initialize\":\"/initialize\",\"root\":\"/Users/leeprince/票务处项目/ticket_office/admin\",\"web-table\":\"/view\",\"web\":\"/web/src\",\"server-service\":\"/service/%s\",\"server-request\":\"/model/%s/request/\",\"web-api\":\"/api\",\"web-form\":\"/view\",\"transfer-restart\":true},\"mysql\":{\"prefix\":\"\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"db-name\":\"gva\",\"username\":\"root\",\"password\":\"leeprince\",\"path\":\"127.0.0.1\",\"engine\":\"\",\"log-mode\":\"error\",\"max-idle-conns\":10,\"max-open-conns\":100,\"singular\":false,\"log-zap\":false},\"mssql\":{\"prefix\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"path\":\"\",\"engine\":\"\",\"log-mode\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"singular\":false,\"log-zap\":false},\"pgsql\":{\"prefix\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"path\":\"\",\"engine\":\"\",\"log-mode\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"singular\":false,\"log-zap\":false},\"oracle\":{\"prefix\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"path\":\"\",\"engine\":\"\",\"log-mode\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"singular\":false,\"log-zap\":false},\"sqlite\":{\"prefix\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"path\":\"\",\"engine\":\"\",\"log-mode\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"singular\":false,\"log-zap\":false},\"db-list\":[{\"type\":\"mysql\",\"alias-name\":\"票务服务\",\"prefix\":\"\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"db-name\":\"ticket\",\"username\":\"root\",\"password\":\"leeprince\",\"path\":\"127.0.0.1\",\"engine\":\"\",\"log-mode\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"singular\":false,\"log-zap\":false,\"disable\":false}],\"local\":{\"path\":\"uploads/file\",\"store-path\":\"uploads/file\"},\"qiniu\":{\"zone\":\"ZoneHuaDong\",\"bucket\":\"\",\"img-path\":\"\",\"access-key\":\"\",\"secret-key\":\"\",\"use-https\":false,\"use-cdn-domains\":false},\"aliyun-oss\":{\"endpoint\":\"yourEndpoint\",\"access-key-id\":\"yourAccessKeyId\",\"access-key-secret\":\"yourAccessKeySecret\",\"bucket-name\":\"yourBucketName\",\"bucket-url\":\"yourBucketUrl\",\"base-path\":\"yourBasePath\"},\"hua-wei-obs\":{\"path\":\"you-path\",\"bucket\":\"you-bucket\",\"endpoint\":\"you-endpoint\",\"access-key\":\"you-access-key\",\"secret-key\":\"you-secret-key\"},\"tencent-cos\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"aws-s3\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"endpoint\":\"\",\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\",\"s3-force-path-style\":false,\"disable-ssl\":false},\"excel\":{\"dir\":\"./resource/excel/\"},\"timer\":{\"spec\":\"@daily\",\"detail\":[{\"tableName\":\"sys_operation_records\",\"compareField\":\"created_at\",\"interval\":\"2160h\"},{\"tableName\":\"jwt_blacklists\",\"compareField\":\"created_at\",\"interval\":\"168h\"}],\"start\":true,\"with_seconds\":false},\"cors\":{\"mode\":\"strict-whitelist\",\"whitelist\":[{\"allow-origin\":\"example1.com\",\"allow-methods\":\"POST, GET\",\"allow-headers\":\"Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true},{\"allow-origin\":\"example2.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true}]}}},\"msg\":\"获取成功\"}', 1);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
BEGIN;
INSERT INTO `sys_user_authority` VALUES (1, 888);
INSERT INTO `sys_user_authority` VALUES (1, 9528);
INSERT INTO `sys_user_authority` VALUES (3, 9528);
COMMIT;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) COLLATE utf8mb4_general_ci DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) COLLATE utf8mb4_general_ci DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) COLLATE utf8mb4_general_ci DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) COLLATE utf8mb4_general_ci DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` bigint(20) unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint(20) DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
BEGIN;
INSERT INTO `sys_users` VALUES (1, '2023-08-23 22:44:11.539', '2023-08-24 01:09:35.419', NULL, 'a9947477-4907-45da-aafa-0a853b52e5d1', 'admin', '$2a$10$ZFXsarXxNx0QZf2hfcsn4uZIrVCgIMRixL7kZdF8kLmn9JFxJJlUO', 'prince.lee', 'dark', 'uploads/file/d97ad29a9f5ea5d419518ee56c70e596_20230824002217.png', '#fff', '#1890ff', 888, '13510480000', '333333333@qq.com', 1);
INSERT INTO `sys_users` VALUES (2, '2023-08-23 22:44:11.539', '2023-08-23 22:54:34.174', '2023-08-24 00:14:21.530', 'ee5745dc-27b4-4c1e-a598-2d75aa51f1c7', 'a303176530', '$2a$10$eIC1OD3MYt1AEoa05t4TIeGuGs9a8wrQp82fVIly9Hzre/L9OlVSC', '管理员', 'dark', 'https:///qmplusimg.henrongyi.top/1572075907logo.png', '#fff', '#1890ff', 9528, '17611111111', '333333333@qq.com', 1);
INSERT INTO `sys_users` VALUES (3, '2023-08-24 00:15:04.034', '2023-08-24 00:22:49.134', NULL, '4f583cc3-92f3-4185-bd6c-139c2ba8b1b6', 'lihuangzi', '$2a$10$a5qf0Ptc6TiYdTYfS/yrKOxrbQ7NdoLlk7n40koLjDeSNgvFmRU5G', '管理员', 'dark', 'uploads/file/27fbd46f8a577320c40eac4a07dffbce_20230823225513.jpg', '#fff', '#1890ff', 9528, '13510480001', '333333333@qq.com', 1);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
