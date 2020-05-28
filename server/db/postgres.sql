/*
 Editer: QinMin
 Date: 27/05/2020 18:44:36
*/
BEGIN;
DROP VIEW IF EXISTS authority_menu;
COMMIT;
-- ----------------------------
-- Table structure for casbin_rule
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS casbin_rule;
CREATE TABLE casbin_rule  (
  p_type varchar(100) DEFAULT NULL,
  v0 varchar(100)  DEFAULT NULL,
  v1 varchar(100)  DEFAULT NULL,
  v2 varchar(100)  DEFAULT NULL,
  v3 varchar(100)  DEFAULT NULL,
  v4 varchar(100)  DEFAULT NULL,
  v5 varchar(100)  DEFAULT NULL
);

-- ----------------------------
-- Records of casbin_rule
-- ----------------------------
INSERT INTO casbin_rule VALUES ('p', '2', '/base/login', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/base/register', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/api/createApi', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/api/getApiList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/api/getApiById', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/api/updateApi', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/user/changePassword', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/user/getUserList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/customer/customer', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/customer/customer', 'PUT', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/customer/customer', 'GET', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '2', '/customer/customerList', 'GET', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/base/login', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/base/register', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/api/createApi', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/api/getApiList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/api/getApiById', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/api/updateApi', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/user/changePassword', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/user/getUserList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/customer/customer', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/customer/customer', 'PUT', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/customer/customer', 'GET', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/customer/customerList', 'GET', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '3', '/autoCode/createTemp', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/base/login', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/base/register', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/api/createApi', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/api/getApiList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/api/getApiById', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/api/deleteApi', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/api/updateApi', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/api/getAllApis', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/authority/createAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/authority/deleteAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/authority/getAuthorityList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/authority/setDataAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/authority/updateAuthority', 'PUT', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/authority/copyAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/menu/getMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/menu/getMenuList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/menu/addBaseMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/menu/getBaseMenuTree', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/menu/addMenuAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/menu/getMenuAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/menu/deleteBaseMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/menu/updateBaseMenu', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/menu/getBaseMenuById', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/user/changePassword', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/user/uploadHeaderImg', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/user/deleteUser', 'DELETE', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/user/getUserList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/user/setUserAuthority', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/fileUploadAndDownload/upload', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/fileUploadAndDownload/getFileList', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/fileUploadAndDownload/deleteFile', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/casbin/updateCasbin', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/casbin/getPolicyPathByAuthorityId', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/casbin/casbinTest/:pathParam', 'GET', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/jwt/jsonInBlacklist', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/system/getSystemConfig', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/system/setSystemConfig', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/customer/customer', 'POST', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/customer/customer', 'PUT', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/customer/customer', 'DELETE', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/customer/customer', 'GET', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/customer/customerList', 'GET', '', '', '');
INSERT INTO casbin_rule VALUES ('p', '1', '/autoCode/createTemp', 'POST', '', '', '');
COMMIT;
-- ----------------------------
-- Table structure for exa_customers
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS exa_customers;
DROP SEQUENCE IF EXISTS exa_customers_id_seq;
CREATE SEQUENCE "exa_customers_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE exa_customers  (
  id integer PRIMARY KEY,
  created_at timestamp(0)  DEFAULT NULL,
  updated_at timestamp(0)  DEFAULT NULL,
  deleted_at timestamp(0)  DEFAULT NULL,
  customer_name varchar(255) DEFAULT NULL,
  customer_phone_data varchar(255)  DEFAULT NULL,
  sys_user_id int DEFAULT NULL,
  sys_user_authority_id varchar(255)  DEFAULT NULL
) ;
alter table exa_customers alter column id set default nextval('exa_customers_id_seq');
DROP INDEX IF EXISTS idx_exa_customers_deleted_at;
CREATE INDEX idx_exa_customers_deleted_at ON exa_customers (deleted_at); 
-- ----------------------------
-- Records of exa_customers
-- ----------------------------
INSERT INTO exa_customers VALUES (1, '2020-02-25 18:01:48', '2020-04-10 12:29:29', NULL, '测试客户', '1761111111', 10, '1');
INSERT INTO exa_customers VALUES (2, '2020-04-10 12:25:53', '2020-04-10 12:25:53', '2020-04-10 13:43:56', 'test', '123123123', 10, '1');
INSERT INTO exa_customers VALUES (3, '2020-04-10 13:44:12', '2020-04-10 13:44:12', '2020-04-10 13:44:13', '123123', '123123', 10, '1');
INSERT INTO exa_customers VALUES (4, '2020-04-10 13:47:10', '2020-04-10 13:47:10', '2020-04-10 13:47:12', '22222222', '222222222222222', 10, '1');
COMMIT;
-- ----------------------------
-- Table structure for exa_file_chunks
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS exa_file_chunks;
DROP SEQUENCE IF EXISTS exa_file_chunks_id_seq;
CREATE SEQUENCE "exa_file_chunks_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE exa_file_chunks  (
  id integer PRIMARY KEY,
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  exa_file_id int DEFAULT NULL,
  file_chunk_path varchar(255)  DEFAULT NULL,
  file_chunk_number int DEFAULT NULL
) ;
alter table exa_file_chunks alter column id set default nextval('exa_file_chunks_id_seq');
DROP INDEX IF EXISTS idx_exa_file_chunks_deleted_at;
CREATE INDEX idx_exa_file_chunks_deleted_at ON exa_file_chunks (deleted_at); 
COMMIT;
-- ----------------------------
-- Table structure for exa_file_upload_and_downloads
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS exa_file_upload_and_downloads;
DROP SEQUENCE IF EXISTS exa_file_upload_and_downloads_id_seq;
CREATE SEQUENCE "exa_file_upload_and_downloads_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE exa_file_upload_and_downloads  (
  id integer PRIMARY KEY,
  created_at timestamp(0) DEFAULT NULL,
  updated_at timestamp(0) DEFAULT NULL,
  deleted_at timestamp(0) DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  url varchar(255)  DEFAULT NULL,
  tag varchar(255)  DEFAULT NULL,
  key varchar(255)  DEFAULT NULL
);
alter table exa_file_upload_and_downloads alter column id set default nextval('exa_file_upload_and_downloads_id_seq');
DROP INDEX IF EXISTS idx_file_upload_and_downloads_deleted_at;
CREATE INDEX idx_file_upload_and_downloads_deleted_at ON exa_file_upload_and_downloads (deleted_at); 
DROP INDEX IF EXISTS idx_exa_file_upload_and_downloads_deleted_at;
CREATE INDEX idx_exa_file_upload_and_downloads_deleted_at ON exa_file_upload_and_downloads (deleted_at); 

