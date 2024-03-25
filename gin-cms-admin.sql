/*
SQLyog Ultimate v13.1.1 (64 bit)
MySQL - 5.7.36 : Database - gin-cms-admin
*********************************************************************
*/

/*!40101 SET NAMES utf8 */;

/*!40101 SET SQL_MODE=''*/;

/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;
CREATE DATABASE /*!32312 IF NOT EXISTS*/`gin-cms-admin` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;

USE `gin-cms-admin`;

/*Table structure for table `casbin_rule` */

DROP TABLE IF EXISTS `casbin_rule`;

CREATE TABLE `casbin_rule` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `ptype` varchar(100) DEFAULT NULL,
  `v0` varchar(100) DEFAULT NULL,
  `v1` varchar(100) DEFAULT NULL,
  `v2` varchar(100) DEFAULT NULL,
  `v3` varchar(100) DEFAULT NULL,
  `v4` varchar(100) DEFAULT NULL,
  `v5` varchar(100) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_casbin_rule` (`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`)
) ENGINE=InnoDB AUTO_INCREMENT=323 DEFAULT CHARSET=utf8mb4;

/*Data for the table `casbin_rule` */

insert  into `casbin_rule`(`id`,`ptype`,`v0`,`v1`,`v2`,`v3`,`v4`,`v5`) values 
(187,'p','888','/api/createApi','POST','','',''),
(188,'p','888','/api/deleteApi','POST','','',''),
(193,'p','888','/api/deleteApisByIds','DELETE','','',''),
(191,'p','888','/api/getAllApis','POST','','',''),
(192,'p','888','/api/getApiById','POST','','',''),
(190,'p','888','/api/getApiList','POST','','',''),
(189,'p','888','/api/updateApi','POST','','',''),
(194,'p','888','/authority/copyAuthority','POST','','',''),
(195,'p','888','/authority/createAuthority','POST','','',''),
(196,'p','888','/authority/deleteAuthority','POST','','',''),
(198,'p','888','/authority/getAuthorityList','POST','','',''),
(199,'p','888','/authority/setDataAuthority','POST','','',''),
(197,'p','888','/authority/updateAuthority','PUT','','',''),
(263,'p','888','/authorityBtn/canRemoveAuthorityBtn','POST','','',''),
(262,'p','888','/authorityBtn/getAuthorityBtn','POST','','',''),
(261,'p','888','/authorityBtn/setAuthorityBtn','POST','','',''),
(235,'p','888','/autoCode/createPackage','POST','','',''),
(232,'p','888','/autoCode/createPlug','POST','','',''),
(229,'p','888','/autoCode/createTemp','POST','','',''),
(237,'p','888','/autoCode/delPackage','POST','','',''),
(241,'p','888','/autoCode/delSysHistory','POST','','',''),
(231,'p','888','/autoCode/getColumn','GET','','',''),
(227,'p','888','/autoCode/getDB','GET','','',''),
(238,'p','888','/autoCode/getMeta','POST','','',''),
(236,'p','888','/autoCode/getPackage','POST','','',''),
(240,'p','888','/autoCode/getSysHistory','POST','','',''),
(228,'p','888','/autoCode/getTables','GET','','',''),
(233,'p','888','/autoCode/installPlugin','POST','','',''),
(230,'p','888','/autoCode/preview','POST','','',''),
(234,'p','888','/autoCode/pubPlug','POST','','',''),
(239,'p','888','/autoCode/rollback','POST','','',''),
(201,'p','888','/casbin/getPolicyPathByAuthorityId','POST','','',''),
(200,'p','888','/casbin/updateCasbin','POST','','',''),
(279,'p','888','/cateMenus/createCateMenus','POST','','',''),
(280,'p','888','/cateMenus/deleteCateMenus','DELETE','','',''),
(281,'p','888','/cateMenus/deleteCateMenusByIds','DELETE','','',''),
(283,'p','888','/cateMenus/findCateMenus','GET','','',''),
(284,'p','888','/cateMenus/getCateMenusList','GET','','',''),
(286,'p','888','/cateMenus/getModelsList','GET','','',''),
(285,'p','888','/cateMenus/getTemplateList','GET','','',''),
(282,'p','888','/cateMenus/updateCateMenus','PUT','','',''),
(316,'p','888','/changeWebconfig','POST','','',''),
(293,'p','888','/class/createClass','POST','','',''),
(294,'p','888','/class/deleteClass','DELETE','','',''),
(295,'p','888','/class/deleteClassByIds','DELETE','','',''),
(297,'p','888','/class/findClass','GET','','',''),
(298,'p','888','/class/getClassList','GET','','',''),
(299,'p','888','/class/getClassListById','GET','','',''),
(296,'p','888','/class/updateClass','PUT','','',''),
(287,'p','888','/course/createCourse','POST','','',''),
(288,'p','888','/course/deleteCourse','DELETE','','',''),
(289,'p','888','/course/deleteCourseByIds','DELETE','','',''),
(291,'p','888','/course/findCourse','GET','','',''),
(292,'p','888','/course/getCourseList','GET','','',''),
(290,'p','888','/course/updateCourse','PUT','','',''),
(312,'p','888','/createWebconfig','POST','','',''),
(224,'p','888','/customer/customer','DELETE','','',''),
(225,'p','888','/customer/customer','GET','','',''),
(223,'p','888','/customer/customer','POST','','',''),
(222,'p','888','/customer/customer','PUT','','',''),
(226,'p','888','/customer/customerList','GET','','',''),
(314,'p','888','/delWebconfig','DELETE','','',''),
(260,'p','888','/email/emailTest','POST','','',''),
(212,'p','888','/fileUploadAndDownload/breakpointContinue','POST','','',''),
(213,'p','888','/fileUploadAndDownload/breakpointContinueFinish','POST','','',''),
(216,'p','888','/fileUploadAndDownload/deleteFile','POST','','',''),
(217,'p','888','/fileUploadAndDownload/editFileName','POST','','',''),
(211,'p','888','/fileUploadAndDownload/findFile','GET','','',''),
(218,'p','888','/fileUploadAndDownload/getFileList','POST','','',''),
(214,'p','888','/fileUploadAndDownload/removeChunk','POST','','',''),
(215,'p','888','/fileUploadAndDownload/upload','POST','','',''),
(313,'p','888','/getWebconfig','GET','','',''),
(176,'p','888','/jwt/jsonInBlacklist','POST','','',''),
(306,'p','888','/links/createLinks','POST','','',''),
(307,'p','888','/links/deleteLinks','DELETE','','',''),
(308,'p','888','/links/deleteLinksByIds','DELETE','','',''),
(310,'p','888','/links/findLinks','GET','','',''),
(311,'p','888','/links/getLinksList','GET','','',''),
(309,'p','888','/links/updateLinks','PUT','','',''),
(202,'p','888','/menu/addBaseMenu','POST','','',''),
(210,'p','888','/menu/addMenuAuthority','POST','','',''),
(204,'p','888','/menu/deleteBaseMenu','POST','','',''),
(206,'p','888','/menu/getBaseMenuById','POST','','',''),
(208,'p','888','/menu/getBaseMenuTree','POST','','',''),
(203,'p','888','/menu/getMenu','POST','','',''),
(209,'p','888','/menu/getMenuAuthority','POST','','',''),
(207,'p','888','/menu/getMenuList','POST','','',''),
(205,'p','888','/menu/updateBaseMenu','POST','','',''),
(300,'p','888','/message/createMessage','POST','','',''),
(301,'p','888','/message/deleteMessage','DELETE','','',''),
(302,'p','888','/message/deleteMessageByIds','DELETE','','',''),
(304,'p','888','/message/findMessage','GET','','',''),
(305,'p','888','/message/getMessageList','GET','','',''),
(303,'p','888','/message/updateMessage','PUT','','',''),
(317,'p','888','/recruit/createRecruit','POST','','',''),
(318,'p','888','/recruit/deleteRecruit','DELETE','','',''),
(319,'p','888','/recruit/deleteRecruitByIds','DELETE','','',''),
(321,'p','888','/recruit/findRecruit','GET','','',''),
(322,'p','888','/recruit/getRecruitList','GET','','',''),
(320,'p','888','/recruit/updateRecruit','PUT','','',''),
(315,'p','888','/setWebconfig','POST','','',''),
(258,'p','888','/simpleUploader/checkFileMd5','GET','','',''),
(259,'p','888','/simpleUploader/mergeFileMd5','GET','','',''),
(257,'p','888','/simpleUploader/upload','POST','','',''),
(273,'p','888','/swiper/createSwiper','POST','','',''),
(274,'p','888','/swiper/deleteSwiper','DELETE','','',''),
(275,'p','888','/swiper/deleteSwiperByIds','DELETE','','',''),
(277,'p','888','/swiper/findSwiper','GET','','',''),
(278,'p','888','/swiper/getSwiperList','GET','','',''),
(276,'p','888','/swiper/updateSwiper','PUT','','',''),
(247,'p','888','/sysDictionary/createSysDictionary','POST','','',''),
(248,'p','888','/sysDictionary/deleteSysDictionary','DELETE','','',''),
(250,'p','888','/sysDictionary/findSysDictionary','GET','','',''),
(251,'p','888','/sysDictionary/getSysDictionaryList','GET','','',''),
(249,'p','888','/sysDictionary/updateSysDictionary','PUT','','',''),
(243,'p','888','/sysDictionaryDetail/createSysDictionaryDetail','POST','','',''),
(244,'p','888','/sysDictionaryDetail/deleteSysDictionaryDetail','DELETE','','',''),
(245,'p','888','/sysDictionaryDetail/findSysDictionaryDetail','GET','','',''),
(246,'p','888','/sysDictionaryDetail/getSysDictionaryDetailList','GET','','',''),
(242,'p','888','/sysDictionaryDetail/updateSysDictionaryDetail','PUT','','',''),
(264,'p','888','/sysExportTemplate/createSysExportTemplate','POST','','',''),
(265,'p','888','/sysExportTemplate/deleteSysExportTemplate','DELETE','','',''),
(266,'p','888','/sysExportTemplate/deleteSysExportTemplateByIds','DELETE','','',''),
(270,'p','888','/sysExportTemplate/exportExcel','GET','','',''),
(271,'p','888','/sysExportTemplate/exportTemplate','GET','','',''),
(268,'p','888','/sysExportTemplate/findSysExportTemplate','GET','','',''),
(269,'p','888','/sysExportTemplate/getSysExportTemplateList','GET','','',''),
(272,'p','888','/sysExportTemplate/importExcel','POST','','',''),
(267,'p','888','/sysExportTemplate/updateSysExportTemplate','PUT','','',''),
(252,'p','888','/sysOperationRecord/createSysOperationRecord','POST','','',''),
(255,'p','888','/sysOperationRecord/deleteSysOperationRecord','DELETE','','',''),
(256,'p','888','/sysOperationRecord/deleteSysOperationRecordByIds','DELETE','','',''),
(253,'p','888','/sysOperationRecord/findSysOperationRecord','GET','','',''),
(254,'p','888','/sysOperationRecord/getSysOperationRecordList','GET','','',''),
(219,'p','888','/system/getServerInfo','POST','','',''),
(220,'p','888','/system/getSystemConfig','POST','','',''),
(221,'p','888','/system/setSystemConfig','POST','','',''),
(178,'p','888','/user/admin_register','POST','','',''),
(184,'p','888','/user/changePassword','POST','','',''),
(177,'p','888','/user/deleteUser','DELETE','','',''),
(182,'p','888','/user/getUserInfo','GET','','',''),
(179,'p','888','/user/getUserList','POST','','',''),
(186,'p','888','/user/resetPassword','POST','','',''),
(181,'p','888','/user/setSelfInfo','PUT','','',''),
(183,'p','888','/user/setUserAuthorities','POST','','',''),
(185,'p','888','/user/setUserAuthority','POST','','',''),
(180,'p','888','/user/setUserInfo','PUT','','',''),
(100,'p','8881','/api/createApi','POST','','',''),
(103,'p','8881','/api/deleteApi','POST','','',''),
(105,'p','8881','/api/getAllApis','POST','','',''),
(102,'p','8881','/api/getApiById','POST','','',''),
(101,'p','8881','/api/getApiList','POST','','',''),
(104,'p','8881','/api/updateApi','POST','','',''),
(106,'p','8881','/authority/createAuthority','POST','','',''),
(107,'p','8881','/authority/deleteAuthority','POST','','',''),
(108,'p','8881','/authority/getAuthorityList','POST','','',''),
(109,'p','8881','/authority/setDataAuthority','POST','','',''),
(127,'p','8881','/casbin/getPolicyPathByAuthorityId','POST','','',''),
(126,'p','8881','/casbin/updateCasbin','POST','','',''),
(133,'p','8881','/customer/customer','DELETE','','',''),
(134,'p','8881','/customer/customer','GET','','',''),
(131,'p','8881','/customer/customer','POST','','',''),
(132,'p','8881','/customer/customer','PUT','','',''),
(135,'p','8881','/customer/customerList','GET','','',''),
(124,'p','8881','/fileUploadAndDownload/deleteFile','POST','','',''),
(125,'p','8881','/fileUploadAndDownload/editFileName','POST','','',''),
(123,'p','8881','/fileUploadAndDownload/getFileList','POST','','',''),
(122,'p','8881','/fileUploadAndDownload/upload','POST','','',''),
(128,'p','8881','/jwt/jsonInBlacklist','POST','','',''),
(112,'p','8881','/menu/addBaseMenu','POST','','',''),
(114,'p','8881','/menu/addMenuAuthority','POST','','',''),
(116,'p','8881','/menu/deleteBaseMenu','POST','','',''),
(118,'p','8881','/menu/getBaseMenuById','POST','','',''),
(113,'p','8881','/menu/getBaseMenuTree','POST','','',''),
(110,'p','8881','/menu/getMenu','POST','','',''),
(115,'p','8881','/menu/getMenuAuthority','POST','','',''),
(111,'p','8881','/menu/getMenuList','POST','','',''),
(117,'p','8881','/menu/updateBaseMenu','POST','','',''),
(129,'p','8881','/system/getSystemConfig','POST','','',''),
(130,'p','8881','/system/setSystemConfig','POST','','',''),
(99,'p','8881','/user/admin_register','POST','','',''),
(119,'p','8881','/user/changePassword','POST','','',''),
(136,'p','8881','/user/getUserInfo','GET','','',''),
(120,'p','8881','/user/getUserList','POST','','',''),
(121,'p','8881','/user/setUserAuthority','POST','','',''),
(138,'p','9528','/api/createApi','POST','','',''),
(141,'p','9528','/api/deleteApi','POST','','',''),
(143,'p','9528','/api/getAllApis','POST','','',''),
(140,'p','9528','/api/getApiById','POST','','',''),
(139,'p','9528','/api/getApiList','POST','','',''),
(142,'p','9528','/api/updateApi','POST','','',''),
(144,'p','9528','/authority/createAuthority','POST','','',''),
(145,'p','9528','/authority/deleteAuthority','POST','','',''),
(146,'p','9528','/authority/getAuthorityList','POST','','',''),
(147,'p','9528','/authority/setDataAuthority','POST','','',''),
(174,'p','9528','/autoCode/createTemp','POST','','',''),
(165,'p','9528','/casbin/getPolicyPathByAuthorityId','POST','','',''),
(164,'p','9528','/casbin/updateCasbin','POST','','',''),
(172,'p','9528','/customer/customer','DELETE','','',''),
(170,'p','9528','/customer/customer','GET','','',''),
(171,'p','9528','/customer/customer','POST','','',''),
(169,'p','9528','/customer/customer','PUT','','',''),
(173,'p','9528','/customer/customerList','GET','','',''),
(162,'p','9528','/fileUploadAndDownload/deleteFile','POST','','',''),
(163,'p','9528','/fileUploadAndDownload/editFileName','POST','','',''),
(161,'p','9528','/fileUploadAndDownload/getFileList','POST','','',''),
(160,'p','9528','/fileUploadAndDownload/upload','POST','','',''),
(166,'p','9528','/jwt/jsonInBlacklist','POST','','',''),
(150,'p','9528','/menu/addBaseMenu','POST','','',''),
(152,'p','9528','/menu/addMenuAuthority','POST','','',''),
(154,'p','9528','/menu/deleteBaseMenu','POST','','',''),
(156,'p','9528','/menu/getBaseMenuById','POST','','',''),
(151,'p','9528','/menu/getBaseMenuTree','POST','','',''),
(148,'p','9528','/menu/getMenu','POST','','',''),
(153,'p','9528','/menu/getMenuAuthority','POST','','',''),
(149,'p','9528','/menu/getMenuList','POST','','',''),
(155,'p','9528','/menu/updateBaseMenu','POST','','',''),
(167,'p','9528','/system/getSystemConfig','POST','','',''),
(168,'p','9528','/system/setSystemConfig','POST','','',''),
(137,'p','9528','/user/admin_register','POST','','',''),
(157,'p','9528','/user/changePassword','POST','','',''),
(175,'p','9528','/user/getUserInfo','GET','','',''),
(158,'p','9528','/user/getUserList','POST','','',''),
(159,'p','9528','/user/setUserAuthority','POST','','','');

/*Table structure for table `cate_menus` */

DROP TABLE IF EXISTS `cate_menus`;

CREATE TABLE `cate_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` smallint(6) DEFAULT NULL,
  `parent_id` smallint(5) unsigned DEFAULT NULL COMMENT '父菜单ID',
  `name` varchar(191) DEFAULT NULL COMMENT '栏目名称',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否隐藏',
  `sort` mediumint(9) DEFAULT NULL COMMENT '排序标记',
  `islink` varchar(191) DEFAULT NULL COMMENT '排序标记',
  `mode_type` tinyint(1) DEFAULT NULL COMMENT '模型管理 1 产品管理 2 文章管理',
  `cate_type` varchar(191) DEFAULT NULL COMMENT '栏目类型 1 封面 2 列表 3 链接',
  `list_template` varchar(191) DEFAULT NULL COMMENT '列表页模板',
  `show_template` varchar(191) DEFAULT NULL COMMENT '内容页模板',
  `cate_thumb` varchar(191) DEFAULT NULL COMMENT '封面缩略图',
  `list_thumb` varchar(191) DEFAULT NULL COMMENT '列表缩略图',
  `desc` text COMMENT '栏目简介',
  `siteid` varchar(32) DEFAULT NULL COMMENT '站点id',
  `short` varchar(191) DEFAULT NULL COMMENT '栏目简称',
  `isjump` tinyint(1) DEFAULT NULL COMMENT '是否跳转子栏目',
  `order_type` varchar(191) DEFAULT '1' COMMENT '排序方式 1 升序 2 降序',
  `page_size` tinyint(4) DEFAULT '10' COMMENT '每页数量',
  PRIMARY KEY (`id`),
  KEY `idx_cate_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=44 DEFAULT CHARSET=utf8mb4;

