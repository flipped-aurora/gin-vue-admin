/*
 Navicat Premium Data Transfer

 Source Server         : qmplus
 Source Server Type    : MySQL
 Source Server Version : 50644
 Source Host           : localhost:3306
 Source Schema         : qmplus

 Target Server Type    : MySQL
 Target Server Version : 50644
 File Encoding         : 65001

 Date: 02/04/2020 14:33:38
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for api_authorities
-- ----------------------------
DROP TABLE IF EXISTS `api_authorities`;
CREATE TABLE `api_authorities`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `authority_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `api_id` int(10) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_api_authorities_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 292 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of api_authorities
-- ----------------------------
INSERT INTO `api_authorities` VALUES (72, '2019-09-30 15:34:19', '2019-09-30 15:34:19', NULL, 9528, 1);
INSERT INTO `api_authorities` VALUES (73, '2019-09-30 15:34:19', '2019-09-30 15:34:19', NULL, 9528, 4);
INSERT INTO `api_authorities` VALUES (74, '2019-09-30 15:34:19', '2019-09-30 15:34:19', NULL, 9528, 5);
INSERT INTO `api_authorities` VALUES (75, '2019-09-30 15:34:19', '2019-09-30 15:34:19', NULL, 9528, 6);
INSERT INTO `api_authorities` VALUES (76, '2019-09-30 15:34:19', '2019-09-30 15:34:19', NULL, 9528, 14);
INSERT INTO `api_authorities` VALUES (77, '2019-09-30 15:34:19', '2019-09-30 15:34:19', NULL, 9528, 15);
INSERT INTO `api_authorities` VALUES (78, '2019-09-30 15:34:19', '2019-09-30 15:34:19', NULL, 9528, 17);
INSERT INTO `api_authorities` VALUES (79, '2019-09-30 15:34:19', '2019-09-30 15:34:19', NULL, 9528, 23);
INSERT INTO `api_authorities` VALUES (80, '2019-09-30 15:34:19', '2019-09-30 15:34:19', NULL, 9528, 24);
INSERT INTO `api_authorities` VALUES (210, '2019-10-09 23:04:56', '2019-10-09 23:04:56', NULL, 999, 1);
INSERT INTO `api_authorities` VALUES (211, '2019-10-09 23:04:56', '2019-10-09 23:04:56', NULL, 999, 2);
INSERT INTO `api_authorities` VALUES (265, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 1);
INSERT INTO `api_authorities` VALUES (266, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 2);
INSERT INTO `api_authorities` VALUES (267, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 3);
INSERT INTO `api_authorities` VALUES (268, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 4);
INSERT INTO `api_authorities` VALUES (269, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 5);
INSERT INTO `api_authorities` VALUES (270, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 6);
INSERT INTO `api_authorities` VALUES (271, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 7);
INSERT INTO `api_authorities` VALUES (272, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 8);
INSERT INTO `api_authorities` VALUES (273, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 9);
INSERT INTO `api_authorities` VALUES (274, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 10);
INSERT INTO `api_authorities` VALUES (275, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 11);
INSERT INTO `api_authorities` VALUES (276, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 12);
INSERT INTO `api_authorities` VALUES (277, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 13);
INSERT INTO `api_authorities` VALUES (278, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 14);
INSERT INTO `api_authorities` VALUES (279, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 15);
INSERT INTO `api_authorities` VALUES (280, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 16);
INSERT INTO `api_authorities` VALUES (281, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 17);
INSERT INTO `api_authorities` VALUES (282, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 18);
INSERT INTO `api_authorities` VALUES (283, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 19);
INSERT INTO `api_authorities` VALUES (284, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 20);
INSERT INTO `api_authorities` VALUES (285, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 21);
INSERT INTO `api_authorities` VALUES (286, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 22);
INSERT INTO `api_authorities` VALUES (287, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 23);
INSERT INTO `api_authorities` VALUES (288, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 24);
INSERT INTO `api_authorities` VALUES (289, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 28);
INSERT INTO `api_authorities` VALUES (290, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 29);
INSERT INTO `api_authorities` VALUES (291, '2019-10-13 17:06:10', '2019-10-13 17:06:10', NULL, 888, 30);

-- ----------------------------
-- Table structure for apis
-- ----------------------------
DROP TABLE IF EXISTS `apis`;
CREATE TABLE `apis`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `authority_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `group` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of apis
-- ----------------------------
INSERT INTO `apis` VALUES (1, '2019-09-28 11:23:49', '2019-09-28 17:06:16', NULL, NULL, '/base/login', '用户登录', 'base');
INSERT INTO `apis` VALUES (2, '2019-09-28 11:32:46', '2019-09-28 17:06:11', NULL, NULL, '/base/regist', '用户注册', 'base');
INSERT INTO `apis` VALUES (3, '2019-09-28 11:33:41', '2019-09-28 17:06:04', NULL, NULL, '/api/createApi', '创建api', 'api');
INSERT INTO `apis` VALUES (4, '2019-09-28 14:09:04', '2019-09-28 17:05:59', NULL, NULL, '/api/getApiList', '获取api列表', 'api');
INSERT INTO `apis` VALUES (5, '2019-09-28 14:15:50', '2019-09-28 17:05:53', NULL, NULL, '/api/getApiById', '获取api详细信息', 'api');
INSERT INTO `apis` VALUES (6, '2019-09-28 14:19:08', '2019-09-28 17:05:48', NULL, NULL, '/api/setAuthAndApi', '设置api和角色关系', 'api');
INSERT INTO `apis` VALUES (7, '2019-09-28 14:19:26', '2019-09-28 17:05:44', NULL, NULL, '/api/deleteApi', '删除Api', 'api');
INSERT INTO `apis` VALUES (8, '2019-09-28 14:19:48', '2019-09-28 17:05:39', NULL, NULL, '/api/updataApi', '更新Api', 'api');
INSERT INTO `apis` VALUES (9, '2019-09-30 15:04:55', '2019-09-30 15:04:55', NULL, NULL, '/api/getAuthAndApi', '获取api和角色关系', 'api');
INSERT INTO `apis` VALUES (10, '2019-09-30 15:05:38', '2019-09-30 15:05:38', NULL, NULL, '/api/getAllApis', '获取所有api', 'api');
INSERT INTO `apis` VALUES (11, '2019-09-30 15:23:09', '2019-09-30 15:23:09', NULL, NULL, '/authority/createAuthority', '创建角色', 'authority');
INSERT INTO `apis` VALUES (12, '2019-09-30 15:23:33', '2019-09-30 15:23:33', NULL, NULL, '/authority/deleteAuthority', '删除角色', 'authority');
INSERT INTO `apis` VALUES (13, '2019-09-30 15:23:57', '2019-09-30 15:23:57', NULL, NULL, '/authority/getAuthorityList', '获取角色列表', 'authority');
INSERT INTO `apis` VALUES (14, '2019-09-30 15:24:20', '2019-09-30 15:24:20', NULL, NULL, '/menu/getMenu', '获取菜单树', 'menu');
INSERT INTO `apis` VALUES (15, '2019-09-30 15:24:50', '2019-09-30 15:24:50', NULL, NULL, '/menu/getMenuList', '分页获取基础menu列表', 'menu');
INSERT INTO `apis` VALUES (16, '2019-09-30 15:25:07', '2019-09-30 15:25:07', NULL, NULL, '/menu/addBaseMenu', '新增菜单', 'menu');
INSERT INTO `apis` VALUES (17, '2019-09-30 15:25:25', '2019-09-30 15:25:25', NULL, NULL, '/menu/getBaseMenuTree', '获取用户动态路由', 'menu');
INSERT INTO `apis` VALUES (18, '2019-09-30 15:25:53', '2019-09-30 15:25:53', NULL, NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', 'menu');
INSERT INTO `apis` VALUES (19, '2019-09-30 15:26:20', '2019-09-30 15:26:20', NULL, NULL, '/menu/getMenuAuthority', '获取指定角色menu', 'menu');
INSERT INTO `apis` VALUES (20, '2019-09-30 15:26:43', '2019-09-30 15:26:43', NULL, NULL, '/menu/deleteBaseMenu', '删除菜单', 'menu');
INSERT INTO `apis` VALUES (21, '2019-09-30 15:28:05', '2019-09-30 15:28:05', NULL, NULL, '/menu/updataBaseMenu', '更新菜单', 'menu');
INSERT INTO `apis` VALUES (22, '2019-09-30 15:28:21', '2019-09-30 15:28:21', NULL, NULL, '/menu/getBaseMenuById', '根据id获取菜单', 'menu');
INSERT INTO `apis` VALUES (23, '2019-09-30 15:29:19', '2019-09-30 15:29:19', NULL, NULL, '/user/changePassword', '修改密码', 'user');
INSERT INTO `apis` VALUES (24, '2019-09-30 15:29:33', '2019-09-30 15:29:33', NULL, NULL, '/user/uploadHeaderImg', '上传头像', 'user');
INSERT INTO `apis` VALUES (25, '2019-09-30 15:30:00', '2019-09-30 15:30:00', '2019-10-09 15:26:37', NULL, '/user/getInfoList', '分页获取用户列表', 'user');
INSERT INTO `apis` VALUES (28, '2019-10-09 15:15:17', '2019-10-09 15:17:07', NULL, NULL, '/user/getUserList', '获取用户列表', 'user');
INSERT INTO `apis` VALUES (29, '2019-10-09 23:01:40', '2019-10-09 23:01:40', NULL, NULL, '/user/setUserAuthority', '修改用户角色', 'user');
INSERT INTO `apis` VALUES (30, '2019-10-13 17:01:24', '2019-10-13 17:01:24', NULL, NULL, '/user/deleteUser', '删除用户', 'user');

-- ----------------------------
-- Table structure for authorities
-- ----------------------------
DROP TABLE IF EXISTS `authorities`;
CREATE TABLE `authorities`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `authority_id` int(10) UNSIGNED NOT NULL,
  `authority_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `authority_id`(`authority_id`) USING BTREE,
  INDEX `idx_authorities_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of authorities
-- ----------------------------
INSERT INTO `authorities` VALUES (1, '2019-09-08 16:18:31', '2019-09-08 16:18:31', '2019-09-18 22:10:26', 9527, '超超超级管理员');
INSERT INTO `authorities` VALUES (2, '2019-09-08 16:18:45', '2019-09-08 16:18:45', NULL, 888, '普通用户');
INSERT INTO `authorities` VALUES (3, '2019-09-18 22:20:28', '2019-09-18 22:20:28', '2019-09-18 22:25:22', 0, '测试角色');
INSERT INTO `authorities` VALUES (6, '2019-09-18 22:23:33', '2019-09-18 22:23:33', NULL, 9528, '测试角色');
INSERT INTO `authorities` VALUES (7, '2019-09-18 22:24:39', '2019-09-18 22:24:39', '2019-09-18 22:25:19', 9529, '测试角色');
INSERT INTO `authorities` VALUES (8, '2019-09-18 22:25:13', '2019-09-18 22:25:13', '2019-09-18 22:25:17', 9522, '测试角色');
INSERT INTO `authorities` VALUES (12, '2019-10-09 23:04:18', '2019-10-09 23:04:18', NULL, 999, '封禁');

-- ----------------------------
-- Table structure for base_menus
-- ----------------------------
DROP TABLE IF EXISTS `base_menus`;
CREATE TABLE `base_menus`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `menu_level` int(10) UNSIGNED NULL DEFAULT NULL,
  `parent_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `hidden` tinyint(1) NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_base_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 19 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of base_menus
-- ----------------------------
INSERT INTO `base_menus` VALUES (1, '2019-09-19 22:05:18', '2019-09-19 22:05:18', NULL, 0, 0, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', '仪表盘');
INSERT INTO `base_menus` VALUES (2, '2019-09-19 22:06:17', '2019-09-30 15:44:50', NULL, 0, 0, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', '测试菜单');
INSERT INTO `base_menus` VALUES (3, '2019-09-19 22:06:38', '2019-09-19 22:06:38', NULL, 0, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', '超级管理员');
INSERT INTO `base_menus` VALUES (4, '2019-09-19 22:11:53', '2019-09-19 22:11:53', NULL, 0, 3, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', '角色管理');
INSERT INTO `base_menus` VALUES (5, '2019-09-19 22:13:18', '2019-09-19 22:13:18', NULL, 0, 3, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', '菜单管理');
INSERT INTO `base_menus` VALUES (6, '2019-09-19 22:13:36', '2019-09-19 22:13:36', NULL, 0, 3, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 'api管理');
INSERT INTO `base_menus` VALUES (7, '2019-09-22 17:17:15', '2019-09-26 13:09:57', '2019-09-26 13:10:14', 0, 0, 'tttt22', '修改测试', 0, '', '修改测试', 'share', 'api管理');
INSERT INTO `base_menus` VALUES (8, '2019-09-26 17:17:48', '2019-09-26 17:18:46', '2019-09-26 17:19:23', 0, 0, 'ASDASDASD', 'tttTTT', 0, 'TTTT', 'TTT', 'TTT', 'TTT');
INSERT INTO `base_menus` VALUES (9, '2019-09-30 18:23:01', '2019-09-30 18:23:01', '2019-09-30 18:25:07', 0, 0, 't\'t', 'tttt', 0, 'tt', 'tt', 'tt', 'tt');
INSERT INTO `base_menus` VALUES (10, '2019-09-30 18:23:29', '2019-09-30 18:23:29', '2019-09-30 18:25:05', 0, 0, 'ttt', 'ttt', 0, 'ttt', 'ttt', 'ttt', 'ttt');
INSERT INTO `base_menus` VALUES (11, '2019-09-30 18:23:59', '2019-09-30 18:23:59', '2019-09-30 18:25:04', 0, 0, '11111111', '11111111111111', 0, '111111111', '1111111111', '11111111', '1111111111');
INSERT INTO `base_menus` VALUES (12, '2019-09-30 18:24:03', '2019-09-30 18:24:03', '2019-09-30 18:25:02', 0, 0, '22222', '11111111111111222222222222222222', 0, '11111111222222221', '1111111111222222222222222', '1111112222222211', '1111111111222222222222222');
INSERT INTO `base_menus` VALUES (13, '2019-09-30 18:24:11', '2019-09-30 18:24:11', '2019-09-30 18:25:13', 0, 0, '11114', '124124', 0, '1241', '124124', '214124', '124124');
INSERT INTO `base_menus` VALUES (14, '2019-09-30 18:24:19', '2019-09-30 18:24:19', '2019-09-30 18:25:12', 0, 0, '22222', '222', 0, '2222', '22', '222', '22');
INSERT INTO `base_menus` VALUES (15, '2019-09-30 18:24:22', '2019-09-30 18:24:22', '2019-09-30 18:25:10', 0, 0, '2222211', '22211', 0, '22221', '2211', '2221', '2211');
INSERT INTO `base_menus` VALUES (16, '2019-09-30 18:24:29', '2019-09-30 18:24:29', '2019-09-30 18:25:09', 0, 0, '3223123', '123123', 0, '123123', '123123', '21312323', '123123');
INSERT INTO `base_menus` VALUES (17, '2019-10-09 15:12:29', '2019-10-09 15:14:44', NULL, 0, 3, 'user', 'user', 0, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', '用户管理');
INSERT INTO `base_menus` VALUES (18, '2019-10-13 17:04:24', '2019-10-13 17:04:24', NULL, 0, 0, 'showTest', 'showTest', 0, 'view/showTest/showTest.vue', '展示用vue文件', 'setting', '展示用vue文件');

-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
DROP TABLE IF EXISTS `casbin_rule`;
CREATE TABLE `casbin_rule`  (
  `p_type` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v0` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v1` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v2` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v3` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v4` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `v5` varchar(100) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/base/regist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/updataApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/updataBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/getInfoList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/casbinPUpdata', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/createExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/updataExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/deleteExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/getExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '8881', '/customer/getExaCustomerList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/base/regist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/updataApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/updataBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/getInfoList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/casbinPUpdata', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/createExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/updataExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/deleteExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/getExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/customer/getExaCustomerList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '9528', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/login', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/base/regist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/createApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getApiById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/updataApi', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/updataBaseMenu', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/changePassword', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getInfoList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/getUserList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/casbinPUpdata', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/casbin/casbinTest/:pathParam', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/createExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/updataExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/deleteExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/getExaCustomer', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/customer/getExaCustomerList', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/createTemp', 'POST', '', '', '');

-- ----------------------------
-- Table structure for ch_cities
-- ----------------------------
DROP TABLE IF EXISTS `ch_cities`;
CREATE TABLE `ch_cities`  (
  `id` int(5) NULL DEFAULT NULL,
  `province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `city` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `code` int(12) NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of ch_cities
-- ----------------------------
INSERT INTO `ch_cities` VALUES (1, '北京市', '北京市', 110100);
INSERT INTO `ch_cities` VALUES (2, '天津市', '天津市', 120100);
INSERT INTO `ch_cities` VALUES (3, '河北省', '石家庄市', 130100);
INSERT INTO `ch_cities` VALUES (4, '河北省', '唐山市', 130200);
INSERT INTO `ch_cities` VALUES (5, '河北省', '秦皇岛市', 130300);
INSERT INTO `ch_cities` VALUES (6, '河北省', '邯郸市', 130400);
INSERT INTO `ch_cities` VALUES (7, '河北省', '邢台市', 130500);
INSERT INTO `ch_cities` VALUES (8, '河北省', '保定市', 130600);
INSERT INTO `ch_cities` VALUES (9, '河北省', '张家口市', 130700);
INSERT INTO `ch_cities` VALUES (10, '河北省', '承德市', 130800);
INSERT INTO `ch_cities` VALUES (11, '河北省', '沧州市', 130900);
INSERT INTO `ch_cities` VALUES (12, '河北省', '廊坊市', 131000);
INSERT INTO `ch_cities` VALUES (13, '河北省', '衡水市', 131100);
INSERT INTO `ch_cities` VALUES (14, '山西省', '太原市', 140100);
INSERT INTO `ch_cities` VALUES (15, '山西省', '大同市', 140200);
INSERT INTO `ch_cities` VALUES (16, '山西省', '阳泉市', 140300);
INSERT INTO `ch_cities` VALUES (17, '山西省', '长治市', 140400);
INSERT INTO `ch_cities` VALUES (18, '山西省', '晋城市', 140500);
INSERT INTO `ch_cities` VALUES (19, '山西省', '朔州市', 140600);
INSERT INTO `ch_cities` VALUES (20, '山西省', '晋中市', 140700);
INSERT INTO `ch_cities` VALUES (21, '山西省', '运城市', 140800);
INSERT INTO `ch_cities` VALUES (22, '山西省', '忻州市', 140900);
INSERT INTO `ch_cities` VALUES (23, '山西省', '临汾市', 141000);
INSERT INTO `ch_cities` VALUES (24, '山西省', '吕梁市', 141100);
INSERT INTO `ch_cities` VALUES (25, '内蒙古自治区', '呼和浩特市', 150100);
INSERT INTO `ch_cities` VALUES (26, '内蒙古自治区', '包头市', 150200);
INSERT INTO `ch_cities` VALUES (27, '内蒙古自治区', '乌海市', 150300);
INSERT INTO `ch_cities` VALUES (28, '内蒙古自治区', '赤峰市', 150400);
INSERT INTO `ch_cities` VALUES (29, '内蒙古自治区', '通辽市', 150500);
INSERT INTO `ch_cities` VALUES (30, '内蒙古自治区', '鄂尔多斯市', 150600);
INSERT INTO `ch_cities` VALUES (31, '内蒙古自治区', '呼伦贝尔市', 150700);
INSERT INTO `ch_cities` VALUES (32, '内蒙古自治区', '巴彦淖尔市', 150800);
INSERT INTO `ch_cities` VALUES (33, '内蒙古自治区', '乌兰察布市', 150900);
INSERT INTO `ch_cities` VALUES (34, '内蒙古自治区', '兴安盟', 152200);
INSERT INTO `ch_cities` VALUES (35, '内蒙古自治区', '锡林郭勒盟', 152500);
INSERT INTO `ch_cities` VALUES (36, '内蒙古自治区', '二连浩特市', 152501);
INSERT INTO `ch_cities` VALUES (37, '内蒙古自治区', '锡林浩特市', 152502);
INSERT INTO `ch_cities` VALUES (38, '内蒙古自治区', '阿拉善盟', 152900);
INSERT INTO `ch_cities` VALUES (39, '辽宁省', '沈阳市', 210100);
INSERT INTO `ch_cities` VALUES (40, '辽宁省', '大连市', 210200);
INSERT INTO `ch_cities` VALUES (41, '辽宁省', '鞍山市', 210300);
INSERT INTO `ch_cities` VALUES (42, '辽宁省', '抚顺市', 210400);
INSERT INTO `ch_cities` VALUES (43, '辽宁省', '本溪市', 210500);
INSERT INTO `ch_cities` VALUES (44, '辽宁省', '丹东市', 210600);
INSERT INTO `ch_cities` VALUES (45, '辽宁省', '锦州市', 210700);
INSERT INTO `ch_cities` VALUES (46, '辽宁省', '营口市', 210800);
INSERT INTO `ch_cities` VALUES (47, '辽宁省', '阜新市', 210900);
INSERT INTO `ch_cities` VALUES (48, '辽宁省', '辽阳市', 211000);
INSERT INTO `ch_cities` VALUES (49, '辽宁省', '盘锦市', 211100);
INSERT INTO `ch_cities` VALUES (50, '辽宁省', '铁岭市', 211200);
INSERT INTO `ch_cities` VALUES (51, '辽宁省', '朝阳市', 211300);
INSERT INTO `ch_cities` VALUES (52, '辽宁省', '葫芦岛市', 211400);
INSERT INTO `ch_cities` VALUES (53, '吉林省', '长春市', 220100);
INSERT INTO `ch_cities` VALUES (54, '吉林省', '吉林市', 220200);
INSERT INTO `ch_cities` VALUES (55, '吉林省', '四平市', 220300);
INSERT INTO `ch_cities` VALUES (56, '吉林省', '辽源市', 220400);
INSERT INTO `ch_cities` VALUES (57, '吉林省', '通化市', 220500);
INSERT INTO `ch_cities` VALUES (58, '吉林省', '白山市', 220600);
INSERT INTO `ch_cities` VALUES (59, '吉林省', '松原市', 220700);
INSERT INTO `ch_cities` VALUES (60, '吉林省', '白城市', 220800);
INSERT INTO `ch_cities` VALUES (61, '吉林省', '延边朝鲜族自治州', 222400);
INSERT INTO `ch_cities` VALUES (62, '黑龙江省', '哈尔滨市', 230100);
INSERT INTO `ch_cities` VALUES (63, '黑龙江省', '齐齐哈尔市', 230200);
INSERT INTO `ch_cities` VALUES (64, '黑龙江省', '鸡西市', 230300);
INSERT INTO `ch_cities` VALUES (65, '黑龙江省', '鹤岗市', 230400);
INSERT INTO `ch_cities` VALUES (66, '黑龙江省', '双鸭山市', 230500);
INSERT INTO `ch_cities` VALUES (67, '黑龙江省', '大庆市', 230600);
INSERT INTO `ch_cities` VALUES (68, '黑龙江省', '伊春市', 230700);
INSERT INTO `ch_cities` VALUES (69, '黑龙江省', '佳木斯市', 230800);
INSERT INTO `ch_cities` VALUES (70, '黑龙江省', '七台河市', 230900);
INSERT INTO `ch_cities` VALUES (71, '黑龙江省', '牡丹江市', 231000);
INSERT INTO `ch_cities` VALUES (72, '黑龙江省', '黑河市', 231100);
INSERT INTO `ch_cities` VALUES (73, '黑龙江省', '绥化市', 231200);
INSERT INTO `ch_cities` VALUES (74, '黑龙江省', '大兴安岭地区', 232700);
INSERT INTO `ch_cities` VALUES (75, '上海市', '上海市', 310100);
INSERT INTO `ch_cities` VALUES (76, '江苏省', '南京市', 320100);
INSERT INTO `ch_cities` VALUES (77, '江苏省', '无锡市', 320200);
INSERT INTO `ch_cities` VALUES (78, '江苏省', '徐州市', 320300);
INSERT INTO `ch_cities` VALUES (79, '江苏省', '常州市', 320400);
INSERT INTO `ch_cities` VALUES (80, '江苏省', '苏州市', 320500);
INSERT INTO `ch_cities` VALUES (81, '江苏省', '昆山市', 320583);
INSERT INTO `ch_cities` VALUES (82, '江苏省', '南通市', 320600);
INSERT INTO `ch_cities` VALUES (83, '江苏省', '连云港市', 320700);
INSERT INTO `ch_cities` VALUES (84, '江苏省', '淮安市', 320800);
INSERT INTO `ch_cities` VALUES (85, '江苏省', '盐城市', 320900);
INSERT INTO `ch_cities` VALUES (86, '江苏省', '扬州市', 321000);
INSERT INTO `ch_cities` VALUES (87, '江苏省', '镇江市', 321100);
INSERT INTO `ch_cities` VALUES (88, '江苏省', '泰州市', 321200);
INSERT INTO `ch_cities` VALUES (89, '江苏省', '宿迁市', 321300);
INSERT INTO `ch_cities` VALUES (90, '浙江省', '杭州市', 330100);
INSERT INTO `ch_cities` VALUES (91, '浙江省', '宁波市', 330200);
INSERT INTO `ch_cities` VALUES (92, '浙江省', '温州市', 330300);
INSERT INTO `ch_cities` VALUES (93, '浙江省', '嘉兴市', 330400);
INSERT INTO `ch_cities` VALUES (94, '浙江省', '湖州市', 330500);
INSERT INTO `ch_cities` VALUES (95, '浙江省', '绍兴市', 330600);
INSERT INTO `ch_cities` VALUES (96, '浙江省', '金华市', 330700);
INSERT INTO `ch_cities` VALUES (97, '浙江省', '衢州市', 330800);
INSERT INTO `ch_cities` VALUES (98, '浙江省', '舟山市', 330900);
INSERT INTO `ch_cities` VALUES (99, '浙江省', '台州市', 331000);
INSERT INTO `ch_cities` VALUES (100, '浙江省', '丽水市', 331100);
INSERT INTO `ch_cities` VALUES (101, '安徽省', '合肥市', 340100);
INSERT INTO `ch_cities` VALUES (102, '安徽省', '芜湖市', 340200);
INSERT INTO `ch_cities` VALUES (103, '安徽省', '蚌埠市', 340300);
INSERT INTO `ch_cities` VALUES (104, '安徽省', '淮南市', 340400);
INSERT INTO `ch_cities` VALUES (105, '安徽省', '马鞍山市', 340500);
INSERT INTO `ch_cities` VALUES (106, '安徽省', '淮北市', 340600);
INSERT INTO `ch_cities` VALUES (107, '安徽省', '铜陵市', 340700);
INSERT INTO `ch_cities` VALUES (108, '安徽省', '安庆市', 340800);
INSERT INTO `ch_cities` VALUES (109, '安徽省', '黄山市', 341000);
INSERT INTO `ch_cities` VALUES (110, '安徽省', '滁州市', 341100);
INSERT INTO `ch_cities` VALUES (111, '安徽省', '阜阳市', 341200);
INSERT INTO `ch_cities` VALUES (112, '安徽省', '宿州市', 341300);
INSERT INTO `ch_cities` VALUES (113, '安徽省', '六安市', 341500);
INSERT INTO `ch_cities` VALUES (114, '安徽省', '亳州市', 341600);
INSERT INTO `ch_cities` VALUES (115, '安徽省', '池州市', 341700);
INSERT INTO `ch_cities` VALUES (116, '安徽省', '宣城市', 341800);
INSERT INTO `ch_cities` VALUES (117, '福建省', '福州市', 350100);
INSERT INTO `ch_cities` VALUES (118, '福建省', '厦门市', 350200);
INSERT INTO `ch_cities` VALUES (119, '福建省', '莆田市', 350300);
INSERT INTO `ch_cities` VALUES (120, '福建省', '三明市', 350400);
INSERT INTO `ch_cities` VALUES (121, '福建省', '泉州市', 350500);
INSERT INTO `ch_cities` VALUES (122, '福建省', '漳州市', 350600);
INSERT INTO `ch_cities` VALUES (123, '福建省', '南平市', 350700);
INSERT INTO `ch_cities` VALUES (124, '福建省', '龙岩市', 350800);
INSERT INTO `ch_cities` VALUES (125, '福建省', '宁德市', 350900);
INSERT INTO `ch_cities` VALUES (126, '江西省', '南昌市', 360100);
INSERT INTO `ch_cities` VALUES (127, '江西省', '景德镇市', 360200);
INSERT INTO `ch_cities` VALUES (128, '江西省', '萍乡市', 360300);
INSERT INTO `ch_cities` VALUES (129, '江西省', '九江市', 360400);
INSERT INTO `ch_cities` VALUES (130, '江西省', '新余市', 360500);
INSERT INTO `ch_cities` VALUES (131, '江西省', '鹰潭市', 360600);
INSERT INTO `ch_cities` VALUES (132, '江西省', '赣州市', 360700);
INSERT INTO `ch_cities` VALUES (133, '江西省', '吉安市', 360800);
INSERT INTO `ch_cities` VALUES (134, '江西省', '宜春市', 360900);
INSERT INTO `ch_cities` VALUES (135, '江西省', '抚州市', 361000);
INSERT INTO `ch_cities` VALUES (136, '江西省', '上饶市', 361100);
INSERT INTO `ch_cities` VALUES (137, '山东省', '济南市', 370100);
INSERT INTO `ch_cities` VALUES (138, '山东省', '青岛市', 370200);
INSERT INTO `ch_cities` VALUES (139, '山东省', '淄博市', 370300);
INSERT INTO `ch_cities` VALUES (140, '山东省', '枣庄市', 370400);
INSERT INTO `ch_cities` VALUES (141, '山东省', '滕州市', 370481);
INSERT INTO `ch_cities` VALUES (142, '山东省', '东营市', 370500);
INSERT INTO `ch_cities` VALUES (143, '山东省', '烟台市', 370600);
INSERT INTO `ch_cities` VALUES (144, '山东省', '潍坊市', 370700);
INSERT INTO `ch_cities` VALUES (145, '山东省', '济宁市', 370800);
INSERT INTO `ch_cities` VALUES (146, '山东省', '泰安市', 370900);
INSERT INTO `ch_cities` VALUES (147, '山东省', '威海市', 371000);
INSERT INTO `ch_cities` VALUES (148, '山东省', '日照市', 371100);
INSERT INTO `ch_cities` VALUES (149, '山东省', '莱芜市', 371200);
INSERT INTO `ch_cities` VALUES (150, '山东省', '临沂市', 371300);
INSERT INTO `ch_cities` VALUES (151, '山东省', '德州市', 371400);
INSERT INTO `ch_cities` VALUES (152, '山东省', '聊城市', 371500);
INSERT INTO `ch_cities` VALUES (153, '山东省', '滨州市', 371600);
INSERT INTO `ch_cities` VALUES (154, '山东省', '菏泽市', 371700);
INSERT INTO `ch_cities` VALUES (155, '河南省', '郑州市', 410100);
INSERT INTO `ch_cities` VALUES (156, '河南省', '开封市', 410200);
INSERT INTO `ch_cities` VALUES (157, '河南省', '洛阳市', 410300);
INSERT INTO `ch_cities` VALUES (158, '河南省', '平顶山市', 410400);
INSERT INTO `ch_cities` VALUES (159, '河南省', '安阳市', 410500);
INSERT INTO `ch_cities` VALUES (160, '河南省', '鹤壁市', 410600);
INSERT INTO `ch_cities` VALUES (161, '河南省', '新乡市', 410700);
INSERT INTO `ch_cities` VALUES (162, '河南省', '焦作市', 410800);
INSERT INTO `ch_cities` VALUES (163, '河南省', '濮阳市', 410900);
INSERT INTO `ch_cities` VALUES (164, '河南省', '许昌市', 411000);
INSERT INTO `ch_cities` VALUES (165, '河南省', '漯河市', 411100);
INSERT INTO `ch_cities` VALUES (166, '河南省', '三门峡市', 411200);
INSERT INTO `ch_cities` VALUES (167, '河南省', '南阳市', 411300);
INSERT INTO `ch_cities` VALUES (168, '河南省', '商丘市', 411400);
INSERT INTO `ch_cities` VALUES (169, '河南省', '信阳市', 411500);
INSERT INTO `ch_cities` VALUES (170, '河南省', '周口市', 411600);
INSERT INTO `ch_cities` VALUES (171, '河南省', '驻马店市', 411700);
INSERT INTO `ch_cities` VALUES (172, '河南省', '济源市', 419001);
INSERT INTO `ch_cities` VALUES (173, '湖北省', '武汉市', 420100);
INSERT INTO `ch_cities` VALUES (174, '湖北省', '黄石市', 420200);
INSERT INTO `ch_cities` VALUES (175, '湖北省', '十堰市', 420300);
INSERT INTO `ch_cities` VALUES (176, '湖北省', '宜昌市', 420500);
INSERT INTO `ch_cities` VALUES (177, '湖北省', '襄阳市', 420600);
INSERT INTO `ch_cities` VALUES (178, '湖北省', '鄂州市', 420700);
INSERT INTO `ch_cities` VALUES (179, '湖北省', '荆门市', 420800);
INSERT INTO `ch_cities` VALUES (180, '湖北省', '孝感市', 420900);
INSERT INTO `ch_cities` VALUES (181, '湖北省', '荆州市', 421000);
INSERT INTO `ch_cities` VALUES (182, '湖北省', '黄冈市', 421100);
INSERT INTO `ch_cities` VALUES (183, '湖北省', '咸宁市', 421200);
INSERT INTO `ch_cities` VALUES (184, '湖北省', '随州市', 421300);
INSERT INTO `ch_cities` VALUES (185, '湖北省', '恩施土家族苗族自治州', 422800);
INSERT INTO `ch_cities` VALUES (186, '湖北省', '潜江市', 429005);
INSERT INTO `ch_cities` VALUES (187, '湖南省', '长沙市', 430100);
INSERT INTO `ch_cities` VALUES (188, '湖南省', '株洲市', 430200);
INSERT INTO `ch_cities` VALUES (189, '湖南省', '湘潭市', 430300);
INSERT INTO `ch_cities` VALUES (190, '湖南省', '衡阳市', 430400);
INSERT INTO `ch_cities` VALUES (191, '湖南省', '邵阳市', 430500);
INSERT INTO `ch_cities` VALUES (192, '湖南省', '岳阳市', 430600);
INSERT INTO `ch_cities` VALUES (193, '湖南省', '常德市', 430700);
INSERT INTO `ch_cities` VALUES (194, '湖南省', '张家界市', 430800);
INSERT INTO `ch_cities` VALUES (195, '湖南省', '益阳市', 430900);
INSERT INTO `ch_cities` VALUES (196, '湖南省', '郴州市', 431000);
INSERT INTO `ch_cities` VALUES (197, '湖南省', '永州市', 431100);
INSERT INTO `ch_cities` VALUES (198, '湖南省', '怀化市', 431200);
INSERT INTO `ch_cities` VALUES (199, '湖南省', '娄底市', 431300);
INSERT INTO `ch_cities` VALUES (200, '湖南省', '湘西土家族苗族自治州', 433100);
INSERT INTO `ch_cities` VALUES (201, '广东省', '广州市', 440100);
INSERT INTO `ch_cities` VALUES (202, '广东省', '韶关市', 440200);
INSERT INTO `ch_cities` VALUES (203, '广东省', '深圳市', 440300);
INSERT INTO `ch_cities` VALUES (204, '广东省', '珠海市', 440400);
INSERT INTO `ch_cities` VALUES (205, '广东省', '汕头市', 440500);
INSERT INTO `ch_cities` VALUES (206, '广东省', '佛山市', 440600);
INSERT INTO `ch_cities` VALUES (207, '广东省', '江门市', 440700);
INSERT INTO `ch_cities` VALUES (208, '广东省', '湛江市', 440800);
INSERT INTO `ch_cities` VALUES (209, '广东省', '茂名市', 440900);
INSERT INTO `ch_cities` VALUES (210, '广东省', '肇庆市', 441200);
INSERT INTO `ch_cities` VALUES (211, '广东省', '惠州市', 441300);
INSERT INTO `ch_cities` VALUES (212, '广东省', '梅州市', 441400);
INSERT INTO `ch_cities` VALUES (213, '广东省', '汕尾市', 441500);
INSERT INTO `ch_cities` VALUES (214, '广东省', '河源市', 441600);
INSERT INTO `ch_cities` VALUES (215, '广东省', '阳江市', 441700);
INSERT INTO `ch_cities` VALUES (216, '广东省', '清远市', 441800);
INSERT INTO `ch_cities` VALUES (217, '广东省', '东莞市', 441900);
INSERT INTO `ch_cities` VALUES (218, '广东省', '中山市', 442000);
INSERT INTO `ch_cities` VALUES (219, '广东省', '潮州市', 445100);
INSERT INTO `ch_cities` VALUES (220, '广东省', '揭阳市', 445200);
INSERT INTO `ch_cities` VALUES (221, '广东省', '云浮市', 445300);
INSERT INTO `ch_cities` VALUES (222, '广西壮族自治区', '南宁市', 450100);
INSERT INTO `ch_cities` VALUES (223, '广西壮族自治区', '柳州市', 450200);
INSERT INTO `ch_cities` VALUES (224, '广西壮族自治区', '桂林市', 450300);
INSERT INTO `ch_cities` VALUES (225, '广西壮族自治区', '梧州市', 450400);
INSERT INTO `ch_cities` VALUES (226, '广西壮族自治区', '北海市', 450500);
INSERT INTO `ch_cities` VALUES (227, '广西壮族自治区', '防城港市', 450600);
INSERT INTO `ch_cities` VALUES (228, '广西壮族自治区', '钦州市', 450700);
INSERT INTO `ch_cities` VALUES (229, '广西壮族自治区', '贵港市', 450800);
INSERT INTO `ch_cities` VALUES (230, '广西壮族自治区', '玉林市', 450900);
INSERT INTO `ch_cities` VALUES (231, '广西壮族自治区', '百色市', 451000);
INSERT INTO `ch_cities` VALUES (232, '广西壮族自治区', '贺州市', 451100);
INSERT INTO `ch_cities` VALUES (233, '广西壮族自治区', '河池市', 451200);
INSERT INTO `ch_cities` VALUES (234, '广西壮族自治区', '来宾市', 451300);
INSERT INTO `ch_cities` VALUES (235, '广西壮族自治区', '崇左市', 451400);
INSERT INTO `ch_cities` VALUES (236, '海南省', '海口市', 460100);
INSERT INTO `ch_cities` VALUES (237, '海南省', '三亚市', 460200);
INSERT INTO `ch_cities` VALUES (238, '海南省', '儋州市', 460400);
INSERT INTO `ch_cities` VALUES (239, '重庆市', '重庆市', 500100);
INSERT INTO `ch_cities` VALUES (240, '四川省', '成都市', 510100);
INSERT INTO `ch_cities` VALUES (241, '四川省', '自贡市', 510300);
INSERT INTO `ch_cities` VALUES (242, '四川省', '攀枝花市', 510400);
INSERT INTO `ch_cities` VALUES (243, '四川省', '泸州市', 510500);
INSERT INTO `ch_cities` VALUES (244, '四川省', '德阳市', 510600);
INSERT INTO `ch_cities` VALUES (245, '四川省', '绵阳市', 510700);
INSERT INTO `ch_cities` VALUES (246, '四川省', '广元市', 510800);
INSERT INTO `ch_cities` VALUES (247, '四川省', '遂宁市', 510900);
INSERT INTO `ch_cities` VALUES (248, '四川省', '内江市', 511000);
INSERT INTO `ch_cities` VALUES (249, '四川省', '乐山市', 511100);
INSERT INTO `ch_cities` VALUES (250, '四川省', '南充市', 511300);
INSERT INTO `ch_cities` VALUES (251, '四川省', '眉山市', 511400);
INSERT INTO `ch_cities` VALUES (252, '四川省', '宜宾市', 511500);
INSERT INTO `ch_cities` VALUES (253, '四川省', '广安市', 511600);
INSERT INTO `ch_cities` VALUES (254, '四川省', '达州市', 511700);
INSERT INTO `ch_cities` VALUES (255, '四川省', '雅安市', 511800);
INSERT INTO `ch_cities` VALUES (256, '四川省', '巴中市', 511900);
INSERT INTO `ch_cities` VALUES (257, '四川省', '资阳市', 512000);
INSERT INTO `ch_cities` VALUES (258, '四川省', '阿坝藏族羌族自治州', 513200);
INSERT INTO `ch_cities` VALUES (259, '四川省', '甘孜藏族自治州', 513300);
INSERT INTO `ch_cities` VALUES (260, '四川省', '凉山彝族自治州', 513400);
INSERT INTO `ch_cities` VALUES (261, '贵州省', '贵阳市', 520100);
INSERT INTO `ch_cities` VALUES (262, '贵州省', '六盘水市', 520200);
INSERT INTO `ch_cities` VALUES (263, '贵州省', '遵义市', 520300);
INSERT INTO `ch_cities` VALUES (264, '贵州省', '安顺市', 520400);
INSERT INTO `ch_cities` VALUES (265, '贵州省', '毕节市', 520500);
INSERT INTO `ch_cities` VALUES (266, '贵州省', '铜仁市', 520600);
INSERT INTO `ch_cities` VALUES (267, '贵州省', '黔西南布依族苗族自治州', 522300);
INSERT INTO `ch_cities` VALUES (268, '贵州省', '黔东南苗族侗族自治州', 522600);
INSERT INTO `ch_cities` VALUES (269, '贵州省', '黔南布依族苗族自治州', 522700);
INSERT INTO `ch_cities` VALUES (270, '云南省', '昆明市', 530100);
INSERT INTO `ch_cities` VALUES (271, '云南省', '曲靖市', 530300);
INSERT INTO `ch_cities` VALUES (272, '云南省', '玉溪市', 530400);
INSERT INTO `ch_cities` VALUES (273, '云南省', '保山市', 530500);
INSERT INTO `ch_cities` VALUES (274, '云南省', '昭通市', 530600);
INSERT INTO `ch_cities` VALUES (275, '云南省', '丽江市', 530700);
INSERT INTO `ch_cities` VALUES (276, '云南省', '普洱市', 530800);
INSERT INTO `ch_cities` VALUES (277, '云南省', '临沧市', 530900);
INSERT INTO `ch_cities` VALUES (278, '云南省', '楚雄彝族自治州', 532300);
INSERT INTO `ch_cities` VALUES (279, '云南省', '红河哈尼族彝族自治州', 532500);
INSERT INTO `ch_cities` VALUES (280, '云南省', '文山壮族苗族自治州', 532600);
INSERT INTO `ch_cities` VALUES (281, '云南省', '西双版纳傣族自治州', 532800);
INSERT INTO `ch_cities` VALUES (282, '云南省', '大理白族自治州', 532900);
INSERT INTO `ch_cities` VALUES (283, '云南省', '德宏傣族景颇族自治州', 533100);
INSERT INTO `ch_cities` VALUES (284, '云南省', '怒江傈僳族自治州', 533300);
INSERT INTO `ch_cities` VALUES (285, '云南省', '迪庆藏族自治州', 533400);
INSERT INTO `ch_cities` VALUES (286, '西藏自治区', '拉萨市', 540100);
INSERT INTO `ch_cities` VALUES (287, '西藏自治区', '日喀则市', 540200);
INSERT INTO `ch_cities` VALUES (288, '陕西省', '西安市', 610100);
INSERT INTO `ch_cities` VALUES (289, '陕西省', '铜川市', 610200);
INSERT INTO `ch_cities` VALUES (290, '陕西省', '宝鸡市', 610300);
INSERT INTO `ch_cities` VALUES (291, '陕西省', '咸阳市', 610400);
INSERT INTO `ch_cities` VALUES (292, '陕西省', '渭南市', 610500);
INSERT INTO `ch_cities` VALUES (293, '陕西省', '延安市', 610600);
INSERT INTO `ch_cities` VALUES (294, '陕西省', '汉中市', 610700);
INSERT INTO `ch_cities` VALUES (295, '陕西省', '榆林市', 610800);
INSERT INTO `ch_cities` VALUES (296, '陕西省', '安康市', 610900);
INSERT INTO `ch_cities` VALUES (297, '陕西省', '商洛市', 611000);
INSERT INTO `ch_cities` VALUES (298, '甘肃省', '兰州市', 620100);
INSERT INTO `ch_cities` VALUES (299, '甘肃省', '嘉峪关市', 620200);
INSERT INTO `ch_cities` VALUES (300, '甘肃省', '金昌市', 620300);
INSERT INTO `ch_cities` VALUES (301, '甘肃省', '白银市', 620400);
INSERT INTO `ch_cities` VALUES (302, '甘肃省', '天水市', 620500);
INSERT INTO `ch_cities` VALUES (303, '甘肃省', '武威市', 620600);
INSERT INTO `ch_cities` VALUES (304, '甘肃省', '张掖市', 620700);
INSERT INTO `ch_cities` VALUES (305, '甘肃省', '平凉市', 620800);
INSERT INTO `ch_cities` VALUES (306, '甘肃省', '酒泉市', 620900);
INSERT INTO `ch_cities` VALUES (307, '甘肃省', '庆阳市', 621000);
INSERT INTO `ch_cities` VALUES (308, '甘肃省', '定西市', 621100);
INSERT INTO `ch_cities` VALUES (309, '甘肃省', '陇南市', 621200);
INSERT INTO `ch_cities` VALUES (310, '甘肃省', '临夏回族自治州', 622900);
INSERT INTO `ch_cities` VALUES (311, '甘肃省', '甘南藏族自治州', 623000);
INSERT INTO `ch_cities` VALUES (312, '青海省', '西宁市', 630100);
INSERT INTO `ch_cities` VALUES (313, '青海省', '海东市', 630200);
INSERT INTO `ch_cities` VALUES (314, '青海省', '海西蒙古族藏族自治州', 632800);
INSERT INTO `ch_cities` VALUES (315, '宁夏回族自治区', '银川市', 640100);
INSERT INTO `ch_cities` VALUES (316, '宁夏回族自治区', '石嘴山市', 640200);
INSERT INTO `ch_cities` VALUES (317, '宁夏回族自治区', '吴忠市', 640300);
INSERT INTO `ch_cities` VALUES (318, '新疆维吾尔自治区', '乌鲁木齐市', 650100);
INSERT INTO `ch_cities` VALUES (319, '新疆维吾尔自治区', '克拉玛依市', 650200);
INSERT INTO `ch_cities` VALUES (320, '新疆维吾尔自治区', '哈密市', 650500);
INSERT INTO `ch_cities` VALUES (321, '新疆维吾尔自治区', '昌吉回族自治州', 652300);
INSERT INTO `ch_cities` VALUES (322, '新疆维吾尔自治区', '巴音郭楞蒙古自治州', 652800);
INSERT INTO `ch_cities` VALUES (323, '新疆维吾尔自治区', '阿克苏地区', 652900);
INSERT INTO `ch_cities` VALUES (324, '新疆维吾尔自治区', '喀什地区', 653100);
INSERT INTO `ch_cities` VALUES (325, '新疆维吾尔自治区', '伊犁哈萨克自治州', 654000);

-- ----------------------------
-- Table structure for ch_provinces
-- ----------------------------
DROP TABLE IF EXISTS `ch_provinces`;
CREATE TABLE `ch_provinces`  (
  `id` int(255) NULL DEFAULT NULL,
  `province` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of ch_provinces
-- ----------------------------
INSERT INTO `ch_provinces` VALUES (1, '北京市');
INSERT INTO `ch_provinces` VALUES (2, '天津市');
INSERT INTO `ch_provinces` VALUES (3, '河北省');
INSERT INTO `ch_provinces` VALUES (4, '山西省');
INSERT INTO `ch_provinces` VALUES (5, '内蒙古自治区');
INSERT INTO `ch_provinces` VALUES (6, '辽宁省');
INSERT INTO `ch_provinces` VALUES (7, '吉林省');
INSERT INTO `ch_provinces` VALUES (8, '黑龙江省');
INSERT INTO `ch_provinces` VALUES (9, '上海市');
INSERT INTO `ch_provinces` VALUES (10, '江苏省');
INSERT INTO `ch_provinces` VALUES (11, '浙江省');
INSERT INTO `ch_provinces` VALUES (12, '安徽省');
INSERT INTO `ch_provinces` VALUES (13, '福建省');
INSERT INTO `ch_provinces` VALUES (14, '江西省');
INSERT INTO `ch_provinces` VALUES (15, '山东省');
INSERT INTO `ch_provinces` VALUES (16, '河南省');
INSERT INTO `ch_provinces` VALUES (17, '湖北省');
INSERT INTO `ch_provinces` VALUES (18, '湖南省');
INSERT INTO `ch_provinces` VALUES (19, '广东省');
INSERT INTO `ch_provinces` VALUES (20, '广西壮族自治区');
INSERT INTO `ch_provinces` VALUES (21, '海南省');
INSERT INTO `ch_provinces` VALUES (22, '重庆市');
INSERT INTO `ch_provinces` VALUES (23, '四川省');
INSERT INTO `ch_provinces` VALUES (24, '贵州省');
INSERT INTO `ch_provinces` VALUES (25, '云南省');
INSERT INTO `ch_provinces` VALUES (26, '西藏自治区');
INSERT INTO `ch_provinces` VALUES (27, '陕西省');
INSERT INTO `ch_provinces` VALUES (28, '甘肃省');
INSERT INTO `ch_provinces` VALUES (29, '青海省');
INSERT INTO `ch_provinces` VALUES (30, '宁夏回族自治区');
INSERT INTO `ch_provinces` VALUES (31, '新疆维吾尔自治区');

-- ----------------------------
-- Table structure for ch_t_d_areainfo
-- ----------------------------
DROP TABLE IF EXISTS `ch_t_d_areainfo`;
CREATE TABLE `ch_t_d_areainfo`  (
  `id` int(11) NOT NULL,
  `name` varchar(48) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '名称',
  `arealevel` tinyint(2) NOT NULL COMMENT '层级标识： 1  省份， 2  市， 3  区县',
  `parent_id` int(11) UNSIGNED ZEROFILL NULL DEFAULT NULL COMMENT '父节点',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of ch_t_d_areainfo
-- ----------------------------
INSERT INTO `ch_t_d_areainfo` VALUES (102100, '天津市', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (110100, '北京市', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (110101, '东城区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110102, '西城区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110105, '朝阳区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110106, '丰台区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110107, '石景山区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110108, '海淀区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110109, '门头沟区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110111, '房山区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110112, '通州区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110113, '顺义区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110114, '昌平区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110115, '大兴区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110116, '怀柔区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110117, '平谷区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110118, '密云区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (110119, '延庆区', 3, 00000110100);
INSERT INTO `ch_t_d_areainfo` VALUES (120101, '和平区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120102, '河东区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120103, '河西区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120104, '南开区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120105, '河北区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120106, '红桥区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120110, '东丽区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120111, '西青区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120112, '津南区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120113, '北辰区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120114, '武清区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120115, '宝坻区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120116, '滨海新区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120117, '宁河区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120118, '静海区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (120119, '蓟州区', 3, 00000120100);
INSERT INTO `ch_t_d_areainfo` VALUES (130000, '河北省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (130100, '石家庄市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (130102, '长安区', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130104, '桥西区', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130105, '新华区', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130107, '井陉矿区', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130108, '裕华区', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130109, '藁城区', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130110, '鹿泉区', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130111, '栾城区', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130121, '井陉县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130123, '正定县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130125, '行唐县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130126, '灵寿县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130127, '高邑县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130128, '深泽县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130129, '赞皇县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130130, '无极县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130131, '平山县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130132, '元氏县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130133, '赵县', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130183, '晋州市', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130184, '新乐市', 3, 00000130100);
INSERT INTO `ch_t_d_areainfo` VALUES (130200, '唐山市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (130202, '路南区', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130203, '路北区', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130204, '古冶区', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130205, '开平区', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130207, '丰南区', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130208, '丰润区', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130209, '曹妃甸区', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130223, '滦县', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130224, '滦南县', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130225, '乐亭县', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130227, '迁西县', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130229, '玉田县', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130281, '遵化市', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130283, '迁安市', 3, 00000130200);
INSERT INTO `ch_t_d_areainfo` VALUES (130300, '秦皇岛市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (130302, '海港区', 3, 00000130300);
INSERT INTO `ch_t_d_areainfo` VALUES (130303, '山海关区', 3, 00000130300);
INSERT INTO `ch_t_d_areainfo` VALUES (130304, '北戴河区', 3, 00000130300);
INSERT INTO `ch_t_d_areainfo` VALUES (130306, '抚宁区', 3, 00000130300);
INSERT INTO `ch_t_d_areainfo` VALUES (130321, '青龙满族自治县', 3, 00000130300);
INSERT INTO `ch_t_d_areainfo` VALUES (130322, '昌黎县', 3, 00000130300);
INSERT INTO `ch_t_d_areainfo` VALUES (130324, '卢龙县', 3, 00000130300);
INSERT INTO `ch_t_d_areainfo` VALUES (130400, '邯郸市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (130402, '邯山区', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130403, '丛台区', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130404, '复兴区', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130406, '峰峰矿区', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130421, '邯郸县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130423, '临漳县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130424, '成安县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130425, '大名县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130426, '涉县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130427, '磁县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130428, '肥乡县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130429, '永年县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130430, '邱县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130431, '鸡泽县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130432, '广平县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130433, '馆陶县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130434, '魏县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130435, '曲周县', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130481, '武安市', 3, 00000130400);
INSERT INTO `ch_t_d_areainfo` VALUES (130500, '邢台市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (130502, '桥东区', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130503, '桥西区', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130521, '邢台县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130522, '临城县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130523, '内丘县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130524, '柏乡县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130525, '隆尧县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130526, '任县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130527, '南和县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130528, '宁晋县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130529, '巨鹿县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130530, '新河县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130531, '广宗县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130532, '平乡县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130533, '威县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130534, '清河县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130535, '临西县', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130581, '南宫市', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130582, '沙河市', 3, 00000130500);
INSERT INTO `ch_t_d_areainfo` VALUES (130600, '保定市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (130602, '竞秀区', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130606, '莲池区', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130607, '满城区', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130608, '清苑区', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130609, '徐水区', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130623, '涞水县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130624, '阜平县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130626, '定兴县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130627, '唐县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130628, '高阳县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130629, '容城县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130630, '涞源县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130631, '望都县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130632, '安新县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130633, '易县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130634, '曲阳县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130635, '蠡县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130636, '顺平县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130637, '博野县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130638, '雄县', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130681, '涿州市', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130683, '安国市', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130684, '高碑店市', 3, 00000130600);
INSERT INTO `ch_t_d_areainfo` VALUES (130700, '张家口市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (130702, '桥东区', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130703, '桥西区', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130705, '宣化区', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130706, '下花园区', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130708, '万全区', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130709, '崇礼区', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130722, '张北县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130723, '康保县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130724, '沽源县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130725, '尚义县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130726, '蔚县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130727, '阳原县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130728, '怀安县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130730, '怀来县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130731, '涿鹿县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130732, '赤城县', 3, 00000130700);
INSERT INTO `ch_t_d_areainfo` VALUES (130800, '承德市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (130802, '双桥区', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130803, '双滦区', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130804, '鹰手营子矿区', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130821, '承德县', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130822, '兴隆县', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130823, '平泉县', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130824, '滦平县', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130825, '隆化县', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130826, '丰宁满族自治县', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130827, '宽城满族自治县', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130828, '围场满族蒙古族自治县', 3, 00000130800);
INSERT INTO `ch_t_d_areainfo` VALUES (130900, '沧州市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (130902, '新华区', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130903, '运河区', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130921, '沧县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130922, '青县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130923, '东光县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130924, '海兴县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130925, '盐山县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130926, '肃宁县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130927, '南皮县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130928, '吴桥县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130929, '献县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130930, '孟村回族自治县', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130981, '泊头市', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130982, '任丘市', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130983, '黄骅市', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (130984, '河间市', 3, 00000130900);
INSERT INTO `ch_t_d_areainfo` VALUES (131000, '廊坊市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (131002, '安次区', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131003, '广阳区', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131022, '固安县', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131023, '永清县', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131024, '香河县', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131025, '大城县', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131026, '文安县', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131028, '大厂回族自治县', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131081, '霸州市', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131082, '三河市', 3, 00000131000);
INSERT INTO `ch_t_d_areainfo` VALUES (131100, '衡水市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (131102, '桃城区', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131103, '冀州区', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131121, '枣强县', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131122, '武邑县', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131123, '武强县', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131124, '饶阳县', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131125, '安平县', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131126, '故城县', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131127, '景县', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131128, '阜城县', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131182, '深州市', 3, 00000131100);
INSERT INTO `ch_t_d_areainfo` VALUES (131200, '定州市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (131300, '辛集市', 2, 00000130000);
INSERT INTO `ch_t_d_areainfo` VALUES (140000, '山西省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (140100, '太原市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (140105, '小店区', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140106, '迎泽区', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140107, '杏花岭区', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140108, '尖草坪区', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140109, '万柏林区', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140110, '晋源区', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140121, '清徐县', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140122, '阳曲县', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140123, '娄烦县', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140181, '古交市', 3, 00000140100);
INSERT INTO `ch_t_d_areainfo` VALUES (140200, '大同市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (140202, '城区', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140203, '矿区', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140211, '南郊区', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140212, '新荣区', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140221, '阳高县', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140222, '天镇县', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140223, '广灵县', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140224, '灵丘县', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140225, '浑源县', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140226, '左云县', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140227, '大同县', 3, 00000140200);
INSERT INTO `ch_t_d_areainfo` VALUES (140300, '阳泉市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (140302, '城区', 3, 00000140300);
INSERT INTO `ch_t_d_areainfo` VALUES (140303, '矿区', 3, 00000140300);
INSERT INTO `ch_t_d_areainfo` VALUES (140311, '郊区', 3, 00000140300);
INSERT INTO `ch_t_d_areainfo` VALUES (140321, '平定县', 3, 00000140300);
INSERT INTO `ch_t_d_areainfo` VALUES (140322, '盂县', 3, 00000140300);
INSERT INTO `ch_t_d_areainfo` VALUES (140400, '长治市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (140402, '城区', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140411, '郊区', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140421, '长治县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140423, '襄垣县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140424, '屯留县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140425, '平顺县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140426, '黎城县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140427, '壶关县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140428, '长子县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140429, '武乡县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140430, '沁县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140431, '沁源县', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140481, '潞城市', 3, 00000140400);
INSERT INTO `ch_t_d_areainfo` VALUES (140500, '晋城市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (140502, '城区', 3, 00000140500);
INSERT INTO `ch_t_d_areainfo` VALUES (140521, '沁水县', 3, 00000140500);
INSERT INTO `ch_t_d_areainfo` VALUES (140522, '阳城县', 3, 00000140500);
INSERT INTO `ch_t_d_areainfo` VALUES (140524, '陵川县', 3, 00000140500);
INSERT INTO `ch_t_d_areainfo` VALUES (140525, '泽州县', 3, 00000140500);
INSERT INTO `ch_t_d_areainfo` VALUES (140581, '高平市', 3, 00000140500);
INSERT INTO `ch_t_d_areainfo` VALUES (140600, '朔州市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (140602, '朔城区', 3, 00000140600);
INSERT INTO `ch_t_d_areainfo` VALUES (140603, '平鲁区', 3, 00000140600);
INSERT INTO `ch_t_d_areainfo` VALUES (140621, '山阴县', 3, 00000140600);
INSERT INTO `ch_t_d_areainfo` VALUES (140622, '应县', 3, 00000140600);
INSERT INTO `ch_t_d_areainfo` VALUES (140623, '右玉县', 3, 00000140600);
INSERT INTO `ch_t_d_areainfo` VALUES (140624, '怀仁县', 3, 00000140600);
INSERT INTO `ch_t_d_areainfo` VALUES (140700, '晋中市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (140702, '榆次区', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140721, '榆社县', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140722, '左权县', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140723, '和顺县', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140724, '昔阳县', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140725, '寿阳县', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140726, '太谷县', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140727, '祁县', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140728, '平遥县', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140729, '灵石县', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140781, '介休市', 3, 00000140700);
INSERT INTO `ch_t_d_areainfo` VALUES (140800, '运城市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (140802, '盐湖区', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140821, '临猗县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140822, '万荣县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140823, '闻喜县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140824, '稷山县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140825, '新绛县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140826, '绛县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140827, '垣曲县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140828, '夏县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140829, '平陆县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140830, '芮城县', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140881, '永济市', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140882, '河津市', 3, 00000140800);
INSERT INTO `ch_t_d_areainfo` VALUES (140900, '忻州市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (140902, '忻府区', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140921, '定襄县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140922, '五台县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140923, '代县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140924, '繁峙县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140925, '宁武县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140926, '静乐县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140927, '神池县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140928, '五寨县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140929, '岢岚县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140930, '河曲县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140931, '保德县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140932, '偏关县', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (140981, '原平市', 3, 00000140900);
INSERT INTO `ch_t_d_areainfo` VALUES (141000, '临汾市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (141002, '尧都区', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141021, '曲沃县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141022, '翼城县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141023, '襄汾县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141024, '洪洞县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141025, '古县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141026, '安泽县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141027, '浮山县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141028, '吉县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141029, '乡宁县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141030, '大宁县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141031, '隰县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141032, '永和县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141033, '蒲县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141034, '汾西县', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141081, '侯马市', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141082, '霍州市', 3, 00000141000);
INSERT INTO `ch_t_d_areainfo` VALUES (141100, '吕梁市', 2, 00000140000);
INSERT INTO `ch_t_d_areainfo` VALUES (141102, '离石区', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141121, '文水县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141122, '交城县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141123, '兴县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141124, '临县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141125, '柳林县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141126, '石楼县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141127, '岚县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141128, '方山县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141129, '中阳县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141130, '交口县', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141181, '孝义市', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (141182, '汾阳市', 3, 00000141100);
INSERT INTO `ch_t_d_areainfo` VALUES (150000, '内蒙古自治区', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (150100, '呼和浩特市', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (150102, '新城区', 3, 00000150100);
INSERT INTO `ch_t_d_areainfo` VALUES (150103, '回民区', 3, 00000150100);
INSERT INTO `ch_t_d_areainfo` VALUES (150104, '玉泉区', 3, 00000150100);
INSERT INTO `ch_t_d_areainfo` VALUES (150105, '赛罕区', 3, 00000150100);
INSERT INTO `ch_t_d_areainfo` VALUES (150121, '土默特左旗', 3, 00000150100);
INSERT INTO `ch_t_d_areainfo` VALUES (150122, '托克托县', 3, 00000150100);
INSERT INTO `ch_t_d_areainfo` VALUES (150123, '和林格尔县', 3, 00000150100);
INSERT INTO `ch_t_d_areainfo` VALUES (150124, '清水河县', 3, 00000150100);
INSERT INTO `ch_t_d_areainfo` VALUES (150125, '武川县', 3, 00000150100);
INSERT INTO `ch_t_d_areainfo` VALUES (150200, '包头市', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (150202, '东河区', 3, 00000150200);
INSERT INTO `ch_t_d_areainfo` VALUES (150203, '昆都仑区', 3, 00000150200);
INSERT INTO `ch_t_d_areainfo` VALUES (150204, '青山区', 3, 00000150200);
INSERT INTO `ch_t_d_areainfo` VALUES (150205, '石拐区', 3, 00000150200);
INSERT INTO `ch_t_d_areainfo` VALUES (150206, '白云鄂博矿区', 3, 00000150200);
INSERT INTO `ch_t_d_areainfo` VALUES (150207, '九原区', 3, 00000150200);
INSERT INTO `ch_t_d_areainfo` VALUES (150221, '土默特右旗', 3, 00000150200);
INSERT INTO `ch_t_d_areainfo` VALUES (150222, '固阳县', 3, 00000150200);
INSERT INTO `ch_t_d_areainfo` VALUES (150223, '达尔罕茂明安联合旗', 3, 00000150200);
INSERT INTO `ch_t_d_areainfo` VALUES (150300, '乌海市', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (150302, '海勃湾区', 3, 00000150300);
INSERT INTO `ch_t_d_areainfo` VALUES (150303, '海南区', 3, 00000150300);
INSERT INTO `ch_t_d_areainfo` VALUES (150304, '乌达区', 3, 00000150300);
INSERT INTO `ch_t_d_areainfo` VALUES (150400, '赤峰市', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (150402, '红山区', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150403, '元宝山区', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150404, '松山区', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150421, '阿鲁科尔沁旗', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150422, '巴林左旗', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150423, '巴林右旗', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150424, '林西县', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150425, '克什克腾旗', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150426, '翁牛特旗', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150428, '喀喇沁旗', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150429, '宁城县', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150430, '敖汉旗', 3, 00000150400);
INSERT INTO `ch_t_d_areainfo` VALUES (150500, '通辽市', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (150502, '科尔沁区', 3, 00000150500);
INSERT INTO `ch_t_d_areainfo` VALUES (150521, '科尔沁左翼中旗', 3, 00000150500);
INSERT INTO `ch_t_d_areainfo` VALUES (150522, '科尔沁左翼后旗', 3, 00000150500);
INSERT INTO `ch_t_d_areainfo` VALUES (150523, '开鲁县', 3, 00000150500);
INSERT INTO `ch_t_d_areainfo` VALUES (150524, '库伦旗', 3, 00000150500);
INSERT INTO `ch_t_d_areainfo` VALUES (150525, '奈曼旗', 3, 00000150500);
INSERT INTO `ch_t_d_areainfo` VALUES (150526, '扎鲁特旗', 3, 00000150500);
INSERT INTO `ch_t_d_areainfo` VALUES (150581, '霍林郭勒市', 3, 00000150500);
INSERT INTO `ch_t_d_areainfo` VALUES (150600, '鄂尔多斯市', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (150602, '东胜区', 3, 00000150600);
INSERT INTO `ch_t_d_areainfo` VALUES (150603, '康巴什区', 3, 00000150600);
INSERT INTO `ch_t_d_areainfo` VALUES (150621, '达拉特旗', 3, 00000150600);
INSERT INTO `ch_t_d_areainfo` VALUES (150622, '准格尔旗', 3, 00000150600);
INSERT INTO `ch_t_d_areainfo` VALUES (150623, '鄂托克前旗', 3, 00000150600);
INSERT INTO `ch_t_d_areainfo` VALUES (150624, '鄂托克旗', 3, 00000150600);
INSERT INTO `ch_t_d_areainfo` VALUES (150625, '杭锦旗', 3, 00000150600);
INSERT INTO `ch_t_d_areainfo` VALUES (150626, '乌审旗', 3, 00000150600);
INSERT INTO `ch_t_d_areainfo` VALUES (150627, '伊金霍洛旗', 3, 00000150600);
INSERT INTO `ch_t_d_areainfo` VALUES (150700, '呼伦贝尔市', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (150702, '海拉尔区', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150703, '扎赉诺尔区', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150721, '阿荣旗', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150722, '莫力达瓦达斡尔族自治旗', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150723, '鄂伦春自治旗', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150724, '鄂温克族自治旗', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150725, '陈巴尔虎旗', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150726, '新巴尔虎左旗', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150727, '新巴尔虎右旗', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150781, '满洲里市', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150782, '牙克石市', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150783, '扎兰屯市', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150784, '额尔古纳市', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150785, '根河市', 3, 00000150700);
INSERT INTO `ch_t_d_areainfo` VALUES (150800, '巴彦淖尔市', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (150802, '临河区', 3, 00000150800);
INSERT INTO `ch_t_d_areainfo` VALUES (150821, '五原县', 3, 00000150800);
INSERT INTO `ch_t_d_areainfo` VALUES (150822, '磴口县', 3, 00000150800);
INSERT INTO `ch_t_d_areainfo` VALUES (150823, '乌拉特前旗', 3, 00000150800);
INSERT INTO `ch_t_d_areainfo` VALUES (150824, '乌拉特中旗', 3, 00000150800);
INSERT INTO `ch_t_d_areainfo` VALUES (150825, '乌拉特后旗', 3, 00000150800);
INSERT INTO `ch_t_d_areainfo` VALUES (150826, '杭锦后旗', 3, 00000150800);
INSERT INTO `ch_t_d_areainfo` VALUES (150900, '乌兰察布市', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (150902, '集宁区', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150921, '卓资县', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150922, '化德县', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150923, '商都县', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150924, '兴和县', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150925, '凉城县', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150926, '察哈尔右翼前旗', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150927, '察哈尔右翼中旗', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150928, '察哈尔右翼后旗', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150929, '四子王旗', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (150981, '丰镇市', 3, 00000150900);
INSERT INTO `ch_t_d_areainfo` VALUES (152200, '兴安盟', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (152201, '乌兰浩特市', 3, 00000152200);
INSERT INTO `ch_t_d_areainfo` VALUES (152202, '阿尔山市', 3, 00000152200);
INSERT INTO `ch_t_d_areainfo` VALUES (152221, '科尔沁右翼前旗', 3, 00000152200);
INSERT INTO `ch_t_d_areainfo` VALUES (152222, '科尔沁右翼中旗', 3, 00000152200);
INSERT INTO `ch_t_d_areainfo` VALUES (152223, '扎赉特旗', 3, 00000152200);
INSERT INTO `ch_t_d_areainfo` VALUES (152224, '突泉县', 3, 00000152200);
INSERT INTO `ch_t_d_areainfo` VALUES (152500, '锡林郭勒盟', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (152501, '二连浩特市', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152502, '锡林浩特市', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152522, '阿巴嘎旗', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152523, '苏尼特左旗', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152524, '苏尼特右旗', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152525, '东乌珠穆沁旗', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152526, '西乌珠穆沁旗', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152527, '太仆寺旗', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152528, '镶黄旗', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152529, '正镶白旗', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152530, '正蓝旗', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152531, '多伦县', 3, 00000152500);
INSERT INTO `ch_t_d_areainfo` VALUES (152900, '阿拉善盟', 2, 00000150000);
INSERT INTO `ch_t_d_areainfo` VALUES (152921, '阿拉善左旗', 3, 00000152900);
INSERT INTO `ch_t_d_areainfo` VALUES (152922, '阿拉善右旗', 3, 00000152900);
INSERT INTO `ch_t_d_areainfo` VALUES (152923, '额济纳旗', 3, 00000152900);
INSERT INTO `ch_t_d_areainfo` VALUES (210000, '辽宁省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (210100, '沈阳市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (210102, '和平区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210103, '沈河区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210104, '大东区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210105, '皇姑区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210106, '铁西区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210111, '苏家屯区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210112, '浑南区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210113, '沈北新区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210114, '于洪区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210115, '辽中区', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210123, '康平县', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210124, '法库县', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210181, '新民市', 3, 00000210100);
INSERT INTO `ch_t_d_areainfo` VALUES (210200, '大连市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (210202, '中山区', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210203, '西岗区', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210204, '沙河口区', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210211, '甘井子区', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210212, '旅顺口区', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210213, '金州区', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210214, '普兰店区', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210224, '长海县', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210281, '瓦房店市', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210283, '庄河市', 3, 00000210200);
INSERT INTO `ch_t_d_areainfo` VALUES (210300, '鞍山市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (210302, '铁东区', 3, 00000210300);
INSERT INTO `ch_t_d_areainfo` VALUES (210303, '铁西区', 3, 00000210300);
INSERT INTO `ch_t_d_areainfo` VALUES (210304, '立山区', 3, 00000210300);
INSERT INTO `ch_t_d_areainfo` VALUES (210311, '千山区', 3, 00000210300);
INSERT INTO `ch_t_d_areainfo` VALUES (210321, '台安县', 3, 00000210300);
INSERT INTO `ch_t_d_areainfo` VALUES (210323, '岫岩满族自治县', 3, 00000210300);
INSERT INTO `ch_t_d_areainfo` VALUES (210381, '海城市', 3, 00000210300);
INSERT INTO `ch_t_d_areainfo` VALUES (210400, '抚顺市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (210402, '新抚区', 3, 00000210400);
INSERT INTO `ch_t_d_areainfo` VALUES (210403, '东洲区', 3, 00000210400);
INSERT INTO `ch_t_d_areainfo` VALUES (210404, '望花区', 3, 00000210400);
INSERT INTO `ch_t_d_areainfo` VALUES (210411, '顺城区', 3, 00000210400);
INSERT INTO `ch_t_d_areainfo` VALUES (210421, '抚顺县', 3, 00000210400);
INSERT INTO `ch_t_d_areainfo` VALUES (210422, '新宾满族自治县', 3, 00000210400);
INSERT INTO `ch_t_d_areainfo` VALUES (210423, '清原满族自治县', 3, 00000210400);
INSERT INTO `ch_t_d_areainfo` VALUES (210500, '本溪市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (210502, '平山区', 3, 00000210500);
INSERT INTO `ch_t_d_areainfo` VALUES (210503, '溪湖区', 3, 00000210500);
INSERT INTO `ch_t_d_areainfo` VALUES (210504, '明山区', 3, 00000210500);
INSERT INTO `ch_t_d_areainfo` VALUES (210505, '南芬区', 3, 00000210500);
INSERT INTO `ch_t_d_areainfo` VALUES (210521, '本溪满族自治县', 3, 00000210500);
INSERT INTO `ch_t_d_areainfo` VALUES (210522, '桓仁满族自治县', 3, 00000210500);
INSERT INTO `ch_t_d_areainfo` VALUES (210600, '丹东市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (210602, '元宝区', 3, 00000210600);
INSERT INTO `ch_t_d_areainfo` VALUES (210603, '振兴区', 3, 00000210600);
INSERT INTO `ch_t_d_areainfo` VALUES (210604, '振安区', 3, 00000210600);
INSERT INTO `ch_t_d_areainfo` VALUES (210624, '宽甸满族自治县', 3, 00000210600);
INSERT INTO `ch_t_d_areainfo` VALUES (210681, '东港市', 3, 00000210600);
INSERT INTO `ch_t_d_areainfo` VALUES (210682, '凤城市', 3, 00000210600);
INSERT INTO `ch_t_d_areainfo` VALUES (210700, '锦州市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (210702, '古塔区', 3, 00000210700);
INSERT INTO `ch_t_d_areainfo` VALUES (210703, '凌河区', 3, 00000210700);
INSERT INTO `ch_t_d_areainfo` VALUES (210711, '太和区', 3, 00000210700);
INSERT INTO `ch_t_d_areainfo` VALUES (210726, '黑山县', 3, 00000210700);
INSERT INTO `ch_t_d_areainfo` VALUES (210727, '义县', 3, 00000210700);
INSERT INTO `ch_t_d_areainfo` VALUES (210781, '凌海市', 3, 00000210700);
INSERT INTO `ch_t_d_areainfo` VALUES (210782, '北镇市', 3, 00000210700);
INSERT INTO `ch_t_d_areainfo` VALUES (210800, '营口市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (210802, '站前区', 3, 00000210800);
INSERT INTO `ch_t_d_areainfo` VALUES (210803, '西市区', 3, 00000210800);
INSERT INTO `ch_t_d_areainfo` VALUES (210804, '鲅鱼圈区', 3, 00000210800);
INSERT INTO `ch_t_d_areainfo` VALUES (210811, '老边区', 3, 00000210800);
INSERT INTO `ch_t_d_areainfo` VALUES (210881, '盖州市', 3, 00000210800);
INSERT INTO `ch_t_d_areainfo` VALUES (210882, '大石桥市', 3, 00000210800);
INSERT INTO `ch_t_d_areainfo` VALUES (210900, '阜新市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (210902, '海州区', 3, 00000210900);
INSERT INTO `ch_t_d_areainfo` VALUES (210903, '新邱区', 3, 00000210900);
INSERT INTO `ch_t_d_areainfo` VALUES (210904, '太平区', 3, 00000210900);
INSERT INTO `ch_t_d_areainfo` VALUES (210905, '清河门区', 3, 00000210900);
INSERT INTO `ch_t_d_areainfo` VALUES (210911, '细河区', 3, 00000210900);
INSERT INTO `ch_t_d_areainfo` VALUES (210921, '阜新蒙古族自治县', 3, 00000210900);
INSERT INTO `ch_t_d_areainfo` VALUES (210922, '彰武县', 3, 00000210900);
INSERT INTO `ch_t_d_areainfo` VALUES (211000, '辽阳市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (211002, '白塔区', 3, 00000211000);
INSERT INTO `ch_t_d_areainfo` VALUES (211003, '文圣区', 3, 00000211000);
INSERT INTO `ch_t_d_areainfo` VALUES (211004, '宏伟区', 3, 00000211000);
INSERT INTO `ch_t_d_areainfo` VALUES (211005, '弓长岭区', 3, 00000211000);
INSERT INTO `ch_t_d_areainfo` VALUES (211011, '太子河区', 3, 00000211000);
INSERT INTO `ch_t_d_areainfo` VALUES (211021, '辽阳县', 3, 00000211000);
INSERT INTO `ch_t_d_areainfo` VALUES (211081, '灯塔市', 3, 00000211000);
INSERT INTO `ch_t_d_areainfo` VALUES (211100, '盘锦市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (211102, '双台子区', 3, 00000211100);
INSERT INTO `ch_t_d_areainfo` VALUES (211103, '兴隆台区', 3, 00000211100);
INSERT INTO `ch_t_d_areainfo` VALUES (211104, '大洼区', 3, 00000211100);
INSERT INTO `ch_t_d_areainfo` VALUES (211122, '盘山县', 3, 00000211100);
INSERT INTO `ch_t_d_areainfo` VALUES (211200, '铁岭市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (211202, '银州区', 3, 00000211200);
INSERT INTO `ch_t_d_areainfo` VALUES (211204, '清河区', 3, 00000211200);
INSERT INTO `ch_t_d_areainfo` VALUES (211221, '铁岭县', 3, 00000211200);
INSERT INTO `ch_t_d_areainfo` VALUES (211223, '西丰县', 3, 00000211200);
INSERT INTO `ch_t_d_areainfo` VALUES (211224, '昌图县', 3, 00000211200);
INSERT INTO `ch_t_d_areainfo` VALUES (211281, '调兵山市', 3, 00000211200);
INSERT INTO `ch_t_d_areainfo` VALUES (211282, '开原市', 3, 00000211200);
INSERT INTO `ch_t_d_areainfo` VALUES (211300, '朝阳市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (211302, '双塔区', 3, 00000211300);
INSERT INTO `ch_t_d_areainfo` VALUES (211303, '龙城区', 3, 00000211300);
INSERT INTO `ch_t_d_areainfo` VALUES (211321, '朝阳县', 3, 00000211300);
INSERT INTO `ch_t_d_areainfo` VALUES (211322, '建平县', 3, 00000211300);
INSERT INTO `ch_t_d_areainfo` VALUES (211324, '喀喇沁左翼蒙古族自治县', 3, 00000211300);
INSERT INTO `ch_t_d_areainfo` VALUES (211381, '北票市', 3, 00000211300);
INSERT INTO `ch_t_d_areainfo` VALUES (211382, '凌源市', 3, 00000211300);
INSERT INTO `ch_t_d_areainfo` VALUES (211400, '葫芦岛市', 2, 00000210000);
INSERT INTO `ch_t_d_areainfo` VALUES (211402, '连山区', 3, 00000211400);
INSERT INTO `ch_t_d_areainfo` VALUES (211403, '龙港区', 3, 00000211400);
INSERT INTO `ch_t_d_areainfo` VALUES (211404, '南票区', 3, 00000211400);
INSERT INTO `ch_t_d_areainfo` VALUES (211421, '绥中县', 3, 00000211400);
INSERT INTO `ch_t_d_areainfo` VALUES (211422, '建昌县', 3, 00000211400);
INSERT INTO `ch_t_d_areainfo` VALUES (211481, '兴城市', 3, 00000211400);
INSERT INTO `ch_t_d_areainfo` VALUES (220000, '吉林省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (220100, '长春市', 2, 00000220000);
INSERT INTO `ch_t_d_areainfo` VALUES (220102, '南关区', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220103, '宽城区', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220104, '朝阳区', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220105, '二道区', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220106, '绿园区', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220112, '双阳区', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220113, '九台区', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220122, '农安县', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220182, '榆树市', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220183, '德惠市', 3, 00000220100);
INSERT INTO `ch_t_d_areainfo` VALUES (220200, '吉林市', 2, 00000220000);
INSERT INTO `ch_t_d_areainfo` VALUES (220202, '昌邑区', 3, 00000220200);
INSERT INTO `ch_t_d_areainfo` VALUES (220203, '龙潭区', 3, 00000220200);
INSERT INTO `ch_t_d_areainfo` VALUES (220204, '船营区', 3, 00000220200);
INSERT INTO `ch_t_d_areainfo` VALUES (220211, '丰满区', 3, 00000220200);
INSERT INTO `ch_t_d_areainfo` VALUES (220221, '永吉县', 3, 00000220200);
INSERT INTO `ch_t_d_areainfo` VALUES (220281, '蛟河市', 3, 00000220200);
INSERT INTO `ch_t_d_areainfo` VALUES (220282, '桦甸市', 3, 00000220200);
INSERT INTO `ch_t_d_areainfo` VALUES (220283, '舒兰市', 3, 00000220200);
INSERT INTO `ch_t_d_areainfo` VALUES (220284, '磐石市', 3, 00000220200);
INSERT INTO `ch_t_d_areainfo` VALUES (220300, '四平市', 2, 00000220000);
INSERT INTO `ch_t_d_areainfo` VALUES (220302, '铁西区', 3, 00000220300);
INSERT INTO `ch_t_d_areainfo` VALUES (220303, '铁东区', 3, 00000220300);
INSERT INTO `ch_t_d_areainfo` VALUES (220322, '梨树县', 3, 00000220300);
INSERT INTO `ch_t_d_areainfo` VALUES (220323, '伊通满族自治县', 3, 00000220300);
INSERT INTO `ch_t_d_areainfo` VALUES (220381, '公主岭市', 3, 00000220300);
INSERT INTO `ch_t_d_areainfo` VALUES (220382, '双辽市', 3, 00000220300);
INSERT INTO `ch_t_d_areainfo` VALUES (220400, '辽源市', 2, 00000220000);
INSERT INTO `ch_t_d_areainfo` VALUES (220402, '龙山区', 3, 00000220400);
INSERT INTO `ch_t_d_areainfo` VALUES (220403, '西安区', 3, 00000220400);
INSERT INTO `ch_t_d_areainfo` VALUES (220421, '东丰县', 3, 00000220400);
INSERT INTO `ch_t_d_areainfo` VALUES (220422, '东辽县', 3, 00000220400);
INSERT INTO `ch_t_d_areainfo` VALUES (220500, '通化市', 2, 00000220000);
INSERT INTO `ch_t_d_areainfo` VALUES (220502, '东昌区', 3, 00000220500);
INSERT INTO `ch_t_d_areainfo` VALUES (220503, '二道江区', 3, 00000220500);
INSERT INTO `ch_t_d_areainfo` VALUES (220521, '通化县', 3, 00000220500);
INSERT INTO `ch_t_d_areainfo` VALUES (220523, '辉南县', 3, 00000220500);
INSERT INTO `ch_t_d_areainfo` VALUES (220524, '柳河县', 3, 00000220500);
INSERT INTO `ch_t_d_areainfo` VALUES (220581, '梅河口市', 3, 00000220500);
INSERT INTO `ch_t_d_areainfo` VALUES (220582, '集安市', 3, 00000220500);
INSERT INTO `ch_t_d_areainfo` VALUES (220600, '白山市', 2, 00000220000);
INSERT INTO `ch_t_d_areainfo` VALUES (220602, '浑江区', 3, 00000220600);
INSERT INTO `ch_t_d_areainfo` VALUES (220605, '江源区', 3, 00000220600);
INSERT INTO `ch_t_d_areainfo` VALUES (220621, '抚松县', 3, 00000220600);
INSERT INTO `ch_t_d_areainfo` VALUES (220622, '靖宇县', 3, 00000220600);
INSERT INTO `ch_t_d_areainfo` VALUES (220623, '长白朝鲜族自治县', 3, 00000220600);
INSERT INTO `ch_t_d_areainfo` VALUES (220681, '临江市', 3, 00000220600);
INSERT INTO `ch_t_d_areainfo` VALUES (220700, '松原市', 2, 00000220000);
INSERT INTO `ch_t_d_areainfo` VALUES (220702, '宁江区', 3, 00000220700);
INSERT INTO `ch_t_d_areainfo` VALUES (220721, '前郭尔罗斯蒙古族自治县', 3, 00000220700);
INSERT INTO `ch_t_d_areainfo` VALUES (220722, '长岭县', 3, 00000220700);
INSERT INTO `ch_t_d_areainfo` VALUES (220723, '乾安县', 3, 00000220700);
INSERT INTO `ch_t_d_areainfo` VALUES (220781, '扶余市', 3, 00000220700);
INSERT INTO `ch_t_d_areainfo` VALUES (220800, '白城市', 2, 00000220000);
INSERT INTO `ch_t_d_areainfo` VALUES (220802, '洮北区', 3, 00000220800);
INSERT INTO `ch_t_d_areainfo` VALUES (220821, '镇赉县', 3, 00000220800);
INSERT INTO `ch_t_d_areainfo` VALUES (220822, '通榆县', 3, 00000220800);
INSERT INTO `ch_t_d_areainfo` VALUES (220881, '洮南市', 3, 00000220800);
INSERT INTO `ch_t_d_areainfo` VALUES (220882, '大安市', 3, 00000220800);
INSERT INTO `ch_t_d_areainfo` VALUES (222400, '延边朝鲜族自治州', 2, 00000220000);
INSERT INTO `ch_t_d_areainfo` VALUES (222401, '延吉市', 3, 00000222400);
INSERT INTO `ch_t_d_areainfo` VALUES (222402, '图们市', 3, 00000222400);
INSERT INTO `ch_t_d_areainfo` VALUES (222403, '敦化市', 3, 00000222400);
INSERT INTO `ch_t_d_areainfo` VALUES (222404, '珲春市', 3, 00000222400);
INSERT INTO `ch_t_d_areainfo` VALUES (222405, '龙井市', 3, 00000222400);
INSERT INTO `ch_t_d_areainfo` VALUES (222406, '和龙市', 3, 00000222400);
INSERT INTO `ch_t_d_areainfo` VALUES (222424, '汪清县', 3, 00000222400);
INSERT INTO `ch_t_d_areainfo` VALUES (222426, '安图县', 3, 00000222400);
INSERT INTO `ch_t_d_areainfo` VALUES (230000, '黑龙江省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (230100, '哈尔滨市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (230102, '道里区', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230103, '南岗区', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230104, '道外区', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230108, '平房区', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230109, '松北区', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230110, '香坊区', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230111, '呼兰区', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230112, '阿城区', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230113, '双城区', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230123, '依兰县', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230124, '方正县', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230125, '宾县', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230126, '巴彦县', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230127, '木兰县', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230128, '通河县', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230129, '延寿县', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230183, '尚志市', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230184, '五常市', 3, 00000230100);
INSERT INTO `ch_t_d_areainfo` VALUES (230200, '齐齐哈尔市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (230202, '龙沙区', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230203, '建华区', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230204, '铁锋区', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230205, '昂昂溪区', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230206, '富拉尔基区', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230207, '碾子山区', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230208, '梅里斯达斡尔族区', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230221, '龙江县', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230223, '依安县', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230224, '泰来县', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230225, '甘南县', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230227, '富裕县', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230229, '克山县', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230230, '克东县', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230231, '拜泉县', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230281, '讷河市', 3, 00000230200);
INSERT INTO `ch_t_d_areainfo` VALUES (230300, '鸡西市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (230302, '鸡冠区', 3, 00000230300);
INSERT INTO `ch_t_d_areainfo` VALUES (230303, '恒山区', 3, 00000230300);
INSERT INTO `ch_t_d_areainfo` VALUES (230304, '滴道区', 3, 00000230300);
INSERT INTO `ch_t_d_areainfo` VALUES (230305, '梨树区', 3, 00000230300);
INSERT INTO `ch_t_d_areainfo` VALUES (230306, '城子河区', 3, 00000230300);
INSERT INTO `ch_t_d_areainfo` VALUES (230307, '麻山区', 3, 00000230300);
INSERT INTO `ch_t_d_areainfo` VALUES (230321, '鸡东县', 3, 00000230300);
INSERT INTO `ch_t_d_areainfo` VALUES (230381, '虎林市', 3, 00000230300);
INSERT INTO `ch_t_d_areainfo` VALUES (230382, '密山市', 3, 00000230300);
INSERT INTO `ch_t_d_areainfo` VALUES (230400, '鹤岗市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (230402, '向阳区', 3, 00000230400);
INSERT INTO `ch_t_d_areainfo` VALUES (230403, '工农区', 3, 00000230400);
INSERT INTO `ch_t_d_areainfo` VALUES (230404, '南山区', 3, 00000230400);
INSERT INTO `ch_t_d_areainfo` VALUES (230405, '兴安区', 3, 00000230400);
INSERT INTO `ch_t_d_areainfo` VALUES (230406, '东山区', 3, 00000230400);
INSERT INTO `ch_t_d_areainfo` VALUES (230407, '兴山区', 3, 00000230400);
INSERT INTO `ch_t_d_areainfo` VALUES (230421, '萝北县', 3, 00000230400);
INSERT INTO `ch_t_d_areainfo` VALUES (230422, '绥滨县', 3, 00000230400);
INSERT INTO `ch_t_d_areainfo` VALUES (230500, '双鸭山市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (230502, '尖山区', 3, 00000230500);
INSERT INTO `ch_t_d_areainfo` VALUES (230503, '岭东区', 3, 00000230500);
INSERT INTO `ch_t_d_areainfo` VALUES (230505, '四方台区', 3, 00000230500);
INSERT INTO `ch_t_d_areainfo` VALUES (230506, '宝山区', 3, 00000230500);
INSERT INTO `ch_t_d_areainfo` VALUES (230521, '集贤县', 3, 00000230500);
INSERT INTO `ch_t_d_areainfo` VALUES (230522, '友谊县', 3, 00000230500);
INSERT INTO `ch_t_d_areainfo` VALUES (230523, '宝清县', 3, 00000230500);
INSERT INTO `ch_t_d_areainfo` VALUES (230524, '饶河县', 3, 00000230500);
INSERT INTO `ch_t_d_areainfo` VALUES (230600, '大庆市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (230602, '萨尔图区', 3, 00000230600);
INSERT INTO `ch_t_d_areainfo` VALUES (230603, '龙凤区', 3, 00000230600);
INSERT INTO `ch_t_d_areainfo` VALUES (230604, '让胡路区', 3, 00000230600);
INSERT INTO `ch_t_d_areainfo` VALUES (230605, '红岗区', 3, 00000230600);
INSERT INTO `ch_t_d_areainfo` VALUES (230606, '大同区', 3, 00000230600);
INSERT INTO `ch_t_d_areainfo` VALUES (230621, '肇州县', 3, 00000230600);
INSERT INTO `ch_t_d_areainfo` VALUES (230622, '肇源县', 3, 00000230600);
INSERT INTO `ch_t_d_areainfo` VALUES (230623, '林甸县', 3, 00000230600);
INSERT INTO `ch_t_d_areainfo` VALUES (230624, '杜尔伯特蒙古族自治县', 3, 00000230600);
INSERT INTO `ch_t_d_areainfo` VALUES (230700, '伊春市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (230702, '伊春区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230703, '南岔区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230704, '友好区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230705, '西林区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230706, '翠峦区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230707, '新青区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230708, '美溪区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230709, '金山屯区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230710, '五营区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230711, '乌马河区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230712, '汤旺河区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230713, '带岭区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230714, '乌伊岭区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230715, '红星区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230716, '上甘岭区', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230722, '嘉荫县', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230781, '铁力市', 3, 00000230700);
INSERT INTO `ch_t_d_areainfo` VALUES (230800, '佳木斯市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (230803, '向阳区', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230804, '前进区', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230805, '东风区', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230811, '郊区', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230822, '桦南县', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230826, '桦川县', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230828, '汤原县', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230881, '同江市', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230882, '富锦市', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230883, '抚远市', 3, 00000230800);
INSERT INTO `ch_t_d_areainfo` VALUES (230900, '七台河市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (230902, '新兴区', 3, 00000230900);
INSERT INTO `ch_t_d_areainfo` VALUES (230903, '桃山区', 3, 00000230900);
INSERT INTO `ch_t_d_areainfo` VALUES (230904, '茄子河区', 3, 00000230900);
INSERT INTO `ch_t_d_areainfo` VALUES (230921, '勃利县', 3, 00000230900);
INSERT INTO `ch_t_d_areainfo` VALUES (231000, '牡丹江市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (231002, '东安区', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231003, '阳明区', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231004, '爱民区', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231005, '西安区', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231025, '林口县', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231081, '绥芬河市', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231083, '海林市', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231084, '宁安市', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231085, '穆棱市', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231086, '东宁市', 3, 00000231000);
INSERT INTO `ch_t_d_areainfo` VALUES (231100, '黑河市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (231102, '爱辉区', 3, 00000231100);
INSERT INTO `ch_t_d_areainfo` VALUES (231121, '嫩江县', 3, 00000231100);
INSERT INTO `ch_t_d_areainfo` VALUES (231123, '逊克县', 3, 00000231100);
INSERT INTO `ch_t_d_areainfo` VALUES (231124, '孙吴县', 3, 00000231100);
INSERT INTO `ch_t_d_areainfo` VALUES (231181, '北安市', 3, 00000231100);
INSERT INTO `ch_t_d_areainfo` VALUES (231182, '五大连池市', 3, 00000231100);
INSERT INTO `ch_t_d_areainfo` VALUES (231200, '绥化市', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (231202, '北林区', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (231221, '望奎县', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (231222, '兰西县', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (231223, '青冈县', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (231224, '庆安县', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (231225, '明水县', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (231226, '绥棱县', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (231281, '安达市', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (231282, '肇东市', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (231283, '海伦市', 3, 00000231200);
INSERT INTO `ch_t_d_areainfo` VALUES (232700, '大兴安岭地区', 2, 00000230000);
INSERT INTO `ch_t_d_areainfo` VALUES (232721, '呼玛县', 3, 00000232700);
INSERT INTO `ch_t_d_areainfo` VALUES (232722, '塔河县', 3, 00000232700);
INSERT INTO `ch_t_d_areainfo` VALUES (232723, '漠河县', 3, 00000232700);
INSERT INTO `ch_t_d_areainfo` VALUES (310000, '上海市', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (310101, '黄浦区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310104, '徐汇区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310105, '长宁区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310106, '静安区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310107, '普陀区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310109, '虹口区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310110, '杨浦区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310112, '闵行区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310113, '宝山区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310114, '嘉定区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310115, '浦东新区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310116, '金山区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310117, '松江区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310118, '青浦区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310120, '奉贤区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (310151, '崇明区', 3, 00000310100);
INSERT INTO `ch_t_d_areainfo` VALUES (320000, '江苏省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (320100, '南京市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (320102, '玄武区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320104, '秦淮区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320105, '建邺区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320106, '鼓楼区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320111, '浦口区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320113, '栖霞区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320114, '雨花台区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320115, '江宁区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320116, '六合区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320117, '溧水区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320118, '高淳区', 3, 00000320100);
INSERT INTO `ch_t_d_areainfo` VALUES (320200, '无锡市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (320205, '锡山区', 3, 00000320200);
INSERT INTO `ch_t_d_areainfo` VALUES (320206, '惠山区', 3, 00000320200);
INSERT INTO `ch_t_d_areainfo` VALUES (320211, '滨湖区', 3, 00000320200);
INSERT INTO `ch_t_d_areainfo` VALUES (320213, '梁溪区', 3, 00000320200);
INSERT INTO `ch_t_d_areainfo` VALUES (320214, '新吴区', 3, 00000320200);
INSERT INTO `ch_t_d_areainfo` VALUES (320281, '江阴市', 3, 00000320200);
INSERT INTO `ch_t_d_areainfo` VALUES (320282, '宜兴市', 3, 00000320200);
INSERT INTO `ch_t_d_areainfo` VALUES (320300, '徐州市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (320302, '鼓楼区', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320303, '云龙区', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320305, '贾汪区', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320311, '泉山区', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320312, '铜山区', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320321, '丰县', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320322, '沛县', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320324, '睢宁县', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320381, '新沂市', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320382, '邳州市', 3, 00000320300);
INSERT INTO `ch_t_d_areainfo` VALUES (320400, '常州市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (320402, '天宁区', 3, 00000320400);
INSERT INTO `ch_t_d_areainfo` VALUES (320404, '钟楼区', 3, 00000320400);
INSERT INTO `ch_t_d_areainfo` VALUES (320411, '新北区', 3, 00000320400);
INSERT INTO `ch_t_d_areainfo` VALUES (320412, '武进区', 3, 00000320400);
INSERT INTO `ch_t_d_areainfo` VALUES (320413, '金坛区', 3, 00000320400);
INSERT INTO `ch_t_d_areainfo` VALUES (320481, '溧阳市', 3, 00000320400);
INSERT INTO `ch_t_d_areainfo` VALUES (320500, '苏州市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (320505, '虎丘区', 3, 00000320500);
INSERT INTO `ch_t_d_areainfo` VALUES (320506, '吴中区', 3, 00000320500);
INSERT INTO `ch_t_d_areainfo` VALUES (320507, '相城区', 3, 00000320500);
INSERT INTO `ch_t_d_areainfo` VALUES (320508, '姑苏区', 3, 00000320500);
INSERT INTO `ch_t_d_areainfo` VALUES (320509, '吴江区', 3, 00000320500);
INSERT INTO `ch_t_d_areainfo` VALUES (320581, '常熟市', 3, 00000320500);
INSERT INTO `ch_t_d_areainfo` VALUES (320582, '张家港市', 3, 00000320500);
INSERT INTO `ch_t_d_areainfo` VALUES (320583, '昆山市', 3, 00000320500);
INSERT INTO `ch_t_d_areainfo` VALUES (320585, '太仓市', 3, 00000320500);
INSERT INTO `ch_t_d_areainfo` VALUES (320600, '南通市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (320602, '崇川区', 3, 00000320600);
INSERT INTO `ch_t_d_areainfo` VALUES (320611, '港闸区', 3, 00000320600);
INSERT INTO `ch_t_d_areainfo` VALUES (320612, '通州区', 3, 00000320600);
INSERT INTO `ch_t_d_areainfo` VALUES (320621, '海安县', 3, 00000320600);
INSERT INTO `ch_t_d_areainfo` VALUES (320623, '如东县', 3, 00000320600);
INSERT INTO `ch_t_d_areainfo` VALUES (320681, '启东市', 3, 00000320600);
INSERT INTO `ch_t_d_areainfo` VALUES (320682, '如皋市', 3, 00000320600);
INSERT INTO `ch_t_d_areainfo` VALUES (320684, '海门市', 3, 00000320600);
INSERT INTO `ch_t_d_areainfo` VALUES (320700, '连云港市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (320703, '连云区', 3, 00000320700);
INSERT INTO `ch_t_d_areainfo` VALUES (320706, '海州区', 3, 00000320700);
INSERT INTO `ch_t_d_areainfo` VALUES (320707, '赣榆区', 3, 00000320700);
INSERT INTO `ch_t_d_areainfo` VALUES (320722, '东海县', 3, 00000320700);
INSERT INTO `ch_t_d_areainfo` VALUES (320723, '灌云县', 3, 00000320700);
INSERT INTO `ch_t_d_areainfo` VALUES (320724, '灌南县', 3, 00000320700);
INSERT INTO `ch_t_d_areainfo` VALUES (320800, '淮安市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (320803, '淮安区', 3, 00000320800);
INSERT INTO `ch_t_d_areainfo` VALUES (320804, '淮阴区', 3, 00000320800);
INSERT INTO `ch_t_d_areainfo` VALUES (320812, '清江浦区', 3, 00000320800);
INSERT INTO `ch_t_d_areainfo` VALUES (320813, '洪泽区', 3, 00000320800);
INSERT INTO `ch_t_d_areainfo` VALUES (320826, '涟水县', 3, 00000320800);
INSERT INTO `ch_t_d_areainfo` VALUES (320830, '盱眙县', 3, 00000320800);
INSERT INTO `ch_t_d_areainfo` VALUES (320831, '金湖县', 3, 00000320800);
INSERT INTO `ch_t_d_areainfo` VALUES (320900, '盐城市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (320902, '亭湖区', 3, 00000320900);
INSERT INTO `ch_t_d_areainfo` VALUES (320903, '盐都区', 3, 00000320900);
INSERT INTO `ch_t_d_areainfo` VALUES (320904, '大丰区', 3, 00000320900);
INSERT INTO `ch_t_d_areainfo` VALUES (320921, '响水县', 3, 00000320900);
INSERT INTO `ch_t_d_areainfo` VALUES (320922, '滨海县', 3, 00000320900);
INSERT INTO `ch_t_d_areainfo` VALUES (320923, '阜宁县', 3, 00000320900);
INSERT INTO `ch_t_d_areainfo` VALUES (320924, '射阳县', 3, 00000320900);
INSERT INTO `ch_t_d_areainfo` VALUES (320925, '建湖县', 3, 00000320900);
INSERT INTO `ch_t_d_areainfo` VALUES (320981, '东台市', 3, 00000320900);
INSERT INTO `ch_t_d_areainfo` VALUES (321000, '扬州市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (321002, '广陵区', 3, 00000321000);
INSERT INTO `ch_t_d_areainfo` VALUES (321003, '邗江区', 3, 00000321000);
INSERT INTO `ch_t_d_areainfo` VALUES (321012, '江都区', 3, 00000321000);
INSERT INTO `ch_t_d_areainfo` VALUES (321023, '宝应县', 3, 00000321000);
INSERT INTO `ch_t_d_areainfo` VALUES (321081, '仪征市', 3, 00000321000);
INSERT INTO `ch_t_d_areainfo` VALUES (321084, '高邮市', 3, 00000321000);
INSERT INTO `ch_t_d_areainfo` VALUES (321100, '镇江市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (321102, '京口区', 3, 00000321100);
INSERT INTO `ch_t_d_areainfo` VALUES (321111, '润州区', 3, 00000321100);
INSERT INTO `ch_t_d_areainfo` VALUES (321112, '丹徒区', 3, 00000321100);
INSERT INTO `ch_t_d_areainfo` VALUES (321181, '丹阳市', 3, 00000321100);
INSERT INTO `ch_t_d_areainfo` VALUES (321182, '扬中市', 3, 00000321100);
INSERT INTO `ch_t_d_areainfo` VALUES (321183, '句容市', 3, 00000321100);
INSERT INTO `ch_t_d_areainfo` VALUES (321200, '泰州市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (321202, '海陵区', 3, 00000321200);
INSERT INTO `ch_t_d_areainfo` VALUES (321203, '高港区', 3, 00000321200);
INSERT INTO `ch_t_d_areainfo` VALUES (321204, '姜堰区', 3, 00000321200);
INSERT INTO `ch_t_d_areainfo` VALUES (321281, '兴化市', 3, 00000321200);
INSERT INTO `ch_t_d_areainfo` VALUES (321282, '靖江市', 3, 00000321200);
INSERT INTO `ch_t_d_areainfo` VALUES (321283, '泰兴市', 3, 00000321200);
INSERT INTO `ch_t_d_areainfo` VALUES (321300, '宿迁市', 2, 00000320000);
INSERT INTO `ch_t_d_areainfo` VALUES (321302, '宿城区', 3, 00000321300);
INSERT INTO `ch_t_d_areainfo` VALUES (321311, '宿豫区', 3, 00000321300);
INSERT INTO `ch_t_d_areainfo` VALUES (321322, '沭阳县', 3, 00000321300);
INSERT INTO `ch_t_d_areainfo` VALUES (321323, '泗阳县', 3, 00000321300);
INSERT INTO `ch_t_d_areainfo` VALUES (321324, '泗洪县', 3, 00000321300);
INSERT INTO `ch_t_d_areainfo` VALUES (330000, '浙江省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (330100, '杭州市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (330102, '上城区', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330103, '下城区', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330104, '江干区', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330105, '拱墅区', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330106, '西湖区', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330108, '滨江区', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330109, '萧山区', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330110, '余杭区', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330111, '富阳区', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330122, '桐庐县', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330127, '淳安县', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330182, '建德市', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330185, '临安市', 3, 00000330100);
INSERT INTO `ch_t_d_areainfo` VALUES (330200, '宁波市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (330203, '海曙区', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330204, '江东区', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330205, '江北区', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330206, '北仑区', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330211, '镇海区', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330212, '鄞州区', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330225, '象山县', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330226, '宁海县', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330281, '余姚市', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330282, '慈溪市', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330283, '奉化市', 3, 00000330200);
INSERT INTO `ch_t_d_areainfo` VALUES (330300, '温州市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (330302, '鹿城区', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330303, '龙湾区', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330304, '瓯海区', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330305, '洞头区', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330324, '永嘉县', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330326, '平阳县', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330327, '苍南县', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330328, '文成县', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330329, '泰顺县', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330381, '瑞安市', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330382, '乐清市', 3, 00000330300);
INSERT INTO `ch_t_d_areainfo` VALUES (330400, '嘉兴市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (330402, '南湖区', 3, 00000330400);
INSERT INTO `ch_t_d_areainfo` VALUES (330411, '秀洲区', 3, 00000330400);
INSERT INTO `ch_t_d_areainfo` VALUES (330421, '嘉善县', 3, 00000330400);
INSERT INTO `ch_t_d_areainfo` VALUES (330424, '海盐县', 3, 00000330400);
INSERT INTO `ch_t_d_areainfo` VALUES (330481, '海宁市', 3, 00000330400);
INSERT INTO `ch_t_d_areainfo` VALUES (330482, '平湖市', 3, 00000330400);
INSERT INTO `ch_t_d_areainfo` VALUES (330483, '桐乡市', 3, 00000330400);
INSERT INTO `ch_t_d_areainfo` VALUES (330500, '湖州市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (330502, '吴兴区', 3, 00000330500);
INSERT INTO `ch_t_d_areainfo` VALUES (330503, '南浔区', 3, 00000330500);
INSERT INTO `ch_t_d_areainfo` VALUES (330521, '德清县', 3, 00000330500);
INSERT INTO `ch_t_d_areainfo` VALUES (330522, '长兴县', 3, 00000330500);
INSERT INTO `ch_t_d_areainfo` VALUES (330523, '安吉县', 3, 00000330500);
INSERT INTO `ch_t_d_areainfo` VALUES (330600, '绍兴市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (330602, '越城区', 3, 00000330600);
INSERT INTO `ch_t_d_areainfo` VALUES (330603, '柯桥区', 3, 00000330600);
INSERT INTO `ch_t_d_areainfo` VALUES (330604, '上虞区', 3, 00000330600);
INSERT INTO `ch_t_d_areainfo` VALUES (330624, '新昌县', 3, 00000330600);
INSERT INTO `ch_t_d_areainfo` VALUES (330681, '诸暨市', 3, 00000330600);
INSERT INTO `ch_t_d_areainfo` VALUES (330683, '嵊州市', 3, 00000330600);
INSERT INTO `ch_t_d_areainfo` VALUES (330700, '金华市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (330702, '婺城区', 3, 00000330700);
INSERT INTO `ch_t_d_areainfo` VALUES (330703, '金东区', 3, 00000330700);
INSERT INTO `ch_t_d_areainfo` VALUES (330723, '武义县', 3, 00000330700);
INSERT INTO `ch_t_d_areainfo` VALUES (330726, '浦江县', 3, 00000330700);
INSERT INTO `ch_t_d_areainfo` VALUES (330727, '磐安县', 3, 00000330700);
INSERT INTO `ch_t_d_areainfo` VALUES (330781, '兰溪市', 3, 00000330700);
INSERT INTO `ch_t_d_areainfo` VALUES (330782, '义乌市', 3, 00000330700);
INSERT INTO `ch_t_d_areainfo` VALUES (330783, '东阳市', 3, 00000330700);
INSERT INTO `ch_t_d_areainfo` VALUES (330784, '永康市', 3, 00000330700);
INSERT INTO `ch_t_d_areainfo` VALUES (330800, '衢州市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (330802, '柯城区', 3, 00000330800);
INSERT INTO `ch_t_d_areainfo` VALUES (330803, '衢江区', 3, 00000330800);
INSERT INTO `ch_t_d_areainfo` VALUES (330822, '常山县', 3, 00000330800);
INSERT INTO `ch_t_d_areainfo` VALUES (330824, '开化县', 3, 00000330800);
INSERT INTO `ch_t_d_areainfo` VALUES (330825, '龙游县', 3, 00000330800);
INSERT INTO `ch_t_d_areainfo` VALUES (330881, '江山市', 3, 00000330800);
INSERT INTO `ch_t_d_areainfo` VALUES (330900, '舟山市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (330902, '定海区', 3, 00000330900);
INSERT INTO `ch_t_d_areainfo` VALUES (330903, '普陀区', 3, 00000330900);
INSERT INTO `ch_t_d_areainfo` VALUES (330921, '岱山县', 3, 00000330900);
INSERT INTO `ch_t_d_areainfo` VALUES (330922, '嵊泗县', 3, 00000330900);
INSERT INTO `ch_t_d_areainfo` VALUES (331000, '台州市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (331002, '椒江区', 3, 00000331000);
INSERT INTO `ch_t_d_areainfo` VALUES (331003, '黄岩区', 3, 00000331000);
INSERT INTO `ch_t_d_areainfo` VALUES (331004, '路桥区', 3, 00000331000);
INSERT INTO `ch_t_d_areainfo` VALUES (331021, '玉环县', 3, 00000331000);
INSERT INTO `ch_t_d_areainfo` VALUES (331022, '三门县', 3, 00000331000);
INSERT INTO `ch_t_d_areainfo` VALUES (331023, '天台县', 3, 00000331000);
INSERT INTO `ch_t_d_areainfo` VALUES (331024, '仙居县', 3, 00000331000);
INSERT INTO `ch_t_d_areainfo` VALUES (331081, '温岭市', 3, 00000331000);
INSERT INTO `ch_t_d_areainfo` VALUES (331082, '临海市', 3, 00000331000);
INSERT INTO `ch_t_d_areainfo` VALUES (331100, '丽水市', 2, 00000330000);
INSERT INTO `ch_t_d_areainfo` VALUES (331102, '莲都区', 3, 00000331100);
INSERT INTO `ch_t_d_areainfo` VALUES (331121, '青田县', 3, 00000331100);
INSERT INTO `ch_t_d_areainfo` VALUES (331122, '缙云县', 3, 00000331100);
INSERT INTO `ch_t_d_areainfo` VALUES (331123, '遂昌县', 3, 00000331100);
INSERT INTO `ch_t_d_areainfo` VALUES (331124, '松阳县', 3, 00000331100);
INSERT INTO `ch_t_d_areainfo` VALUES (331125, '云和县', 3, 00000331100);
INSERT INTO `ch_t_d_areainfo` VALUES (331126, '庆元县', 3, 00000331100);
INSERT INTO `ch_t_d_areainfo` VALUES (331127, '景宁畲族自治县', 3, 00000331100);
INSERT INTO `ch_t_d_areainfo` VALUES (331181, '龙泉市', 3, 00000331100);
INSERT INTO `ch_t_d_areainfo` VALUES (340000, '安徽省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (340100, '合肥市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (340102, '瑶海区', 3, 00000340100);
INSERT INTO `ch_t_d_areainfo` VALUES (340103, '庐阳区', 3, 00000340100);
INSERT INTO `ch_t_d_areainfo` VALUES (340104, '蜀山区', 3, 00000340100);
INSERT INTO `ch_t_d_areainfo` VALUES (340111, '包河区', 3, 00000340100);
INSERT INTO `ch_t_d_areainfo` VALUES (340121, '长丰县', 3, 00000340100);
INSERT INTO `ch_t_d_areainfo` VALUES (340122, '肥东县', 3, 00000340100);
INSERT INTO `ch_t_d_areainfo` VALUES (340123, '肥西县', 3, 00000340100);
INSERT INTO `ch_t_d_areainfo` VALUES (340124, '庐江县', 3, 00000340100);
INSERT INTO `ch_t_d_areainfo` VALUES (340181, '巢湖市', 3, 00000340100);
INSERT INTO `ch_t_d_areainfo` VALUES (340200, '芜湖市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (340202, '镜湖区', 3, 00000340200);
INSERT INTO `ch_t_d_areainfo` VALUES (340203, '弋江区', 3, 00000340200);
INSERT INTO `ch_t_d_areainfo` VALUES (340207, '鸠江区', 3, 00000340200);
INSERT INTO `ch_t_d_areainfo` VALUES (340208, '三山区', 3, 00000340200);
INSERT INTO `ch_t_d_areainfo` VALUES (340221, '芜湖县', 3, 00000340200);
INSERT INTO `ch_t_d_areainfo` VALUES (340222, '繁昌县', 3, 00000340200);
INSERT INTO `ch_t_d_areainfo` VALUES (340223, '南陵县', 3, 00000340200);
INSERT INTO `ch_t_d_areainfo` VALUES (340225, '无为县', 3, 00000340200);
INSERT INTO `ch_t_d_areainfo` VALUES (340300, '蚌埠市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (340302, '龙子湖区', 3, 00000340300);
INSERT INTO `ch_t_d_areainfo` VALUES (340303, '蚌山区', 3, 00000340300);
INSERT INTO `ch_t_d_areainfo` VALUES (340304, '禹会区', 3, 00000340300);
INSERT INTO `ch_t_d_areainfo` VALUES (340311, '淮上区', 3, 00000340300);
INSERT INTO `ch_t_d_areainfo` VALUES (340321, '怀远县', 3, 00000340300);
INSERT INTO `ch_t_d_areainfo` VALUES (340322, '五河县', 3, 00000340300);
INSERT INTO `ch_t_d_areainfo` VALUES (340323, '固镇县', 3, 00000340300);
INSERT INTO `ch_t_d_areainfo` VALUES (340400, '淮南市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (340402, '大通区', 3, 00000340400);
INSERT INTO `ch_t_d_areainfo` VALUES (340403, '田家庵区', 3, 00000340400);
INSERT INTO `ch_t_d_areainfo` VALUES (340404, '谢家集区', 3, 00000340400);
INSERT INTO `ch_t_d_areainfo` VALUES (340405, '八公山区', 3, 00000340400);
INSERT INTO `ch_t_d_areainfo` VALUES (340406, '潘集区', 3, 00000340400);
INSERT INTO `ch_t_d_areainfo` VALUES (340421, '凤台县', 3, 00000340400);
INSERT INTO `ch_t_d_areainfo` VALUES (340422, '寿县', 3, 00000340400);
INSERT INTO `ch_t_d_areainfo` VALUES (340500, '马鞍山市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (340503, '花山区', 3, 00000340500);
INSERT INTO `ch_t_d_areainfo` VALUES (340504, '雨山区', 3, 00000340500);
INSERT INTO `ch_t_d_areainfo` VALUES (340506, '博望区', 3, 00000340500);
INSERT INTO `ch_t_d_areainfo` VALUES (340521, '当涂县', 3, 00000340500);
INSERT INTO `ch_t_d_areainfo` VALUES (340522, '含山县', 3, 00000340500);
INSERT INTO `ch_t_d_areainfo` VALUES (340523, '和县', 3, 00000340500);
INSERT INTO `ch_t_d_areainfo` VALUES (340600, '淮北市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (340602, '杜集区', 3, 00000340600);
INSERT INTO `ch_t_d_areainfo` VALUES (340603, '相山区', 3, 00000340600);
INSERT INTO `ch_t_d_areainfo` VALUES (340604, '烈山区', 3, 00000340600);
INSERT INTO `ch_t_d_areainfo` VALUES (340621, '濉溪县', 3, 00000340600);
INSERT INTO `ch_t_d_areainfo` VALUES (340700, '铜陵市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (340705, '铜官区', 3, 00000340700);
INSERT INTO `ch_t_d_areainfo` VALUES (340706, '义安区', 3, 00000340700);
INSERT INTO `ch_t_d_areainfo` VALUES (340711, '郊区', 3, 00000340700);
INSERT INTO `ch_t_d_areainfo` VALUES (340722, '枞阳县', 3, 00000340700);
INSERT INTO `ch_t_d_areainfo` VALUES (340800, '安庆市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (340802, '迎江区', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (340803, '大观区', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (340811, '宜秀区', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (340822, '怀宁县', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (340824, '潜山县', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (340825, '太湖县', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (340826, '宿松县', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (340827, '望江县', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (340828, '岳西县', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (340881, '桐城市', 3, 00000340800);
INSERT INTO `ch_t_d_areainfo` VALUES (341000, '黄山市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (341002, '屯溪区', 3, 00000341000);
INSERT INTO `ch_t_d_areainfo` VALUES (341003, '黄山区', 3, 00000341000);
INSERT INTO `ch_t_d_areainfo` VALUES (341004, '徽州区', 3, 00000341000);
INSERT INTO `ch_t_d_areainfo` VALUES (341021, '歙县', 3, 00000341000);
INSERT INTO `ch_t_d_areainfo` VALUES (341022, '休宁县', 3, 00000341000);
INSERT INTO `ch_t_d_areainfo` VALUES (341023, '黟县', 3, 00000341000);
INSERT INTO `ch_t_d_areainfo` VALUES (341024, '祁门县', 3, 00000341000);
INSERT INTO `ch_t_d_areainfo` VALUES (341100, '滁州市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (341102, '琅琊区', 3, 00000341100);
INSERT INTO `ch_t_d_areainfo` VALUES (341103, '南谯区', 3, 00000341100);
INSERT INTO `ch_t_d_areainfo` VALUES (341122, '来安县', 3, 00000341100);
INSERT INTO `ch_t_d_areainfo` VALUES (341124, '全椒县', 3, 00000341100);
INSERT INTO `ch_t_d_areainfo` VALUES (341125, '定远县', 3, 00000341100);
INSERT INTO `ch_t_d_areainfo` VALUES (341126, '凤阳县', 3, 00000341100);
INSERT INTO `ch_t_d_areainfo` VALUES (341181, '天长市', 3, 00000341100);
INSERT INTO `ch_t_d_areainfo` VALUES (341182, '明光市', 3, 00000341100);
INSERT INTO `ch_t_d_areainfo` VALUES (341200, '阜阳市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (341202, '颍州区', 3, 00000341200);
INSERT INTO `ch_t_d_areainfo` VALUES (341203, '颍东区', 3, 00000341200);
INSERT INTO `ch_t_d_areainfo` VALUES (341204, '颍泉区', 3, 00000341200);
INSERT INTO `ch_t_d_areainfo` VALUES (341221, '临泉县', 3, 00000341200);
INSERT INTO `ch_t_d_areainfo` VALUES (341222, '太和县', 3, 00000341200);
INSERT INTO `ch_t_d_areainfo` VALUES (341225, '阜南县', 3, 00000341200);
INSERT INTO `ch_t_d_areainfo` VALUES (341226, '颍上县', 3, 00000341200);
INSERT INTO `ch_t_d_areainfo` VALUES (341282, '界首市', 3, 00000341200);
INSERT INTO `ch_t_d_areainfo` VALUES (341300, '宿州市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (341302, '埇桥区', 3, 00000341300);
INSERT INTO `ch_t_d_areainfo` VALUES (341321, '砀山县', 3, 00000341300);
INSERT INTO `ch_t_d_areainfo` VALUES (341322, '萧县', 3, 00000341300);
INSERT INTO `ch_t_d_areainfo` VALUES (341323, '灵璧县', 3, 00000341300);
INSERT INTO `ch_t_d_areainfo` VALUES (341324, '泗县', 3, 00000341300);
INSERT INTO `ch_t_d_areainfo` VALUES (341500, '六安市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (341502, '金安区', 3, 00000341500);
INSERT INTO `ch_t_d_areainfo` VALUES (341503, '裕安区', 3, 00000341500);
INSERT INTO `ch_t_d_areainfo` VALUES (341504, '叶集区', 3, 00000341500);
INSERT INTO `ch_t_d_areainfo` VALUES (341522, '霍邱县', 3, 00000341500);
INSERT INTO `ch_t_d_areainfo` VALUES (341523, '舒城县', 3, 00000341500);
INSERT INTO `ch_t_d_areainfo` VALUES (341524, '金寨县', 3, 00000341500);
INSERT INTO `ch_t_d_areainfo` VALUES (341525, '霍山县', 3, 00000341500);
INSERT INTO `ch_t_d_areainfo` VALUES (341600, '亳州市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (341602, '谯城区', 3, 00000341600);
INSERT INTO `ch_t_d_areainfo` VALUES (341621, '涡阳县', 3, 00000341600);
INSERT INTO `ch_t_d_areainfo` VALUES (341622, '蒙城县', 3, 00000341600);
INSERT INTO `ch_t_d_areainfo` VALUES (341623, '利辛县', 3, 00000341600);
INSERT INTO `ch_t_d_areainfo` VALUES (341700, '池州市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (341702, '贵池区', 3, 00000341700);
INSERT INTO `ch_t_d_areainfo` VALUES (341721, '东至县', 3, 00000341700);
INSERT INTO `ch_t_d_areainfo` VALUES (341722, '石台县', 3, 00000341700);
INSERT INTO `ch_t_d_areainfo` VALUES (341723, '青阳县', 3, 00000341700);
INSERT INTO `ch_t_d_areainfo` VALUES (341800, '宣城市', 2, 00000340000);
INSERT INTO `ch_t_d_areainfo` VALUES (341802, '宣州区', 3, 00000341800);
INSERT INTO `ch_t_d_areainfo` VALUES (341821, '郎溪县', 3, 00000341800);
INSERT INTO `ch_t_d_areainfo` VALUES (341822, '广德县', 3, 00000341800);
INSERT INTO `ch_t_d_areainfo` VALUES (341823, '泾县', 3, 00000341800);
INSERT INTO `ch_t_d_areainfo` VALUES (341824, '绩溪县', 3, 00000341800);
INSERT INTO `ch_t_d_areainfo` VALUES (341825, '旌德县', 3, 00000341800);
INSERT INTO `ch_t_d_areainfo` VALUES (341881, '宁国市', 3, 00000341800);
INSERT INTO `ch_t_d_areainfo` VALUES (350000, '福建省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (350100, '福州市', 2, 00000350000);
INSERT INTO `ch_t_d_areainfo` VALUES (350102, '鼓楼区', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350103, '台江区', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350104, '仓山区', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350105, '马尾区', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350111, '晋安区', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350121, '闽侯县', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350122, '连江县', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350123, '罗源县', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350124, '闽清县', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350125, '永泰县', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350128, '平潭县', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350181, '福清市', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350182, '长乐市', 3, 00000350100);
INSERT INTO `ch_t_d_areainfo` VALUES (350200, '厦门市', 2, 00000350000);
INSERT INTO `ch_t_d_areainfo` VALUES (350203, '思明区', 3, 00000350200);
INSERT INTO `ch_t_d_areainfo` VALUES (350205, '海沧区', 3, 00000350200);
INSERT INTO `ch_t_d_areainfo` VALUES (350206, '湖里区', 3, 00000350200);
INSERT INTO `ch_t_d_areainfo` VALUES (350211, '集美区', 3, 00000350200);
INSERT INTO `ch_t_d_areainfo` VALUES (350212, '同安区', 3, 00000350200);
INSERT INTO `ch_t_d_areainfo` VALUES (350213, '翔安区', 3, 00000350200);
INSERT INTO `ch_t_d_areainfo` VALUES (350300, '莆田市', 2, 00000350000);
INSERT INTO `ch_t_d_areainfo` VALUES (350302, '城厢区', 3, 00000350300);
INSERT INTO `ch_t_d_areainfo` VALUES (350303, '涵江区', 3, 00000350300);
INSERT INTO `ch_t_d_areainfo` VALUES (350304, '荔城区', 3, 00000350300);
INSERT INTO `ch_t_d_areainfo` VALUES (350305, '秀屿区', 3, 00000350300);
INSERT INTO `ch_t_d_areainfo` VALUES (350322, '仙游县', 3, 00000350300);
INSERT INTO `ch_t_d_areainfo` VALUES (350400, '三明市', 2, 00000350000);
INSERT INTO `ch_t_d_areainfo` VALUES (350402, '梅列区', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350403, '三元区', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350421, '明溪县', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350423, '清流县', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350424, '宁化县', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350425, '大田县', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350426, '尤溪县', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350427, '沙县', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350428, '将乐县', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350429, '泰宁县', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350430, '建宁县', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350481, '永安市', 3, 00000350400);
INSERT INTO `ch_t_d_areainfo` VALUES (350500, '泉州市', 2, 00000350000);
INSERT INTO `ch_t_d_areainfo` VALUES (350502, '鲤城区', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350503, '丰泽区', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350504, '洛江区', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350505, '泉港区', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350521, '惠安县', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350524, '安溪县', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350525, '永春县', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350526, '德化县', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350527, '金门县', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350581, '石狮市', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350582, '晋江市', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350583, '南安市', 3, 00000350500);
INSERT INTO `ch_t_d_areainfo` VALUES (350600, '漳州市', 2, 00000350000);
INSERT INTO `ch_t_d_areainfo` VALUES (350602, '芗城区', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350603, '龙文区', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350622, '云霄县', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350623, '漳浦县', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350624, '诏安县', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350625, '长泰县', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350626, '东山县', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350627, '南靖县', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350628, '平和县', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350629, '华安县', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350681, '龙海市', 3, 00000350600);
INSERT INTO `ch_t_d_areainfo` VALUES (350700, '南平市', 2, 00000350000);
INSERT INTO `ch_t_d_areainfo` VALUES (350702, '延平区', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350703, '建阳区', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350721, '顺昌县', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350722, '浦城县', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350723, '光泽县', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350724, '松溪县', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350725, '政和县', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350781, '邵武市', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350782, '武夷山市', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350783, '建瓯市', 3, 00000350700);
INSERT INTO `ch_t_d_areainfo` VALUES (350800, '龙岩市', 2, 00000350000);
INSERT INTO `ch_t_d_areainfo` VALUES (350802, '新罗区', 3, 00000350800);
INSERT INTO `ch_t_d_areainfo` VALUES (350803, '永定区', 3, 00000350800);
INSERT INTO `ch_t_d_areainfo` VALUES (350821, '长汀县', 3, 00000350800);
INSERT INTO `ch_t_d_areainfo` VALUES (350823, '上杭县', 3, 00000350800);
INSERT INTO `ch_t_d_areainfo` VALUES (350824, '武平县', 3, 00000350800);
INSERT INTO `ch_t_d_areainfo` VALUES (350825, '连城县', 3, 00000350800);
INSERT INTO `ch_t_d_areainfo` VALUES (350881, '漳平市', 3, 00000350800);
INSERT INTO `ch_t_d_areainfo` VALUES (350900, '宁德市', 2, 00000350000);
INSERT INTO `ch_t_d_areainfo` VALUES (350902, '蕉城区', 3, 00000350900);
INSERT INTO `ch_t_d_areainfo` VALUES (350921, '霞浦县', 3, 00000350900);
INSERT INTO `ch_t_d_areainfo` VALUES (350922, '古田县', 3, 00000350900);
INSERT INTO `ch_t_d_areainfo` VALUES (350923, '屏南县', 3, 00000350900);
INSERT INTO `ch_t_d_areainfo` VALUES (350924, '寿宁县', 3, 00000350900);
INSERT INTO `ch_t_d_areainfo` VALUES (350925, '周宁县', 3, 00000350900);
INSERT INTO `ch_t_d_areainfo` VALUES (350926, '柘荣县', 3, 00000350900);
INSERT INTO `ch_t_d_areainfo` VALUES (350981, '福安市', 3, 00000350900);
INSERT INTO `ch_t_d_areainfo` VALUES (350982, '福鼎市', 3, 00000350900);
INSERT INTO `ch_t_d_areainfo` VALUES (360000, '江西省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (360100, '南昌市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (360102, '东湖区', 3, 00000360100);
INSERT INTO `ch_t_d_areainfo` VALUES (360103, '西湖区', 3, 00000360100);
INSERT INTO `ch_t_d_areainfo` VALUES (360104, '青云谱区', 3, 00000360100);
INSERT INTO `ch_t_d_areainfo` VALUES (360105, '湾里区', 3, 00000360100);
INSERT INTO `ch_t_d_areainfo` VALUES (360111, '青山湖区', 3, 00000360100);
INSERT INTO `ch_t_d_areainfo` VALUES (360112, '新建区', 3, 00000360100);
INSERT INTO `ch_t_d_areainfo` VALUES (360121, '南昌县', 3, 00000360100);
INSERT INTO `ch_t_d_areainfo` VALUES (360123, '安义县', 3, 00000360100);
INSERT INTO `ch_t_d_areainfo` VALUES (360124, '进贤县', 3, 00000360100);
INSERT INTO `ch_t_d_areainfo` VALUES (360200, '景德镇市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (360202, '昌江区', 3, 00000360200);
INSERT INTO `ch_t_d_areainfo` VALUES (360203, '珠山区', 3, 00000360200);
INSERT INTO `ch_t_d_areainfo` VALUES (360222, '浮梁县', 3, 00000360200);
INSERT INTO `ch_t_d_areainfo` VALUES (360281, '乐平市', 3, 00000360200);
INSERT INTO `ch_t_d_areainfo` VALUES (360300, '萍乡市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (360302, '安源区', 3, 00000360300);
INSERT INTO `ch_t_d_areainfo` VALUES (360313, '湘东区', 3, 00000360300);
INSERT INTO `ch_t_d_areainfo` VALUES (360321, '莲花县', 3, 00000360300);
INSERT INTO `ch_t_d_areainfo` VALUES (360322, '上栗县', 3, 00000360300);
INSERT INTO `ch_t_d_areainfo` VALUES (360323, '芦溪县', 3, 00000360300);
INSERT INTO `ch_t_d_areainfo` VALUES (360400, '九江市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (360402, '濂溪区', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360403, '浔阳区', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360421, '九江县', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360423, '武宁县', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360424, '修水县', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360425, '永修县', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360426, '德安县', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360428, '都昌县', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360429, '湖口县', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360430, '彭泽县', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360481, '瑞昌市', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360482, '共青城市', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360483, '庐山市', 3, 00000360400);
INSERT INTO `ch_t_d_areainfo` VALUES (360500, '新余市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (360502, '渝水区', 3, 00000360500);
INSERT INTO `ch_t_d_areainfo` VALUES (360521, '分宜县', 3, 00000360500);
INSERT INTO `ch_t_d_areainfo` VALUES (360600, '鹰潭市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (360602, '月湖区', 3, 00000360600);
INSERT INTO `ch_t_d_areainfo` VALUES (360622, '余江县', 3, 00000360600);
INSERT INTO `ch_t_d_areainfo` VALUES (360681, '贵溪市', 3, 00000360600);
INSERT INTO `ch_t_d_areainfo` VALUES (360700, '赣州市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (360702, '章贡区', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360703, '南康区', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360721, '赣县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360722, '信丰县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360723, '大余县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360724, '上犹县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360725, '崇义县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360726, '安远县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360727, '龙南县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360728, '定南县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360729, '全南县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360730, '宁都县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360731, '于都县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360732, '兴国县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360733, '会昌县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360734, '寻乌县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360735, '石城县', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360781, '瑞金市', 3, 00000360700);
INSERT INTO `ch_t_d_areainfo` VALUES (360800, '吉安市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (360802, '吉州区', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360803, '青原区', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360821, '吉安县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360822, '吉水县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360823, '峡江县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360824, '新干县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360825, '永丰县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360826, '泰和县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360827, '遂川县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360828, '万安县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360829, '安福县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360830, '永新县', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360881, '井冈山市', 3, 00000360800);
INSERT INTO `ch_t_d_areainfo` VALUES (360900, '宜春市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (360902, '袁州区', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (360921, '奉新县', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (360922, '万载县', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (360923, '上高县', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (360924, '宜丰县', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (360925, '靖安县', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (360926, '铜鼓县', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (360981, '丰城市', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (360982, '樟树市', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (360983, '高安市', 3, 00000360900);
INSERT INTO `ch_t_d_areainfo` VALUES (361000, '抚州市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (361002, '临川区', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361021, '南城县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361022, '黎川县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361023, '南丰县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361024, '崇仁县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361025, '乐安县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361026, '宜黄县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361027, '金溪县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361028, '资溪县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361029, '东乡县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361030, '广昌县', 3, 00000361000);
INSERT INTO `ch_t_d_areainfo` VALUES (361100, '上饶市', 2, 00000360000);
INSERT INTO `ch_t_d_areainfo` VALUES (361102, '信州区', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361103, '广丰区', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361121, '上饶县', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361123, '玉山县', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361124, '铅山县', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361125, '横峰县', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361126, '弋阳县', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361127, '余干县', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361128, '鄱阳县', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361129, '万年县', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361130, '婺源县', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (361181, '德兴市', 3, 00000361100);
INSERT INTO `ch_t_d_areainfo` VALUES (370000, '山东省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (370100, '济南市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (370102, '历下区', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370103, '市中区', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370104, '槐荫区', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370105, '天桥区', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370112, '历城区', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370113, '长清区', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370124, '平阴县', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370125, '济阳县', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370126, '商河县', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370181, '章丘市', 3, 00000370100);
INSERT INTO `ch_t_d_areainfo` VALUES (370200, '青岛市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (370202, '市南区', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370203, '市北区', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370211, '黄岛区', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370212, '崂山区', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370213, '李沧区', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370214, '城阳区', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370281, '胶州市', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370282, '即墨市', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370283, '平度市', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370285, '莱西市', 3, 00000370200);
INSERT INTO `ch_t_d_areainfo` VALUES (370300, '淄博市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (370302, '淄川区', 3, 00000370300);
INSERT INTO `ch_t_d_areainfo` VALUES (370303, '张店区', 3, 00000370300);
INSERT INTO `ch_t_d_areainfo` VALUES (370304, '博山区', 3, 00000370300);
INSERT INTO `ch_t_d_areainfo` VALUES (370305, '临淄区', 3, 00000370300);
INSERT INTO `ch_t_d_areainfo` VALUES (370306, '周村区', 3, 00000370300);
INSERT INTO `ch_t_d_areainfo` VALUES (370321, '桓台县', 3, 00000370300);
INSERT INTO `ch_t_d_areainfo` VALUES (370322, '高青县', 3, 00000370300);
INSERT INTO `ch_t_d_areainfo` VALUES (370323, '沂源县', 3, 00000370300);
INSERT INTO `ch_t_d_areainfo` VALUES (370400, '枣庄市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (370402, '市中区', 3, 00000370400);
INSERT INTO `ch_t_d_areainfo` VALUES (370403, '薛城区', 3, 00000370400);
INSERT INTO `ch_t_d_areainfo` VALUES (370404, '峄城区', 3, 00000370400);
INSERT INTO `ch_t_d_areainfo` VALUES (370405, '台儿庄区', 3, 00000370400);
INSERT INTO `ch_t_d_areainfo` VALUES (370406, '山亭区', 3, 00000370400);
INSERT INTO `ch_t_d_areainfo` VALUES (370481, '滕州市', 3, 00000370400);
INSERT INTO `ch_t_d_areainfo` VALUES (370500, '东营市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (370502, '东营区', 3, 00000370500);
INSERT INTO `ch_t_d_areainfo` VALUES (370503, '河口区', 3, 00000370500);
INSERT INTO `ch_t_d_areainfo` VALUES (370505, '垦利区', 3, 00000370500);
INSERT INTO `ch_t_d_areainfo` VALUES (370522, '利津县', 3, 00000370500);
INSERT INTO `ch_t_d_areainfo` VALUES (370523, '广饶县', 3, 00000370500);
INSERT INTO `ch_t_d_areainfo` VALUES (370600, '烟台市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (370602, '芝罘区', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370611, '福山区', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370612, '牟平区', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370613, '莱山区', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370634, '长岛县', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370681, '龙口市', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370682, '莱阳市', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370683, '莱州市', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370684, '蓬莱市', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370685, '招远市', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370686, '栖霞市', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370687, '海阳市', 3, 00000370600);
INSERT INTO `ch_t_d_areainfo` VALUES (370700, '潍坊市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (370702, '潍城区', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370703, '寒亭区', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370704, '坊子区', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370705, '奎文区', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370724, '临朐县', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370725, '昌乐县', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370781, '青州市', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370782, '诸城市', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370783, '寿光市', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370784, '安丘市', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370785, '高密市', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370786, '昌邑市', 3, 00000370700);
INSERT INTO `ch_t_d_areainfo` VALUES (370800, '济宁市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (370811, '任城区', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370812, '兖州区', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370826, '微山县', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370827, '鱼台县', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370828, '金乡县', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370829, '嘉祥县', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370830, '汶上县', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370831, '泗水县', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370832, '梁山县', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370881, '曲阜市', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370883, '邹城市', 3, 00000370800);
INSERT INTO `ch_t_d_areainfo` VALUES (370900, '泰安市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (370902, '泰山区', 3, 00000370900);
INSERT INTO `ch_t_d_areainfo` VALUES (370911, '岱岳区', 3, 00000370900);
INSERT INTO `ch_t_d_areainfo` VALUES (370921, '宁阳县', 3, 00000370900);
INSERT INTO `ch_t_d_areainfo` VALUES (370923, '东平县', 3, 00000370900);
INSERT INTO `ch_t_d_areainfo` VALUES (370982, '新泰市', 3, 00000370900);
INSERT INTO `ch_t_d_areainfo` VALUES (370983, '肥城市', 3, 00000370900);
INSERT INTO `ch_t_d_areainfo` VALUES (371000, '威海市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (371002, '环翠区', 3, 00000371000);
INSERT INTO `ch_t_d_areainfo` VALUES (371003, '文登区', 3, 00000371000);
INSERT INTO `ch_t_d_areainfo` VALUES (371082, '荣成市', 3, 00000371000);
INSERT INTO `ch_t_d_areainfo` VALUES (371083, '乳山市', 3, 00000371000);
INSERT INTO `ch_t_d_areainfo` VALUES (371100, '日照市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (371102, '东港区', 3, 00000371100);
INSERT INTO `ch_t_d_areainfo` VALUES (371103, '岚山区', 3, 00000371100);
INSERT INTO `ch_t_d_areainfo` VALUES (371121, '五莲县', 3, 00000371100);
INSERT INTO `ch_t_d_areainfo` VALUES (371122, '莒县', 3, 00000371100);
INSERT INTO `ch_t_d_areainfo` VALUES (371200, '莱芜市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (371202, '莱城区', 3, 00000371200);
INSERT INTO `ch_t_d_areainfo` VALUES (371203, '钢城区', 3, 00000371200);
INSERT INTO `ch_t_d_areainfo` VALUES (371300, '临沂市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (371302, '兰山区', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371311, '罗庄区', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371312, '河东区', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371321, '沂南县', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371322, '郯城县', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371323, '沂水县', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371324, '兰陵县', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371325, '费县', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371326, '平邑县', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371327, '莒南县', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371328, '蒙阴县', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371329, '临沭县', 3, 00000371300);
INSERT INTO `ch_t_d_areainfo` VALUES (371400, '德州市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (371402, '德城区', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371403, '陵城区', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371422, '宁津县', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371423, '庆云县', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371424, '临邑县', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371425, '齐河县', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371426, '平原县', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371427, '夏津县', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371428, '武城县', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371481, '乐陵市', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371482, '禹城市', 3, 00000371400);
INSERT INTO `ch_t_d_areainfo` VALUES (371500, '聊城市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (371502, '东昌府区', 3, 00000371500);
INSERT INTO `ch_t_d_areainfo` VALUES (371521, '阳谷县', 3, 00000371500);
INSERT INTO `ch_t_d_areainfo` VALUES (371522, '莘县', 3, 00000371500);
INSERT INTO `ch_t_d_areainfo` VALUES (371523, '茌平县', 3, 00000371500);
INSERT INTO `ch_t_d_areainfo` VALUES (371524, '东阿县', 3, 00000371500);
INSERT INTO `ch_t_d_areainfo` VALUES (371525, '冠县', 3, 00000371500);
INSERT INTO `ch_t_d_areainfo` VALUES (371526, '高唐县', 3, 00000371500);
INSERT INTO `ch_t_d_areainfo` VALUES (371581, '临清市', 3, 00000371500);
INSERT INTO `ch_t_d_areainfo` VALUES (371600, '滨州市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (371602, '滨城区', 3, 00000371600);
INSERT INTO `ch_t_d_areainfo` VALUES (371603, '沾化区', 3, 00000371600);
INSERT INTO `ch_t_d_areainfo` VALUES (371621, '惠民县', 3, 00000371600);
INSERT INTO `ch_t_d_areainfo` VALUES (371622, '阳信县', 3, 00000371600);
INSERT INTO `ch_t_d_areainfo` VALUES (371623, '无棣县', 3, 00000371600);
INSERT INTO `ch_t_d_areainfo` VALUES (371625, '博兴县', 3, 00000371600);
INSERT INTO `ch_t_d_areainfo` VALUES (371626, '邹平县', 3, 00000371600);
INSERT INTO `ch_t_d_areainfo` VALUES (371700, '菏泽市', 2, 00000370000);
INSERT INTO `ch_t_d_areainfo` VALUES (371702, '牡丹区', 3, 00000371700);
INSERT INTO `ch_t_d_areainfo` VALUES (371703, '定陶区', 3, 00000371700);
INSERT INTO `ch_t_d_areainfo` VALUES (371721, '曹县', 3, 00000371700);
INSERT INTO `ch_t_d_areainfo` VALUES (371722, '单县', 3, 00000371700);
INSERT INTO `ch_t_d_areainfo` VALUES (371723, '成武县', 3, 00000371700);
INSERT INTO `ch_t_d_areainfo` VALUES (371724, '巨野县', 3, 00000371700);
INSERT INTO `ch_t_d_areainfo` VALUES (371725, '郓城县', 3, 00000371700);
INSERT INTO `ch_t_d_areainfo` VALUES (371726, '鄄城县', 3, 00000371700);
INSERT INTO `ch_t_d_areainfo` VALUES (371728, '东明县', 3, 00000371700);
INSERT INTO `ch_t_d_areainfo` VALUES (410000, '河南省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (410100, '郑州市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (410102, '中原区', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410103, '二七区', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410104, '管城回族区', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410105, '金水区', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410106, '上街区', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410108, '惠济区', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410122, '中牟县', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410181, '巩义市', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410182, '荥阳市', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410183, '新密市', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410184, '新郑市', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410185, '登封市', 3, 00000410100);
INSERT INTO `ch_t_d_areainfo` VALUES (410200, '开封市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (410202, '龙亭区', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410203, '顺河回族区', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410204, '鼓楼区', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410205, '禹王台区', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410211, '金明区', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410212, '祥符区', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410221, '杞县', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410222, '通许县', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410223, '尉氏县', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410225, '兰考县', 3, 00000410200);
INSERT INTO `ch_t_d_areainfo` VALUES (410300, '洛阳市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (410302, '老城区', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410303, '西工区', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410304, '瀍河回族区', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410305, '涧西区', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410306, '吉利区', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410311, '洛龙区', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410322, '孟津县', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410323, '新安县', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410324, '栾川县', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410325, '嵩县', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410326, '汝阳县', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410327, '宜阳县', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410328, '洛宁县', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410329, '伊川县', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410381, '偃师市', 3, 00000410300);
INSERT INTO `ch_t_d_areainfo` VALUES (410400, '平顶山市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (410402, '新华区', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410403, '卫东区', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410404, '石龙区', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410411, '湛河区', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410421, '宝丰县', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410422, '叶县', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410423, '鲁山县', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410425, '郏县', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410481, '舞钢市', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410482, '汝州市', 3, 00000410400);
INSERT INTO `ch_t_d_areainfo` VALUES (410500, '安阳市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (410502, '文峰区', 3, 00000410500);
INSERT INTO `ch_t_d_areainfo` VALUES (410503, '北关区', 3, 00000410500);
INSERT INTO `ch_t_d_areainfo` VALUES (410505, '殷都区', 3, 00000410500);
INSERT INTO `ch_t_d_areainfo` VALUES (410506, '龙安区', 3, 00000410500);
INSERT INTO `ch_t_d_areainfo` VALUES (410522, '安阳县', 3, 00000410500);
INSERT INTO `ch_t_d_areainfo` VALUES (410523, '汤阴县', 3, 00000410500);
INSERT INTO `ch_t_d_areainfo` VALUES (410526, '滑县', 3, 00000410500);
INSERT INTO `ch_t_d_areainfo` VALUES (410527, '内黄县', 3, 00000410500);
INSERT INTO `ch_t_d_areainfo` VALUES (410581, '林州市', 3, 00000410500);
INSERT INTO `ch_t_d_areainfo` VALUES (410600, '鹤壁市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (410602, '鹤山区', 3, 00000410600);
INSERT INTO `ch_t_d_areainfo` VALUES (410603, '山城区', 3, 00000410600);
INSERT INTO `ch_t_d_areainfo` VALUES (410611, '淇滨区', 3, 00000410600);
INSERT INTO `ch_t_d_areainfo` VALUES (410621, '浚县', 3, 00000410600);
INSERT INTO `ch_t_d_areainfo` VALUES (410622, '淇县', 3, 00000410600);
INSERT INTO `ch_t_d_areainfo` VALUES (410700, '新乡市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (410702, '红旗区', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410703, '卫滨区', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410704, '凤泉区', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410711, '牧野区', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410721, '新乡县', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410724, '获嘉县', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410725, '原阳县', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410726, '延津县', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410727, '封丘县', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410728, '长垣县', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410781, '卫辉市', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410782, '辉县市', 3, 00000410700);
INSERT INTO `ch_t_d_areainfo` VALUES (410800, '焦作市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (410802, '解放区', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410803, '中站区', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410804, '马村区', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410811, '山阳区', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410821, '修武县', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410822, '博爱县', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410823, '武陟县', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410825, '温县', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410882, '沁阳市', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410883, '孟州市', 3, 00000410800);
INSERT INTO `ch_t_d_areainfo` VALUES (410900, '濮阳市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (410902, '华龙区', 3, 00000410900);
INSERT INTO `ch_t_d_areainfo` VALUES (410922, '清丰县', 3, 00000410900);
INSERT INTO `ch_t_d_areainfo` VALUES (410923, '南乐县', 3, 00000410900);
INSERT INTO `ch_t_d_areainfo` VALUES (410926, '范县', 3, 00000410900);
INSERT INTO `ch_t_d_areainfo` VALUES (410927, '台前县', 3, 00000410900);
INSERT INTO `ch_t_d_areainfo` VALUES (410928, '濮阳县', 3, 00000410900);
INSERT INTO `ch_t_d_areainfo` VALUES (411000, '许昌市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (411002, '魏都区', 3, 00000411000);
INSERT INTO `ch_t_d_areainfo` VALUES (411023, '许昌县', 3, 00000411000);
INSERT INTO `ch_t_d_areainfo` VALUES (411024, '鄢陵县', 3, 00000411000);
INSERT INTO `ch_t_d_areainfo` VALUES (411025, '襄城县', 3, 00000411000);
INSERT INTO `ch_t_d_areainfo` VALUES (411081, '禹州市', 3, 00000411000);
INSERT INTO `ch_t_d_areainfo` VALUES (411082, '长葛市', 3, 00000411000);
INSERT INTO `ch_t_d_areainfo` VALUES (411100, '漯河市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (411102, '源汇区', 3, 00000411100);
INSERT INTO `ch_t_d_areainfo` VALUES (411103, '郾城区', 3, 00000411100);
INSERT INTO `ch_t_d_areainfo` VALUES (411104, '召陵区', 3, 00000411100);
INSERT INTO `ch_t_d_areainfo` VALUES (411121, '舞阳县', 3, 00000411100);
INSERT INTO `ch_t_d_areainfo` VALUES (411122, '临颍县', 3, 00000411100);
INSERT INTO `ch_t_d_areainfo` VALUES (411200, '三门峡市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (411202, '湖滨区', 3, 00000411200);
INSERT INTO `ch_t_d_areainfo` VALUES (411203, '陕州区', 3, 00000411200);
INSERT INTO `ch_t_d_areainfo` VALUES (411221, '渑池县', 3, 00000411200);
INSERT INTO `ch_t_d_areainfo` VALUES (411224, '卢氏县', 3, 00000411200);
INSERT INTO `ch_t_d_areainfo` VALUES (411281, '义马市', 3, 00000411200);
INSERT INTO `ch_t_d_areainfo` VALUES (411282, '灵宝市', 3, 00000411200);
INSERT INTO `ch_t_d_areainfo` VALUES (411300, '南阳市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (411302, '宛城区', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411303, '卧龙区', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411321, '南召县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411322, '方城县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411323, '西峡县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411324, '镇平县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411325, '内乡县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411326, '淅川县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411327, '社旗县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411328, '唐河县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411329, '新野县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411330, '桐柏县', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411381, '邓州市', 3, 00000411300);
INSERT INTO `ch_t_d_areainfo` VALUES (411400, '商丘市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (411402, '梁园区', 3, 00000411400);
INSERT INTO `ch_t_d_areainfo` VALUES (411403, '睢阳区', 3, 00000411400);
INSERT INTO `ch_t_d_areainfo` VALUES (411421, '民权县', 3, 00000411400);
INSERT INTO `ch_t_d_areainfo` VALUES (411422, '睢县', 3, 00000411400);
INSERT INTO `ch_t_d_areainfo` VALUES (411423, '宁陵县', 3, 00000411400);
INSERT INTO `ch_t_d_areainfo` VALUES (411424, '柘城县', 3, 00000411400);
INSERT INTO `ch_t_d_areainfo` VALUES (411425, '虞城县', 3, 00000411400);
INSERT INTO `ch_t_d_areainfo` VALUES (411426, '夏邑县', 3, 00000411400);
INSERT INTO `ch_t_d_areainfo` VALUES (411481, '永城市', 3, 00000411400);
INSERT INTO `ch_t_d_areainfo` VALUES (411500, '信阳市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (411502, '浉河区', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411503, '平桥区', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411521, '罗山县', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411522, '光山县', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411523, '新县', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411524, '商城县', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411525, '固始县', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411526, '潢川县', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411527, '淮滨县', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411528, '息县', 3, 00000411500);
INSERT INTO `ch_t_d_areainfo` VALUES (411600, '周口市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (411602, '川汇区', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411621, '扶沟县', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411622, '西华县', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411623, '商水县', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411624, '沈丘县', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411625, '郸城县', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411626, '淮阳县', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411627, '太康县', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411628, '鹿邑县', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411681, '项城市', 3, 00000411600);
INSERT INTO `ch_t_d_areainfo` VALUES (411700, '驻马店市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (411702, '驿城区', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411721, '西平县', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411722, '上蔡县', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411723, '平舆县', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411724, '正阳县', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411725, '确山县', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411726, '泌阳县', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411727, '汝南县', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411728, '遂平县', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411729, '新蔡县', 3, 00000411700);
INSERT INTO `ch_t_d_areainfo` VALUES (411800, '济源市', 2, 00000410000);
INSERT INTO `ch_t_d_areainfo` VALUES (420000, '湖北省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (420100, '武汉市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (420102, '江岸区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420103, '江汉区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420104, '硚口区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420105, '汉阳区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420106, '武昌区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420107, '青山区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420111, '洪山区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420112, '东西湖区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420113, '汉南区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420114, '蔡甸区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420115, '江夏区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420116, '黄陂区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420117, '新洲区', 3, 00000420100);
INSERT INTO `ch_t_d_areainfo` VALUES (420200, '黄石市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (420202, '黄石港区', 3, 00000420200);
INSERT INTO `ch_t_d_areainfo` VALUES (420203, '西塞山区', 3, 00000420200);
INSERT INTO `ch_t_d_areainfo` VALUES (420204, '下陆区', 3, 00000420200);
INSERT INTO `ch_t_d_areainfo` VALUES (420205, '铁山区', 3, 00000420200);
INSERT INTO `ch_t_d_areainfo` VALUES (420222, '阳新县', 3, 00000420200);
INSERT INTO `ch_t_d_areainfo` VALUES (420281, '大冶市', 3, 00000420200);
INSERT INTO `ch_t_d_areainfo` VALUES (420300, '十堰市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (420302, '茅箭区', 3, 00000420300);
INSERT INTO `ch_t_d_areainfo` VALUES (420303, '张湾区', 3, 00000420300);
INSERT INTO `ch_t_d_areainfo` VALUES (420304, '郧阳区', 3, 00000420300);
INSERT INTO `ch_t_d_areainfo` VALUES (420322, '郧西县', 3, 00000420300);
INSERT INTO `ch_t_d_areainfo` VALUES (420323, '竹山县', 3, 00000420300);
INSERT INTO `ch_t_d_areainfo` VALUES (420324, '竹溪县', 3, 00000420300);
INSERT INTO `ch_t_d_areainfo` VALUES (420325, '房县', 3, 00000420300);
INSERT INTO `ch_t_d_areainfo` VALUES (420381, '丹江口市', 3, 00000420300);
INSERT INTO `ch_t_d_areainfo` VALUES (420500, '宜昌市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (420502, '西陵区', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420503, '伍家岗区', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420504, '点军区', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420505, '猇亭区', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420506, '夷陵区', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420525, '远安县', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420526, '兴山县', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420527, '秭归县', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420528, '长阳土家族自治县', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420529, '五峰土家族自治县', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420581, '宜都市', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420582, '当阳市', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420583, '枝江市', 3, 00000420500);
INSERT INTO `ch_t_d_areainfo` VALUES (420600, '襄阳市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (420602, '襄城区', 3, 00000420600);
INSERT INTO `ch_t_d_areainfo` VALUES (420606, '樊城区', 3, 00000420600);
INSERT INTO `ch_t_d_areainfo` VALUES (420607, '襄州区', 3, 00000420600);
INSERT INTO `ch_t_d_areainfo` VALUES (420624, '南漳县', 3, 00000420600);
INSERT INTO `ch_t_d_areainfo` VALUES (420625, '谷城县', 3, 00000420600);
INSERT INTO `ch_t_d_areainfo` VALUES (420626, '保康县', 3, 00000420600);
INSERT INTO `ch_t_d_areainfo` VALUES (420682, '老河口市', 3, 00000420600);
INSERT INTO `ch_t_d_areainfo` VALUES (420683, '枣阳市', 3, 00000420600);
INSERT INTO `ch_t_d_areainfo` VALUES (420684, '宜城市', 3, 00000420600);
INSERT INTO `ch_t_d_areainfo` VALUES (420700, '鄂州市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (420702, '梁子湖区', 3, 00000420700);
INSERT INTO `ch_t_d_areainfo` VALUES (420703, '华容区', 3, 00000420700);
INSERT INTO `ch_t_d_areainfo` VALUES (420704, '鄂城区', 3, 00000420700);
INSERT INTO `ch_t_d_areainfo` VALUES (420800, '荆门市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (420802, '东宝区', 3, 00000420800);
INSERT INTO `ch_t_d_areainfo` VALUES (420804, '掇刀区', 3, 00000420800);
INSERT INTO `ch_t_d_areainfo` VALUES (420821, '京山县', 3, 00000420800);
INSERT INTO `ch_t_d_areainfo` VALUES (420822, '沙洋县', 3, 00000420800);
INSERT INTO `ch_t_d_areainfo` VALUES (420881, '钟祥市', 3, 00000420800);
INSERT INTO `ch_t_d_areainfo` VALUES (420900, '孝感市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (420902, '孝南区', 3, 00000420900);
INSERT INTO `ch_t_d_areainfo` VALUES (420921, '孝昌县', 3, 00000420900);
INSERT INTO `ch_t_d_areainfo` VALUES (420922, '大悟县', 3, 00000420900);
INSERT INTO `ch_t_d_areainfo` VALUES (420923, '云梦县', 3, 00000420900);
INSERT INTO `ch_t_d_areainfo` VALUES (420981, '应城市', 3, 00000420900);
INSERT INTO `ch_t_d_areainfo` VALUES (420982, '安陆市', 3, 00000420900);
INSERT INTO `ch_t_d_areainfo` VALUES (420984, '汉川市', 3, 00000420900);
INSERT INTO `ch_t_d_areainfo` VALUES (421000, '荆州市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (421002, '沙市区', 3, 00000421000);
INSERT INTO `ch_t_d_areainfo` VALUES (421003, '荆州区', 3, 00000421000);
INSERT INTO `ch_t_d_areainfo` VALUES (421022, '公安县', 3, 00000421000);
INSERT INTO `ch_t_d_areainfo` VALUES (421023, '监利县', 3, 00000421000);
INSERT INTO `ch_t_d_areainfo` VALUES (421024, '江陵县', 3, 00000421000);
INSERT INTO `ch_t_d_areainfo` VALUES (421081, '石首市', 3, 00000421000);
INSERT INTO `ch_t_d_areainfo` VALUES (421083, '洪湖市', 3, 00000421000);
INSERT INTO `ch_t_d_areainfo` VALUES (421087, '松滋市', 3, 00000421000);
INSERT INTO `ch_t_d_areainfo` VALUES (421100, '黄冈市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (421102, '黄州区', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421121, '团风县', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421122, '红安县', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421123, '罗田县', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421124, '英山县', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421125, '浠水县', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421126, '蕲春县', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421127, '黄梅县', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421181, '麻城市', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421182, '武穴市', 3, 00000421100);
INSERT INTO `ch_t_d_areainfo` VALUES (421200, '咸宁市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (421202, '咸安区', 3, 00000421200);
INSERT INTO `ch_t_d_areainfo` VALUES (421221, '嘉鱼县', 3, 00000421200);
INSERT INTO `ch_t_d_areainfo` VALUES (421222, '通城县', 3, 00000421200);
INSERT INTO `ch_t_d_areainfo` VALUES (421223, '崇阳县', 3, 00000421200);
INSERT INTO `ch_t_d_areainfo` VALUES (421224, '通山县', 3, 00000421200);
INSERT INTO `ch_t_d_areainfo` VALUES (421281, '赤壁市', 3, 00000421200);
INSERT INTO `ch_t_d_areainfo` VALUES (421300, '随州市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (421303, '曾都区', 3, 00000421300);
INSERT INTO `ch_t_d_areainfo` VALUES (421321, '随县', 3, 00000421300);
INSERT INTO `ch_t_d_areainfo` VALUES (421381, '广水市', 3, 00000421300);
INSERT INTO `ch_t_d_areainfo` VALUES (421400, '仙桃市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (421500, '潜江市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (421600, '天门市', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (421700, '神农架林区', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (422800, '恩施土家族苗族自治州', 2, 00000420000);
INSERT INTO `ch_t_d_areainfo` VALUES (422801, '恩施市', 3, 00000422800);
INSERT INTO `ch_t_d_areainfo` VALUES (422802, '利川市', 3, 00000422800);
INSERT INTO `ch_t_d_areainfo` VALUES (422822, '建始县', 3, 00000422800);
INSERT INTO `ch_t_d_areainfo` VALUES (422823, '巴东县', 3, 00000422800);
INSERT INTO `ch_t_d_areainfo` VALUES (422825, '宣恩县', 3, 00000422800);
INSERT INTO `ch_t_d_areainfo` VALUES (422826, '咸丰县', 3, 00000422800);
INSERT INTO `ch_t_d_areainfo` VALUES (422827, '来凤县', 3, 00000422800);
INSERT INTO `ch_t_d_areainfo` VALUES (422828, '鹤峰县', 3, 00000422800);
INSERT INTO `ch_t_d_areainfo` VALUES (430000, '湖南省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (430100, '长沙市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (430102, '芙蓉区', 3, 00000430100);
INSERT INTO `ch_t_d_areainfo` VALUES (430103, '天心区', 3, 00000430100);
INSERT INTO `ch_t_d_areainfo` VALUES (430104, '岳麓区', 3, 00000430100);
INSERT INTO `ch_t_d_areainfo` VALUES (430105, '开福区', 3, 00000430100);
INSERT INTO `ch_t_d_areainfo` VALUES (430111, '雨花区', 3, 00000430100);
INSERT INTO `ch_t_d_areainfo` VALUES (430112, '望城区', 3, 00000430100);
INSERT INTO `ch_t_d_areainfo` VALUES (430121, '长沙县', 3, 00000430100);
INSERT INTO `ch_t_d_areainfo` VALUES (430124, '宁乡县', 3, 00000430100);
INSERT INTO `ch_t_d_areainfo` VALUES (430181, '浏阳市', 3, 00000430100);
INSERT INTO `ch_t_d_areainfo` VALUES (430200, '株洲市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (430202, '荷塘区', 3, 00000430200);
INSERT INTO `ch_t_d_areainfo` VALUES (430203, '芦淞区', 3, 00000430200);
INSERT INTO `ch_t_d_areainfo` VALUES (430204, '石峰区', 3, 00000430200);
INSERT INTO `ch_t_d_areainfo` VALUES (430211, '天元区', 3, 00000430200);
INSERT INTO `ch_t_d_areainfo` VALUES (430221, '株洲县', 3, 00000430200);
INSERT INTO `ch_t_d_areainfo` VALUES (430223, '攸县', 3, 00000430200);
INSERT INTO `ch_t_d_areainfo` VALUES (430224, '茶陵县', 3, 00000430200);
INSERT INTO `ch_t_d_areainfo` VALUES (430225, '炎陵县', 3, 00000430200);
INSERT INTO `ch_t_d_areainfo` VALUES (430281, '醴陵市', 3, 00000430200);
INSERT INTO `ch_t_d_areainfo` VALUES (430300, '湘潭市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (430302, '雨湖区', 3, 00000430300);
INSERT INTO `ch_t_d_areainfo` VALUES (430304, '岳塘区', 3, 00000430300);
INSERT INTO `ch_t_d_areainfo` VALUES (430321, '湘潭县', 3, 00000430300);
INSERT INTO `ch_t_d_areainfo` VALUES (430381, '湘乡市', 3, 00000430300);
INSERT INTO `ch_t_d_areainfo` VALUES (430382, '韶山市', 3, 00000430300);
INSERT INTO `ch_t_d_areainfo` VALUES (430400, '衡阳市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (430405, '珠晖区', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430406, '雁峰区', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430407, '石鼓区', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430408, '蒸湘区', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430412, '南岳区', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430421, '衡阳县', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430422, '衡南县', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430423, '衡山县', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430424, '衡东县', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430426, '祁东县', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430481, '耒阳市', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430482, '常宁市', 3, 00000430400);
INSERT INTO `ch_t_d_areainfo` VALUES (430500, '邵阳市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (430502, '双清区', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430503, '大祥区', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430511, '北塔区', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430521, '邵东县', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430522, '新邵县', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430523, '邵阳县', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430524, '隆回县', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430525, '洞口县', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430527, '绥宁县', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430528, '新宁县', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430529, '城步苗族自治县', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430581, '武冈市', 3, 00000430500);
INSERT INTO `ch_t_d_areainfo` VALUES (430600, '岳阳市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (430602, '岳阳楼区', 3, 00000430600);
INSERT INTO `ch_t_d_areainfo` VALUES (430603, '云溪区', 3, 00000430600);
INSERT INTO `ch_t_d_areainfo` VALUES (430611, '君山区', 3, 00000430600);
INSERT INTO `ch_t_d_areainfo` VALUES (430621, '岳阳县', 3, 00000430600);
INSERT INTO `ch_t_d_areainfo` VALUES (430623, '华容县', 3, 00000430600);
INSERT INTO `ch_t_d_areainfo` VALUES (430624, '湘阴县', 3, 00000430600);
INSERT INTO `ch_t_d_areainfo` VALUES (430626, '平江县', 3, 00000430600);
INSERT INTO `ch_t_d_areainfo` VALUES (430681, '汨罗市', 3, 00000430600);
INSERT INTO `ch_t_d_areainfo` VALUES (430682, '临湘市', 3, 00000430600);
INSERT INTO `ch_t_d_areainfo` VALUES (430700, '常德市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (430702, '武陵区', 3, 00000430700);
INSERT INTO `ch_t_d_areainfo` VALUES (430703, '鼎城区', 3, 00000430700);
INSERT INTO `ch_t_d_areainfo` VALUES (430721, '安乡县', 3, 00000430700);
INSERT INTO `ch_t_d_areainfo` VALUES (430722, '汉寿县', 3, 00000430700);
INSERT INTO `ch_t_d_areainfo` VALUES (430723, '澧县', 3, 00000430700);
INSERT INTO `ch_t_d_areainfo` VALUES (430724, '临澧县', 3, 00000430700);
INSERT INTO `ch_t_d_areainfo` VALUES (430725, '桃源县', 3, 00000430700);
INSERT INTO `ch_t_d_areainfo` VALUES (430726, '石门县', 3, 00000430700);
INSERT INTO `ch_t_d_areainfo` VALUES (430781, '津市市', 3, 00000430700);
INSERT INTO `ch_t_d_areainfo` VALUES (430800, '张家界市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (430802, '永定区', 3, 00000430800);
INSERT INTO `ch_t_d_areainfo` VALUES (430811, '武陵源区', 3, 00000430800);
INSERT INTO `ch_t_d_areainfo` VALUES (430821, '慈利县', 3, 00000430800);
INSERT INTO `ch_t_d_areainfo` VALUES (430822, '桑植县', 3, 00000430800);
INSERT INTO `ch_t_d_areainfo` VALUES (430900, '益阳市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (430902, '资阳区', 3, 00000430900);
INSERT INTO `ch_t_d_areainfo` VALUES (430903, '赫山区', 3, 00000430900);
INSERT INTO `ch_t_d_areainfo` VALUES (430921, '南县', 3, 00000430900);
INSERT INTO `ch_t_d_areainfo` VALUES (430922, '桃江县', 3, 00000430900);
INSERT INTO `ch_t_d_areainfo` VALUES (430923, '安化县', 3, 00000430900);
INSERT INTO `ch_t_d_areainfo` VALUES (430981, '沅江市', 3, 00000430900);
INSERT INTO `ch_t_d_areainfo` VALUES (431000, '郴州市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (431002, '北湖区', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431003, '苏仙区', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431021, '桂阳县', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431022, '宜章县', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431023, '永兴县', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431024, '嘉禾县', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431025, '临武县', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431026, '汝城县', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431027, '桂东县', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431028, '安仁县', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431081, '资兴市', 3, 00000431000);
INSERT INTO `ch_t_d_areainfo` VALUES (431100, '永州市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (431102, '零陵区', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431103, '冷水滩区', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431121, '祁阳县', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431122, '东安县', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431123, '双牌县', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431124, '道县', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431125, '江永县', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431126, '宁远县', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431127, '蓝山县', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431128, '新田县', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431129, '江华瑶族自治县', 3, 00000431100);
INSERT INTO `ch_t_d_areainfo` VALUES (431200, '怀化市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (431202, '鹤城区', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431221, '中方县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431222, '沅陵县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431223, '辰溪县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431224, '溆浦县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431225, '会同县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431226, '麻阳苗族自治县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431227, '新晃侗族自治县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431228, '芷江侗族自治县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431229, '靖州苗族侗族自治县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431230, '通道侗族自治县', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431281, '洪江市', 3, 00000431200);
INSERT INTO `ch_t_d_areainfo` VALUES (431300, '娄底市', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (431302, '娄星区', 3, 00000431300);
INSERT INTO `ch_t_d_areainfo` VALUES (431321, '双峰县', 3, 00000431300);
INSERT INTO `ch_t_d_areainfo` VALUES (431322, '新化县', 3, 00000431300);
INSERT INTO `ch_t_d_areainfo` VALUES (431381, '冷水江市', 3, 00000431300);
INSERT INTO `ch_t_d_areainfo` VALUES (431382, '涟源市', 3, 00000431300);
INSERT INTO `ch_t_d_areainfo` VALUES (433100, '湘西土家族苗族自治州', 2, 00000430000);
INSERT INTO `ch_t_d_areainfo` VALUES (433101, '吉首市', 3, 00000433100);
INSERT INTO `ch_t_d_areainfo` VALUES (433122, '泸溪县', 3, 00000433100);
INSERT INTO `ch_t_d_areainfo` VALUES (433123, '凤凰县', 3, 00000433100);
INSERT INTO `ch_t_d_areainfo` VALUES (433124, '花垣县', 3, 00000433100);
INSERT INTO `ch_t_d_areainfo` VALUES (433125, '保靖县', 3, 00000433100);
INSERT INTO `ch_t_d_areainfo` VALUES (433126, '古丈县', 3, 00000433100);
INSERT INTO `ch_t_d_areainfo` VALUES (433127, '永顺县', 3, 00000433100);
INSERT INTO `ch_t_d_areainfo` VALUES (433130, '龙山县', 3, 00000433100);
INSERT INTO `ch_t_d_areainfo` VALUES (440000, '广东省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (440100, '广州市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (440103, '荔湾区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440104, '越秀区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440105, '海珠区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440106, '天河区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440111, '白云区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440112, '黄埔区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440113, '番禺区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440114, '花都区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440115, '南沙区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440117, '从化区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440118, '增城区', 3, 00000440100);
INSERT INTO `ch_t_d_areainfo` VALUES (440200, '韶关市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (440203, '武江区', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440204, '浈江区', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440205, '曲江区', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440222, '始兴县', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440224, '仁化县', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440229, '翁源县', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440232, '乳源瑶族自治县', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440233, '新丰县', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440281, '乐昌市', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440282, '南雄市', 3, 00000440200);
INSERT INTO `ch_t_d_areainfo` VALUES (440300, '深圳市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (440303, '罗湖区', 3, 00000440300);
INSERT INTO `ch_t_d_areainfo` VALUES (440304, '福田区', 3, 00000440300);
INSERT INTO `ch_t_d_areainfo` VALUES (440305, '南山区', 3, 00000440300);
INSERT INTO `ch_t_d_areainfo` VALUES (440306, '宝安区', 3, 00000440300);
INSERT INTO `ch_t_d_areainfo` VALUES (440307, '龙岗区', 3, 00000440300);
INSERT INTO `ch_t_d_areainfo` VALUES (440308, '盐田区', 3, 00000440300);
INSERT INTO `ch_t_d_areainfo` VALUES (440400, '珠海市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (440402, '香洲区', 3, 00000440400);
INSERT INTO `ch_t_d_areainfo` VALUES (440403, '斗门区', 3, 00000440400);
INSERT INTO `ch_t_d_areainfo` VALUES (440404, '金湾区', 3, 00000440400);
INSERT INTO `ch_t_d_areainfo` VALUES (440500, '汕头市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (440507, '龙湖区', 3, 00000440500);
INSERT INTO `ch_t_d_areainfo` VALUES (440511, '金平区', 3, 00000440500);
INSERT INTO `ch_t_d_areainfo` VALUES (440512, '濠江区', 3, 00000440500);
INSERT INTO `ch_t_d_areainfo` VALUES (440513, '潮阳区', 3, 00000440500);
INSERT INTO `ch_t_d_areainfo` VALUES (440514, '潮南区', 3, 00000440500);
INSERT INTO `ch_t_d_areainfo` VALUES (440515, '澄海区', 3, 00000440500);
INSERT INTO `ch_t_d_areainfo` VALUES (440523, '南澳县', 3, 00000440500);
INSERT INTO `ch_t_d_areainfo` VALUES (440600, '佛山市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (440604, '禅城区', 3, 00000440600);
INSERT INTO `ch_t_d_areainfo` VALUES (440605, '南海区', 3, 00000440600);
INSERT INTO `ch_t_d_areainfo` VALUES (440606, '顺德区', 3, 00000440600);
INSERT INTO `ch_t_d_areainfo` VALUES (440607, '三水区', 3, 00000440600);
INSERT INTO `ch_t_d_areainfo` VALUES (440608, '高明区', 3, 00000440600);
INSERT INTO `ch_t_d_areainfo` VALUES (440700, '江门市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (440703, '蓬江区', 3, 00000440700);
INSERT INTO `ch_t_d_areainfo` VALUES (440704, '江海区', 3, 00000440700);
INSERT INTO `ch_t_d_areainfo` VALUES (440705, '新会区', 3, 00000440700);
INSERT INTO `ch_t_d_areainfo` VALUES (440781, '台山市', 3, 00000440700);
INSERT INTO `ch_t_d_areainfo` VALUES (440783, '开平市', 3, 00000440700);
INSERT INTO `ch_t_d_areainfo` VALUES (440784, '鹤山市', 3, 00000440700);
INSERT INTO `ch_t_d_areainfo` VALUES (440785, '恩平市', 3, 00000440700);
INSERT INTO `ch_t_d_areainfo` VALUES (440800, '湛江市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (440802, '赤坎区', 3, 00000440800);
INSERT INTO `ch_t_d_areainfo` VALUES (440803, '霞山区', 3, 00000440800);
INSERT INTO `ch_t_d_areainfo` VALUES (440804, '坡头区', 3, 00000440800);
INSERT INTO `ch_t_d_areainfo` VALUES (440811, '麻章区', 3, 00000440800);
INSERT INTO `ch_t_d_areainfo` VALUES (440823, '遂溪县', 3, 00000440800);
INSERT INTO `ch_t_d_areainfo` VALUES (440825, '徐闻县', 3, 00000440800);
INSERT INTO `ch_t_d_areainfo` VALUES (440881, '廉江市', 3, 00000440800);
INSERT INTO `ch_t_d_areainfo` VALUES (440882, '雷州市', 3, 00000440800);
INSERT INTO `ch_t_d_areainfo` VALUES (440883, '吴川市', 3, 00000440800);
INSERT INTO `ch_t_d_areainfo` VALUES (440900, '茂名市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (440902, '茂南区', 3, 00000440900);
INSERT INTO `ch_t_d_areainfo` VALUES (440904, '电白区', 3, 00000440900);
INSERT INTO `ch_t_d_areainfo` VALUES (440981, '高州市', 3, 00000440900);
INSERT INTO `ch_t_d_areainfo` VALUES (440982, '化州市', 3, 00000440900);
INSERT INTO `ch_t_d_areainfo` VALUES (440983, '信宜市', 3, 00000440900);
INSERT INTO `ch_t_d_areainfo` VALUES (441200, '肇庆市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (441202, '端州区', 3, 00000441200);
INSERT INTO `ch_t_d_areainfo` VALUES (441203, '鼎湖区', 3, 00000441200);
INSERT INTO `ch_t_d_areainfo` VALUES (441204, '高要区', 3, 00000441200);
INSERT INTO `ch_t_d_areainfo` VALUES (441223, '广宁县', 3, 00000441200);
INSERT INTO `ch_t_d_areainfo` VALUES (441224, '怀集县', 3, 00000441200);
INSERT INTO `ch_t_d_areainfo` VALUES (441225, '封开县', 3, 00000441200);
INSERT INTO `ch_t_d_areainfo` VALUES (441226, '德庆县', 3, 00000441200);
INSERT INTO `ch_t_d_areainfo` VALUES (441284, '四会市', 3, 00000441200);
INSERT INTO `ch_t_d_areainfo` VALUES (441300, '惠州市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (441302, '惠城区', 3, 00000441300);
INSERT INTO `ch_t_d_areainfo` VALUES (441303, '惠阳区', 3, 00000441300);
INSERT INTO `ch_t_d_areainfo` VALUES (441322, '博罗县', 3, 00000441300);
INSERT INTO `ch_t_d_areainfo` VALUES (441323, '惠东县', 3, 00000441300);
INSERT INTO `ch_t_d_areainfo` VALUES (441324, '龙门县', 3, 00000441300);
INSERT INTO `ch_t_d_areainfo` VALUES (441400, '梅州市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (441402, '梅江区', 3, 00000441400);
INSERT INTO `ch_t_d_areainfo` VALUES (441403, '梅县区', 3, 00000441400);
INSERT INTO `ch_t_d_areainfo` VALUES (441422, '大埔县', 3, 00000441400);
INSERT INTO `ch_t_d_areainfo` VALUES (441423, '丰顺县', 3, 00000441400);
INSERT INTO `ch_t_d_areainfo` VALUES (441424, '五华县', 3, 00000441400);
INSERT INTO `ch_t_d_areainfo` VALUES (441426, '平远县', 3, 00000441400);
INSERT INTO `ch_t_d_areainfo` VALUES (441427, '蕉岭县', 3, 00000441400);
INSERT INTO `ch_t_d_areainfo` VALUES (441481, '兴宁市', 3, 00000441400);
INSERT INTO `ch_t_d_areainfo` VALUES (441500, '汕尾市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (441502, '城区', 3, 00000441500);
INSERT INTO `ch_t_d_areainfo` VALUES (441521, '海丰县', 3, 00000441500);
INSERT INTO `ch_t_d_areainfo` VALUES (441523, '陆河县', 3, 00000441500);
INSERT INTO `ch_t_d_areainfo` VALUES (441581, '陆丰市', 3, 00000441500);
INSERT INTO `ch_t_d_areainfo` VALUES (441600, '河源市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (441602, '源城区', 3, 00000441600);
INSERT INTO `ch_t_d_areainfo` VALUES (441621, '紫金县', 3, 00000441600);
INSERT INTO `ch_t_d_areainfo` VALUES (441622, '龙川县', 3, 00000441600);
INSERT INTO `ch_t_d_areainfo` VALUES (441623, '连平县', 3, 00000441600);
INSERT INTO `ch_t_d_areainfo` VALUES (441624, '和平县', 3, 00000441600);
INSERT INTO `ch_t_d_areainfo` VALUES (441625, '东源县', 3, 00000441600);
INSERT INTO `ch_t_d_areainfo` VALUES (441700, '阳江市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (441702, '江城区', 3, 00000441700);
INSERT INTO `ch_t_d_areainfo` VALUES (441704, '阳东区', 3, 00000441700);
INSERT INTO `ch_t_d_areainfo` VALUES (441721, '阳西县', 3, 00000441700);
INSERT INTO `ch_t_d_areainfo` VALUES (441781, '阳春市', 3, 00000441700);
INSERT INTO `ch_t_d_areainfo` VALUES (441800, '清远市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (441802, '清城区', 3, 00000441800);
INSERT INTO `ch_t_d_areainfo` VALUES (441803, '清新区', 3, 00000441800);
INSERT INTO `ch_t_d_areainfo` VALUES (441821, '佛冈县', 3, 00000441800);
INSERT INTO `ch_t_d_areainfo` VALUES (441823, '阳山县', 3, 00000441800);
INSERT INTO `ch_t_d_areainfo` VALUES (441825, '连山壮族瑶族自治县', 3, 00000441800);
INSERT INTO `ch_t_d_areainfo` VALUES (441826, '连南瑶族自治县', 3, 00000441800);
INSERT INTO `ch_t_d_areainfo` VALUES (441881, '英德市', 3, 00000441800);
INSERT INTO `ch_t_d_areainfo` VALUES (441882, '连州市', 3, 00000441800);
INSERT INTO `ch_t_d_areainfo` VALUES (441900, '东莞市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (442000, '中山市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (445100, '潮州市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (445102, '湘桥区', 3, 00000445100);
INSERT INTO `ch_t_d_areainfo` VALUES (445103, '潮安区', 3, 00000445100);
INSERT INTO `ch_t_d_areainfo` VALUES (445122, '饶平县', 3, 00000445100);
INSERT INTO `ch_t_d_areainfo` VALUES (445200, '揭阳市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (445202, '榕城区', 3, 00000445200);
INSERT INTO `ch_t_d_areainfo` VALUES (445203, '揭东区', 3, 00000445200);
INSERT INTO `ch_t_d_areainfo` VALUES (445222, '揭西县', 3, 00000445200);
INSERT INTO `ch_t_d_areainfo` VALUES (445224, '惠来县', 3, 00000445200);
INSERT INTO `ch_t_d_areainfo` VALUES (445281, '普宁市', 3, 00000445200);
INSERT INTO `ch_t_d_areainfo` VALUES (445300, '云浮市', 2, 00000440000);
INSERT INTO `ch_t_d_areainfo` VALUES (445302, '云城区', 3, 00000445300);
INSERT INTO `ch_t_d_areainfo` VALUES (445303, '云安区', 3, 00000445300);
INSERT INTO `ch_t_d_areainfo` VALUES (445321, '新兴县', 3, 00000445300);
INSERT INTO `ch_t_d_areainfo` VALUES (445322, '郁南县', 3, 00000445300);
INSERT INTO `ch_t_d_areainfo` VALUES (445381, '罗定市', 3, 00000445300);
INSERT INTO `ch_t_d_areainfo` VALUES (450000, '广西壮族自治区', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (450100, '南宁市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (450102, '兴宁区', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450103, '青秀区', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450105, '江南区', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450107, '西乡塘区', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450108, '良庆区', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450109, '邕宁区', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450110, '武鸣区', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450123, '隆安县', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450124, '马山县', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450125, '上林县', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450126, '宾阳县', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450127, '横县', 3, 00000450100);
INSERT INTO `ch_t_d_areainfo` VALUES (450200, '柳州市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (450202, '城中区', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450203, '鱼峰区', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450204, '柳南区', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450205, '柳北区', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450206, '柳江区', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450222, '柳城县', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450223, '鹿寨县', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450224, '融安县', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450225, '融水苗族自治县', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450226, '三江侗族自治县', 3, 00000450200);
INSERT INTO `ch_t_d_areainfo` VALUES (450300, '桂林市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (450302, '秀峰区', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450303, '叠彩区', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450304, '象山区', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450305, '七星区', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450311, '雁山区', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450312, '临桂区', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450321, '阳朔县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450323, '灵川县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450324, '全州县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450325, '兴安县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450326, '永福县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450327, '灌阳县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450328, '龙胜各族自治县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450329, '资源县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450330, '平乐县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450331, '荔浦县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450332, '恭城瑶族自治县', 3, 00000450300);
INSERT INTO `ch_t_d_areainfo` VALUES (450400, '梧州市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (450403, '万秀区', 3, 00000450400);
INSERT INTO `ch_t_d_areainfo` VALUES (450405, '长洲区', 3, 00000450400);
INSERT INTO `ch_t_d_areainfo` VALUES (450406, '龙圩区', 3, 00000450400);
INSERT INTO `ch_t_d_areainfo` VALUES (450421, '苍梧县', 3, 00000450400);
INSERT INTO `ch_t_d_areainfo` VALUES (450422, '藤县', 3, 00000450400);
INSERT INTO `ch_t_d_areainfo` VALUES (450423, '蒙山县', 3, 00000450400);
INSERT INTO `ch_t_d_areainfo` VALUES (450481, '岑溪市', 3, 00000450400);
INSERT INTO `ch_t_d_areainfo` VALUES (450500, '北海市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (450502, '海城区', 3, 00000450500);
INSERT INTO `ch_t_d_areainfo` VALUES (450503, '银海区', 3, 00000450500);
INSERT INTO `ch_t_d_areainfo` VALUES (450512, '铁山港区', 3, 00000450500);
INSERT INTO `ch_t_d_areainfo` VALUES (450521, '合浦县', 3, 00000450500);
INSERT INTO `ch_t_d_areainfo` VALUES (450600, '防城港市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (450602, '港口区', 3, 00000450600);
INSERT INTO `ch_t_d_areainfo` VALUES (450603, '防城区', 3, 00000450600);
INSERT INTO `ch_t_d_areainfo` VALUES (450621, '上思县', 3, 00000450600);
INSERT INTO `ch_t_d_areainfo` VALUES (450681, '东兴市', 3, 00000450600);
INSERT INTO `ch_t_d_areainfo` VALUES (450700, '钦州市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (450702, '钦南区', 3, 00000450700);
INSERT INTO `ch_t_d_areainfo` VALUES (450703, '钦北区', 3, 00000450700);
INSERT INTO `ch_t_d_areainfo` VALUES (450721, '灵山县', 3, 00000450700);
INSERT INTO `ch_t_d_areainfo` VALUES (450722, '浦北县', 3, 00000450700);
INSERT INTO `ch_t_d_areainfo` VALUES (450800, '贵港市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (450802, '港北区', 3, 00000450800);
INSERT INTO `ch_t_d_areainfo` VALUES (450803, '港南区', 3, 00000450800);
INSERT INTO `ch_t_d_areainfo` VALUES (450804, '覃塘区', 3, 00000450800);
INSERT INTO `ch_t_d_areainfo` VALUES (450821, '平南县', 3, 00000450800);
INSERT INTO `ch_t_d_areainfo` VALUES (450881, '桂平市', 3, 00000450800);
INSERT INTO `ch_t_d_areainfo` VALUES (450900, '玉林市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (450902, '玉州区', 3, 00000450900);
INSERT INTO `ch_t_d_areainfo` VALUES (450903, '福绵区', 3, 00000450900);
INSERT INTO `ch_t_d_areainfo` VALUES (450921, '容县', 3, 00000450900);
INSERT INTO `ch_t_d_areainfo` VALUES (450922, '陆川县', 3, 00000450900);
INSERT INTO `ch_t_d_areainfo` VALUES (450923, '博白县', 3, 00000450900);
INSERT INTO `ch_t_d_areainfo` VALUES (450924, '兴业县', 3, 00000450900);
INSERT INTO `ch_t_d_areainfo` VALUES (450981, '北流市', 3, 00000450900);
INSERT INTO `ch_t_d_areainfo` VALUES (451000, '百色市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (451002, '右江区', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451021, '田阳县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451022, '田东县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451023, '平果县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451024, '德保县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451026, '那坡县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451027, '凌云县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451028, '乐业县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451029, '田林县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451030, '西林县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451031, '隆林各族自治县', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451081, '靖西市', 3, 00000451000);
INSERT INTO `ch_t_d_areainfo` VALUES (451100, '贺州市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (451102, '八步区', 3, 00000451100);
INSERT INTO `ch_t_d_areainfo` VALUES (451103, '平桂区', 3, 00000451100);
INSERT INTO `ch_t_d_areainfo` VALUES (451121, '昭平县', 3, 00000451100);
INSERT INTO `ch_t_d_areainfo` VALUES (451122, '钟山县', 3, 00000451100);
INSERT INTO `ch_t_d_areainfo` VALUES (451123, '富川瑶族自治县', 3, 00000451100);
INSERT INTO `ch_t_d_areainfo` VALUES (451200, '河池市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (451202, '金城江区', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451221, '南丹县', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451222, '天峨县', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451223, '凤山县', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451224, '东兰县', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451225, '罗城仫佬族自治县', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451226, '环江毛南族自治县', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451227, '巴马瑶族自治县', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451228, '都安瑶族自治县', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451229, '大化瑶族自治县', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451281, '宜州市', 3, 00000451200);
INSERT INTO `ch_t_d_areainfo` VALUES (451300, '来宾市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (451302, '兴宾区', 3, 00000451300);
INSERT INTO `ch_t_d_areainfo` VALUES (451321, '忻城县', 3, 00000451300);
INSERT INTO `ch_t_d_areainfo` VALUES (451322, '象州县', 3, 00000451300);
INSERT INTO `ch_t_d_areainfo` VALUES (451323, '武宣县', 3, 00000451300);
INSERT INTO `ch_t_d_areainfo` VALUES (451324, '金秀瑶族自治县', 3, 00000451300);
INSERT INTO `ch_t_d_areainfo` VALUES (451381, '合山市', 3, 00000451300);
INSERT INTO `ch_t_d_areainfo` VALUES (451400, '崇左市', 2, 00000450000);
INSERT INTO `ch_t_d_areainfo` VALUES (451402, '江州区', 3, 00000451400);
INSERT INTO `ch_t_d_areainfo` VALUES (451421, '扶绥县', 3, 00000451400);
INSERT INTO `ch_t_d_areainfo` VALUES (451422, '宁明县', 3, 00000451400);
INSERT INTO `ch_t_d_areainfo` VALUES (451423, '龙州县', 3, 00000451400);
INSERT INTO `ch_t_d_areainfo` VALUES (451424, '大新县', 3, 00000451400);
INSERT INTO `ch_t_d_areainfo` VALUES (451425, '天等县', 3, 00000451400);
INSERT INTO `ch_t_d_areainfo` VALUES (451481, '凭祥市', 3, 00000451400);
INSERT INTO `ch_t_d_areainfo` VALUES (460000, '海南省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (460100, '海口市', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (460105, '秀英区', 3, 00000460100);
INSERT INTO `ch_t_d_areainfo` VALUES (460106, '龙华区', 3, 00000460100);
INSERT INTO `ch_t_d_areainfo` VALUES (460107, '琼山区', 3, 00000460100);
INSERT INTO `ch_t_d_areainfo` VALUES (460108, '美兰区', 3, 00000460100);
INSERT INTO `ch_t_d_areainfo` VALUES (460200, '三亚市', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (460202, '海棠区', 3, 00000460200);
INSERT INTO `ch_t_d_areainfo` VALUES (460203, '吉阳区', 3, 00000460200);
INSERT INTO `ch_t_d_areainfo` VALUES (460204, '天涯区', 3, 00000460200);
INSERT INTO `ch_t_d_areainfo` VALUES (460205, '崖州区', 3, 00000460200);
INSERT INTO `ch_t_d_areainfo` VALUES (460300, '三沙市', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (460400, '儋州市', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (460500, '五指山市', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (460600, '琼海市', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (460700, '文昌市', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (460800, '万宁市', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (460900, '东方市', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461000, '定安县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461100, '屯昌县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461200, '澄迈县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461300, '临高县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461400, '白沙黎族自治县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461500, '昌江黎族自治县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461600, '乐东黎族自治县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461700, '陵水黎族自治县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461800, '保亭黎族苗族自治县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (461900, '琼中黎族苗族自治县', 2, 00000460000);
INSERT INTO `ch_t_d_areainfo` VALUES (500000, '重庆市', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (500101, '万州区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500102, '涪陵区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500103, '渝中区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500104, '大渡口区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500105, '江北区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500106, '沙坪坝区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500107, '九龙坡区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500108, '南岸区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500109, '北碚区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500110, '綦江区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500111, '大足区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500112, '渝北区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500113, '巴南区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500114, '黔江区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500115, '长寿区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500116, '江津区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500117, '合川区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500118, '永川区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500119, '南川区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500120, '璧山区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500128, '梁平县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500129, '城口县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500130, '丰都县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500131, '垫江县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500132, '武隆县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500133, '忠县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500135, '云阳县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500136, '奉节县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500137, '巫山县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500138, '巫溪县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500140, '石柱土家族自治县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500141, '秀山土家族苗族自治县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500142, '酉阳土家族苗族自治县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500143, '彭水苗族土家族自治县', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500151, '铜梁区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500152, '潼南区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500153, '荣昌区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (500154, '开州区', 3, 00000500000);
INSERT INTO `ch_t_d_areainfo` VALUES (510000, '四川省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (510100, '成都市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (510104, '锦江区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510105, '青羊区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510106, '金牛区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510107, '武侯区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510108, '成华区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510109, '高新区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510112, '龙泉驿区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510113, '青白江区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510114, '新都区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510115, '温江区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510116, '双流区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510121, '金堂县', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510124, '郫都区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510129, '大邑县', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510131, '蒲江县', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510132, '新津县', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510141, '天府新区', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510181, '都江堰市', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510182, '彭州市', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510183, '邛崃市', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510184, '崇州市', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510185, '简阳市', 3, 00000510100);
INSERT INTO `ch_t_d_areainfo` VALUES (510300, '自贡市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (510302, '自流井区', 3, 00000510300);
INSERT INTO `ch_t_d_areainfo` VALUES (510303, '贡井区', 3, 00000510300);
INSERT INTO `ch_t_d_areainfo` VALUES (510304, '大安区', 3, 00000510300);
INSERT INTO `ch_t_d_areainfo` VALUES (510311, '沿滩区', 3, 00000510300);
INSERT INTO `ch_t_d_areainfo` VALUES (510321, '荣县', 3, 00000510300);
INSERT INTO `ch_t_d_areainfo` VALUES (510322, '富顺县', 3, 00000510300);
INSERT INTO `ch_t_d_areainfo` VALUES (510400, '攀枝花市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (510402, '东区', 3, 00000510400);
INSERT INTO `ch_t_d_areainfo` VALUES (510403, '西区', 3, 00000510400);
INSERT INTO `ch_t_d_areainfo` VALUES (510411, '仁和区', 3, 00000510400);
INSERT INTO `ch_t_d_areainfo` VALUES (510421, '米易县', 3, 00000510400);
INSERT INTO `ch_t_d_areainfo` VALUES (510422, '盐边县', 3, 00000510400);
INSERT INTO `ch_t_d_areainfo` VALUES (510500, '泸州市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (510502, '江阳区', 3, 00000510500);
INSERT INTO `ch_t_d_areainfo` VALUES (510503, '纳溪区', 3, 00000510500);
INSERT INTO `ch_t_d_areainfo` VALUES (510504, '龙马潭区', 3, 00000510500);
INSERT INTO `ch_t_d_areainfo` VALUES (510521, '泸县', 3, 00000510500);
INSERT INTO `ch_t_d_areainfo` VALUES (510522, '合江县', 3, 00000510500);
INSERT INTO `ch_t_d_areainfo` VALUES (510524, '叙永县', 3, 00000510500);
INSERT INTO `ch_t_d_areainfo` VALUES (510525, '古蔺县', 3, 00000510500);
INSERT INTO `ch_t_d_areainfo` VALUES (510600, '德阳市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (510603, '旌阳区', 3, 00000510600);
INSERT INTO `ch_t_d_areainfo` VALUES (510623, '中江县', 3, 00000510600);
INSERT INTO `ch_t_d_areainfo` VALUES (510626, '罗江县', 3, 00000510600);
INSERT INTO `ch_t_d_areainfo` VALUES (510681, '广汉市', 3, 00000510600);
INSERT INTO `ch_t_d_areainfo` VALUES (510682, '什邡市', 3, 00000510600);
INSERT INTO `ch_t_d_areainfo` VALUES (510683, '绵竹市', 3, 00000510600);
INSERT INTO `ch_t_d_areainfo` VALUES (510700, '绵阳市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (510703, '涪城区', 3, 00000510700);
INSERT INTO `ch_t_d_areainfo` VALUES (510704, '游仙区', 3, 00000510700);
INSERT INTO `ch_t_d_areainfo` VALUES (510705, '安州区', 3, 00000510700);
INSERT INTO `ch_t_d_areainfo` VALUES (510722, '三台县', 3, 00000510700);
INSERT INTO `ch_t_d_areainfo` VALUES (510723, '盐亭县', 3, 00000510700);
INSERT INTO `ch_t_d_areainfo` VALUES (510725, '梓潼县', 3, 00000510700);
INSERT INTO `ch_t_d_areainfo` VALUES (510726, '北川羌族自治县', 3, 00000510700);
INSERT INTO `ch_t_d_areainfo` VALUES (510727, '平武县', 3, 00000510700);
INSERT INTO `ch_t_d_areainfo` VALUES (510781, '江油市', 3, 00000510700);
INSERT INTO `ch_t_d_areainfo` VALUES (510800, '广元市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (510802, '利州区', 3, 00000510800);
INSERT INTO `ch_t_d_areainfo` VALUES (510811, '昭化区', 3, 00000510800);
INSERT INTO `ch_t_d_areainfo` VALUES (510812, '朝天区', 3, 00000510800);
INSERT INTO `ch_t_d_areainfo` VALUES (510821, '旺苍县', 3, 00000510800);
INSERT INTO `ch_t_d_areainfo` VALUES (510822, '青川县', 3, 00000510800);
INSERT INTO `ch_t_d_areainfo` VALUES (510823, '剑阁县', 3, 00000510800);
INSERT INTO `ch_t_d_areainfo` VALUES (510824, '苍溪县', 3, 00000510800);
INSERT INTO `ch_t_d_areainfo` VALUES (510900, '遂宁市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (510903, '船山区', 3, 00000510900);
INSERT INTO `ch_t_d_areainfo` VALUES (510904, '安居区', 3, 00000510900);
INSERT INTO `ch_t_d_areainfo` VALUES (510921, '蓬溪县', 3, 00000510900);
INSERT INTO `ch_t_d_areainfo` VALUES (510922, '射洪县', 3, 00000510900);
INSERT INTO `ch_t_d_areainfo` VALUES (510923, '大英县', 3, 00000510900);
INSERT INTO `ch_t_d_areainfo` VALUES (511000, '内江市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (511002, '市中区', 3, 00000511000);
INSERT INTO `ch_t_d_areainfo` VALUES (511011, '东兴区', 3, 00000511000);
INSERT INTO `ch_t_d_areainfo` VALUES (511024, '威远县', 3, 00000511000);
INSERT INTO `ch_t_d_areainfo` VALUES (511025, '资中县', 3, 00000511000);
INSERT INTO `ch_t_d_areainfo` VALUES (511028, '隆昌县', 3, 00000511000);
INSERT INTO `ch_t_d_areainfo` VALUES (511100, '乐山市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (511102, '市中区', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511111, '沙湾区', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511112, '五通桥区', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511113, '金口河区', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511123, '犍为县', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511124, '井研县', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511126, '夹江县', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511129, '沐川县', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511132, '峨边彝族自治县', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511133, '马边彝族自治县', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511181, '峨眉山市', 3, 00000511100);
INSERT INTO `ch_t_d_areainfo` VALUES (511300, '南充市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (511302, '顺庆区', 3, 00000511300);
INSERT INTO `ch_t_d_areainfo` VALUES (511303, '高坪区', 3, 00000511300);
INSERT INTO `ch_t_d_areainfo` VALUES (511304, '嘉陵区', 3, 00000511300);
INSERT INTO `ch_t_d_areainfo` VALUES (511321, '南部县', 3, 00000511300);
INSERT INTO `ch_t_d_areainfo` VALUES (511322, '营山县', 3, 00000511300);
INSERT INTO `ch_t_d_areainfo` VALUES (511323, '蓬安县', 3, 00000511300);
INSERT INTO `ch_t_d_areainfo` VALUES (511324, '仪陇县', 3, 00000511300);
INSERT INTO `ch_t_d_areainfo` VALUES (511325, '西充县', 3, 00000511300);
INSERT INTO `ch_t_d_areainfo` VALUES (511381, '阆中市', 3, 00000511300);
INSERT INTO `ch_t_d_areainfo` VALUES (511400, '眉山市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (511402, '东坡区', 3, 00000511400);
INSERT INTO `ch_t_d_areainfo` VALUES (511403, '彭山区', 3, 00000511400);
INSERT INTO `ch_t_d_areainfo` VALUES (511421, '仁寿县', 3, 00000511400);
INSERT INTO `ch_t_d_areainfo` VALUES (511423, '洪雅县', 3, 00000511400);
INSERT INTO `ch_t_d_areainfo` VALUES (511424, '丹棱县', 3, 00000511400);
INSERT INTO `ch_t_d_areainfo` VALUES (511425, '青神县', 3, 00000511400);
INSERT INTO `ch_t_d_areainfo` VALUES (511500, '宜宾市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (511502, '翠屏区', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511503, '南溪区', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511521, '宜宾县', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511523, '江安县', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511524, '长宁县', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511525, '高县', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511526, '珙县', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511527, '筠连县', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511528, '兴文县', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511529, '屏山县', 3, 00000511500);
INSERT INTO `ch_t_d_areainfo` VALUES (511600, '广安市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (511602, '广安区', 3, 00000511600);
INSERT INTO `ch_t_d_areainfo` VALUES (511603, '前锋区', 3, 00000511600);
INSERT INTO `ch_t_d_areainfo` VALUES (511621, '岳池县', 3, 00000511600);
INSERT INTO `ch_t_d_areainfo` VALUES (511622, '武胜县', 3, 00000511600);
INSERT INTO `ch_t_d_areainfo` VALUES (511623, '邻水县', 3, 00000511600);
INSERT INTO `ch_t_d_areainfo` VALUES (511681, '华蓥市', 3, 00000511600);
INSERT INTO `ch_t_d_areainfo` VALUES (511700, '达州市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (511702, '通川区', 3, 00000511700);
INSERT INTO `ch_t_d_areainfo` VALUES (511703, '达川区', 3, 00000511700);
INSERT INTO `ch_t_d_areainfo` VALUES (511722, '宣汉县', 3, 00000511700);
INSERT INTO `ch_t_d_areainfo` VALUES (511723, '开江县', 3, 00000511700);
INSERT INTO `ch_t_d_areainfo` VALUES (511724, '大竹县', 3, 00000511700);
INSERT INTO `ch_t_d_areainfo` VALUES (511725, '渠县', 3, 00000511700);
INSERT INTO `ch_t_d_areainfo` VALUES (511781, '万源市', 3, 00000511700);
INSERT INTO `ch_t_d_areainfo` VALUES (511800, '雅安市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (511802, '雨城区', 3, 00000511800);
INSERT INTO `ch_t_d_areainfo` VALUES (511803, '名山区', 3, 00000511800);
INSERT INTO `ch_t_d_areainfo` VALUES (511822, '荥经县', 3, 00000511800);
INSERT INTO `ch_t_d_areainfo` VALUES (511823, '汉源县', 3, 00000511800);
INSERT INTO `ch_t_d_areainfo` VALUES (511824, '石棉县', 3, 00000511800);
INSERT INTO `ch_t_d_areainfo` VALUES (511825, '天全县', 3, 00000511800);
INSERT INTO `ch_t_d_areainfo` VALUES (511826, '芦山县', 3, 00000511800);
INSERT INTO `ch_t_d_areainfo` VALUES (511827, '宝兴县', 3, 00000511800);
INSERT INTO `ch_t_d_areainfo` VALUES (511900, '巴中市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (511902, '巴州区', 3, 00000511900);
INSERT INTO `ch_t_d_areainfo` VALUES (511903, '恩阳区', 3, 00000511900);
INSERT INTO `ch_t_d_areainfo` VALUES (511921, '通江县', 3, 00000511900);
INSERT INTO `ch_t_d_areainfo` VALUES (511922, '南江县', 3, 00000511900);
INSERT INTO `ch_t_d_areainfo` VALUES (511923, '平昌县', 3, 00000511900);
INSERT INTO `ch_t_d_areainfo` VALUES (512000, '资阳市', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (512002, '雁江区', 3, 00000512000);
INSERT INTO `ch_t_d_areainfo` VALUES (512021, '安岳县', 3, 00000512000);
INSERT INTO `ch_t_d_areainfo` VALUES (512022, '乐至县', 3, 00000512000);
INSERT INTO `ch_t_d_areainfo` VALUES (513200, '阿坝藏族羌族自治州', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (513201, '马尔康市', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513221, '汶川县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513222, '理县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513223, '茂县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513224, '松潘县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513225, '九寨沟县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513226, '金川县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513227, '小金县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513228, '黑水县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513230, '壤塘县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513231, '阿坝县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513232, '若尔盖县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513233, '红原县', 3, 00000513200);
INSERT INTO `ch_t_d_areainfo` VALUES (513300, '甘孜藏族自治州', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (513301, '康定市', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513322, '泸定县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513323, '丹巴县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513324, '九龙县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513325, '雅江县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513326, '道孚县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513327, '炉霍县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513328, '甘孜县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513329, '新龙县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513330, '德格县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513331, '白玉县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513332, '石渠县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513333, '色达县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513334, '理塘县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513335, '巴塘县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513336, '乡城县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513337, '稻城县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513338, '得荣县', 3, 00000513300);
INSERT INTO `ch_t_d_areainfo` VALUES (513400, '凉山彝族自治州', 2, 00000510000);
INSERT INTO `ch_t_d_areainfo` VALUES (513401, '西昌市', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513422, '木里藏族自治县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513423, '盐源县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513424, '德昌县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513425, '会理县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513426, '会东县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513427, '宁南县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513428, '普格县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513429, '布拖县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513430, '金阳县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513431, '昭觉县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513432, '喜德县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513433, '冕宁县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513434, '越西县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513435, '甘洛县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513436, '美姑县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (513437, '雷波县', 3, 00000513400);
INSERT INTO `ch_t_d_areainfo` VALUES (520000, '贵州省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (520100, '贵阳市', 2, 00000520000);
INSERT INTO `ch_t_d_areainfo` VALUES (520102, '南明区', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520103, '云岩区', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520111, '花溪区', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520112, '乌当区', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520113, '白云区', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520115, '观山湖区', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520121, '开阳县', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520122, '息烽县', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520123, '修文县', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520181, '清镇市', 3, 00000520100);
INSERT INTO `ch_t_d_areainfo` VALUES (520200, '六盘水市', 2, 00000520000);
INSERT INTO `ch_t_d_areainfo` VALUES (520201, '钟山区', 3, 00000520200);
INSERT INTO `ch_t_d_areainfo` VALUES (520203, '六枝特区', 3, 00000520200);
INSERT INTO `ch_t_d_areainfo` VALUES (520221, '水城县', 3, 00000520200);
INSERT INTO `ch_t_d_areainfo` VALUES (520222, '盘县', 3, 00000520200);
INSERT INTO `ch_t_d_areainfo` VALUES (520300, '遵义市', 2, 00000520000);
INSERT INTO `ch_t_d_areainfo` VALUES (520302, '红花岗区', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520303, '汇川区', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520304, '播州区', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520322, '桐梓县', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520323, '绥阳县', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520324, '正安县', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520325, '道真仡佬族苗族自治县', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520326, '务川仡佬族苗族自治县', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520327, '凤冈县', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520328, '湄潭县', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520329, '余庆县', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520330, '习水县', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520381, '赤水市', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520382, '仁怀市', 3, 00000520300);
INSERT INTO `ch_t_d_areainfo` VALUES (520400, '安顺市', 2, 00000520000);
INSERT INTO `ch_t_d_areainfo` VALUES (520402, '西秀区', 3, 00000520400);
INSERT INTO `ch_t_d_areainfo` VALUES (520403, '平坝区', 3, 00000520400);
INSERT INTO `ch_t_d_areainfo` VALUES (520422, '普定县', 3, 00000520400);
INSERT INTO `ch_t_d_areainfo` VALUES (520423, '镇宁布依族苗族自治县', 3, 00000520400);
INSERT INTO `ch_t_d_areainfo` VALUES (520424, '关岭布依族苗族自治县', 3, 00000520400);
INSERT INTO `ch_t_d_areainfo` VALUES (520425, '紫云苗族布依族自治县', 3, 00000520400);
INSERT INTO `ch_t_d_areainfo` VALUES (520500, '毕节市', 2, 00000520000);
INSERT INTO `ch_t_d_areainfo` VALUES (520502, '七星关区', 3, 00000520500);
INSERT INTO `ch_t_d_areainfo` VALUES (520521, '大方县', 3, 00000520500);
INSERT INTO `ch_t_d_areainfo` VALUES (520522, '黔西县', 3, 00000520500);
INSERT INTO `ch_t_d_areainfo` VALUES (520523, '金沙县', 3, 00000520500);
INSERT INTO `ch_t_d_areainfo` VALUES (520524, '织金县', 3, 00000520500);
INSERT INTO `ch_t_d_areainfo` VALUES (520525, '纳雍县', 3, 00000520500);
INSERT INTO `ch_t_d_areainfo` VALUES (520526, '威宁彝族回族苗族自治县', 3, 00000520500);
INSERT INTO `ch_t_d_areainfo` VALUES (520527, '赫章县', 3, 00000520500);
INSERT INTO `ch_t_d_areainfo` VALUES (520600, '铜仁市', 2, 00000520000);
INSERT INTO `ch_t_d_areainfo` VALUES (520602, '碧江区', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (520603, '万山区', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (520621, '江口县', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (520622, '玉屏侗族自治县', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (520623, '石阡县', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (520624, '思南县', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (520625, '印江土家族苗族自治县', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (520626, '德江县', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (520627, '沿河土家族自治县', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (520628, '松桃苗族自治县', 3, 00000520600);
INSERT INTO `ch_t_d_areainfo` VALUES (522300, '黔西南布依族苗族自治州', 2, 00000520000);
INSERT INTO `ch_t_d_areainfo` VALUES (522301, '兴义市', 3, 00000522300);
INSERT INTO `ch_t_d_areainfo` VALUES (522322, '兴仁县', 3, 00000522300);
INSERT INTO `ch_t_d_areainfo` VALUES (522323, '普安县', 3, 00000522300);
INSERT INTO `ch_t_d_areainfo` VALUES (522324, '晴隆县', 3, 00000522300);
INSERT INTO `ch_t_d_areainfo` VALUES (522325, '贞丰县', 3, 00000522300);
INSERT INTO `ch_t_d_areainfo` VALUES (522326, '望谟县', 3, 00000522300);
INSERT INTO `ch_t_d_areainfo` VALUES (522327, '册亨县', 3, 00000522300);
INSERT INTO `ch_t_d_areainfo` VALUES (522328, '安龙县', 3, 00000522300);
INSERT INTO `ch_t_d_areainfo` VALUES (522600, '黔东南苗族侗族自治州', 2, 00000520000);
INSERT INTO `ch_t_d_areainfo` VALUES (522601, '凯里市', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522622, '黄平县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522623, '施秉县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522624, '三穗县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522625, '镇远县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522626, '岑巩县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522627, '天柱县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522628, '锦屏县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522629, '剑河县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522630, '台江县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522631, '黎平县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522632, '榕江县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522633, '从江县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522634, '雷山县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522635, '麻江县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522636, '丹寨县', 3, 00000522600);
INSERT INTO `ch_t_d_areainfo` VALUES (522700, '黔南布依族苗族自治州', 2, 00000520000);
INSERT INTO `ch_t_d_areainfo` VALUES (522701, '都匀市', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522702, '福泉市', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522722, '荔波县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522723, '贵定县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522725, '瓮安县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522726, '独山县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522727, '平塘县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522728, '罗甸县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522729, '长顺县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522730, '龙里县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522731, '惠水县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (522732, '三都水族自治县', 3, 00000522700);
INSERT INTO `ch_t_d_areainfo` VALUES (530000, '云南省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (530100, '昆明市', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (530102, '五华区', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530103, '盘龙区', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530111, '官渡区', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530112, '西山区', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530113, '东川区', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530114, '呈贡区', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530122, '晋宁县', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530124, '富民县', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530125, '宜良县', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530126, '石林彝族自治县', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530127, '嵩明县', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530128, '禄劝彝族苗族自治县', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530129, '寻甸回族彝族自治县', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530181, '安宁市', 3, 00000530100);
INSERT INTO `ch_t_d_areainfo` VALUES (530300, '曲靖市', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (530302, '麒麟区', 3, 00000530300);
INSERT INTO `ch_t_d_areainfo` VALUES (530303, '沾益区', 3, 00000530300);
INSERT INTO `ch_t_d_areainfo` VALUES (530321, '马龙县', 3, 00000530300);
INSERT INTO `ch_t_d_areainfo` VALUES (530322, '陆良县', 3, 00000530300);
INSERT INTO `ch_t_d_areainfo` VALUES (530323, '师宗县', 3, 00000530300);
INSERT INTO `ch_t_d_areainfo` VALUES (530324, '罗平县', 3, 00000530300);
INSERT INTO `ch_t_d_areainfo` VALUES (530325, '富源县', 3, 00000530300);
INSERT INTO `ch_t_d_areainfo` VALUES (530326, '会泽县', 3, 00000530300);
INSERT INTO `ch_t_d_areainfo` VALUES (530381, '宣威市', 3, 00000530300);
INSERT INTO `ch_t_d_areainfo` VALUES (530400, '玉溪市', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (530402, '红塔区', 3, 00000530400);
INSERT INTO `ch_t_d_areainfo` VALUES (530403, '江川区', 3, 00000530400);
INSERT INTO `ch_t_d_areainfo` VALUES (530422, '澄江县', 3, 00000530400);
INSERT INTO `ch_t_d_areainfo` VALUES (530423, '通海县', 3, 00000530400);
INSERT INTO `ch_t_d_areainfo` VALUES (530424, '华宁县', 3, 00000530400);
INSERT INTO `ch_t_d_areainfo` VALUES (530425, '易门县', 3, 00000530400);
INSERT INTO `ch_t_d_areainfo` VALUES (530426, '峨山彝族自治县', 3, 00000530400);
INSERT INTO `ch_t_d_areainfo` VALUES (530427, '新平彝族傣族自治县', 3, 00000530400);
INSERT INTO `ch_t_d_areainfo` VALUES (530428, '元江哈尼族彝族傣族自治县', 3, 00000530400);
INSERT INTO `ch_t_d_areainfo` VALUES (530500, '保山市', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (530502, '隆阳区', 3, 00000530500);
INSERT INTO `ch_t_d_areainfo` VALUES (530521, '施甸县', 3, 00000530500);
INSERT INTO `ch_t_d_areainfo` VALUES (530523, '龙陵县', 3, 00000530500);
INSERT INTO `ch_t_d_areainfo` VALUES (530524, '昌宁县', 3, 00000530500);
INSERT INTO `ch_t_d_areainfo` VALUES (530581, '腾冲市', 3, 00000530500);
INSERT INTO `ch_t_d_areainfo` VALUES (530600, '昭通市', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (530602, '昭阳区', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530621, '鲁甸县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530622, '巧家县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530623, '盐津县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530624, '大关县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530625, '永善县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530626, '绥江县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530627, '镇雄县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530628, '彝良县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530629, '威信县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530630, '水富县', 3, 00000530600);
INSERT INTO `ch_t_d_areainfo` VALUES (530700, '丽江市', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (530702, '古城区', 3, 00000530700);
INSERT INTO `ch_t_d_areainfo` VALUES (530721, '玉龙纳西族自治县', 3, 00000530700);
INSERT INTO `ch_t_d_areainfo` VALUES (530722, '永胜县', 3, 00000530700);
INSERT INTO `ch_t_d_areainfo` VALUES (530723, '华坪县', 3, 00000530700);
INSERT INTO `ch_t_d_areainfo` VALUES (530724, '宁蒗彝族自治县', 3, 00000530700);
INSERT INTO `ch_t_d_areainfo` VALUES (530800, '普洱市', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (530802, '思茅区', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530821, '宁洱哈尼族彝族自治县', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530822, '墨江哈尼族自治县', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530823, '景东彝族自治县', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530824, '景谷傣族彝族自治县', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530825, '镇沅彝族哈尼族拉祜族自治县', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530826, '江城哈尼族彝族自治县', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530827, '孟连傣族拉祜族佤族自治县', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530828, '澜沧拉祜族自治县', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530829, '西盟佤族自治县', 3, 00000530800);
INSERT INTO `ch_t_d_areainfo` VALUES (530900, '临沧市', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (530902, '临翔区', 3, 00000530900);
INSERT INTO `ch_t_d_areainfo` VALUES (530921, '凤庆县', 3, 00000530900);
INSERT INTO `ch_t_d_areainfo` VALUES (530922, '云县', 3, 00000530900);
INSERT INTO `ch_t_d_areainfo` VALUES (530923, '永德县', 3, 00000530900);
INSERT INTO `ch_t_d_areainfo` VALUES (530924, '镇康县', 3, 00000530900);
INSERT INTO `ch_t_d_areainfo` VALUES (530925, '双江拉祜族佤族布朗族傣族自治县', 3, 00000530900);
INSERT INTO `ch_t_d_areainfo` VALUES (530926, '耿马傣族佤族自治县', 3, 00000530900);
INSERT INTO `ch_t_d_areainfo` VALUES (530927, '沧源佤族自治县', 3, 00000530900);
INSERT INTO `ch_t_d_areainfo` VALUES (532300, '楚雄彝族自治州', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (532301, '楚雄市', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532322, '双柏县', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532323, '牟定县', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532324, '南华县', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532325, '姚安县', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532326, '大姚县', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532327, '永仁县', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532328, '元谋县', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532329, '武定县', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532331, '禄丰县', 3, 00000532300);
INSERT INTO `ch_t_d_areainfo` VALUES (532500, '红河哈尼族彝族自治州', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (532501, '个旧市', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532502, '开远市', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532503, '蒙自市', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532504, '弥勒市', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532523, '屏边苗族自治县', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532524, '建水县', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532525, '石屏县', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532527, '泸西县', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532528, '元阳县', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532529, '红河县', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532530, '金平苗族瑶族傣族自治县', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532531, '绿春县', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532532, '河口瑶族自治县', 3, 00000532500);
INSERT INTO `ch_t_d_areainfo` VALUES (532600, '文山壮族苗族自治州', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (532601, '文山市', 3, 00000532600);
INSERT INTO `ch_t_d_areainfo` VALUES (532622, '砚山县', 3, 00000532600);
INSERT INTO `ch_t_d_areainfo` VALUES (532623, '西畴县', 3, 00000532600);
INSERT INTO `ch_t_d_areainfo` VALUES (532624, '麻栗坡县', 3, 00000532600);
INSERT INTO `ch_t_d_areainfo` VALUES (532625, '马关县', 3, 00000532600);
INSERT INTO `ch_t_d_areainfo` VALUES (532626, '丘北县', 3, 00000532600);
INSERT INTO `ch_t_d_areainfo` VALUES (532627, '广南县', 3, 00000532600);
INSERT INTO `ch_t_d_areainfo` VALUES (532628, '富宁县', 3, 00000532600);
INSERT INTO `ch_t_d_areainfo` VALUES (532800, '西双版纳傣族自治州', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (532801, '景洪市', 3, 00000532800);
INSERT INTO `ch_t_d_areainfo` VALUES (532822, '勐海县', 3, 00000532800);
INSERT INTO `ch_t_d_areainfo` VALUES (532823, '勐腊县', 3, 00000532800);
INSERT INTO `ch_t_d_areainfo` VALUES (532900, '大理白族自治州', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (532901, '大理市', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532922, '漾濞彝族自治县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532923, '祥云县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532924, '宾川县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532925, '弥渡县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532926, '南涧彝族自治县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532927, '巍山彝族回族自治县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532928, '永平县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532929, '云龙县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532930, '洱源县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532931, '剑川县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (532932, '鹤庆县', 3, 00000532900);
INSERT INTO `ch_t_d_areainfo` VALUES (533100, '德宏傣族景颇族自治州', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (533102, '瑞丽市', 3, 00000533100);
INSERT INTO `ch_t_d_areainfo` VALUES (533103, '芒市', 3, 00000533100);
INSERT INTO `ch_t_d_areainfo` VALUES (533122, '梁河县', 3, 00000533100);
INSERT INTO `ch_t_d_areainfo` VALUES (533123, '盈江县', 3, 00000533100);
INSERT INTO `ch_t_d_areainfo` VALUES (533124, '陇川县', 3, 00000533100);
INSERT INTO `ch_t_d_areainfo` VALUES (533300, '怒江傈僳族自治州', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (533301, '泸水市', 3, 00000533300);
INSERT INTO `ch_t_d_areainfo` VALUES (533323, '福贡县', 3, 00000533300);
INSERT INTO `ch_t_d_areainfo` VALUES (533324, '贡山独龙族怒族自治县', 3, 00000533300);
INSERT INTO `ch_t_d_areainfo` VALUES (533325, '兰坪白族普米族自治县', 3, 00000533300);
INSERT INTO `ch_t_d_areainfo` VALUES (533400, '迪庆藏族自治州', 2, 00000530000);
INSERT INTO `ch_t_d_areainfo` VALUES (533401, '香格里拉市', 3, 00000533400);
INSERT INTO `ch_t_d_areainfo` VALUES (533422, '德钦县', 3, 00000533400);
INSERT INTO `ch_t_d_areainfo` VALUES (533423, '维西傈僳族自治县', 3, 00000533400);
INSERT INTO `ch_t_d_areainfo` VALUES (540000, '西藏自治区', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (540100, '拉萨市', 2, 00000540000);
INSERT INTO `ch_t_d_areainfo` VALUES (540102, '城关区', 3, 00000540100);
INSERT INTO `ch_t_d_areainfo` VALUES (540103, '堆龙德庆区', 3, 00000540100);
INSERT INTO `ch_t_d_areainfo` VALUES (540121, '林周县', 3, 00000540100);
INSERT INTO `ch_t_d_areainfo` VALUES (540122, '当雄县', 3, 00000540100);
INSERT INTO `ch_t_d_areainfo` VALUES (540123, '尼木县', 3, 00000540100);
INSERT INTO `ch_t_d_areainfo` VALUES (540124, '曲水县', 3, 00000540100);
INSERT INTO `ch_t_d_areainfo` VALUES (540126, '达孜县', 3, 00000540100);
INSERT INTO `ch_t_d_areainfo` VALUES (540127, '墨竹工卡县', 3, 00000540100);
INSERT INTO `ch_t_d_areainfo` VALUES (540200, '日喀则市', 2, 00000540000);
INSERT INTO `ch_t_d_areainfo` VALUES (540202, '桑珠孜区', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540221, '南木林县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540222, '江孜县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540223, '定日县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540224, '萨迦县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540225, '拉孜县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540226, '昂仁县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540227, '谢通门县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540228, '白朗县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540229, '仁布县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540230, '康马县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540231, '定结县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540232, '仲巴县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540233, '亚东县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540234, '吉隆县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540235, '聂拉木县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540236, '萨嘎县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540237, '岗巴县', 3, 00000540200);
INSERT INTO `ch_t_d_areainfo` VALUES (540300, '昌都市', 2, 00000540000);
INSERT INTO `ch_t_d_areainfo` VALUES (540302, '卡若区', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540321, '江达县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540322, '贡觉县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540323, '类乌齐县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540324, '丁青县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540325, '察雅县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540326, '八宿县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540327, '左贡县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540328, '芒康县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540329, '洛隆县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540330, '边坝县', 3, 00000540300);
INSERT INTO `ch_t_d_areainfo` VALUES (540400, '林芝市', 2, 00000540000);
INSERT INTO `ch_t_d_areainfo` VALUES (540402, '巴宜区', 3, 00000540400);
INSERT INTO `ch_t_d_areainfo` VALUES (540421, '工布江达县', 3, 00000540400);
INSERT INTO `ch_t_d_areainfo` VALUES (540422, '米林县', 3, 00000540400);
INSERT INTO `ch_t_d_areainfo` VALUES (540423, '墨脱县', 3, 00000540400);
INSERT INTO `ch_t_d_areainfo` VALUES (540424, '波密县', 3, 00000540400);
INSERT INTO `ch_t_d_areainfo` VALUES (540425, '察隅县', 3, 00000540400);
INSERT INTO `ch_t_d_areainfo` VALUES (540426, '朗县', 3, 00000540400);
INSERT INTO `ch_t_d_areainfo` VALUES (540500, '山南市', 2, 00000540000);
INSERT INTO `ch_t_d_areainfo` VALUES (540502, '乃东区', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540521, '扎囊县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540522, '贡嘎县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540523, '桑日县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540524, '琼结县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540525, '曲松县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540526, '措美县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540527, '洛扎县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540528, '加查县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540529, '隆子县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540530, '错那县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (540531, '浪卡子县', 3, 00000540500);
INSERT INTO `ch_t_d_areainfo` VALUES (542400, '那曲地区', 2, 00000540000);
INSERT INTO `ch_t_d_areainfo` VALUES (542421, '那曲县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542422, '嘉黎县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542423, '比如县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542424, '聂荣县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542425, '安多县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542426, '申扎县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542427, '索县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542428, '班戈县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542429, '巴青县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542430, '尼玛县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542431, '双湖县', 3, 00000542400);
INSERT INTO `ch_t_d_areainfo` VALUES (542500, '阿里地区', 2, 00000540000);
INSERT INTO `ch_t_d_areainfo` VALUES (542521, '普兰县', 3, 00000542500);
INSERT INTO `ch_t_d_areainfo` VALUES (542522, '札达县', 3, 00000542500);
INSERT INTO `ch_t_d_areainfo` VALUES (542523, '噶尔县', 3, 00000542500);
INSERT INTO `ch_t_d_areainfo` VALUES (542524, '日土县', 3, 00000542500);
INSERT INTO `ch_t_d_areainfo` VALUES (542525, '革吉县', 3, 00000542500);
INSERT INTO `ch_t_d_areainfo` VALUES (542526, '改则县', 3, 00000542500);
INSERT INTO `ch_t_d_areainfo` VALUES (542527, '措勤县', 3, 00000542500);
INSERT INTO `ch_t_d_areainfo` VALUES (610000, '陕西省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (610100, '西安市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (610102, '新城区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610103, '碑林区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610104, '莲湖区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610111, '灞桥区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610112, '未央区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610113, '雁塔区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610114, '阎良区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610115, '临潼区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610116, '长安区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610117, '高陵区', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610122, '蓝田县', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610124, '周至县', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610125, '户县', 3, 00000610100);
INSERT INTO `ch_t_d_areainfo` VALUES (610200, '铜川市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (610202, '王益区', 3, 00000610200);
INSERT INTO `ch_t_d_areainfo` VALUES (610203, '印台区', 3, 00000610200);
INSERT INTO `ch_t_d_areainfo` VALUES (610204, '耀州区', 3, 00000610200);
INSERT INTO `ch_t_d_areainfo` VALUES (610222, '宜君县', 3, 00000610200);
INSERT INTO `ch_t_d_areainfo` VALUES (610300, '宝鸡市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (610302, '渭滨区', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610303, '金台区', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610304, '陈仓区', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610322, '凤翔县', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610323, '岐山县', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610324, '扶风县', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610326, '眉县', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610327, '陇县', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610328, '千阳县', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610329, '麟游县', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610330, '凤县', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610331, '太白县', 3, 00000610300);
INSERT INTO `ch_t_d_areainfo` VALUES (610400, '咸阳市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (610402, '秦都区', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610403, '杨陵区', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610404, '渭城区', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610422, '三原县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610423, '泾阳县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610424, '乾县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610425, '礼泉县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610426, '永寿县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610427, '彬县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610428, '长武县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610429, '旬邑县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610430, '淳化县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610431, '武功县', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610481, '兴平市', 3, 00000610400);
INSERT INTO `ch_t_d_areainfo` VALUES (610500, '渭南市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (610502, '临渭区', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610503, '华州区', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610522, '潼关县', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610523, '大荔县', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610524, '合阳县', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610525, '澄城县', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610526, '蒲城县', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610527, '白水县', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610528, '富平县', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610581, '韩城市', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610582, '华阴市', 3, 00000610500);
INSERT INTO `ch_t_d_areainfo` VALUES (610600, '延安市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (610602, '宝塔区', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610603, '安塞区', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610621, '延长县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610622, '延川县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610623, '子长县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610625, '志丹县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610626, '吴起县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610627, '甘泉县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610628, '富县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610629, '洛川县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610630, '宜川县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610631, '黄龙县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610632, '黄陵县', 3, 00000610600);
INSERT INTO `ch_t_d_areainfo` VALUES (610700, '汉中市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (610702, '汉台区', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610721, '南郑县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610722, '城固县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610723, '洋县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610724, '西乡县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610725, '勉县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610726, '宁强县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610727, '略阳县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610728, '镇巴县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610729, '留坝县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610730, '佛坪县', 3, 00000610700);
INSERT INTO `ch_t_d_areainfo` VALUES (610800, '榆林市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (610802, '榆阳区', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610803, '横山区', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610821, '神木县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610822, '府谷县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610824, '靖边县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610825, '定边县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610826, '绥德县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610827, '米脂县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610828, '佳县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610829, '吴堡县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610830, '清涧县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610831, '子洲县', 3, 00000610800);
INSERT INTO `ch_t_d_areainfo` VALUES (610900, '安康市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (610902, '汉滨区', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (610921, '汉阴县', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (610922, '石泉县', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (610923, '宁陕县', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (610924, '紫阳县', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (610925, '岚皋县', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (610926, '平利县', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (610927, '镇坪县', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (610928, '旬阳县', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (610929, '白河县', 3, 00000610900);
INSERT INTO `ch_t_d_areainfo` VALUES (611000, '商洛市', 2, 00000610000);
INSERT INTO `ch_t_d_areainfo` VALUES (611002, '商州区', 3, 00000611000);
INSERT INTO `ch_t_d_areainfo` VALUES (611021, '洛南县', 3, 00000611000);
INSERT INTO `ch_t_d_areainfo` VALUES (611022, '丹凤县', 3, 00000611000);
INSERT INTO `ch_t_d_areainfo` VALUES (611023, '商南县', 3, 00000611000);
INSERT INTO `ch_t_d_areainfo` VALUES (611024, '山阳县', 3, 00000611000);
INSERT INTO `ch_t_d_areainfo` VALUES (611025, '镇安县', 3, 00000611000);
INSERT INTO `ch_t_d_areainfo` VALUES (611026, '柞水县', 3, 00000611000);
INSERT INTO `ch_t_d_areainfo` VALUES (620000, '甘肃省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (620100, '兰州市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (620102, '城关区', 3, 00000620100);
INSERT INTO `ch_t_d_areainfo` VALUES (620103, '七里河区', 3, 00000620100);
INSERT INTO `ch_t_d_areainfo` VALUES (620104, '西固区', 3, 00000620100);
INSERT INTO `ch_t_d_areainfo` VALUES (620105, '安宁区', 3, 00000620100);
INSERT INTO `ch_t_d_areainfo` VALUES (620111, '红古区', 3, 00000620100);
INSERT INTO `ch_t_d_areainfo` VALUES (620121, '永登县', 3, 00000620100);
INSERT INTO `ch_t_d_areainfo` VALUES (620122, '皋兰县', 3, 00000620100);
INSERT INTO `ch_t_d_areainfo` VALUES (620123, '榆中县', 3, 00000620100);
INSERT INTO `ch_t_d_areainfo` VALUES (620200, '嘉峪关市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (620300, '金昌市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (620302, '金川区', 3, 00000620300);
INSERT INTO `ch_t_d_areainfo` VALUES (620321, '永昌县', 3, 00000620300);
INSERT INTO `ch_t_d_areainfo` VALUES (620400, '白银市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (620402, '白银区', 3, 00000620400);
INSERT INTO `ch_t_d_areainfo` VALUES (620403, '平川区', 3, 00000620400);
INSERT INTO `ch_t_d_areainfo` VALUES (620421, '靖远县', 3, 00000620400);
INSERT INTO `ch_t_d_areainfo` VALUES (620422, '会宁县', 3, 00000620400);
INSERT INTO `ch_t_d_areainfo` VALUES (620423, '景泰县', 3, 00000620400);
INSERT INTO `ch_t_d_areainfo` VALUES (620500, '天水市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (620502, '秦州区', 3, 00000620500);
INSERT INTO `ch_t_d_areainfo` VALUES (620503, '麦积区', 3, 00000620500);
INSERT INTO `ch_t_d_areainfo` VALUES (620521, '清水县', 3, 00000620500);
INSERT INTO `ch_t_d_areainfo` VALUES (620522, '秦安县', 3, 00000620500);
INSERT INTO `ch_t_d_areainfo` VALUES (620523, '甘谷县', 3, 00000620500);
INSERT INTO `ch_t_d_areainfo` VALUES (620524, '武山县', 3, 00000620500);
INSERT INTO `ch_t_d_areainfo` VALUES (620525, '张家川回族自治县', 3, 00000620500);
INSERT INTO `ch_t_d_areainfo` VALUES (620600, '武威市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (620602, '凉州区', 3, 00000620600);
INSERT INTO `ch_t_d_areainfo` VALUES (620621, '民勤县', 3, 00000620600);
INSERT INTO `ch_t_d_areainfo` VALUES (620622, '古浪县', 3, 00000620600);
INSERT INTO `ch_t_d_areainfo` VALUES (620623, '天祝藏族自治县', 3, 00000620600);
INSERT INTO `ch_t_d_areainfo` VALUES (620700, '张掖市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (620702, '甘州区', 3, 00000620700);
INSERT INTO `ch_t_d_areainfo` VALUES (620721, '肃南裕固族自治县', 3, 00000620700);
INSERT INTO `ch_t_d_areainfo` VALUES (620722, '民乐县', 3, 00000620700);
INSERT INTO `ch_t_d_areainfo` VALUES (620723, '临泽县', 3, 00000620700);
INSERT INTO `ch_t_d_areainfo` VALUES (620724, '高台县', 3, 00000620700);
INSERT INTO `ch_t_d_areainfo` VALUES (620725, '山丹县', 3, 00000620700);
INSERT INTO `ch_t_d_areainfo` VALUES (620800, '平凉市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (620802, '崆峒区', 3, 00000620800);
INSERT INTO `ch_t_d_areainfo` VALUES (620821, '泾川县', 3, 00000620800);
INSERT INTO `ch_t_d_areainfo` VALUES (620822, '灵台县', 3, 00000620800);
INSERT INTO `ch_t_d_areainfo` VALUES (620823, '崇信县', 3, 00000620800);
INSERT INTO `ch_t_d_areainfo` VALUES (620824, '华亭县', 3, 00000620800);
INSERT INTO `ch_t_d_areainfo` VALUES (620825, '庄浪县', 3, 00000620800);
INSERT INTO `ch_t_d_areainfo` VALUES (620826, '静宁县', 3, 00000620800);
INSERT INTO `ch_t_d_areainfo` VALUES (620900, '酒泉市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (620902, '肃州区', 3, 00000620900);
INSERT INTO `ch_t_d_areainfo` VALUES (620921, '金塔县', 3, 00000620900);
INSERT INTO `ch_t_d_areainfo` VALUES (620922, '瓜州县', 3, 00000620900);
INSERT INTO `ch_t_d_areainfo` VALUES (620923, '肃北蒙古族自治县', 3, 00000620900);
INSERT INTO `ch_t_d_areainfo` VALUES (620924, '阿克塞哈萨克族自治县', 3, 00000620900);
INSERT INTO `ch_t_d_areainfo` VALUES (620981, '玉门市', 3, 00000620900);
INSERT INTO `ch_t_d_areainfo` VALUES (620982, '敦煌市', 3, 00000620900);
INSERT INTO `ch_t_d_areainfo` VALUES (621000, '庆阳市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (621002, '西峰区', 3, 00000621000);
INSERT INTO `ch_t_d_areainfo` VALUES (621021, '庆城县', 3, 00000621000);
INSERT INTO `ch_t_d_areainfo` VALUES (621022, '环县', 3, 00000621000);
INSERT INTO `ch_t_d_areainfo` VALUES (621023, '华池县', 3, 00000621000);
INSERT INTO `ch_t_d_areainfo` VALUES (621024, '合水县', 3, 00000621000);
INSERT INTO `ch_t_d_areainfo` VALUES (621025, '正宁县', 3, 00000621000);
INSERT INTO `ch_t_d_areainfo` VALUES (621026, '宁县', 3, 00000621000);
INSERT INTO `ch_t_d_areainfo` VALUES (621027, '镇原县', 3, 00000621000);
INSERT INTO `ch_t_d_areainfo` VALUES (621100, '定西市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (621102, '安定区', 3, 00000621100);
INSERT INTO `ch_t_d_areainfo` VALUES (621121, '通渭县', 3, 00000621100);
INSERT INTO `ch_t_d_areainfo` VALUES (621122, '陇西县', 3, 00000621100);
INSERT INTO `ch_t_d_areainfo` VALUES (621123, '渭源县', 3, 00000621100);
INSERT INTO `ch_t_d_areainfo` VALUES (621124, '临洮县', 3, 00000621100);
INSERT INTO `ch_t_d_areainfo` VALUES (621125, '漳县', 3, 00000621100);
INSERT INTO `ch_t_d_areainfo` VALUES (621126, '岷县', 3, 00000621100);
INSERT INTO `ch_t_d_areainfo` VALUES (621200, '陇南市', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (621202, '武都区', 3, 00000621200);
INSERT INTO `ch_t_d_areainfo` VALUES (621221, '成县', 3, 00000621200);
INSERT INTO `ch_t_d_areainfo` VALUES (621222, '文县', 3, 00000621200);
INSERT INTO `ch_t_d_areainfo` VALUES (621223, '宕昌县', 3, 00000621200);
INSERT INTO `ch_t_d_areainfo` VALUES (621224, '康县', 3, 00000621200);
INSERT INTO `ch_t_d_areainfo` VALUES (621225, '西和县', 3, 00000621200);
INSERT INTO `ch_t_d_areainfo` VALUES (621226, '礼县', 3, 00000621200);
INSERT INTO `ch_t_d_areainfo` VALUES (621227, '徽县', 3, 00000621200);
INSERT INTO `ch_t_d_areainfo` VALUES (621228, '两当县', 3, 00000621200);
INSERT INTO `ch_t_d_areainfo` VALUES (622900, '临夏回族自治州', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (622901, '临夏市', 3, 00000622900);
INSERT INTO `ch_t_d_areainfo` VALUES (622921, '临夏县', 3, 00000622900);
INSERT INTO `ch_t_d_areainfo` VALUES (622922, '康乐县', 3, 00000622900);
INSERT INTO `ch_t_d_areainfo` VALUES (622923, '永靖县', 3, 00000622900);
INSERT INTO `ch_t_d_areainfo` VALUES (622924, '广河县', 3, 00000622900);
INSERT INTO `ch_t_d_areainfo` VALUES (622925, '和政县', 3, 00000622900);
INSERT INTO `ch_t_d_areainfo` VALUES (622926, '东乡族自治县', 3, 00000622900);
INSERT INTO `ch_t_d_areainfo` VALUES (622927, '积石山保安族东乡族撒拉族自治县', 3, 00000622900);
INSERT INTO `ch_t_d_areainfo` VALUES (623000, '甘南藏族自治州', 2, 00000620000);
INSERT INTO `ch_t_d_areainfo` VALUES (623001, '合作市', 3, 00000623000);
INSERT INTO `ch_t_d_areainfo` VALUES (623021, '临潭县', 3, 00000623000);
INSERT INTO `ch_t_d_areainfo` VALUES (623022, '卓尼县', 3, 00000623000);
INSERT INTO `ch_t_d_areainfo` VALUES (623023, '舟曲县', 3, 00000623000);
INSERT INTO `ch_t_d_areainfo` VALUES (623024, '迭部县', 3, 00000623000);
INSERT INTO `ch_t_d_areainfo` VALUES (623025, '玛曲县', 3, 00000623000);
INSERT INTO `ch_t_d_areainfo` VALUES (623026, '碌曲县', 3, 00000623000);
INSERT INTO `ch_t_d_areainfo` VALUES (623027, '夏河县', 3, 00000623000);
INSERT INTO `ch_t_d_areainfo` VALUES (630000, '青海省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (630100, '西宁市', 2, 00000630000);
INSERT INTO `ch_t_d_areainfo` VALUES (630102, '城东区', 3, 00000630100);
INSERT INTO `ch_t_d_areainfo` VALUES (630103, '城中区', 3, 00000630100);
INSERT INTO `ch_t_d_areainfo` VALUES (630104, '城西区', 3, 00000630100);
INSERT INTO `ch_t_d_areainfo` VALUES (630105, '城北区', 3, 00000630100);
INSERT INTO `ch_t_d_areainfo` VALUES (630121, '大通回族土族自治县', 3, 00000630100);
INSERT INTO `ch_t_d_areainfo` VALUES (630122, '湟中县', 3, 00000630100);
INSERT INTO `ch_t_d_areainfo` VALUES (630123, '湟源县', 3, 00000630100);
INSERT INTO `ch_t_d_areainfo` VALUES (630200, '海东市', 2, 00000630000);
INSERT INTO `ch_t_d_areainfo` VALUES (630202, '乐都区', 3, 00000630200);
INSERT INTO `ch_t_d_areainfo` VALUES (630203, '平安区', 3, 00000630200);
INSERT INTO `ch_t_d_areainfo` VALUES (630222, '民和回族土族自治县', 3, 00000630200);
INSERT INTO `ch_t_d_areainfo` VALUES (630223, '互助土族自治县', 3, 00000630200);
INSERT INTO `ch_t_d_areainfo` VALUES (630224, '化隆回族自治县', 3, 00000630200);
INSERT INTO `ch_t_d_areainfo` VALUES (630225, '循化撒拉族自治县', 3, 00000630200);
INSERT INTO `ch_t_d_areainfo` VALUES (632200, '海北藏族自治州', 2, 00000630000);
INSERT INTO `ch_t_d_areainfo` VALUES (632221, '门源回族自治县', 3, 00000632200);
INSERT INTO `ch_t_d_areainfo` VALUES (632222, '祁连县', 3, 00000632200);
INSERT INTO `ch_t_d_areainfo` VALUES (632223, '海晏县', 3, 00000632200);
INSERT INTO `ch_t_d_areainfo` VALUES (632224, '刚察县', 3, 00000632200);
INSERT INTO `ch_t_d_areainfo` VALUES (632300, '黄南藏族自治州', 2, 00000630000);
INSERT INTO `ch_t_d_areainfo` VALUES (632321, '同仁县', 3, 00000632300);
INSERT INTO `ch_t_d_areainfo` VALUES (632322, '尖扎县', 3, 00000632300);
INSERT INTO `ch_t_d_areainfo` VALUES (632323, '泽库县', 3, 00000632300);
INSERT INTO `ch_t_d_areainfo` VALUES (632324, '河南蒙古族自治县', 3, 00000632300);
INSERT INTO `ch_t_d_areainfo` VALUES (632500, '海南藏族自治州', 2, 00000630000);
INSERT INTO `ch_t_d_areainfo` VALUES (632521, '共和县', 3, 00000632500);
INSERT INTO `ch_t_d_areainfo` VALUES (632522, '同德县', 3, 00000632500);
INSERT INTO `ch_t_d_areainfo` VALUES (632523, '贵德县', 3, 00000632500);
INSERT INTO `ch_t_d_areainfo` VALUES (632524, '兴海县', 3, 00000632500);
INSERT INTO `ch_t_d_areainfo` VALUES (632525, '贵南县', 3, 00000632500);
INSERT INTO `ch_t_d_areainfo` VALUES (632600, '果洛藏族自治州', 2, 00000630000);
INSERT INTO `ch_t_d_areainfo` VALUES (632621, '玛沁县', 3, 00000632600);
INSERT INTO `ch_t_d_areainfo` VALUES (632622, '班玛县', 3, 00000632600);
INSERT INTO `ch_t_d_areainfo` VALUES (632623, '甘德县', 3, 00000632600);
INSERT INTO `ch_t_d_areainfo` VALUES (632624, '达日县', 3, 00000632600);
INSERT INTO `ch_t_d_areainfo` VALUES (632625, '久治县', 3, 00000632600);
INSERT INTO `ch_t_d_areainfo` VALUES (632626, '玛多县', 3, 00000632600);
INSERT INTO `ch_t_d_areainfo` VALUES (632700, '玉树藏族自治州', 2, 00000630000);
INSERT INTO `ch_t_d_areainfo` VALUES (632701, '玉树市', 3, 00000632700);
INSERT INTO `ch_t_d_areainfo` VALUES (632722, '杂多县', 3, 00000632700);
INSERT INTO `ch_t_d_areainfo` VALUES (632723, '称多县', 3, 00000632700);
INSERT INTO `ch_t_d_areainfo` VALUES (632724, '治多县', 3, 00000632700);
INSERT INTO `ch_t_d_areainfo` VALUES (632725, '囊谦县', 3, 00000632700);
INSERT INTO `ch_t_d_areainfo` VALUES (632726, '曲麻莱县', 3, 00000632700);
INSERT INTO `ch_t_d_areainfo` VALUES (632800, '海西蒙古族藏族自治州', 2, 00000630000);
INSERT INTO `ch_t_d_areainfo` VALUES (632801, '格尔木市', 3, 00000632800);
INSERT INTO `ch_t_d_areainfo` VALUES (632802, '德令哈市', 3, 00000632800);
INSERT INTO `ch_t_d_areainfo` VALUES (632821, '乌兰县', 3, 00000632800);
INSERT INTO `ch_t_d_areainfo` VALUES (632822, '都兰县', 3, 00000632800);
INSERT INTO `ch_t_d_areainfo` VALUES (632823, '天峻县', 3, 00000632800);
INSERT INTO `ch_t_d_areainfo` VALUES (640000, '宁夏回族自治区', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (640100, '银川市', 2, 00000640000);
INSERT INTO `ch_t_d_areainfo` VALUES (640104, '兴庆区', 3, 00000640100);
INSERT INTO `ch_t_d_areainfo` VALUES (640105, '西夏区', 3, 00000640100);
INSERT INTO `ch_t_d_areainfo` VALUES (640106, '金凤区', 3, 00000640100);
INSERT INTO `ch_t_d_areainfo` VALUES (640121, '永宁县', 3, 00000640100);
INSERT INTO `ch_t_d_areainfo` VALUES (640122, '贺兰县', 3, 00000640100);
INSERT INTO `ch_t_d_areainfo` VALUES (640181, '灵武市', 3, 00000640100);
INSERT INTO `ch_t_d_areainfo` VALUES (640200, '石嘴山市', 2, 00000640000);
INSERT INTO `ch_t_d_areainfo` VALUES (640202, '大武口区', 3, 00000640200);
INSERT INTO `ch_t_d_areainfo` VALUES (640205, '惠农区', 3, 00000640200);
INSERT INTO `ch_t_d_areainfo` VALUES (640221, '平罗县', 3, 00000640200);
INSERT INTO `ch_t_d_areainfo` VALUES (640300, '吴忠市', 2, 00000640000);
INSERT INTO `ch_t_d_areainfo` VALUES (640302, '利通区', 3, 00000640300);
INSERT INTO `ch_t_d_areainfo` VALUES (640303, '红寺堡区', 3, 00000640300);
INSERT INTO `ch_t_d_areainfo` VALUES (640323, '盐池县', 3, 00000640300);
INSERT INTO `ch_t_d_areainfo` VALUES (640324, '同心县', 3, 00000640300);
INSERT INTO `ch_t_d_areainfo` VALUES (640381, '青铜峡市', 3, 00000640300);
INSERT INTO `ch_t_d_areainfo` VALUES (640400, '固原市', 2, 00000640000);
INSERT INTO `ch_t_d_areainfo` VALUES (640402, '原州区', 3, 00000640400);
INSERT INTO `ch_t_d_areainfo` VALUES (640422, '西吉县', 3, 00000640400);
INSERT INTO `ch_t_d_areainfo` VALUES (640423, '隆德县', 3, 00000640400);
INSERT INTO `ch_t_d_areainfo` VALUES (640424, '泾源县', 3, 00000640400);
INSERT INTO `ch_t_d_areainfo` VALUES (640425, '彭阳县', 3, 00000640400);
INSERT INTO `ch_t_d_areainfo` VALUES (640500, '中卫市', 2, 00000640000);
INSERT INTO `ch_t_d_areainfo` VALUES (640502, '沙坡头区', 3, 00000640500);
INSERT INTO `ch_t_d_areainfo` VALUES (640521, '中宁县', 3, 00000640500);
INSERT INTO `ch_t_d_areainfo` VALUES (640522, '海原县', 3, 00000640500);
INSERT INTO `ch_t_d_areainfo` VALUES (650000, '新疆维吾尔自治区', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (650100, '乌鲁木齐市', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (650102, '天山区', 3, 00000650100);
INSERT INTO `ch_t_d_areainfo` VALUES (650103, '沙依巴克区', 3, 00000650100);
INSERT INTO `ch_t_d_areainfo` VALUES (650104, '新市区', 3, 00000650100);
INSERT INTO `ch_t_d_areainfo` VALUES (650105, '水磨沟区', 3, 00000650100);
INSERT INTO `ch_t_d_areainfo` VALUES (650106, '头屯河区', 3, 00000650100);
INSERT INTO `ch_t_d_areainfo` VALUES (650107, '达坂城区', 3, 00000650100);
INSERT INTO `ch_t_d_areainfo` VALUES (650109, '米东区', 3, 00000650100);
INSERT INTO `ch_t_d_areainfo` VALUES (650121, '乌鲁木齐县', 3, 00000650100);
INSERT INTO `ch_t_d_areainfo` VALUES (650200, '克拉玛依市', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (650202, '独山子区', 3, 00000650200);
INSERT INTO `ch_t_d_areainfo` VALUES (650203, '克拉玛依区', 3, 00000650200);
INSERT INTO `ch_t_d_areainfo` VALUES (650204, '白碱滩区', 3, 00000650200);
INSERT INTO `ch_t_d_areainfo` VALUES (650205, '乌尔禾区', 3, 00000650200);
INSERT INTO `ch_t_d_areainfo` VALUES (650400, '吐鲁番市', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (650402, '高昌区', 3, 00000650400);
INSERT INTO `ch_t_d_areainfo` VALUES (650421, '鄯善县', 3, 00000650400);
INSERT INTO `ch_t_d_areainfo` VALUES (650422, '托克逊县', 3, 00000650400);
INSERT INTO `ch_t_d_areainfo` VALUES (650500, '哈密市', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (650502, '伊州区', 3, 00000650500);
INSERT INTO `ch_t_d_areainfo` VALUES (650521, '巴里坤哈萨克自治县', 3, 00000650500);
INSERT INTO `ch_t_d_areainfo` VALUES (650522, '伊吾县', 3, 00000650500);
INSERT INTO `ch_t_d_areainfo` VALUES (652300, '昌吉回族自治州', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (652301, '昌吉市', 3, 00000652300);
INSERT INTO `ch_t_d_areainfo` VALUES (652302, '阜康市', 3, 00000652300);
INSERT INTO `ch_t_d_areainfo` VALUES (652323, '呼图壁县', 3, 00000652300);
INSERT INTO `ch_t_d_areainfo` VALUES (652324, '玛纳斯县', 3, 00000652300);
INSERT INTO `ch_t_d_areainfo` VALUES (652325, '奇台县', 3, 00000652300);
INSERT INTO `ch_t_d_areainfo` VALUES (652327, '吉木萨尔县', 3, 00000652300);
INSERT INTO `ch_t_d_areainfo` VALUES (652328, '木垒哈萨克自治县', 3, 00000652300);
INSERT INTO `ch_t_d_areainfo` VALUES (652700, '博尔塔拉蒙古自治州', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (652701, '博乐市', 3, 00000652700);
INSERT INTO `ch_t_d_areainfo` VALUES (652702, '阿拉山口市', 3, 00000652700);
INSERT INTO `ch_t_d_areainfo` VALUES (652722, '精河县', 3, 00000652700);
INSERT INTO `ch_t_d_areainfo` VALUES (652723, '温泉县', 3, 00000652700);
INSERT INTO `ch_t_d_areainfo` VALUES (652800, '巴音郭楞蒙古自治州', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (652801, '库尔勒市', 3, 00000652800);
INSERT INTO `ch_t_d_areainfo` VALUES (652822, '轮台县', 3, 00000652800);
INSERT INTO `ch_t_d_areainfo` VALUES (652823, '尉犁县', 3, 00000652800);
INSERT INTO `ch_t_d_areainfo` VALUES (652824, '若羌县', 3, 00000652800);
INSERT INTO `ch_t_d_areainfo` VALUES (652825, '且末县', 3, 00000652800);
INSERT INTO `ch_t_d_areainfo` VALUES (652826, '焉耆回族自治县', 3, 00000652800);
INSERT INTO `ch_t_d_areainfo` VALUES (652827, '和静县', 3, 00000652800);
INSERT INTO `ch_t_d_areainfo` VALUES (652828, '和硕县', 3, 00000652800);
INSERT INTO `ch_t_d_areainfo` VALUES (652829, '博湖县', 3, 00000652800);
INSERT INTO `ch_t_d_areainfo` VALUES (652900, '阿克苏地区', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (652901, '阿克苏市', 3, 00000652900);
INSERT INTO `ch_t_d_areainfo` VALUES (652922, '温宿县', 3, 00000652900);
INSERT INTO `ch_t_d_areainfo` VALUES (652923, '库车县', 3, 00000652900);
INSERT INTO `ch_t_d_areainfo` VALUES (652924, '沙雅县', 3, 00000652900);
INSERT INTO `ch_t_d_areainfo` VALUES (652925, '新和县', 3, 00000652900);
INSERT INTO `ch_t_d_areainfo` VALUES (652926, '拜城县', 3, 00000652900);
INSERT INTO `ch_t_d_areainfo` VALUES (652927, '乌什县', 3, 00000652900);
INSERT INTO `ch_t_d_areainfo` VALUES (652928, '阿瓦提县', 3, 00000652900);
INSERT INTO `ch_t_d_areainfo` VALUES (652929, '柯坪县', 3, 00000652900);
INSERT INTO `ch_t_d_areainfo` VALUES (653000, '克孜勒苏柯尔克孜自治州', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (653001, '阿图什市', 3, 00000653000);
INSERT INTO `ch_t_d_areainfo` VALUES (653022, '阿克陶县', 3, 00000653000);
INSERT INTO `ch_t_d_areainfo` VALUES (653023, '阿合奇县', 3, 00000653000);
INSERT INTO `ch_t_d_areainfo` VALUES (653024, '乌恰县', 3, 00000653000);
INSERT INTO `ch_t_d_areainfo` VALUES (653100, '喀什地区', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (653101, '喀什市', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653121, '疏附县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653122, '疏勒县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653123, '英吉沙县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653124, '泽普县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653125, '莎车县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653126, '叶城县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653127, '麦盖提县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653128, '岳普湖县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653129, '伽师县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653130, '巴楚县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653131, '塔什库尔干塔吉克自治县', 3, 00000653100);
INSERT INTO `ch_t_d_areainfo` VALUES (653200, '和田地区', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (653201, '和田市', 3, 00000653200);
INSERT INTO `ch_t_d_areainfo` VALUES (653221, '和田县', 3, 00000653200);
INSERT INTO `ch_t_d_areainfo` VALUES (653222, '墨玉县', 3, 00000653200);
INSERT INTO `ch_t_d_areainfo` VALUES (653223, '皮山县', 3, 00000653200);
INSERT INTO `ch_t_d_areainfo` VALUES (653224, '洛浦县', 3, 00000653200);
INSERT INTO `ch_t_d_areainfo` VALUES (653225, '策勒县', 3, 00000653200);
INSERT INTO `ch_t_d_areainfo` VALUES (653226, '于田县', 3, 00000653200);
INSERT INTO `ch_t_d_areainfo` VALUES (653227, '民丰县', 3, 00000653200);
INSERT INTO `ch_t_d_areainfo` VALUES (654000, '伊犁哈萨克自治州', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (654002, '伊宁市', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654003, '奎屯市', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654004, '霍尔果斯市', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654021, '伊宁县', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654022, '察布查尔锡伯自治县', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654023, '霍城县', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654024, '巩留县', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654025, '新源县', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654026, '昭苏县', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654027, '特克斯县', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654028, '尼勒克县', 3, 00000654000);
INSERT INTO `ch_t_d_areainfo` VALUES (654200, '塔城地区', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (654201, '塔城市', 3, 00000654200);
INSERT INTO `ch_t_d_areainfo` VALUES (654202, '乌苏市', 3, 00000654200);
INSERT INTO `ch_t_d_areainfo` VALUES (654221, '额敏县', 3, 00000654200);
INSERT INTO `ch_t_d_areainfo` VALUES (654223, '沙湾县', 3, 00000654200);
INSERT INTO `ch_t_d_areainfo` VALUES (654224, '托里县', 3, 00000654200);
INSERT INTO `ch_t_d_areainfo` VALUES (654225, '裕民县', 3, 00000654200);
INSERT INTO `ch_t_d_areainfo` VALUES (654226, '和布克赛尔蒙古自治县', 3, 00000654200);
INSERT INTO `ch_t_d_areainfo` VALUES (654300, '阿勒泰地区', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (654301, '阿勒泰市', 3, 00000654300);
INSERT INTO `ch_t_d_areainfo` VALUES (654321, '布尔津县', 3, 00000654300);
INSERT INTO `ch_t_d_areainfo` VALUES (654322, '富蕴县', 3, 00000654300);
INSERT INTO `ch_t_d_areainfo` VALUES (654323, '福海县', 3, 00000654300);
INSERT INTO `ch_t_d_areainfo` VALUES (654324, '哈巴河县', 3, 00000654300);
INSERT INTO `ch_t_d_areainfo` VALUES (654325, '青河县', 3, 00000654300);
INSERT INTO `ch_t_d_areainfo` VALUES (654326, '吉木乃县', 3, 00000654300);
INSERT INTO `ch_t_d_areainfo` VALUES (659000, '自治区直辖县级行政区划', 2, 00000650000);
INSERT INTO `ch_t_d_areainfo` VALUES (659001, '石河子市', 3, 00000659000);
INSERT INTO `ch_t_d_areainfo` VALUES (659002, '阿拉尔市', 3, 00000659000);
INSERT INTO `ch_t_d_areainfo` VALUES (659003, '图木舒克市', 3, 00000659000);
INSERT INTO `ch_t_d_areainfo` VALUES (659004, '五家渠市', 3, 00000659000);
INSERT INTO `ch_t_d_areainfo` VALUES (659006, '铁门关市', 3, 00000659000);
INSERT INTO `ch_t_d_areainfo` VALUES (710000, '台湾省', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (810000, '香港特别行政区', 1, NULL);
INSERT INTO `ch_t_d_areainfo` VALUES (820000, '澳门特别行政区', 1, NULL);

-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
DROP TABLE IF EXISTS `exa_customers`;
CREATE TABLE `exa_customers`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `customer_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `customer_phone_data` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sys_user_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `sys_user_authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_customers_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of exa_customers
-- ----------------------------
INSERT INTO `exa_customers` VALUES (1, '2020-02-25 18:01:48', '2020-02-25 18:01:48', NULL, '测试客户123', '1761111111', 10, '888');

-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_chunks`;
CREATE TABLE `exa_file_chunks`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `exa_file_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `file_chunk_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `file_chunk_number` int(11) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_file_chunks_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;
CREATE TABLE `exa_file_upload_and_downloads`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `key` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_file_upload_and_downloads_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_exa_file_upload_and_downloads_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO `exa_file_upload_and_downloads` VALUES (7, '2019-10-26 22:46:32', '2019-10-26 22:46:32', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1572101191logo.png', 'png', '1572101191logo.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (10, '2019-10-26 23:10:44', '2019-10-26 23:10:44', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1572102643logo.png', 'png', '1572102643logo.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (12, '2019-10-26 23:14:08', '2019-10-26 23:14:08', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1572102846logo.png', 'png', '1572102846logo.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (13, '2019-10-26 23:18:17', '2019-10-26 23:18:17', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1572103096logo.png', 'png', '1572103096logo.png');
INSERT INTO `exa_file_upload_and_downloads` VALUES (15, '2019-12-15 14:31:00', '2019-12-15 14:31:00', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1576391451logo.png', 'png', '1576391451logo.png');

-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
DROP TABLE IF EXISTS `exa_files`;
CREATE TABLE `exa_files`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `file_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `file_md5` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `file_path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `chunk_total` int(11) NULL DEFAULT NULL,
  `is_finish` tinyint(1) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_exa_files_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
DROP TABLE IF EXISTS `jwt_blacklists`;
CREATE TABLE `jwt_blacklists`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `jwt` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_jwt_blacklists_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 39 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO `jwt_blacklists` VALUES (3, '2019-12-28 18:29:05', '2019-12-28 18:29:05', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MTMzNzM2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc1Mjc5MzZ9.T7ikGw-lgAAQlfMne7zPIF-PlfQMg37uBCYJ24Y_B38');
INSERT INTO `jwt_blacklists` VALUES (4, '2019-12-28 18:31:02', '2019-12-28 18:31:02', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MTMzODUzLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc1MjgwNTN9.tDzUm4KNFeJCErNfZGfuF2tcuolga2f_2dE0nTl_UZU');
INSERT INTO `jwt_blacklists` VALUES (5, '2019-12-28 18:31:25', '2019-12-28 18:31:25', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MTMzODcwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc1MjgwNzB9.mspXy9sqQO_5PusPReLalodo_ybWRKxb3Ownf2r2HxE');
INSERT INTO `jwt_blacklists` VALUES (6, '2019-12-30 14:20:10', '2019-12-30 14:20:10', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkxNTc2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODU3NzZ9.AR2KYShboFKsHTjwohxEkA3lytttfZqRH849sl2fNdw');
INSERT INTO `jwt_blacklists` VALUES (7, '2019-12-30 14:21:14', '2019-12-30 14:21:14', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkxNjE2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODU4MTZ9.h8zbDVHM_QbBI-ejGXeQpw0S9oYHJyP4U-TwsVFus9Q');
INSERT INTO `jwt_blacklists` VALUES (8, '2019-12-30 14:21:57', '2019-12-30 14:21:57', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkxNjgxLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODU4ODF9.CSjolDGVpU0g7YG6TaPAlWAMdhtvnBhAi-XYYWZ6RLo');
INSERT INTO `jwt_blacklists` VALUES (9, '2019-12-30 14:25:01', '2019-12-30 14:25:01', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkxODIyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODYwMjJ9.Y_s22Vh5J2ah6Kh1nZQQ8XIQspbT4I7tzc_YJqWrRWM');
INSERT INTO `jwt_blacklists` VALUES (10, '2019-12-30 14:29:26', '2019-12-30 14:29:26', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkyMTU0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODYzNTR9.4HJdx-sfYE5TUUefdwi3yZ6dY_jG7WwEC_55WuGawY8');
INSERT INTO `jwt_blacklists` VALUES (11, '2019-12-30 14:43:43', '2019-12-30 14:43:43', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkyMTcwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODYzNzB9.YEhupQVwjMVBB2eAcAoGG-vJczoxuUyn6KR-tDWU86I');
INSERT INTO `jwt_blacklists` VALUES (12, '2019-12-30 14:55:13', '2019-12-30 14:55:13', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkzMDI3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODcyMjd9.r_sE_Z31cFdS2nCf3iyQjuiZe0Z3HPR07wKBGlUHsnk');
INSERT INTO `jwt_blacklists` VALUES (13, '2019-12-30 14:58:31', '2019-12-30 14:58:31', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkzNzY2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODc5NjZ9.dYFlmyIKQZjzTCKu56wCmxXiW6zOayN_YgygCcvCyLk');
INSERT INTO `jwt_blacklists` VALUES (14, '2019-12-30 14:58:38', '2019-12-30 14:58:38', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkzOTEwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODgxMTB9.pPmzsHU4UceZuPFT_G-SDdxe6FD3MuL47HkovpI-_0c');
INSERT INTO `jwt_blacklists` VALUES (15, '2019-12-30 14:58:58', '2019-12-30 14:58:58', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MjkzOTE4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc2ODgxMTh9.irf98R0belbXtb8x9SxsvuhiYsbHMPbHbFDxaaH0z6Q');
INSERT INTO `jwt_blacklists` VALUES (16, '2020-01-06 16:32:31', '2020-01-06 16:32:31', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTAzMjk5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTc0OTl9.jgLfjvek7sQyuZ2TABQvLOyu_ifNw_KYzfY3VTLL4fw');
INSERT INTO `jwt_blacklists` VALUES (17, '2020-01-06 16:33:08', '2020-01-06 16:33:08', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA0MzU4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTg1NTh9.89r6xHZUBDjfmNpmF02RjQXYTBGUiJvOEDP8pydNt-A');
INSERT INTO `jwt_blacklists` VALUES (18, '2020-01-06 16:33:18', '2020-01-06 16:33:18', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA0MzkyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTg1OTJ9.6Yv9ZYhN-TH9H4SoZEAkjevKVX0vLHL1lVQGFpfBr2U');
INSERT INTO `jwt_blacklists` VALUES (19, '2020-01-06 16:36:06', '2020-01-06 16:36:06', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA0NDA5LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTg2MDl9._9zRRK76XH_KgrW1X9P5GTLW9dwfIixB4QUsC7M3RHA');
INSERT INTO `jwt_blacklists` VALUES (20, '2020-01-06 16:44:06', '2020-01-06 16:44:06', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA0NTcxLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTg3NzF9.5ki0TZooCorK81xWpYa-OO3RR-Bpp5am_uNCNPh4250');
INSERT INTO `jwt_blacklists` VALUES (21, '2020-01-06 16:45:50', '2020-01-06 16:45:50', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MDUwLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTkyNTB9.A0n5faE0X0TyRb_1RvAQBLooY-peapPTD0LnJD03Ul0');
INSERT INTO `jwt_blacklists` VALUES (22, '2020-01-06 16:46:24', '2020-01-06 16:46:24', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MTU0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTkzNTR9.VtqTOJ-MQY2K3w4tM7HgT0z73CEOd3CDqmYqKCjXxnc');
INSERT INTO `jwt_blacklists` VALUES (23, '2020-01-06 16:47:20', '2020-01-06 16:47:20', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MTg3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTkzODd9.fwL1QakF30SHSaGDkPo3weIg0l7kiAGwNq_fKsFxquc');
INSERT INTO `jwt_blacklists` VALUES (24, '2020-01-06 16:47:57', '2020-01-06 16:47:57', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MjQ0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk0NDR9.VoKdA0-brmUlQ5bYufIdMWrS-cCQ2ARm7_jeVtfvCpc');
INSERT INTO `jwt_blacklists` VALUES (25, '2020-01-06 16:49:08', '2020-01-06 16:49:08', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1Mjg1LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk0ODV9.a8-zmyIlJJGdonhXAzNvNH9C-nMa-Voq4bhTbiVKJzE');
INSERT INTO `jwt_blacklists` VALUES (26, '2020-01-06 16:49:32', '2020-01-06 16:49:32', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1MzUyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk1NTJ9.l4e3rjtrDgRsqnQwizJ-ZXVUVM8ywSJcNJkkEVYbdzU');
INSERT INTO `jwt_blacklists` VALUES (27, '2020-01-06 16:49:58', '2020-01-06 16:49:58', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1Mzc3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk1Nzd9.mXUPYvmXbntrdywpBNM0j9sP991cwfhc9b0KvUM4dG4');
INSERT INTO `jwt_blacklists` VALUES (28, '2020-01-06 16:50:56', '2020-01-06 16:50:56', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NDExLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk2MTF9.Z21e8nWHKV5XvYg61CZCz3nMK25m_FmlxncxGMpMS0k');
INSERT INTO `jwt_blacklists` VALUES (29, '2020-01-06 16:52:03', '2020-01-06 16:52:03', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NDY0LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk2NjR9.qzptIyCcL_SPm6TGwXML8Rih3qYqj9GLUpWzTpSPPuI');
INSERT INTO `jwt_blacklists` VALUES (30, '2020-01-06 16:52:36', '2020-01-06 16:52:36', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NTI3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk3Mjd9.D9e8qbx44CLX0ZInwNlIqTGS_sSE069TRIDkQAk7tVY');
INSERT INTO `jwt_blacklists` VALUES (31, '2020-01-06 16:54:35', '2020-01-06 16:54:35', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NTY1LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk3NjV9.D4EZmVAJ96kxcyIfWkT_LA81t1JCuQZcYmQkkoNhtPo');
INSERT INTO `jwt_blacklists` VALUES (32, '2020-01-06 16:55:40', '2020-01-06 16:55:40', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NjgzLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk4ODN9.SJL2fFMbe5VL2YWBzMlrhxbBIJhIHTUeodkEpgH1Xgo');
INSERT INTO `jwt_blacklists` VALUES (33, '2020-01-06 16:57:28', '2020-01-06 16:57:28', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1NzU4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgyOTk5NTh9.6y12UkOeW7vz7gGTcYaN3Y-2Ut2QmjgU9WEuy_pneGM');
INSERT INTO `jwt_blacklists` VALUES (34, '2020-01-06 16:59:02', '2020-01-06 16:59:02', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1ODU1LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgzMDAwNTV9.G0q9X7Ld3cN_BO-K219b7tFAHgtpiAwqLPoxVNKsEl8');
INSERT INTO `jwt_blacklists` VALUES (35, '2020-01-06 16:59:26', '2020-01-06 16:59:26', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4OTA1OTQ2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1NzgzMDAxNDZ9.cmBgWiztsnh7zF3OUNIDQKv8wzGJF7fllUv-4LlYxu8');
INSERT INTO `jwt_blacklists` VALUES (36, '2020-03-21 14:46:14', '2020-03-21 14:46:14', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg1Mzc3ODY3LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODQ3NzIwNjd9.DLhWhD1FdcWLyFLcXQynKJnenbVHrSiKhlDGFRzgo5k');
INSERT INTO `jwt_blacklists` VALUES (37, '2020-03-31 14:24:35', '2020-03-31 14:24:35', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg2MTM4MTA4LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODU1MzIzMDh9.Ro2F2dZLfOk2Z_OPRbweOuCpchr6HlHfQIF5qjfc8y4');
INSERT INTO `jwt_blacklists` VALUES (38, '2020-04-01 16:07:57', '2020-04-01 16:07:57', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTg2MjQwNzQyLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1ODU2MzQ5NDJ9.9qaOFu7D5cq4vxTfLi4pyO_JGcKjVAEJIcoStJWJlYg');

-- ----------------------------
-- Table structure for menus
-- ----------------------------
DROP TABLE IF EXISTS `menus`;
CREATE TABLE `menus`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `menu_level` int(10) UNSIGNED NULL DEFAULT NULL,
  `authority_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `hidden` tinyint(1) NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `parent_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `menu_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 241 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (195, '2019-09-19 22:05:18', '2019-09-19 22:05:18', NULL, 0, 9528, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘');
INSERT INTO `menus` VALUES (196, '2019-09-19 22:06:17', '2019-09-30 15:44:50', NULL, 0, 9528, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', 0, '2', '测试菜单');
INSERT INTO `menus` VALUES (197, '2019-09-19 22:06:38', '2019-09-19 22:06:38', NULL, 0, 9528, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', 0, '3', '超级管理员');
INSERT INTO `menus` VALUES (198, '2019-09-19 22:11:53', '2019-09-19 22:11:53', NULL, 0, 9528, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', 3, '4', '角色管理');
INSERT INTO `menus` VALUES (199, '2019-09-19 22:13:18', '2019-09-19 22:13:18', NULL, 0, 9528, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', 3, '5', '菜单管理');
INSERT INTO `menus` VALUES (231, '2019-09-19 22:05:18', '2019-09-19 22:05:18', NULL, 0, 999, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘');
INSERT INTO `menus` VALUES (232, '2019-09-19 22:06:17', '2019-09-30 15:44:50', NULL, 0, 999, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', 0, '2', '测试菜单');
INSERT INTO `menus` VALUES (233, '2019-09-19 22:05:18', '2019-09-19 22:05:18', NULL, 0, 888, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘');
INSERT INTO `menus` VALUES (234, '2019-09-19 22:06:17', '2019-09-30 15:44:50', NULL, 0, 888, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', 0, '2', '测试菜单');
INSERT INTO `menus` VALUES (235, '2019-09-19 22:06:38', '2019-09-19 22:06:38', NULL, 0, 888, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', 0, '3', '超级管理员');
INSERT INTO `menus` VALUES (236, '2019-09-19 22:11:53', '2019-09-19 22:11:53', NULL, 0, 888, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', 3, '4', '角色管理');
INSERT INTO `menus` VALUES (237, '2019-09-19 22:13:18', '2019-09-19 22:13:18', NULL, 0, 888, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', 3, '5', '菜单管理');
INSERT INTO `menus` VALUES (238, '2019-09-19 22:13:36', '2019-09-19 22:13:36', NULL, 0, 888, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 3, '6', 'api管理');
INSERT INTO `menus` VALUES (239, '2019-10-09 15:12:29', '2019-10-09 15:14:44', NULL, 0, 888, 'user', 'user', 0, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', 3, '17', '用户管理');
INSERT INTO `menus` VALUES (240, '2019-10-13 17:04:24', '2019-10-13 17:04:24', NULL, 0, 888, 'showTest', 'showTest', 0, 'view/showTest/showTest.vue', '展示用vue文件', 'setting', 0, '18', '展示用vue文件');

-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
DROP TABLE IF EXISTS `sys_apis`;
CREATE TABLE `sys_apis`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `authority_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `group` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `method` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'POST',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_apis_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_apis_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 46 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO `sys_apis` VALUES (1, '2019-09-28 11:23:49', '2019-09-28 17:06:16', NULL, NULL, '/base/login', '用户登录', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (2, '2019-09-28 11:32:46', '2019-09-28 17:06:11', NULL, NULL, '/base/regist', '用户注册', 'base', 'POST');
INSERT INTO `sys_apis` VALUES (3, '2019-09-28 11:33:41', '2019-12-11 16:51:41', NULL, NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (4, '2019-09-28 14:09:04', '2019-09-28 17:05:59', NULL, NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (5, '2019-09-28 14:15:50', '2019-09-28 17:05:53', NULL, NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (7, '2019-09-28 14:19:26', '2019-09-28 17:05:44', NULL, NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (8, '2019-09-28 14:19:48', '2019-09-28 17:05:39', NULL, NULL, '/api/updataApi', '更新Api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (10, '2019-09-30 15:05:38', '2019-09-30 15:05:38', NULL, NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO `sys_apis` VALUES (11, '2019-09-30 15:23:09', '2019-09-30 15:23:09', NULL, NULL, '/authority/createAuthority', '创建角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (12, '2019-09-30 15:23:33', '2019-09-30 15:23:33', NULL, NULL, '/authority/deleteAuthority', '删除角色', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (13, '2019-09-30 15:23:57', '2019-09-30 15:23:57', NULL, NULL, '/authority/getAuthorityList', '获取角色列表', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (14, '2019-09-30 15:24:20', '2019-09-30 15:24:20', NULL, NULL, '/menu/getMenu', '获取菜单树', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (15, '2019-09-30 15:24:50', '2019-09-30 15:24:50', NULL, NULL, '/menu/getMenuList', '分页获取基础menu列表', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (16, '2019-09-30 15:25:07', '2019-09-30 15:25:07', NULL, NULL, '/menu/addBaseMenu', '新增菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (17, '2019-09-30 15:25:25', '2019-09-30 15:25:25', NULL, NULL, '/menu/getBaseMenuTree', '获取用户动态路由', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (18, '2019-09-30 15:25:53', '2019-09-30 15:25:53', NULL, NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (19, '2019-09-30 15:26:20', '2019-09-30 15:26:20', NULL, NULL, '/menu/getMenuAuthority', '获取指定角色menu', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (20, '2019-09-30 15:26:43', '2019-09-30 15:26:43', NULL, NULL, '/menu/deleteBaseMenu', '删除菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (21, '2019-09-30 15:28:05', '2019-09-30 15:28:05', NULL, NULL, '/menu/updataBaseMenu', '更新菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (22, '2019-09-30 15:28:21', '2019-09-30 15:28:21', NULL, NULL, '/menu/getBaseMenuById', '根据id获取菜单', 'menu', 'POST');
INSERT INTO `sys_apis` VALUES (23, '2019-09-30 15:29:19', '2019-09-30 15:29:19', NULL, NULL, '/user/changePassword', '修改密码', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (24, '2019-09-30 15:29:33', '2019-09-30 15:29:33', NULL, NULL, '/user/uploadHeaderImg', '上传头像', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (25, '2019-09-30 15:30:00', '2019-09-30 15:30:00', NULL, NULL, '/user/getInfoList', '分页获取用户列表', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (28, '2019-10-09 15:15:17', '2019-10-09 15:17:07', NULL, NULL, '/user/getUserList', '获取用户列表', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (29, '2019-10-09 23:01:40', '2019-10-09 23:01:40', NULL, NULL, '/user/setUserAuthority', '修改用户角色', 'user', 'POST');
INSERT INTO `sys_apis` VALUES (30, '2019-10-26 20:14:38', '2019-10-26 20:14:38', NULL, NULL, '/fileUploadAndDownload/upload', '文件上传示例', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES (31, '2019-10-26 20:14:59', '2019-10-26 20:14:59', NULL, NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES (32, '2019-12-12 13:28:47', '2019-12-12 13:28:47', NULL, NULL, '/casbin/casbinPUpdata', '更改角色api权限', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (33, '2019-12-12 13:28:59', '2019-12-12 13:28:59', NULL, NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO `sys_apis` VALUES (34, '2019-12-12 17:02:15', '2019-12-12 17:02:15', NULL, NULL, '/fileUploadAndDownload/deleteFile', '删除文件', 'fileUploadAndDownload', 'POST');
INSERT INTO `sys_apis` VALUES (35, '2019-12-28 18:18:07', '2019-12-28 18:18:07', NULL, NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单', 'jwt', 'POST');
INSERT INTO `sys_apis` VALUES (36, '2020-01-06 17:56:36', '2020-01-06 17:56:36', NULL, NULL, '/authority/setDataAuthority', '设置角色资源权限', 'authority', 'POST');
INSERT INTO `sys_apis` VALUES (37, '2020-01-13 14:04:05', '2020-01-13 14:04:05', NULL, NULL, '/system/getSystemConfig', '获取配置文件内容', 'system', 'POST');
INSERT INTO `sys_apis` VALUES (38, '2020-01-13 15:02:06', '2020-01-13 15:02:06', NULL, NULL, '/system/setSystemConfig', '设置配置文件内容', 'system', 'POST');
INSERT INTO `sys_apis` VALUES (39, '2020-02-25 15:32:39', '2020-02-25 15:32:39', NULL, NULL, '/customer/createExaCustomer', '创建客户', 'customer', 'POST');
INSERT INTO `sys_apis` VALUES (40, '2020-02-25 15:32:51', '2020-02-25 15:34:56', NULL, NULL, '/customer/updataExaCustomer', '更新客户', 'customer', 'POST');
INSERT INTO `sys_apis` VALUES (41, '2020-02-25 15:33:57', '2020-02-25 15:33:57', NULL, NULL, '/customer/deleteExaCustomer', '删除客户', 'customer', 'POST');
INSERT INTO `sys_apis` VALUES (42, '2020-02-25 15:36:48', '2020-02-25 15:37:16', NULL, NULL, '/customer/getExaCustomer', '获取单一客户', 'customer', 'POST');
INSERT INTO `sys_apis` VALUES (43, '2020-02-25 15:37:06', '2020-02-25 15:37:06', NULL, NULL, '/customer/getExaCustomerList', '获取客户列表', 'customer', 'POST');
INSERT INTO `sys_apis` VALUES (44, '2020-03-12 14:36:54', '2020-03-12 14:56:50', NULL, NULL, '/casbin/casbinTest/:pathParam', 'RESTFUL模式测试', 'casbin', 'GET');
INSERT INTO `sys_apis` VALUES (45, '2020-03-29 23:01:28', '2020-03-29 23:01:28', NULL, NULL, '/autoCode/createTemp', '自动化代码', 'autoCode', 'POST');

-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
DROP TABLE IF EXISTS `sys_authorities`;
CREATE TABLE `sys_authorities`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `authority_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `parent_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `authority_id`(`authority_id`) USING BTREE,
  INDEX `idx_authorities_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_authorities_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO `sys_authorities` VALUES (2, '2019-09-08 16:18:45', '2019-09-08 16:18:45', NULL, '888', '普通用户', '0');
INSERT INTO `sys_authorities` VALUES (6, '2019-09-18 22:23:33', '2019-09-18 22:23:33', NULL, '9528', '测试角色', '0');
INSERT INTO `sys_authorities` VALUES (8, '2019-12-28 18:19:13', '2019-12-28 18:19:13', NULL, '8881', '普通用户子角色', '888');

-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_base_menus`;
CREATE TABLE `sys_base_menus`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `menu_level` int(10) UNSIGNED NULL DEFAULT NULL,
  `parent_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `hidden` tinyint(1) NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_base_menus_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_base_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 43 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO `sys_base_menus` VALUES (1, '2019-09-19 22:05:18', '2019-12-12 16:51:25', NULL, 0, 0, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', '仪表盘', '1');
INSERT INTO `sys_base_menus` VALUES (2, '2019-09-19 22:06:17', '2020-03-27 20:33:58', NULL, 0, 0, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', '测试菜单', '2');
INSERT INTO `sys_base_menus` VALUES (3, '2019-09-19 22:06:38', '2019-12-12 16:51:31', NULL, 0, 0, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', '超级管理员', '3');
INSERT INTO `sys_base_menus` VALUES (4, '2019-09-19 22:11:53', '2019-09-19 22:11:53', NULL, 0, 3, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', '角色管理', '1');
INSERT INTO `sys_base_menus` VALUES (5, '2019-09-19 22:13:18', '2019-12-12 16:57:20', NULL, 0, 3, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', '菜单管理', '2');
INSERT INTO `sys_base_menus` VALUES (6, '2019-09-19 22:13:36', '2019-12-12 16:57:30', NULL, 0, 3, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 'api管理', '3');
INSERT INTO `sys_base_menus` VALUES (17, '2019-10-09 15:12:29', '2019-12-12 16:57:25', NULL, 0, 3, 'user', 'user', 0, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', '用户管理', '4');
INSERT INTO `sys_base_menus` VALUES (18, '2019-10-15 22:27:22', '2019-12-12 16:51:33', NULL, 0, 0, 'person', 'person', 1, 'view/person/person.vue', '个人信息', 'user-solid', '个人信息', '4');
INSERT INTO `sys_base_menus` VALUES (19, '2019-10-20 11:14:42', '2020-03-29 21:39:18', NULL, 0, 0, 'example', 'example', 0, 'view/example/index.vue', '示例文件', 's-management', '示例文件', '6');
INSERT INTO `sys_base_menus` VALUES (20, '2019-10-20 11:18:11', '2019-10-20 11:18:11', NULL, 0, 19, 'table', 'table', 0, 'view/example/table/table.vue', '表格示例', 's-order', '表格示例', '1');
INSERT INTO `sys_base_menus` VALUES (21, '2019-10-20 11:19:52', '2019-12-12 16:58:15', NULL, 0, 19, 'form', 'form', 0, 'view/example/form/form.vue', '表单示例', 'document', '表单示例', '2');
INSERT INTO `sys_base_menus` VALUES (22, '2019-10-20 11:22:19', '2019-12-12 16:58:20', NULL, 0, 19, 'rte', 'rte', 0, 'view/example/rte/rte.vue', '富文本编辑器', 'reading', '富文本编辑器', '3');
INSERT INTO `sys_base_menus` VALUES (23, '2019-10-20 11:23:39', '2019-12-12 16:58:23', NULL, 0, 19, 'excel', 'excel', 0, 'view/example/excel/excel.vue', 'excel导入导出', 's-marketing', 'excel导入导出', '4');
INSERT INTO `sys_base_menus` VALUES (26, '2019-10-20 11:27:02', '2019-12-12 16:58:27', NULL, 0, 19, 'upload', 'upload', 0, 'view/example/upload/upload.vue', '上传下载', 'upload', '上传下载', '5');
INSERT INTO `sys_base_menus` VALUES (29, '2019-12-04 10:05:57', '2019-12-12 16:51:39', '2019-12-12 17:00:50', 0, 0, 'workflow', 'workflow', 0, 'view/workflow/index.vue', '工作流', 'share', '工作流', '6');
INSERT INTO `sys_base_menus` VALUES (30, '2019-12-04 10:06:36', '2019-12-04 10:06:36', '2019-12-12 17:00:48', 0, 29, 'workflowCreate', 'workflowCreate', 0, 'view/workflow/workflowCreate/workflowCreate', '创建工作流', '', '创建工作流', '1');
INSERT INTO `sys_base_menus` VALUES (31, '2019-12-17 10:08:02', '2019-12-17 10:08:58', '2019-12-17 10:09:27', 0, 0, 'testtest', 'testtest', 0, 'view/test/index.vue', '测试menu', '', '测试menu', '8');
INSERT INTO `sys_base_menus` VALUES (32, '2020-01-13 14:03:21', '2020-01-13 14:05:19', '2020-03-29 21:31:23', 0, 3, 'system', 'system', 0, 'view/superAdmin/system/system.vue', '配置管理', 'setting', '配置管理', '5');
INSERT INTO `sys_base_menus` VALUES (33, '2020-02-17 16:20:47', '2020-02-24 19:45:40', NULL, 0, 19, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', '断点续传', 'upload', '断点续传', '6');
INSERT INTO `sys_base_menus` VALUES (34, '2020-02-24 19:48:37', '2020-03-27 20:10:02', NULL, 0, 19, 'customer', 'customer', 0, 'view/example/customer/customer.vue', '客户列表（资源示例）', 's-custom', '客户列表（资源示例）', '7');
INSERT INTO `sys_base_menus` VALUES (35, '2020-03-22 17:13:38', '2020-03-22 17:13:38', '2020-03-22 17:13:49', 0, 0, 'autoCode', 'autoCode', 0, 'view/superAdmin/autoCode/index.vue', '代码构建工具', 's-platform', '代码构建工具', '6');
INSERT INTO `sys_base_menus` VALUES (36, '2020-03-22 17:14:24', '2020-03-22 17:15:23', '2020-03-29 21:31:18', 0, 3, 'autoCode', 'autoCode', 0, 'view/superAdmin/autoCode/index.vue', '自动化代码', 'cpu', '自动化代码', '6');
INSERT INTO `sys_base_menus` VALUES (37, '2020-03-28 23:43:39', '2020-03-28 23:43:39', '2020-03-29 21:31:15', 0, 3, 'formCreate', 'formCreate', 0, 'view/superAdmin/formCreate/index.vue', '表单生成器', 'setting', '表单生成器', '8');
INSERT INTO `sys_base_menus` VALUES (38, '2020-03-29 21:31:03', '2020-03-29 21:31:03', NULL, 0, 0, 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', '系统工具', 's-cooperation', '系统工具', '5');
INSERT INTO `sys_base_menus` VALUES (39, '2020-03-29 21:34:06', '2020-03-30 13:56:24', '2020-03-30 15:58:32', 0, 38, 'system', 'system', 0, 'view/systemTools/system/system.vue', '系统配置', 's-order', '系统配置', '0');
INSERT INTO `sys_base_menus` VALUES (40, '2020-03-29 21:35:10', '2020-03-29 21:35:10', NULL, 0, 38, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', '代码生成器', 'cpu', '代码生成器', '1');
INSERT INTO `sys_base_menus` VALUES (41, '2020-03-29 21:36:26', '2020-03-29 21:36:26', NULL, 0, 38, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', '表单生成器', 'magic-stick', '表单生成器', '2');
INSERT INTO `sys_base_menus` VALUES (42, '2020-04-02 14:19:36', '2020-04-02 14:20:16', NULL, 0, 38, 'system', 'system', 0, 'view/systemTools/system/system.vue', '系统配置', 's-operation', '系统配置', '3');

-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
DROP TABLE IF EXISTS `sys_data_authority_id`;
CREATE TABLE `sys_data_authority_id`  (
  `sys_authority_id` int(10) UNSIGNED NOT NULL DEFAULT 0,
  `data_id` int(10) UNSIGNED NOT NULL DEFAULT 0,
  PRIMARY KEY (`sys_authority_id`, `data_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO `sys_data_authority_id` VALUES (2, 2);
INSERT INTO `sys_data_authority_id` VALUES (2, 6);
INSERT INTO `sys_data_authority_id` VALUES (2, 8);
INSERT INTO `sys_data_authority_id` VALUES (6, 2);
INSERT INTO `sys_data_authority_id` VALUES (6, 6);
INSERT INTO `sys_data_authority_id` VALUES (8, 2);
INSERT INTO `sys_data_authority_id` VALUES (8, 6);
INSERT INTO `sys_data_authority_id` VALUES (8, 8);

-- ----------------------------
-- Table structure for sys_menus
-- ----------------------------
DROP TABLE IF EXISTS `sys_menus`;
CREATE TABLE `sys_menus`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `menu_level` int(10) UNSIGNED NULL DEFAULT NULL,
  `authority_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `path` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `hidden` tinyint(1) NULL DEFAULT NULL,
  `component` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `title` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `icon` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `parent_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `menu_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `sort` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_menus_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_menus_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 580 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_menus
-- ----------------------------
INSERT INTO `sys_menus` VALUES (231, '2019-09-19 22:05:18', '2019-12-12 16:51:25', NULL, 0, 999, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘', '1');
INSERT INTO `sys_menus` VALUES (232, '2019-09-19 22:06:17', '2020-03-27 20:33:58', NULL, 0, 999, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', 0, '2', '测试菜单', '2');
INSERT INTO `sys_menus` VALUES (451, '2019-09-19 22:05:18', '2019-12-12 16:51:25', NULL, 0, 8881, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘', '1');
INSERT INTO `sys_menus` VALUES (452, '2019-09-19 22:06:38', '2019-12-12 16:51:31', NULL, 0, 8881, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', 0, '3', '超级管理员', '3');
INSERT INTO `sys_menus` VALUES (453, '2019-09-19 22:11:53', '2019-09-19 22:11:53', NULL, 0, 8881, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', 3, '4', '角色管理', '1');
INSERT INTO `sys_menus` VALUES (454, '2019-09-19 22:13:18', '2019-12-12 16:57:20', NULL, 0, 8881, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', 3, '5', '菜单管理', '2');
INSERT INTO `sys_menus` VALUES (455, '2019-09-19 22:13:36', '2019-12-12 16:57:30', NULL, 0, 8881, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 3, '6', 'api管理', '3');
INSERT INTO `sys_menus` VALUES (456, '2019-10-09 15:12:29', '2019-12-12 16:57:25', NULL, 0, 8881, 'user', 'user', 0, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', 3, '17', '用户管理', '4');
INSERT INTO `sys_menus` VALUES (458, '2019-10-15 22:27:22', '2019-12-12 16:51:33', NULL, 0, 8881, 'person', 'person', 1, 'view/person/person.vue', '个人信息', 'user-solid', 0, '18', '个人信息', '4');
INSERT INTO `sys_menus` VALUES (459, '2019-10-20 11:14:42', '2020-03-29 21:39:18', NULL, 0, 8881, 'example', 'example', 0, 'view/example/index.vue', '示例文件', 's-management', 0, '19', '示例文件', '6');
INSERT INTO `sys_menus` VALUES (460, '2019-10-20 11:18:11', '2019-10-20 11:18:11', NULL, 0, 8881, 'table', 'table', 0, 'view/example/table/table.vue', '表格示例', 's-order', 19, '20', '表格示例', '1');
INSERT INTO `sys_menus` VALUES (461, '2019-10-20 11:19:52', '2019-12-12 16:58:15', NULL, 0, 8881, 'form', 'form', 0, 'view/example/form/form.vue', '表单示例', 'document', 19, '21', '表单示例', '2');
INSERT INTO `sys_menus` VALUES (462, '2019-10-20 11:22:19', '2019-12-12 16:58:20', NULL, 0, 8881, 'rte', 'rte', 0, 'view/example/rte/rte.vue', '富文本编辑器', 'reading', 19, '22', '富文本编辑器', '3');
INSERT INTO `sys_menus` VALUES (463, '2019-10-20 11:23:39', '2019-12-12 16:58:23', NULL, 0, 8881, 'excel', 'excel', 0, 'view/example/excel/excel.vue', 'excel导入导出', 's-marketing', 19, '23', 'excel导入导出', '4');
INSERT INTO `sys_menus` VALUES (464, '2019-10-20 11:27:02', '2019-12-12 16:58:27', NULL, 0, 8881, 'upload', 'upload', 0, 'view/example/upload/upload.vue', '上传下载', 'upload', 19, '26', '上传下载', '5');
INSERT INTO `sys_menus` VALUES (465, '2020-02-17 16:20:47', '2020-02-24 19:45:40', NULL, 0, 8881, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', '断点续传', 'upload', 19, '33', '断点续传', '6');
INSERT INTO `sys_menus` VALUES (466, '2020-02-24 19:48:37', '2020-03-27 20:05:39', NULL, 0, 8881, 'customer', 'customer', 0, 'view/example/customer/customer.vue', '客户列表（资源示例）', 's-custom', 19, '34', '客户列表（资源示例）', '7');
INSERT INTO `sys_menus` VALUES (503, '2019-09-19 22:05:18', '2019-12-12 16:51:25', NULL, 0, 9528, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘', '1');
INSERT INTO `sys_menus` VALUES (504, '2019-09-19 22:06:17', '2020-03-27 20:33:58', NULL, 0, 9528, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', 0, '2', '测试菜单', '2');
INSERT INTO `sys_menus` VALUES (505, '2019-09-19 22:06:38', '2019-12-12 16:51:31', NULL, 0, 9528, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', 0, '3', '超级管理员', '3');
INSERT INTO `sys_menus` VALUES (506, '2019-09-19 22:11:53', '2019-09-19 22:11:53', NULL, 0, 9528, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', 3, '4', '角色管理', '1');
INSERT INTO `sys_menus` VALUES (507, '2019-09-19 22:13:18', '2019-12-12 16:57:20', NULL, 0, 9528, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', 3, '5', '菜单管理', '2');
INSERT INTO `sys_menus` VALUES (508, '2019-09-19 22:13:36', '2019-12-12 16:57:30', NULL, 0, 9528, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 3, '6', 'api管理', '3');
INSERT INTO `sys_menus` VALUES (509, '2019-10-09 15:12:29', '2019-12-12 16:57:25', NULL, 0, 9528, 'user', 'user', 0, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', 3, '17', '用户管理', '4');
INSERT INTO `sys_menus` VALUES (512, '2019-10-15 22:27:22', '2019-12-12 16:51:33', NULL, 0, 9528, 'person', 'person', 1, 'view/person/person.vue', '个人信息', 'user-solid', 0, '18', '个人信息', '4');
INSERT INTO `sys_menus` VALUES (513, '2019-10-20 11:14:42', '2020-03-29 21:39:18', NULL, 0, 9528, 'example', 'example', 0, 'view/example/index.vue', '示例文件', 's-management', 0, '19', '示例文件', '6');
INSERT INTO `sys_menus` VALUES (514, '2019-10-20 11:18:11', '2019-10-20 11:18:11', NULL, 0, 9528, 'table', 'table', 0, 'view/example/table/table.vue', '表格示例', 's-order', 19, '20', '表格示例', '1');
INSERT INTO `sys_menus` VALUES (515, '2019-10-20 11:19:52', '2019-12-12 16:58:15', NULL, 0, 9528, 'form', 'form', 0, 'view/example/form/form.vue', '表单示例', 'document', 19, '21', '表单示例', '2');
INSERT INTO `sys_menus` VALUES (516, '2019-10-20 11:22:19', '2019-12-12 16:58:20', NULL, 0, 9528, 'rte', 'rte', 0, 'view/example/rte/rte.vue', '富文本编辑器', 'reading', 19, '22', '富文本编辑器', '3');
INSERT INTO `sys_menus` VALUES (517, '2019-10-20 11:23:39', '2019-12-12 16:58:23', NULL, 0, 9528, 'excel', 'excel', 0, 'view/example/excel/excel.vue', 'excel导入导出', 's-marketing', 19, '23', 'excel导入导出', '4');
INSERT INTO `sys_menus` VALUES (518, '2019-10-20 11:27:02', '2019-12-12 16:58:27', NULL, 0, 9528, 'upload', 'upload', 0, 'view/example/upload/upload.vue', '上传下载', 'upload', 19, '26', '上传下载', '5');
INSERT INTO `sys_menus` VALUES (519, '2020-02-17 16:20:47', '2020-02-24 19:45:40', NULL, 0, 9528, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', '断点续传', 'upload', 19, '33', '断点续传', '6');
INSERT INTO `sys_menus` VALUES (520, '2020-02-24 19:48:37', '2020-03-27 20:05:38', NULL, 0, 9528, 'customer', 'customer', 0, 'view/example/customer/customer.vue', '客户列表（资源示例）', 's-custom', 19, '34', '客户列表（资源示例）', '7');
INSERT INTO `sys_menus` VALUES (560, '2019-09-19 22:05:18', '2019-12-12 16:51:25', NULL, 0, 888, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘', '1');
INSERT INTO `sys_menus` VALUES (561, '2019-09-19 22:06:17', '2020-03-27 20:33:58', NULL, 0, 888, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', 0, '2', '测试菜单', '2');
INSERT INTO `sys_menus` VALUES (562, '2019-09-19 22:06:38', '2019-12-12 16:51:31', NULL, 0, 888, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', 0, '3', '超级管理员', '3');
INSERT INTO `sys_menus` VALUES (563, '2019-09-19 22:11:53', '2019-09-19 22:11:53', NULL, 0, 888, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', 3, '4', '角色管理', '1');
INSERT INTO `sys_menus` VALUES (564, '2019-09-19 22:13:18', '2019-12-12 16:57:20', NULL, 0, 888, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', 3, '5', '菜单管理', '2');
INSERT INTO `sys_menus` VALUES (565, '2019-09-19 22:13:36', '2019-12-12 16:57:30', NULL, 0, 888, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 3, '6', 'api管理', '3');
INSERT INTO `sys_menus` VALUES (566, '2019-10-09 15:12:29', '2019-12-12 16:57:25', NULL, 0, 888, 'user', 'user', 0, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', 3, '17', '用户管理', '4');
INSERT INTO `sys_menus` VALUES (567, '2019-10-15 22:27:22', '2019-12-12 16:51:33', NULL, 0, 888, 'person', 'person', 1, 'view/person/person.vue', '个人信息', 'user-solid', 0, '18', '个人信息', '4');
INSERT INTO `sys_menus` VALUES (568, '2020-03-29 21:31:03', '2020-03-29 21:31:03', NULL, 0, 888, 'systemTools', 'systemTools', 0, 'view/systemTools/index.vue', '系统工具', 's-cooperation', 0, '38', '系统工具', '5');
INSERT INTO `sys_menus` VALUES (569, '2020-03-29 21:35:10', '2020-03-29 21:35:10', NULL, 0, 888, 'autoCode', 'autoCode', 0, 'view/systemTools/autoCode/index.vue', '代码生成器', 'cpu', 38, '40', '代码生成器', '1');
INSERT INTO `sys_menus` VALUES (570, '2020-03-29 21:36:26', '2020-03-29 21:36:26', NULL, 0, 888, 'formCreate', 'formCreate', 0, 'view/systemTools/formCreate/index.vue', '表单生成器', 'magic-stick', 38, '41', '表单生成器', '2');
INSERT INTO `sys_menus` VALUES (571, '2020-04-02 14:19:36', '2020-04-02 14:20:16', NULL, 0, 888, 'system', 'system', 0, 'view/systemTools/system/system.vue', '系统配置', 's-operation', 38, '42', '系统配置', '3');
INSERT INTO `sys_menus` VALUES (572, '2019-10-20 11:14:42', '2020-03-29 21:39:18', NULL, 0, 888, 'example', 'example', 0, 'view/example/index.vue', '示例文件', 's-management', 0, '19', '示例文件', '6');
INSERT INTO `sys_menus` VALUES (573, '2019-10-20 11:18:11', '2019-10-20 11:18:11', NULL, 0, 888, 'table', 'table', 0, 'view/example/table/table.vue', '表格示例', 's-order', 19, '20', '表格示例', '1');
INSERT INTO `sys_menus` VALUES (574, '2019-10-20 11:19:52', '2019-12-12 16:58:15', NULL, 0, 888, 'form', 'form', 0, 'view/example/form/form.vue', '表单示例', 'document', 19, '21', '表单示例', '2');
INSERT INTO `sys_menus` VALUES (575, '2019-10-20 11:22:19', '2019-12-12 16:58:20', NULL, 0, 888, 'rte', 'rte', 0, 'view/example/rte/rte.vue', '富文本编辑器', 'reading', 19, '22', '富文本编辑器', '3');
INSERT INTO `sys_menus` VALUES (576, '2019-10-20 11:23:39', '2019-12-12 16:58:23', NULL, 0, 888, 'excel', 'excel', 0, 'view/example/excel/excel.vue', 'excel导入导出', 's-marketing', 19, '23', 'excel导入导出', '4');
INSERT INTO `sys_menus` VALUES (577, '2019-10-20 11:27:02', '2019-12-12 16:58:27', NULL, 0, 888, 'upload', 'upload', 0, 'view/example/upload/upload.vue', '上传下载', 'upload', 19, '26', '上传下载', '5');
INSERT INTO `sys_menus` VALUES (578, '2020-02-17 16:20:47', '2020-02-24 19:45:40', NULL, 0, 888, 'breakpoint', 'breakpoint', 0, 'view/example/breakpoint/breakpoint.vue', '断点续传', 'upload', 19, '33', '断点续传', '6');
INSERT INTO `sys_menus` VALUES (579, '2020-02-24 19:48:37', '2020-03-27 20:10:02', NULL, 0, 888, 'customer', 'customer', 0, 'view/example/customer/customer.vue', '客户列表（资源示例）', 's-custom', 19, '34', '客户列表（资源示例）', '7');

-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
DROP TABLE IF EXISTS `sys_users`;
CREATE TABLE `sys_users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `uuid` varbinary(255) NULL DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pass_word` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'QMPlusUser',
  `header_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'http://www.henrongyi.top/avatar/lufu.jpg',
  `authority_id` double NULL DEFAULT 888,
  `authority_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `phone_data` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `manager` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 12 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO `sys_users` VALUES (10, '2019-09-13 17:23:46', '2019-10-21 11:16:03', NULL, 0x63653064363638352D633135662D343132362D613562342D383930626339643233353664, NULL, NULL, '超级管理员', 'http://qmplusimg.henrongyi.top/1571627762timg.jpg', 888, NULL, 'admin', 'e10adc3949ba59abbe56e057f20f883e', NULL, NULL);
INSERT INTO `sys_users` VALUES (11, '2019-09-13 17:27:29', '2019-09-13 17:27:29', NULL, 0x66643665663739622D393434632D343838382D383337372D616265326432363038383538, NULL, NULL, 'QMPlusUser', 'http://qmplusimg.henrongyi.top/1572075907logo.png', 9528, NULL, 'a303176530', '3ec063004a6f31642261936a379fde3d', NULL, NULL);

-- ----------------------------
-- Table structure for sys_workflow_step_infos
-- ----------------------------
DROP TABLE IF EXISTS `sys_workflow_step_infos`;
CREATE TABLE `sys_workflow_step_infos`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `workflow_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `is_strat` tinyint(1) NULL DEFAULT NULL,
  `step_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `step_no` double NULL DEFAULT NULL,
  `step_authority_id` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `is_end` tinyint(1) NULL DEFAULT NULL,
  `sys_workflow_id` int(10) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_workflow_step_infos_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_workflow_step_infos_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for sys_workflows
-- ----------------------------
DROP TABLE IF EXISTS `sys_workflows`;
CREATE TABLE `sys_workflows`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `workflow_nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `workflow_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `workflow_description` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_workflows_deleted_at`(`deleted_at`) USING BTREE,
  INDEX `idx_sys_workflows_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of sys_workflows
-- ----------------------------
INSERT INTO `sys_workflows` VALUES (8, '2019-12-09 15:20:21', '2019-12-09 15:20:21', NULL, '测试改版1', 'test', '123123');

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` timestamp(0) NULL DEFAULT NULL,
  `updated_at` timestamp(0) NULL DEFAULT NULL,
  `deleted_at` timestamp(0) NULL DEFAULT NULL,
  `uuid` varbinary(255) NULL DEFAULT NULL,
  `user_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `pass_word` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `nick_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'QMPlusUser',
  `header_img` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT 'http://www.henrongyi.top/avatar/lufu.jpg',
  `authority_id` double NULL DEFAULT 888,
  `authority_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `username` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_users_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 16 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (10, '2019-09-13 17:23:46', '2019-09-13 17:23:46', NULL, 0x63653064363638352D633135662D343132362D613562342D383930626339643233353664, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, '303176530', '3ec063004a6f31642261936a379fde3d');
INSERT INTO `users` VALUES (11, '2019-09-13 17:27:29', '2019-09-13 17:27:29', NULL, 0x66643665663739622D393434632D343838382D383337372D616265326432363038383538, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, 'a303176530', '3ec063004a6f31642261936a379fde3d');
INSERT INTO `users` VALUES (12, '2019-09-13 17:28:56', '2019-09-13 17:28:56', NULL, 0x65373939636563362D346337662D343338632D383634372D376435633333393734353165, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, 'a30317465', '3ec063004a6f31642261936a379fde3d');
INSERT INTO `users` VALUES (13, '2019-09-13 17:29:26', '2019-09-13 17:29:26', NULL, 0x65653764353932322D323333312D343162632D393363322D613665366461306465343230, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, 'a30317465', '3ec063004a6f31642261936a379fde3d');
INSERT INTO `users` VALUES (14, '2019-09-13 17:29:28', '2019-09-13 17:29:28', NULL, 0x35623464333461322D343266352D343763352D613932642D613637616536643461643334, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, 'a30317465', '3ec063004a6f31642261936a379fde3d');
INSERT INTO `users` VALUES (15, '2019-09-13 17:31:16', '2019-10-09 23:04:28', NULL, 0x34306437393436612D363732382D346536662D396434302D313432356566653831363032, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 999, NULL, 'a303146523', 'e10adc3949ba59abbe56e057f20f883e');

SET FOREIGN_KEY_CHECKS = 1;