-- ----------------------------
-- Records of exa_file_upload_and_downloads
-- ----------------------------
INSERT INTO exa_file_upload_and_downloads VALUES (1, '2020-04-26 11:51:39', '2020-04-26 11:51:39', NULL, '10.png', 'http://qmplusimg.henrongyi.top/158787308910.png', 'png', '158787308910.png');
INSERT INTO exa_file_upload_and_downloads VALUES (2, '2020-04-27 15:48:38', '2020-04-27 15:48:38', NULL, 'logo.png', 'http://qmplusimg.henrongyi.top/1587973709logo.png', 'png', '1587973709logo.png');
COMMIT;
-- ----------------------------
-- Table structure for exa_files
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS exa_files;
DROP SEQUENCE IF EXISTS exa_files_id_seq;
CREATE SEQUENCE "exa_files_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE exa_files  (
  id integer PRIMARY KEY,
  created_at timestamp(0)  DEFAULT NULL,
  updated_at timestamp(0)  DEFAULT NULL,
  deleted_at timestamp(0)  DEFAULT NULL,
  file_name varchar(255)  DEFAULT NULL,
  file_md5 varchar(255)  DEFAULT NULL,
  file_path varchar(255)  DEFAULT NULL,
  chunk_total int DEFAULT NULL,
  is_finish int  DEFAULT NULL
) ;
alter table exa_file_chunks alter column id set default nextval('exa_files_id_seq');
DROP INDEX IF EXISTS idx_exa_files_deleted_at;
CREATE INDEX idx_exa_files_deleted_at ON exa_files (deleted_at); 
COMMIT;
-- ----------------------------
-- Table structure for jwt_blacklists
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS jwt_blacklists;
DROP SEQUENCE IF EXISTS jwt_blacklists_id_seq;
CREATE SEQUENCE "jwt_blacklists_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE jwt_blacklists  (
  id integer PRIMARY KEY,
  created_at timestamp(0)  DEFAULT NULL,
  updated_at timestamp(0)  DEFAULT NULL,
  deleted_at timestamp(0)  DEFAULT NULL,
  jwt text DEFAULT NULL
) ;
alter table jwt_blacklists alter column id set default nextval('jwt_blacklists_id_seq');
DROP INDEX IF EXISTS idx_jwt_blacklists_deleted_at;
CREATE INDEX idx_jwt_blacklists_deleted_at ON jwt_blacklists (deleted_at); 