/*Data for the table `cate_menus` */

insert  into `cate_menus`(`id`,`created_at`,`updated_at`,`deleted_at`,`menu_level`,`parent_id`,`name`,`hidden`,`sort`,`islink`,`mode_type`,`cate_type`,`list_template`,`show_template`,`cate_thumb`,`list_thumb`,`desc`,`siteid`,`short`,`isjump`,`order_type`,`page_size`) values 
(1,'2023-07-25 16:48:21.641','2023-10-03 10:52:19.098',NULL,0,0,'关于禹冰',0,1,'#',2,'1','category_about.html','','uploads/file/5c1be23a094f1e7863a660721fe0e2fa_20230821141351.jpg','','<p>追求卓越服务，争创一流设计企业</p>','1','About Us',1,'1',10),
(2,'2023-07-25 16:48:21.641','2023-08-29 15:34:58.530',NULL,0,0,'业务范围',0,2,'#',2,'2','products.html','','/uploadsfile/file/c3d87184fe029356303f16554fa6c258_20230829151305.jpg','','<p><span style=\"font-size: 16px;\">公司的经营范围主要为：水土保持方案编制、水土保持验收评估、水土保持小流域治理、水利工程设计（水库除险加固、河道整治、灌溉排涝、引水供水等）、农田水利规划设计、农业综合开发、农业生态工程、土地治理工程、土地整理、土地复垦、行洪论证、水资源论证、入河排污口论证、工程造价、可研报告编制。</span></p>','1','BUSINESS',1,'1',10),
(3,'2023-07-25 16:48:21.641','2023-09-07 15:38:06.815',NULL,0,0,'工程项目',0,3,'#',2,'2','category_process.html','','/uploadsfile/file/3a537b2726e92aa95124936e14813013_20230829171819.jpg','','','1','PROJECT',0,'2',10),
(4,'2023-07-25 16:48:21.641','2023-09-06 21:47:33.613',NULL,0,0,'新闻资讯',0,4,'#',2,'2','category_news.html','','/uploadsfile/file/0ea3428ffd6f10d36e176f7b1845c643_20230906214641.jpg','','<p><br></p>','1','News',0,'1',10),
(5,'2023-07-25 16:48:21.641','2023-09-20 09:29:34.565',NULL,0,0,'人力资源',0,5,'#',2,'1','list_rlzy.html','','/uploadsfile/file/8d4abdf64c4a22d9480cddb0e2f50d40_20230920092841.jpg','','','1','RESOURCES OF TALENTS',1,'1',10),
(6,'2023-07-25 16:48:21.641','2023-09-22 17:53:42.710',NULL,0,0,'联系我们',0,6,'#',2,'1','list_contact.html','','/uploadsfile/file/55ba7dad103a8190e9f71e140923b5df_20230922175338.jpg','','','1','Contact Us',0,'1',10),
(7,'2023-07-25 16:55:22.115','2023-10-03 11:06:30.159',NULL,0,1,'公司概况',0,0,'',2,'2','list_about.html','','','','<div data-w-e-type=\"video\" data-w-e-is-void=\"\"><video style=\"width: 958px; height: 479px;\" poster=\"\" controls=\"controls\" width=\"958\" height=\"479\"><source src=\"/uploadsfile/file/video(5).mp4\" type=\"video/mp4\" /></video></div>\n<p>&nbsp;</p>\n<p>&nbsp;</p>','1','gsgk',1,'1',10),
(8,'2023-07-25 16:55:22.115','2024-02-01 14:01:01.533',NULL,1,1,'董事长致辞',0,2,'',2,'2','list_about.html','','','','<p>董事长致辞简介</p>\n<table class=\"MsoTableGrid\" style=\"width: 452.9pt; margin-left: 6.75pt; margin-right: 6.75pt; border: none;\" border=\"1\" cellspacing=\"0\" align=\"center\">\n<tbody>\n<tr style=\"height: 38.3500pt;\">\n<td style=\"width: 60.8pt; padding: 0pt 5.4pt; border-width: 1pt; border-color: windowtext;\" valign=\"top\" width=\"81\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">编号</span></p>\n</td>\n<td style=\"width: 60.8pt; padding: 0pt 5.4pt; border-width: 1pt; border-color: windowtext;\" valign=\"top\" width=\"81\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">姓名</span></p>\n</td>\n<td style=\"width: 60.8pt; padding: 0pt 5.4pt; border-width: 1pt; border-color: windowtext;\" valign=\"top\" width=\"81\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">性别</span></p>\n</td>\n<td style=\"width: 183.6pt; padding: 0pt 5.4pt; border-width: 1pt; border-color: windowtext;\" valign=\"top\" width=\"244\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">职务</span></p>\n</td>\n<td style=\"width: 86.9pt; padding: 0pt 5.4pt; border-width: 1pt; border-color: windowtext;\" valign=\"top\" width=\"115\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">学历</span></p>\n</td>\n</tr>\n<tr style=\"height: 50.1000pt;\">\n<td style=\"width: 60.8pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"81\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">01</span></p>\n</td>\n<td style=\"width: 60.8pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"81\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">樊敬周</span></p>\n</td>\n<td style=\"width: 60.8pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"81\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">男</span></p>\n</td>\n<td style=\"width: 183.6pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"244\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">郑州聚瓷博物馆</span><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">监事</span></p>\n</td>\n<td style=\"width: 86.9pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"115\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">本科</span></p>\n</td>\n</tr>\n<tr style=\"height: 50.6000pt;\">\n<td style=\"width: 60.8pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"81\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">02</span></p>\n</td>\n<td style=\"width: 60.8pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"81\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">任梦真</span></p>\n</td>\n<td style=\"width: 60.8pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"81\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">女</span></p>\n</td>\n<td style=\"width: 183.6pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"244\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">郑州聚瓷博物馆</span><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">监事</span></p>\n</td>\n<td style=\"width: 86.9pt; padding: 0pt 5.4pt; border-left-width: 1pt; border-left-color: windowtext; border-right-width: 1pt; border-right-color: windowtext; border-top: none; border-bottom-width: 1pt; border-bottom-color: windowtext;\" valign=\"top\" width=\"115\">\n<p class=\"MsoNormal\" style=\"margin-left: 0.0000pt; mso-para-margin-left: 0.0000gd; text-indent: 0.0000pt; mso-char-indent-count: 0.0000; text-align: center;\" align=\"center\"><span style=\"font-family: 仿宋; font-size: 15.0000pt; mso-font-kerning: 1.0000pt;\">本科</span></p>\n</td>\n</tr>\n</tbody>\n</table>','1','',0,'1',10),
(9,'2023-07-25 16:55:22.115','2023-08-29 11:28:24.610',NULL,1,1,'资质荣誉',0,1,'',2,'1','list_about_zz.html','','/uploadsfile/file/9eebcb3aa4f0858b3ff299b53019d6fc_20230828165105.jpg','','','1','',1,'1',10),
(10,'2023-07-25 16:55:22.115','2023-10-03 11:11:33.067',NULL,0,1,'企业文化',0,0,'',2,'2','list_about_wh.html','','/uploadsfile/file/32e6d0b53071cd8176d5c7aac8433f01_20230828154411.jpg','','<p><span style=\"font-size: 24px;\">水润万物 利国益民</span></p>','1','',0,'1',10),
(11,'2023-07-25 16:56:00.881','2023-08-29 15:33:58.842',NULL,0,2,'水利市政设计',0,0,'',2,'1','products_details.html','','/uploadsfile/file/e78038471be5cfc890cd6bdf746516ee_20230829144326.jpg','','<p><span style=\"font-size: 16px;\">公司的经营范围主要为：水土保持方案编制、水土保持验收评估、水土保持小流域治理、水利工程设计（水库除险加固、河道整治、灌溉排涝、引水供水等）、农田水利规划设计、农业综合开发、农业生态工程、土地治理工程、土地整理、土地复垦、行洪论证、水资源论证、入河排污口论证、工程造价、可研报告编制。</span></p>','1','',0,'1',10),
(12,'2023-07-25 16:56:00.881','2023-08-29 15:32:48.261',NULL,0,2,'工程勘察',0,0,'',2,'1','products_details.html','','/uploadsfile/file/f4977a01004caec41d7d2f61817f4f72_20230829144403.jpg','','<p><span style=\"font-size: 16px;\">公司的经营范围主要为：水土保持方案编制、水土保持验收评估、水土保持小流域治理、水利工程设计（水库除险加固、河道整治、灌溉排涝、引水供水等）、农田水利规划设计、农业综合开发、农业生态工程、土地治理工程、土地整理、土地复垦、行洪论证、水资源论证、入河排污口论证、工程造价、可研报告编制。</span></p>','1','',0,'1',10),
(13,'2023-07-25 16:56:00.881','2023-08-29 15:34:24.659',NULL,0,2,'工程造价咨询',0,0,'',2,'1','products_details.html','','/uploadsfile/file/c05caeed821a2f2af964f596a84f67fb_20230829150422.jpg','','<p><span style=\"font-size: 16px;\">公司的经营范围主要为：水土保持方案编制、水土保持验收评估、水土保持小流域治理、水利工程设计（水库除险加固、河道整治、灌溉排涝、引水供水等）、农田水利规划设计、农业综合开发、农业生态工程、土地治理工程、土地整理、土地复垦、行洪论证、水资源论证、入河排污口论证、工程造价、可研报告编制。</span></p>','1','',0,'1',10),
(14,'2023-07-25 16:56:00.881','2023-08-29 15:34:36.257',NULL,0,2,'工程监理监测',0,0,'',2,'1','products_details.html','','/uploadsfile/file/ccbaf918a1c658d006ae423a7b60e024_20230829144625.jpg','','<p><span style=\"font-size: 16px;\">公司的经营范围主要为：水土保持方案编制、水土保持验收评估、水土保持小流域治理、水利工程设计（水库除险加固、河道整治、灌溉排涝、引水供水等）、农田水利规划设计、农业综合开发、农业生态工程、土地治理工程、土地整理、土地复垦、行洪论证、水资源论证、入河排污口论证、工程造价、可研报告编制。</span></p>','1','',0,'1',10),
(15,'2023-07-25 16:56:00.881','2023-08-29 15:34:41.948',NULL,0,2,'招标代理资质',0,0,'',2,'1','products_details.html','','/uploadsfile/file/ebfdae050f6aac7dc603ceabb937ac1b_20230829150512.jpg','','<p><span style=\"font-size: 16px;\">公司的经营范围主要为：水土保持方案编制、水土保持验收评估、水土保持小流域治理、水利工程设计（水库除险加固、河道整治、灌溉排涝、引水供水等）、农田水利规划设计、农业综合开发、农业生态工程、土地治理工程、土地整理、土地复垦、行洪论证、水资源论证、入河排污口论证、工程造价、可研报告编制。</span></p>','1','',0,'1',10),
(16,'2023-07-25 16:56:14.234','2023-09-07 15:37:03.712',NULL,0,3,'水利市政设计',0,0,'',2,'2','list_process.html','process_details.html','','','','1','',0,'2',10),
(17,'2023-07-25 16:56:14.234','2023-09-04 09:50:05.727',NULL,0,3,'工程勘察',0,0,'',2,'2','list_process.html','process_details.html','','','','1','',0,'1',10),
(18,'2023-07-25 17:00:40.299','2023-09-20 09:07:33.820',NULL,0,4,'企业新闻',0,0,'',2,'2','list_news.html','news_details.html','','','','1','Corporate News',0,'1',10),
(19,'2023-07-25 17:00:40.299','2023-09-20 09:07:53.609',NULL,0,4,'行业动态',0,0,'',2,'2','list_news.html','news_details.html','','','','1','Industry News',0,'1',10),
(20,'2023-07-25 17:01:00.345','2023-09-22 16:53:11.137',NULL,0,5,'人才招聘',0,0,'',1,'2','list_rlzy.html','','','','','1','RESOURCES OF TALENTS',0,'1',10),
(21,'2023-08-18 11:54:07.966','2023-08-28 14:17:09.307',NULL,0,0,'About Us',0,1,'#',2,'1','category_about.html','news.html','uploads/file/5c1be23a094f1e7863a660721fe0e2fa_20230821141351.jpg','','<p style=\"text-indent: 40px; text-align: start; line-height: 2;\">四川禹冰工程勘察设计有限公司，成立于2015年8月，办公及注册地位于成都市青羊区日月大道一段978号2栋1单元914-918号，已通过 IS09001质量管理体系认证、IS014001环境管理体系认证、IS045001职业健康安全管理体系认证，属国家高新技术企业。<br></p><p style=\"text-indent: 40px; text-align: start; line-height: 2;\">公司下设股东大会、监事会、战略委员会、董事会、财务中心、法务中心、审计中心、行政部、经营部、技术部、总工办，现有职工50余人，各专业齐全，其中高级工程师12人、工程师18人，助理工程师及其他20余人，注册类人员包括：水工结构师、岩土工程师、造价工程师、咨询工程师、建造师、监理工程师等。<br></p><p style=\"text-indent: 40px; text-align: start; line-height: 2;\">公司目前拥有国家建设部门颁发的：工程设计水利行业（河道整治）乙级、 工程设计水利行业丙级、工程勘察专业类（工程测量、水文地质）乙级、工程勘察专业类（岩土工程）丙级、工程造价咨询资质、工程招标代理资质；国家水土保持学会颁发的：水土保持方案编制二星；国家水利水电勘测设计协会颁发的：水资源论证乙级、水文水资源调查评价乙级；国家水利部颁发的：水利工程施工监理乙级等多项资质。</p><p style=\"text-indent: 40px; text-align: start; line-height: 2;\"><br></p><p><img src=\"/uploadsfile/file/166f1c810ac94b9b6cc49933af3e9173_20230828091153.jpg\" alt=\"img5.jpg\" data-href=\"/uploadsfile/file/166f1c810ac94b9b6cc49933af3e9173_20230828091153.jpg\" style=\"\"/></p>','2',NULL,NULL,'1',10),
(22,'2023-08-18 11:54:07.966','2023-08-24 11:19:28.807',NULL,0,0,'business scope',0,2,'#',2,'1','about.html','news.html','','','<p><img src=\"uploads/file/7e1c57ddccab6457ebb6bba09642a008_20230822163634.png\" alt=\"\" data-href=\"\" style=\"\"/></p><p><img src=\"uploads/file/fe5067706fde605fcc635835a1e52fc8_20230821181512.jpg\" alt=\"\" data-href=\"\" style=\"\"/></p><p><img src=\"uploads/file/630715656ec118c0e861fa0718d22f92_20230824105846.png\" alt=\"uploads/file/630715656ec118c0e861fa0718d22f92_20230824105846.png/\" data-href=\"\" style=\"\"/></p><p><img src=\"uploads/file/aad11223d63a90ab3c4875d708ca0027_20230823135417.jpg\" alt=\"\" data-href=\"\" style=\"\"/><img src=\"uploads/file/aad11223d63a90ab3c4875d708ca0027_20230823134753.jpg\" alt=\"\" data-href=\"\" style=\"\"/></p>','2',NULL,NULL,'1',10),
(23,'2023-08-18 11:54:07.966','2023-08-20 08:29:51.767',NULL,0,0,'Project',0,4,'#',2,'1','about.html','news.html','','','','2',NULL,NULL,'1',10),
(24,'2023-08-18 11:54:07.966','2023-08-20 08:29:32.340',NULL,0,0,'News',0,3,'#',2,'1','about.html','news.html','','','','2',NULL,NULL,'1',10),
(25,'2023-08-18 11:54:07.966','2023-08-20 08:29:57.837',NULL,0,0,'Resources',0,5,'#',2,'1','about.html','news.html','','','','2',NULL,NULL,'1',10),
(26,'2023-08-18 11:54:07.966','2023-08-20 08:30:02.306',NULL,0,0,'Contact Us',0,6,'#',2,'1','about.html','news.html','','','','2',NULL,NULL,'1',10),
(27,'2023-08-18 11:56:05.845','2023-10-08 16:58:15.907',NULL,0,21,'Company Overview',0,0,'#',2,'2','list_about.html','news_details.html','','','<p style=\"text-indent: 40px; text-align: start; line-height: 2;\">四川禹冰工程勘察设计有限公司，成立于2015年8月，办公及注册地位于成都市青羊区日月大道一段978号2栋1单元914-918号，已通过 IS09001质量管理体系认证、IS014001环境管理体系认证、IS045001职业健康安全管理体系认证，属国家高新技术企业。<br></p><p style=\"text-indent: 40px; text-align: start; line-height: 2;\">公司下设股东大会、监事会、战略委员会、董事会、财务中心、法务中心、审计中心、行政部、经营部、技术部、总工办，现有职工50余人，各专业齐全，其中高级工程师12人、工程师18人，助理工程师及其他20余人，注册类人员包括：水工结构师、岩土工程师、造价工程师、咨询工程师、建造师、监理工程师等。<br></p><p style=\"text-indent: 40px; text-align: start; line-height: 2;\">公司目前拥有国家建设部门颁发的：工程设计水利行业（河道整治）乙级、 工程设计水利行业丙级、工程勘察专业类（工程测量、水文地质）乙级、工程勘察专业类（岩土工程）丙级、工程造价咨询资质、工程招标代理资质；国家水土保持学会颁发的：水土保持方案编制二星；国家水利水电勘测设计协会颁发的：水资源论证乙级、水文水资源调查评价乙级；国家水利部颁发的：水利工程施工监理乙级等多项资质。</p><p style=\"text-indent: 40px; text-align: start; line-height: 2;\"><br></p><p><img src=\"/uploadsfile/file/166f1c810ac94b9b6cc49933af3e9173_20230828091153.jpg\" alt=\"img5.jpg\" data-href=\"/uploadsfile/file/166f1c810ac94b9b6cc49933af3e9173_20230828091153.jpg\" style=\"\"/></p>','2','',0,'1',10),
(28,'2023-08-18 11:56:05.845','2023-08-18 11:56:05.845',NULL,0,21,'Corporate Culture',0,0,'#',2,'2','news.html','news_details.html','','','','2',NULL,NULL,'1',10),
(29,'2023-08-18 11:56:05.845','2023-08-18 11:56:05.845',NULL,0,21,'Honors',0,0,'#',2,'2','news.html','news_details.html','','','','2',NULL,NULL,'1',10),
(30,'2023-08-18 11:56:05.845','2023-08-28 14:39:10.746',NULL,0,21,'Chairman\'s speech',0,0,'#',2,'2','list_about.html','news_details.html','','','<p><span style=\"color: rgb(51, 51, 51); background-color: rgb(255, 255, 255); font-size: 16px;\">四川禹冰工程勘察设计有限公司，成立于2015年8月，办公及注册地位于成都市青羊区日月大道一段978号2栋1单元914-918号，已通过 IS09001质量管理体系认证、IS014001环境管理体系认证、IS045001职业健康安全管理体系认证，属国家高新技术企业。</span></p>','2',NULL,NULL,'1',10),
(31,'2023-08-18 11:57:04.711','2023-08-18 11:57:04.711',NULL,0,22,'分类名称',0,0,'#',2,'2','news.html','news_details.html','','','','2',NULL,NULL,'1',10),
(32,'2023-08-18 11:57:04.711','2023-08-18 11:57:04.711',NULL,0,22,'分类名称',0,0,'#',2,'2','news.html','news_details.html','','','','2',NULL,NULL,'1',10),
(33,'2023-08-18 11:57:21.470','2023-08-18 11:57:21.470',NULL,0,23,'分类名称',0,0,'#',2,'2','news.html','news_details.html','','','','2',NULL,NULL,'1',10),
(34,'2023-08-18 11:57:21.470','2023-08-18 11:57:21.470',NULL,0,23,'分类名称',0,0,'#',2,'2','news.html','news_details.html','','','','2',NULL,NULL,'1',10),
(35,'2023-08-18 11:57:21.470','2023-08-18 11:57:21.470',NULL,0,23,'分类名称',0,0,'#',2,'2','news.html','news_details.html','','','','2',NULL,NULL,'1',10),
(36,'2023-08-18 11:58:07.304','2023-10-08 16:57:57.957',NULL,0,24,'Corporate News',0,0,'#',2,'2','news.html','news_details.html','','','','2','',0,'1',10),
(37,'2023-08-18 11:58:07.304','2023-10-08 16:58:04.635',NULL,0,24,'Industry News',0,0,'#',2,'2','news.html','news_details.html','','','','2','',0,'1',10),
(38,'2023-08-18 11:58:32.374','2023-10-08 17:00:33.794',NULL,0,25,'Talent recruitment',0,0,'#',1,'2','news.html','news_details.html','','','','2','',0,'1',10),
(39,'2023-08-20 08:29:10.961','2023-10-08 18:35:29.716',NULL,0,0,'Home',0,0,'http://en.yhystudy.top/',2,'3','','','','','','2','',0,'1',10),
(40,'2023-08-20 08:49:38.315','2023-10-08 18:33:09.520',NULL,0,0,'首页',0,0,'http://rxshop.yhystudy.top/',2,'3','','','','','<p><br></p>','1','',0,'1',10),
(41,'2023-08-28 18:03:43.478','2023-08-29 13:44:13.450',NULL,0,9,'企业资质',0,0,'#',2,'2','list_about_zz.html','','','','','1','',0,'1',10),
(42,'2023-08-28 18:03:43.478','2023-08-29 13:44:18.817',NULL,0,9,'工程奖项',0,0,'#',2,'2','list_about_zz.html','','','','','1','',0,'1',10),
(43,'2023-08-28 18:03:43.478','2023-08-29 13:44:24.180',NULL,0,9,'社会荣誉',0,0,'#',2,'2','list_about_zz.html','','','','','1','',0,'1',10);

