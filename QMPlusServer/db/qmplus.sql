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

 Date: 13/10/2019 19:51:33
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
) ENGINE = InnoDB AUTO_INCREMENT = 212 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

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
INSERT INTO `api_authorities` VALUES (184, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 1);
INSERT INTO `api_authorities` VALUES (185, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 2);
INSERT INTO `api_authorities` VALUES (186, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 3);
INSERT INTO `api_authorities` VALUES (187, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 4);
INSERT INTO `api_authorities` VALUES (188, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 5);
INSERT INTO `api_authorities` VALUES (189, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 6);
INSERT INTO `api_authorities` VALUES (190, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 7);
INSERT INTO `api_authorities` VALUES (191, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 8);
INSERT INTO `api_authorities` VALUES (192, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 9);
INSERT INTO `api_authorities` VALUES (193, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 10);
INSERT INTO `api_authorities` VALUES (194, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 11);
INSERT INTO `api_authorities` VALUES (195, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 12);
INSERT INTO `api_authorities` VALUES (196, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 13);
INSERT INTO `api_authorities` VALUES (197, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 14);
INSERT INTO `api_authorities` VALUES (198, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 15);
INSERT INTO `api_authorities` VALUES (199, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 16);
INSERT INTO `api_authorities` VALUES (200, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 17);
INSERT INTO `api_authorities` VALUES (201, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 18);
INSERT INTO `api_authorities` VALUES (202, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 19);
INSERT INTO `api_authorities` VALUES (203, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 20);
INSERT INTO `api_authorities` VALUES (204, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 21);
INSERT INTO `api_authorities` VALUES (205, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 22);
INSERT INTO `api_authorities` VALUES (206, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 23);
INSERT INTO `api_authorities` VALUES (207, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 24);
INSERT INTO `api_authorities` VALUES (208, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 28);
INSERT INTO `api_authorities` VALUES (209, '2019-10-09 23:01:52', '2019-10-09 23:01:52', NULL, 888, 29);
INSERT INTO `api_authorities` VALUES (210, '2019-10-09 23:04:56', '2019-10-09 23:04:56', NULL, 999, 1);
INSERT INTO `api_authorities` VALUES (211, '2019-10-09 23:04:56', '2019-10-09 23:04:56', NULL, 999, 2);

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
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

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
) ENGINE = InnoDB AUTO_INCREMENT = 18 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

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
) ENGINE = InnoDB AUTO_INCREMENT = 233 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Compact;

-- ----------------------------
-- Records of menus
-- ----------------------------
INSERT INTO `menus` VALUES (195, '2019-09-19 22:05:18', '2019-09-19 22:05:18', NULL, 0, 9528, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘');
INSERT INTO `menus` VALUES (196, '2019-09-19 22:06:17', '2019-09-30 15:44:50', NULL, 0, 9528, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', 0, '2', '测试菜单');
INSERT INTO `menus` VALUES (197, '2019-09-19 22:06:38', '2019-09-19 22:06:38', NULL, 0, 9528, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', 0, '3', '超级管理员');
INSERT INTO `menus` VALUES (198, '2019-09-19 22:11:53', '2019-09-19 22:11:53', NULL, 0, 9528, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', 3, '4', '角色管理');
INSERT INTO `menus` VALUES (199, '2019-09-19 22:13:18', '2019-09-19 22:13:18', NULL, 0, 9528, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', 3, '5', '菜单管理');
INSERT INTO `menus` VALUES (224, '2019-09-19 22:05:18', '2019-09-19 22:05:18', NULL, 0, 888, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘');
INSERT INTO `menus` VALUES (225, '2019-09-19 22:06:17', '2019-09-30 15:44:50', NULL, 0, 888, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', 0, '2', '测试菜单');
INSERT INTO `menus` VALUES (226, '2019-09-19 22:06:38', '2019-09-19 22:06:38', NULL, 0, 888, 'admin', 'superAdmin', 0, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', 0, '3', '超级管理员');
INSERT INTO `menus` VALUES (227, '2019-09-19 22:11:53', '2019-09-19 22:11:53', NULL, 0, 888, 'authority', 'authority', 0, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', 3, '4', '角色管理');
INSERT INTO `menus` VALUES (228, '2019-09-19 22:13:18', '2019-09-19 22:13:18', NULL, 0, 888, 'menu', 'menu', 0, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', 3, '5', '菜单管理');
INSERT INTO `menus` VALUES (229, '2019-09-19 22:13:36', '2019-09-19 22:13:36', NULL, 0, 888, 'api', 'api', 0, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 3, '6', 'api管理');
INSERT INTO `menus` VALUES (230, '2019-10-09 15:12:29', '2019-10-09 15:14:44', NULL, 0, 888, 'user', 'user', 0, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', 3, '17', '用户管理');
INSERT INTO `menus` VALUES (231, '2019-09-19 22:05:18', '2019-09-19 22:05:18', NULL, 0, 999, 'dashbord', 'dashbord', 0, 'view/dashbord/index.vue', '仪表盘', 'setting', 0, '1', '仪表盘');
INSERT INTO `menus` VALUES (232, '2019-09-19 22:06:17', '2019-09-30 15:44:50', NULL, 0, 999, 'test', 'test', 0, 'view/test/index.vue', '测试菜单', 'info', 0, '2', '测试菜单');

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
INSERT INTO `users` VALUES (10, '2019-09-13 17:23:46', '2019-09-13 17:23:46', NULL, 0x63653064363638352D633135662D343132362D613562342D383930626339643233353664, NULL, NULL, '超级管理员', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, 'admin', 'e10adc3949ba59abbe56e057f20f883e');
INSERT INTO `users` VALUES (11, '2019-09-13 17:27:29', '2019-09-13 17:27:29', NULL, 0x66643665663739622D393434632D343838382D383337372D616265326432363038383538, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, 'a303176530', '3ec063004a6f31642261936a379fde3d');
INSERT INTO `users` VALUES (12, '2019-09-13 17:28:56', '2019-09-13 17:28:56', NULL, 0x65373939636563362D346337662D343338632D383634372D376435633333393734353165, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, 'a30317465', '3ec063004a6f31642261936a379fde3d');
INSERT INTO `users` VALUES (13, '2019-09-13 17:29:26', '2019-09-13 17:29:26', NULL, 0x65653764353932322D323333312D343162632D393363322D613665366461306465343230, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, 'a30317465', '3ec063004a6f31642261936a379fde3d');
INSERT INTO `users` VALUES (14, '2019-09-13 17:29:28', '2019-09-13 17:29:28', NULL, 0x35623464333461322D343266352D343763352D613932642D613637616536643461643334, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 888, NULL, 'a30317465', '3ec063004a6f31642261936a379fde3d');
INSERT INTO `users` VALUES (15, '2019-09-13 17:31:16', '2019-10-09 23:04:28', NULL, 0x34306437393436612D363732382D346536662D396434302D313432356566653831363032, NULL, NULL, 'QMPlusUser', 'http://www.henrongyi.top/avatar/lufu.jpg', 999, NULL, 'a303146523', 'e10adc3949ba59abbe56e057f20f883e');

SET FOREIGN_KEY_CHECKS = 1;