-- ----------------------------
-- Records of jwt_blacklists
-- ----------------------------
INSERT INTO jwt_blacklists VALUES (1, '2019-12-28 18:29:05', '2019-12-28 18:29:05', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MTMzNzM2LCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc1Mjc5MzZ9.T7ikGw-lgAAQlfMne7zPIF-PlfQMg37uBCYJ24Y_B90');
INSERT INTO jwt_blacklists VALUES (2, '2019-12-28 18:31:02', '2019-12-28 18:31:02', NULL, 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiY2UwZDY2ODUtYzE1Zi00MTI2LWE1YjQtODkwYmM5ZDIzNTZkIiwiSUQiOjEwLCJOaWNrTmFtZSI6Iui2hee6p-euoeeQhuWRmCIsIkF1dGhvcml0eUlkIjoiODg4IiwiZXhwIjoxNTc4MTMzODUzLCJpc3MiOiJxbVBsdXMiLCJuYmYiOjE1Nzc1MjgwNTN9.tDzUm4KNFeJCErNfZGfuF2tcuolga2f_2dE0nTl_U80');
COMMIT;
-- ----------------------------
-- Table structure for sys_apis
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS sys_apis;
DROP SEQUENCE IF EXISTS sys_apis_id_seq;
CREATE SEQUENCE "sys_apis_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE sys_apis  (
  id integer PRIMARY KEY,
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  authority_id int DEFAULT NULL,
  path varchar(255) DEFAULT NULL,
  description varchar(255) DEFAULT NULL,
  api_group varchar(255) DEFAULT NULL,
  method varchar(255) DEFAULT 'POST'
) ;
alter table sys_apis alter column id set default nextval('sys_apis_id_seq');
DROP INDEX IF EXISTS idx_apis_deleted_at;
CREATE INDEX idx_apis_deleted_at ON sys_apis (deleted_at); 
DROP INDEX IF EXISTS idx_sys_apis_deleted_at;
CREATE INDEX idx_sys_apis_deleted_at ON sys_apis (deleted_at); 

-- ----------------------------
-- Records of sys_apis
-- ----------------------------
INSERT INTO sys_apis VALUES (1, '2019-09-28 11:23:49', '2019-09-28 17:06:16', NULL, NULL, '/base/login', '用户登录', 'base', 'POST');
INSERT INTO sys_apis VALUES (2, '2019-09-28 11:32:46', '2019-09-28 17:06:11', NULL, NULL, '/base/register', '用户注册', 'base', 'POST');
INSERT INTO sys_apis VALUES (3, '2019-09-28 11:33:41', '2020-05-09 17:43:15', NULL, NULL, '/api/createApi', '创建api', 'api', 'POST');
INSERT INTO sys_apis VALUES (4, '2019-09-28 14:09:04', '2019-09-28 17:05:59', NULL, NULL, '/api/getApiList', '获取api列表', 'api', 'POST');
INSERT INTO sys_apis VALUES (5, '2019-09-28 14:15:50', '2019-09-28 17:05:53', NULL, NULL, '/api/getApiById', '获取api详细信息', 'api', 'POST');
INSERT INTO sys_apis VALUES (7, '2019-09-28 14:19:26', '2019-09-28 17:05:44', NULL, NULL, '/api/deleteApi', '删除Api', 'api', 'POST');
INSERT INTO sys_apis VALUES (8, '2019-09-28 14:19:48', '2019-09-28 17:05:39', NULL, NULL, '/api/updateApi', '更新Api', 'api', 'POST');
INSERT INTO sys_apis VALUES (10, '2019-09-30 15:05:38', '2019-09-30 15:05:38', NULL, NULL, '/api/getAllApis', '获取所有api', 'api', 'POST');
INSERT INTO sys_apis VALUES (11, '2019-09-30 15:23:09', '2019-09-30 15:23:09', NULL, NULL, '/authority/createAuthority', '创建角色', 'authority', 'POST');
INSERT INTO sys_apis VALUES (12, '2019-09-30 15:23:33', '2019-09-30 15:23:33', NULL, NULL, '/authority/deleteAuthority', '删除角色', 'authority', 'POST');
INSERT INTO sys_apis VALUES (13, '2019-09-30 15:23:57', '2019-09-30 15:23:57', NULL, NULL, '/authority/getAuthorityList', '获取角色列表', 'authority', 'POST');
INSERT INTO sys_apis VALUES (14, '2019-09-30 15:24:20', '2019-09-30 15:24:20', NULL, NULL, '/menu/getMenu', '获取菜单树', 'menu', 'POST');
INSERT INTO sys_apis VALUES (15, '2019-09-30 15:24:50', '2019-09-30 15:24:50', NULL, NULL, '/menu/getMenuList', '分页获取基础menu列表', 'menu', 'POST');
INSERT INTO sys_apis VALUES (16, '2019-09-30 15:25:07', '2019-09-30 15:25:07', NULL, NULL, '/menu/addBaseMenu', '新增菜单', 'menu', 'POST');
INSERT INTO sys_apis VALUES (17, '2019-09-30 15:25:25', '2019-09-30 15:25:25', NULL, NULL, '/menu/getBaseMenuTree', '获取用户动态路由', 'menu', 'POST');
INSERT INTO sys_apis VALUES (18, '2019-09-30 15:25:53', '2019-09-30 15:25:53', NULL, NULL, '/menu/addMenuAuthority', '增加menu和角色关联关系', 'menu', 'POST');
INSERT INTO sys_apis VALUES (19, '2019-09-30 15:26:20', '2019-09-30 15:26:20', NULL, NULL, '/menu/getMenuAuthority', '获取指定角色menu', 'menu', 'POST');
INSERT INTO sys_apis VALUES (20, '2019-09-30 15:26:43', '2019-09-30 15:26:43', NULL, NULL, '/menu/deleteBaseMenu', '删除菜单', 'menu', 'POST');
INSERT INTO sys_apis VALUES (21, '2019-09-30 15:28:05', '2019-09-30 15:28:05', NULL, NULL, '/menu/updateBaseMenu', '更新菜单', 'menu', 'POST');
INSERT INTO sys_apis VALUES (22, '2019-09-30 15:28:21', '2019-09-30 15:28:21', NULL, NULL, '/menu/getBaseMenuById', '根据id获取菜单', 'menu', 'POST');
INSERT INTO sys_apis VALUES (23, '2019-09-30 15:29:19', '2019-09-30 15:29:19', NULL, NULL, '/user/changePassword', '修改密码', 'user', 'POST');
INSERT INTO sys_apis VALUES (24, '2019-09-30 15:29:33', '2019-09-30 15:29:33', NULL, NULL, '/user/uploadHeaderImg', '上传头像', 'user', 'POST');
INSERT INTO sys_apis VALUES (25, '2019-09-30 15:30:00', '2020-05-06 16:03:47', NULL, NULL, '/user/deleteUser', '删除用户', 'user', 'DELETE');
INSERT INTO sys_apis VALUES (28, '2019-10-09 15:15:17', '2019-10-09 15:17:07', NULL, NULL, '/user/getUserList', '获取用户列表', 'user', 'POST');
INSERT INTO sys_apis VALUES (29, '2019-10-09 23:01:40', '2019-10-09 23:01:40', NULL, NULL, '/user/setUserAuthority', '修改用户角色', 'user', 'POST');
INSERT INTO sys_apis VALUES (30, '2019-10-26 20:14:38', '2019-10-26 20:14:38', NULL, NULL, '/fileUploadAndDownload/upload', '文件上传示例', 'fileUploadAndDownload', 'POST');
INSERT INTO sys_apis VALUES (31, '2019-10-26 20:14:59', '2019-10-26 20:14:59', NULL, NULL, '/fileUploadAndDownload/getFileList', '获取上传文件列表', 'fileUploadAndDownload', 'POST');
INSERT INTO sys_apis VALUES (32, '2019-12-12 13:28:47', '2019-12-12 13:28:47', NULL, NULL, '/casbin/updateCasbin', '更改角色api权限', 'casbin', 'POST');
INSERT INTO sys_apis VALUES (33, '2019-12-12 13:28:59', '2019-12-12 13:28:59', NULL, NULL, '/casbin/getPolicyPathByAuthorityId', '获取权限列表', 'casbin', 'POST');
INSERT INTO sys_apis VALUES (34, '2019-12-12 17:02:15', '2019-12-12 17:02:15', NULL, NULL, '/fileUploadAndDownload/deleteFile', '删除文件', 'fileUploadAndDownload', 'POST');
INSERT INTO sys_apis VALUES (35, '2019-12-28 18:18:07', '2019-12-28 18:18:07', NULL, NULL, '/jwt/jsonInBlacklist', 'jwt加入黑名单', 'jwt', 'POST');
INSERT INTO sys_apis VALUES (36, '2020-01-06 17:56:36', '2020-01-06 17:56:36', NULL, NULL, '/authority/setDataAuthority', '设置角色资源权限', 'authority', 'POST');
INSERT INTO sys_apis VALUES (37, '2020-01-13 14:04:05', '2020-01-13 14:04:05', NULL, NULL, '/system/getSystemConfig', '获取配置文件内容', 'system', 'POST');
INSERT INTO sys_apis VALUES (38, '2020-01-13 15:02:06', '2020-01-13 15:02:06', NULL, NULL, '/system/setSystemConfig', '设置配置文件内容', 'system', 'POST');
INSERT INTO sys_apis VALUES (39, '2020-02-25 15:32:39', '2020-02-25 15:32:39', NULL, NULL, '/customer/customer', '创建客户', 'customer', 'POST');
INSERT INTO sys_apis VALUES (40, '2020-02-25 15:32:51', '2020-02-25 15:34:56', NULL, NULL, '/customer/customer', '更新客户', 'customer', 'PUT');
INSERT INTO sys_apis VALUES (41, '2020-02-25 15:33:57', '2020-02-25 15:33:57', NULL, NULL, '/customer/customer', '删除客户', 'customer', 'DELETE');
INSERT INTO sys_apis VALUES (42, '2020-02-25 15:36:48', '2020-02-25 15:37:16', NULL, NULL, '/customer/customer', '获取单一客户', 'customer', 'GET');
INSERT INTO sys_apis VALUES (43, '2020-02-25 15:37:06', '2020-02-25 15:37:06', NULL, NULL, '/customer/customerList', '获取客户列表', 'customer', 'GET');
INSERT INTO sys_apis VALUES (44, '2020-03-12 14:36:54', '2020-03-12 14:56:50', NULL, NULL, '/casbin/casbinTest/:pathParam', 'RESTFUL模式测试', 'casbin', 'GET');
INSERT INTO sys_apis VALUES (45, '2020-03-29 23:01:28', '2020-03-29 23:01:28', NULL, NULL, '/autoCode/createTemp', '自动化代码', 'autoCode', 'POST');
INSERT INTO sys_apis VALUES (46, '2020-04-15 12:46:58', '2020-04-15 12:46:58', NULL, NULL, '/authority/updateAuthority', '更新角色信息', 'authority', 'PUT');
INSERT INTO sys_apis VALUES (47, '2020-04-20 15:14:25', '2020-04-20 15:14:25', NULL, NULL, '/authority/copyAuthority', '拷贝角色', 'authority', 'POST');
COMMIT;
-- ----------------------------
-- Table structure for sys_authorities
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS sys_authorities;
CREATE TABLE sys_authorities  (
  authority_id varchar(255)  NOT NULL   PRIMARY KEY,
  authority_name varchar(255) DEFAULT NULL,
  parent_id varchar(255)  DEFAULT NULL,
  created_at timestamp(0)  DEFAULT NULL,
  updated_at timestamp(0)  DEFAULT NULL,
  deleted_at timestamp(0)  DEFAULT NULL
) ;
DROP INDEX IF EXISTS idx_sys_authorities_deleted_at;
CREATE INDEX idx_sys_authorities_deleted_at ON sys_authorities (deleted_at); 
DROP INDEX IF EXISTS authority_id;
CREATE unique index authority_id ON sys_authorities (authority_id); 


-- ----------------------------
-- Records of sys_authorities
-- ----------------------------
INSERT INTO sys_authorities VALUES ('1', '管理用户', '0', '2020-04-04 11:44:56', '2020-05-09 17:41:29', NULL);
INSERT INTO sys_authorities VALUES ('2', '管理用户子角色', '1', '2020-04-04 11:44:56', '2020-05-09 17:41:29', NULL);
INSERT INTO sys_authorities VALUES ('3', '测试角色', '0', '2020-04-04 11:44:56', '2020-05-09 17:41:29', NULL);
COMMIT;
-- ----------------------------
-- Table structure for sys_authority_menus
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS sys_authority_menus;
CREATE TABLE sys_authority_menus (
  sys_authority_authority_id varchar(255)  NOT NULL,
  sys_base_menu_id INT NOT NULL DEFAULT 0,
  PRIMARY KEY (sys_authority_authority_id, sys_base_menu_id) 
) ;
DROP INDEX IF EXISTS sys_authority_authority_id;
CREATE INDEX sys_authority_authority_id ON sys_authority_menus (sys_authority_authority_id); 

-- ----------------------------
-- Records of sys_authority_menus
-- ----------------------------
INSERT INTO sys_authority_menus VALUES ('1', 1);
INSERT INTO sys_authority_menus VALUES ('1', 2);
INSERT INTO sys_authority_menus VALUES ('1', 3);
INSERT INTO sys_authority_menus VALUES ('1', 4);
INSERT INTO sys_authority_menus VALUES ('1', 5);
INSERT INTO sys_authority_menus VALUES ('1', 6);
INSERT INTO sys_authority_menus VALUES ('1', 17);
INSERT INTO sys_authority_menus VALUES ('1', 18);
INSERT INTO sys_authority_menus VALUES ('1', 19);
INSERT INTO sys_authority_menus VALUES ('1', 20);
INSERT INTO sys_authority_menus VALUES ('1', 21);
INSERT INTO sys_authority_menus VALUES ('1', 22);
INSERT INTO sys_authority_menus VALUES ('1', 23);
INSERT INTO sys_authority_menus VALUES ('1', 26);
INSERT INTO sys_authority_menus VALUES ('1', 33);
INSERT INTO sys_authority_menus VALUES ('1', 34);
INSERT INTO sys_authority_menus VALUES ('1', 38);
INSERT INTO sys_authority_menus VALUES ('1', 40);
INSERT INTO sys_authority_menus VALUES ('1', 41);
INSERT INTO sys_authority_menus VALUES ('1', 42);
INSERT INTO sys_authority_menus VALUES ('1', 45);
INSERT INTO sys_authority_menus VALUES ('2', 1);
INSERT INTO sys_authority_menus VALUES ('2', 2);
INSERT INTO sys_authority_menus VALUES ('2', 18);
INSERT INTO sys_authority_menus VALUES ('2', 38);
INSERT INTO sys_authority_menus VALUES ('2', 40);
INSERT INTO sys_authority_menus VALUES ('2', 41);
INSERT INTO sys_authority_menus VALUES ('2', 42);
INSERT INTO sys_authority_menus VALUES ('3', 1);
INSERT INTO sys_authority_menus VALUES ('3', 2);
INSERT INTO sys_authority_menus VALUES ('3', 3);
INSERT INTO sys_authority_menus VALUES ('3', 4);
INSERT INTO sys_authority_menus VALUES ('3', 5);
INSERT INTO sys_authority_menus VALUES ('3', 6);
INSERT INTO sys_authority_menus VALUES ('3', 17);
INSERT INTO sys_authority_menus VALUES ('3', 18);
INSERT INTO sys_authority_menus VALUES ('3', 19);
INSERT INTO sys_authority_menus VALUES ('3', 20);
INSERT INTO sys_authority_menus VALUES ('3', 21);
INSERT INTO sys_authority_menus VALUES ('3', 22);
INSERT INTO sys_authority_menus VALUES ('3', 23);
INSERT INTO sys_authority_menus VALUES ('3', 26);
INSERT INTO sys_authority_menus VALUES ('3', 33);
INSERT INTO sys_authority_menus VALUES ('3', 34);
INSERT INTO sys_authority_menus VALUES ('3', 38);
INSERT INTO sys_authority_menus VALUES ('3', 40);
INSERT INTO sys_authority_menus VALUES ('3', 41);
INSERT INTO sys_authority_menus VALUES ('3', 42);
COMMIT;
-- ----------------------------
-- Table structure for sys_base_menus
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS sys_base_menus;
DROP SEQUENCE IF EXISTS sys_base_menus_id_seq;
CREATE SEQUENCE "sys_base_menus_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE sys_base_menus  (
  id integer PRIMARY KEY,
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  menu_level int DEFAULT NULL,
  parent_id int DEFAULT NULL,
  path varchar(255)  DEFAULT NULL,
  name varchar(255) DEFAULT NULL,
  hidden boolean DEFAULT NULL,
  component varchar(255) DEFAULT NULL,
  title varchar(255)  DEFAULT NULL,
  icon varchar(255)  DEFAULT NULL,
  nick_name varchar(255)  DEFAULT NULL,
  sort int  DEFAULT NULL,
  keep_alive boolean  DEFAULT NULL,
  default_menu boolean  DEFAULT NULL
);
alter table sys_base_menus alter column id set default nextval('sys_base_menus_id_seq');
DROP INDEX IF EXISTS idx_base_menus_deleted_at;
CREATE INDEX idx_base_menus_deleted_at ON sys_base_menus (deleted_at);
DROP INDEX IF EXISTS idx_sys_base_menus_deleted_at;
CREATE INDEX idx_sys_base_menus_deleted_at ON sys_base_menus (deleted_at);