/*Table structure for table `class` */

DROP TABLE IF EXISTS `class`;

CREATE TABLE `class` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `cate_id` smallint(5) unsigned DEFAULT NULL COMMENT '栏目分类id',
  `title` varchar(191) DEFAULT NULL COMMENT '标题',
  `thumb` varchar(191) DEFAULT NULL COMMENT '缩略图',
  `desc` varchar(1024) DEFAULT NULL COMMENT '介绍',
  `content` text COMMENT '介绍',
  `enable` tinyint(1) DEFAULT NULL COMMENT '状态 1显示',
  `siteid` varchar(32) DEFAULT NULL COMMENT '站点id',
  PRIMARY KEY (`id`),
  KEY `idx_class_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8mb4;

/*Data for the table `class` */

insert  into `class`(`id`,`created_at`,`updated_at`,`deleted_at`,`cate_id`,`title`,`thumb`,`desc`,`content`,`enable`,`siteid`) values 
(1,'2023-08-28 16:01:13.490','2024-02-01 14:00:43.263',NULL,10,'价值观','/uploadsfile/file/57ffc2a8786ce2bcf29cd50b93eb3099_20230828160104.jpg','','<p>顾客至上 质量第一<br />团结合作 诚实守信<br />感恩奉献 迎接变化<br />积极向上 结果为王</p>',1,'1'),
(2,'2023-08-29 13:42:52.984','2023-08-29 17:26:26.773',NULL,41,'工程勘察专业类(工程测量、水文地质）乙级、岩土工程勘察丙级','/uploadsfile/file/8e20420342b829f723c8ff0f03398299_20230829134243.jpg','','<p>hello</p>',1,'1'),
(3,'2023-08-29 13:43:31.210','2023-08-29 17:26:30.813',NULL,41,'工程勘察专业类(工程测量、水文地质）乙级、岩土工程勘察丙级','/uploadsfile/file/8e20420342b829f723c8ff0f03398299_20230829134243.jpg','','<p><br></p>',1,'1'),
(4,'2023-08-29 17:26:18.880','2023-08-29 17:26:18.880',NULL,16,'汶川县三江水库工程','/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230829172507.jpg','该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡','<p><span style=\"font-size: 12px;\">该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 </span><img src=\"/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230829172507.jpg\" alt=\"/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230829172507.jpg\" data-href=\"\" style=\"\"/></p><p><span style=\"font-size: 12px;\">该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供 </span></p>',1,'1'),
(5,'2023-08-29 17:28:32.289','2023-08-29 17:58:04.967',NULL,16,'汶川县三江水库工程2','/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230829172507.jpg','汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2','<p>汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2汶川县三江水库工程2</p><p><img src=\"/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230829172507.jpg\" alt=\"/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230829172507.jpg\" data-href=\"\" style=\"\"/></p>',1,'1'),
(6,'2023-09-06 21:58:59.536','2023-09-06 21:58:59.536',NULL,18,'汶川县三江水库工程','/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230906215740.jpg','该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³，','<p> &nbsp; &nbsp; &nbsp; &nbsp;该工程位于汶川县三江镇境内，属小（1） 型 水 库 ， 总 库 容798万m³， 兴 利 库容420万m³，该水库工程任务为城乡供水、农业灌溉、水生态开发建设、水生态综合治理、河湖康养基地打造，防洪安全等综合开发利用。工程由水库枢纽、供水工程、灌溉工程组成，工程总投资36480万元。</p><p style=\"text-align: center;\"><img src=\"/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230906215740.jpg\" alt=\"/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230906215740.jpg\" data-href=\"\" style=\"\"><br></p>',1,'1'),
(7,'2023-10-03 19:46:01.007','2023-10-08 16:09:14.821',NULL,19,'测试新闻','/uploadsfile/file/c05caeed821a2f2af964f596a84f67fb_20230829150422.jpg','6666就是六六六六六六','<p>6666就是六六六六六六</p>',1,'1'),
(8,'2023-10-08 18:05:29.346','2023-10-08 18:05:29.346',NULL,36,'news tests','','','<p>news testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews testsnews tests</p>',0,'2'),
(9,'2023-10-08 18:07:37.993','2023-10-08 18:07:37.993',NULL,37,'Industry News tests','','','<p>Industry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News testsIndustry News tests</p>',0,'2');

/*Table structure for table `course` */

DROP TABLE IF EXISTS `course`;

CREATE TABLE `course` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `cate_id` varchar(10) DEFAULT NULL COMMENT '栏目分类id',
  `title` varchar(191) DEFAULT NULL COMMENT '产品名称',
  `thumb` varchar(191) DEFAULT NULL COMMENT '缩略图',
  `product_thumbs` varchar(9999) DEFAULT NULL COMMENT '产品图片集',
  `product_details` text COMMENT '产品详细信息',
  `product_clause` text COMMENT '付款和装运条款',
  `desc` text COMMENT '产品简介',
  `created_by` varchar(191) DEFAULT NULL COMMENT '上传用户id',
  `enable` tinyint(1) DEFAULT '1' COMMENT '状态 1 显示',
  `sort` mediumint(9) DEFAULT '0' COMMENT '排序标记',
  PRIMARY KEY (`id`),
  KEY `idx_course_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `course` */

/*Table structure for table `exa_customers` */

DROP TABLE IF EXISTS `exa_customers`;

CREATE TABLE `exa_customers` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `customer_name` varchar(191) DEFAULT NULL COMMENT '客户名',
  `customer_phone_data` varchar(191) DEFAULT NULL COMMENT '客户手机号',
  `sys_user_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理ID',
  `sys_user_authority_id` bigint(20) unsigned DEFAULT NULL COMMENT '管理角色ID',
  PRIMARY KEY (`id`),
  KEY `idx_exa_customers_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `exa_customers` */

/*Table structure for table `exa_file_chunks` */

DROP TABLE IF EXISTS `exa_file_chunks`;

CREATE TABLE `exa_file_chunks` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `exa_file_id` bigint(20) unsigned DEFAULT NULL,
  `file_chunk_number` bigint(20) DEFAULT NULL,
  `file_chunk_path` varchar(191) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_chunks_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `exa_file_chunks` */

/*Table structure for table `exa_file_upload_and_downloads` */

DROP TABLE IF EXISTS `exa_file_upload_and_downloads`;

CREATE TABLE `exa_file_upload_and_downloads` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '文件名',
  `url` varchar(191) DEFAULT NULL COMMENT '文件地址',
  `tag` varchar(191) DEFAULT NULL COMMENT '文件标签',
  `key` varchar(191) DEFAULT NULL COMMENT '编号',
  PRIMARY KEY (`id`),
  KEY `idx_exa_file_upload_and_downloads_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=74 DEFAULT CHARSET=utf8mb4;

/*Data for the table `exa_file_upload_and_downloads` */

insert  into `exa_file_upload_and_downloads`(`id`,`created_at`,`updated_at`,`deleted_at`,`name`,`url`,`tag`,`key`) values 
(1,'2023-07-13 14:59:08.470','2023-07-13 14:59:08.470',NULL,'10.png','https://qmplusimg.henrongyi.top/gvalogo.png','png','158787308910.png'),
(2,'2023-07-13 14:59:08.470','2023-07-13 14:59:08.470',NULL,'logo.png','https://qmplusimg.henrongyi.top/1576554439myAvatar.png','png','1587973709logo.png'),
(3,'2023-07-25 15:59:56.112','2023-07-25 15:59:56.112',NULL,'logo.png','uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png','png','96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png'),
(4,'2023-07-25 16:00:23.720','2023-07-25 16:00:23.720',NULL,'img4.jpg','uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg','jpg','94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg'),
(5,'2023-07-27 18:09:12.535','2023-07-27 18:09:12.535',NULL,'banner1.jpg','uploads/file/7c9a778d594c025dc99213019cdf6bcf_20230727180912.jpg','jpg','7c9a778d594c025dc99213019cdf6bcf_20230727180912.jpg'),
(6,'2023-08-17 18:53:22.581','2023-08-17 18:53:22.581',NULL,'banner1.jpg','uploads/file/7c9a778d594c025dc99213019cdf6bcf_20230817185322.jpg','jpg','7c9a778d594c025dc99213019cdf6bcf_20230817185322.jpg'),
(7,'2023-08-18 10:04:40.903','2023-08-18 10:04:40.903',NULL,'wap_banner1.jpg','uploads/file/b8f144bfcd83307f42d6211740d031bb_20230818100440.jpg','jpg','b8f144bfcd83307f42d6211740d031bb_20230818100440.jpg'),
(8,'2023-08-21 14:13:51.662','2023-08-21 14:13:51.662',NULL,'n_banner1.jpg','uploads/file/5c1be23a094f1e7863a660721fe0e2fa_20230821141351.jpg','jpg','5c1be23a094f1e7863a660721fe0e2fa_20230821141351.jpg'),
(9,'2023-08-21 18:15:12.518','2023-08-21 18:15:12.518',NULL,'img2.jpg','uploads/file/fe5067706fde605fcc635835a1e52fc8_20230821181512.jpg','jpg','fe5067706fde605fcc635835a1e52fc8_20230821181512.jpg'),
(10,'2023-08-22 15:51:39.471','2023-08-22 15:51:39.471',NULL,'企业微信截图_16926847192722.png','uploads/file/15ffdc8db6797864c63d29f25206cc8f_20230822155139.png','png','15ffdc8db6797864c63d29f25206cc8f_20230822155139.png'),
(11,'2023-08-22 16:36:34.382','2023-08-22 16:36:34.382',NULL,'企业微信截图_16915789278288.png','uploads/file/7e1c57ddccab6457ebb6bba09642a008_20230822163634.png','png','7e1c57ddccab6457ebb6bba09642a008_20230822163634.png'),
(12,'2023-08-22 16:37:39.810','2023-08-22 16:37:39.810',NULL,'企业微信截图_16915788114501.png','uploads/file/b2811745e691fb161ef7b52f80800a4e_20230822163739.png','png','b2811745e691fb161ef7b52f80800a4e_20230822163739.png'),
(23,'2023-08-28 09:00:27.286','2023-08-28 09:00:27.286',NULL,'1(1).jpg','/uploadsfile/file/aad11223d63a90ab3c4875d708ca0027_20230828090027.jpg','jpg','aad11223d63a90ab3c4875d708ca0027_20230828090027.jpg'),
(24,'2023-08-28 09:09:10.928','2023-08-28 09:09:10.928',NULL,'img5.jpg','/uploadsfile/file/166f1c810ac94b9b6cc49933af3e9173_20230828090910.jpg','jpg','166f1c810ac94b9b6cc49933af3e9173_20230828090910.jpg'),
(25,'2023-08-28 09:10:46.485','2023-08-28 09:10:46.485',NULL,'img5.jpg','/uploadsfile/file/166f1c810ac94b9b6cc49933af3e9173_20230828091046.jpg','jpg','166f1c810ac94b9b6cc49933af3e9173_20230828091046.jpg'),
(26,'2023-08-28 09:11:54.034','2023-08-28 09:11:54.034',NULL,'img5.jpg','/uploadsfile/file/166f1c810ac94b9b6cc49933af3e9173_20230828091153.jpg','jpg','166f1c810ac94b9b6cc49933af3e9173_20230828091153.jpg'),
(27,'2023-08-28 15:44:11.355','2023-08-28 15:44:11.355',NULL,'n_banner4.jpg','/uploadsfile/file/32e6d0b53071cd8176d5c7aac8433f01_20230828154411.jpg','jpg','32e6d0b53071cd8176d5c7aac8433f01_20230828154411.jpg'),
(28,'2023-08-28 16:01:04.466','2023-08-28 16:01:04.466',NULL,'img6.jpg','/uploadsfile/file/57ffc2a8786ce2bcf29cd50b93eb3099_20230828160104.jpg','jpg','57ffc2a8786ce2bcf29cd50b93eb3099_20230828160104.jpg'),
(29,'2023-08-28 16:51:05.527','2023-08-28 16:51:05.527',NULL,'n_banner3.jpg','/uploadsfile/file/9eebcb3aa4f0858b3ff299b53019d6fc_20230828165105.jpg','jpg','9eebcb3aa4f0858b3ff299b53019d6fc_20230828165105.jpg'),
(30,'2023-08-29 13:42:43.051','2023-08-29 13:42:43.051',NULL,'img3.jpg','/uploadsfile/file/8e20420342b829f723c8ff0f03398299_20230829134243.jpg','jpg','8e20420342b829f723c8ff0f03398299_20230829134243.jpg'),
(31,'2023-08-29 14:43:26.063','2023-08-29 14:43:26.063',NULL,'img10.jpg','/uploadsfile/file/e78038471be5cfc890cd6bdf746516ee_20230829144326.jpg','jpg','e78038471be5cfc890cd6bdf746516ee_20230829144326.jpg'),
(32,'2023-08-29 14:44:03.715','2023-08-29 14:44:03.715',NULL,'img11.jpg','/uploadsfile/file/f4977a01004caec41d7d2f61817f4f72_20230829144403.jpg','jpg','f4977a01004caec41d7d2f61817f4f72_20230829144403.jpg'),
(33,'2023-08-29 14:46:25.578','2023-08-29 14:46:25.578',NULL,'img12.jpg','/uploadsfile/file/ccbaf918a1c658d006ae423a7b60e024_20230829144625.jpg','jpg','ccbaf918a1c658d006ae423a7b60e024_20230829144625.jpg'),
(34,'2023-08-29 15:04:22.237','2023-08-29 15:04:22.237',NULL,'img13.jpg','/uploadsfile/file/c05caeed821a2f2af964f596a84f67fb_20230829150422.jpg','jpg','c05caeed821a2f2af964f596a84f67fb_20230829150422.jpg'),
(35,'2023-08-29 15:05:12.521','2023-08-29 15:05:12.521',NULL,'img14.jpg','/uploadsfile/file/ebfdae050f6aac7dc603ceabb937ac1b_20230829150512.jpg','jpg','ebfdae050f6aac7dc603ceabb937ac1b_20230829150512.jpg'),
(36,'2023-08-29 15:13:05.390','2023-08-29 15:13:05.390',NULL,'n_banner5.jpg','/uploadsfile/file/c3d87184fe029356303f16554fa6c258_20230829151305.jpg','jpg','c3d87184fe029356303f16554fa6c258_20230829151305.jpg'),
(37,'2023-08-29 17:18:19.745','2023-08-29 17:18:19.745',NULL,'n_banner6.jpg','/uploadsfile/file/3a537b2726e92aa95124936e14813013_20230829171819.jpg','jpg','3a537b2726e92aa95124936e14813013_20230829171819.jpg'),
(38,'2023-08-29 17:25:07.951','2023-08-29 17:25:07.951',NULL,'img2.jpg','/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230829172507.jpg','jpg','fe5067706fde605fcc635835a1e52fc8_20230829172507.jpg'),
(61,'2023-09-05 13:47:13.710','2023-09-05 13:47:13.710',NULL,'1692257179535.jpg','/uploadsfile/file/567307944b2a26084f11b6d2347d10e5_20230905134713.jpg','jpg','567307944b2a26084f11b6d2347d10e5_20230905134713.jpg'),
(66,'2023-09-05 17:42:52.555','2023-09-05 17:42:52.555',NULL,'video(5).mp4','/uploadsfile/file/video(5).mp4','mp4','video(5).mp4'),
(67,'2023-09-06 21:46:41.881','2023-09-06 21:46:41.881',NULL,'n_banner8.jpg','/uploadsfile/file/0ea3428ffd6f10d36e176f7b1845c643_20230906214641.jpg','jpg','0ea3428ffd6f10d36e176f7b1845c643_20230906214641.jpg'),
(68,'2023-09-06 21:57:40.127','2023-09-06 21:57:40.127',NULL,'img2.jpg','/uploadsfile/file/fe5067706fde605fcc635835a1e52fc8_20230906215740.jpg','jpg','fe5067706fde605fcc635835a1e52fc8_20230906215740.jpg'),
(69,'2023-09-20 09:28:41.213','2023-09-20 09:28:41.213',NULL,'n_banner9.jpg','/uploadsfile/file/8d4abdf64c4a22d9480cddb0e2f50d40_20230920092841.jpg','jpg','8d4abdf64c4a22d9480cddb0e2f50d40_20230920092841.jpg'),
(70,'2023-09-22 17:53:38.388','2023-09-22 17:53:38.388',NULL,'n_banner10.jpg','/uploadsfile/file/55ba7dad103a8190e9f71e140923b5df_20230922175338.jpg','jpg','55ba7dad103a8190e9f71e140923b5df_20230922175338.jpg');

/*Table structure for table `exa_files` */

DROP TABLE IF EXISTS `exa_files`;

CREATE TABLE `exa_files` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `file_name` varchar(191) DEFAULT NULL,
  `file_md5` varchar(191) DEFAULT NULL,
  `file_path` varchar(191) DEFAULT NULL,
  `chunk_total` bigint(20) DEFAULT NULL,
  `is_finish` tinyint(1) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_exa_files_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `exa_files` */

/*Table structure for table `jwt_blacklists` */

DROP TABLE IF EXISTS `jwt_blacklists`;

CREATE TABLE `jwt_blacklists` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `jwt` text COMMENT 'jwt',
  PRIMARY KEY (`id`),
  KEY `idx_jwt_blacklists_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `jwt_blacklists` */

