/*
 Navicat Premium Data Transfer

 Source Server         : 8.134.131.67
 Source Server Type    : MySQL
 Source Server Version : 101102 (10.11.2-MariaDB-1:10.11.2+maria~ubu2204)
 Source Host           : 8.134.131.67:3306
 Source Schema         : clothing

 Target Server Type    : MySQL
 Target Server Version : 101102 (10.11.2-MariaDB-1:10.11.2+maria~ubu2204)
 File Encoding         : 65001

 Date: 11/05/2023 21:57:00
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for app_jwt_blacklist
-- ----------------------------
DROP TABLE IF EXISTS `app_jwt_blacklist`;
CREATE TABLE `app_jwt_blacklist`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_app_jwt_blacklist_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_jwt_blacklist
-- ----------------------------

-- ----------------------------
-- Table structure for app_role
-- ----------------------------
DROP TABLE IF EXISTS `app_role`;
CREATE TABLE `app_role`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_app_role_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_role
-- ----------------------------
INSERT INTO `app_role` VALUES (1, '2023-05-04 11:10:02.539', '2023-05-04 11:10:02.539', NULL, '老板', 1, 0, 0);
INSERT INTO `app_role` VALUES (2, '2023-05-04 11:10:08.171', '2023-05-04 11:10:08.171', NULL, '裁缝', 1, 0, 0);
INSERT INTO `app_role` VALUES (3, '2023-05-04 11:10:12.834', '2023-05-04 11:10:12.834', NULL, '组长', 1, 0, 0);
INSERT INTO `app_role` VALUES (4, '2023-05-04 11:10:17.205', '2023-05-04 11:10:17.205', NULL, '车工', 1, 0, 0);

-- ----------------------------
-- Table structure for app_user
-- ----------------------------
DROP TABLE IF EXISTS `app_user`;
CREATE TABLE `app_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `nickname` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `wages` decimal(10, 2) NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `phone_num` varchar(11) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `open_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `union_id` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_app_user_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of app_user
-- ----------------------------
INSERT INTO `app_user` VALUES (1, '2023-05-04 18:32:26.163', '2023-05-04 18:32:26.163', NULL, 'lph', '$2a$10$/H9ShAKz.CVeZazugunWfeajmO0z7v19cPxehjIIwSuPPh7NbwyMm', '', 0.00, NULL, '15625538114', '', '', 0, 0, 0);
INSERT INTO `app_user` VALUES (2, '2023-05-08 10:39:37.435', '2023-05-08 10:39:37.435', NULL, 'lph2', '$2a$10$NxFGW69wc/eWte3SvadAd.1/BSKszku.erbYcWMEvpKPRl2k4WnIq', '', 0.00, NULL, '15625560423', '', '', 0, 0, 0);
INSERT INTO `app_user` VALUES (3, '2023-05-08 14:29:10.528', '2023-05-08 14:29:10.528', NULL, 'lph4', '$2a$10$fX5P2ew7WAV16no3a6/zFeYSdvncf1PmosBtvZok/f2uPpjsV35la', '', 0.00, NULL, '15625560425', '', '', 0, 0, 0);
INSERT INTO `app_user` VALUES (4, '2023-05-08 14:48:40.626', '2023-05-08 14:48:40.626', NULL, 'lph3', '$2a$10$GtN8DLVZy6WarrISyZOCm.dM61uRanMlnj5G6sBncCYks0fdJ1/B2', '', 0.00, NULL, '15625560424', '', '', 0, 0, 0);

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `content` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `media_url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sort` bigint UNSIGNED NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `type` bigint UNSIGNED NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_banner_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of banner
-- ----------------------------

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v6` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `v7` varchar(25) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `idx_casbin_rule`(`ptype` ASC, `v0` ASC, `v1` ASC, `v2` ASC, `v3` ASC, `v4` ASC, `v5` ASC, `v6` ASC, `v7` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 295 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES (185, 'p', '888', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (186, 'p', '888', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (191, 'p', '888', '/api/deleteApisByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (189, 'p', '888', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (190, 'p', '888', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (188, 'p', '888', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (187, 'p', '888', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (271, 'p', '888', '/appRole/createAppRole', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (272, 'p', '888', '/appRole/deleteAppRole', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (273, 'p', '888', '/appRole/deleteAppRoleByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (275, 'p', '888', '/appRole/findAppRole', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (276, 'p', '888', '/appRole/getAppRoleList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (274, 'p', '888', '/appRole/updateAppRole', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (283, 'p', '888', '/appUser/createAppUser', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (284, 'p', '888', '/appUser/deleteAppUser', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (285, 'p', '888', '/appUser/deleteAppUserByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (287, 'p', '888', '/appUser/findAppUser', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (288, 'p', '888', '/appUser/getAppUserList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (286, 'p', '888', '/appUser/updateAppUser', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (192, 'p', '888', '/authority/copyAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (193, 'p', '888', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (194, 'p', '888', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (196, 'p', '888', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (197, 'p', '888', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (195, 'p', '888', '/authority/updateAuthority', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (260, 'p', '888', '/authorityBtn/canRemoveAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (259, 'p', '888', '/authorityBtn/getAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (258, 'p', '888', '/authorityBtn/setAuthorityBtn', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (232, 'p', '888', '/autoCode/createPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (230, 'p', '888', '/autoCode/createPlug', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (227, 'p', '888', '/autoCode/createTemp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (234, 'p', '888', '/autoCode/delPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (238, 'p', '888', '/autoCode/delSysHistory', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (229, 'p', '888', '/autoCode/getColumn', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (225, 'p', '888', '/autoCode/getDB', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (235, 'p', '888', '/autoCode/getMeta', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (233, 'p', '888', '/autoCode/getPackage', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (237, 'p', '888', '/autoCode/getSysHistory', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (226, 'p', '888', '/autoCode/getTables', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (231, 'p', '888', '/autoCode/installPlugin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (228, 'p', '888', '/autoCode/preview', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (236, 'p', '888', '/autoCode/rollback', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (173, 'p', '888', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (199, 'p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (198, 'p', '888', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (262, 'p', '888', '/chatGpt/createSK', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (264, 'p', '888', '/chatGpt/deleteSK', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (263, 'p', '888', '/chatGpt/getSK', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (261, 'p', '888', '/chatGpt/getTable', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (265, 'p', '888', '/company/createCompany', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (266, 'p', '888', '/company/deleteCompany', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (267, 'p', '888', '/company/deleteCompanyByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (269, 'p', '888', '/company/findCompany', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (270, 'p', '888', '/company/getCompanyList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (268, 'p', '888', '/company/updateCompany', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (222, 'p', '888', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (223, 'p', '888', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (221, 'p', '888', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (220, 'p', '888', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (224, 'p', '888', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (257, 'p', '888', '/email/emailTest', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (210, 'p', '888', '/fileUploadAndDownload/breakpointContinue', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (211, 'p', '888', '/fileUploadAndDownload/breakpointContinueFinish', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (214, 'p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (215, 'p', '888', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (209, 'p', '888', '/fileUploadAndDownload/findFile', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (216, 'p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (212, 'p', '888', '/fileUploadAndDownload/removeChunk', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (213, 'p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (174, 'p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (200, 'p', '888', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (208, 'p', '888', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (202, 'p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (204, 'p', '888', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (206, 'p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (201, 'p', '888', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (207, 'p', '888', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (205, 'p', '888', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (203, 'p', '888', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (255, 'p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (256, 'p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (254, 'p', '888', '/simpleUploader/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (244, 'p', '888', '/sysDictionary/createSysDictionary', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (245, 'p', '888', '/sysDictionary/deleteSysDictionary', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (247, 'p', '888', '/sysDictionary/findSysDictionary', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (248, 'p', '888', '/sysDictionary/getSysDictionaryList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (246, 'p', '888', '/sysDictionary/updateSysDictionary', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (240, 'p', '888', '/sysDictionaryDetail/createSysDictionaryDetail', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (241, 'p', '888', '/sysDictionaryDetail/deleteSysDictionaryDetail', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (242, 'p', '888', '/sysDictionaryDetail/findSysDictionaryDetail', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (243, 'p', '888', '/sysDictionaryDetail/getSysDictionaryDetailList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (239, 'p', '888', '/sysDictionaryDetail/updateSysDictionaryDetail', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (249, 'p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (252, 'p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (253, 'p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (250, 'p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (251, 'p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (217, 'p', '888', '/system/getServerInfo', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (218, 'p', '888', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (219, 'p', '888', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (277, 'p', '888', '/team/createTeam', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (278, 'p', '888', '/team/deleteTeam', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (279, 'p', '888', '/team/deleteTeamByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (281, 'p', '888', '/team/findTeam', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (282, 'p', '888', '/team/getTeamList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (280, 'p', '888', '/team/updateTeam', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (176, 'p', '888', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (182, 'p', '888', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (175, 'p', '888', '/user/deleteUser', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (180, 'p', '888', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (177, 'p', '888', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (184, 'p', '888', '/user/resetPassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (179, 'p', '888', '/user/setSelfInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (181, 'p', '888', '/user/setUserAuthorities', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (183, 'p', '888', '/user/setUserAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (178, 'p', '888', '/user/setUserInfo', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (289, 'p', '888', '/userRole/createUserRole', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (290, 'p', '888', '/userRole/deleteUserRole', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (291, 'p', '888', '/userRole/deleteUserRoleByIds', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (293, 'p', '888', '/userRole/findUserRole', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (294, 'p', '888', '/userRole/getUserRoleList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (292, 'p', '888', '/userRole/updateUserRole', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (96, 'p', '8881', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (99, 'p', '8881', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (101, 'p', '8881', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (98, 'p', '8881', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (97, 'p', '8881', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (100, 'p', '8881', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (102, 'p', '8881', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (103, 'p', '8881', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (104, 'p', '8881', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (105, 'p', '8881', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (94, 'p', '8881', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (123, 'p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (122, 'p', '8881', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (129, 'p', '8881', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (130, 'p', '8881', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (127, 'p', '8881', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (128, 'p', '8881', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (131, 'p', '8881', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (120, 'p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (121, 'p', '8881', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (119, 'p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (118, 'p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (124, 'p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (108, 'p', '8881', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (110, 'p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (112, 'p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (114, 'p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (109, 'p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (106, 'p', '8881', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (111, 'p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (107, 'p', '8881', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (113, 'p', '8881', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (125, 'p', '8881', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (126, 'p', '8881', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (95, 'p', '8881', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (115, 'p', '8881', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (132, 'p', '8881', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (116, 'p', '8881', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (117, 'p', '8881', '/user/setUserAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (135, 'p', '9528', '/api/createApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (138, 'p', '9528', '/api/deleteApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (140, 'p', '9528', '/api/getAllApis', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (137, 'p', '9528', '/api/getApiById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (136, 'p', '9528', '/api/getApiList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (139, 'p', '9528', '/api/updateApi', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (141, 'p', '9528', '/authority/createAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (142, 'p', '9528', '/authority/deleteAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (143, 'p', '9528', '/authority/getAuthorityList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (144, 'p', '9528', '/authority/setDataAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (171, 'p', '9528', '/autoCode/createTemp', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (133, 'p', '9528', '/base/login', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (162, 'p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (161, 'p', '9528', '/casbin/updateCasbin', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (169, 'p', '9528', '/customer/customer', 'DELETE', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (167, 'p', '9528', '/customer/customer', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (168, 'p', '9528', '/customer/customer', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (166, 'p', '9528', '/customer/customer', 'PUT', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (170, 'p', '9528', '/customer/customerList', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (159, 'p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (160, 'p', '9528', '/fileUploadAndDownload/editFileName', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (158, 'p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (157, 'p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (163, 'p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (147, 'p', '9528', '/menu/addBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (149, 'p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (151, 'p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (153, 'p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (148, 'p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (145, 'p', '9528', '/menu/getMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (150, 'p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (146, 'p', '9528', '/menu/getMenuList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (152, 'p', '9528', '/menu/updateBaseMenu', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (164, 'p', '9528', '/system/getSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (165, 'p', '9528', '/system/setSystemConfig', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (134, 'p', '9528', '/user/admin_register', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (154, 'p', '9528', '/user/changePassword', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (172, 'p', '9528', '/user/getUserInfo', 'GET', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (155, 'p', '9528', '/user/getUserList', 'POST', '', '', '', '', '');
INSERT INTO `casbin_rule` VALUES (156, 'p', '9528', '/user/setUserAuthority', 'POST', '', '', '', '', '');

-- ----------------------------
-- Table structure for cloth
-- ----------------------------
DROP TABLE IF EXISTS `cloth`;
CREATE TABLE `cloth`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `company_id` bigint UNSIGNED NULL DEFAULT NULL,
  `style_id` bigint UNSIGNED NULL DEFAULT NULL,
  `color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `length` decimal(10, 2) NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cloth_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cloth
-- ----------------------------

-- ----------------------------
-- Table structure for company
-- ----------------------------
DROP TABLE IF EXISTS `company`;
CREATE TABLE `company`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_company_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of company
-- ----------------------------
INSERT INTO `company` VALUES (1, '2023-05-08 10:37:12.593', '2023-05-08 10:37:12.593', NULL, 1, 'test', 1, 0, 0, 0);

-- ----------------------------
-- Table structure for company_apply
-- ----------------------------
DROP TABLE IF EXISTS `company_apply`;
CREATE TABLE `company_apply`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `company_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `role_id` bigint UNSIGNED NULL DEFAULT NULL,
  `remark` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  `status` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_company_apply_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of company_apply
-- ----------------------------
INSERT INTO `company_apply` VALUES (1, '2023-05-08 14:49:51.826', '2023-05-09 09:14:08.979', NULL, 1, 2, 2, '', 0, 0, 0, 1);
INSERT INTO `company_apply` VALUES (2, '2023-05-08 14:49:58.469', '2023-05-09 09:15:11.096', NULL, 1, 4, 3, 'test', 0, 0, 0, 1);

-- ----------------------------
-- Table structure for cropping_record
-- ----------------------------
DROP TABLE IF EXISTS `cropping_record`;
CREATE TABLE `cropping_record`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `company_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `style_id` bigint UNSIGNED NULL DEFAULT NULL,
  `color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `length` decimal(10, 2) NULL DEFAULT NULL,
  `size` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `quantity` bigint UNSIGNED NULL DEFAULT NULL,
  `step` bigint UNSIGNED NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_cropping_record_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of cropping_record
-- ----------------------------

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `customer_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_customers_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `exa_file_id` bigint UNSIGNED NULL DEFAULT NULL,
  `file_chunk_number` bigint NULL DEFAULT NULL,
  `file_chunk_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_chunks_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of exa_file_chunks
-- ----------------------------

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_upload_and_downloads_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO `exa_file_upload_and_downloads` VALUES (1, '2023-05-02 17:43:06.934', '2023-05-02 17:43:06.934', NULL, '10.png', 'https://qmplusimg.henrongyi.top/gvalogo.png', 'png', '158787308910.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (2, '2023-05-02 17:43:06.934', '2023-05-02 17:43:06.934', NULL, 'logo.png', 'https://qmplusimg.henrongyi.top/1576554439myAvatar.png', 'png', '1587973709logo.png');

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `file_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_md5` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `file_path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `chunk_total` bigint NULL DEFAULT NULL,
  `is_finish` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_files_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of exa_files
-- ----------------------------

-- ----------------------------
-- Table structure for job
-- ----------------------------
DROP TABLE IF EXISTS `job`;
CREATE TABLE `job`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `cropping_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `process_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `price` decimal(10, 2) NULL DEFAULT NULL,
  `quantity` bigint UNSIGNED NULL DEFAULT NULL,
  `income` decimal(10, 2) NULL DEFAULT NULL,
  `real_quantity` bigint UNSIGNED NULL DEFAULT NULL,
  `real_income` decimal(10, 2) NULL DEFAULT NULL,
  `step` bigint UNSIGNED NULL DEFAULT NULL,
  `job_type` bigint UNSIGNED NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_job_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of job
-- ----------------------------

-- ----------------------------
-- Table structure for job_question
-- ----------------------------
DROP TABLE IF EXISTS `job_question`;
CREATE TABLE `job_question`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `job_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `content` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_job_question_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of job_question
-- ----------------------------

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT 'jwt',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES (2, '2023-05-04 13:36:48.574', '2023-05-04 13:36:48.574', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiNDY4NDgyMjYtODMwZS00ZmRkLTgwMDQtODdiMjY4YzMwNDVmIiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6Ik1yLuWlh-a3vCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJxbVBsdXMiLCJhdWQiOlsiR1ZBIl0sImV4cCI6MTY4Mzc4MjEyNiwibmJmIjoxNjgzMTc3MzI2fQ.03BWXABvdL2zjT96KIViGDO7cs3p53KXokywC-8Ra6M');

-- ----------------------------
-- Table structure for msg_box
-- ----------------------------
DROP TABLE IF EXISTS `msg_box`;
CREATE TABLE `msg_box`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `msg_type` bigint UNSIGNED NULL DEFAULT NULL,
  `msg_id` bigint UNSIGNED NULL DEFAULT NULL,
  `status` tinyint(1) NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  `from_user` bigint UNSIGNED NULL DEFAULT NULL,
  `to_user` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_msg_box_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of msg_box
-- ----------------------------
INSERT INTO `msg_box` VALUES (1, '2023-05-08 14:49:51.872', '2023-05-08 16:00:01.406', NULL, 1, 1, 1, 0, 0, 0, 2, 1);
INSERT INTO `msg_box` VALUES (2, '2023-05-08 14:49:58.521', '2023-05-08 14:49:58.521', NULL, 1, 2, 0, 0, 0, 0, 4, 1);
INSERT INTO `msg_box` VALUES (3, '2023-05-09 16:16:15.534', '2023-05-09 16:16:15.534', NULL, 2, 1, 0, 0, 0, 0, 3, 4);
INSERT INTO `msg_box` VALUES (4, '2023-05-09 16:16:17.878', '2023-05-09 16:16:17.878', NULL, 2, 2, 0, 0, 0, 0, 3, 4);

-- ----------------------------
-- Table structure for process
-- ----------------------------
DROP TABLE IF EXISTS `process`;
CREATE TABLE `process`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `style_id` bigint UNSIGNED NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `price` decimal(10, 2) NULL DEFAULT NULL,
  `percent` decimal(10, 2) NULL DEFAULT NULL,
  `sort` bigint UNSIGNED NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_process_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of process
-- ----------------------------

-- ----------------------------
-- Table structure for style
-- ----------------------------
DROP TABLE IF EXISTS `style`;
CREATE TABLE `style`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `company_id` bigint UNSIGNED NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `price` decimal(10, 2) NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_style_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of style
-- ----------------------------

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 220 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/base/login', '用户登录(必选)', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单(退出，必选)', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (3, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/deleteUser', '删除用户', '系统用户', 'DELETE');
INSERT INTO `sys_apis` VALUES (4, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/admin_register', '用户注册', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/getUserList', '获取用户列表', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (6, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/setUserInfo', '设置用户信息', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (7, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/setSelfInfo', '设置自身信息(必选)', '系统用户', 'PUT');
INSERT INTO `sys_apis` VALUES (8, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/getUserInfo', '获取自身信息(必选)', '系统用户', 'GET');
INSERT INTO `sys_apis` VALUES (9, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/setUserAuthorities', '设置权限组', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/changePassword', '修改密码（建议选择)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/setUserAuthority', '修改用户角色(必选)', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/user/resetPassword', '重置用户密码', '系统用户', 'POST');
INSERT INTO `sys_apis` VALUES (13, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (19, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/api/deleteApisByIds', '批量删除api', 'api', 'DELETE');
INSERT INTO `sys_apis` VALUES (20, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/authority/copyAuthority', '拷贝角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (21, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/authority/createAuthority', '创建角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (22, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/authority/deleteAuthority', '删除角色', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/authority/updateAuthority', '更新角色信息', '角色', 'PUT');
INSERT INTO `sys_apis` VALUES (24, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/authority/getAuthorityList', '获取角色列表', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/authority/setDataAuthority', '设置角色资源权限', '角色', 'POST');
INSERT INTO `sys_apis` VALUES (26, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (27, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (28, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/menu/addBaseMenu', '新增菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (29, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/menu/getMenu', '获取菜单树(必选)', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/menu/deleteBaseMenu', '删除菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/menu/updateBaseMenu', '更新菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (32, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/menu/getBaseMenuById', '根据id获取菜单', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/menu/getMenuList', '分页获取基础menu列表', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/menu/getBaseMenuTree', '获取用户动态路由', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/menu/getMenuAuthority', '获取指定角色menu', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (36, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', '菜单', 'POST');
INSERT INTO `sys_apis` VALUES (37, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/fileUploadAndDownload/findFile', '寻找目标文件（秒传）', '分片上传', 'GET');
INSERT INTO `sys_apis` VALUES (38, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/fileUploadAndDownload/breakpointContinue', '断点续传', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (39, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/fileUploadAndDownload/breakpointContinueFinish', '断点续传完成', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (40, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/fileUploadAndDownload/removeChunk', '上传完成移除文件', '分片上传', 'POST');
INSERT INTO `sys_apis` VALUES (41, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/fileUploadAndDownload/upload', '文件上传示例', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (42, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/fileUploadAndDownload/deleteFile', '删除文件', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (43, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/fileUploadAndDownload/editFileName', '文件名或者备注编辑', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (44, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', '文件上传与下载', 'POST');
INSERT INTO `sys_apis` VALUES (45, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/system/getServerInfo', '获取服务器信息', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (46, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/system/getSystemConfig', '获取配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (47, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/system/setSystemConfig', '设置配置文件内容', '系统服务', 'POST');
INSERT INTO `sys_apis` VALUES (48, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/customer/customer', '更新客户', '客户', 'PUT');
INSERT INTO `sys_apis` VALUES (49, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/customer/customer', '创建客户', '客户', 'POST');
INSERT INTO `sys_apis` VALUES (50, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/customer/customer', '删除客户', '客户', 'DELETE');
INSERT INTO `sys_apis` VALUES (51, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/customer/customer', '获取单一客户', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (52, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/customer/customerList', '获取客户列表', '客户', 'GET');
INSERT INTO `sys_apis` VALUES (53, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/getDB', '获取所有数据库', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (54, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/getTables', '获取数据库表', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (55, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/createTemp', '自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (56, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/preview', '预览自动化代码', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (57, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/getColumn', '获取所选table的所有字段', '代码生成器', 'GET');
INSERT INTO `sys_apis` VALUES (58, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/createPlug', '自动创建插件包', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (59, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/installPlugin', '安装插件', '代码生成器', 'POST');
INSERT INTO `sys_apis` VALUES (60, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/createPackage', '生成包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES (61, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/getPackage', '获取所有包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES (62, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/delPackage', '删除包(package)', '包（pkg）生成器', 'POST');
INSERT INTO `sys_apis` VALUES (63, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/getMeta', '获取meta信息', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (64, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/rollback', '回滚自动生成代码', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (65, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/getSysHistory', '查询回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (66, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/autoCode/delSysHistory', '删除回滚记录', '代码生成器历史', 'POST');
INSERT INTO `sys_apis` VALUES (67, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionaryDetail/updateSysDictionaryDetail', '更新字典内容', '系统字典详情', 'PUT');
INSERT INTO `sys_apis` VALUES (68, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionaryDetail/createSysDictionaryDetail', '新增字典内容', '系统字典详情', 'POST');
INSERT INTO `sys_apis` VALUES (69, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionaryDetail/deleteSysDictionaryDetail', '删除字典内容', '系统字典详情', 'DELETE');
INSERT INTO `sys_apis` VALUES (70, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionaryDetail/findSysDictionaryDetail', '根据ID获取字典内容', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (71, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionaryDetail/getSysDictionaryDetailList', '获取字典内容列表', '系统字典详情', 'GET');
INSERT INTO `sys_apis` VALUES (72, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionary/createSysDictionary', '新增字典', '系统字典', 'POST');
INSERT INTO `sys_apis` VALUES (73, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionary/deleteSysDictionary', '删除字典', '系统字典', 'DELETE');
INSERT INTO `sys_apis` VALUES (74, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionary/updateSysDictionary', '更新字典', '系统字典', 'PUT');
INSERT INTO `sys_apis` VALUES (75, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionary/findSysDictionary', '根据ID获取字典', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (76, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysDictionary/getSysDictionaryList', '获取字典列表', '系统字典', 'GET');
INSERT INTO `sys_apis` VALUES (77, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', '操作记录', 'POST');
INSERT INTO `sys_apis` VALUES (78, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (79, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', '操作记录', 'GET');
INSERT INTO `sys_apis` VALUES (80, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (81, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', '操作记录', 'DELETE');
INSERT INTO `sys_apis` VALUES (82, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/simpleUploader/upload', '插件版分片上传', '断点续传(插件版)', 'POST');
INSERT INTO `sys_apis` VALUES (83, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (84, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', '断点续传(插件版)', 'GET');
INSERT INTO `sys_apis` VALUES (85, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/email/emailTest', '发送测试邮件', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (86, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/email/emailSend', '发送邮件示例', 'email', 'POST');
INSERT INTO `sys_apis` VALUES (87, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/authorityBtn/setAuthorityBtn', '设置按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (88, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/authorityBtn/getAuthorityBtn', '获取已有按钮权限', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (89, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/authorityBtn/canRemoveAuthorityBtn', '删除按钮', '按钮权限', 'POST');
INSERT INTO `sys_apis` VALUES (90, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/chatGpt/getTable', '通过gpt获取内容', '万用表格', 'POST');
INSERT INTO `sys_apis` VALUES (91, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/chatGpt/createSK', '录入sk', '万用表格', 'POST');
INSERT INTO `sys_apis` VALUES (92, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/chatGpt/getSK', '获取sk', '万用表格', 'GET');
INSERT INTO `sys_apis` VALUES (93, '2023-05-02 17:42:55.896', '2023-05-02 17:42:55.896', NULL, '/chatGpt/deleteSK', '删除sk', '万用表格', 'DELETE');
INSERT INTO `sys_apis` VALUES (94, '2023-05-02 20:29:59.702', '2023-05-02 20:29:59.702', '2023-05-02 20:53:41.727', '/appUser/createAppUser', '新增app用户', 'appUser', 'POST');
INSERT INTO `sys_apis` VALUES (95, '2023-05-02 20:29:59.866', '2023-05-02 20:29:59.866', '2023-05-02 20:53:41.727', '/appUser/deleteAppUser', '删除app用户', 'appUser', 'DELETE');
INSERT INTO `sys_apis` VALUES (96, '2023-05-02 20:30:00.029', '2023-05-02 20:30:00.029', '2023-05-02 20:53:41.727', '/appUser/deleteAppUserByIds', '批量删除app用户', 'appUser', 'DELETE');
INSERT INTO `sys_apis` VALUES (97, '2023-05-02 20:30:00.194', '2023-05-02 20:30:00.194', '2023-05-02 20:53:41.727', '/appUser/updateAppUser', '更新app用户', 'appUser', 'PUT');
INSERT INTO `sys_apis` VALUES (98, '2023-05-02 20:30:00.357', '2023-05-02 20:30:00.357', '2023-05-02 20:53:41.727', '/appUser/findAppUser', '根据ID获取app用户', 'appUser', 'GET');
INSERT INTO `sys_apis` VALUES (99, '2023-05-02 20:30:00.522', '2023-05-02 20:30:00.522', '2023-05-02 20:53:41.727', '/appUser/getAppUserList', '获取app用户列表', 'appUser', 'GET');
INSERT INTO `sys_apis` VALUES (100, '2023-05-02 20:34:43.960', '2023-05-02 20:34:43.960', NULL, '/company/createCompany', '新增公司', 'company', 'POST');
INSERT INTO `sys_apis` VALUES (101, '2023-05-02 20:34:44.123', '2023-05-02 20:34:44.123', NULL, '/company/deleteCompany', '删除公司', 'company', 'DELETE');
INSERT INTO `sys_apis` VALUES (102, '2023-05-02 20:34:44.286', '2023-05-02 20:34:44.286', NULL, '/company/deleteCompanyByIds', '批量删除公司', 'company', 'DELETE');
INSERT INTO `sys_apis` VALUES (103, '2023-05-02 20:34:44.449', '2023-05-02 20:34:44.449', NULL, '/company/updateCompany', '更新公司', 'company', 'PUT');
INSERT INTO `sys_apis` VALUES (104, '2023-05-02 20:34:44.611', '2023-05-02 20:34:44.611', NULL, '/company/findCompany', '根据ID获取公司', 'company', 'GET');
INSERT INTO `sys_apis` VALUES (105, '2023-05-02 20:34:44.774', '2023-05-02 20:34:44.774', NULL, '/company/getCompanyList', '获取公司列表', 'company', 'GET');
INSERT INTO `sys_apis` VALUES (106, '2023-05-02 20:36:43.537', '2023-05-02 20:36:43.537', NULL, '/appRole/createAppRole', '新增app角色', 'appRole', 'POST');
INSERT INTO `sys_apis` VALUES (107, '2023-05-02 20:36:43.701', '2023-05-02 20:36:43.701', NULL, '/appRole/deleteAppRole', '删除app角色', 'appRole', 'DELETE');
INSERT INTO `sys_apis` VALUES (108, '2023-05-02 20:36:43.865', '2023-05-02 20:36:43.865', NULL, '/appRole/deleteAppRoleByIds', '批量删除app角色', 'appRole', 'DELETE');
INSERT INTO `sys_apis` VALUES (109, '2023-05-02 20:36:44.029', '2023-05-02 20:36:44.029', NULL, '/appRole/updateAppRole', '更新app角色', 'appRole', 'PUT');
INSERT INTO `sys_apis` VALUES (110, '2023-05-02 20:36:44.192', '2023-05-02 20:36:44.192', NULL, '/appRole/findAppRole', '根据ID获取app角色', 'appRole', 'GET');
INSERT INTO `sys_apis` VALUES (111, '2023-05-02 20:36:44.356', '2023-05-02 20:36:44.356', NULL, '/appRole/getAppRoleList', '获取app角色列表', 'appRole', 'GET');
INSERT INTO `sys_apis` VALUES (112, '2023-05-02 20:49:38.859', '2023-05-02 20:49:38.859', NULL, '/team/createTeam', '新增组', 'team', 'POST');
INSERT INTO `sys_apis` VALUES (113, '2023-05-02 20:49:39.021', '2023-05-02 20:49:39.021', NULL, '/team/deleteTeam', '删除组', 'team', 'DELETE');
INSERT INTO `sys_apis` VALUES (114, '2023-05-02 20:49:39.195', '2023-05-02 20:49:39.195', NULL, '/team/deleteTeamByIds', '批量删除组', 'team', 'DELETE');
INSERT INTO `sys_apis` VALUES (115, '2023-05-02 20:49:39.358', '2023-05-02 20:49:39.358', NULL, '/team/updateTeam', '更新组', 'team', 'PUT');
INSERT INTO `sys_apis` VALUES (116, '2023-05-02 20:49:39.521', '2023-05-02 20:49:39.521', NULL, '/team/findTeam', '根据ID获取组', 'team', 'GET');
INSERT INTO `sys_apis` VALUES (117, '2023-05-02 20:49:39.683', '2023-05-02 20:49:39.683', NULL, '/team/getTeamList', '获取组列表', 'team', 'GET');
INSERT INTO `sys_apis` VALUES (118, '2023-05-02 20:57:35.271', '2023-05-02 20:57:35.271', NULL, '/appUser/createAppUser', '新增app用户', 'appUser', 'POST');
INSERT INTO `sys_apis` VALUES (119, '2023-05-02 20:57:35.430', '2023-05-02 20:57:35.430', NULL, '/appUser/deleteAppUser', '删除app用户', 'appUser', 'DELETE');
INSERT INTO `sys_apis` VALUES (120, '2023-05-02 20:57:35.589', '2023-05-02 20:57:35.589', NULL, '/appUser/deleteAppUserByIds', '批量删除app用户', 'appUser', 'DELETE');
INSERT INTO `sys_apis` VALUES (121, '2023-05-02 20:57:35.748', '2023-05-02 20:57:35.748', NULL, '/appUser/updateAppUser', '更新app用户', 'appUser', 'PUT');
INSERT INTO `sys_apis` VALUES (122, '2023-05-02 20:57:35.906', '2023-05-02 20:57:35.906', NULL, '/appUser/findAppUser', '根据ID获取app用户', 'appUser', 'GET');
INSERT INTO `sys_apis` VALUES (123, '2023-05-02 20:57:36.065', '2023-05-02 20:57:36.065', NULL, '/appUser/getAppUserList', '获取app用户列表', 'appUser', 'GET');
INSERT INTO `sys_apis` VALUES (124, '2023-05-02 21:08:32.418', '2023-05-02 21:08:32.418', NULL, '/userRole/createUserRole', '新增用户角色', 'userRole', 'POST');
INSERT INTO `sys_apis` VALUES (125, '2023-05-02 21:08:32.582', '2023-05-02 21:08:32.582', NULL, '/userRole/deleteUserRole', '删除用户角色', 'userRole', 'DELETE');
INSERT INTO `sys_apis` VALUES (126, '2023-05-02 21:08:32.748', '2023-05-02 21:08:32.748', NULL, '/userRole/deleteUserRoleByIds', '批量删除用户角色', 'userRole', 'DELETE');
INSERT INTO `sys_apis` VALUES (127, '2023-05-02 21:08:32.912', '2023-05-02 21:08:32.912', NULL, '/userRole/updateUserRole', '更新用户角色', 'userRole', 'PUT');
INSERT INTO `sys_apis` VALUES (128, '2023-05-02 21:08:33.076', '2023-05-02 21:08:33.076', NULL, '/userRole/findUserRole', '根据ID获取用户角色', 'userRole', 'GET');
INSERT INTO `sys_apis` VALUES (129, '2023-05-02 21:08:33.240', '2023-05-02 21:08:33.240', NULL, '/userRole/getUserRoleList', '获取用户角色列表', 'userRole', 'GET');
INSERT INTO `sys_apis` VALUES (130, '2023-05-02 21:26:23.277', '2023-05-02 21:26:23.277', NULL, '/userWallet/createUserWallet', '新增用户钱包', 'userWallet', 'POST');
INSERT INTO `sys_apis` VALUES (131, '2023-05-02 21:26:23.440', '2023-05-02 21:26:23.440', NULL, '/userWallet/deleteUserWallet', '删除用户钱包', 'userWallet', 'DELETE');
INSERT INTO `sys_apis` VALUES (132, '2023-05-02 21:26:23.626', '2023-05-02 21:26:23.626', NULL, '/userWallet/deleteUserWalletByIds', '批量删除用户钱包', 'userWallet', 'DELETE');
INSERT INTO `sys_apis` VALUES (133, '2023-05-02 21:26:23.789', '2023-05-02 21:26:23.789', NULL, '/userWallet/updateUserWallet', '更新用户钱包', 'userWallet', 'PUT');
INSERT INTO `sys_apis` VALUES (134, '2023-05-02 21:26:23.952', '2023-05-02 21:26:23.952', NULL, '/userWallet/findUserWallet', '根据ID获取用户钱包', 'userWallet', 'GET');
INSERT INTO `sys_apis` VALUES (135, '2023-05-02 21:26:24.115', '2023-05-02 21:26:24.115', NULL, '/userWallet/getUserWalletList', '获取用户钱包列表', 'userWallet', 'GET');
INSERT INTO `sys_apis` VALUES (136, '2023-05-03 09:21:48.439', '2023-05-03 09:21:48.439', NULL, '/teamUser/createTeamUser', '新增组员', 'teamUser', 'POST');
INSERT INTO `sys_apis` VALUES (137, '2023-05-03 09:21:48.655', '2023-05-03 09:21:48.655', NULL, '/teamUser/deleteTeamUser', '删除组员', 'teamUser', 'DELETE');
INSERT INTO `sys_apis` VALUES (138, '2023-05-03 09:21:49.229', '2023-05-03 09:21:49.229', NULL, '/teamUser/deleteTeamUserByIds', '批量删除组员', 'teamUser', 'DELETE');
INSERT INTO `sys_apis` VALUES (139, '2023-05-03 09:21:49.449', '2023-05-03 09:21:49.449', NULL, '/teamUser/updateTeamUser', '更新组员', 'teamUser', 'PUT');
INSERT INTO `sys_apis` VALUES (140, '2023-05-03 09:21:49.669', '2023-05-03 09:21:49.669', NULL, '/teamUser/findTeamUser', '根据ID获取组员', 'teamUser', 'GET');
INSERT INTO `sys_apis` VALUES (141, '2023-05-03 09:21:49.883', '2023-05-03 09:21:49.883', NULL, '/teamUser/getTeamUserList', '获取组员列表', 'teamUser', 'GET');
INSERT INTO `sys_apis` VALUES (142, '2023-05-03 10:02:19.495', '2023-05-03 10:02:19.495', '2023-05-03 11:05:48.717', '/cloth/createCloth', '新增布料入库单', 'cloth', 'POST');
INSERT INTO `sys_apis` VALUES (143, '2023-05-03 10:02:19.719', '2023-05-03 10:02:19.719', '2023-05-03 11:05:48.717', '/cloth/deleteCloth', '删除布料入库单', 'cloth', 'DELETE');
INSERT INTO `sys_apis` VALUES (144, '2023-05-03 10:02:19.941', '2023-05-03 10:02:19.941', '2023-05-03 11:05:48.717', '/cloth/deleteClothByIds', '批量删除布料入库单', 'cloth', 'DELETE');
INSERT INTO `sys_apis` VALUES (145, '2023-05-03 10:02:20.156', '2023-05-03 10:02:20.156', '2023-05-03 11:05:48.717', '/cloth/updateCloth', '更新布料入库单', 'cloth', 'PUT');
INSERT INTO `sys_apis` VALUES (146, '2023-05-03 10:02:20.371', '2023-05-03 10:02:20.371', '2023-05-03 11:05:48.717', '/cloth/findCloth', '根据ID获取布料入库单', 'cloth', 'GET');
INSERT INTO `sys_apis` VALUES (147, '2023-05-03 10:02:20.587', '2023-05-03 10:02:20.587', '2023-05-03 11:05:48.717', '/cloth/getClothList', '获取布料入库单列表', 'cloth', 'GET');
INSERT INTO `sys_apis` VALUES (148, '2023-05-03 11:07:32.957', '2023-05-03 11:07:32.957', NULL, '/cloth/createCloth', '新增布料入库单', 'cloth', 'POST');
INSERT INTO `sys_apis` VALUES (149, '2023-05-03 11:07:33.177', '2023-05-03 11:07:33.177', NULL, '/cloth/deleteCloth', '删除布料入库单', 'cloth', 'DELETE');
INSERT INTO `sys_apis` VALUES (150, '2023-05-03 11:07:33.393', '2023-05-03 11:07:33.393', NULL, '/cloth/deleteClothByIds', '批量删除布料入库单', 'cloth', 'DELETE');
INSERT INTO `sys_apis` VALUES (151, '2023-05-03 11:07:33.606', '2023-05-03 11:07:33.606', NULL, '/cloth/updateCloth', '更新布料入库单', 'cloth', 'PUT');
INSERT INTO `sys_apis` VALUES (152, '2023-05-03 11:07:33.855', '2023-05-03 11:07:33.855', NULL, '/cloth/findCloth', '根据ID获取布料入库单', 'cloth', 'GET');
INSERT INTO `sys_apis` VALUES (153, '2023-05-03 11:07:34.104', '2023-05-03 11:07:34.104', NULL, '/cloth/getClothList', '获取布料入库单列表', 'cloth', 'GET');
INSERT INTO `sys_apis` VALUES (154, '2023-05-03 11:12:34.102', '2023-05-03 11:12:34.102', NULL, '/style/createStyle', '新增款式', 'style', 'POST');
INSERT INTO `sys_apis` VALUES (155, '2023-05-03 11:12:34.300', '2023-05-03 11:12:34.300', NULL, '/style/deleteStyle', '删除款式', 'style', 'DELETE');
INSERT INTO `sys_apis` VALUES (156, '2023-05-03 11:12:34.516', '2023-05-03 11:12:34.516', NULL, '/style/deleteStyleByIds', '批量删除款式', 'style', 'DELETE');
INSERT INTO `sys_apis` VALUES (157, '2023-05-03 11:12:34.738', '2023-05-03 11:12:34.738', NULL, '/style/updateStyle', '更新款式', 'style', 'PUT');
INSERT INTO `sys_apis` VALUES (158, '2023-05-03 11:12:35.315', '2023-05-03 11:12:35.315', NULL, '/style/findStyle', '根据ID获取款式', 'style', 'GET');
INSERT INTO `sys_apis` VALUES (159, '2023-05-03 11:12:35.822', '2023-05-03 11:12:35.822', NULL, '/style/getStyleList', '获取款式列表', 'style', 'GET');
INSERT INTO `sys_apis` VALUES (160, '2023-05-03 11:17:42.768', '2023-05-03 11:17:42.768', NULL, '/process/createProcess', '新增工序', 'process', 'POST');
INSERT INTO `sys_apis` VALUES (161, '2023-05-03 11:17:42.977', '2023-05-03 11:17:42.977', NULL, '/process/deleteProcess', '删除工序', 'process', 'DELETE');
INSERT INTO `sys_apis` VALUES (162, '2023-05-03 11:17:43.190', '2023-05-03 11:17:43.190', NULL, '/process/deleteProcessByIds', '批量删除工序', 'process', 'DELETE');
INSERT INTO `sys_apis` VALUES (163, '2023-05-03 11:17:43.410', '2023-05-03 11:17:43.410', NULL, '/process/updateProcess', '更新工序', 'process', 'PUT');
INSERT INTO `sys_apis` VALUES (164, '2023-05-03 11:17:43.624', '2023-05-03 11:17:43.624', NULL, '/process/findProcess', '根据ID获取工序', 'process', 'GET');
INSERT INTO `sys_apis` VALUES (165, '2023-05-03 11:17:43.846', '2023-05-03 11:17:43.846', NULL, '/process/getProcessList', '获取工序列表', 'process', 'GET');
INSERT INTO `sys_apis` VALUES (166, '2023-05-03 15:16:45.452', '2023-05-03 15:16:45.452', NULL, '/croppingRecord/createCroppingRecord', '新增裁剪单', 'croppingRecord', 'POST');
INSERT INTO `sys_apis` VALUES (167, '2023-05-03 15:16:45.674', '2023-05-03 15:16:45.674', NULL, '/croppingRecord/deleteCroppingRecord', '删除裁剪单', 'croppingRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES (168, '2023-05-03 15:16:45.891', '2023-05-03 15:16:45.891', NULL, '/croppingRecord/deleteCroppingRecordByIds', '批量删除裁剪单', 'croppingRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES (169, '2023-05-03 15:16:46.107', '2023-05-03 15:16:46.107', NULL, '/croppingRecord/updateCroppingRecord', '更新裁剪单', 'croppingRecord', 'PUT');
INSERT INTO `sys_apis` VALUES (170, '2023-05-03 15:16:46.327', '2023-05-03 15:16:46.327', NULL, '/croppingRecord/findCroppingRecord', '根据ID获取裁剪单', 'croppingRecord', 'GET');
INSERT INTO `sys_apis` VALUES (171, '2023-05-03 15:16:46.543', '2023-05-03 15:16:46.543', NULL, '/croppingRecord/getCroppingRecordList', '获取裁剪单列表', 'croppingRecord', 'GET');
INSERT INTO `sys_apis` VALUES (172, '2023-05-03 15:58:05.802', '2023-05-03 15:58:05.802', NULL, '/job/createJob', '新增工单', 'job', 'POST');
INSERT INTO `sys_apis` VALUES (173, '2023-05-03 15:58:05.996', '2023-05-03 15:58:05.996', NULL, '/job/deleteJob', '删除工单', 'job', 'DELETE');
INSERT INTO `sys_apis` VALUES (174, '2023-05-03 15:58:06.216', '2023-05-03 15:58:06.216', NULL, '/job/deleteJobByIds', '批量删除工单', 'job', 'DELETE');
INSERT INTO `sys_apis` VALUES (175, '2023-05-03 15:58:06.430', '2023-05-03 15:58:06.430', NULL, '/job/updateJob', '更新工单', 'job', 'PUT');
INSERT INTO `sys_apis` VALUES (176, '2023-05-03 15:58:06.654', '2023-05-03 15:58:06.654', NULL, '/job/findJob', '根据ID获取工单', 'job', 'GET');
INSERT INTO `sys_apis` VALUES (177, '2023-05-03 15:58:06.876', '2023-05-03 15:58:06.876', NULL, '/job/getJobList', '获取工单列表', 'job', 'GET');
INSERT INTO `sys_apis` VALUES (178, '2023-05-04 09:23:01.550', '2023-05-04 09:23:01.550', NULL, '/jobQuestion/createJobQuestion', '新增工单问题', 'jobQuestion', 'POST');
INSERT INTO `sys_apis` VALUES (179, '2023-05-04 09:23:01.764', '2023-05-04 09:23:01.764', NULL, '/jobQuestion/deleteJobQuestion', '删除工单问题', 'jobQuestion', 'DELETE');
INSERT INTO `sys_apis` VALUES (180, '2023-05-04 09:23:01.986', '2023-05-04 09:23:01.986', NULL, '/jobQuestion/deleteJobQuestionByIds', '批量删除工单问题', 'jobQuestion', 'DELETE');
INSERT INTO `sys_apis` VALUES (181, '2023-05-04 09:23:02.251', '2023-05-04 09:23:02.251', NULL, '/jobQuestion/updateJobQuestion', '更新工单问题', 'jobQuestion', 'PUT');
INSERT INTO `sys_apis` VALUES (182, '2023-05-04 09:23:02.451', '2023-05-04 09:23:02.451', NULL, '/jobQuestion/findJobQuestion', '根据ID获取工单问题', 'jobQuestion', 'GET');
INSERT INTO `sys_apis` VALUES (183, '2023-05-04 09:23:02.666', '2023-05-04 09:23:02.666', NULL, '/jobQuestion/getJobQuestionList', '获取工单问题列表', 'jobQuestion', 'GET');
INSERT INTO `sys_apis` VALUES (184, '2023-05-04 09:26:07.915', '2023-05-04 09:26:07.915', NULL, '/teamApply/createTeamApply', '新增进组申请', 'teamApply', 'POST');
INSERT INTO `sys_apis` VALUES (185, '2023-05-04 09:26:08.137', '2023-05-04 09:26:08.137', NULL, '/teamApply/deleteTeamApply', '删除进组申请', 'teamApply', 'DELETE');
INSERT INTO `sys_apis` VALUES (186, '2023-05-04 09:26:08.361', '2023-05-04 09:26:08.361', NULL, '/teamApply/deleteTeamApplyByIds', '批量删除进组申请', 'teamApply', 'DELETE');
INSERT INTO `sys_apis` VALUES (187, '2023-05-04 09:26:08.575', '2023-05-04 09:26:08.575', NULL, '/teamApply/updateTeamApply', '更新进组申请', 'teamApply', 'PUT');
INSERT INTO `sys_apis` VALUES (188, '2023-05-04 09:26:08.792', '2023-05-04 09:26:08.792', NULL, '/teamApply/findTeamApply', '根据ID获取进组申请', 'teamApply', 'GET');
INSERT INTO `sys_apis` VALUES (189, '2023-05-04 09:26:09.013', '2023-05-04 09:26:09.013', NULL, '/teamApply/getTeamApplyList', '获取进组申请列表', 'teamApply', 'GET');
INSERT INTO `sys_apis` VALUES (190, '2023-05-04 09:34:02.975', '2023-05-04 09:34:02.975', NULL, '/msgBox/createMsgBox', '新增消息盒子', 'msgBox', 'POST');
INSERT INTO `sys_apis` VALUES (191, '2023-05-04 09:34:03.185', '2023-05-04 09:34:03.185', NULL, '/msgBox/deleteMsgBox', '删除消息盒子', 'msgBox', 'DELETE');
INSERT INTO `sys_apis` VALUES (192, '2023-05-04 09:34:03.403', '2023-05-04 09:34:03.403', NULL, '/msgBox/deleteMsgBoxByIds', '批量删除消息盒子', 'msgBox', 'DELETE');
INSERT INTO `sys_apis` VALUES (193, '2023-05-04 09:34:03.621', '2023-05-04 09:34:03.621', NULL, '/msgBox/updateMsgBox', '更新消息盒子', 'msgBox', 'PUT');
INSERT INTO `sys_apis` VALUES (194, '2023-05-04 09:34:03.836', '2023-05-04 09:34:03.836', NULL, '/msgBox/findMsgBox', '根据ID获取消息盒子', 'msgBox', 'GET');
INSERT INTO `sys_apis` VALUES (195, '2023-05-04 09:34:04.060', '2023-05-04 09:34:04.060', NULL, '/msgBox/getMsgBoxList', '获取消息盒子列表', 'msgBox', 'GET');
INSERT INTO `sys_apis` VALUES (196, '2023-05-04 09:38:55.984', '2023-05-04 09:38:55.984', '2023-05-04 09:40:07.772', '/banner/createBanner', '新增轮播图', 'banner', 'POST');
INSERT INTO `sys_apis` VALUES (197, '2023-05-04 09:38:56.181', '2023-05-04 09:38:56.181', '2023-05-04 09:40:07.772', '/banner/deleteBanner', '删除轮播图', 'banner', 'DELETE');
INSERT INTO `sys_apis` VALUES (198, '2023-05-04 09:38:56.369', '2023-05-04 09:38:56.369', '2023-05-04 09:40:07.772', '/banner/deleteBannerByIds', '批量删除轮播图', 'banner', 'DELETE');
INSERT INTO `sys_apis` VALUES (199, '2023-05-04 09:38:56.553', '2023-05-04 09:38:56.553', '2023-05-04 09:40:07.772', '/banner/updateBanner', '更新轮播图', 'banner', 'PUT');
INSERT INTO `sys_apis` VALUES (200, '2023-05-04 09:38:56.740', '2023-05-04 09:38:56.740', '2023-05-04 09:40:07.772', '/banner/findBanner', '根据ID获取轮播图', 'banner', 'GET');
INSERT INTO `sys_apis` VALUES (201, '2023-05-04 09:38:56.926', '2023-05-04 09:38:56.926', '2023-05-04 09:40:07.772', '/banner/getBannerList', '获取轮播图列表', 'banner', 'GET');
INSERT INTO `sys_apis` VALUES (202, '2023-05-04 09:43:06.364', '2023-05-04 09:43:06.364', NULL, '/banner/createBanner', '新增轮播图', 'banner', 'POST');
INSERT INTO `sys_apis` VALUES (203, '2023-05-04 09:43:06.585', '2023-05-04 09:43:06.585', NULL, '/banner/deleteBanner', '删除轮播图', 'banner', 'DELETE');
INSERT INTO `sys_apis` VALUES (204, '2023-05-04 09:43:06.806', '2023-05-04 09:43:06.806', NULL, '/banner/deleteBannerByIds', '批量删除轮播图', 'banner', 'DELETE');
INSERT INTO `sys_apis` VALUES (205, '2023-05-04 09:43:07.020', '2023-05-04 09:43:07.020', NULL, '/banner/updateBanner', '更新轮播图', 'banner', 'PUT');
INSERT INTO `sys_apis` VALUES (206, '2023-05-04 09:43:07.237', '2023-05-04 09:43:07.237', NULL, '/banner/findBanner', '根据ID获取轮播图', 'banner', 'GET');
INSERT INTO `sys_apis` VALUES (207, '2023-05-04 09:43:07.453', '2023-05-04 09:43:07.453', NULL, '/banner/getBannerList', '获取轮播图列表', 'banner', 'GET');
INSERT INTO `sys_apis` VALUES (208, '2023-05-05 10:44:52.000', '2023-05-05 10:44:52.000', NULL, '/companyApply/createCompanyApply', '新增公司申请', 'companyApply', 'POST');
INSERT INTO `sys_apis` VALUES (209, '2023-05-05 10:44:52.004', '2023-05-05 10:44:52.004', NULL, '/companyApply/deleteCompanyApply', '删除公司申请', 'companyApply', 'DELETE');
INSERT INTO `sys_apis` VALUES (210, '2023-05-05 10:44:52.008', '2023-05-05 10:44:52.008', NULL, '/companyApply/deleteCompanyApplyByIds', '批量删除公司申请', 'companyApply', 'DELETE');
INSERT INTO `sys_apis` VALUES (211, '2023-05-05 10:44:52.012', '2023-05-05 10:44:52.012', NULL, '/companyApply/updateCompanyApply', '更新公司申请', 'companyApply', 'PUT');
INSERT INTO `sys_apis` VALUES (212, '2023-05-05 10:44:52.014', '2023-05-05 10:44:52.014', NULL, '/companyApply/findCompanyApply', '根据ID获取公司申请', 'companyApply', 'GET');
INSERT INTO `sys_apis` VALUES (213, '2023-05-05 10:44:52.017', '2023-05-05 10:44:52.017', NULL, '/companyApply/getCompanyApplyList', '获取公司申请列表', 'companyApply', 'GET');
INSERT INTO `sys_apis` VALUES (214, '2023-05-11 21:00:29.974', '2023-05-11 21:00:29.974', NULL, '/jobApply/createJobApply', '新增工单申请', 'jobApply', 'POST');
INSERT INTO `sys_apis` VALUES (215, '2023-05-11 21:00:30.007', '2023-05-11 21:00:30.007', NULL, '/jobApply/deleteJobApply', '删除工单申请', 'jobApply', 'DELETE');
INSERT INTO `sys_apis` VALUES (216, '2023-05-11 21:00:30.039', '2023-05-11 21:00:30.039', NULL, '/jobApply/deleteJobApplyByIds', '批量删除工单申请', 'jobApply', 'DELETE');
INSERT INTO `sys_apis` VALUES (217, '2023-05-11 21:00:30.071', '2023-05-11 21:00:30.071', NULL, '/jobApply/updateJobApply', '更新工单申请', 'jobApply', 'PUT');
INSERT INTO `sys_apis` VALUES (218, '2023-05-11 21:00:30.102', '2023-05-11 21:00:30.102', NULL, '/jobApply/findJobApply', '根据ID获取工单申请', 'jobApply', 'GET');
INSERT INTO `sys_apis` VALUES (219, '2023-05-11 21:00:30.134', '2023-05-11 21:00:30.134', NULL, '/jobApply/getJobApplyList', '获取工单申请列表', 'jobApply', 'GET');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `authority_id` bigint UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`) USING BTREE,
  UNIQUE INDEX `authority_id`(`authority_id` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9529 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES ('2023-05-02 17:42:56.660', '2023-05-02 21:16:40.000', NULL, 888, '普通用户', 0, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2023-05-02 17:42:56.660', '2023-05-02 17:43:06.453', NULL, 8881, '普通用户子角色', 888, 'dashboard');
INSERT INTO `sys_authorities` VALUES ('2023-05-02 17:42:56.660', '2023-05-02 17:43:05.322', NULL, 9528, '测试角色', 0, 'dashboard');

-- ----------------------------
-- Table structure for sys_authority_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_btns`;
CREATE TABLE `sys_authority_btns`  (
  `authority_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_authority_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_authority_menus`;
CREATE TABLE `sys_authority_menus`  (
  `sys_base_menu_id` bigint UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO `sys_authority_menus` VALUES (1, 888);
INSERT INTO `sys_authority_menus` VALUES (1, 8881);
INSERT INTO `sys_authority_menus` VALUES (1, 9528);
INSERT INTO `sys_authority_menus` VALUES (2, 888);
INSERT INTO `sys_authority_menus` VALUES (2, 8881);
INSERT INTO `sys_authority_menus` VALUES (2, 9528);
INSERT INTO `sys_authority_menus` VALUES (3, 888);
INSERT INTO `sys_authority_menus` VALUES (4, 888);
INSERT INTO `sys_authority_menus` VALUES (4, 8881);
INSERT INTO `sys_authority_menus` VALUES (5, 888);
INSERT INTO `sys_authority_menus` VALUES (5, 8881);
INSERT INTO `sys_authority_menus` VALUES (6, 888);
INSERT INTO `sys_authority_menus` VALUES (6, 8881);
INSERT INTO `sys_authority_menus` VALUES (7, 888);
INSERT INTO `sys_authority_menus` VALUES (7, 8881);
INSERT INTO `sys_authority_menus` VALUES (8, 888);
INSERT INTO `sys_authority_menus` VALUES (8, 8881);
INSERT INTO `sys_authority_menus` VALUES (8, 9528);
INSERT INTO `sys_authority_menus` VALUES (9, 888);
INSERT INTO `sys_authority_menus` VALUES (9, 8881);
INSERT INTO `sys_authority_menus` VALUES (10, 888);
INSERT INTO `sys_authority_menus` VALUES (10, 8881);
INSERT INTO `sys_authority_menus` VALUES (11, 888);
INSERT INTO `sys_authority_menus` VALUES (11, 8881);
INSERT INTO `sys_authority_menus` VALUES (12, 888);
INSERT INTO `sys_authority_menus` VALUES (13, 888);
INSERT INTO `sys_authority_menus` VALUES (13, 8881);
INSERT INTO `sys_authority_menus` VALUES (14, 888);
INSERT INTO `sys_authority_menus` VALUES (14, 8881);
INSERT INTO `sys_authority_menus` VALUES (15, 888);
INSERT INTO `sys_authority_menus` VALUES (15, 8881);
INSERT INTO `sys_authority_menus` VALUES (16, 888);
INSERT INTO `sys_authority_menus` VALUES (16, 8881);
INSERT INTO `sys_authority_menus` VALUES (17, 888);
INSERT INTO `sys_authority_menus` VALUES (17, 8881);
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
INSERT INTO `sys_authority_menus` VALUES (33, 888);
INSERT INTO `sys_authority_menus` VALUES (34, 888);

-- ----------------------------
-- Table structure for sys_auto_code_histories
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_code_histories`;
CREATE TABLE `sys_auto_code_histories`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `package` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `business_db` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `table_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `request_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `auto_code_path` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `injection_meta` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL,
  `struct_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `struct_cn_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `api_ids` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `flag` bigint NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_code_histories_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 22 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_auto_code_histories
-- ----------------------------
INSERT INTO `sys_auto_code_histories` VALUES (1, '2023-05-02 20:30:00.821', '2023-05-02 20:53:42.876', '2023-05-02 20:57:47.474', 'clothing', '', 'app_user', '{\"structName\":\"AppUser\",\"tableName\":\"app_user\",\"packageName\":\"appUser\",\"humpPackageName\":\"app_user\",\"abbreviation\":\"appUser\",\"description\":\"app用户\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"UserName\",\"fieldDesc\":\"用户名\",\"fieldType\":\"string\",\"fieldJson\":\"userName\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Password\",\"fieldDesc\":\"密码\",\"fieldType\":\"string\",\"fieldJson\":\"password\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"password\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RoleID\",\"fieldDesc\":\"角色id\",\"fieldType\":\"int\",\"fieldJson\":\"roleID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"role_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Nickname\",\"fieldDesc\":\"昵称\",\"fieldType\":\"string\",\"fieldJson\":\"nickname\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"nickname\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompanyID\",\"fieldDesc\":\"公司id\",\"fieldType\":\"int\",\"fieldJson\":\"companyID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"company_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Wages\",\"fieldDesc\":\"工资\",\"fieldType\":\"float64\",\"fieldJson\":\"wages\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"wages\",\"fieldSearchType\":\"BETWEEN\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":true},{\"fieldName\":\"WorkType\",\"fieldDesc\":\"工作计算方式\",\"fieldType\":\"int\",\"fieldJson\":\"workType\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"work_type\",\"fieldSearchType\":\"=\",\"dictType\":\"workType\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"status\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\app_user.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\app_user.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\request\\app_user.go;d:\\project\\gin-vue-admin\\server\\router\\clothing\\app_user.go;d:\\project\\gin-vue-admin\\server\\service\\clothing\\app_user.go;d:\\project\\gin-vue-admin\\web\\src\\api\\appUser.js;d:\\project\\gin-vue-admin\\web\\src\\view\\appUser\\appUserForm.vue;d:\\project\\gin-vue-admin\\web\\src\\view\\appUser\\appUser.vue;', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\enter.go@ApiGroup@AppUserApi;d:\\project\\gin-vue-admin\\server\\router\\clothing\\enter.go@RouterGroup@AppUserRouter;d:\\project\\gin-vue-admin\\server\\service\\clothing\\enter.go@ServiceGroup@AppUserService;', 'AppUser', 'app用户', '94;95;96;97;98;99;', 1);
INSERT INTO `sys_auto_code_histories` VALUES (2, '2023-05-02 20:34:45.077', '2023-05-02 20:34:45.077', NULL, 'clothing', '', 'company', '{\"structName\":\"Company\",\"tableName\":\"company\",\"packageName\":\"company\",\"humpPackageName\":\"company\",\"abbreviation\":\"company\",\"description\":\"公司\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"Name\",\"fieldDesc\":\"公司名\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"status\",\"fieldSearchType\":\"=\",\"dictType\":\"status\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\company.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\company.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\request\\company.go;d:\\project\\gin-vue-admin\\server\\router\\clothing\\company.go;d:\\project\\gin-vue-admin\\server\\service\\clothing\\company.go;d:\\project\\gin-vue-admin\\web\\src\\api\\company.js;d:\\project\\gin-vue-admin\\web\\src\\view\\company\\companyForm.vue;d:\\project\\gin-vue-admin\\web\\src\\view\\company\\company.vue;', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\enter.go@ApiGroup@CompanyApi;d:\\project\\gin-vue-admin\\server\\router\\clothing\\enter.go@RouterGroup@CompanyRouter;d:\\project\\gin-vue-admin\\server\\service\\clothing\\enter.go@ServiceGroup@CompanyService;', 'Company', '公司', '100;101;102;103;104;105;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (3, '2023-05-02 20:36:44.753', '2023-05-02 20:36:44.753', NULL, 'clothing', '', 'app_role', '{\"structName\":\"AppRole\",\"tableName\":\"app_role\",\"packageName\":\"appRole\",\"humpPackageName\":\"app_role\",\"abbreviation\":\"appRole\",\"description\":\"app角色\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"Name\",\"fieldDesc\":\"角色名\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\app_role.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\app_role.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\request\\app_role.go;d:\\project\\gin-vue-admin\\server\\router\\clothing\\app_role.go;d:\\project\\gin-vue-admin\\server\\service\\clothing\\app_role.go;d:\\project\\gin-vue-admin\\web\\src\\api\\appRole.js;d:\\project\\gin-vue-admin\\web\\src\\view\\appRole\\appRoleForm.vue;d:\\project\\gin-vue-admin\\web\\src\\view\\appRole\\appRole.vue;', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\enter.go@ApiGroup@AppRoleApi;d:\\project\\gin-vue-admin\\server\\router\\clothing\\enter.go@RouterGroup@AppRoleRouter;d:\\project\\gin-vue-admin\\server\\service\\clothing\\enter.go@ServiceGroup@AppRoleService;', 'AppRole', 'app角色', '106;107;108;109;110;111;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (4, '2023-05-02 20:49:40.044', '2023-05-02 20:49:40.044', NULL, 'clothing', '', 'team', '{\"structName\":\"Team\",\"tableName\":\"team\",\"packageName\":\"team\",\"humpPackageName\":\"team\",\"abbreviation\":\"team\",\"description\":\"组\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"CompanyID\",\"fieldDesc\":\"公司ID\",\"fieldType\":\"int\",\"fieldJson\":\"companyID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"company_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"组名\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\team.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\team.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\request\\team.go;d:\\project\\gin-vue-admin\\server\\router\\clothing\\team.go;d:\\project\\gin-vue-admin\\server\\service\\clothing\\team.go;d:\\project\\gin-vue-admin\\web\\src\\api\\team.js;d:\\project\\gin-vue-admin\\web\\src\\view\\team\\teamForm.vue;d:\\project\\gin-vue-admin\\web\\src\\view\\team\\team.vue;', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\enter.go@ApiGroup@TeamApi;d:\\project\\gin-vue-admin\\server\\router\\clothing\\enter.go@RouterGroup@TeamRouter;d:\\project\\gin-vue-admin\\server\\service\\clothing\\enter.go@ServiceGroup@TeamService;', 'Team', '组', '112;113;114;115;116;117;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (5, '2023-05-02 20:57:36.411', '2023-05-02 20:57:36.411', NULL, 'clothing', '', 'app_user', '{\"structName\":\"AppUser\",\"tableName\":\"app_user\",\"packageName\":\"appUser\",\"humpPackageName\":\"app_user\",\"abbreviation\":\"appUser\",\"description\":\"app用户\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"UserName\",\"fieldDesc\":\"用户名\",\"fieldType\":\"string\",\"fieldJson\":\"userName\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Password\",\"fieldDesc\":\"密码\",\"fieldType\":\"string\",\"fieldJson\":\"password\",\"dataTypeLong\":\"64\",\"comment\":\"\",\"columnName\":\"password\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Nickname\",\"fieldDesc\":\"昵称\",\"fieldType\":\"string\",\"fieldJson\":\"nickname\",\"dataTypeLong\":\"20\",\"comment\":\"\",\"columnName\":\"nickname\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Wages\",\"fieldDesc\":\"工资\",\"fieldType\":\"float64\",\"fieldJson\":\"wages\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"wages\",\"fieldSearchType\":\"BETWEEN\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":true},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"1\",\"comment\":\"\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"status\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PhoneNum\",\"fieldDesc\":\"手机号\",\"fieldType\":\"string\",\"fieldJson\":\"phoneNum\",\"dataTypeLong\":\"11\",\"comment\":\"\",\"columnName\":\"phone_num\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"OpenID\",\"fieldDesc\":\"openID\",\"fieldType\":\"string\",\"fieldJson\":\"openID\",\"dataTypeLong\":\"64\",\"comment\":\"\",\"columnName\":\"open_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UnionID\",\"fieldDesc\":\"unionID\",\"fieldType\":\"string\",\"fieldJson\":\"unionID\",\"dataTypeLong\":\"64\",\"comment\":\"\",\"columnName\":\"union_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\app_user.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\app_user.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\request\\app_user.go;d:\\project\\gin-vue-admin\\server\\router\\clothing\\app_user.go;d:\\project\\gin-vue-admin\\server\\service\\clothing\\app_user.go;d:\\project\\gin-vue-admin\\web\\src\\api\\appUser.js;d:\\project\\gin-vue-admin\\web\\src\\view\\appUser\\appUserForm.vue;d:\\project\\gin-vue-admin\\web\\src\\view\\appUser\\appUser.vue;', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\enter.go@ApiGroup@AppUserApi;d:\\project\\gin-vue-admin\\server\\router\\clothing\\enter.go@RouterGroup@AppUserRouter;d:\\project\\gin-vue-admin\\server\\service\\clothing\\enter.go@ServiceGroup@AppUserService;', 'AppUser', 'app用户', '118;119;120;121;122;123;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (6, '2023-05-02 21:08:33.545', '2023-05-02 21:08:33.545', NULL, 'clothing', '', 'user_role', '{\"structName\":\"UserRole\",\"tableName\":\"user_role\",\"packageName\":\"userRole\",\"humpPackageName\":\"user_role\",\"abbreviation\":\"userRole\",\"description\":\"用户角色\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"UserID\",\"fieldDesc\":\"用户id\",\"fieldType\":\"int\",\"fieldJson\":\"userID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RoleID\",\"fieldDesc\":\"角色id\",\"fieldType\":\"int\",\"fieldJson\":\"roleID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"role_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompanyID\",\"fieldDesc\":\"公司id\",\"fieldType\":\"int\",\"fieldJson\":\"companyID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"company_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\user_role.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\user_role.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\request\\user_role.go;d:\\project\\gin-vue-admin\\server\\router\\clothing\\user_role.go;d:\\project\\gin-vue-admin\\server\\service\\clothing\\user_role.go;d:\\project\\gin-vue-admin\\web\\src\\api\\userRole.js;d:\\project\\gin-vue-admin\\web\\src\\view\\userRole\\userRoleForm.vue;d:\\project\\gin-vue-admin\\web\\src\\view\\userRole\\userRole.vue;', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\enter.go@ApiGroup@UserRoleApi;d:\\project\\gin-vue-admin\\server\\router\\clothing\\enter.go@RouterGroup@UserRoleRouter;d:\\project\\gin-vue-admin\\server\\service\\clothing\\enter.go@ServiceGroup@UserRoleService;', 'UserRole', '用户角色', '124;125;126;127;128;129;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (7, '2023-05-02 21:26:24.413', '2023-05-02 21:26:24.413', NULL, 'clothing', '', 'user_wallet', '{\"structName\":\"UserWallet\",\"tableName\":\"user_wallet\",\"packageName\":\"userWallet\",\"humpPackageName\":\"user_wallet\",\"abbreviation\":\"userWallet\",\"description\":\"用户钱包\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"UserID\",\"fieldDesc\":\"用户id\",\"fieldType\":\"int\",\"fieldJson\":\"userID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"CompanyID\",\"fieldDesc\":\"公司id\",\"fieldType\":\"int\",\"fieldJson\":\"companyID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"company_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Wages\",\"fieldDesc\":\"工资\",\"fieldType\":\"float64\",\"fieldJson\":\"wages\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"wages\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\user_wallet.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\user_wallet.go;d:\\project\\gin-vue-admin\\server\\model\\clothing\\request\\user_wallet.go;d:\\project\\gin-vue-admin\\server\\router\\clothing\\user_wallet.go;d:\\project\\gin-vue-admin\\server\\service\\clothing\\user_wallet.go;d:\\project\\gin-vue-admin\\web\\src\\api\\userWallet.js;d:\\project\\gin-vue-admin\\web\\src\\view\\userWallet\\userWalletForm.vue;d:\\project\\gin-vue-admin\\web\\src\\view\\userWallet\\userWallet.vue;', 'd:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\enter.go@ApiGroup@UserWalletApi;d:\\project\\gin-vue-admin\\server\\router\\clothing\\enter.go@RouterGroup@UserWalletRouter;d:\\project\\gin-vue-admin\\server\\service\\clothing\\enter.go@ServiceGroup@UserWalletService;', 'UserWallet', '用户钱包', '130;131;132;133;134;135;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (8, '2023-05-03 09:21:50.338', '2023-05-03 09:21:50.338', NULL, 'clothing', '', 'team_user', '{\"structName\":\"TeamUser\",\"tableName\":\"team_user\",\"packageName\":\"teamUser\",\"humpPackageName\":\"team_user\",\"abbreviation\":\"teamUser\",\"description\":\"组员\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TeamID\",\"fieldDesc\":\"组id\",\"fieldType\":\"int\",\"fieldJson\":\"teamID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"team_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserID\",\"fieldDesc\":\"用户id\",\"fieldType\":\"int\",\"fieldJson\":\"userID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\team_user.go;d:\\project\\clothing\\server\\model\\clothing\\team_user.go;d:\\project\\clothing\\server\\model\\clothing\\request\\team_user.go;d:\\project\\clothing\\server\\router\\clothing\\team_user.go;d:\\project\\clothing\\server\\service\\clothing\\team_user.go;d:\\project\\clothing\\web\\src\\api\\teamUser.js;d:\\project\\clothing\\web\\src\\view\\teamUser\\teamUserForm.vue;d:\\project\\clothing\\web\\src\\view\\teamUser\\teamUser.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@TeamUserApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@TeamUserRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@TeamUserService;', 'TeamUser', '组员', '136;137;138;139;140;141;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (9, '2023-05-03 10:02:20.957', '2023-05-03 11:05:50.408', '2023-05-03 11:10:13.725', 'clothing', '', 'cloth', '{\"structName\":\"Cloth\",\"tableName\":\"cloth\",\"packageName\":\"cloth\",\"humpPackageName\":\"cloth\",\"abbreviation\":\"cloth\",\"description\":\"布料入库单\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"CompanyID\",\"fieldDesc\":\"公司id\",\"fieldType\":\"int\",\"fieldJson\":\"companyID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"company_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"PackageNo\",\"fieldDesc\":\"包号\",\"fieldType\":\"int\",\"fieldJson\":\"packageNo\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"package_no\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Style\",\"fieldDesc\":\"款式\",\"fieldType\":\"string\",\"fieldJson\":\"style\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"style\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Color\",\"fieldDesc\":\"颜色\",\"fieldType\":\"string\",\"fieldJson\":\"color\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"color\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Length\",\"fieldDesc\":\"长度\",\"fieldType\":\"float64\",\"fieldJson\":\"length\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"length\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RemainingLength\",\"fieldDesc\":\"剩余长度\",\"fieldType\":\"float64\",\"fieldJson\":\"remainingLength\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"remaining_length\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\cloth.go;d:\\project\\clothing\\server\\model\\clothing\\cloth.go;d:\\project\\clothing\\server\\model\\clothing\\request\\cloth.go;d:\\project\\clothing\\server\\router\\clothing\\cloth.go;d:\\project\\clothing\\server\\service\\clothing\\cloth.go;d:\\project\\clothing\\web\\src\\api\\cloth.js;d:\\project\\clothing\\web\\src\\view\\cloth\\clothForm.vue;d:\\project\\clothing\\web\\src\\view\\cloth\\cloth.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@ClothApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@ClothRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@ClothService;', 'Cloth', '布料入库单', '142;143;144;145;146;147;', 1);
INSERT INTO `sys_auto_code_histories` VALUES (10, '2023-05-03 11:07:34.540', '2023-05-03 11:07:34.540', NULL, 'clothing', '', 'cloth', '{\"structName\":\"Cloth\",\"tableName\":\"cloth\",\"packageName\":\"cloth\",\"humpPackageName\":\"cloth\",\"abbreviation\":\"cloth\",\"description\":\"布料入库单\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"CompanyID\",\"fieldDesc\":\"公司id\",\"fieldType\":\"int\",\"fieldJson\":\"companyID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"company_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"StyleID\",\"fieldDesc\":\"款式ID\",\"fieldType\":\"int\",\"fieldJson\":\"styleID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"style_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Color\",\"fieldDesc\":\"颜色\",\"fieldType\":\"string\",\"fieldJson\":\"color\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"color\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Length\",\"fieldDesc\":\"长度\",\"fieldType\":\"float64\",\"fieldJson\":\"length\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"length\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\cloth.go;d:\\project\\clothing\\server\\model\\clothing\\cloth.go;d:\\project\\clothing\\server\\model\\clothing\\request\\cloth.go;d:\\project\\clothing\\server\\router\\clothing\\cloth.go;d:\\project\\clothing\\server\\service\\clothing\\cloth.go;d:\\project\\clothing\\web\\src\\api\\cloth.js;d:\\project\\clothing\\web\\src\\view\\cloth\\clothForm.vue;d:\\project\\clothing\\web\\src\\view\\cloth\\cloth.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@ClothApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@ClothRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@ClothService;', 'Cloth', '布料入库单', '148;149;150;151;152;153;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (11, '2023-05-03 11:12:36.338', '2023-05-03 11:12:36.338', NULL, 'clothing', '', 'style', '{\"structName\":\"Style\",\"tableName\":\"style\",\"packageName\":\"style\",\"humpPackageName\":\"style\",\"abbreviation\":\"style\",\"description\":\"款式\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"CompanyID\",\"fieldDesc\":\"公司id\",\"fieldType\":\"int\",\"fieldJson\":\"companyID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"company_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"name\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"name\",\"fieldSearchType\":\"LIKE\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Price\",\"fieldDesc\":\"成衣价格\",\"fieldType\":\"float64\",\"fieldJson\":\"price\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"price\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\style.go;d:\\project\\clothing\\server\\model\\clothing\\style.go;d:\\project\\clothing\\server\\model\\clothing\\request\\style.go;d:\\project\\clothing\\server\\router\\clothing\\style.go;d:\\project\\clothing\\server\\service\\clothing\\style.go;d:\\project\\clothing\\web\\src\\api\\style.js;d:\\project\\clothing\\web\\src\\view\\style\\styleForm.vue;d:\\project\\clothing\\web\\src\\view\\style\\style.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@StyleApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@StyleRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@StyleService;', 'Style', '款式', '154;155;156;157;158;159;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (12, '2023-05-03 11:17:44.205', '2023-05-03 11:17:44.205', NULL, 'clothing', '', 'process', '{\"structName\":\"Process\",\"tableName\":\"process\",\"packageName\":\"process\",\"humpPackageName\":\"process\",\"abbreviation\":\"process\",\"description\":\"工序\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"StyleID\",\"fieldDesc\":\"款式id\",\"fieldType\":\"int\",\"fieldJson\":\"styleID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"style_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Name\",\"fieldDesc\":\"工序名称\",\"fieldType\":\"string\",\"fieldJson\":\"name\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Price\",\"fieldDesc\":\"价格\",\"fieldType\":\"float64\",\"fieldJson\":\"price\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"price\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Percent\",\"fieldDesc\":\"百分比\",\"fieldType\":\"float64\",\"fieldJson\":\"percent\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"percent\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\process.go;d:\\project\\clothing\\server\\model\\clothing\\process.go;d:\\project\\clothing\\server\\model\\clothing\\request\\process.go;d:\\project\\clothing\\server\\router\\clothing\\process.go;d:\\project\\clothing\\server\\service\\clothing\\process.go;d:\\project\\clothing\\web\\src\\api\\process.js;d:\\project\\clothing\\web\\src\\view\\process\\processForm.vue;d:\\project\\clothing\\web\\src\\view\\process\\process.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@ProcessApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@ProcessRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@ProcessService;', 'Process', '工序', '160;161;162;163;164;165;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (13, '2023-05-03 15:16:47.065', '2023-05-03 15:16:47.065', NULL, 'clothing', '', 'cropping_record', '{\"structName\":\"CroppingRecord\",\"tableName\":\"cropping_record\",\"packageName\":\"croppingRecord\",\"humpPackageName\":\"cropping_record\",\"abbreviation\":\"croppingRecord\",\"description\":\"裁剪单\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"CompanyID\",\"fieldDesc\":\"公司id\",\"fieldType\":\"int\",\"fieldJson\":\"companyID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"company_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"StyleID\",\"fieldDesc\":\"款式id\",\"fieldType\":\"int\",\"fieldJson\":\"styleID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"style_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Color\",\"fieldDesc\":\"颜色\",\"fieldType\":\"string\",\"fieldJson\":\"color\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"color\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Length\",\"fieldDesc\":\"长度\",\"fieldType\":\"float64\",\"fieldJson\":\"length\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"length\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Size\",\"fieldDesc\":\"尺码\",\"fieldType\":\"string\",\"fieldJson\":\"size\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"size\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Quantity\",\"fieldDesc\":\"数量\",\"fieldType\":\"int\",\"fieldJson\":\"quantity\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"quantity\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Step\",\"fieldDesc\":\"步骤\",\"fieldType\":\"int\",\"fieldJson\":\"step\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"step\",\"fieldSearchType\":\"\",\"dictType\":\"step\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\cropping_record.go;d:\\project\\clothing\\server\\model\\clothing\\cropping_record.go;d:\\project\\clothing\\server\\model\\clothing\\request\\cropping_record.go;d:\\project\\clothing\\server\\router\\clothing\\cropping_record.go;d:\\project\\clothing\\server\\service\\clothing\\cropping_record.go;d:\\project\\clothing\\web\\src\\api\\croppingRecord.js;d:\\project\\clothing\\web\\src\\view\\croppingRecord\\croppingRecordForm.vue;d:\\project\\clothing\\web\\src\\view\\croppingRecord\\croppingRecord.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@CroppingRecordApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@CroppingRecordRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@CroppingRecordService;', 'CroppingRecord', '裁剪单', '166;167;168;169;170;171;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (14, '2023-05-03 15:58:07.214', '2023-05-03 15:58:07.214', NULL, 'clothing', '', 'job', '{\"structName\":\"Job\",\"tableName\":\"job\",\"packageName\":\"job\",\"humpPackageName\":\"job\",\"abbreviation\":\"job\",\"description\":\"工单\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"CroppingID\",\"fieldDesc\":\"裁剪单id\",\"fieldType\":\"int\",\"fieldJson\":\"croppingID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"cropping_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserID\",\"fieldDesc\":\"用户id\",\"fieldType\":\"int\",\"fieldJson\":\"userID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ProcessName\",\"fieldDesc\":\"工序名称\",\"fieldType\":\"string\",\"fieldJson\":\"processName\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"process_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Price\",\"fieldDesc\":\"单价\",\"fieldType\":\"float64\",\"fieldJson\":\"price\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"price\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Quantity\",\"fieldDesc\":\"申请数量\",\"fieldType\":\"int\",\"fieldJson\":\"quantity\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"quantity\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Income\",\"fieldDesc\":\"预估收入\",\"fieldType\":\"float64\",\"fieldJson\":\"income\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"income\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RealQuantity\",\"fieldDesc\":\"真实数量\",\"fieldType\":\"int\",\"fieldJson\":\"realQuantity\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"real_quantity\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RealIncome\",\"fieldDesc\":\"真实收入\",\"fieldType\":\"float64\",\"fieldJson\":\"realIncome\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"real_income\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Step\",\"fieldDesc\":\"步骤\",\"fieldType\":\"int\",\"fieldJson\":\"step\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"step\",\"fieldSearchType\":\"=\",\"dictType\":\"step\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"JobType\",\"fieldDesc\":\"工作类型\",\"fieldType\":\"int\",\"fieldJson\":\"jobType\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"job_type\",\"fieldSearchType\":\"\",\"dictType\":\"jobType\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\job.go;d:\\project\\clothing\\server\\model\\clothing\\job.go;d:\\project\\clothing\\server\\model\\clothing\\request\\job.go;d:\\project\\clothing\\server\\router\\clothing\\job.go;d:\\project\\clothing\\server\\service\\clothing\\job.go;d:\\project\\clothing\\web\\src\\api\\job.js;d:\\project\\clothing\\web\\src\\view\\job\\jobForm.vue;d:\\project\\clothing\\web\\src\\view\\job\\job.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@JobApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@JobRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@JobService;', 'Job', '工单', '172;173;174;175;176;177;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (15, '2023-05-04 09:23:03.147', '2023-05-04 09:23:03.147', NULL, 'clothing', '', 'job_question', '{\"structName\":\"JobQuestion\",\"tableName\":\"job_question\",\"packageName\":\"jobQuestion\",\"humpPackageName\":\"job_question\",\"abbreviation\":\"jobQuestion\",\"description\":\"工单问题\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"JobID\",\"fieldDesc\":\"工单id\",\"fieldType\":\"int\",\"fieldJson\":\"jobID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"job_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserID\",\"fieldDesc\":\"用户id\",\"fieldType\":\"int\",\"fieldJson\":\"userID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Content\",\"fieldDesc\":\"内容\",\"fieldType\":\"string\",\"fieldJson\":\"content\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"handleStatus\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\job_question.go;d:\\project\\clothing\\server\\model\\clothing\\job_question.go;d:\\project\\clothing\\server\\model\\clothing\\request\\job_question.go;d:\\project\\clothing\\server\\router\\clothing\\job_question.go;d:\\project\\clothing\\server\\service\\clothing\\job_question.go;d:\\project\\clothing\\web\\src\\api\\jobQuestion.js;d:\\project\\clothing\\web\\src\\view\\jobQuestion\\jobQuestionForm.vue;d:\\project\\clothing\\web\\src\\view\\jobQuestion\\jobQuestion.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@JobQuestionApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@JobQuestionRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@JobQuestionService;', 'JobQuestion', '工单问题', '178;179;180;181;182;183;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (16, '2023-05-04 09:26:09.373', '2023-05-04 09:26:09.373', NULL, 'clothing', '', 'team_apply', '{\"structName\":\"TeamApply\",\"tableName\":\"team_apply\",\"packageName\":\"teamApply\",\"humpPackageName\":\"team_apply\",\"abbreviation\":\"teamApply\",\"description\":\"进组申请\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"TeamID\",\"fieldDesc\":\"组id\",\"fieldType\":\"int\",\"fieldJson\":\"teamID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"team_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserID\",\"fieldDesc\":\"申请人ID\",\"fieldType\":\"int\",\"fieldJson\":\"userID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"handleStatus\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\team_apply.go;d:\\project\\clothing\\server\\model\\clothing\\team_apply.go;d:\\project\\clothing\\server\\model\\clothing\\request\\team_apply.go;d:\\project\\clothing\\server\\router\\clothing\\team_apply.go;d:\\project\\clothing\\server\\service\\clothing\\team_apply.go;d:\\project\\clothing\\web\\src\\api\\teamApply.js;d:\\project\\clothing\\web\\src\\view\\teamApply\\teamApplyForm.vue;d:\\project\\clothing\\web\\src\\view\\teamApply\\teamApply.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@TeamApplyApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@TeamApplyRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@TeamApplyService;', 'TeamApply', '进组申请', '184;185;186;187;188;189;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (17, '2023-05-04 09:34:04.419', '2023-05-04 09:34:04.419', NULL, 'clothing', '', 'msg_box', '{\"structName\":\"MsgBox\",\"tableName\":\"msg_box\",\"packageName\":\"msgBox\",\"humpPackageName\":\"msg_box\",\"abbreviation\":\"msgBox\",\"description\":\"消息盒子\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"UserID\",\"fieldDesc\":\"用户id\",\"fieldType\":\"int\",\"fieldJson\":\"userID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MsgType\",\"fieldDesc\":\"消息类型\",\"fieldType\":\"int\",\"fieldJson\":\"msgType\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"msg_type\",\"fieldSearchType\":\"=\",\"dictType\":\"msgType\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MsgID\",\"fieldDesc\":\"消息id\",\"fieldType\":\"int\",\"fieldJson\":\"msgID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"msg_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"status\",\"fieldSearchType\":\"\",\"dictType\":\"handleStatus\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\msg_box.go;d:\\project\\clothing\\server\\model\\clothing\\msg_box.go;d:\\project\\clothing\\server\\model\\clothing\\request\\msg_box.go;d:\\project\\clothing\\server\\router\\clothing\\msg_box.go;d:\\project\\clothing\\server\\service\\clothing\\msg_box.go;d:\\project\\clothing\\web\\src\\api\\msgBox.js;d:\\project\\clothing\\web\\src\\view\\msgBox\\msgBoxForm.vue;d:\\project\\clothing\\web\\src\\view\\msgBox\\msgBox.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@MsgBoxApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@MsgBoxRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@MsgBoxService;', 'MsgBox', '消息盒子', '190;191;192;193;194;195;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (18, '2023-05-04 09:38:57.294', '2023-05-04 09:40:09.009', NULL, 'clothing', '', 'banner', '{\"structName\":\"Banner\",\"tableName\":\"banner\",\"packageName\":\"banner\",\"humpPackageName\":\"banner\",\"abbreviation\":\"banner\",\"description\":\"轮播图\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"Url\",\"fieldDesc\":\"跳转链接\",\"fieldType\":\"string\",\"fieldJson\":\"url\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"url\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Title\",\"fieldDesc\":\"标题\",\"fieldType\":\"string\",\"fieldJson\":\"title\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"title\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Content\",\"fieldDesc\":\"内容\",\"fieldType\":\"string\",\"fieldJson\":\"content\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Img\",\"fieldDesc\":\"图片链接\",\"fieldType\":\"string\",\"fieldJson\":\"img\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"img\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"status\",\"fieldSearchType\":\"=\",\"dictType\":\"status\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\banner.go;d:\\project\\clothing\\server\\model\\clothing\\banner.go;d:\\project\\clothing\\server\\model\\clothing\\request\\banner.go;d:\\project\\clothing\\server\\router\\clothing\\banner.go;d:\\project\\clothing\\server\\service\\clothing\\banner.go;d:\\project\\clothing\\web\\src\\api\\banner.js;d:\\project\\clothing\\web\\src\\view\\banner\\bannerForm.vue;d:\\project\\clothing\\web\\src\\view\\banner\\banner.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@BannerApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@BannerRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@BannerService;', 'Banner', '轮播图', '196;197;198;199;200;201;', 1);
INSERT INTO `sys_auto_code_histories` VALUES (19, '2023-05-04 09:43:07.813', '2023-05-04 09:43:07.813', NULL, 'clothing', '', 'banner', '{\"structName\":\"Banner\",\"tableName\":\"banner\",\"packageName\":\"banner\",\"humpPackageName\":\"banner\",\"abbreviation\":\"banner\",\"description\":\"轮播图\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"Url\",\"fieldDesc\":\"跳转链接\",\"fieldType\":\"string\",\"fieldJson\":\"url\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"url\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Title\",\"fieldDesc\":\"标题\",\"fieldType\":\"string\",\"fieldJson\":\"title\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"title\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Content\",\"fieldDesc\":\"内容\",\"fieldType\":\"string\",\"fieldJson\":\"content\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"MediaUrl\",\"fieldDesc\":\"媒体链接\",\"fieldType\":\"string\",\"fieldJson\":\"mediaUrl\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"media_url\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Sort\",\"fieldDesc\":\"排序\",\"fieldType\":\"int\",\"fieldJson\":\"sort\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"sort\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"status\",\"fieldSearchType\":\"=\",\"dictType\":\"status\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Type\",\"fieldDesc\":\"媒体类型\",\"fieldType\":\"int\",\"fieldJson\":\"type\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"type\",\"fieldSearchType\":\"\",\"dictType\":\"mediaType\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\banner.go;d:\\project\\clothing\\server\\model\\clothing\\banner.go;d:\\project\\clothing\\server\\model\\clothing\\request\\banner.go;d:\\project\\clothing\\server\\router\\clothing\\banner.go;d:\\project\\clothing\\server\\service\\clothing\\banner.go;d:\\project\\clothing\\web\\src\\api\\banner.js;d:\\project\\clothing\\web\\src\\view\\banner\\bannerForm.vue;d:\\project\\clothing\\web\\src\\view\\banner\\banner.vue;', 'd:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@BannerApi;d:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@BannerRouter;d:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@BannerService;', 'Banner', '轮播图', '202;203;204;205;206;207;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (20, '2023-05-05 10:44:52.190', '2023-05-05 10:44:52.190', NULL, 'clothing', '', 'company_apply', '{\"structName\":\"CompanyApply\",\"tableName\":\"company_apply\",\"packageName\":\"companyApply\",\"humpPackageName\":\"company_apply\",\"abbreviation\":\"companyApply\",\"description\":\"公司申请\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"CompanyID\",\"fieldDesc\":\"公司id\",\"fieldType\":\"int\",\"fieldJson\":\"companyID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"company_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserID\",\"fieldDesc\":\"申请人id\",\"fieldType\":\"int\",\"fieldJson\":\"userID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"RoleID\",\"fieldDesc\":\"角色id\",\"fieldType\":\"int\",\"fieldJson\":\"roleID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"role_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Remark\",\"fieldDesc\":\"备注\",\"fieldType\":\"string\",\"fieldJson\":\"remark\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"remark\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'D:\\project\\clothing\\server\\api\\v1\\clothing\\company_apply.go;D:\\project\\clothing\\server\\model\\clothing\\company_apply.go;D:\\project\\clothing\\server\\model\\clothing\\request\\company_apply.go;D:\\project\\clothing\\server\\router\\clothing\\company_apply.go;D:\\project\\clothing\\server\\service\\clothing\\company_apply.go;D:\\project\\clothing\\web\\src\\api\\companyApply.js;D:\\project\\clothing\\web\\src\\view\\companyApply\\companyApplyForm.vue;D:\\project\\clothing\\web\\src\\view\\companyApply\\companyApply.vue;', 'D:\\project\\clothing\\server\\api\\v1\\clothing\\enter.go@ApiGroup@CompanyApplyApi;D:\\project\\clothing\\server\\router\\clothing\\enter.go@RouterGroup@CompanyApplyRouter;D:\\project\\clothing\\server\\service\\clothing\\enter.go@ServiceGroup@CompanyApplyService;', 'CompanyApply', '公司申请', '208;209;210;211;212;213;', 0);
INSERT INTO `sys_auto_code_histories` VALUES (21, '2023-05-11 21:00:30.259', '2023-05-11 21:00:30.259', NULL, 'clothing', '', 'job_apply', '{\"structName\":\"JobApply\",\"tableName\":\"job_apply\",\"packageName\":\"jobApply\",\"humpPackageName\":\"job_apply\",\"abbreviation\":\"jobApply\",\"description\":\"工单申请\",\"autoCreateApiToSql\":true,\"autoCreateResource\":true,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"CroppingID\",\"fieldDesc\":\"裁剪单id\",\"fieldType\":\"int\",\"fieldJson\":\"croppingID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"cropping_id\",\"fieldSearchType\":\"=\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"UserID\",\"fieldDesc\":\"用户id\",\"fieldType\":\"int\",\"fieldJson\":\"userID\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"user_id\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"ProcessName\",\"fieldDesc\":\"工序名\",\"fieldType\":\"string\",\"fieldJson\":\"processName\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"process_name\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Price\",\"fieldDesc\":\"价格\",\"fieldType\":\"float64\",\"fieldJson\":\"price\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"price\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Quantity\",\"fieldDesc\":\"数量\",\"fieldType\":\"int\",\"fieldJson\":\"quantity\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"quantity\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"JobType\",\"fieldDesc\":\"计价类型\",\"fieldType\":\"int\",\"fieldJson\":\"jobType\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"job_type\",\"fieldSearchType\":\"\",\"dictType\":\"jobType\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Status\",\"fieldDesc\":\"审核状态\",\"fieldType\":\"int\",\"fieldJson\":\"status\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"status\",\"fieldSearchType\":\"=\",\"dictType\":\"handleStatus\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"clothing\"}', 'D:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\job_apply.go;D:\\project\\gin-vue-admin\\server\\model\\clothing\\job_apply.go;D:\\project\\gin-vue-admin\\server\\model\\clothing\\request\\job_apply.go;D:\\project\\gin-vue-admin\\server\\router\\clothing\\job_apply.go;D:\\project\\gin-vue-admin\\server\\service\\clothing\\job_apply.go;D:\\project\\gin-vue-admin\\web\\src\\api\\jobApply.js;D:\\project\\gin-vue-admin\\web\\src\\view\\jobApply\\jobApplyForm.vue;D:\\project\\gin-vue-admin\\web\\src\\view\\jobApply\\jobApply.vue;', 'D:\\project\\gin-vue-admin\\server\\api\\v1\\clothing\\enter.go@ApiGroup@JobApplyApi;D:\\project\\gin-vue-admin\\server\\router\\clothing\\enter.go@RouterGroup@JobApplyRouter;D:\\project\\gin-vue-admin\\server\\service\\clothing\\enter.go@ServiceGroup@JobApplyService;', 'JobApply', '工单申请', '214;215;216;217;218;219;', 0);

-- ----------------------------
-- Table structure for sys_auto_codes
-- ----------------------------
DROP TABLE IF EXISTS `sys_auto_codes`;
CREATE TABLE `sys_auto_codes`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `package_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '包名',
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示名',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_auto_codes_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_auto_codes
-- ----------------------------
INSERT INTO `sys_auto_codes` VALUES (1, '2023-05-02 20:17:20.457', '2023-05-02 20:17:20.457', NULL, 'clothing', '', '');

-- ----------------------------
-- Table structure for sys_base_menu_btns
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_btns`;
CREATE TABLE `sys_base_menu_btns`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `sys_base_menu_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_btns_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_base_menu_btns
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menu_parameters
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menu_parameters`;
CREATE TABLE `sys_base_menu_parameters`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `sys_base_menu_id` bigint UNSIGNED NULL DEFAULT NULL,
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menu_parameters_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_base_menu_parameters
-- ----------------------------

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `menu_level` bigint UNSIGNED NULL DEFAULT NULL,
  `parent_id` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) NULL DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint NULL DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) NULL DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 35 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '0', 'dashboard', 'dashboard', 0, 'view/dashboard/index.vue', 1, '', 0, 0, '仪表盘', 'odometer', 0);
INSERT INTO `sys_base_menus` VALUES (2, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '0', 'about', 'about', 0, 'view/about/index.vue', 9, '', 0, 0, '关于我们', 'info-filled', 0);
INSERT INTO `sys_base_menus` VALUES (3, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '0', 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', 3, '', 0, 0, '超级管理员', 'user', 0);
INSERT INTO `sys_base_menus` VALUES (4, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '3', 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', 1, '', 0, 0, '角色管理', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (5, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '3', 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', 2, '', 1, 0, '菜单管理', 'tickets', 0);
INSERT INTO `sys_base_menus` VALUES (6, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '3', 'api', 'api', 0, 'view/superAdmin/api/api.vue', 3, '', 1, 0, 'api管理', 'platform', 0);
INSERT INTO `sys_base_menus` VALUES (7, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '3', 'user', 'user', 0, 'view/superAdmin/user/user.vue', 4, '', 0, 0, '用户管理', 'coordinate', 0);
INSERT INTO `sys_base_menus` VALUES (8, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '3', 'dictionary', 'dictionary', 0, 'view/superAdmin/dictionary/sysDictionary.vue', 5, '', 0, 0, '字典管理', 'notebook', 0);
INSERT INTO `sys_base_menus` VALUES (9, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '3', 'dictionaryDetail/:id', 'dictionaryDetail', 1, 'view/superAdmin/dictionary/sysDictionaryDetail.vue', 1, 'dictionary', 0, 0, '字典详情-${id}', 'list', 0);
INSERT INTO `sys_base_menus` VALUES (10, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '3', 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', 6, '', 0, 0, '操作历史', 'pie-chart', 0);
INSERT INTO `sys_base_menus` VALUES (11, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '0', 'person', 'person', 1, 'view/person/person.vue', 4, '', 0, 0, '个人信息', 'message', 0);
INSERT INTO `sys_base_menus` VALUES (12, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '0', 'example', 'example', 0, 'view/example/index.vue', 7, '', 0, 0, '示例文件', 'management', 0);
INSERT INTO `sys_base_menus` VALUES (13, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '12', 'upload', 'upload', 0, 'view/example/upload/upload.vue', 5, '', 0, 0, '媒体库（上传下载）', 'upload', 0);
INSERT INTO `sys_base_menus` VALUES (14, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '12', 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', 6, '', 0, 0, '断点续传', 'upload-filled', 0);
INSERT INTO `sys_base_menus` VALUES (15, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '12', 'customer', 'customer', 0, 'view/example/customer/customer.vue', 7, '', 0, 0, '客户列表（资源示例）', 'avatar', 0);
INSERT INTO `sys_base_menus` VALUES (16, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '0', 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', 5, '', 0, 0, '系统工具', 'tools', 0);
INSERT INTO `sys_base_menus` VALUES (17, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '16', 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', 1, '', 1, 0, '代码生成器', 'cpu', 0);
INSERT INTO `sys_base_menus` VALUES (18, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '16', 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', 2, '', 1, 0, '表单生成器', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (19, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '16', 'system', 'system', 0, 'view/systemTools/system/system.vue', 3, '', 0, 0, '系统配置', 'operation', 0);
INSERT INTO `sys_base_menus` VALUES (20, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '16', 'autoCodeAdmin', 'autoCodeAdmin', 0, 'view/systemTools/autoCodeAdmin/index.vue', 1, '', 0, 0, '自动化代码管理', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (21, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '16', 'autoCodeEdit/:id', 'autoCodeEdit', 1, 'view/systemTools/autoCode/index.vue', 0, '', 0, 0, '自动化代码-${id}', 'magic-stick', 0);
INSERT INTO `sys_base_menus` VALUES (22, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '16', 'autoPkg', 'autoPkg', 0, 'view/systemTools/autoPkg/autoPkg.vue', 0, '', 0, 0, '自动化package', 'folder', 0);
INSERT INTO `sys_base_menus` VALUES (23, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '0', 'https://www.gin-vue-admin.com', 'https://www.gin-vue-admin.com', 0, '/', 0, '', 0, 0, '官方网站', 'home-filled', 0);
INSERT INTO `sys_base_menus` VALUES (24, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '0', 'state', 'state', 0, 'view/system/state.vue', 8, '', 0, 0, '服务器状态', 'cloudy', 0);
INSERT INTO `sys_base_menus` VALUES (25, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '0', 'plugin', 'plugin', 0, 'view/routerHolder.vue', 6, '', 0, 0, '插件系统', 'cherry', 0);
INSERT INTO `sys_base_menus` VALUES (26, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '25', 'https://plugin.gin-vue-admin.com/', 'https://plugin.gin-vue-admin.com/', 0, 'https://plugin.gin-vue-admin.com/', 0, '', 0, 0, '插件市场', 'shop', 0);
INSERT INTO `sys_base_menus` VALUES (27, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '25', 'installPlugin', 'installPlugin', 0, 'view/systemTools/installPlugin/index.vue', 1, '', 0, 0, '插件安装', 'box', 0);
INSERT INTO `sys_base_menus` VALUES (28, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '25', 'autoPlug', 'autoPlug', 0, 'view/systemTools/autoPlug/autoPlug.vue', 2, '', 0, 0, '插件模板', 'folder', 0);
INSERT INTO `sys_base_menus` VALUES (29, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '25', 'plugin-email', 'plugin-email', 0, 'plugin/email/view/index.vue', 3, '', 0, 0, '邮件插件', 'message', 0);
INSERT INTO `sys_base_menus` VALUES (30, '2023-05-02 17:43:02.810', '2023-05-02 17:43:02.810', NULL, 0, '16', 'chatTable', 'chatTable', 0, 'view/chatgpt/chatTable.vue', 6, '', 0, 0, '万用表格', 'chat-dot-square', 0);
INSERT INTO `sys_base_menus` VALUES (31, '2023-05-02 21:12:56.337', '2023-05-02 21:12:56.337', NULL, 0, '0', 'appUserManage', 'appUserManage', 0, 'view/routerHolder.vue', 0, '', 0, 0, '用户管理', 'user', 0);
INSERT INTO `sys_base_menus` VALUES (32, '2023-05-02 21:14:08.036', '2023-05-02 21:15:41.524', NULL, 0, '31', 'appUser', 'appUser', 0, 'view/appUser/appUser.vue', 2, '', 0, 0, '用户列表', 'user-filled', 0);
INSERT INTO `sys_base_menus` VALUES (33, '2023-05-02 21:15:04.958', '2023-05-02 21:16:27.588', NULL, 0, '31', 'appRole', 'appRole', 0, 'view/appRole/appRole.vue', 0, '', 0, 0, '角色列表', 'aim', 0);
INSERT INTO `sys_base_menus` VALUES (34, '2023-05-02 21:16:17.353', '2023-05-02 21:16:17.353', NULL, 0, '31', 'company', 'company', 0, 'view/company/company.vue', 2, '', 0, 0, '公司列表', 'aim', 0);

-- ----------------------------
-- Table structure for sys_chat_gpt_options
-- ----------------------------
DROP TABLE IF EXISTS `sys_chat_gpt_options`;
CREATE TABLE `sys_chat_gpt_options`  (
  `sk` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_chat_gpt_options
-- ----------------------------

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`, `data_authority_id_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES (888, 888);
INSERT INTO `sys_data_authority_id` VALUES (888, 8881);
INSERT INTO `sys_data_authority_id` VALUES (888, 9528);
INSERT INTO `sys_data_authority_id` VALUES (9528, 8881);
INSERT INTO `sys_data_authority_id` VALUES (9528, 9528);

-- ----------------------------
-- Table structure for sys_dictionaries
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionaries`;
CREATE TABLE `sys_dictionaries`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionaries_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 14 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dictionaries
-- ----------------------------
INSERT INTO `sys_dictionaries` VALUES (1, '2023-05-02 17:42:58.056', '2023-05-02 17:42:58.426', NULL, '性别', 'gender', 1, '性别字典');
INSERT INTO `sys_dictionaries` VALUES (2, '2023-05-02 17:42:58.056', '2023-05-02 17:42:58.875', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (3, '2023-05-02 17:42:58.056', '2023-05-02 17:42:59.327', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2023-05-02 17:42:58.056', '2023-05-02 17:43:00.131', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (5, '2023-05-02 17:42:58.056', '2023-05-02 17:43:00.589', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (6, '2023-05-02 17:42:58.056', '2023-05-02 17:43:02.267', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');
INSERT INTO `sys_dictionaries` VALUES (7, '2023-05-02 20:26:25.165', '2023-05-02 20:26:25.165', NULL, '工作计算方式', 'workType', 1, '工作计算方式');
INSERT INTO `sys_dictionaries` VALUES (8, '2023-05-02 20:28:14.579', '2023-05-02 20:28:14.579', NULL, '状态', 'status', 1, '状态');
INSERT INTO `sys_dictionaries` VALUES (9, '2023-05-03 15:10:09.110', '2023-05-03 15:10:09.110', NULL, '裁剪单步骤', 'step', 1, '裁剪单步骤');
INSERT INTO `sys_dictionaries` VALUES (10, '2023-05-03 15:36:14.828', '2023-05-03 15:36:14.828', NULL, '工作类型', 'jobType', 1, '工作类型');
INSERT INTO `sys_dictionaries` VALUES (11, '2023-05-04 09:21:35.199', '2023-05-04 09:21:35.199', NULL, '处理状态', 'handleStatus', 1, '处理状态');
INSERT INTO `sys_dictionaries` VALUES (12, '2023-05-04 09:29:10.650', '2023-05-04 09:31:07.505', NULL, '消息类型', 'msgType', 1, '消息类型');
INSERT INTO `sys_dictionaries` VALUES (13, '2023-05-04 09:41:24.624', '2023-05-04 09:41:24.624', NULL, '媒体类型', 'mediaType', 1, '媒体类型');

-- ----------------------------
-- Table structure for sys_dictionary_details
-- ----------------------------
DROP TABLE IF EXISTS `sys_dictionary_details`;
CREATE TABLE `sys_dictionary_details`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `label` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '展示值',
  `value` bigint NULL DEFAULT NULL COMMENT '字典值',
  `status` tinyint(1) NULL DEFAULT NULL COMMENT '启用状态',
  `sort` bigint NULL DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_dictionary_details_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 48 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_dictionary_details
-- ----------------------------
INSERT INTO `sys_dictionary_details` VALUES (1, '2023-05-02 17:42:58.515', '2023-05-02 17:42:58.515', NULL, '男', 1, 1, 1, 1);
INSERT INTO `sys_dictionary_details` VALUES (2, '2023-05-02 17:42:58.515', '2023-05-02 17:42:58.515', NULL, '女', 2, 1, 2, 1);
INSERT INTO `sys_dictionary_details` VALUES (3, '2023-05-02 17:42:58.962', '2023-05-02 17:42:58.962', NULL, 'smallint', 1, 1, 1, 2);
INSERT INTO `sys_dictionary_details` VALUES (4, '2023-05-02 17:42:58.962', '2023-05-02 17:42:58.962', NULL, 'mediumint', 2, 1, 2, 2);
INSERT INTO `sys_dictionary_details` VALUES (5, '2023-05-02 17:42:58.962', '2023-05-02 17:42:58.962', NULL, 'int', 3, 1, 3, 2);
INSERT INTO `sys_dictionary_details` VALUES (6, '2023-05-02 17:42:58.962', '2023-05-02 17:42:58.962', NULL, 'bigint', 4, 1, 4, 2);
INSERT INTO `sys_dictionary_details` VALUES (7, '2023-05-02 17:42:59.414', '2023-05-02 17:42:59.414', NULL, 'date', 0, 1, 0, 3);
INSERT INTO `sys_dictionary_details` VALUES (8, '2023-05-02 17:42:59.414', '2023-05-02 17:42:59.414', NULL, 'time', 1, 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (9, '2023-05-02 17:42:59.414', '2023-05-02 17:42:59.414', NULL, 'year', 2, 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (10, '2023-05-02 17:42:59.414', '2023-05-02 17:42:59.414', NULL, 'datetime', 3, 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (11, '2023-05-02 17:42:59.414', '2023-05-02 17:42:59.414', NULL, 'timestamp', 5, 1, 5, 3);
INSERT INTO `sys_dictionary_details` VALUES (12, '2023-05-02 17:43:00.219', '2023-05-02 17:43:00.219', NULL, 'float', 0, 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (13, '2023-05-02 17:43:00.219', '2023-05-02 17:43:00.219', NULL, 'double', 1, 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (14, '2023-05-02 17:43:00.219', '2023-05-02 17:43:00.219', NULL, 'decimal', 2, 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (15, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'char', 0, 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (16, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'varchar', 1, 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (17, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'tinyblob', 2, 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (18, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'tinytext', 3, 1, 3, 5);
INSERT INTO `sys_dictionary_details` VALUES (19, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'text', 4, 1, 4, 5);
INSERT INTO `sys_dictionary_details` VALUES (20, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'blob', 5, 1, 5, 5);
INSERT INTO `sys_dictionary_details` VALUES (21, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'mediumblob', 6, 1, 6, 5);
INSERT INTO `sys_dictionary_details` VALUES (22, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'mediumtext', 7, 1, 7, 5);
INSERT INTO `sys_dictionary_details` VALUES (23, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'longblob', 8, 1, 8, 5);
INSERT INTO `sys_dictionary_details` VALUES (24, '2023-05-02 17:43:00.678', '2023-05-02 17:43:00.678', NULL, 'longtext', 9, 1, 9, 5);
INSERT INTO `sys_dictionary_details` VALUES (25, '2023-05-02 17:43:02.357', '2023-05-02 17:43:02.357', NULL, 'tinyint', 0, 1, 0, 6);
INSERT INTO `sys_dictionary_details` VALUES (26, '2023-05-02 20:26:38.719', '2023-05-02 20:26:38.719', NULL, '工序', 1, 1, 1, 7);
INSERT INTO `sys_dictionary_details` VALUES (27, '2023-05-02 20:26:50.882', '2023-05-02 20:26:50.882', NULL, '整件', 2, 1, 2, 7);
INSERT INTO `sys_dictionary_details` VALUES (28, '2023-05-02 20:28:34.379', '2023-05-02 20:29:17.881', NULL, '开启', 1, 1, 1, 8);
INSERT INTO `sys_dictionary_details` VALUES (29, '2023-05-02 20:28:53.009', '2023-05-02 20:29:23.079', NULL, '停用', 0, 1, 0, 8);
INSERT INTO `sys_dictionary_details` VALUES (30, '2023-05-03 15:14:10.825', '2023-05-03 15:14:10.825', NULL, '待处理', 1, 1, 1, 9);
INSERT INTO `sys_dictionary_details` VALUES (31, '2023-05-03 15:14:31.469', '2023-05-03 15:14:31.469', NULL, '处理中', 2, 1, 2, 9);
INSERT INTO `sys_dictionary_details` VALUES (32, '2023-05-03 15:14:43.573', '2023-05-03 15:14:43.573', NULL, '待审核', 3, 1, 3, 9);
INSERT INTO `sys_dictionary_details` VALUES (33, '2023-05-03 15:14:56.795', '2023-05-03 15:14:56.795', NULL, '已完成', 4, 1, 4, 9);
INSERT INTO `sys_dictionary_details` VALUES (34, '2023-05-03 15:36:40.111', '2023-05-03 15:36:40.111', NULL, '工序', 1, 1, 1, 10);
INSERT INTO `sys_dictionary_details` VALUES (35, '2023-05-03 15:36:48.800', '2023-05-03 15:36:48.800', NULL, '成衣', 2, 1, 2, 10);
INSERT INTO `sys_dictionary_details` VALUES (36, '2023-05-04 09:22:00.222', '2023-05-04 09:22:00.222', NULL, '待处理', 0, 1, 0, 11);
INSERT INTO `sys_dictionary_details` VALUES (37, '2023-05-04 09:22:22.952', '2023-05-08 16:49:14.838', NULL, '已通过', 1, 1, 1, 11);
INSERT INTO `sys_dictionary_details` VALUES (38, '2023-05-04 09:29:32.242', '2023-05-08 13:37:58.135', NULL, '进组申请', 2, 1, 2, 12);
INSERT INTO `sys_dictionary_details` VALUES (39, '2023-05-04 09:30:02.489', '2023-05-08 13:38:06.349', NULL, '工单问题', 4, 1, 4, 12);
INSERT INTO `sys_dictionary_details` VALUES (40, '2023-05-04 09:30:17.475', '2023-05-08 13:38:11.988', NULL, '工单审核', 5, 1, 5, 12);
INSERT INTO `sys_dictionary_details` VALUES (41, '2023-05-04 09:41:36.101', '2023-05-04 09:41:36.101', NULL, '图片', 1, 1, 1, 13);
INSERT INTO `sys_dictionary_details` VALUES (42, '2023-05-04 09:41:45.135', '2023-05-04 09:41:45.135', NULL, '视频', 2, 1, 2, 13);
INSERT INTO `sys_dictionary_details` VALUES (43, '2023-05-04 11:33:25.288', '2023-05-08 13:37:37.855', NULL, '裁缝、组长申请', 1, 1, 1, 12);
INSERT INTO `sys_dictionary_details` VALUES (44, '2023-05-04 11:33:41.396', '2023-05-04 11:33:41.396', '2023-05-05 10:45:59.102', '组长申请', 2, 1, 2, 12);
INSERT INTO `sys_dictionary_details` VALUES (45, '2023-05-08 09:06:08.670', '2023-05-08 13:38:02.892', NULL, '工单申请', 3, 1, 3, 12);
INSERT INTO `sys_dictionary_details` VALUES (46, '2023-05-08 09:46:06.961', '2023-05-08 09:46:28.561', '2023-05-08 13:37:52.432', '组长申请', 2, 1, 2, 12);
INSERT INTO `sys_dictionary_details` VALUES (47, '2023-05-08 16:49:22.415', '2023-05-08 16:49:22.415', NULL, '已拒绝', 2, 1, 2, 11);

-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `ip` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '请求路径',
  `status` bigint NULL DEFAULT NULL COMMENT '请求状态',
  `latency` bigint NULL DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '错误信息',
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '请求Body',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL COMMENT '响应Body',
  `user_id` bigint UNSIGNED NULL DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 77 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_operation_records
-- ----------------------------
INSERT INTO `sys_operation_records` VALUES (1, '2023-05-02 17:51:08.895', '2023-05-02 17:51:08.895', NULL, '127.0.0.1', 'GET', '/chatGpt/getSK', 200, 45035500, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{}', '{\"code\":0,\"data\":{\"ok\":false},\"msg\":\"无sk或获取失败\"}', 1);
INSERT INTO `sys_operation_records` VALUES (2, '2023-05-02 17:51:28.600', '2023-05-02 17:51:28.600', NULL, '127.0.0.1', 'POST', '/system/getSystemConfig', 200, 567200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '', '{\"code\":0,\"data\":{\"config\":{\"jwt\":{\"signing-key\":\"4b611845-6352-4cb2-a65c-db25bdc43383\",\"expires-time\":\"7d\",\"buffer-time\":\"1d\",\"issuer\":\"qmPlus\"},\"zap\":{\"level\":\"info\",\"prefix\":\"[github.com/flipped-aurora/gin-vue-admin/server]\",\"format\":\"console\",\"director\":\"log\",\"encode-level\":\"LowercaseColorLevelEncoder\",\"stacktrace-key\":\"stacktrace\",\"max-age\":0,\"show-line\":true,\"log-in-console\":true},\"redis\":{\"db\":0,\"addr\":\"127.0.0.1:6379\",\"password\":\"\"},\"email\":{\"to\":\"xxx@qq.com\",\"port\":465,\"from\":\"xxx@163.com\",\"host\":\"smtp.163.com\",\"is-ssl\":true,\"secret\":\"xxx\",\"nickname\":\"test\"},\"system\":{\"env\":\"public\",\"addr\":8889,\"db-type\":\"mysql\",\"oss-type\":\"local\",\"use-multipoint\":false,\"use-redis\":false,\"iplimit-count\":15000,\"iplimit-time\":3600,\"router-prefix\":\"\"},\"captcha\":{\"key-long\":6,\"img-width\":240,\"img-height\":80,\"open-captcha\":0,\"open-captcha-timeout\":3600},\"autocode\":{\"transfer-restart\":true,\"root\":\"d:\\\\project\\\\clothing\",\"server\":\"/server\",\"server-api\":\"/api/v1/%s\",\"server-plug\":\"/plugin/%s\",\"server-initialize\":\"/initialize\",\"server-model\":\"/model/%s\",\"server-request\":\"/model/%s/request/\",\"server-router\":\"/router/%s\",\"server-service\":\"/service/%s\",\"web\":\"/web/src\",\"web-api\":\"/api\",\"web-form\":\"/view\",\"web-table\":\"/view\"},\"mysql\":{\"path\":\"39.105.214.148\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"db-name\":\"gva\",\"username\":\"root\",\"password\":\"local123\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"error\",\"log-zap\":false},\"mssql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"pgsql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"oracle\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"db-list\":[{\"disable\":true,\"type\":\"\",\"alias-name\":\"\",\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false}],\"local\":{\"path\":\"uploads/file\",\"store-path\":\"uploads/file\"},\"qiniu\":{\"zone\":\"ZoneHuaDong\",\"bucket\":\"\",\"img-path\":\"\",\"use-https\":false,\"access-key\":\"\",\"secret-key\":\"\",\"use-cdn-domains\":false},\"aliyun-oss\":{\"endpoint\":\"yourEndpoint\",\"access-key-id\":\"yourAccessKeyId\",\"access-key-secret\":\"yourAccessKeySecret\",\"bucket-name\":\"yourBucketName\",\"bucket-url\":\"yourBucketUrl\",\"base-path\":\"yourBasePath\"},\"hua-wei-obs\":{\"path\":\"you-path\",\"bucket\":\"you-bucket\",\"endpoint\":\"you-endpoint\",\"access-key\":\"you-access-key\",\"secret-key\":\"you-secret-key\"},\"tencent-cos\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"aws-s3\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"endpoint\":\"\",\"s3-force-path-style\":false,\"disable-ssl\":false,\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"excel\":{\"dir\":\"./resource/excel/\"},\"timer\":{\"start\":true,\"spec\":\"@daily\",\"with_seconds\":false,\"detail\":[{\"tableName\":\"sys_operation_records\",\"compareField\":\"created_at\",\"interval\":\"2160h\"},{\"tableName\":\"jwt_blacklists\",\"compareField\":\"created_at\",\"interval\":\"168h\"}]},\"cors\":{\"mode\":\"strict-whitelist\",\"whitelist\":[{\"allow-origin\":\"example1.com\",\"allow-methods\":\"POST, GET\",\"allow-headers\":\"Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true},{\"allow-origin\":\"example2.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true}]}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (3, '2023-05-02 18:07:53.442', '2023-05-02 18:07:53.442', NULL, '127.0.0.1', 'POST', '/system/getSystemConfig', 200, 1000200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '', '{\"code\":0,\"data\":{\"config\":{\"jwt\":{\"signing-key\":\"4b611845-6352-4cb2-a65c-db25bdc43383\",\"expires-time\":\"7d\",\"buffer-time\":\"1d\",\"issuer\":\"qmPlus\"},\"zap\":{\"level\":\"info\",\"prefix\":\"[github.com/flipped-aurora/gin-vue-admin/server]\",\"format\":\"console\",\"director\":\"log\",\"encode-level\":\"LowercaseColorLevelEncoder\",\"stacktrace-key\":\"stacktrace\",\"max-age\":0,\"show-line\":true,\"log-in-console\":true},\"redis\":{\"db\":0,\"addr\":\"127.0.0.1:6379\",\"password\":\"\"},\"email\":{\"to\":\"xxx@qq.com\",\"port\":465,\"from\":\"xxx@163.com\",\"host\":\"smtp.163.com\",\"is-ssl\":true,\"secret\":\"xxx\",\"nickname\":\"test\"},\"system\":{\"env\":\"public\",\"addr\":8889,\"db-type\":\"mysql\",\"oss-type\":\"local\",\"use-multipoint\":false,\"use-redis\":false,\"iplimit-count\":15000,\"iplimit-time\":3600,\"router-prefix\":\"\"},\"captcha\":{\"key-long\":6,\"img-width\":240,\"img-height\":80,\"open-captcha\":0,\"open-captcha-timeout\":3600},\"autocode\":{\"transfer-restart\":true,\"root\":\"d:\\\\project\\\\clothing\",\"server\":\"/server\",\"server-api\":\"/api/v1/%s\",\"server-plug\":\"/plugin/%s\",\"server-initialize\":\"/initialize\",\"server-model\":\"/model/%s\",\"server-request\":\"/model/%s/request/\",\"server-router\":\"/router/%s\",\"server-service\":\"/service/%s\",\"web\":\"/web/src\",\"web-api\":\"/api\",\"web-form\":\"/view\",\"web-table\":\"/view\"},\"mysql\":{\"path\":\"39.105.214.148\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"db-name\":\"gva\",\"username\":\"root\",\"password\":\"local123\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"error\",\"log-zap\":false},\"mssql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"pgsql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"oracle\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"db-list\":[{\"disable\":true,\"type\":\"\",\"alias-name\":\"\",\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false}],\"local\":{\"path\":\"uploads/file\",\"store-path\":\"uploads/file\"},\"qiniu\":{\"zone\":\"ZoneHuaDong\",\"bucket\":\"\",\"img-path\":\"\",\"use-https\":false,\"access-key\":\"\",\"secret-key\":\"\",\"use-cdn-domains\":false},\"aliyun-oss\":{\"endpoint\":\"yourEndpoint\",\"access-key-id\":\"yourAccessKeyId\",\"access-key-secret\":\"yourAccessKeySecret\",\"bucket-name\":\"yourBucketName\",\"bucket-url\":\"yourBucketUrl\",\"base-path\":\"yourBasePath\"},\"hua-wei-obs\":{\"path\":\"you-path\",\"bucket\":\"you-bucket\",\"endpoint\":\"you-endpoint\",\"access-key\":\"you-access-key\",\"secret-key\":\"you-secret-key\"},\"tencent-cos\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"aws-s3\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"endpoint\":\"\",\"s3-force-path-style\":false,\"disable-ssl\":false,\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"excel\":{\"dir\":\"./resource/excel/\"},\"timer\":{\"start\":true,\"spec\":\"@daily\",\"with_seconds\":false,\"detail\":[{\"tableName\":\"sys_operation_records\",\"compareField\":\"created_at\",\"interval\":\"2160h\"},{\"tableName\":\"jwt_blacklists\",\"compareField\":\"created_at\",\"interval\":\"168h\"}]},\"cors\":{\"mode\":\"strict-whitelist\",\"whitelist\":[{\"allow-origin\":\"example1.com\",\"allow-methods\":\"POST, GET\",\"allow-headers\":\"Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true},{\"allow-origin\":\"example2.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true}]}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (4, '2023-05-02 20:26:25.330', '2023-05-02 20:26:25.330', NULL, '127.0.0.1', 'POST', '/sysDictionary/createSysDictionary', 200, 246370800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"工作计算方式\",\"type\":\"workType\",\"status\":true,\"desc\":\"工作计算方式\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (5, '2023-05-02 20:26:38.883', '2023-05-02 20:26:38.883', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 163498100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"工序\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":7}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (6, '2023-05-02 20:26:51.046', '2023-05-02 20:26:51.046', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 163357000, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"整件\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":7}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (7, '2023-05-02 20:28:14.743', '2023-05-02 20:28:14.743', NULL, '127.0.0.1', 'POST', '/sysDictionary/createSysDictionary', 200, 245545800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"状态\",\"type\":\"status\",\"status\":true,\"desc\":\"状态\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (8, '2023-05-02 20:28:34.542', '2023-05-02 20:28:34.542', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 163333000, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"开启\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":8}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (9, '2023-05-02 20:28:53.173', '2023-05-02 20:28:53.173', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 163528400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"停用\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":8}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (10, '2023-05-02 20:28:59.487', '2023-05-02 20:28:59.487', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 163846400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":28,\"CreatedAt\":\"2023-05-02T20:28:34.379+08:00\",\"UpdatedAt\":\"2023-05-02T20:28:34.379+08:00\",\"label\":\"开启\",\"value\":0,\"status\":true,\"sort\":0,\"sysDictionaryID\":8}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (11, '2023-05-02 20:29:18.046', '2023-05-02 20:29:18.046', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 163935100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":28,\"CreatedAt\":\"2023-05-02T20:28:34.379+08:00\",\"UpdatedAt\":\"2023-05-02T20:28:59.322+08:00\",\"label\":\"开启\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":8}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (12, '2023-05-02 20:29:23.243', '2023-05-02 20:29:23.243', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 163388300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":29,\"CreatedAt\":\"2023-05-02T20:28:53.009+08:00\",\"UpdatedAt\":\"2023-05-02T20:28:53.009+08:00\",\"label\":\"停用\",\"value\":0,\"status\":true,\"sort\":0,\"sysDictionaryID\":8}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (13, '2023-05-02 21:12:56.496', '2023-05-02 21:12:56.496', NULL, '127.0.0.1', 'POST', '/menu/addBaseMenu', 200, 238006900, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":0,\"path\":\"appUserManage\",\"name\":\"appUserManage\",\"hidden\":false,\"parentId\":\"0\",\"component\":\"view/routerHolder.vue\",\"meta\":{\"activeName\":\"\",\"title\":\"用户管理\",\"icon\":\"user\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false},\"parameters\":[],\"menuBtn\":[]}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (14, '2023-05-02 21:14:08.196', '2023-05-02 21:14:08.196', NULL, '127.0.0.1', 'POST', '/menu/addBaseMenu', 200, 238622500, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":0,\"path\":\"appUser\",\"name\":\"appUser\",\"hidden\":false,\"parentId\":\"31\",\"component\":\"view/appUser/appUser.vue\",\"meta\":{\"title\":\"用户列表\",\"icon\":\"user-filled\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false},\"sort\":0}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (15, '2023-05-02 21:15:05.118', '2023-05-02 21:15:05.118', NULL, '127.0.0.1', 'POST', '/menu/addBaseMenu', 200, 239785000, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":0,\"path\":\"appRole\",\"name\":\"appRole\",\"hidden\":false,\"parentId\":\"0\",\"component\":\"view/appRole/appRole.vue\",\"meta\":{\"title\":\"角色管理\",\"icon\":\"aim\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false},\"sort\":0}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (16, '2023-05-02 21:15:33.323', '2023-05-02 21:15:33.323', NULL, '127.0.0.1', 'POST', '/menu/updateBaseMenu', 200, 397274600, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":33,\"CreatedAt\":\"2023-05-02T21:15:04.958+08:00\",\"UpdatedAt\":\"2023-05-02T21:15:04.958+08:00\",\"parentId\":\"31\",\"path\":\"appRole\",\"name\":\"appRole\",\"hidden\":false,\"component\":\"view/appRole/appRole.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (17, '2023-05-02 21:15:41.682', '2023-05-02 21:15:41.682', NULL, '127.0.0.1', 'POST', '/menu/updateBaseMenu', 200, 396139000, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":32,\"CreatedAt\":\"2023-05-02T21:14:08.036+08:00\",\"UpdatedAt\":\"2023-05-02T21:14:08.036+08:00\",\"parentId\":\"31\",\"path\":\"appUser\",\"name\":\"appUser\",\"hidden\":false,\"component\":\"view/appUser/appUser.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户列表\",\"icon\":\"user-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (18, '2023-05-02 21:16:17.512', '2023-05-02 21:16:17.512', NULL, '127.0.0.1', 'POST', '/menu/addBaseMenu', 200, 238105400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":0,\"path\":\"company\",\"name\":\"company\",\"hidden\":false,\"parentId\":\"31\",\"component\":\"view/company/company.vue\",\"meta\":{\"title\":\"公司列表\",\"icon\":\"aim\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false},\"sort\":2}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (19, '2023-05-02 21:16:27.746', '2023-05-02 21:16:27.746', NULL, '127.0.0.1', 'POST', '/menu/updateBaseMenu', 200, 396155600, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":33,\"CreatedAt\":\"2023-05-02T21:15:04.958+08:00\",\"UpdatedAt\":\"2023-05-02T21:15:33.164+08:00\",\"parentId\":\"31\",\"path\":\"appRole\",\"name\":\"appRole\",\"hidden\":false,\"component\":\"view/appRole/appRole.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色列表\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (20, '2023-05-02 21:16:40.519', '2023-05-02 21:16:40.519', NULL, '127.0.0.1', 'POST', '/menu/addMenuAuthority', 200, 759001800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"menus\":[{\"ID\":23,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"0\",\"path\":\"https://www.gin-vue-admin.com\",\"name\":\"https://www.gin-vue-admin.com\",\"hidden\":false,\"component\":\"/\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"官方网站\",\"icon\":\"home-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":31,\"CreatedAt\":\"2023-05-02T21:12:56.337+08:00\",\"UpdatedAt\":\"2023-05-02T21:12:56.337+08:00\",\"parentId\":\"0\",\"path\":\"appUserManage\",\"name\":\"appUserManage\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"user\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":33,\"CreatedAt\":\"2023-05-02T21:15:04.958+08:00\",\"UpdatedAt\":\"2023-05-02T21:16:27.588+08:00\",\"parentId\":\"31\",\"path\":\"appRole\",\"name\":\"appRole\",\"hidden\":false,\"component\":\"view/appRole/appRole.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色列表\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":32,\"CreatedAt\":\"2023-05-02T21:14:08.036+08:00\",\"UpdatedAt\":\"2023-05-02T21:15:41.524+08:00\",\"parentId\":\"31\",\"path\":\"appUser\",\"name\":\"appUser\",\"hidden\":false,\"component\":\"view/appUser/appUser.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户列表\",\"icon\":\"user-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2023-05-02T21:16:17.353+08:00\",\"UpdatedAt\":\"2023-05-02T21:16:17.353+08:00\",\"parentId\":\"31\",\"path\":\"company\",\"name\":\"company\",\"hidden\":false,\"component\":\"view/company/company.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"公司列表\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":33,\"CreatedAt\":\"2023-05-02T21:15:04.958+08:00\",\"UpdatedAt\":\"2023-05-02T21:16:27.588+08:00\",\"parentId\":\"31\",\"path\":\"appRole\",\"name\":\"appRole\",\"hidden\":false,\"component\":\"view/appRole/appRole.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色列表\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":32,\"CreatedAt\":\"2023-05-02T21:14:08.036+08:00\",\"UpdatedAt\":\"2023-05-02T21:15:41.524+08:00\",\"parentId\":\"31\",\"path\":\"appUser\",\"name\":\"appUser\",\"hidden\":false,\"component\":\"view/appUser/appUser.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户列表\",\"icon\":\"user-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":34,\"CreatedAt\":\"2023-05-02T21:16:17.353+08:00\",\"UpdatedAt\":\"2023-05-02T21:16:17.353+08:00\",\"parentId\":\"31\",\"path\":\"company\",\"name\":\"company\",\"hidden\":false,\"component\":\"view/company/company.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"公司列表\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":1,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"0\",\"path\":\"dashboard\",\"name\":\"dashboard\",\"hidden\":false,\"component\":\"view/dashboard/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"仪表盘\",\"icon\":\"odometer\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":3,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"0\",\"path\":\"admin\",\"name\":\"superAdmin\",\"hidden\":false,\"component\":\"view/superAdmin/index.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"超级管理员\",\"icon\":\"user\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":4,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":9,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"activeName\":\"dictionary\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情-${id}\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":10,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":4,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"authority\",\"name\":\"authority\",\"hidden\":false,\"component\":\"view/superAdmin/authority/authority.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"角色管理\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":9,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"dictionaryDetail/:id\",\"name\":\"dictionaryDetail\",\"hidden\":true,\"component\":\"view/superAdmin/dictionary/sysDictionaryDetail.vue\",\"sort\":1,\"meta\":{\"activeName\":\"dictionary\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典详情-${id}\",\"icon\":\"list\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":5,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"menu\",\"name\":\"menu\",\"hidden\":false,\"component\":\"view/superAdmin/menu/menu.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"菜单管理\",\"icon\":\"tickets\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":6,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"api\",\"name\":\"api\",\"hidden\":false,\"component\":\"view/superAdmin/api/api.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"api管理\",\"icon\":\"platform\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":7,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"user\",\"name\":\"user\",\"hidden\":false,\"component\":\"view/superAdmin/user/user.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"用户管理\",\"icon\":\"coordinate\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":8,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"dictionary\",\"name\":\"dictionary\",\"hidden\":false,\"component\":\"view/superAdmin/dictionary/sysDictionary.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"字典管理\",\"icon\":\"notebook\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":10,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"3\",\"path\":\"operation\",\"name\":\"operation\",\"hidden\":false,\"component\":\"view/superAdmin/operation/sysOperationRecord.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"操作历史\",\"icon\":\"pie-chart\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":11,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"0\",\"path\":\"person\",\"name\":\"person\",\"hidden\":true,\"component\":\"view/person/person.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"个人信息\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":16,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"0\",\"path\":\"systemTools\",\"name\":\"systemTools\",\"hidden\":false,\"component\":\"view/systemTools/index.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统工具\",\"icon\":\"tools\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":21,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码-${id}\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":22,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"autoPkg\",\"name\":\"autoPkg\",\"hidden\":false,\"component\":\"view/systemTools/autoPkg/autoPkg.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化package\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":30,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"chatTable\",\"name\":\"chatTable\",\"hidden\":false,\"component\":\"view/chatgpt/chatTable.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"万用表格\",\"icon\":\"chat-dot-square\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":21,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"autoCodeEdit/:id\",\"name\":\"autoCodeEdit\",\"hidden\":true,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码-${id}\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":22,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"autoPkg\",\"name\":\"autoPkg\",\"hidden\":false,\"component\":\"view/systemTools/autoPkg/autoPkg.vue\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化package\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":17,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"autoCode\",\"name\":\"autoCode\",\"hidden\":false,\"component\":\"view/systemTools/autoCode/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"代码生成器\",\"icon\":\"cpu\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":20,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"autoCodeAdmin\",\"name\":\"autoCodeAdmin\",\"hidden\":false,\"component\":\"view/systemTools/autoCodeAdmin/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"自动化代码管理\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":18,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"formCreate\",\"name\":\"formCreate\",\"hidden\":false,\"component\":\"view/systemTools/formCreate/index.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":true,\"defaultMenu\":false,\"title\":\"表单生成器\",\"icon\":\"magic-stick\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":19,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"system\",\"name\":\"system\",\"hidden\":false,\"component\":\"view/systemTools/system/system.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"系统配置\",\"icon\":\"operation\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":30,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"16\",\"path\":\"chatTable\",\"name\":\"chatTable\",\"hidden\":false,\"component\":\"view/chatgpt/chatTable.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"万用表格\",\"icon\":\"chat-dot-square\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":25,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"0\",\"path\":\"plugin\",\"name\":\"plugin\",\"hidden\":false,\"component\":\"view/routerHolder.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件系统\",\"icon\":\"cherry\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":26,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"25\",\"path\":\"https://plugin.gin-vue-admin.com/\",\"name\":\"https://plugin.gin-vue-admin.com/\",\"hidden\":false,\"component\":\"https://plugin.gin-vue-admin.com/\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件市场\",\"icon\":\"shop\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":27,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"25\",\"path\":\"installPlugin\",\"name\":\"installPlugin\",\"hidden\":false,\"component\":\"view/systemTools/installPlugin/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件安装\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":28,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"25\",\"path\":\"autoPlug\",\"name\":\"autoPlug\",\"hidden\":false,\"component\":\"view/systemTools/autoPlug/autoPlug.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件模板\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":29,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"25\",\"path\":\"plugin-email\",\"name\":\"plugin-email\",\"hidden\":false,\"component\":\"plugin/email/view/index.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"邮件插件\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":26,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"25\",\"path\":\"https://plugin.gin-vue-admin.com/\",\"name\":\"https://plugin.gin-vue-admin.com/\",\"hidden\":false,\"component\":\"https://plugin.gin-vue-admin.com/\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件市场\",\"icon\":\"shop\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":27,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"25\",\"path\":\"installPlugin\",\"name\":\"installPlugin\",\"hidden\":false,\"component\":\"view/systemTools/installPlugin/index.vue\",\"sort\":1,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件安装\",\"icon\":\"box\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":28,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"25\",\"path\":\"autoPlug\",\"name\":\"autoPlug\",\"hidden\":false,\"component\":\"view/systemTools/autoPlug/autoPlug.vue\",\"sort\":2,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"插件模板\",\"icon\":\"folder\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":29,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"25\",\"path\":\"plugin-email\",\"name\":\"plugin-email\",\"hidden\":false,\"component\":\"plugin/email/view/index.vue\",\"sort\":3,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"邮件插件\",\"icon\":\"message\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":12,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"0\",\"path\":\"example\",\"name\":\"example\",\"hidden\":false,\"component\":\"view/example/index.vue\",\"sort\":7,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"示例文件\",\"icon\":\"management\",\"closeTab\":false},\"authoritys\":null,\"children\":[{\"ID\":13,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"12\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":14,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"12\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"12\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"parameters\":[],\"menuBtn\":[]},{\"ID\":13,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"12\",\"path\":\"upload\",\"name\":\"upload\",\"hidden\":false,\"component\":\"view/example/upload/upload.vue\",\"sort\":5,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"媒体库（上传下载）\",\"icon\":\"upload\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":14,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"12\",\"path\":\"breakpoint\",\"name\":\"breakpoint\",\"hidden\":false,\"component\":\"view/example/breakpoint/breakpoint.vue\",\"sort\":6,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"断点续传\",\"icon\":\"upload-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":15,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"12\",\"path\":\"customer\",\"name\":\"customer\",\"hidden\":false,\"component\":\"view/example/customer/customer.vue\",\"sort\":7,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"客户列表（资源示例）\",\"icon\":\"avatar\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":24,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"0\",\"path\":\"state\",\"name\":\"state\",\"hidden\":false,\"component\":\"view/system/state.vue\",\"sort\":8,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"服务器状态\",\"icon\":\"cloudy\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]},{\"ID\":2,\"CreatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"UpdatedAt\":\"2023-05-02T17:43:02.81+08:00\",\"parentId\":\"0\",\"path\":\"about\",\"name\":\"about\",\"hidden\":false,\"component\":\"view/about/index.vue\",\"sort\":9,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"关于我们\",\"icon\":\"info-filled\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}],\"authorityId\":888}', '{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (21, '2023-05-02 21:16:53.086', '2023-05-02 21:16:53.086', NULL, '127.0.0.1', 'POST', '/casbin/updateCasbin', 200, 362621000, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"authorityId\":888,\"casbinInfos\":[{\"path\":\"/base/login\",\"method\":\"POST\"},{\"path\":\"/jwt/jsonInBlacklist\",\"method\":\"POST\"},{\"path\":\"/user/deleteUser\",\"method\":\"DELETE\"},{\"path\":\"/user/admin_register\",\"method\":\"POST\"},{\"path\":\"/user/getUserList\",\"method\":\"POST\"},{\"path\":\"/user/setUserInfo\",\"method\":\"PUT\"},{\"path\":\"/user/setSelfInfo\",\"method\":\"PUT\"},{\"path\":\"/user/getUserInfo\",\"method\":\"GET\"},{\"path\":\"/user/setUserAuthorities\",\"method\":\"POST\"},{\"path\":\"/user/changePassword\",\"method\":\"POST\"},{\"path\":\"/user/setUserAuthority\",\"method\":\"POST\"},{\"path\":\"/user/resetPassword\",\"method\":\"POST\"},{\"path\":\"/api/createApi\",\"method\":\"POST\"},{\"path\":\"/api/deleteApi\",\"method\":\"POST\"},{\"path\":\"/api/updateApi\",\"method\":\"POST\"},{\"path\":\"/api/getApiList\",\"method\":\"POST\"},{\"path\":\"/api/getAllApis\",\"method\":\"POST\"},{\"path\":\"/api/getApiById\",\"method\":\"POST\"},{\"path\":\"/api/deleteApisByIds\",\"method\":\"DELETE\"},{\"path\":\"/authority/copyAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/createAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/deleteAuthority\",\"method\":\"POST\"},{\"path\":\"/authority/updateAuthority\",\"method\":\"PUT\"},{\"path\":\"/authority/getAuthorityList\",\"method\":\"POST\"},{\"path\":\"/authority/setDataAuthority\",\"method\":\"POST\"},{\"path\":\"/casbin/updateCasbin\",\"method\":\"POST\"},{\"path\":\"/casbin/getPolicyPathByAuthorityId\",\"method\":\"POST\"},{\"path\":\"/menu/addBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getMenu\",\"method\":\"POST\"},{\"path\":\"/menu/deleteBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/updateBaseMenu\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuById\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuList\",\"method\":\"POST\"},{\"path\":\"/menu/getBaseMenuTree\",\"method\":\"POST\"},{\"path\":\"/menu/getMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/menu/addMenuAuthority\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/findFile\",\"method\":\"GET\"},{\"path\":\"/fileUploadAndDownload/breakpointContinue\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/breakpointContinueFinish\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/removeChunk\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/upload\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/deleteFile\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/editFileName\",\"method\":\"POST\"},{\"path\":\"/fileUploadAndDownload/getFileList\",\"method\":\"POST\"},{\"path\":\"/system/getServerInfo\",\"method\":\"POST\"},{\"path\":\"/system/getSystemConfig\",\"method\":\"POST\"},{\"path\":\"/system/setSystemConfig\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"PUT\"},{\"path\":\"/customer/customer\",\"method\":\"POST\"},{\"path\":\"/customer/customer\",\"method\":\"DELETE\"},{\"path\":\"/customer/customer\",\"method\":\"GET\"},{\"path\":\"/customer/customerList\",\"method\":\"GET\"},{\"path\":\"/autoCode/getDB\",\"method\":\"GET\"},{\"path\":\"/autoCode/getTables\",\"method\":\"GET\"},{\"path\":\"/autoCode/createTemp\",\"method\":\"POST\"},{\"path\":\"/autoCode/preview\",\"method\":\"POST\"},{\"path\":\"/autoCode/getColumn\",\"method\":\"GET\"},{\"path\":\"/autoCode/createPlug\",\"method\":\"POST\"},{\"path\":\"/autoCode/installPlugin\",\"method\":\"POST\"},{\"path\":\"/autoCode/createPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/getPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/delPackage\",\"method\":\"POST\"},{\"path\":\"/autoCode/getMeta\",\"method\":\"POST\"},{\"path\":\"/autoCode/rollback\",\"method\":\"POST\"},{\"path\":\"/autoCode/getSysHistory\",\"method\":\"POST\"},{\"path\":\"/autoCode/delSysHistory\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/updateSysDictionaryDetail\",\"method\":\"PUT\"},{\"path\":\"/sysDictionaryDetail/createSysDictionaryDetail\",\"method\":\"POST\"},{\"path\":\"/sysDictionaryDetail/deleteSysDictionaryDetail\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionaryDetail/findSysDictionaryDetail\",\"method\":\"GET\"},{\"path\":\"/sysDictionaryDetail/getSysDictionaryDetailList\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/createSysDictionary\",\"method\":\"POST\"},{\"path\":\"/sysDictionary/deleteSysDictionary\",\"method\":\"DELETE\"},{\"path\":\"/sysDictionary/updateSysDictionary\",\"method\":\"PUT\"},{\"path\":\"/sysDictionary/findSysDictionary\",\"method\":\"GET\"},{\"path\":\"/sysDictionary/getSysDictionaryList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/createSysOperationRecord\",\"method\":\"POST\"},{\"path\":\"/sysOperationRecord/findSysOperationRecord\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/getSysOperationRecordList\",\"method\":\"GET\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecord\",\"method\":\"DELETE\"},{\"path\":\"/sysOperationRecord/deleteSysOperationRecordByIds\",\"method\":\"DELETE\"},{\"path\":\"/simpleUploader/upload\",\"method\":\"POST\"},{\"path\":\"/simpleUploader/checkFileMd5\",\"method\":\"GET\"},{\"path\":\"/simpleUploader/mergeFileMd5\",\"method\":\"GET\"},{\"path\":\"/email/emailTest\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/setAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/getAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/authorityBtn/canRemoveAuthorityBtn\",\"method\":\"POST\"},{\"path\":\"/chatGpt/getTable\",\"method\":\"POST\"},{\"path\":\"/chatGpt/createSK\",\"method\":\"POST\"},{\"path\":\"/chatGpt/getSK\",\"method\":\"GET\"},{\"path\":\"/chatGpt/deleteSK\",\"method\":\"DELETE\"},{\"path\":\"/company/createCompany\",\"method\":\"POST\"},{\"path\":\"/company/deleteCompany\",\"method\":\"DELETE\"},{\"path\":\"/company/deleteCompanyByIds\",\"method\":\"DELETE\"},{\"path\":\"/company/updateCompany\",\"method\":\"PUT\"},{\"path\":\"/company/findCompany\",\"method\":\"GET\"},{\"path\":\"/company/getCompanyList\",\"method\":\"GET\"},{\"path\":\"/appRole/createAppRole\",\"method\":\"POST\"},{\"path\":\"/appRole/deleteAppRole\",\"method\":\"DELETE\"},{\"path\":\"/appRole/deleteAppRoleByIds\",\"method\":\"DELETE\"},{\"path\":\"/appRole/updateAppRole\",\"method\":\"PUT\"},{\"path\":\"/appRole/findAppRole\",\"method\":\"GET\"},{\"path\":\"/appRole/getAppRoleList\",\"method\":\"GET\"},{\"path\":\"/team/createTeam\",\"method\":\"POST\"},{\"path\":\"/team/deleteTeam\",\"method\":\"DELETE\"},{\"path\":\"/team/deleteTeamByIds\",\"method\":\"DELETE\"},{\"path\":\"/team/updateTeam\",\"method\":\"PUT\"},{\"path\":\"/team/findTeam\",\"method\":\"GET\"},{\"path\":\"/team/getTeamList\",\"method\":\"GET\"},{\"path\":\"/appUser/createAppUser\",\"method\":\"POST\"},{\"path\":\"/appUser/deleteAppUser\",\"method\":\"DELETE\"},{\"path\":\"/appUser/deleteAppUserByIds\",\"method\":\"DELETE\"},{\"path\":\"/appUser/updateAppUser\",\"method\":\"PUT\"},{\"path\":\"/appUser/findAppUser\",\"method\":\"GET\"},{\"path\":\"/appUser/getAppUserList\",\"method\":\"GET\"},{\"path\":\"/userRole/createUserRole\",\"method\":\"POST\"},{\"path\":\"/userRole/deleteUserRole\",\"method\":\"DELETE\"},{\"path\":\"/userRole/deleteUserRoleByIds\",\"method\":\"DELETE\"},{\"path\":\"/userRole/updateUserRole\",\"method\":\"PUT\"},{\"path\":\"/userRole/findUserRole\",\"method\":\"GET\"},{\"path\":\"/userRole/getUserRoleList\",\"method\":\"GET\"}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (22, '2023-05-03 09:10:29.341', '2023-05-03 09:10:29.341', NULL, '127.0.0.1', 'POST', '/system/getSystemConfig', 200, 520400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '', '{\"code\":0,\"data\":{\"config\":{\"jwt\":{\"signing-key\":\"4b611845-6352-4cb2-a65c-db25bdc43383\",\"expires-time\":\"7d\",\"buffer-time\":\"1d\",\"issuer\":\"qmPlus\"},\"zap\":{\"level\":\"info\",\"prefix\":\"[github.com/flipped-aurora/gin-vue-admin/server]\",\"format\":\"console\",\"director\":\"log\",\"encode-level\":\"LowercaseColorLevelEncoder\",\"stacktrace-key\":\"stacktrace\",\"max-age\":0,\"show-line\":true,\"log-in-console\":true},\"redis\":{\"db\":0,\"addr\":\"127.0.0.1:6379\",\"password\":\"\"},\"email\":{\"to\":\"xxx@qq.com\",\"port\":465,\"from\":\"xxx@163.com\",\"host\":\"smtp.163.com\",\"is-ssl\":true,\"secret\":\"xxx\",\"nickname\":\"test\"},\"system\":{\"env\":\"public\",\"addr\":8889,\"db-type\":\"mysql\",\"oss-type\":\"local\",\"use-multipoint\":false,\"use-redis\":false,\"iplimit-count\":15000,\"iplimit-time\":3600,\"router-prefix\":\"\"},\"captcha\":{\"key-long\":6,\"img-width\":240,\"img-height\":80,\"open-captcha\":0,\"open-captcha-timeout\":3600},\"autocode\":{\"transfer-restart\":true,\"root\":\"d:\\\\project\\\\clothing\",\"server\":\"/server\",\"server-api\":\"/api/v1/%s\",\"server-plug\":\"/plugin/%s\",\"server-initialize\":\"/initialize\",\"server-model\":\"/model/%s\",\"server-request\":\"/model/%s/request/\",\"server-router\":\"/router/%s\",\"server-service\":\"/service/%s\",\"web\":\"/web/src\",\"web-api\":\"/api\",\"web-form\":\"/view\",\"web-table\":\"/view\"},\"mysql\":{\"path\":\"39.105.214.148\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"db-name\":\"gva\",\"username\":\"root\",\"password\":\"local123\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"error\",\"log-zap\":false},\"mssql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"pgsql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"oracle\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"db-list\":[{\"disable\":true,\"type\":\"\",\"alias-name\":\"\",\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false}],\"local\":{\"path\":\"uploads/file\",\"store-path\":\"uploads/file\"},\"qiniu\":{\"zone\":\"ZoneHuaDong\",\"bucket\":\"\",\"img-path\":\"\",\"use-https\":false,\"access-key\":\"\",\"secret-key\":\"\",\"use-cdn-domains\":false},\"aliyun-oss\":{\"endpoint\":\"yourEndpoint\",\"access-key-id\":\"yourAccessKeyId\",\"access-key-secret\":\"yourAccessKeySecret\",\"bucket-name\":\"yourBucketName\",\"bucket-url\":\"yourBucketUrl\",\"base-path\":\"yourBasePath\"},\"hua-wei-obs\":{\"path\":\"you-path\",\"bucket\":\"you-bucket\",\"endpoint\":\"you-endpoint\",\"access-key\":\"you-access-key\",\"secret-key\":\"you-secret-key\"},\"tencent-cos\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"aws-s3\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"endpoint\":\"\",\"s3-force-path-style\":false,\"disable-ssl\":false,\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"excel\":{\"dir\":\"./resource/excel/\"},\"timer\":{\"start\":true,\"spec\":\"@daily\",\"with_seconds\":false,\"detail\":[{\"tableName\":\"sys_operation_records\",\"compareField\":\"created_at\",\"interval\":\"2160h\"},{\"tableName\":\"jwt_blacklists\",\"compareField\":\"created_at\",\"interval\":\"168h\"}]},\"cors\":{\"mode\":\"strict-whitelist\",\"whitelist\":[{\"allow-origin\":\"example1.com\",\"allow-methods\":\"POST, GET\",\"allow-headers\":\"Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true},{\"allow-origin\":\"example2.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true}]}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (23, '2023-05-03 15:10:09.293', '2023-05-03 15:10:09.293', NULL, '127.0.0.1', 'POST', '/sysDictionary/createSysDictionary', 200, 281916700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"裁剪单步骤\",\"type\":\"step\",\"status\":true,\"desc\":\"裁剪单步骤\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (24, '2023-05-03 15:14:11.005', '2023-05-03 15:14:11.005', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 178924100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"待处理\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":9}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (25, '2023-05-03 15:14:31.644', '2023-05-03 15:14:31.644', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 178705800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"处理中\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":9}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (26, '2023-05-03 15:14:43.752', '2023-05-03 15:14:43.752', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 175303700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"待审核\",\"value\":3,\"status\":true,\"sort\":3,\"sysDictionaryID\":9}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (27, '2023-05-03 15:14:56.971', '2023-05-03 15:14:56.971', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 177806700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"已完成\",\"value\":4,\"status\":true,\"sort\":4,\"sysDictionaryID\":9}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (28, '2023-05-03 15:36:15.009', '2023-05-03 15:36:15.009', NULL, '127.0.0.1', 'POST', '/sysDictionary/createSysDictionary', 200, 508203500, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"工作类型\",\"type\":\"jobType\",\"status\":true,\"desc\":\"工作类型\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (29, '2023-05-03 15:36:40.327', '2023-05-03 15:36:40.327', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 251540300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"工序\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":10}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (30, '2023-05-03 15:36:49.106', '2023-05-03 15:36:49.106', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 326180300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"成衣\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":10}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (31, '2023-05-04 09:21:35.380', '2023-05-04 09:21:35.380', NULL, '127.0.0.1', 'POST', '/sysDictionary/createSysDictionary', 200, 268307100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"处理状态\",\"type\":\"handleStatus\",\"status\":true,\"desc\":\"处理状态\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (32, '2023-05-04 09:22:00.398', '2023-05-04 09:22:00.398', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 180639100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"待处理\",\"value\":0,\"status\":true,\"sort\":0,\"sysDictionaryID\":11}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (33, '2023-05-04 09:22:23.132', '2023-05-04 09:22:23.132', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 178483900, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"已完成\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":11}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (34, '2023-05-04 09:29:10.824', '2023-05-04 09:29:10.824', NULL, '127.0.0.1', 'POST', '/sysDictionary/createSysDictionary', 200, 263835100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"消息类型\",\"type\":\"noticeType\",\"status\":true,\"desc\":\"消息类型\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (35, '2023-05-04 09:29:32.418', '2023-05-04 09:29:32.418', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 174654400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"进组申请\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (36, '2023-05-04 09:30:02.665', '2023-05-04 09:30:02.665', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 173633200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"工单问题\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (37, '2023-05-04 09:30:17.650', '2023-05-04 09:30:17.650', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 174607200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"工单审核\",\"value\":3,\"status\":true,\"sort\":3,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (38, '2023-05-04 09:31:07.679', '2023-05-04 09:31:07.679', NULL, '127.0.0.1', 'PUT', '/sysDictionary/updateSysDictionary', 200, 360646800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":12,\"CreatedAt\":\"2023-05-04T09:29:10.65+08:00\",\"UpdatedAt\":\"2023-05-04T09:29:10.65+08:00\",\"name\":\"消息类型\",\"type\":\"msgType\",\"status\":true,\"desc\":\"消息类型\",\"sysDictionaryDetails\":[{\"ID\":38,\"CreatedAt\":\"2023-05-04T09:29:32.242+08:00\",\"UpdatedAt\":\"2023-05-04T09:29:32.242+08:00\",\"label\":\"进组申请\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":12},{\"ID\":39,\"CreatedAt\":\"2023-05-04T09:30:02.489+08:00\",\"UpdatedAt\":\"2023-05-04T09:30:02.489+08:00\",\"label\":\"工单问题\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":12},{\"ID\":40,\"CreatedAt\":\"2023-05-04T09:30:17.475+08:00\",\"UpdatedAt\":\"2023-05-04T09:30:17.475+08:00\",\"label\":\"工单审核\",\"value\":3,\"status\":true,\"sort\":3,\"sysDictionaryID\":12}]}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (39, '2023-05-04 09:41:24.797', '2023-05-04 09:41:24.797', NULL, '127.0.0.1', 'POST', '/sysDictionary/createSysDictionary', 200, 260514600, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"媒体类型\",\"type\":\"mediaType\",\"status\":true,\"desc\":\"媒体类型\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (40, '2023-05-04 09:41:36.278', '2023-05-04 09:41:36.278', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 172092300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"图片\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":13}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (41, '2023-05-04 09:41:45.309', '2023-05-04 09:41:45.309', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 172180000, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"视频\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":13}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (42, '2023-05-04 11:10:02.744', '2023-05-04 11:10:02.744', NULL, '127.0.0.1', 'POST', '/appRole/createAppRole', 200, 205675800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"老板\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (43, '2023-05-04 11:10:08.345', '2023-05-04 11:10:08.345', NULL, '127.0.0.1', 'POST', '/appRole/createAppRole', 200, 170564900, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"裁缝\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (44, '2023-05-04 11:10:13.006', '2023-05-04 11:10:13.006', NULL, '127.0.0.1', 'POST', '/appRole/createAppRole', 200, 171174700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"组长\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (45, '2023-05-04 11:10:17.383', '2023-05-04 11:10:17.383', NULL, '127.0.0.1', 'POST', '/appRole/createAppRole', 200, 177455400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"name\":\"车工\"}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (46, '2023-05-04 11:33:25.469', '2023-05-04 11:33:25.469', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 177929800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"裁缝申请\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (47, '2023-05-04 11:33:41.577', '2023-05-04 11:33:41.577', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 178883200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"组长申请\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (48, '2023-05-04 11:33:48.758', '2023-05-04 11:33:48.758', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 178656800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":38,\"CreatedAt\":\"2023-05-04T09:29:32.242+08:00\",\"UpdatedAt\":\"2023-05-04T09:29:32.242+08:00\",\"label\":\"进组申请\",\"value\":3,\"status\":true,\"sort\":3,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (49, '2023-05-04 11:33:55.867', '2023-05-04 11:33:55.867', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 180006500, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":43,\"CreatedAt\":\"2023-05-04T11:33:25.288+08:00\",\"UpdatedAt\":\"2023-05-04T11:33:25.288+08:00\",\"label\":\"裁缝申请\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (50, '2023-05-04 11:34:05.032', '2023-05-04 11:34:05.032', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 190914200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":39,\"CreatedAt\":\"2023-05-04T09:30:02.489+08:00\",\"UpdatedAt\":\"2023-05-04T09:30:02.489+08:00\",\"label\":\"工单问题\",\"value\":4,\"status\":true,\"sort\":4,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (51, '2023-05-04 11:34:12.098', '2023-05-04 11:34:12.098', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 183971100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":40,\"CreatedAt\":\"2023-05-04T09:30:17.475+08:00\",\"UpdatedAt\":\"2023-05-04T09:30:17.475+08:00\",\"label\":\"工单审核\",\"value\":5,\"status\":true,\"sort\":5,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (52, '2023-05-04 13:39:22.641', '2023-05-04 13:39:22.641', NULL, '127.0.0.1', 'POST', '/system/getSystemConfig', 200, 521700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '', '{\"code\":0,\"data\":{\"config\":{\"jwt\":{\"signing-key\":\"b84badd0-4e04-4ee9-a914-a38116489b6a\",\"expires-time\":\"7d\",\"buffer-time\":\"1d\",\"issuer\":\"qmPlus\"},\"zap\":{\"level\":\"info\",\"prefix\":\"[github.com/flipped-aurora/gin-vue-admin/server]\",\"format\":\"console\",\"director\":\"log\",\"encode-level\":\"LowercaseColorLevelEncoder\",\"stacktrace-key\":\"stacktrace\",\"max-age\":0,\"show-line\":true,\"log-in-console\":true},\"redis\":{\"db\":0,\"addr\":\"127.0.0.1:6379\",\"password\":\"\",\"prefix\":\"\"},\"email\":{\"to\":\"xxx@qq.com\",\"port\":465,\"from\":\"xxx@163.com\",\"host\":\"smtp.163.com\",\"is-ssl\":true,\"secret\":\"xxx\",\"nickname\":\"test\"},\"system\":{\"env\":\"public\",\"addr\":8889,\"db-type\":\"mysql\",\"oss-type\":\"local\",\"use-multipoint\":false,\"use-redis\":false,\"iplimit-count\":15000,\"iplimit-time\":3600,\"router-prefix\":\"\"},\"captcha\":{\"key-long\":6,\"img-width\":240,\"img-height\":80,\"open-captcha\":0,\"open-captcha-timeout\":3600},\"autocode\":{\"transfer-restart\":true,\"root\":\"D:\\\\project\\\\clothing\",\"server\":\"/server\",\"server-api\":\"/api/v1/%s\",\"server-plug\":\"/plugin/%s\",\"server-initialize\":\"/initialize\",\"server-model\":\"/model/%s\",\"server-request\":\"/model/%s/request/\",\"server-router\":\"/router/%s\",\"server-service\":\"/service/%s\",\"web\":\"/web/src\",\"web-api\":\"/api\",\"web-form\":\"/view\",\"web-table\":\"/view\"},\"mysql\":{\"path\":\"127.0.0.1\",\"port\":\"3306\",\"config\":\"charset=utf8mb4\\u0026parseTime=True\\u0026loc=Local\",\"db-name\":\"clothing\",\"username\":\"root\",\"password\":\"local123\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"error\",\"log-zap\":false},\"mssql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"pgsql\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"oracle\":{\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false},\"db-list\":[{\"disable\":true,\"type\":\"\",\"alias-name\":\"\",\"path\":\"\",\"port\":\"\",\"config\":\"\",\"db-name\":\"\",\"username\":\"\",\"password\":\"\",\"prefix\":\"\",\"singular\":false,\"engine\":\"\",\"max-idle-conns\":10,\"max-open-conns\":100,\"log-mode\":\"\",\"log-zap\":false}],\"local\":{\"path\":\"uploads/file\",\"store-path\":\"uploads/file\"},\"qiniu\":{\"zone\":\"ZoneHuaDong\",\"bucket\":\"\",\"img-path\":\"\",\"use-https\":false,\"access-key\":\"\",\"secret-key\":\"\",\"use-cdn-domains\":false},\"aliyun-oss\":{\"endpoint\":\"yourEndpoint\",\"access-key-id\":\"yourAccessKeyId\",\"access-key-secret\":\"yourAccessKeySecret\",\"bucket-name\":\"yourBucketName\",\"bucket-url\":\"yourBucketUrl\",\"base-path\":\"yourBasePath\"},\"hua-wei-obs\":{\"path\":\"you-path\",\"bucket\":\"you-bucket\",\"endpoint\":\"you-endpoint\",\"access-key\":\"you-access-key\",\"secret-key\":\"you-secret-key\"},\"tencent-cos\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"aws-s3\":{\"bucket\":\"xxxxx-10005608\",\"region\":\"ap-shanghai\",\"endpoint\":\"\",\"s3-force-path-style\":false,\"disable-ssl\":false,\"secret-id\":\"your-secret-id\",\"secret-key\":\"your-secret-key\",\"base-url\":\"https://gin.vue.admin\",\"path-prefix\":\"github.com/flipped-aurora/gin-vue-admin/server\"},\"excel\":{\"dir\":\"./resource/excel/\"},\"timer\":{\"start\":true,\"spec\":\"@daily\",\"with_seconds\":false,\"detail\":[{\"tableName\":\"sys_operation_records\",\"compareField\":\"created_at\",\"interval\":\"2160h\"},{\"tableName\":\"jwt_blacklists\",\"compareField\":\"created_at\",\"interval\":\"168h\"}]},\"cors\":{\"mode\":\"strict-whitelist\",\"whitelist\":[{\"allow-origin\":\"example1.com\",\"allow-methods\":\"POST, GET\",\"allow-headers\":\"Content-Type,AccessToken,X-CSRF-Token, Authorization, Token,X-Token,X-User-Id\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true},{\"allow-origin\":\"example2.com\",\"allow-methods\":\"GET, POST\",\"allow-headers\":\"content-type\",\"expose-headers\":\"Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Content-Type\",\"allow-credentials\":true}]}}},\"msg\":\"获取成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (53, '2023-05-05 10:45:50.362', '2023-05-05 10:45:50.362', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 15003800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":43,\"CreatedAt\":\"2023-05-04T11:33:25.288+08:00\",\"UpdatedAt\":\"2023-05-04T11:33:55.687+08:00\",\"label\":\"裁缝/组长申请\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (54, '2023-05-05 10:45:59.109', '2023-05-05 10:45:59.109', NULL, '127.0.0.1', 'DELETE', '/sysDictionaryDetail/deleteSysDictionaryDetail', 200, 5649200, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":44}', '{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (55, '2023-05-05 10:46:03.848', '2023-05-05 10:46:03.848', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 7841300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":38,\"CreatedAt\":\"2023-05-04T09:29:32.242+08:00\",\"UpdatedAt\":\"2023-05-04T11:33:48.576+08:00\",\"label\":\"进组申请\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (56, '2023-05-05 10:46:10.332', '2023-05-05 10:46:10.332', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 6339400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":39,\"CreatedAt\":\"2023-05-04T09:30:02.489+08:00\",\"UpdatedAt\":\"2023-05-04T11:34:04.838+08:00\",\"label\":\"工单问题\",\"value\":3,\"status\":true,\"sort\":3,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (57, '2023-05-05 10:46:14.254', '2023-05-05 10:46:14.254', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 4288300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":40,\"CreatedAt\":\"2023-05-04T09:30:17.475+08:00\",\"UpdatedAt\":\"2023-05-04T11:34:11.914+08:00\",\"label\":\"工单审核\",\"value\":4,\"status\":true,\"sort\":4,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (58, '2023-05-08 09:06:08.749', '2023-05-08 09:06:08.749', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 165308400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"工单申请\",\"value\":3,\"status\":true,\"sort\":3,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (59, '2023-05-08 09:06:16.527', '2023-05-08 09:06:16.527', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 37250900, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":39,\"CreatedAt\":\"2023-05-04T09:30:02.489+08:00\",\"UpdatedAt\":\"2023-05-05T10:46:10.327+08:00\",\"label\":\"工单问题\",\"value\":4,\"status\":true,\"sort\":4,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (60, '2023-05-08 09:06:21.694', '2023-05-08 09:06:21.694', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 58388600, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":40,\"CreatedAt\":\"2023-05-04T09:30:17.475+08:00\",\"UpdatedAt\":\"2023-05-05T10:46:14.25+08:00\",\"label\":\"工单审核\",\"value\":5,\"status\":true,\"sort\":5,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (61, '2023-05-08 09:46:06.995', '2023-05-08 09:46:06.995', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 34532700, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"组长申请\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (62, '2023-05-08 09:46:12.798', '2023-05-08 09:46:12.798', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 41695000, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":43,\"CreatedAt\":\"2023-05-04T11:33:25.288+08:00\",\"UpdatedAt\":\"2023-05-05T10:45:50.348+08:00\",\"label\":\"裁缝申请\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (63, '2023-05-08 09:46:17.226', '2023-05-08 09:46:17.226', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 37396600, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":46,\"CreatedAt\":\"2023-05-08T09:46:06.961+08:00\",\"UpdatedAt\":\"2023-05-08T09:46:06.961+08:00\",\"label\":\"组长申请\",\"value\":3,\"status\":true,\"sort\":3,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (64, '2023-05-08 09:46:28.595', '2023-05-08 09:46:28.595', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 34677100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":46,\"CreatedAt\":\"2023-05-08T09:46:06.961+08:00\",\"UpdatedAt\":\"2023-05-08T09:46:17.187+08:00\",\"label\":\"组长申请\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (65, '2023-05-08 09:46:38.331', '2023-05-08 09:46:38.331', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 35042400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":38,\"CreatedAt\":\"2023-05-04T09:29:32.242+08:00\",\"UpdatedAt\":\"2023-05-05T10:46:03.842+08:00\",\"label\":\"进组申请\",\"value\":3,\"status\":true,\"sort\":3,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (66, '2023-05-08 09:46:43.500', '2023-05-08 09:46:43.500', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 39902100, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":45,\"CreatedAt\":\"2023-05-08T09:06:08.67+08:00\",\"UpdatedAt\":\"2023-05-08T09:06:08.67+08:00\",\"label\":\"工单申请\",\"value\":4,\"status\":true,\"sort\":4,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (67, '2023-05-08 09:46:50.396', '2023-05-08 09:46:50.396', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 34994600, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":39,\"CreatedAt\":\"2023-05-04T09:30:02.489+08:00\",\"UpdatedAt\":\"2023-05-08T09:06:16.487+08:00\",\"label\":\"工单问题\",\"value\":5,\"status\":true,\"sort\":5,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (68, '2023-05-08 09:50:23.589', '2023-05-08 09:50:23.589', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 34873900, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":40,\"CreatedAt\":\"2023-05-04T09:30:17.475+08:00\",\"UpdatedAt\":\"2023-05-08T09:06:21.632+08:00\",\"label\":\"工单审核\",\"value\":6,\"status\":true,\"sort\":6,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (69, '2023-05-08 13:37:37.890', '2023-05-08 13:37:37.890', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 37238800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":43,\"CreatedAt\":\"2023-05-04T11:33:25.288+08:00\",\"UpdatedAt\":\"2023-05-08T09:46:12.757+08:00\",\"label\":\"裁缝、组长申请\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (70, '2023-05-08 13:37:52.465', '2023-05-08 13:37:52.465', NULL, '127.0.0.1', 'DELETE', '/sysDictionaryDetail/deleteSysDictionaryDetail', 200, 32952400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":46}', '{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (71, '2023-05-08 13:37:58.169', '2023-05-08 13:37:58.169', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 36058300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":38,\"CreatedAt\":\"2023-05-04T09:29:32.242+08:00\",\"UpdatedAt\":\"2023-05-08T09:46:38.296+08:00\",\"label\":\"进组申请\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (72, '2023-05-08 13:38:02.930', '2023-05-08 13:38:02.930', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 38928800, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":45,\"CreatedAt\":\"2023-05-08T09:06:08.67+08:00\",\"UpdatedAt\":\"2023-05-08T09:46:43.457+08:00\",\"label\":\"工单申请\",\"value\":3,\"status\":true,\"sort\":3,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (73, '2023-05-08 13:38:06.385', '2023-05-08 13:38:06.385', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 34710300, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":39,\"CreatedAt\":\"2023-05-04T09:30:02.489+08:00\",\"UpdatedAt\":\"2023-05-08T09:46:50.36+08:00\",\"label\":\"工单问题\",\"value\":4,\"status\":true,\"sort\":4,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (74, '2023-05-08 13:38:12.023', '2023-05-08 13:38:12.023', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 33318400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":40,\"CreatedAt\":\"2023-05-04T09:30:17.475+08:00\",\"UpdatedAt\":\"2023-05-08T09:50:23.555+08:00\",\"label\":\"工单审核\",\"value\":5,\"status\":true,\"sort\":5,\"sysDictionaryID\":12}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (75, '2023-05-08 16:49:14.881', '2023-05-08 16:49:14.881', NULL, '127.0.0.1', 'PUT', '/sysDictionaryDetail/updateSysDictionaryDetail', 200, 47861400, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"ID\":37,\"CreatedAt\":\"2023-05-04T09:22:22.952+08:00\",\"UpdatedAt\":\"2023-05-04T09:22:22.952+08:00\",\"label\":\"已通过\",\"value\":1,\"status\":true,\"sort\":1,\"sysDictionaryID\":11}', '{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}', 1);
INSERT INTO `sys_operation_records` VALUES (76, '2023-05-08 16:49:22.452', '2023-05-08 16:49:22.452', NULL, '127.0.0.1', 'POST', '/sysDictionaryDetail/createSysDictionaryDetail', 200, 36420500, 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 Edg/112.0.1722.64', '', '{\"label\":\"已拒绝\",\"value\":2,\"status\":true,\"sort\":2,\"sysDictionaryID\":11}', '{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}', 1);

-- ----------------------------
-- Table structure for sys_user_authority
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_authority`;
CREATE TABLE `sys_user_authority`  (
  `sys_user_id` bigint UNSIGNED NOT NULL,
  `sys_authority_authority_id` bigint UNSIGNED NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`, `sys_authority_authority_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_user_authority
-- ----------------------------
INSERT INTO `sys_user_authority` VALUES (1, 888);
INSERT INTO `sys_user_authority` VALUES (1, 8881);
INSERT INTO `sys_user_authority` VALUES (1, 9528);
INSERT INTO `sys_user_authority` VALUES (2, 888);

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `uuid` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` bigint UNSIGNED NULL DEFAULT 888 COMMENT '用户角色ID',
  `phone` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint NULL DEFAULT 1 COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at` ASC) USING BTREE,
  INDEX `idx_sys_users_uuid`(`uuid` ASC) USING BTREE,
  INDEX `idx_sys_users_username`(`username` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (1, '2023-05-02 17:43:03.221', '2023-05-02 17:43:03.396', NULL, '88098e5d-b8c6-46fa-9713-499d924f4b87', 'admin', '$2a$10$i81guTFWJ6.KuLl88a.u/elAiOCsxa0JaICZHQm3W6cvEV64hULLq', 'Mr.奇淼', 'dark', 'https://qmplusimg.henrongyi.top/gva_header.jpg', '#fff', '#1890ff', 888, '17611111111', '333333333@qq.com', 1);
INSERT INTO `sys_users` VALUES (2, '2023-05-02 17:43:03.221', '2023-05-02 17:43:03.966', NULL, 'e959bf30-01ad-4a05-a7a5-de0c0ea92c7f', 'a303176530', '$2a$10$QH73aZ/Qy.2gi8zF/HoOhucTN/m.iq3QxgabrLPd0zYjRN447pCo6', '用户1', 'dark', 'https:///qmplusimg.henrongyi.top/1572075907logo.png', '#fff', '#1890ff', 9528, '17611111111', '333333333@qq.com', 1);

-- ----------------------------
-- Table structure for team
-- ----------------------------
DROP TABLE IF EXISTS `team`;
CREATE TABLE `team`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `company_id` bigint UNSIGNED NULL DEFAULT NULL,
  `name` varchar(191) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_team_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of team
-- ----------------------------
INSERT INTO `team` VALUES (1, '2023-05-09 09:15:10.880', '2023-05-09 09:15:10.880', NULL, 1, 'test', 0, 0, 0, 4);

-- ----------------------------
-- Table structure for team_apply
-- ----------------------------
DROP TABLE IF EXISTS `team_apply`;
CREATE TABLE `team_apply`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `team_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `status` bigint NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_team_apply_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of team_apply
-- ----------------------------
INSERT INTO `team_apply` VALUES (1, '2023-05-09 16:16:15.476', '2023-05-09 16:22:40.157', NULL, 1, 3, 1, 0, 0, 0);
INSERT INTO `team_apply` VALUES (2, '2023-05-09 16:16:17.823', '2023-05-09 16:21:37.129', NULL, 1, 3, 1, 0, 0, 0);

-- ----------------------------
-- Table structure for team_user
-- ----------------------------
DROP TABLE IF EXISTS `team_user`;
CREATE TABLE `team_user`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `team_id` bigint UNSIGNED NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_team_user_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of team_user
-- ----------------------------
INSERT INTO `team_user` VALUES (1, '2023-05-09 16:21:37.073', '2023-05-09 16:21:37.073', NULL, 1, 4, 0, 0, 0);
INSERT INTO `team_user` VALUES (2, '2023-05-09 16:21:37.073', '2023-05-09 16:21:37.073', NULL, 1, 3, 0, 0, 0);

-- ----------------------------
-- Table structure for user_role
-- ----------------------------
DROP TABLE IF EXISTS `user_role`;
CREATE TABLE `user_role`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `role_id` bigint UNSIGNED NULL DEFAULT NULL,
  `company_id` bigint UNSIGNED NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_role_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 7 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_role
-- ----------------------------
INSERT INTO `user_role` VALUES (1, '2023-05-08 10:37:12.638', '2023-05-08 10:37:12.638', NULL, 1, 1, 1, 0, 0, 0);
INSERT INTO `user_role` VALUES (4, '2023-05-09 09:14:08.875', '2023-05-09 09:14:08.875', NULL, 2, 2, 1, 0, 0, 0);
INSERT INTO `user_role` VALUES (5, '2023-05-09 09:15:10.990', '2023-05-09 09:15:10.990', NULL, 4, 3, 1, 0, 0, 0);
INSERT INTO `user_role` VALUES (6, '2023-05-09 16:16:15.414', '2023-05-09 16:16:15.414', NULL, 3, 4, 1, 0, 0, 0);

-- ----------------------------
-- Table structure for user_wallet
-- ----------------------------
DROP TABLE IF EXISTS `user_wallet`;
CREATE TABLE `user_wallet`  (
  `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) NULL DEFAULT NULL,
  `updated_at` datetime(3) NULL DEFAULT NULL,
  `deleted_at` datetime(3) NULL DEFAULT NULL,
  `user_id` bigint UNSIGNED NULL DEFAULT NULL,
  `company_id` bigint UNSIGNED NULL DEFAULT NULL,
  `wages` decimal(10, 2) NULL DEFAULT NULL,
  `created_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '创建者',
  `updated_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '更新者',
  `deleted_by` bigint UNSIGNED NULL DEFAULT NULL COMMENT '删除者',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_user_wallet_deleted_at`(`deleted_at` ASC) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of user_wallet
-- ----------------------------

SET FOREIGN_KEY_CHECKS = 1;