-- ----------------------------
-- Records of sys_base_menus
-- ----------------------------
INSERT INTO sys_base_menus VALUES (1, '2019-09-19 22:05:18', '2020-05-09 17:41:51', NULL, 0, 0, 'dashboard', 'dashboard', false, 'view/dashboard/index.vue', '仪表盘', 'setting', '仪表盘', 1, false, false);
INSERT INTO sys_base_menus VALUES (2, '2019-09-19 22:06:17', '2020-05-09 17:41:29', NULL, 0, 0, 'about', 'about', false, 'view/about/index.vue', '关于我们', 'info', '测试菜单', 7, false, false);
INSERT INTO sys_base_menus VALUES (3, '2019-09-19 22:06:38', '2020-05-09 17:41:29', NULL, 0, 0, 'admin', 'superAdmin', false, 'view/superAdmin/index.vue', '超级管理员', 'user-solid', '超级管理员', 3, false, false);
INSERT INTO sys_base_menus VALUES (4, '2019-09-19 22:11:53', '2020-05-09 17:41:29', NULL, 0, 3, 'authority', 'authority', false, 'view/superAdmin/authority/authority.vue', '角色管理', 's-custom', '角色管理', 1, false, false);
INSERT INTO sys_base_menus VALUES (5, '2019-09-19 22:13:18', '2020-05-09 17:41:29', NULL, 0, 3, 'menu', 'menu', false, 'view/superAdmin/menu/menu.vue', '菜单管理', 's-order', '菜单管理', 2, true, false);
INSERT INTO sys_base_menus VALUES (6, '2019-09-19 22:13:36', '2020-05-09 17:41:29', NULL, 0, 3, 'api', 'api', false, 'view/superAdmin/api/api.vue', 'api管理', 's-platform', 'api管理', 3, true, false);
INSERT INTO sys_base_menus VALUES (17, '2019-10-09 15:12:29', '2020-05-09 17:41:29', NULL, 0, 3, 'user', 'user', false, 'view/superAdmin/user/user.vue', '用户管理', 'coordinate', '用户管理', 4, false, false);
INSERT INTO sys_base_menus VALUES (18, '2019-10-15 22:27:22', '2020-05-09 17:41:29', NULL, 0, 0, 'person', 'person', true, 'view/person/person.vue', '个人信息', 'user-solid', '个人信息', 4, false, false);
INSERT INTO sys_base_menus VALUES (19, '2019-10-20 11:14:42', '2020-05-09 17:41:29', NULL, 0, 0, 'example', 'example', false, 'view/example/index.vue', '示例文件', 's-management', '示例文件', 6, false, false);
INSERT INTO sys_base_menus VALUES (20, '2019-10-20 11:18:11', '2020-05-09 17:41:29', NULL, 0, 19, 'table', 'table', false, 'view/example/table/table.vue', '表格示例', 's-order', '表格示例', 1, false, false);
INSERT INTO sys_base_menus VALUES (21, '2019-10-20 11:19:52', '2020-05-09 17:41:29', NULL, 0, 19, 'form', 'form', false, 'view/example/form/form.vue', '表单示例', 'document', '表单示例', 2, false, false);
INSERT INTO sys_base_menus VALUES (22, '2019-10-20 11:22:19', '2020-05-09 17:41:29', NULL, 0, 19, 'rte', 'rte', false, 'view/example/rte/rte.vue', '富文本编辑器', 'reading', '富文本编辑器', 3, false, false);
INSERT INTO sys_base_menus VALUES (23, '2019-10-20 11:23:39', '2020-05-09 17:41:29', NULL, 0, 19, 'excel', 'excel', false, 'view/example/excel/excel.vue', 'excel导入导出', 's-marketing', 'excel导入导出', 4, false, false);
INSERT INTO sys_base_menus VALUES (26, '2019-10-20 11:27:02', '2020-05-09 17:41:29', NULL, 0, 19, 'upload', 'upload', false, 'view/example/upload/upload.vue', '上传下载', 'upload', '上传下载', 5, false, false);
INSERT INTO sys_base_menus VALUES (33, '2020-02-17 16:20:47', '2020-05-09 17:41:29', NULL, 0, 19, 'breakpoint', 'breakpoint', false, 'view/example/breakpoint/breakpoint.vue', '断点续传', 'upload', '断点续传', 6, false, false);
INSERT INTO sys_base_menus VALUES (34, '2020-02-24 19:48:37', '2020-05-09 17:41:29', NULL, 0, 19, 'customer', 'customer', false, 'view/example/customer/customer.vue', '客户列表（资源示例）', 's-custom', '客户列表（资源示例）', 7, false, false);
INSERT INTO sys_base_menus VALUES (38, '2020-03-29 21:31:03', '2020-05-09 17:41:29', NULL, 0, 0, 'systemTools', 'systemTools', false, 'view/systemTools/index.vue', '系统工具', 's-cooperation', '系统工具', 5, false, false);
INSERT INTO sys_base_menus VALUES (40, '2020-03-29 21:35:10', '2020-05-09 17:41:29', NULL, 0, 38, 'autoCode', 'autoCode', false, 'view/systemTools/autoCode/index.vue', '代码生成器', 'cpu', '代码生成器', 1, false, false);
INSERT INTO sys_base_menus VALUES (41, '2020-03-29 21:36:26', '2020-05-09 17:41:29', NULL, 0, 38, 'formCreate', 'formCreate', false, 'view/systemTools/formCreate/index.vue', '表单生成器', 'magic-stick', '表单生成器', 2, false, false);
INSERT INTO sys_base_menus VALUES (42, '2020-04-02 14:19:36', '2020-05-09 17:41:29', NULL, 0, 38, 'system', 'system', false, 'view/systemTools/system/system.vue', '系统配置', 's-operation', '系统配置', 3, false, false);
INSERT INTO sys_base_menus VALUES (45, '2020-04-29 17:19:34', '2020-05-09 17:42:23', NULL, 0, 0, 'iconList', 'iconList', false, 'view/iconList/index.vue', '图标集合', 'star-on', NULL, 2, false, false);
COMMIT;
-- ----------------------------
-- Table structure for sys_data_authority_id
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS sys_data_authority_id;
CREATE TABLE sys_data_authority_id  (
  sys_authority_authority_id varchar(255)  NOT NULL,
  data_authority_id varchar(255)  NOT NULL,
  PRIMARY KEY (sys_authority_authority_id, data_authority_id) 
) ;
DROP INDEX IF EXISTS sys_authority_authority_id;
CREATE INDEX sys_authority_authority_id ON sys_data_authority_id (sys_authority_authority_id);
DROP INDEX IF EXISTS data_authority_id;
CREATE INDEX data_authority_id ON sys_data_authority_id (data_authority_id);