/*Table structure for table `links` */

DROP TABLE IF EXISTS `links`;

CREATE TABLE `links` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL COMMENT '名称',
  `url` varchar(32) DEFAULT NULL COMMENT '跳转链接',
  `thumb` varchar(191) DEFAULT NULL COMMENT '网站logo',
  `type` enum('文字','图片') DEFAULT NULL COMMENT '类型 1 文字 2 图片',
  `enable` tinyint(1) DEFAULT NULL COMMENT '用户是否显示 1正常 2冻结',
  PRIMARY KEY (`id`),
  KEY `idx_links_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `links` */

/*Table structure for table `message` */

DROP TABLE IF EXISTS `message`;

CREATE TABLE `message` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(32) DEFAULT NULL COMMENT '姓名',
  `email` varchar(32) DEFAULT NULL COMMENT '邮箱',
  `tel` varchar(32) DEFAULT NULL COMMENT '电话',
  `country` varchar(32) DEFAULT NULL COMMENT '国家',
  `company` varchar(32) DEFAULT NULL COMMENT '公司',
  `content` varchar(191) DEFAULT NULL COMMENT '内容',
  PRIMARY KEY (`id`),
  KEY `idx_message_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `message` */

/*Table structure for table `recruit` */

DROP TABLE IF EXISTS `recruit`;

CREATE TABLE `recruit` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `job` varchar(32) DEFAULT NULL,
  `count` varchar(4) DEFAULT NULL,
  `address` varchar(191) DEFAULT NULL,
  `content` text,
  `enable` tinyint(1) DEFAULT NULL COMMENT '状态 1显示',
  `cate_id` smallint(5) unsigned DEFAULT NULL COMMENT '栏目分类id',
  `siteid` varchar(32) DEFAULT NULL COMMENT '站点id',
  PRIMARY KEY (`id`),
  KEY `idx_recruit_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

/*Data for the table `recruit` */

insert  into `recruit`(`id`,`created_at`,`updated_at`,`deleted_at`,`job`,`count`,`address`,`content`,`enable`,`cate_id`,`siteid`) values 
(1,'2023-09-22 14:59:39.794','2023-09-22 16:21:24.820',NULL,'跟单专员','2人','成都','<p>岗位职责：<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;1.负责分析项目具体业务，参与测试需求沟通与分析；<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;2.组织搭建测试环境，设计并执行测试用例；<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;3.执行测试，报告软件问题；<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;4.总结软件测试工作，编写测试报告；<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;5.与测试业务相关的临时工作。<br>任职要求：<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;1.本科及以上学历，5年以上软件测试工作经验；<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;2.熟悉测试体系，能独立分析测试需求，设计测试数据及开发测试脚本，能编写测试用例，能对软件质量进行客观评价；<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;3.熟悉常用的数据库操作方法，至少掌握一种脚本语言；<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;4.熟悉软件测试流程，熟练掌握常用的软件测试方法，了解一些软件测试理论、标准；<br> &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp; &nbsp;5.工作耐心、严谨、主动，上进心强；</p>',1,20,NULL),
(2,'2023-09-22 16:49:36.763','2023-09-22 16:52:55.556',NULL,'golang  工程师','1人','北京','<p>暂无</p>',1,20,NULL),
(3,'2023-10-03 19:44:59.870','2023-10-03 19:44:59.870',NULL,'测试9527','10','成都','<p>钱钱钱钱钱</p>',1,20,NULL);

/*Table structure for table `swiper` */

DROP TABLE IF EXISTS `swiper`;

CREATE TABLE `swiper` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `title` varchar(32) DEFAULT NULL COMMENT '标题',
  `desc` varchar(255) DEFAULT NULL COMMENT '描述',
  `thumb` varchar(255) DEFAULT NULL COMMENT '轮播图',
  `wap_thumb` varchar(255) DEFAULT NULL COMMENT 'wap轮播图',
  `url` varchar(255) DEFAULT NULL COMMENT '跳转链接',
  `enable` tinyint(1) DEFAULT NULL COMMENT '是否显示 1是 2否',
  `siteid` varchar(32) DEFAULT NULL COMMENT '站点id',
  PRIMARY KEY (`id`),
  KEY `idx_swiper_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `swiper` */

insert  into `swiper`(`id`,`created_at`,`updated_at`,`deleted_at`,`title`,`desc`,`thumb`,`wap_thumb`,`url`,`enable`,`siteid`) values 
(1,'2023-07-27 18:09:21.452','2023-09-04 09:32:31.873',NULL,'中文站点1','中文站点1','uploads/file/7c9a778d594c025dc99213019cdf6bcf_20230727180912.jpg','uploads/file/b8f144bfcd83307f42d6211740d031bb_20230818100440.jpg','',1,'1'),
(2,'2023-07-27 18:14:29.089','2023-08-18 10:04:44.451',NULL,'英文站测试1','英文站测试1','uploads/file/7c9a778d594c025dc99213019cdf6bcf_20230817185322.jpg','uploads/file/b8f144bfcd83307f42d6211740d031bb_20230818100440.jpg','',1,'2');

/*Table structure for table `sys_apis` */

DROP TABLE IF EXISTS `sys_apis`;

CREATE TABLE `sys_apis` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `path` varchar(191) DEFAULT NULL COMMENT 'api路径',
  `description` varchar(191) DEFAULT NULL COMMENT 'api中文描述',
  `api_group` varchar(191) DEFAULT NULL COMMENT 'api组',
  `method` varchar(191) DEFAULT 'POST' COMMENT '方法',
  PRIMARY KEY (`id`),
  KEY `idx_sys_apis_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=163 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_apis` */

