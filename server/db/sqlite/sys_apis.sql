INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
1,1569641029000,1569661576000,NULL,NULL,'/base/login','用户登录','base','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
2,1569641566000,1569661571000,NULL,NULL,'/base/register','用户注册','base','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
3,1569641621000,1576054301000,NULL,NULL,'/api/createApi','创建api','api','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
4,1569650944000,1569661559000,NULL,NULL,'/api/getApiList','获取api列表','api','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
5,1569651350000,1569661553000,NULL,NULL,'/api/getApiById','获取api详细信息','api','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
7,1569651566000,1569661544000,NULL,NULL,'/api/deleteApi','删除Api','api','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
8,1569651588000,1569661539000,NULL,NULL,'/api/updateApi','更新Api','api','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
10,1569827138000,1569827138000,NULL,NULL,'/api/getAllApis','获取所有api','api','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
11,1569828189000,1569828189000,NULL,NULL,'/authority/createAuthority','创建角色','authority','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
12,1569828213000,1569828213000,NULL,NULL,'/authority/deleteAuthority','删除角色','authority','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
13,1569828237000,1569828237000,NULL,NULL,'/authority/getAuthorityList','获取角色列表','authority','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
14,1569828260000,1569828260000,NULL,NULL,'/menu/getMenu','获取菜单树','menu','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
15,1569828290000,1569828290000,NULL,NULL,'/menu/getMenuList','分页获取基础menu列表','menu','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
16,1569828307000,1569828307000,NULL,NULL,'/menu/addBaseMenu','新增菜单','menu','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
17,1569828325000,1569828325000,NULL,NULL,'/menu/getBaseMenuTree','获取用户动态路由','menu','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
18,1569828353000,1569828353000,NULL,NULL,'/menu/addMenuAuthority','增加menu和角色关联关系','menu','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
19,1569828380000,1569828380000,NULL,NULL,'/menu/getMenuAuthority','获取指定角色menu','menu','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
20,1569828403000,1569828403000,NULL,NULL,'/menu/deleteBaseMenu','删除菜单','menu','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
21,1569828485000,1569828485000,NULL,NULL,'/menu/updateBaseMenu','更新菜单','menu','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
22,1569828501000,1569828501000,NULL,NULL,'/menu/getBaseMenuById','根据id获取菜单','menu','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
23,1569828559000,1569828559000,NULL,NULL,'/user/changePassword','修改密码','user','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
24,1569828573000,1569828573000,NULL,NULL,'/user/uploadHeaderImg','上传头像','user','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
25,1569828600000,1569828600000,NULL,NULL,'/user/getInfoList','分页获取用户列表','user','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
28,1570605317000,1570605427000,NULL,NULL,'/user/getUserList','获取用户列表','user','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
29,1570633300000,1570633300000,NULL,NULL,'/user/setUserAuthority','修改用户角色','user','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
30,1572092078000,1572092078000,NULL,NULL,'/fileUploadAndDownload/upload','文件上传示例','fileUploadAndDownload','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
31,1572092099000,1572092099000,NULL,NULL,'/fileUploadAndDownload/getFileList','获取上传文件列表','fileUploadAndDownload','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
32,1576128527000,1576128527000,NULL,NULL,'/casbin/updateCasbin','更改角色api权限','casbin','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
33,1576128539000,1576128539000,NULL,NULL,'/casbin/getPolicyPathByAuthorityId','获取权限列表','casbin','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
34,1576141335000,1576141335000,NULL,NULL,'/fileUploadAndDownload/deleteFile','删除文件','fileUploadAndDownload','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
35,1577528287000,1577528287000,NULL,NULL,'/jwt/jsonInBlacklist','jwt加入黑名单','jwt','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
36,1578304596000,1578304596000,NULL,NULL,'/authority/setDataAuthority','设置角色资源权限','authority','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
37,1578895445000,1578895445000,NULL,NULL,'/system/getSystemConfig','获取配置文件内容','system','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
38,1578898926000,1578898926000,NULL,NULL,'/system/setSystemConfig','设置配置文件内容','system','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
39,1582615959000,1582615959000,NULL,NULL,'/customer/customer','创建客户','customer','POST');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
40,1582615971000,1582616096000,NULL,NULL,'/customer/customer','更新客户','customer','PUT');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
41,1582616037000,1582616037000,NULL,NULL,'/customer/customer','删除客户','customer','DELETE');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
42,1582616208000,1582616236000,NULL,NULL,'/customer/customer','获取单一客户','customer','GET');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
43,1582616226000,1582616226000,NULL,NULL,'/customer/customerList','获取客户列表','customer','GET');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
44,1583995014000,1583996210000,NULL,NULL,'/casbin/casbinTest/:pathParam','RESTFUL模式测试','casbin','GET');
INSERT INTO sys_apis (id,created_at,updated_at,deleted_at,authority_id,"path",description,api_group,"method") VALUES (
45,1585494088000,1585494088000,NULL,NULL,'/autoCode/createTemp','自动化代码','autoCode','POST');