-- ----------------------------
-- Records of sys_data_authority_id
-- ----------------------------
INSERT INTO sys_data_authority_id VALUES ('1', '1');
INSERT INTO sys_data_authority_id VALUES ('1', '2');
INSERT INTO sys_data_authority_id VALUES ('1', '3');
INSERT INTO sys_data_authority_id VALUES ('88822', '1');
INSERT INTO sys_data_authority_id VALUES ('88822', '2');
INSERT INTO sys_data_authority_id VALUES ('88822', '3');
INSERT INTO sys_data_authority_id VALUES ('888222', '1');
INSERT INTO sys_data_authority_id VALUES ('888222', '2');
INSERT INTO sys_data_authority_id VALUES ('888222', '3');
INSERT INTO sys_data_authority_id VALUES ('8883', '1');
INSERT INTO sys_data_authority_id VALUES ('8883', '2');
INSERT INTO sys_data_authority_id VALUES ('8883', '3');
INSERT INTO sys_data_authority_id VALUES ('3', '2');
INSERT INTO sys_data_authority_id VALUES ('3', '3');
COMMIT;
-- ----------------------------
-- Table structure for sys_users
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS sys_users;
DROP SEQUENCE IF EXISTS sys_users_id_seq;
CREATE SEQUENCE "sys_users_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE sys_users  (
  id integer PRIMARY KEY,
  created_at timestamp(0)  DEFAULT NULL,
  updated_at timestamp(0)  DEFAULT NULL,
  deleted_at timestamp(0)  DEFAULT NULL,
  uuid varchar(255)  DEFAULT NULL,
  user_name varchar(255) DEFAULT NULL,
  pass_word varchar(255)  DEFAULT NULL,
  nick_name varchar(255)  DEFAULT 'QMPlusUser',
  header_img varchar(255)  DEFAULT 'http://www.henrongyi.top/avatar/lufu.jpg',
  authority_id varchar(255)  DEFAULT '1',
  authority_name varchar(255)  DEFAULT NULL,
  username varchar(255)  DEFAULT NULL,
  password varchar(255)  DEFAULT NULL,
  phone_data varchar(255)  DEFAULT NULL,
  manager varchar(255)  DEFAULT NULL
) ;
alter table sys_users alter column id set default nextval('sys_users_id_seq');
DROP INDEX IF EXISTS idx_users_deleted_at;
CREATE INDEX idx_users_deleted_at ON sys_users (deleted_at);
DROP INDEX IF EXISTS idx_sys_users_deleted_at;
CREATE INDEX idx_sys_users_deleted_at ON sys_users (deleted_at);