insert  into `sys_apis`(`id`,`created_at`,`updated_at`,`deleted_at`,`path`,`description`,`api_group`,`method`) values 
(1,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/jwt/jsonInBlacklist','jwt加入黑名单(退出，必选)','jwt','POST'),
(2,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/deleteUser','删除用户','系统用户','DELETE'),
(3,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/admin_register','用户注册','系统用户','POST'),
(4,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/getUserList','获取用户列表','系统用户','POST'),
(5,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/setUserInfo','设置用户信息','系统用户','PUT'),
(6,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/setSelfInfo','设置自身信息(必选)','系统用户','PUT'),
(7,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/getUserInfo','获取自身信息(必选)','系统用户','GET'),
(8,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/setUserAuthorities','设置权限组','系统用户','POST'),
(9,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/changePassword','修改密码（建议选择)','系统用户','POST'),
(10,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/setUserAuthority','修改用户角色(必选)','系统用户','POST'),
(11,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/user/resetPassword','重置用户密码','系统用户','POST'),
(12,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/api/createApi','创建api','api','POST'),
(13,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/api/deleteApi','删除Api','api','POST'),
(14,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/api/updateApi','更新Api','api','POST'),
(15,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/api/getApiList','获取api列表','api','POST'),
(16,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/api/getAllApis','获取所有api','api','POST'),
(17,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/api/getApiById','获取api详细信息','api','POST'),
(18,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/api/deleteApisByIds','批量删除api','api','DELETE'),
(19,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/authority/copyAuthority','拷贝角色','角色','POST'),
(20,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/authority/createAuthority','创建角色','角色','POST'),
(21,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/authority/deleteAuthority','删除角色','角色','POST'),
(22,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/authority/updateAuthority','更新角色信息','角色','PUT'),
(23,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/authority/getAuthorityList','获取角色列表','角色','POST'),
(24,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/authority/setDataAuthority','设置角色资源权限','角色','POST'),
(25,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/casbin/updateCasbin','更改角色api权限','casbin','POST'),
(26,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/casbin/getPolicyPathByAuthorityId','获取权限列表','casbin','POST'),
(27,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/menu/addBaseMenu','新增菜单','菜单','POST'),
(28,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/menu/getMenu','获取菜单树(必选)','菜单','POST'),
(29,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/menu/deleteBaseMenu','删除菜单','菜单','POST'),
(30,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/menu/updateBaseMenu','更新菜单','菜单','POST'),
(31,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/menu/getBaseMenuById','根据id获取菜单','菜单','POST'),
(32,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/menu/getMenuList','分页获取基础menu列表','菜单','POST'),
(33,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/menu/getBaseMenuTree','获取用户动态路由','菜单','POST'),
(34,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/menu/getMenuAuthority','获取指定角色menu','菜单','POST'),
(35,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/menu/addMenuAuthority','增加menu和角色关联关系','菜单','POST'),
(36,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/fileUploadAndDownload/findFile','寻找目标文件（秒传）','分片上传','GET'),
(37,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/fileUploadAndDownload/breakpointContinue','断点续传','分片上传','POST'),
(38,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/fileUploadAndDownload/breakpointContinueFinish','断点续传完成','分片上传','POST'),
(39,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/fileUploadAndDownload/removeChunk','上传完成移除文件','分片上传','POST'),
(40,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/fileUploadAndDownload/upload','文件上传示例','文件上传与下载','POST'),
(41,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/fileUploadAndDownload/deleteFile','删除文件','文件上传与下载','POST'),
(42,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/fileUploadAndDownload/editFileName','文件名或者备注编辑','文件上传与下载','POST'),
(43,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/fileUploadAndDownload/getFileList','获取上传文件列表','文件上传与下载','POST'),
(44,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/system/getServerInfo','获取服务器信息','系统服务','POST'),
(45,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/system/getSystemConfig','获取配置文件内容','系统服务','POST'),
(46,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/system/setSystemConfig','设置配置文件内容','系统服务','POST'),
(47,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/customer/customer','更新客户','客户','PUT'),
(48,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/customer/customer','创建客户','客户','POST'),
(49,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/customer/customer','删除客户','客户','DELETE'),
(50,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/customer/customer','获取单一客户','客户','GET'),
(51,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/customer/customerList','获取客户列表','客户','GET'),
(52,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/getDB','获取所有数据库','代码生成器','GET'),
(53,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/getTables','获取数据库表','代码生成器','GET'),
(54,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/createTemp','自动化代码','代码生成器','POST'),
(55,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/preview','预览自动化代码','代码生成器','POST'),
(56,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/getColumn','获取所选table的所有字段','代码生成器','GET'),
(57,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/createPlug','自动创建插件包','代码生成器','POST'),
(58,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/installPlugin','安装插件','代码生成器','POST'),
(59,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/pubPlug','打包插件','代码生成器','POST'),
(60,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/createPackage','生成包(package)','包（pkg）生成器','POST'),
(61,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/getPackage','获取所有包(package)','包（pkg）生成器','POST'),
(62,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/delPackage','删除包(package)','包（pkg）生成器','POST'),
(63,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/getMeta','获取meta信息','代码生成器历史','POST'),
(64,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/rollback','回滚自动生成代码','代码生成器历史','POST'),
(65,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/getSysHistory','查询回滚记录','代码生成器历史','POST'),
(66,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/autoCode/delSysHistory','删除回滚记录','代码生成器历史','POST'),
(67,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionaryDetail/updateSysDictionaryDetail','更新字典内容','系统字典详情','PUT'),
(68,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionaryDetail/createSysDictionaryDetail','新增字典内容','系统字典详情','POST'),
(69,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionaryDetail/deleteSysDictionaryDetail','删除字典内容','系统字典详情','DELETE'),
(70,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionaryDetail/findSysDictionaryDetail','根据ID获取字典内容','系统字典详情','GET'),
(71,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionaryDetail/getSysDictionaryDetailList','获取字典内容列表','系统字典详情','GET'),
(72,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionary/createSysDictionary','新增字典','系统字典','POST'),
(73,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionary/deleteSysDictionary','删除字典','系统字典','DELETE'),
(74,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionary/updateSysDictionary','更新字典','系统字典','PUT'),
(75,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionary/findSysDictionary','根据ID获取字典','系统字典','GET'),
(76,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysDictionary/getSysDictionaryList','获取字典列表','系统字典','GET'),
(77,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysOperationRecord/createSysOperationRecord','新增操作记录','操作记录','POST'),
(78,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysOperationRecord/findSysOperationRecord','根据ID获取操作记录','操作记录','GET'),
(79,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysOperationRecord/getSysOperationRecordList','获取操作记录列表','操作记录','GET'),
(80,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysOperationRecord/deleteSysOperationRecord','删除操作记录','操作记录','DELETE'),
(81,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysOperationRecord/deleteSysOperationRecordByIds','批量删除操作历史','操作记录','DELETE'),
(82,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/simpleUploader/upload','插件版分片上传','断点续传(插件版)','POST'),
(83,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/simpleUploader/checkFileMd5','文件完整度验证','断点续传(插件版)','GET'),
(84,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/simpleUploader/mergeFileMd5','上传完成合并文件','断点续传(插件版)','GET'),
(85,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/email/emailTest','发送测试邮件','email','POST'),
(86,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/email/emailSend','发送邮件示例','email','POST'),
(87,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/authorityBtn/setAuthorityBtn','设置按钮权限','按钮权限','POST'),
(88,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/authorityBtn/getAuthorityBtn','获取已有按钮权限','按钮权限','POST'),
(89,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/authorityBtn/canRemoveAuthorityBtn','删除按钮','按钮权限','POST'),
(90,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysExportTemplate/createSysExportTemplate','新增导出模板','表格模板','POST'),
(91,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysExportTemplate/deleteSysExportTemplate','删除导出模板','表格模板','DELETE'),
(92,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysExportTemplate/deleteSysExportTemplateByIds','批量删除导出模板','表格模板','DELETE'),
(93,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysExportTemplate/updateSysExportTemplate','更新导出模板','表格模板','PUT'),
(94,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysExportTemplate/findSysExportTemplate','根据ID获取导出模板','表格模板','GET'),
(95,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysExportTemplate/getSysExportTemplateList','获取导出模板列表','表格模板','GET'),
(96,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysExportTemplate/exportExcel','导出Excel','表格模板','GET'),
(97,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysExportTemplate/exportTemplate','下载模板','表格模板','GET'),
(98,'2024-03-20 11:40:39.096','2024-03-20 11:40:39.096',NULL,'/sysExportTemplate/importExcel','导入Excel','表格模板','POST'),
(99,'2023-07-13 15:33:05.422','2023-07-13 15:33:05.422',NULL,'/system/reloadSystem','重启服务','系统服务','POST'),
(100,'2023-02-20 14:02:11.489','2023-02-20 14:02:11.489',NULL,'/swiper/createSwiper','新增swiper表','swiper','POST'),
(101,'2023-02-20 14:02:11.648','2023-02-20 14:02:11.648',NULL,'/swiper/deleteSwiper','删除swiper表','swiper','DELETE'),
(102,'2023-02-20 14:02:11.817','2023-02-20 14:02:11.817',NULL,'/swiper/deleteSwiperByIds','批量删除swiper表','swiper','DELETE'),
(103,'2023-02-20 14:02:11.985','2023-02-20 14:02:11.985',NULL,'/swiper/updateSwiper','更新swiper表','swiper','PUT'),
(104,'2023-02-20 14:02:12.146','2023-02-20 14:02:12.146',NULL,'/swiper/findSwiper','根据ID获取swiper表','swiper','GET'),
(105,'2023-02-20 14:02:12.309','2023-02-20 14:02:12.309',NULL,'/swiper/getSwiperList','获取swiper表列表','swiper','GET'),
(106,'2023-02-22 15:31:15.909','2023-02-22 15:31:15.909',NULL,'/cateMenus/createCateMenus','新增cateMenus表','cateMenus','POST'),
(107,'2023-02-22 15:31:16.071','2023-02-22 15:31:16.071',NULL,'/cateMenus/deleteCateMenus','删除cateMenus表','cateMenus','DELETE'),
(108,'2023-02-22 15:31:16.233','2023-02-22 15:31:16.233',NULL,'/cateMenus/deleteCateMenusByIds','批量删除cateMenus表','cateMenus','DELETE'),
(109,'2023-02-22 15:31:16.394','2023-02-22 15:31:16.394',NULL,'/cateMenus/updateCateMenus','更新cateMenus表','cateMenus','PUT'),
(110,'2023-02-22 15:31:16.554','2023-02-22 15:31:16.554',NULL,'/cateMenus/findCateMenus','根据ID获取cateMenus表','cateMenus','GET'),
(111,'2023-02-22 15:31:16.715','2023-02-22 15:31:16.715',NULL,'/cateMenus/getCateMenusList','获取cateMenus表列表','cateMenus','GET'),
(112,'2023-02-23 10:21:54.492','2023-02-23 10:21:54.492',NULL,'/course/createCourse','新增course表','course','POST'),
(113,'2023-02-23 10:21:54.648','2023-02-23 10:21:54.648',NULL,'/course/deleteCourse','删除course表','course','DELETE'),
(114,'2023-02-23 10:21:54.816','2023-02-23 10:21:54.816',NULL,'/course/deleteCourseByIds','批量删除course表','course','DELETE'),
(115,'2023-02-23 10:21:54.972','2023-02-23 10:21:54.972',NULL,'/course/updateCourse','更新course表','course','PUT'),
(116,'2023-02-23 10:21:55.128','2023-02-23 10:21:55.128',NULL,'/course/findCourse','根据ID获取course表','course','GET'),
(117,'2023-02-23 10:21:55.289','2023-02-23 10:21:55.289',NULL,'/course/getCourseList','获取course表列表','course','GET'),
(118,'2023-02-25 15:30:31.140','2023-02-25 15:30:31.140',NULL,'/class/createClass','新增class表','class','POST'),
(119,'2023-02-25 15:30:31.313','2023-02-25 15:30:31.313',NULL,'/class/deleteClass','删除class表','class','DELETE'),
(120,'2023-02-25 15:30:31.486','2023-02-25 15:30:31.486',NULL,'/class/deleteClassByIds','批量删除class表','class','DELETE'),
(121,'2023-02-25 15:30:31.660','2023-02-25 15:30:31.660',NULL,'/class/updateClass','更新class表','class','PUT'),
(122,'2023-02-25 15:30:31.833','2023-02-25 15:30:31.833',NULL,'/class/findClass','根据ID获取class表','class','GET'),
(123,'2023-02-25 15:30:32.006','2023-02-25 15:30:32.006',NULL,'/class/getClassList','获取class表列表','class','GET'),
(124,'2023-02-27 16:02:10.352','2023-02-27 16:02:10.352','2023-03-05 09:33:31.235','/course/getCourseIdName','获取course表id name','course','GET'),
(125,'2023-02-27 18:02:21.461','2023-02-27 18:02:21.461','2023-02-28 22:25:26.564','/teacher/createTeacher','新增teacher表','teacher','POST'),
(126,'2023-02-27 18:02:21.621','2023-02-27 18:02:21.621','2023-02-28 22:25:26.564','/teacher/deleteTeacher','删除teacher表','teacher','DELETE'),
(127,'2023-02-27 18:02:21.778','2023-02-27 18:02:21.778','2023-02-28 22:25:26.564','/teacher/deleteTeacherByIds','批量删除teacher表','teacher','DELETE'),
(128,'2023-02-27 18:02:21.937','2023-02-27 18:02:21.937','2023-02-28 22:25:26.564','/teacher/updateTeacher','更新teacher表','teacher','PUT'),
(129,'2023-02-27 18:02:22.095','2023-02-27 18:02:22.095','2023-02-28 22:25:26.564','/teacher/findTeacher','根据ID获取teacher表','teacher','GET'),
(130,'2023-02-27 18:02:22.251','2023-02-27 18:02:22.251','2023-02-28 22:25:26.564','/teacher/getTeacherList','获取teacher表列表','teacher','GET'),
(131,'2023-02-28 18:13:24.553','2023-02-28 18:13:24.553',NULL,'/class/getClassListById','通过课程id获取Class列表','class','GET'),
(132,'2023-03-03 10:45:13.243','2023-03-03 11:55:31.287',NULL,'/cateMenus/getTemplateList','获取模板名称列表','cateMenus','GET'),
(133,'2023-03-11 15:12:21.119','2023-03-11 15:12:21.119','2023-07-14 13:55:56.513','/courseEmail/createCourseEmail','新增courseEmail表','courseEmail','POST'),
(134,'2023-03-11 15:12:21.293','2023-03-11 15:12:21.293','2023-07-14 13:55:56.513','/courseEmail/deleteCourseEmail','删除courseEmail表','courseEmail','DELETE'),
(135,'2023-03-11 15:12:21.468','2023-03-11 15:12:21.468','2023-07-14 13:55:56.513','/courseEmail/deleteCourseEmailByIds','批量删除courseEmail表','courseEmail','DELETE'),
(136,'2023-03-11 15:12:21.642','2023-03-11 15:12:21.642','2023-07-14 13:55:56.513','/courseEmail/updateCourseEmail','更新courseEmail表','courseEmail','PUT'),
(137,'2023-03-11 15:12:21.816','2023-03-11 15:12:21.816','2023-07-14 13:55:56.513','/courseEmail/findCourseEmail','根据ID获取courseEmail表','courseEmail','GET'),
(138,'2023-03-11 15:12:21.990','2023-03-11 15:12:21.990','2023-07-14 13:55:56.513','/courseEmail/getCourseEmailList','获取courseEmail表列表','courseEmail','GET'),
(139,'2023-03-13 15:35:00.394','2023-03-13 15:35:00.394',NULL,'/message/createMessage','新增message表','message','POST'),
(140,'2023-03-13 15:35:00.531','2023-03-13 15:35:00.531',NULL,'/message/deleteMessage','删除message表','message','DELETE'),
(141,'2023-03-13 15:35:00.667','2023-03-13 15:35:00.667',NULL,'/message/deleteMessageByIds','批量删除message表','message','DELETE'),
(142,'2023-03-13 15:35:00.807','2023-03-13 15:35:00.807',NULL,'/message/updateMessage','更新message表','message','PUT'),
(143,'2023-03-13 15:35:00.945','2023-03-13 15:35:00.945',NULL,'/message/findMessage','根据ID获取message表','message','GET'),
(144,'2023-03-13 15:35:01.088','2023-03-13 15:35:01.088',NULL,'/message/getMessageList','获取message表列表','message','GET'),
(145,'2023-03-13 18:35:37.656','2023-03-13 18:35:37.656',NULL,'/links/createLinks','新增links表','links','POST'),
(146,'2023-03-13 18:35:37.809','2023-03-13 18:35:37.809',NULL,'/links/deleteLinks','删除links表','links','DELETE'),
(147,'2023-03-13 18:35:37.963','2023-03-13 18:35:37.963',NULL,'/links/deleteLinksByIds','批量删除links表','links','DELETE'),
(148,'2023-03-13 18:35:38.133','2023-03-13 18:35:38.133',NULL,'/links/updateLinks','更新links表','links','PUT'),
(149,'2023-03-13 18:35:38.284','2023-03-13 18:35:38.284',NULL,'/links/findLinks','根据ID获取links表','links','GET'),
(150,'2023-03-13 18:35:38.435','2023-03-13 18:35:38.435',NULL,'/links/getLinksList','获取links表列表','links','GET'),
(151,'2023-07-27 16:49:19.437','2023-07-27 16:49:19.437',NULL,'/createWebconfig','网站配置','webconfig','POST'),
(152,'2023-07-27 16:49:55.252','2023-07-27 16:49:55.252',NULL,'/getWebconfig','获取站点配置','webconfig','GET'),
(153,'2023-07-27 16:50:25.835','2023-07-27 16:50:25.835',NULL,'/delWebconfig','删除网站配置','webconfig','DELETE'),
(154,'2023-07-27 16:50:52.415','2023-07-27 16:50:52.415',NULL,'/setWebconfig','设置网站配置','webconfig','POST'),
(155,'2023-07-27 16:51:26.452','2023-07-27 16:51:26.452',NULL,'/changeWebconfig','切换站点','webconfig','POST'),
(156,'2023-09-20 10:11:11.544','2023-09-20 10:11:11.544',NULL,'/recruit/createRecruit','新增招聘模型','recruit','POST'),
(157,'2023-09-20 10:11:11.690','2023-09-20 10:11:11.690',NULL,'/recruit/deleteRecruit','删除招聘模型','recruit','DELETE'),
(158,'2023-09-20 10:11:11.835','2023-09-20 10:11:11.835',NULL,'/recruit/deleteRecruitByIds','批量删除招聘模型','recruit','DELETE'),
(159,'2023-09-20 10:11:11.980','2023-09-20 10:11:11.980',NULL,'/recruit/updateRecruit','更新招聘模型','recruit','PUT'),
(160,'2023-09-20 10:11:12.126','2023-09-20 10:11:12.126',NULL,'/recruit/findRecruit','根据ID获取招聘模型','recruit','GET'),
(161,'2023-09-20 10:11:12.271','2023-09-20 10:11:12.271',NULL,'/recruit/getRecruitList','获取招聘模型列表','recruit','GET'),
(162,'2023-09-20 11:34:03.907','2023-09-20 11:34:03.907',NULL,'/cateMenus/getModelsList','获取模型列表','cateMenus','GET');

/*Table structure for table `sys_authorities` */

DROP TABLE IF EXISTS `sys_authorities`;

CREATE TABLE `sys_authorities` (
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `authority_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT COMMENT '角色ID',
  `authority_name` varchar(191) DEFAULT NULL COMMENT '角色名',
  `parent_id` bigint(20) unsigned DEFAULT NULL COMMENT '父角色ID',
  `default_router` varchar(191) DEFAULT 'dashboard' COMMENT '默认菜单',
  PRIMARY KEY (`authority_id`),
  UNIQUE KEY `authority_id` (`authority_id`)
) ENGINE=InnoDB AUTO_INCREMENT=9529 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_authorities` */

insert  into `sys_authorities`(`created_at`,`updated_at`,`deleted_at`,`authority_id`,`authority_name`,`parent_id`,`default_router`) values 
('2024-03-20 11:40:39.294','2024-03-21 11:45:25.514',NULL,888,'普通用户',0,'dashboard'),
('2024-03-20 11:40:39.294','2024-03-20 11:40:40.915',NULL,8881,'普通用户子角色',888,'dashboard'),
('2024-03-20 11:40:39.294','2024-03-20 11:40:40.720',NULL,9528,'测试角色',0,'dashboard');

/*Table structure for table `sys_authority_btns` */

DROP TABLE IF EXISTS `sys_authority_btns`;

CREATE TABLE `sys_authority_btns` (
  `authority_id` bigint(20) unsigned DEFAULT NULL COMMENT '角色ID',
  `sys_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  `sys_base_menu_btn_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单按钮ID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_authority_btns` */

/*Table structure for table `sys_authority_menus` */

DROP TABLE IF EXISTS `sys_authority_menus`;

CREATE TABLE `sys_authority_menus` (
  `sys_base_menu_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_base_menu_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_authority_menus` */

insert  into `sys_authority_menus`(`sys_base_menu_id`,`sys_authority_authority_id`) values 
(1,888),
(1,8881),
(1,9528),
(2,888),
(2,8881),
(2,9528),
(3,888),
(4,888),
(4,8881),
(5,888),
(5,8881),
(6,888),
(6,8881),
(7,888),
(7,8881),
(8,888),
(8,8881),
(8,9528),
(9,888),
(9,8881),
(10,888),
(10,8881),
(11,888),
(11,8881),
(12,888),
(13,888),
(13,8881),
(14,888),
(14,8881),
(15,888),
(15,8881),
(16,888),
(16,8881),
(17,888),
(17,8881),
(18,888),
(19,888),
(20,888),
(21,888),
(22,888),
(23,888),
(30,888),
(31,888),
(32,888),
(33,888),
(34,888),
(35,888),
(36,888),
(37,888);

/*Table structure for table `sys_auto_code_histories` */

DROP TABLE IF EXISTS `sys_auto_code_histories`;

CREATE TABLE `sys_auto_code_histories` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package` varchar(191) DEFAULT NULL,
  `business_db` varchar(191) DEFAULT NULL,
  `table_name` varchar(191) DEFAULT NULL,
  `request_meta` text,
  `auto_code_path` text,
  `injection_meta` text,
  `struct_name` varchar(191) DEFAULT NULL,
  `struct_cn_name` varchar(191) DEFAULT NULL,
  `api_ids` varchar(191) DEFAULT NULL,
  `flag` bigint(20) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_code_histories_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_auto_code_histories` */

insert  into `sys_auto_code_histories`(`id`,`created_at`,`updated_at`,`deleted_at`,`package`,`business_db`,`table_name`,`request_meta`,`auto_code_path`,`injection_meta`,`struct_name`,`struct_cn_name`,`api_ids`,`flag`) values 
(1,'2023-09-20 10:11:14.265','2023-09-20 10:11:14.265',NULL,'webcms','','recruit','{\"structName\":\"Recruit\",\"tableName\":\"recruit\",\"packageName\":\"recruit\",\"humpPackageName\":\"recruit\",\"abbreviation\":\"recruit\",\"description\":\"招聘模型\",\"autoCreateApiToSql\":true,\"autoCreateResource\":false,\"autoMoveFile\":true,\"businessDB\":\"\",\"fields\":[{\"fieldName\":\"Job\",\"fieldDesc\":\"招聘职位\",\"fieldType\":\"string\",\"fieldJson\":\"job\",\"dataTypeLong\":\"32\",\"comment\":\"\",\"columnName\":\"job\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":true,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Count\",\"fieldDesc\":\"招聘人数\",\"fieldType\":\"string\",\"fieldJson\":\"count\",\"dataTypeLong\":\"4\",\"comment\":\"\",\"columnName\":\"count\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Address\",\"fieldDesc\":\"工作地点\",\"fieldType\":\"string\",\"fieldJson\":\"address\",\"dataTypeLong\":\"191\",\"comment\":\"\",\"columnName\":\"address\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false},{\"fieldName\":\"Content\",\"fieldDesc\":\"岗位描述\",\"fieldType\":\"string\",\"fieldJson\":\"content\",\"dataTypeLong\":\"\",\"comment\":\"\",\"columnName\":\"content\",\"fieldSearchType\":\"\",\"dictType\":\"\",\"require\":false,\"errorText\":\"\",\"clearable\":true,\"sort\":false}],\"HasTimer\":false,\"package\":\"webcms\"}','C:\\Users\\Administrator\\go\\baseAdmmin\\server\\api\\v1\\webcms\\recruit.go;C:\\Users\\Administrator\\go\\baseAdmmin\\server\\model\\webcms\\recruit.go;C:\\Users\\Administrator\\go\\baseAdmmin\\server\\model\\webcms\\request\\recruit.go;C:\\Users\\Administrator\\go\\baseAdmmin\\server\\router\\webcms\\recruit.go;C:\\Users\\Administrator\\go\\baseAdmmin\\server\\service\\webcms\\recruit.go;C:\\Users\\Administrator\\go\\baseAdmmin\\web\\src\\api\\recruit.js;C:\\Users\\Administrator\\go\\baseAdmmin\\web\\src\\view\\recruit\\recruitForm.vue;C:\\Users\\Administrator\\go\\baseAdmmin\\web\\src\\view\\recruit\\recruit.vue;','C:\\Users\\Administrator\\go\\baseAdmmin\\server\\api\\v1\\webcms\\enter.go@ApiGroup@RecruitApi;C:\\Users\\Administrator\\go\\baseAdmmin\\server\\router\\webcms\\enter.go@RouterGroup@RecruitRouter;C:\\Users\\Administrator\\go\\baseAdmmin\\server\\service\\webcms\\enter.go@ServiceGroup@RecruitService;','Recruit','招聘模型','151;152;153;154;155;156;',0),
(2,'2023-09-20 10:34:18.000','2023-09-20 10:34:20.000',NULL,'webcms',NULL,'class',NULL,NULL,NULL,'Class','文章模型',NULL,0);

/*Table structure for table `sys_auto_codes` */

DROP TABLE IF EXISTS `sys_auto_codes`;

CREATE TABLE `sys_auto_codes` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `package_name` varchar(191) DEFAULT NULL COMMENT '包名',
  `label` varchar(191) DEFAULT NULL COMMENT '展示名',
  `desc` varchar(191) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_auto_codes_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_auto_codes` */

/*Table structure for table `sys_base_menu_btns` */

DROP TABLE IF EXISTS `sys_base_menu_btns`;

CREATE TABLE `sys_base_menu_btns` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '按钮关键key',
  `desc` varchar(191) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL COMMENT '菜单ID',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_btns_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_base_menu_btns` */

/*Table structure for table `sys_base_menu_parameters` */

DROP TABLE IF EXISTS `sys_base_menu_parameters`;

CREATE TABLE `sys_base_menu_parameters` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `sys_base_menu_id` bigint(20) unsigned DEFAULT NULL,
  `type` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数为params还是query',
  `key` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的key',
  `value` varchar(191) DEFAULT NULL COMMENT '地址栏携带参数的值',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menu_parameters_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_base_menu_parameters` */

/*Table structure for table `sys_base_menus` */

DROP TABLE IF EXISTS `sys_base_menus`;

CREATE TABLE `sys_base_menus` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `menu_level` bigint(20) unsigned DEFAULT NULL,
  `parent_id` varchar(191) DEFAULT NULL COMMENT '父菜单ID',
  `path` varchar(191) DEFAULT NULL COMMENT '路由path',
  `name` varchar(191) DEFAULT NULL COMMENT '路由name',
  `hidden` tinyint(1) DEFAULT NULL COMMENT '是否在列表隐藏',
  `component` varchar(191) DEFAULT NULL COMMENT '对应前端文件路径',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `active_name` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `keep_alive` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `default_menu` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  `title` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `icon` varchar(191) DEFAULT NULL COMMENT '附加属性',
  `close_tab` tinyint(1) DEFAULT NULL COMMENT '附加属性',
  PRIMARY KEY (`id`),
  KEY `idx_sys_base_menus_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_base_menus` */

insert  into `sys_base_menus`(`id`,`created_at`,`updated_at`,`deleted_at`,`menu_level`,`parent_id`,`path`,`name`,`hidden`,`component`,`sort`,`active_name`,`keep_alive`,`default_menu`,`title`,`icon`,`close_tab`) values 
(1,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'0','dashboard','dashboard',0,'view/dashboard/index.vue',1,'',0,0,'仪表盘','odometer',0),
(2,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'0','about','about',0,'view/about/index.vue',9,'',0,0,'关于我们','info-filled',0),
(3,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'0','admin','superAdmin',0,'view/superAdmin/index.vue',3,'',0,0,'超级管理员','user',0),
(4,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'3','authority','authority',0,'view/superAdmin/authority/authority.vue',1,'',0,0,'角色管理','avatar',0),
(5,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'3','menu','menu',0,'view/superAdmin/menu/menu.vue',2,'',1,0,'菜单管理','tickets',0),
(6,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'3','api','api',0,'view/superAdmin/api/api.vue',3,'',1,0,'api管理','platform',0),
(7,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'3','user','user',0,'view/superAdmin/user/user.vue',4,'',0,0,'用户管理','coordinate',0),
(8,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'3','dictionary','dictionary',0,'view/superAdmin/dictionary/sysDictionary.vue',5,'',0,0,'字典管理','notebook',0),
(9,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'3','operation','operation',0,'view/superAdmin/operation/sysOperationRecord.vue',6,'',0,0,'操作历史','pie-chart',0),
(10,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'0','person','person',1,'view/person/person.vue',4,'',0,0,'个人信息','message',0),
(11,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'0','example','example',0,'view/example/index.vue',7,'',0,0,'示例文件','management',0),
(12,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'11','upload','upload',0,'view/example/upload/upload.vue',5,'',0,0,'媒体库（上传下载）','upload',0),
(13,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'11','breakpoint','breakpoint',0,'view/example/breakpoint/breakpoint.vue',6,'',0,0,'断点续传','upload-filled',0),
(14,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'11','customer','customer',0,'view/example/customer/customer.vue',7,'',0,0,'客户列表（资源示例）','avatar',0),
(15,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'0','systemTools','systemTools',0,'view/systemTools/index.vue',5,'',0,0,'系统工具','tools',0),
(16,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'15','autoCode','autoCode',0,'view/systemTools/autoCode/index.vue',1,'',1,0,'代码生成器','cpu',0),
(17,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'15','formCreate','formCreate',0,'view/systemTools/formCreate/index.vue',2,'',1,0,'表单生成器','magic-stick',0),
(18,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'15','system','system',0,'view/systemTools/system/system.vue',3,'',0,0,'系统配置','operation',0),
(19,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'15','autoCodeAdmin','autoCodeAdmin',0,'view/systemTools/autoCodeAdmin/index.vue',1,'',0,0,'自动化代码管理','magic-stick',0),
(20,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'15','autoCodeEdit/:id','autoCodeEdit',1,'view/systemTools/autoCode/index.vue',0,'',0,0,'自动化代码-${id}','magic-stick',0),
(21,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'15','autoPkg','autoPkg',0,'view/systemTools/autoPkg/autoPkg.vue',0,'',0,0,'自动化package','folder',0),
(22,'2024-03-20 11:40:40.118','2024-03-20 17:52:55.729',NULL,0,'0','/','/',0,'/',0,'',0,0,'官方网站','customer-gva',0),
(23,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'0','state','state',0,'view/system/state.vue',8,'',0,0,'服务器状态','cloudy',0),
(24,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118','2024-03-20 17:39:46.557',0,'0','plugin','plugin',0,'view/routerHolder.vue',6,'',0,0,'插件系统','cherry',0),
(25,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118','2024-03-20 17:39:24.821',0,'24','https://plugin.gin-vue-admin.com/','https://plugin.gin-vue-admin.com/',0,'https://plugin.gin-vue-admin.com/',0,'',0,0,'插件市场','shop',0),
(26,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118','2024-03-20 17:39:28.314',0,'24','installPlugin','installPlugin',0,'view/systemTools/installPlugin/index.vue',1,'',0,0,'插件安装','box',0),
(27,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118','2024-03-20 17:39:31.696',0,'24','autoPlug','autoPlug',0,'view/systemTools/autoPlug/autoPlug.vue',2,'',0,0,'插件模板','folder',0),
(28,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118','2024-03-20 17:39:34.576',0,'24','pubPlug','pubPlug',0,'view/systemTools/pubPlug/pubPlug.vue',3,'',0,0,'打包插件','files',0),
(29,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118','2024-03-20 17:39:41.428',0,'24','plugin-email','plugin-email',0,'plugin/email/view/index.vue',4,'',0,0,'邮件插件','message',0),
(30,'2024-03-20 11:40:40.118','2024-03-20 11:40:40.118',NULL,0,'15','exportTemplate','exportTemplate',0,'view/systemTools/exportTemplate/exportTemplate.vue',10,'',0,0,'表格模板','reading',0),
(31,'2024-03-21 09:43:00.077','2024-03-21 09:54:59.354',NULL,0,'0','project','project',0,'view/webcms/index.vue',4,'',0,0,'cms管理','aim',0),
(32,'2024-03-21 09:47:03.425','2024-03-21 09:47:03.425',NULL,0,'31','webcinfig','webcinfig',0,'view/webcms/config/index.vue',0,'',0,0,'网站配置','aim',0),
(33,'2024-03-21 09:48:18.699','2024-03-21 09:48:18.699',NULL,0,'31','swiper','swiper',0,'view/webcms/swiper/swiper.vue',0,'',0,0,'轮播图','picture',0),
(34,'2024-03-21 09:49:21.977','2024-03-21 09:49:21.977',NULL,0,'31','cateMenus','cateMenus',0,'view/webcms/cateMenus/cateMenus.vue',0,'',0,0,'栏目管理','calendar',0),
(35,'2024-03-21 09:51:03.122','2024-03-21 09:51:03.122',NULL,0,'31','content','content',0,'view/webcms/content/content.vue',0,'',0,0,'内容管理','edit',0),
(36,'2024-03-21 09:51:41.121','2024-03-21 09:51:41.121',NULL,0,'31','message','message',0,'view/webcms/message/message.vue',0,'',0,0,'留言管理','aim',0),
(37,'2024-03-21 09:52:10.155','2024-03-21 09:52:10.155',NULL,0,'31','links','links',0,'view/webcms/links/links.vue',0,'',0,0,'友情链接','aim',0);

/*Table structure for table `sys_data_authority_id` */

DROP TABLE IF EXISTS `sys_data_authority_id`;

CREATE TABLE `sys_data_authority_id` (
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  `data_authority_id_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_authority_authority_id`,`data_authority_id_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_data_authority_id` */

insert  into `sys_data_authority_id`(`sys_authority_authority_id`,`data_authority_id_authority_id`) values 
(888,888),
(888,8881),
(888,9528),
(9528,8881),
(9528,9528);

/*Table structure for table `sys_dictionaries` */

DROP TABLE IF EXISTS `sys_dictionaries`;

CREATE TABLE `sys_dictionaries` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` varchar(191) DEFAULT NULL COMMENT '字典名（中）',
  `type` varchar(191) DEFAULT NULL COMMENT '字典名（英）',
  `status` tinyint(1) DEFAULT NULL COMMENT '状态',
  `desc` varchar(191) DEFAULT NULL COMMENT '描述',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionaries_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=9 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_dictionaries` */

insert  into `sys_dictionaries`(`id`,`created_at`,`updated_at`,`deleted_at`,`name`,`type`,`status`,`desc`) values 
(1,'2024-03-20 11:40:39.547','2024-03-20 11:40:39.630',NULL,'性别','gender',1,'性别字典'),
(2,'2024-03-20 11:40:39.547','2024-03-20 11:40:39.771',NULL,'数据库int类型','int',1,'int类型对应的数据库类型'),
(3,'2024-03-20 11:40:39.547','2024-03-20 11:40:39.841',NULL,'数据库时间日期类型','time.Time',1,'数据库时间日期类型'),
(4,'2024-03-20 11:40:39.547','2024-03-20 11:40:39.910',NULL,'数据库浮点型','float64',1,'数据库浮点型'),
(5,'2024-03-20 11:40:39.547','2024-03-20 11:40:39.967',NULL,'数据库字符串','string',1,'数据库字符串'),
(6,'2024-03-20 11:40:39.547','2024-03-20 11:40:40.046',NULL,'数据库bool类型','bool',1,'数据库bool类型'),
(7,'2024-03-22 16:38:37.397','2024-03-22 16:38:37.397',NULL,'栏目类型','catetype',1,'栏目类型'),
(8,'2024-03-22 16:40:26.949','2024-03-22 16:40:26.949',NULL,'排序方式','ordertype',1,'列表页排序方式');

/*Table structure for table `sys_dictionary_details` */

DROP TABLE IF EXISTS `sys_dictionary_details`;

CREATE TABLE `sys_dictionary_details` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `label` varchar(191) DEFAULT NULL COMMENT '展示值',
  `value` varchar(191) DEFAULT NULL COMMENT '字典值',
  `extend` varchar(191) DEFAULT NULL COMMENT '扩展值',
  `status` tinyint(1) DEFAULT NULL COMMENT '启用状态',
  `sort` bigint(20) DEFAULT NULL COMMENT '排序标记',
  `sys_dictionary_id` bigint(20) unsigned DEFAULT NULL COMMENT '关联标记',
  PRIMARY KEY (`id`),
  KEY `idx_sys_dictionary_details_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=39 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_dictionary_details` */

insert  into `sys_dictionary_details`(`id`,`created_at`,`updated_at`,`deleted_at`,`label`,`value`,`extend`,`status`,`sort`,`sys_dictionary_id`) values 
(1,'2024-03-20 11:40:39.632','2024-03-20 11:40:39.632',NULL,'男','1','',1,1,1),
(2,'2024-03-20 11:40:39.632','2024-03-20 11:40:39.632',NULL,'女','2','',1,2,1),
(3,'2024-03-20 11:40:39.777','2024-03-20 11:40:39.777',NULL,'smallint','1','mysql',1,1,2),
(4,'2024-03-20 11:40:39.777','2024-03-20 11:40:39.777',NULL,'mediumint','2','mysql',1,2,2),
(5,'2024-03-20 11:40:39.777','2024-03-20 11:40:39.777',NULL,'int','3','mysql',1,3,2),
(6,'2024-03-20 11:40:39.777','2024-03-20 11:40:39.777',NULL,'bigint','4','mysql',1,4,2),
(7,'2024-03-20 11:40:39.777','2024-03-20 11:40:39.777',NULL,'int2','5','pgsql',1,5,2),
(8,'2024-03-20 11:40:39.777','2024-03-20 11:40:39.777',NULL,'int4','6','pgsql',1,6,2),
(9,'2024-03-20 11:40:39.777','2024-03-20 11:40:39.777',NULL,'int6','7','pgsql',1,7,2),
(10,'2024-03-20 11:40:39.777','2024-03-20 11:40:39.777',NULL,'int8','8','pgsql',1,8,2),
(11,'2024-03-20 11:40:39.847','2024-03-20 11:40:39.847',NULL,'date','','',1,0,3),
(12,'2024-03-20 11:40:39.847','2024-03-20 11:40:39.847',NULL,'time','1','mysql',1,1,3),
(13,'2024-03-20 11:40:39.847','2024-03-20 11:40:39.847',NULL,'year','2','mysql',1,2,3),
(14,'2024-03-20 11:40:39.847','2024-03-20 11:40:39.847',NULL,'datetime','3','mysql',1,3,3),
(15,'2024-03-20 11:40:39.847','2024-03-20 11:40:39.847',NULL,'timestamp','5','mysql',1,5,3),
(16,'2024-03-20 11:40:39.847','2024-03-20 11:40:39.847',NULL,'timestamptz','6','pgsql',1,5,3),
(17,'2024-03-20 11:40:39.917','2024-03-20 11:40:39.917',NULL,'float','','',1,0,4),
(18,'2024-03-20 11:40:39.917','2024-03-20 11:40:39.917',NULL,'double','1','mysql',1,1,4),
(19,'2024-03-20 11:40:39.917','2024-03-20 11:40:39.917',NULL,'decimal','2','mysql',1,2,4),
(20,'2024-03-20 11:40:39.917','2024-03-20 11:40:39.917',NULL,'numeric','3','pgsql',1,3,4),
(21,'2024-03-20 11:40:39.917','2024-03-20 11:40:39.917',NULL,'smallserial','4','pgsql',1,4,4),
(22,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'char','','',1,0,5),
(23,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'varchar','1','mysql',1,1,5),
(24,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'tinyblob','2','mysql',1,2,5),
(25,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'tinytext','3','mysql',1,3,5),
(26,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'text','4','mysql',1,4,5),
(27,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'blob','5','mysql',1,5,5),
(28,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'mediumblob','6','mysql',1,6,5),
(29,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'mediumtext','7','mysql',1,7,5),
(30,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'longblob','8','mysql',1,8,5),
(31,'2024-03-20 11:40:39.972','2024-03-20 11:40:39.972',NULL,'longtext','9','mysql',1,9,5),
(32,'2024-03-20 11:40:40.051','2024-03-20 11:40:40.051',NULL,'tinyint','1','mysql',1,0,6),
(33,'2024-03-20 11:40:40.051','2024-03-20 11:40:40.051',NULL,'bool','2','pgsql',1,0,6),
(34,'2024-03-22 16:39:16.842','2024-03-22 16:39:16.842',NULL,'封面','1','',1,0,7),
(35,'2024-03-22 16:39:31.161','2024-03-22 16:39:31.161',NULL,'列表','2','',1,1,7),
(36,'2024-03-22 16:39:50.626','2024-03-22 16:39:50.626',NULL,'链接','3','',1,2,7),
(37,'2024-03-22 16:40:43.214','2024-03-22 16:40:43.214',NULL,'升序','1','',1,0,8),
(38,'2024-03-22 16:40:55.254','2024-03-22 16:40:55.254',NULL,'降序','2','',1,1,8);

/*Table structure for table `sys_export_template_condition` */

DROP TABLE IF EXISTS `sys_export_template_condition`;

CREATE TABLE `sys_export_template_condition` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `template_id` varchar(191) DEFAULT NULL COMMENT '模板标识',
  `from` varchar(191) DEFAULT NULL COMMENT '条件取的key',
  `column` varchar(191) DEFAULT NULL COMMENT '作为查询条件的字段',
  `operator` varchar(191) DEFAULT NULL COMMENT '操作符',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_template_condition_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_export_template_condition` */

/*Table structure for table `sys_export_templates` */

DROP TABLE IF EXISTS `sys_export_templates`;

CREATE TABLE `sys_export_templates` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `db_name` varchar(191) DEFAULT NULL COMMENT '数据库名称',
  `name` varchar(191) DEFAULT NULL COMMENT '模板名称',
  `table_name` varchar(191) DEFAULT NULL COMMENT '表名称',
  `template_id` varchar(191) DEFAULT NULL COMMENT '模板标识',
  `template_info` text,
  `limit` bigint(20) DEFAULT NULL COMMENT '导出限制',
  `order` varchar(191) DEFAULT NULL COMMENT '排序',
  PRIMARY KEY (`id`),
  KEY `idx_sys_export_templates_deleted_at` (`deleted_at`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_export_templates` */

/*Table structure for table `sys_operation_records` */

DROP TABLE IF EXISTS `sys_operation_records`;

CREATE TABLE `sys_operation_records` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `ip` varchar(191) DEFAULT NULL COMMENT '请求ip',
  `method` varchar(191) DEFAULT NULL COMMENT '请求方法',
  `path` varchar(191) DEFAULT NULL COMMENT '请求路径',
  `status` bigint(20) DEFAULT NULL COMMENT '请求状态',
  `latency` bigint(20) DEFAULT NULL COMMENT '延迟',
  `agent` varchar(191) DEFAULT NULL COMMENT '代理',
  `error_message` varchar(191) DEFAULT NULL COMMENT '错误信息',
  `body` text COMMENT '请求Body',
  `resp` text COMMENT '响应Body',
  `user_id` bigint(20) unsigned DEFAULT NULL COMMENT '用户id',
  PRIMARY KEY (`id`),
  KEY `idx_sys_operation_records_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=118 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_operation_records` */

insert  into `sys_operation_records`(`id`,`created_at`,`updated_at`,`deleted_at`,`ip`,`method`,`path`,`status`,`latency`,`agent`,`error_message`,`body`,`resp`,`user_id`) values 
(1,'2024-03-20 17:39:25.009','2024-03-20 17:39:25.009',NULL,'127.0.0.1','POST','/menu/deleteBaseMenu',200,207381700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":25}','{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}',1),
(2,'2024-03-20 17:39:28.438','2024-03-20 17:39:28.438',NULL,'127.0.0.1','POST','/menu/deleteBaseMenu',200,137496400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":26}','{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}',1),
(3,'2024-03-20 17:39:31.804','2024-03-20 17:39:31.804',NULL,'127.0.0.1','POST','/menu/deleteBaseMenu',200,119663900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":27}','{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}',1),
(4,'2024-03-20 17:39:34.688','2024-03-20 17:39:34.688',NULL,'127.0.0.1','POST','/menu/deleteBaseMenu',200,123920600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":28}','{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}',1),
(5,'2024-03-20 17:39:41.576','2024-03-20 17:39:41.576',NULL,'127.0.0.1','POST','/menu/deleteBaseMenu',200,158604300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":29}','{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}',1),
(6,'2024-03-20 17:39:46.696','2024-03-20 17:39:46.696',NULL,'127.0.0.1','POST','/menu/deleteBaseMenu',200,155871700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":24}','{\"code\":0,\"data\":{},\"msg\":\"删除成功\"}',1),
(7,'2024-03-20 17:50:39.981','2024-03-20 17:50:39.981',NULL,'127.0.0.1','POST','/system/getServerInfo',200,271002700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','','{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.20.12\",\"numGoroutine\":10},\"cpu\":{\"cpus\":[7.142857142857142,21.428571428571427,14.285714285714285,0,7.142857142857142,7.142857142857142,7.142857142857142,7.142857142857142],\"cores\":4},\"ram\":{\"usedMb\":11647,\"totalMb\":16283,\"usedPercent\":71},\"disk\":{\"usedMb\":222224,\"usedGb\":217,\"totalMb\":228934,\"totalGb\":223,\"usedPercent\":97}}},\"msg\":\"获取成功\"}',1),
(8,'2024-03-20 17:50:49.945','2024-03-20 17:50:49.945',NULL,'127.0.0.1','POST','/system/getServerInfo',200,230227300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','','{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.20.12\",\"numGoroutine\":10},\"cpu\":{\"cpus\":[7.6923076923076925,0,0,0,0,0,0,0],\"cores\":4},\"ram\":{\"usedMb\":11626,\"totalMb\":16283,\"usedPercent\":71},\"disk\":{\"usedMb\":222223,\"usedGb\":217,\"totalMb\":228934,\"totalGb\":223,\"usedPercent\":97}}},\"msg\":\"获取成功\"}',1),
(9,'2024-03-20 17:50:59.263','2024-03-20 17:50:59.263',NULL,'127.0.0.1','PUT','/user/setSelfInfo',200,45112000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"sideMode\":\"light\"}','{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}',1),
(10,'2024-03-20 17:50:59.965','2024-03-20 17:50:59.965',NULL,'127.0.0.1','POST','/system/getServerInfo',200,257103000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','','{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.20.12\",\"numGoroutine\":10},\"cpu\":{\"cpus\":[0,0,14.285714285714285,0,7.142857142857142,0,0,0],\"cores\":4},\"ram\":{\"usedMb\":11631,\"totalMb\":16283,\"usedPercent\":71},\"disk\":{\"usedMb\":222223,\"usedGb\":217,\"totalMb\":228934,\"totalGb\":223,\"usedPercent\":97}}},\"msg\":\"获取成功\"}',1),
(11,'2024-03-20 17:51:02.907','2024-03-20 17:51:02.907',NULL,'127.0.0.1','PUT','/user/setSelfInfo',200,57187000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"sideMode\":\"dark\"}','{\"code\":0,\"data\":{},\"msg\":\"设置成功\"}',1),
(12,'2024-03-20 17:52:55.774','2024-03-20 17:52:55.774',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,59886300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":22,\"CreatedAt\":\"2024-03-20T11:40:40.118+08:00\",\"UpdatedAt\":\"2024-03-20T11:40:40.118+08:00\",\"parentId\":\"0\",\"path\":\"/\",\"name\":\"/\",\"hidden\":false,\"component\":\"/\",\"sort\":0,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"官方网站\",\"icon\":\"customer-gva\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),
(13,'2024-03-20 18:03:11.311','2024-03-20 18:03:11.311',NULL,'127.0.0.1','POST','/system/getServerInfo',200,250149100,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','','{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.20.12\",\"numGoroutine\":14},\"cpu\":{\"cpus\":[7.142857142857142,7.142857142857142,7.142857142857142,28.57142857142857,7.142857142857142,7.142857142857142,14.285714285714285,57.14285714285714],\"cores\":4},\"ram\":{\"usedMb\":11763,\"totalMb\":16283,\"usedPercent\":72},\"disk\":{\"usedMb\":222251,\"usedGb\":217,\"totalMb\":228934,\"totalGb\":223,\"usedPercent\":97}}},\"msg\":\"获取成功\"}',1),
(14,'2024-03-21 08:26:34.438','2024-03-21 08:26:34.438',NULL,'127.0.0.1','POST','/system/getServerInfo',200,279497400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','','{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.20.12\",\"numGoroutine\":9},\"cpu\":{\"cpus\":[0,7.6923076923076925,7.6923076923076925,0,15.384615384615385,0,7.6923076923076925,15.384615384615385],\"cores\":4},\"ram\":{\"usedMb\":12032,\"totalMb\":16283,\"usedPercent\":73},\"disk\":{\"usedMb\":220764,\"usedGb\":215,\"totalMb\":228934,\"totalGb\":223,\"usedPercent\":96}}},\"msg\":\"获取成功\"}',1),
(15,'2024-03-21 08:26:44.435','2024-03-21 08:26:44.435',NULL,'127.0.0.1','POST','/system/getServerInfo',200,271807600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','','{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.20.12\",\"numGoroutine\":9},\"cpu\":{\"cpus\":[7.142857142857142,0,14.285714285714285,7.142857142857142,21.428571428571427,0,7.142857142857142,14.285714285714285],\"cores\":4},\"ram\":{\"usedMb\":11936,\"totalMb\":16283,\"usedPercent\":73},\"disk\":{\"usedMb\":220764,\"usedGb\":215,\"totalMb\":228934,\"totalGb\":223,\"usedPercent\":96}}},\"msg\":\"获取成功\"}',1),
(16,'2024-03-21 08:26:54.401','2024-03-21 08:26:54.401',NULL,'127.0.0.1','POST','/system/getServerInfo',200,237204800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','','{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.20.12\",\"numGoroutine\":9},\"cpu\":{\"cpus\":[7.6923076923076925,7.6923076923076925,0,7.6923076923076925,7.142857142857142,0,0,0],\"cores\":4},\"ram\":{\"usedMb\":11901,\"totalMb\":16283,\"usedPercent\":73},\"disk\":{\"usedMb\":220764,\"usedGb\":215,\"totalMb\":228934,\"totalGb\":223,\"usedPercent\":96}}},\"msg\":\"获取成功\"}',1),
(17,'2024-03-21 09:32:51.080','2024-03-21 09:32:51.080',NULL,'127.0.0.1','POST','/system/getServerInfo',200,240722000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','','{\"code\":0,\"data\":{\"server\":{\"os\":{\"goos\":\"windows\",\"numCpu\":8,\"compiler\":\"gc\",\"goVersion\":\"go1.20.12\",\"numGoroutine\":9},\"cpu\":{\"cpus\":[0,0,7.6923076923076925,0,7.6923076923076925,0,7.6923076923076925,0],\"cores\":4},\"ram\":{\"usedMb\":10008,\"totalMb\":16283,\"usedPercent\":61},\"disk\":{\"usedMb\":220641,\"usedGb\":215,\"totalMb\":228934,\"totalGb\":223,\"usedPercent\":96}}},\"msg\":\"获取成功\"}',1),
(18,'2024-03-21 09:43:00.160','2024-03-21 09:43:00.160',NULL,'127.0.0.1','POST','/menu/addBaseMenu',200,89893400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":0,\"path\":\"project\",\"name\":\"project\",\"hidden\":false,\"parentId\":\"0\",\"component\":\"view/studypro/index.vue\",\"meta\":{\"title\":\"cms管理\",\"icon\":\"aim\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false}}','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),
(19,'2024-03-21 09:43:15.304','2024-03-21 09:43:15.304',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,42662000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":31,\"CreatedAt\":\"2024-03-21T09:43:00.077+08:00\",\"UpdatedAt\":\"2024-03-21T09:43:00.077+08:00\",\"parentId\":\"0\",\"path\":\"project\",\"name\":\"project\",\"hidden\":false,\"component\":\"view/studypro/index.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"cms管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),
(20,'2024-03-21 09:47:03.506','2024-03-21 09:47:03.506',NULL,'127.0.0.1','POST','/menu/addBaseMenu',200,84404500,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":0,\"path\":\"webcinfig\",\"name\":\"webcinfig\",\"hidden\":false,\"parentId\":\"31\",\"component\":\"view/webcms/config/index.vue\",\"meta\":{\"title\":\"网站配置\",\"icon\":\"aim\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false}}','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),
(21,'2024-03-21 09:48:18.772','2024-03-21 09:48:18.772',NULL,'127.0.0.1','POST','/menu/addBaseMenu',200,77364200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":0,\"path\":\"swiper\",\"name\":\"swiper\",\"hidden\":false,\"parentId\":\"31\",\"component\":\"view/webcms/swiper/swiper.vue\",\"meta\":{\"title\":\"轮播图\",\"icon\":\"picture\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false}}','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),
(22,'2024-03-21 09:49:22.052','2024-03-21 09:49:22.052',NULL,'127.0.0.1','POST','/menu/addBaseMenu',200,78848800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":0,\"path\":\"cateMenus\",\"name\":\"cateMenus\",\"hidden\":false,\"parentId\":\"31\",\"component\":\"view/webcms/cateMenus/cateMenus.vue\",\"meta\":{\"title\":\"栏目管理\",\"icon\":\"calendar\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false}}','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),
(23,'2024-03-21 09:51:03.413','2024-03-21 09:51:03.413',NULL,'127.0.0.1','POST','/menu/addBaseMenu',200,293279400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":0,\"path\":\"content\",\"name\":\"content\",\"hidden\":false,\"parentId\":\"31\",\"component\":\"view/webcms/content/content.vue\",\"meta\":{\"title\":\"内容管理\",\"icon\":\"edit\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false}}','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),
(24,'2024-03-21 09:51:41.198','2024-03-21 09:51:41.198',NULL,'127.0.0.1','POST','/menu/addBaseMenu',200,80721400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":0,\"path\":\"message\",\"name\":\"message\",\"hidden\":false,\"parentId\":\"31\",\"component\":\"view/webcms/message/message.vue\",\"meta\":{\"title\":\"留言管理\",\"icon\":\"aim\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false}}','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),
(25,'2024-03-21 09:52:10.199','2024-03-21 09:52:10.199',NULL,'127.0.0.1','POST','/menu/addBaseMenu',200,48575400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":0,\"path\":\"links\",\"name\":\"links\",\"hidden\":false,\"parentId\":\"31\",\"component\":\"view/webcms/links/links.vue\",\"meta\":{\"title\":\"友情链接\",\"icon\":\"aim\",\"defaultMenu\":false,\"closeTab\":false,\"keepAlive\":false}}','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),
(26,'2024-03-21 09:54:59.427','2024-03-21 09:54:59.427',NULL,'127.0.0.1','POST','/menu/updateBaseMenu',200,79981000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"ID\":31,\"CreatedAt\":\"2024-03-21T09:43:00.077+08:00\",\"UpdatedAt\":\"2024-03-21T09:43:15.272+08:00\",\"parentId\":\"0\",\"path\":\"project\",\"name\":\"project\",\"hidden\":false,\"component\":\"view/webcms/index.vue\",\"sort\":4,\"meta\":{\"activeName\":\"\",\"keepAlive\":false,\"defaultMenu\":false,\"title\":\"cms管理\",\"icon\":\"aim\",\"closeTab\":false},\"authoritys\":null,\"children\":null,\"parameters\":[],\"menuBtn\":[]}','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),
(27,'2024-03-21 11:45:23.699','2024-03-21 11:45:23.699',NULL,'127.0.0.1','POST','/menu/addMenuAuthority',200,107219100,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),
(28,'2024-03-21 11:45:25.610','2024-03-21 11:45:25.610',NULL,'127.0.0.1','POST','/menu/addMenuAuthority',200,106404600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"添加成功\"}',1),
(29,'2024-03-21 14:26:26.795','2024-03-21 14:26:26.795',NULL,'127.0.0.1','POST','/casbin/updateCasbin',200,324107300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','[超出记录长度]','{\"code\":0,\"data\":{},\"msg\":\"更新成功\"}',1),
(30,'2024-03-22 09:44:11.421','2024-03-22 09:44:11.421',NULL,'127.0.0.1','GET','/getWebconfig',200,10843300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":7,\"data\":{},\"msg\":\"获取失败\"}',1),
(31,'2024-03-22 10:52:59.495','2024-03-22 10:52:59.495',NULL,'127.0.0.1','GET','/getWebconfig',200,12427900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(32,'2024-03-22 11:11:12.921','2024-03-22 11:11:12.921',NULL,'127.0.0.1','GET','/getWebconfig',200,4004800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(33,'2024-03-22 11:51:45.853','2024-03-22 11:51:45.853',NULL,'127.0.0.1','GET','/getWebconfig',200,3864500,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(34,'2024-03-22 11:52:02.236','2024-03-22 11:52:02.236',NULL,'127.0.0.1','GET','/getWebconfig',200,2565800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(35,'2024-03-22 11:52:26.087','2024-03-22 11:52:26.087',NULL,'127.0.0.1','GET','/getWebconfig',200,2611200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(36,'2024-03-22 11:52:28.208','2024-03-22 11:52:28.208',NULL,'127.0.0.1','POST','/setWebconfig',200,42802600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(37,'2024-03-22 11:52:37.101','2024-03-22 11:52:37.101',NULL,'127.0.0.1','POST','/setWebconfig',200,31918500,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(38,'2024-03-22 14:22:27.744','2024-03-22 14:22:27.744',NULL,'127.0.0.1','GET','/getWebconfig',200,3372300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(39,'2024-03-22 14:22:31.789','2024-03-22 14:22:31.789',NULL,'127.0.0.1','POST','/setWebconfig',200,87712900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(40,'2024-03-22 14:22:36.064','2024-03-22 14:22:36.064',NULL,'127.0.0.1','GET','/getWebconfig',200,3893300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(41,'2024-03-22 14:23:44.757','2024-03-22 14:23:44.757',NULL,'127.0.0.1','GET','/getWebconfig',200,4089700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(42,'2024-03-22 14:23:56.732','2024-03-22 14:23:56.732',NULL,'127.0.0.1','POST','/setWebconfig',200,44540200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(43,'2024-03-22 14:24:01.377','2024-03-22 14:24:01.377',NULL,'127.0.0.1','GET','/getWebconfig',200,6666300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(44,'2024-03-22 15:36:56.440','2024-03-22 15:36:56.440',NULL,'127.0.0.1','GET','/getWebconfig',200,3774000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(45,'2024-03-22 15:38:22.876','2024-03-22 15:38:22.876',NULL,'127.0.0.1','POST','/setWebconfig',200,58280600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(46,'2024-03-22 15:38:25.652','2024-03-22 15:38:25.652',NULL,'127.0.0.1','POST','/setWebconfig',200,52340800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"}','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(47,'2024-03-22 15:49:49.851','2024-03-22 15:49:49.851',NULL,'127.0.0.1','GET','/getWebconfig',200,2169900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(48,'2024-03-22 15:54:57.715','2024-03-22 15:54:57.715',NULL,'127.0.0.1','GET','/getWebconfig',200,1746700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(49,'2024-03-22 15:55:01.128','2024-03-22 15:55:01.128',NULL,'127.0.0.1','POST','/changeWebconfig',200,507200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','2','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(50,'2024-03-22 15:55:23.762','2024-03-22 15:55:23.762',NULL,'127.0.0.1','GET','/getWebconfig',200,1071900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(51,'2024-03-22 15:55:51.460','2024-03-22 15:55:51.460',NULL,'127.0.0.1','GET','/getWebconfig',200,1715000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(52,'2024-03-22 15:55:53.696','2024-03-22 15:55:53.696',NULL,'127.0.0.1','POST','/changeWebconfig',200,35300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','2','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(53,'2024-03-22 15:56:40.810','2024-03-22 15:56:40.810',NULL,'127.0.0.1','GET','/getWebconfig',200,1663000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(54,'2024-03-22 15:56:46.838','2024-03-22 15:56:46.838',NULL,'127.0.0.1','GET','/getWebconfig',200,1099300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(55,'2024-03-22 15:56:48.968','2024-03-22 15:56:48.968',NULL,'127.0.0.1','POST','/changeWebconfig',200,0,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','2','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(56,'2024-03-22 15:56:50.227','2024-03-22 15:56:50.227',NULL,'127.0.0.1','GET','/getWebconfig',200,1638300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(57,'2024-03-22 15:57:21.033','2024-03-22 15:57:21.033',NULL,'127.0.0.1','GET','/getWebconfig',200,2171700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(58,'2024-03-22 16:30:03.751','2024-03-22 16:30:03.751',NULL,'127.0.0.1','GET','/getWebconfig',200,1645600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(59,'2024-03-22 16:30:06.876','2024-03-22 16:30:06.876',NULL,'127.0.0.1','POST','/changeWebconfig',200,545900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','2','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(60,'2024-03-22 16:30:07.981','2024-03-22 16:30:07.981',NULL,'127.0.0.1','GET','/getWebconfig',200,1664300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(61,'2024-03-22 16:35:13.298','2024-03-22 16:35:13.298',NULL,'127.0.0.1','GET','/getWebconfig',200,1687000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(62,'2024-03-22 16:35:16.569','2024-03-22 16:35:16.569',NULL,'127.0.0.1','GET','/getWebconfig',200,1150100,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(63,'2024-03-22 16:36:35.099','2024-03-22 16:36:35.099',NULL,'127.0.0.1','GET','/getWebconfig',200,1613800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(64,'2024-03-22 16:37:11.397','2024-03-22 16:37:11.397',NULL,'127.0.0.1','GET','/getWebconfig',200,1098200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(65,'2024-03-22 16:37:11.463','2024-03-22 16:37:11.463',NULL,'127.0.0.1','GET','/getWebconfig',200,1785300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(66,'2024-03-22 16:38:37.476','2024-03-22 16:38:37.476',NULL,'127.0.0.1','POST','/sysDictionary/createSysDictionary',200,83432100,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"name\":\"栏目类型\",\"type\":\"catetype\",\"status\":true,\"desc\":\"栏目类型\"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),
(67,'2024-03-22 16:39:16.924','2024-03-22 16:39:16.924',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,82268200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"label\":\"封面\",\"value\":\"1\",\"status\":true,\"sort\":0,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),
(68,'2024-03-22 16:39:31.205','2024-03-22 16:39:31.205',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,44252900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"label\":\"列表\",\"value\":\"2\",\"status\":true,\"sort\":1,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),
(69,'2024-03-22 16:39:50.671','2024-03-22 16:39:50.671',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,46398200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"label\":\"链接\",\"value\":\"3\",\"status\":true,\"sort\":2,\"sysDictionaryID\":7}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),
(70,'2024-03-22 16:40:27.021','2024-03-22 16:40:27.021',NULL,'127.0.0.1','POST','/sysDictionary/createSysDictionary',200,74756900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"name\":\"排序方式\",\"type\":\"ordertype\",\"status\":true,\"desc\":\"列表页排序方式\"}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),
(71,'2024-03-22 16:40:43.260','2024-03-22 16:40:43.260',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,47123300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"label\":\"升序\",\"value\":\"1\",\"status\":true,\"sort\":0,\"sysDictionaryID\":8}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),
(72,'2024-03-22 16:40:55.299','2024-03-22 16:40:55.299',NULL,'127.0.0.1','POST','/sysDictionaryDetail/createSysDictionaryDetail',200,45736800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{\"label\":\"降序\",\"value\":\"2\",\"status\":true,\"sort\":1,\"sysDictionaryID\":8}','{\"code\":0,\"data\":{},\"msg\":\"创建成功\"}',1),
(73,'2024-03-22 16:41:23.937','2024-03-22 16:41:23.937',NULL,'127.0.0.1','GET','/getWebconfig',200,1658300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(74,'2024-03-22 16:41:33.199','2024-03-22 16:41:33.199',NULL,'127.0.0.1','GET','/getWebconfig',200,1624600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(75,'2024-03-22 17:00:25.996','2024-03-22 17:00:25.996',NULL,'127.0.0.1','GET','/getWebconfig',200,2326200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(76,'2024-03-22 17:04:23.618','2024-03-22 17:04:23.618',NULL,'127.0.0.1','GET','/getWebconfig',200,2348100,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(77,'2024-03-22 17:06:06.462','2024-03-22 17:06:06.462',NULL,'127.0.0.1','GET','/getWebconfig',200,2184800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(78,'2024-03-22 17:07:00.245','2024-03-22 17:07:00.245',NULL,'127.0.0.1','POST','/changeWebconfig',200,2053800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','2','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(79,'2024-03-22 17:07:02.272','2024-03-22 17:07:02.272',NULL,'127.0.0.1','GET','/getWebconfig',200,2223900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(80,'2024-03-22 17:20:43.413','2024-03-22 17:20:43.413',NULL,'127.0.0.1','POST','/changeWebconfig',200,566800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','1','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(81,'2024-03-22 17:20:44.741','2024-03-22 17:20:44.741',NULL,'127.0.0.1','GET','/getWebconfig',200,1635900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(82,'2024-03-22 17:20:47.855','2024-03-22 17:20:47.855',NULL,'127.0.0.1','POST','/changeWebconfig',200,549500,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','2','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(83,'2024-03-22 17:20:49.127','2024-03-22 17:20:49.127',NULL,'127.0.0.1','GET','/getWebconfig',200,1685400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(84,'2024-03-22 17:21:08.186','2024-03-22 17:21:08.186',NULL,'127.0.0.1','GET','/getWebconfig',200,1827800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.289 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(85,'2024-03-22 17:21:18.002','2024-03-22 17:21:18.002',NULL,'127.0.0.1','POST','/changeWebconfig',200,0,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.289 Safari/537.36','','1','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(86,'2024-03-22 17:21:19.183','2024-03-22 17:21:19.183',NULL,'127.0.0.1','GET','/getWebconfig',200,1646500,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.289 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(87,'2024-03-22 17:21:28.943','2024-03-22 17:21:28.943',NULL,'127.0.0.1','GET','/getWebconfig',200,1103200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(88,'2024-03-25 09:05:40.922','2024-03-25 09:05:40.922',NULL,'127.0.0.1','GET','/getWebconfig',200,1620900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(89,'2024-03-25 09:06:44.658','2024-03-25 09:06:44.658',NULL,'127.0.0.1','GET','/getWebconfig',200,1135000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(90,'2024-03-25 09:06:55.638','2024-03-25 09:06:55.638',NULL,'127.0.0.1','POST','/changeWebconfig',200,508900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','2','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(91,'2024-03-25 09:06:56.222','2024-03-25 09:06:56.222',NULL,'127.0.0.1','GET','/getWebconfig',200,1077400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(92,'2024-03-25 09:06:58.429','2024-03-25 09:06:58.429',NULL,'127.0.0.1','POST','/changeWebconfig',200,0,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','1','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(93,'2024-03-25 09:06:58.971','2024-03-25 09:06:58.971',NULL,'127.0.0.1','GET','/getWebconfig',200,2196300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(94,'2024-03-25 09:07:47.239','2024-03-25 09:07:47.239',NULL,'127.0.0.1','GET','/getWebconfig',200,1660800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(95,'2024-03-25 09:08:19.448','2024-03-25 09:08:19.448',NULL,'127.0.0.1','POST','/changeWebconfig',200,528000,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','2','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(96,'2024-03-25 09:08:20.690','2024-03-25 09:08:20.690',NULL,'127.0.0.1','GET','/getWebconfig',200,1626600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(97,'2024-03-25 09:16:55.224','2024-03-25 09:16:55.224',NULL,'127.0.0.1','GET','/getWebconfig',200,2261700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(98,'2024-03-25 09:18:53.101','2024-03-25 09:18:53.101',NULL,'127.0.0.1','GET','/getWebconfig',200,1631400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(99,'2024-03-25 09:25:33.310','2024-03-25 09:25:33.310',NULL,'127.0.0.1','GET','/getWebconfig',200,1699500,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(100,'2024-03-25 14:36:50.412','2024-03-25 14:36:50.412',NULL,'127.0.0.1','GET','/getWebconfig',200,1781500,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(101,'2024-03-25 15:12:53.393','2024-03-25 15:12:53.393',NULL,'127.0.0.1','GET','/getWebconfig',200,1668200,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(102,'2024-03-25 15:16:18.275','2024-03-25 15:16:18.275',NULL,'127.0.0.1','GET','/getWebconfig',200,1568300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(103,'2024-03-25 15:22:41.568','2024-03-25 15:22:41.568',NULL,'127.0.0.1','GET','/getWebconfig',200,1640400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(104,'2024-03-25 15:27:47.859','2024-03-25 15:27:47.859',NULL,'127.0.0.1','GET','/getWebconfig',200,1103700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(105,'2024-03-25 15:28:38.835','2024-03-25 15:28:38.835',NULL,'127.0.0.1','GET','/getWebconfig',200,1726700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(106,'2024-03-25 15:29:44.711','2024-03-25 15:29:44.711',NULL,'127.0.0.1','GET','/getWebconfig',200,1788100,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(107,'2024-03-25 15:30:54.439','2024-03-25 15:30:54.439',NULL,'127.0.0.1','GET','/getWebconfig',200,1106500,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(108,'2024-03-25 15:31:22.611','2024-03-25 15:31:22.611',NULL,'127.0.0.1','GET','/getWebconfig',200,1088900,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(109,'2024-03-25 15:34:42.431','2024-03-25 15:34:42.431',NULL,'127.0.0.1','GET','/getWebconfig',200,1633300,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(110,'2024-03-25 15:36:56.899','2024-03-25 15:36:56.899',NULL,'127.0.0.1','GET','/getWebconfig',200,1625400,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(111,'2024-03-25 15:40:10.718','2024-03-25 15:40:10.718',NULL,'127.0.0.1','POST','/changeWebconfig',200,0,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','2','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(112,'2024-03-25 15:40:12.027','2024-03-25 15:40:12.027',NULL,'127.0.0.1','GET','/getWebconfig',200,1650500,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(113,'2024-03-25 15:40:14.498','2024-03-25 15:40:14.498',NULL,'127.0.0.1','POST','/changeWebconfig',200,0,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','1','{\"code\":0,\"data\":{},\"msg\":\"操作成功\"}',1),
(114,'2024-03-25 15:40:15.700','2024-03-25 15:40:15.700',NULL,'127.0.0.1','GET','/getWebconfig',200,1634800,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(115,'2024-03-25 15:54:13.499','2024-03-25 15:54:13.499',NULL,'127.0.0.1','GET','/getWebconfig',200,1676600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(116,'2024-03-25 15:54:39.525','2024-03-25 15:54:39.525',NULL,'127.0.0.1','GET','/getWebconfig',200,1756600,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1),
(117,'2024-03-25 16:25:48.232','2024-03-25 16:25:48.232',NULL,'127.0.0.1','GET','/getWebconfig',200,2249700,'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36','','{}','{\"code\":0,\"data\":{\"list\":[{\"id\":1,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"028-85593003\",\"phone\":\"+86 1361682301077777\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://rxshop.yhystudy.top/\"},{\"id\":2,\"logo\":\"uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png\",\"hotLine\":\"11111133333\",\"phone\":\"222222244444\",\"wechat\":\"uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg\",\"email\":\"admin@chinachuanhong.com\",\"beian\":\" 蜀ICP备2021020602号-1\",\"address\":\"Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China\",\"company\":\"四川禹冰工程勘察设计有限公司_英文站\",\"keywords\":\"四川禹冰工程勘察设计有限公司\",\"description\":\"四川禹冰工程勘察设计有限公司\",\"siteUrl\":\"http://en.yhystudy.top/\"}],\"siteid\":\"\"},\"msg\":\"获取成功\"}',1);

/*Table structure for table `sys_user_authority` */

DROP TABLE IF EXISTS `sys_user_authority`;

CREATE TABLE `sys_user_authority` (
  `sys_user_id` bigint(20) unsigned NOT NULL,
  `sys_authority_authority_id` bigint(20) unsigned NOT NULL COMMENT '角色ID',
  PRIMARY KEY (`sys_user_id`,`sys_authority_authority_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_user_authority` */

insert  into `sys_user_authority`(`sys_user_id`,`sys_authority_authority_id`) values 
(1,888),
(1,8881),
(1,9528),
(2,888);

/*Table structure for table `sys_users` */

DROP TABLE IF EXISTS `sys_users`;

CREATE TABLE `sys_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `uuid` varchar(191) DEFAULT NULL COMMENT '用户UUID',
  `username` varchar(191) DEFAULT NULL COMMENT '用户登录名',
  `password` varchar(191) DEFAULT NULL COMMENT '用户登录密码',
  `nick_name` varchar(191) DEFAULT '系统用户' COMMENT '用户昵称',
  `side_mode` varchar(191) DEFAULT 'dark' COMMENT '用户侧边主题',
  `header_img` varchar(191) DEFAULT 'https://qmplusimg.henrongyi.top/gva_header.jpg' COMMENT '用户头像',
  `base_color` varchar(191) DEFAULT '#fff' COMMENT '基础颜色',
  `active_color` varchar(191) DEFAULT '#1890ff' COMMENT '活跃颜色',
  `authority_id` bigint(20) unsigned DEFAULT '888' COMMENT '用户角色ID',
  `phone` varchar(191) DEFAULT NULL COMMENT '用户手机号',
  `email` varchar(191) DEFAULT NULL COMMENT '用户邮箱',
  `enable` bigint(20) DEFAULT '1' COMMENT '用户是否被冻结 1正常 2冻结',
  PRIMARY KEY (`id`),
  KEY `idx_sys_users_deleted_at` (`deleted_at`),
  KEY `idx_sys_users_uuid` (`uuid`),
  KEY `idx_sys_users_username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `sys_users` */

insert  into `sys_users`(`id`,`created_at`,`updated_at`,`deleted_at`,`uuid`,`username`,`password`,`nick_name`,`side_mode`,`header_img`,`base_color`,`active_color`,`authority_id`,`phone`,`email`,`enable`) values 
(1,'2024-03-20 11:40:40.317','2024-03-20 17:51:02.849',NULL,'f89c9e6c-372d-4000-8471-3f53940e9865','admin','$2a$10$/RKQ3lgEEHt.bNm22L8N3e8FNJyoYLCV8k9Rk3An03nJcZcj4PFWC','Mr.奇淼','dark','https://qmplusimg.henrongyi.top/gva_header.jpg','#fff','#1890ff',888,'17611111111','333333333@qq.com',1),
(2,'2024-03-20 11:40:40.317','2024-03-20 11:40:40.506',NULL,'a4dfe4fa-4b47-4383-abc0-d977133a7da5','a303176530','$2a$10$RRypL8E1be6m2FYUpJr7aO3.ATiSgcrn3ec9C1oaJGvDG6IWEXkY.','用户1','dark','https:///qmplusimg.henrongyi.top/1572075907logo.png','#fff','#1890ff',9528,'17611111111','333333333@qq.com',1);

/*Table structure for table `webconfig` */

DROP TABLE IF EXISTS `webconfig`;

CREATE TABLE `webconfig` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `logo` varchar(191) DEFAULT NULL COMMENT '网站logo',
  `hot_line` varchar(32) DEFAULT NULL COMMENT '网站热线',
  `phone` varchar(32) DEFAULT NULL COMMENT '手机号码',
  `wechat` varchar(191) DEFAULT NULL COMMENT 'WeChat',
  `email` varchar(32) DEFAULT NULL COMMENT 'email',
  `beian` varchar(191) DEFAULT NULL COMMENT '备案',
  `address` varchar(191) DEFAULT NULL COMMENT '公司地址',
  `company` varchar(191) DEFAULT NULL COMMENT '公司名称',
  `keywords` varchar(191) DEFAULT NULL COMMENT '网站关键字',
  `description` varchar(191) DEFAULT NULL COMMENT '网站描述',
  `site_url` varchar(191) DEFAULT NULL COMMENT '网站域名',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

/*Data for the table `webconfig` */

insert  into `webconfig`(`id`,`logo`,`hot_line`,`phone`,`wechat`,`email`,`beian`,`address`,`company`,`keywords`,`description`,`site_url`) values 
(1,'uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png','028-85593003','+86 1361682301077777','uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg','admin@chinachuanhong.com',' 蜀ICP备2021020602号-1','Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China','四川禹冰工程勘察设计有限公司','四川禹冰工程勘察设计有限公司','四川禹冰工程勘察设计有限公司','http://rxshop.yhystudy.top/'),
(2,'uploads/file/96d6f2e7e1f705ab5e59c84a6dc009b2_20230725155956.png','11111133333','222222244444','uploads/file/94da5fdf2dc2a2e41b4331bd83b38488_20230725160023.jpg','admin@chinachuanhong.com',' 蜀ICP备2021020602号-1','Qiao Nan Industrial Zone, Hongqiao Town, Yueqing City,Zhejiang Province, China','四川禹冰工程勘察设计有限公司_英文站','四川禹冰工程勘察设计有限公司','四川禹冰工程勘察设计有限公司','http://en.yhystudy.top/');

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
