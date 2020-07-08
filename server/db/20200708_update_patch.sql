INSERT INTO `sys_base_menus` VALUES (52, '2020-06-29 13:31:17', '2020-07-07 16:05:34', NULL, 0, 3, 'operation', 'operation', 0, 'view/superAdmin/operation/sysOperationRecord.vue', '操作历史', 'time', NULL, 6, 0, 0);


INSERT INTO `sys_dictionaries` VALUES (3, '2020-07-05 15:27:31', '2020-07-05 15:27:31', NULL, '数据库int类型', 'int', 1, 'int类型对应的数据库类型');
INSERT INTO `sys_dictionaries` VALUES (4, '2020-07-05 15:33:07', '2020-07-05 16:07:18', NULL, '数据库时间日期类型', 'time.Time', 1, '数据库时间日期类型');
INSERT INTO `sys_dictionaries` VALUES (5, '2020-07-05 15:34:23', '2020-07-05 15:52:45', NULL, '数据库浮点型', 'float64', 1, '数据库浮点型');
INSERT INTO `sys_dictionaries` VALUES (6, '2020-07-05 15:35:05', '2020-07-05 15:35:05', NULL, '数据库字符串', 'string', 1, '数据库字符串');
INSERT INTO `sys_dictionaries` VALUES (7, '2020-07-05 15:36:48', '2020-07-05 15:36:48', NULL, '数据库bool类型', 'bool', 1, '数据库bool类型');


INSERT INTO `sys_dictionary_details` VALUES (12, '2020-07-05 15:31:41', '2020-07-05 15:31:41', NULL, 'smallint', 1, 1, 1, 3);
INSERT INTO `sys_dictionary_details` VALUES (13, '2020-07-05 15:31:52', '2020-07-05 15:31:52', NULL, 'mediumint', 2, 1, 2, 3);
INSERT INTO `sys_dictionary_details` VALUES (14, '2020-07-05 15:32:04', '2020-07-05 15:32:04', NULL, 'int', 3, 1, 3, 3);
INSERT INTO `sys_dictionary_details` VALUES (15, '2020-07-05 15:32:11', '2020-07-05 15:32:11', NULL, 'bigint', 4, 1, 4, 3);
INSERT INTO `sys_dictionary_details` VALUES (19, '2020-07-05 15:33:16', '2020-07-05 15:33:16', NULL, 'data', 0, 1, 0, 4);
INSERT INTO `sys_dictionary_details` VALUES (20, '2020-07-05 15:33:21', '2020-07-05 15:33:21', NULL, 'time', 1, 1, 1, 4);
INSERT INTO `sys_dictionary_details` VALUES (21, '2020-07-05 15:33:25', '2020-07-05 15:33:25', NULL, 'year', 2, 1, 2, 4);
INSERT INTO `sys_dictionary_details` VALUES (22, '2020-07-05 15:33:35', '2020-07-05 15:33:35', NULL, 'datetime', 3, 1, 3, 4);
INSERT INTO `sys_dictionary_details` VALUES (23, '2020-07-05 15:33:42', '2020-07-05 15:33:42', NULL, 'timestamp', 5, 1, 5, 4);
INSERT INTO `sys_dictionary_details` VALUES (24, '2020-07-05 15:34:30', '2020-07-05 15:34:30', NULL, 'float', 0, 1, 0, 5);
INSERT INTO `sys_dictionary_details` VALUES (25, '2020-07-05 15:34:35', '2020-07-05 15:34:35', NULL, 'double', 1, 1, 1, 5);
INSERT INTO `sys_dictionary_details` VALUES (26, '2020-07-05 15:34:41', '2020-07-05 15:34:41', NULL, 'decimal', 2, 1, 2, 5);
INSERT INTO `sys_dictionary_details` VALUES (27, '2020-07-05 15:37:45', '2020-07-05 15:37:45', NULL, 'tinyint', 0, 1, 0, 7);
INSERT INTO `sys_dictionary_details` VALUES (28, '2020-07-05 15:53:25', '2020-07-05 15:53:25', NULL, 'char', 0, 1, 0, 6);
INSERT INTO `sys_dictionary_details` VALUES (29, '2020-07-05 15:53:29', '2020-07-05 15:53:29', NULL, 'varchar', 1, 1, 1, 6);
INSERT INTO `sys_dictionary_details` VALUES (30, '2020-07-05 15:53:35', '2020-07-05 15:53:35', NULL, 'tinyblob', 2, 1, 2, 6);
INSERT INTO `sys_dictionary_details` VALUES (31, '2020-07-05 15:53:40', '2020-07-05 15:53:40', NULL, 'tinytext', 3, 1, 3, 6);
INSERT INTO `sys_dictionary_details` VALUES (32, '2020-07-05 15:53:48', '2020-07-05 15:53:48', NULL, 'text', 4, 1, 4, 6);
INSERT INTO `sys_dictionary_details` VALUES (33, '2020-07-05 15:53:55', '2020-07-05 15:53:55', NULL, 'blob', 5, 1, 5, 6);
INSERT INTO `sys_dictionary_details` VALUES (34, '2020-07-05 15:54:02', '2020-07-05 15:54:02', NULL, 'mediumblob', 6, 1, 6, 6);
INSERT INTO `sys_dictionary_details` VALUES (35, '2020-07-05 15:54:09', '2020-07-05 15:54:09', NULL, 'mediumtext', 7, 1, 7, 6);
INSERT INTO `sys_dictionary_details` VALUES (36, '2020-07-05 15:54:16', '2020-07-05 15:54:16', NULL, 'longblob', 8, 1, 8, 6);
INSERT INTO `sys_dictionary_details` VALUES (37, '2020-07-05 15:54:24', '2020-07-05 15:54:24', NULL, 'longtext', 9, 1, 9, 6);