-- ----------------------------
-- Records of sys_users
-- ----------------------------
INSERT INTO sys_users VALUES (1, '2019-09-13 17:23:46', '2020-05-06 16:09:15', NULL, 'ce0d6685-c15f-4126-a5b4-890bc9d2356d', NULL, NULL, '超级管理员', 'http://qmplusimg.henrongyi.top/15887525450B978439-F04A-4a09-A8D3-DE7DE2677142.png', '1', NULL, 'admin', 'e10adc3949ba59abbe56e057f20f883e', NULL, NULL);
INSERT INTO sys_users VALUES (2, '2019-09-13 17:27:29', '2020-05-09 17:43:44', NULL, 'fd6ef79b-944c-4888-8377-abe2d2608858', NULL, NULL, 'QMPlusUser', 'http://qmplusimg.henrongyi.top/1572075907logo.png', '3', NULL, 'a303176530', '3ec063004a6f31642261936a379fde3d', NULL, NULL);
COMMIT;
-- ----------------------------
-- Table structure for sys_workflow_step_infos
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS sys_workflow_step_infos;
DROP SEQUENCE IF EXISTS sys_workflow_step_infos_id_seq;
CREATE SEQUENCE "sys_workflow_step_infos_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE sys_workflow_step_infos  (
  id integer PRIMARY KEY,
  created_at timestamp(0) NULL DEFAULT NULL,
  updated_at timestamp(0) NULL DEFAULT NULL,
  deleted_at timestamp(0) NULL DEFAULT NULL,
  workflow_id int  DEFAULT NULL,
  is_strat int  DEFAULT NULL,
  step_name varchar(255) DEFAULT NULL,
  step_no float  DEFAULT NULL,
  step_authority_id varchar(255) DEFAULT NULL,
  is_end int DEFAULT NULL,
  sys_workflow_id int DEFAULT NULL
) ;
alter table sys_workflow_step_infos alter column id set default nextval('sys_workflow_step_infos_id_seq');
DROP INDEX IF EXISTS idx_workflow_step_infos_deleted_at;
CREATE INDEX idx_workflow_step_infos_deleted_at ON sys_workflow_step_infos (deleted_at);
DROP INDEX IF EXISTS idx_sys_workflow_step_infos_deleted_at;
CREATE INDEX idx_sys_workflow_step_infos_deleted_at ON sys_workflow_step_infos (deleted_at);
COMMIT;
-- ----------------------------
-- Table structure for sys_workflows
-- ----------------------------
BEGIN;
DROP TABLE IF EXISTS sys_workflows;
DROP SEQUENCE IF EXISTS sys_workflows_id_seq;
CREATE SEQUENCE "sys_workflows_id_seq" INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 99999999 CACHE 1;
CREATE TABLE sys_workflows  (
  id integer PRIMARY KEY,
  created_at timestamp(0)  DEFAULT NULL,
  updated_at timestamp(0)  DEFAULT NULL,
  deleted_at timestamp(0)  DEFAULT NULL,
  workflow_nick_name varchar(255) DEFAULT NULL,
  workflow_name varchar(255) DEFAULT NULL,
  workflow_description varchar(255) DEFAULT NULL
) ;
alter table sys_workflows alter column id set default nextval('sys_workflows_id_seq');
DROP INDEX IF EXISTS idx_workflows_deleted_at;
CREATE INDEX idx_workflows_deleted_at ON sys_workflows (deleted_at);
DROP INDEX IF EXISTS idx_sys_workflows_deleted_at;
CREATE INDEX idx_sys_workflows_deleted_at ON sys_workflows (deleted_at);

