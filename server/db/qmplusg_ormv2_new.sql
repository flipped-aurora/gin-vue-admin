/*
 Navicat Premium Data Transfer

 Source Server         : pandora
 Source Server Type    : MySQL
 Source Server Version : 50640
 Source Host           : localhost:3306
 Source Schema         : qmplus

 Target Server Type    : MySQL
 Target Server Version : 50640
 File Encoding         : 65001

 Date: 31/08/2020 23:32:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/login', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/register', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/createApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiById', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/updateApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getAllApis', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/createAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/deleteAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/getAuthorityList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/setDataAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/updateAuthority', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/copyAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuTree', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addMenuAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/deleteBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/updateBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuById', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/changePassword', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/deleteUser', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/upload', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/getFileList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/deleteFile', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/updateCasbin', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/casbinTest/:pathParam', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/jwt/jsonInBlacklist', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/getSystemConfig', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/setSystemConfig', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customer', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/customerList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/createTemp', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getTables', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getDB', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getColume', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/createSysDictionary', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/updateSysDictionary', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/findSysDictionary', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysDictionary/getSysDictionaryList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/updateSysOperationRecord', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserInfo', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/base/login', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/base/register', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/createApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiById', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/deleteApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/updateApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getAllApis', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/createAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/deleteAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/getAuthorityList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/setDataAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuTree', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addMenuAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/deleteBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/updateBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuById', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/changePassword', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/getUserList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/setUserAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/upload', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/getFileList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/updateCasbin', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/jwt/jsonInBlacklist', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/getSystemConfig', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/setSystemConfig', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customer', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/customerList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/base/login', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/base/register', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/createApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiById', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/deleteApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/updateApi', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getAllApis', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/createAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/deleteAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/getAuthorityList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/setDataAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuTree', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addMenuAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/deleteBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/updateBaseMenu', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuById', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/changePassword', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/getUserList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/setUserAuthority', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/upload', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/getFileList', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/updateCasbin', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/jwt/jsonInBlacklist', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/getSystemConfig', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/setSystemConfig', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'POST', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'PUT', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'DELETE', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customer', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/customerList', 'GET', NULL, NULL, NULL);
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/autoCode/createTemp', 'POST', NULL, NULL, NULL);

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'客户名\'',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'客户手机号\'',
  `sys_user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '\'管理ID\'',
  `sys_user_authority_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'管理角色ID\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_customers_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------
INSERT INTO `exa_customers` VALUES (1, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '测试客户', '1761111111', 1, '888');

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `exa_file_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `file_chunk_number` bigint(20) NULL DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_chunks_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'文件名\'',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'文件地址\'',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'文件标签\'',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'编号\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_upload_and_downloads_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO `exa_file_upload_and_downloads` VALUES (1, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '10.png', 'http://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (2, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `chunk_total` bigint(20) NULL DEFAULT NULL,
  `is_finish` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_files_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for exa_simple_uploaders
-- ----------------------------
DROP TABLE IF EXISTS `exa_simple_uploaders`;
CREATE TABLE `exa_simple_uploaders`  (
  `chunk_number` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'当前切片标记\'',
  `current_chunk_size` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'当前切片容量\'',
  `current_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'切片本地路径\'',
  `total_size` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'总容量\'',
  `identifier` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'文件标识（md5）\'',
  `filename` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'文件名\'',
  `total_chunks` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'切片总数\'',
  `is_done` tinyint(1) NULL DEFAULT NULL COMMENT '\'是否上传完成\'',
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'文件本地路径\''
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '\'jwt\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'api路径\'',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'api中文描述\'',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'api组\'',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT 'POST',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 66 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/base/login', '用户登录', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/base/register', '用户注册', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (3, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (4, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (6, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (7, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (8, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (9, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/authority/createAuthority', '创建角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/authority/deleteAuthority', '删除角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/authority/getAuthorityList', '获取角色列表', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/menu/getMenu', '获取菜单树', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (13, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/menu/getMenuList', '分页获取基础menu列表', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/menu/addBaseMenu', '新增菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/menu/getMenuAuthority', '获取指定角色menu', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/menu/deleteBaseMenu', '删除菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (19, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/menu/updateBaseMenu', '更新菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (20, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/menu/getBaseMenuById', '根据id获取菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (21, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/user/changePassword', '修改密码', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/user/getUserList', '获取用户列表', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (24, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/user/setUserAuthority', '修改用户角色', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/fileUploadAndDownload/upload', '文件上传示例', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES (26, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES (27, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (28, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (29, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/authority/setDataAuthority', '设置角色资源权限', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (32, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/system/getSystemConfig', '获取配置文件内容', 'system', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/system/setSystemConfig', '设置配置文件内容', 'system', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/customer/customer', '创建客户', 'customer', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/customer/customer', '更新客户', 'customer', 'PUT');
INSERT INTO `sys_apis` VALUES (36, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/customer/customer', '删除客户', 'customer', 'DELETE');
INSERT INTO `sys_apis` VALUES (37, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/customer/customer', '获取单一客户', 'customer', 'GET');
INSERT INTO `sys_apis` VALUES (38, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/customer/customerList', '获取客户列表', 'customer', 'GET');
INSERT INTO `sys_apis` VALUES (39, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/casbin/casbinTest/:pathParam', 'RESTFUL模式测试', 'casbin', 'GET');
INSERT INTO `sys_apis` VALUES (40, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/autoCode/createTemp', '自动化代码', 'autoCode', 'POST');
INSERT INTO `sys_apis` VALUES (41, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/authority/updateAuthority', '更新角色信息', 'authority', 'PUT');
INSERT INTO `sys_apis` VALUES (42, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/authority/copyAuthority', '拷贝角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (43, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/user/deleteUser', '删除用户', 'user', 'DELETE');
INSERT INTO `sys_apis` VALUES (44, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', 'sysDictionaryDetail', 'POST');
INSERT INTO `sys_apis` VALUES (45, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', 'sysDictionaryDetail', 'DELETE');
INSERT INTO `sys_apis` VALUES (46, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', 'sysDictionaryDetail', 'PUT');
INSERT INTO `sys_apis` VALUES (47, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', 'sysDictionaryDetail', 'GET');
INSERT INTO `sys_apis` VALUES (48, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', 'sysDictionaryDetail', 'GET');
INSERT INTO `sys_apis` VALUES (49, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionary/createSysDictionary', '新增字典', 'sysDictionary', 'POST');
INSERT INTO `sys_apis` VALUES (50, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', 'sysDictionary', 'DELETE');
INSERT INTO `sys_apis` VALUES (51, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionary/updateSysDictionary', '更新字典', 'sysDictionary', 'PUT');
INSERT INTO `sys_apis` VALUES (52, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', 'sysDictionary', 'GET');
INSERT INTO `sys_apis` VALUES (53, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', 'sysDictionary', 'GET');
INSERT INTO `sys_apis` VALUES (54, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', 'sysOperationRecord', 'POST');
INSERT INTO `sys_apis` VALUES (55, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', 'sysOperationRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES (56, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', 'sysOperationRecord', 'GET');
INSERT INTO `sys_apis` VALUES (57, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', 'sysOperationRecord', 'GET');
INSERT INTO `sys_apis` VALUES (58, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/autoCode/getTables', '获取数据库表', 'autoCode', 'GET');
INSERT INTO `sys_apis` VALUES (59, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/autoCode/getDB', '获取所有数据库', 'autoCode', 'GET');
INSERT INTO `sys_apis` VALUES (60, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/autoCode/getColume', '获取所选table的所有字段', 'autoCode', 'GET');
INSERT INTO `sys_apis` VALUES (61, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', 'sysOperationRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES (62, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/simpleUploader/upload', '插件版分片上传', 'simpleUploader', 'POST');
INSERT INTO `sys_apis` VALUES (63, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', 'simpleUploader', 'GET');
INSERT INTO `sys_apis` VALUES (64, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', 'simpleUploader', 'GET');
INSERT INTO `sys_apis` VALUES (65, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '/user/setUserInfo', '设置用户信息', 'user', 'PUT');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '\'角色ID\'',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'角色名\'',
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'父角色ID\'',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE INDEX `authority_id`(`authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '888', '普通用户', '0');
INSERT INTO `sys_authorities` VALUES ('2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '8881', '普通用户子角色', '888');
INSERT INTO `sys_authorities` VALUES ('2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '9528', '测试角色', '0');

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_base_menu_id` bigint(20) UNSIGNED NOT NULL DEFAULT 0,
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '\'角色ID\'',
  PRIMARY KEY (`sys_base_menu_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES (1, '888');
INSERT INTO `sys_authority_menus` VALUES (1, '8881');
INSERT INTO `sys_authority_menus` VALUES (1, '9528');
INSERT INTO `sys_authority_menus` VALUES (2, '888');
INSERT INTO `sys_authority_menus` VALUES (2, '8881');
INSERT INTO `sys_authority_menus` VALUES (2, '9528');
INSERT INTO `sys_authority_menus` VALUES (3, '888');
INSERT INTO `sys_authority_menus` VALUES (3, '9528');
INSERT INTO `sys_authority_menus` VALUES (4, '888');
INSERT INTO `sys_authority_menus` VALUES (4, '9528');
INSERT INTO `sys_authority_menus` VALUES (5, '888');
INSERT INTO `sys_authority_menus` VALUES (5, '9528');
INSERT INTO `sys_authority_menus` VALUES (6, '888');
INSERT INTO `sys_authority_menus` VALUES (6, '9528');
INSERT INTO `sys_authority_menus` VALUES (7, '888');
INSERT INTO `sys_authority_menus` VALUES (7, '9528');
INSERT INTO `sys_authority_menus` VALUES (8, '888');
INSERT INTO `sys_authority_menus` VALUES (8, '8881');
INSERT INTO `sys_authority_menus` VALUES (8, '9528');
INSERT INTO `sys_authority_menus` VALUES (9, '888');
INSERT INTO `sys_authority_menus` VALUES (9, '9528');
INSERT INTO `sys_authority_menus` VALUES (10, '888');
INSERT INTO `sys_authority_menus` VALUES (10, '9528');
INSERT INTO `sys_authority_menus` VALUES (11, '888');
INSERT INTO `sys_authority_menus` VALUES (11, '9528');
INSERT INTO `sys_authority_menus` VALUES (12, '888');
INSERT INTO `sys_authority_menus` VALUES (12, '9528');
INSERT INTO `sys_authority_menus` VALUES (13, '888');
INSERT INTO `sys_authority_menus` VALUES (13, '9528');
INSERT INTO `sys_authority_menus` VALUES (14, '888');
INSERT INTO `sys_authority_menus` VALUES (14, '9528');
INSERT INTO `sys_authority_menus` VALUES (15, '888');
INSERT INTO `sys_authority_menus` VALUES (15, '9528');
INSERT INTO `sys_authority_menus` VALUES (16, '888');
INSERT INTO `sys_authority_menus` VALUES (17, '888');
INSERT INTO `sys_authority_menus` VALUES (17, '8881');
INSERT INTO `sys_authority_menus` VALUES (17, '9528');
INSERT INTO `sys_authority_menus` VALUES (18, '888');
INSERT INTO `sys_authority_menus` VALUES (18, '8881');
INSERT INTO `sys_authority_menus` VALUES (18, '9528');
INSERT INTO `sys_authority_menus` VALUES (19, '888');
INSERT INTO `sys_authority_menus` VALUES (19, '8881');
INSERT INTO `sys_authority_menus` VALUES (19, '9528');
INSERT INTO `sys_authority_menus` VALUES (20, '888');
INSERT INTO `sys_authority_menus` VALUES (20, '8881');
INSERT INTO `sys_authority_menus` VALUES (20, '9528');
INSERT INTO `sys_authority_menus` VALUES (21, '888');
INSERT INTO `sys_authority_menus` VALUES (22, '888');
INSERT INTO `sys_authority_menus` VALUES (23, '888');
INSERT INTO `sys_authority_menus` VALUES (24, '888');
INSERT INTO `sys_authority_menus` VALUES (25, '888');

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `sys_base_menu_id` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `menu_level` bigint(20) UNSIGNED NULL DEFAULT NULL,
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'父菜单ID\'',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'路由path\'',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'路由name\'',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '\'是否在列表隐藏\'',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'对应前端文件路径\'',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '\'排序标记\'',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '\'附加属性\'',
  `default_menu` tinyint(1) NULL DEFAULT NULL COMMENT '\'附加属性\'',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'附加属性\'',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'附加属性\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 26 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, 0, 0, '仪表盘', 'setting');
INSERT INTO `sys_base_menus` VALUES (2, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '0', 'about', 'about', 0, 'view/about/index.vue', 7, 0, 0, '关于我们', 'info');
INSERT INTO `sys_base_menus` VALUES (3, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, 0, 0, '超级管理员', 'user-solid');
INSERT INTO `sys_base_menus` VALUES (4, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, 0, 0, '角色管理', 's-custom');
INSERT INTO `sys_base_menus` VALUES (5, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, 1, 0, '菜单管理', 's-order');
INSERT INTO `sys_base_menus` VALUES (6, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, 1, 0, 'api管理', 's-platform');
INSERT INTO `sys_base_menus` VALUES (7, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, 0, 0, '用户管理', 'coordinate');
INSERT INTO `sys_base_menus` VALUES (8, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 4, 0, 0, '个人信息', 'message-solid');
INSERT INTO `sys_base_menus` VALUES (9, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 6, 0, 0, '示例文件', 's-management');
INSERT INTO `sys_base_menus` VALUES (10, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '9', 'table', 'table', 0, 'view/example/table/table.vue', 1, 0, 0, '表格示例', 's-order');
INSERT INTO `sys_base_menus` VALUES (11, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '9', 'form', 'form', 0, 'view/example/form/form.vue', 2, 0, 0, '表单示例', 'document');
INSERT INTO `sys_base_menus` VALUES (12, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '9', 'rte', 'rte', 0, 'view/example/rte/rte.vue', 3, 0, 0, '富文本编辑器', 'reading');
INSERT INTO `sys_base_menus` VALUES (13, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '9', 'excel', 'excel', 0, 'view/example/excel/excel.vue', 4, 0, 0, 'excel导入导出', 's-marketing');
INSERT INTO `sys_base_menus` VALUES (14, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '9', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, 0, 0, '上传下载', 'upload');
INSERT INTO `sys_base_menus` VALUES (15, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '9', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, 0, 0, '断点续传', 'upload');
INSERT INTO `sys_base_menus` VALUES (16, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '9', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, 0, 0, '客户列表（资源示例）', 's-custom');
INSERT INTO `sys_base_menus` VALUES (17, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, 0, 0, '系统工具', 's-cooperation');
INSERT INTO `sys_base_menus` VALUES (18, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '17', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, 1, 0, '代码生成器', 'cpu');
INSERT INTO `sys_base_menus` VALUES (19, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '17', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, 1, 0, '表单生成器', 'magic-stick');
INSERT INTO `sys_base_menus` VALUES (20, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '17', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, 0, 0, '系统配置', 's-operation');
INSERT INTO `sys_base_menus` VALUES (21, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '0', 'iconList', 'iconList', 0, 'view/iconList/index.vue', 2, 0, 0, '图标集合', 'star-on');
INSERT INTO `sys_base_menus` VALUES (22, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, 0, 0, '字典管理', 'notebook-2');
INSERT INTO `sys_base_menus` VALUES (23, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 0, 0, '字典详情', 's-order');
INSERT INTO `sys_base_menus` VALUES (24, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, 0, 0, '操作历史', 'time');
INSERT INTO `sys_base_menus` VALUES (25, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 0, '9', 'simpleUploader', 'simpleUploader', 0, 'view/example/simpleUploader/simpleUploader', 6, 0, 0, '断点续传（插件版）', 'upload');

-- ----------------------------
-- Table structure for sys_data_authority_ids
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_ids`;
CREATE TABLE `sys_data_authority_ids`  (
  `sys_authority_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '\'角色ID\'',
  `data_authority_id_authority_id` varchar(90) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL COMMENT '\'角色ID\'',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_data_authority_ids
-- ----------------------------
INSERT INTO `sys_data_authority_ids` VALUES ('888', '888');
INSERT INTO `sys_data_authority_ids` VALUES ('888', '8881');
INSERT INTO `sys_data_authority_ids` VALUES ('888', '9528');
INSERT INTO `sys_data_authority_ids` VALUES ('9528', '8881');
INSERT INTO `sys_data_authority_ids` VALUES ('9528', '9528');

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'字典名（中）\'',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'字典名（英）\'',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '\'状态\'',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'描述\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '性别', 'sex', 1, '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'展示值\'',
  `value` bigint(20) NULL DEFAULT NULL COMMENT '\'字典值\'',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '\'启用状态\'',
  `sort` bigint(20) NULL DEFAULT NULL COMMENT '\'排序标记\'',
  `sys_dictionary_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '\'关联标记\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 24 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'smallint', 1, 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (2, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'mediumint', 2, 1, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (3, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'int', 3, 1, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'bigint', 4, 1, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'data', 0, 1, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (6, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'time', 1, 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (7, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'year', 2, 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (8, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'datetime', 3, 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (9, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'timestamp', 5, 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (10, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'float', 0, 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (11, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'double', 1, 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (12, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'decimal', 2, 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (13, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'char', 0, 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (14, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'varchar', 1, 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (15, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'tinyblob', 2, 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (16, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'tinytext', 3, 1, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (17, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'text', 4, 1, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (18, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'blob', 5, 1, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (19, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'mediumblob', 6, 1, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (20, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'mediumtext', 7, 1, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (21, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'longblob', 8, 1, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (22, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'longtext', 9, 1, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'tinyint', 0, 1, 0, 6);

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'请求ip\'',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'\'',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'\'',
  `status` bigint(20) NULL DEFAULT NULL COMMENT '\'\'',
  `latency` bigint(20) NULL DEFAULT NULL COMMENT '\'\'',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'\'',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'\'',
  `body` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'请求Body\'',
  `resp` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '\'响应Body\'',
  `user_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '\'\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'用户UUID\'',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'用户登录名\'',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'用户登录密码\'',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '系统用户' COMMENT '\'用户昵称\'',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT 'http://qmplusimg.henrongyi.top/head.png' COMMENT '\'用户头像\'',
  `authority_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT '888' COMMENT '\'用户角色ID\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, 'e8cfdf92-c7c4-40e2-bc70-2da2bf7df0a3', 'admin', 'e10adc3949ba59abbe56e057f20f883e', '超级管理员', 'http://qmplusimg.henrongyi.top/1571627762timg.jpg', '888');
INSERT INTO `sys_users` VALUES (2, '2020-08-31 23:31:07', '2020-08-31 23:31:07', NULL, '64b55645-7d14-45d7-9e1f-491bd8b8fde5', 'a303176530', '3ec063004a6f31642261936a379fde3d', 'QMPlusUser', 'http://qmplusimg.henrongyi.top/1572075907logo.png', '9528');

-- ----------------------------
-- Table structure for sys_workflow_step_infos
-- ----------------------------
DROP TABLE IF EXISTS `sys_workflow_step_infos`;
CREATE TABLE `sys_workflow_step_infos`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `sys_workflow_id` bigint(20) UNSIGNED NULL DEFAULT NULL COMMENT '\'所属工作流ID\'',
  `is_strat` tinyint(1) NULL DEFAULT NULL COMMENT '\'是否是开始流节点\'',
  `step_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'工作流节点名称\'',
  `step_no` double NULL DEFAULT NULL COMMENT '\'步骤id （第几步）\'',
  `step_authority_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'操作者级别id\'',
  `is_end` tinyint(1) NULL DEFAULT NULL COMMENT '\'是否是完结流节点\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_workflow_step_infos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sys_workflows
-- ----------------------------
DROP TABLE IF EXISTS `sys_workflows`;
CREATE TABLE `sys_workflows`  (
  `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `workflow_nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'工作流中文名称\'',
  `workflow_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'工作流英文名称\'',
  `workflow_description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '\'工作流描述\'',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_workflows_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

-- ----------------------------
-- View structure for authority_menu
-- ----------------------------
DROP VIEW IF EXISTS `authority_menu`;
CREATE ALGORITHM = UNDEFINED SQL SECURITY DEFINER VIEW `authority_menu` AS select `sys_base_menus`.`id` AS `id`,`sys_base_menus`.`created_at` AS `created_at`,`sys_base_menus`.`updated_at` AS `updated_at`,`sys_base_menus`.`deleted_at` AS `deleted_at`,`sys_base_menus`.`menu_level` AS `menu_level`,`sys_base_menus`.`parent_id` AS `parent_id`,`sys_base_menus`.`path` AS `path`,`sys_base_menus`.`name` AS `name`,`sys_base_menus`.`hidden` AS `hidden`,`sys_base_menus`.`component` AS `component`,`sys_base_menus`.`title` AS `title`,`sys_base_menus`.`icon` AS `icon`,`sys_base_menus`.`sort` AS `sort`,`sys_authority_menus`.`sys_authority_authority_id` AS `authority_id`,`sys_authority_menus`.`sys_base_menu_id` AS `menu_id`,`sys_base_menus`.`keep_alive` AS `keep_alive`,`sys_base_menus`.`default_menu` AS `default_menu` from (`sys_authority_menus` join `sys_base_menus` on((`sys_authority_menus`.`sys_base_menu_id` = `sys_base_menus`.`id`)));

SET FOREIGN_KEY_CHECKS = 1;