INSERT INTO `sys_apis` VALUES (91, '2020-06-29 13:21:35', '2020-06-29 13:21:35', NULL, NULL, '/sysOperationRecord/createSysOperationRecord', '新增操作记录', 'sysOperationRecord', 'POST');
INSERT INTO `sys_apis` VALUES (92, '2020-06-29 13:21:35', '2020-06-29 13:21:35', NULL, NULL, '/sysOperationRecord/deleteSysOperationRecord', '删除操作记录', 'sysOperationRecord', 'DELETE');
INSERT INTO `sys_apis` VALUES (93, '2020-06-29 13:21:35', '2020-06-29 13:21:35', NULL, NULL, '/sysOperationRecord/updateSysOperationRecord', '更新操作记录', 'sysOperationRecord', 'PUT');
INSERT INTO `sys_apis` VALUES (94, '2020-06-29 13:21:35', '2020-06-29 13:21:35', NULL, NULL, '/sysOperationRecord/findSysOperationRecord', '根据ID获取操作记录', 'sysOperationRecord', 'GET');
INSERT INTO `sys_apis` VALUES (95, '2020-06-29 13:21:35', '2020-06-29 13:21:35', NULL, NULL, '/sysOperationRecord/getSysOperationRecordList', '获取操作记录列表', 'sysOperationRecord', 'GET');
INSERT INTO `sys_apis` VALUES (96, '2020-07-05 14:34:20', '2020-07-05 14:34:20', NULL, NULL, '/autoCode/getTables', '获取数据库表', 'autoCode', 'GET');
INSERT INTO `sys_apis` VALUES (97, '2020-07-05 15:02:07', '2020-07-05 15:02:07', NULL, NULL, '/autoCode/getDB', '获取所有数据库', 'autoCode', 'GET');
INSERT INTO `sys_apis` VALUES (98, '2020-07-05 16:32:08', '2020-07-05 16:32:08', NULL, NULL, '/autoCode/getColume', '获取所选table的所有字段', 'autoCode', 'GET');
INSERT INTO `sys_apis` VALUES (99, '2020-07-07 15:59:53', '2020-07-07 15:59:53', NULL, NULL, '/sysOperationRecord/deleteSysOperationRecordByIds', '批量删除操作历史', 'sysOperationRecord', 'DELETE');





-- ----------------------------
-- Table structure for sys_operation_records
-- ----------------------------
DROP TABLE IF EXISTS `sys_operation_records`;
CREATE TABLE `sys_operation_records`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` datetime(0) NULL DEFAULT NULL,
  `updated_at` datetime(0) NULL DEFAULT NULL,
  `deleted_at` datetime(0) NULL DEFAULT NULL,
  `ip` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求ip',
  `method` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求方法',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '请求路由',
  `status` int(11) NULL DEFAULT NULL COMMENT '状态',
  `latency` bigint(20) NULL DEFAULT NULL,
  `agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `error_message` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL,
  `body` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '请求Body',
  `user_id` int(11) NULL DEFAULT NULL COMMENT '用户id',
  `resp` text CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL COMMENT '响应Body',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_sys_operation_records_deleted_at`(`deleted_at`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 342 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;



INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/createSysOperationRecord', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecord', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/updateSysOperationRecord', 'PUT', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/findSysOperationRecord', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/getSysOperationRecordList', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/sysOperationRecord/deleteSysOperationRecordByIds', 'DELETE', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getTables', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getDB', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/autoCode/getColume', 'GET', '', '', '');


INSERT INTO `sys_authority_menus` VALUES ('888', 52);


-- 2020/07/08版本更新补丁
-- 增加了共自动化用的字典
-- 增加了操作记录 menu
-- 增加了操作记录相关api
-- 增加了操作记录表
-- 增加了角色和操作记录关联api
-- 增加了角色和menu关联记录