-- ----------------------------
-- Records of sys_workflows
-- ----------------------------
INSERT INTO sys_workflows VALUES (1, '2019-12-09 15:20:21', '2019-12-09 15:20:21', NULL, '测试改版1', 'test', '123123');
COMMIT;
-- ----------------------------
-- View structure for authority_menu
-- ----------------------------
BEGIN;
CREATE VIEW authority_menu AS select sys_base_menus.id AS id,sys_base_menus.created_at AS created_at,sys_base_menus.updated_at AS updated_at,sys_base_menus.deleted_at AS deleted_at,sys_base_menus.menu_level AS menu_level,sys_base_menus.parent_id AS parent_id,sys_base_menus.path AS path,sys_base_menus.name AS name,sys_base_menus.hidden AS hidden,sys_base_menus.component AS component,sys_base_menus.title AS title,sys_base_menus.icon AS icon,sys_base_menus.nick_name AS nick_name,sys_base_menus.sort AS sort,sys_authority_menus.sys_authority_authority_id AS authority_id,sys_authority_menus.sys_base_menu_id AS menu_id,sys_base_menus.keep_alive AS keep_alive,sys_base_menus.default_menu AS default_menu from (sys_authority_menus join sys_base_menus on((sys_authority_menus.sys_base_menu_id = sys_base_menus.id)));
COMMIT;

