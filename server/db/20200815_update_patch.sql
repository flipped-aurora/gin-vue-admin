INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/upload', 'POST', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/checkFileMd5', 'GET', '', '', '');
INSERT INTO `casbin_rule` VALUES ('p', '888', '/simpleUploader/mergeFileMd5', 'GET', '', '', '');

-- ----------------------------
-- Table structure for exa_simple_uploaders
-- ----------------------------
DROP TABLE IF EXISTS `exa_simple_uploaders`;
CREATE TABLE `exa_simple_uploaders`  (
  `chunk_number` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '当前切片标记',
  `current_chunk_size` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '当前切片容量',
  `current_chunk_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '切片本地路径',
  `total_size` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '总容量',
  `identifier` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '文件标识（md5）',
  `filename` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '文件名',
  `total_chunks` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '切片总数',
  `is_done` tinyint(1) NULL DEFAULT NULL COMMENT '是否上传完成',
  `file_path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NULL DEFAULT NULL COMMENT '文件本地路径'
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_bin ROW_FORMAT = Compact;

INSERT INTO `sys_apis` VALUES (100, '2020-08-15 12:10:55', '2020-08-15 12:10:55', NULL, NULL, '/simpleUploader/upload', '插件版分片上传', 'simpleUploader', 'POST');
INSERT INTO `sys_apis` VALUES (101, '2020-08-15 19:53:53', '2020-08-15 19:53:53', NULL, NULL, '/simpleUploader/checkFileMd5', '文件完整度验证', 'simpleUploader', 'GET');
INSERT INTO `sys_apis` VALUES (102, '2020-08-15 22:28:04', '2020-08-15 22:28:04', NULL, NULL, '/simpleUploader/mergeFileMd5', '上传完成合并文件', 'simpleUploader', 'GET');

INSERT INTO `sys_authority_menus` VALUES ('888', 53);

INSERT INTO `sys_base_menus` VALUES (53, '2020-08-15 11:41:49', '2020-08-15 11:43:15', NULL, 0, 19, 'simpleUploader', 'simpleUploader', 0, 'view/example/simpleUploader/simpleUploader', '断点续传（插件版）', 'upload', NULL, 6, 0, 0);
