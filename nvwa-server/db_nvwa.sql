# ************************************************************
# Sequel Pro SQL dump
# Version 4499
#
# http://www.sequelpro.com/
# https://github.com/sequelpro/sequelpro
#
# Host: 127.0.0.1 (MySQL 5.6.27)
# Database: db_nvwa
# Generation Time: 2019-01-27 17:28:25 +0000
# ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


# Dump of table app
# ------------------------------------------------------------

DROP TABLE IF EXISTS `app`;

CREATE TABLE `app` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者 ID',
  `project_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属项目 ID',
  `name` varchar(256) NOT NULL DEFAULT '' COMMENT '应用名称(格式：英文数字-_)，如：nvwa-server',
  `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '项目介绍',
  `deploy_type` tinyint(4) unsigned NOT NULL DEFAULT '1' COMMENT '部署类型，1：Git 项目部署；2：Jenkins 持续集成打包部署；',
  `app_type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '应用类型，1：自定义项目；2：SpringBoot；3：NodeJs；4：PM2 部署；',
  `app_config` text NOT NULL COMMENT '（应用类型为非自定义）项目场景化部署的参数配置',
  `repo_url` varchar(1024) NOT NULL DEFAULT '' COMMENT '仓库地址，SSH 协议或 HTTP 协议',
  `repo_username` varchar(256) NOT NULL DEFAULT '' COMMENT 'HTTP 协议，仓库用户名',
  `repo_password` varchar(1024) NOT NULL DEFAULT '' COMMENT 'HTTP 协议，用户密码（ base64 encode）',
  `public_key` varchar(2048) NOT NULL DEFAULT '' COMMENT 'SSH 协议，指定使用的密钥',
  `repo_type` varchar(16) NOT NULL DEFAULT '1' COMMENT 'git/svn',
  `local_repo_workspace` varchar(1024) NOT NULL DEFAULT '' COMMENT '[本地]本地 clone 的工作空间',
  `local_build_workspace` varchar(1024) NOT NULL DEFAULT '' COMMENT '[本地]本地构建时，用来执行构建的临时工作空间',
  `local_pkg_workspace` varchar(1024) NOT NULL DEFAULT '' COMMENT '[本地]本地构建完打包的应用版本包根路径',
  `excludes` text NOT NULL COMMENT '不参与打包的文件，逗号分隔',
  `files` text NOT NULL COMMENT '指定要打包的文件',
  `cmd_build` text NOT NULL COMMENT '构建命令（Shell）',
  `cmd_before_deploy` text NOT NULL COMMENT '部署前命令',
  `cmd_after_deploy` text NOT NULL COMMENT '部署后命令',
  `cmd_health_check` text NOT NULL COMMENT '应用部署完健康检查的命令',
  `cmd_online` text NOT NULL COMMENT '服务实例切到线上的命令',
  `cmd_timeout` int(10) NOT NULL DEFAULT '600' COMMENT '命令执行的超时时间，单位：秒',
  `cmd_get_pid` varchar(1024) NOT NULL DEFAULT '' COMMENT '获取进程 ID 的 shell 命令',
  `deploy_user` varchar(64) NOT NULL COMMENT '部署应用的主机用户',
  `deploy_path` varchar(1024) NOT NULL COMMENT '[服务器]部署应用的路径',
  `remote_pkg_workspace` varchar(1024) NOT NULL DEFAULT '' COMMENT '[服务器]服务器版本包保存根路径',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1：正常；2：删除',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `project_id` (`project_id`),
  KEY `name` (`name`(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用表';

LOCK TABLES `app` WRITE;
/*!40000 ALTER TABLE `app` DISABLE KEYS */;

INSERT INTO `app` (`id`, `uid`, `project_id`, `name`, `description`, `deploy_type`, `app_type`, `app_config`, `repo_url`, `repo_username`, `repo_password`, `public_key`, `repo_type`, `local_repo_workspace`, `local_build_workspace`, `local_pkg_workspace`, `excludes`, `files`, `cmd_build`, `cmd_before_deploy`, `cmd_after_deploy`, `cmd_health_check`, `cmd_online`, `cmd_timeout`, `cmd_get_pid`, `deploy_user`, `deploy_path`, `remote_pkg_workspace`, `enabled`, `ctime`, `utime`)
VALUES
	(1,1,8,'demo-01','测试 APP',1,2,'{}','git@code.aliyun.com:542792857/demo-01.git','nvwa-io','@nvwa-io','&&&KHDKKDDJkdjka','git','/data/nvwa/repos/demo-01','/data/nvwa/builds/demo-01','/data/nvwa/packages/demo-01','.git\nREADME.md','','echo \'hello world\'','echo \'before deploy\'','echo \'after deploy\'','echo \'health check\'','echo \'online\'',3600,'','root','/data/nvwa/deploys/demo-01','/data/nvwa/packages/demo-01',1,'2019-01-01 22:57:49','2019-01-09 00:46:55'),
	(2,1,8,'demo-02','测试 APP',1,2,'{}','git@code.aliyun.com:542792857/demo-01.git','nvwa-io','@nvwa-io','&&&KHDKKDDJkdjka','git','/data/nvwa/repos/demo-01','/data/nvwa/builds/demo-01','/data/nvwa/packages/demo-01','.git\nREADME.md','','echo \'hello world\'','echo \'before deploy\'','echo \'after deploy\'','echo \'health check\'','echo \'online\'',3600,'','root','/data/nvwa/deploys/demo-01','/data/nvwa/packages/demo-01',1,'2019-01-02 22:14:13','2019-01-09 00:46:53'),
	(4,1,8,'demo-001','demo-001',0,0,'','http://github.com','','','','','','','','','','','echo \'cmd before deploy\'','','','',3600,'','nvwa','/data/nvwa/deploys/demo-001','',1,'2019-01-19 18:41:19','2019-01-19 22:48:51'),
	(5,1,8,'demo-002','demo-001',0,0,'','http://github.com','','','','','','','','.git\nREADME.md','file01\nfile02','echo \'build\'','echo \'before deploy\'','echo \'after deploy\'','','',3600,'','nvwa','/data/nvwa/deploys/demo-002','',1,'2019-01-19 18:44:36','2019-01-19 23:36:52');

/*!40000 ALTER TABLE `app` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table audit
# ------------------------------------------------------------

DROP TABLE IF EXISTS `audit`;

CREATE TABLE `audit` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `deployment_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '部署单 ID',
  `project_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '项目 ID',
  `app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '应用 ID',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户 ID，发起人',
  `audit_uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '审核人 uid',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '10' COMMENT '审核单状态，10：待审核；40：通过；50：驳回；60：取消；',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1：正常；2：删除',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`,`status`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='审核表';



# Dump of table build
# ------------------------------------------------------------

DROP TABLE IF EXISTS `build`;

CREATE TABLE `build` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '应用 ID',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户 ID',
  `branch` varchar(256) NOT NULL DEFAULT '' COMMENT '所选分支',
  `tag` varchar(256) NOT NULL DEFAULT '' COMMENT '所选的 Tag',
  `commit_id` varchar(64) NOT NULL DEFAULT '' COMMENT '构建的 commit_id',
  `package_name` varchar(256) NOT NULL DEFAULT '' COMMENT '构建完打包的包名',
  `log` text NOT NULL COMMENT '构建日志',
  `jenkins_build_num` int(10) unsigned NOT NULL COMMENT 'Jenkins 构建号',
  `notified` tinyint(4) NOT NULL COMMENT '构建结束，是否已经通知',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '10' COMMENT '构建状态，10：创建任务；20：构建中；30：构建成功；40：构建失败；50：版本包保存成功；60：版本包保存失败；70：完成',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='应用的 Jenkins 持续构建的构建记录';

LOCK TABLES `build` WRITE;
/*!40000 ALTER TABLE `build` DISABLE KEYS */;

INSERT INTO `build` (`id`, `app_id`, `uid`, `branch`, `tag`, `commit_id`, `package_name`, `log`, `jenkins_build_num`, `notified`, `status`, `enabled`, `ctime`, `utime`)
VALUES
	(1,1,1,'master','','abcdefgh','','',5,0,70,1,'2019-01-02 21:12:59','2019-01-12 15:58:19'),
	(2,1,1,'master','','abcdefgh','','',0,0,60,1,'2019-01-02 21:13:45','2019-01-21 02:10:23'),
	(3,1,1,'master','','abcdefgh','','',0,0,50,1,'2019-01-02 21:13:55','2019-01-21 02:10:25'),
	(4,1,1,'master','','abcdefgh','','',0,0,40,1,'2019-01-02 21:14:23','2019-01-21 02:10:27'),
	(5,1,1,'master','','abcdefgh','','',0,0,30,1,'2019-01-02 21:14:24','2019-01-21 02:10:30'),
	(6,1,1,'master','','abcdefgh','demo-01.6.master.9cf0947c.20190107010713.tar.gz','[app] app id: 1\nApp demo-01 repository updated ... ok\n\necho \'hello world\'\nhello world\n\n[pack version package]\ncd /data/nvwa/builds/demo-01/build_id_6 && tar -p  --exclude=\'.git\'  --exclude=\'README.md\'  -cz -f \'/data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713.tar.gz\' *\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n[app] app id: 1\nApp demo-01 repository updated ... ok\n\necho \'hello world\'\nhello world\n\n[pack version package]\ncd /data/nvwa/builds/demo-01/build_id_6 && tar -p  --exclude=\'.git\'  --exclude=\'README.md\'  -cz -f \'/data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713.tar.gz\' *\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n[app] app id: 1\nApp demo-01 repository updated ... ok\n\necho \'hello world\'\nhello world\n\n[pack version package]\ncd /data/nvwa/builds/demo-01/build_id_6 && tar -p  --exclude=\'.git\'  --exclude=\'README.md\'  -cz -f \'/data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713.tar.gz\' *\n',0,0,20,1,'2019-01-06 21:37:18','2019-01-21 02:31:54'),
	(7,1,1,'f-01','','','','',0,0,10,1,'2019-01-25 01:16:56','2019-01-25 01:16:56'),
	(8,1,1,'f-02','','','','',0,0,10,1,'2019-01-25 01:18:07','2019-01-25 01:18:07'),
	(9,1,1,'master','','','','',0,0,10,1,'2019-01-25 01:18:21','2019-01-25 01:18:21'),
	(10,1,1,'f-04','','','','',0,0,10,1,'2019-01-25 01:19:42','2019-01-25 01:19:42'),
	(11,1,1,'master','','','','',0,0,10,1,'2019-01-25 02:33:24','2019-01-25 02:33:24');

/*!40000 ALTER TABLE `build` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table cluster
# ------------------------------------------------------------

DROP TABLE IF EXISTS `cluster`;

CREATE TABLE `cluster` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '应用 ID',
  `env_id` int(10) unsigned NOT NULL DEFAULT '0',
  `uid` int(10) unsigned NOT NULL DEFAULT '0',
  `name` varchar(128) DEFAULT NULL COMMENT '分组名称',
  `hosts` text COMMENT '目标机器列表，逗号分隔',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='主机分组';

LOCK TABLES `cluster` WRITE;
/*!40000 ALTER TABLE `cluster` DISABLE KEYS */;

INSERT INTO `cluster` (`id`, `app_id`, `env_id`, `uid`, `name`, `hosts`, `enabled`, `ctime`, `utime`)
VALUES
	(1,1,1,1,'开发集群01','localhost,app01.bjgt.nvwa.io,m01',1,'2019-01-02 01:54:25','2019-01-02 01:54:25'),
	(2,2,2,1,'默认分组','m01',1,'2019-01-19 18:41:23','2019-01-02 22:14:13'),
	(3,2,3,1,'默认分组','localhost,app01.bjgt.nvwa.io',1,'2019-01-19 18:41:23','2019-01-02 22:14:13'),
	(4,2,4,1,'默认分组','',1,'2019-01-19 18:41:23','2019-01-02 22:14:13'),
	(5,2,5,1,'默认分组','',1,'2019-01-19 18:41:23','2019-01-02 22:14:13'),
	(6,4,6,1,'默认分组','',1,'2019-01-19 18:41:23','2019-01-19 18:41:23'),
	(7,4,7,1,'默认分组','',1,'2019-01-19 18:41:23','2019-01-19 18:41:23'),
	(8,4,8,1,'默认分组','',1,'2019-01-19 18:41:23','2019-01-19 18:41:23'),
	(9,4,9,1,'默认分组','',1,'2019-01-19 18:41:23','2019-01-19 18:41:23'),
	(10,5,10,1,'默认分组','',1,'2019-01-19 18:44:36','2019-01-19 18:44:36'),
	(11,5,11,1,'默认分组','',1,'2019-01-19 18:44:36','2019-01-19 18:44:36'),
	(12,5,12,1,'默认分组','',1,'2019-01-19 18:44:36','2019-01-19 18:44:36'),
	(13,5,13,1,'默认分组','',1,'2019-01-19 18:44:36','2019-01-19 18:44:36'),
	(15,1,16,1,'默认分组','me01,m30',1,'2019-01-20 23:25:41','2019-01-20 23:25:41'),
	(17,1,1,1,'开发集群 02','m02,m03',1,'2019-01-21 00:31:51','2019-01-21 00:31:51'),
	(18,1,16,1,'测试集群 02','m02',1,'2019-01-21 00:33:46','2019-01-21 00:33:46'),
	(19,1,16,1,'测试集群 03','m09',1,'2019-01-21 00:36:57','2019-01-21 00:36:57');

/*!40000 ALTER TABLE `cluster` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table deployment
# ------------------------------------------------------------

DROP TABLE IF EXISTS `deployment`;

CREATE TABLE `deployment` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int(10) unsigned NOT NULL DEFAULT '0',
  `app_id` int(10) unsigned NOT NULL DEFAULT '0',
  `uid` int(10) unsigned NOT NULL DEFAULT '0',
  `env_id` int(10) unsigned NOT NULL DEFAULT '0',
  `cluster_ids` varchar(512) NOT NULL COMMENT '部署单要部署的分组 id',
  `cluster_hosts` varchar(2048) NOT NULL COMMENT '分组对应的主机',
  `pkg_id` int(10) unsigned NOT NULL COMMENT '版本包的 ID',
  `pkg` varchar(512) NOT NULL DEFAULT '' COMMENT '待部署的版本包',
  `is_auto_deploy` tinyint(3) unsigned NOT NULL DEFAULT '2' COMMENT '自动部署，1：自动部署；2：手动部署；',
  `is_all_cluster` tinyint(11) unsigned NOT NULL DEFAULT '1' COMMENT '1：所有分组；2：指定分组；',
  `link_id` varchar(512) DEFAULT NULL COMMENT '上线的软链名',
  `branch` varchar(128) DEFAULT NULL COMMENT 'Git 项目部署，所选分支',
  `commit_id` varchar(128) DEFAULT NULL COMMENT 'Git 项目所选 commit_id',
  `file_list` text COMMENT 'Git 项目，增量发布，所选择的文件列表',
  `file_deploy_mode` tinyint(3) unsigned NOT NULL COMMENT 'Git 项目，上线文件模式: 1.全量所有文件 2.指定文件列表',
  `latest_link_id` varchar(512) DEFAULT '' COMMENT 'Git 项目，增量发布类型，指定上一个版本的软链号',
  `is_need_audit` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否需要审核',
  `status` tinyint(3) unsigned NOT NULL DEFAULT '10' COMMENT '10: 新创建；20：免审核；30：待审核；40：审核通过；50：驳回；60：取消；70: 部署中；80: 完成；90: 失败；',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部署单';

LOCK TABLES `deployment` WRITE;
/*!40000 ALTER TABLE `deployment` DISABLE KEYS */;

INSERT INTO `deployment` (`id`, `project_id`, `app_id`, `uid`, `env_id`, `cluster_ids`, `cluster_hosts`, `pkg_id`, `pkg`, `is_auto_deploy`, `is_all_cluster`, `link_id`, `branch`, `commit_id`, `file_list`, `file_deploy_mode`, `latest_link_id`, `is_need_audit`, `status`, `enabled`, `ctime`, `utime`)
VALUES
	(1,8,2,1,3,'1,2,3','{\"3\":[\"\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,10,1,'2019-01-03 00:59:58','2019-01-18 02:02:40'),
	(2,8,2,1,3,'1,2,3','{\"3\":[\"localhost\",\"app01.bjgt.nvwa.io\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,20,1,'2019-01-03 01:02:47','2019-01-23 22:50:36'),
	(3,8,2,1,3,'1,2,3','{\"3\":[\"localhost\",\"app01.bjgt.nvwa.io\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,30,1,'2019-01-03 01:02:47','2019-01-23 22:50:39'),
	(4,8,2,1,3,'1,2,3','{\"3\":[\"localhost\",\"app01.bjgt.nvwa.io\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,30,1,'2019-01-03 01:02:47','2019-01-23 22:50:42'),
	(5,8,2,1,3,'1,2,3','{\"3\":[\"localhost\",\"app01.bjgt.nvwa.io\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,50,1,'2019-01-03 01:02:47','2019-01-23 22:50:46'),
	(6,8,2,1,3,'1,2,3','{\"3\":[\"localhost\",\"app01.bjgt.nvwa.io\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,60,1,'2019-01-03 01:02:47','2019-01-23 22:50:48'),
	(7,8,2,1,3,'1,2,3','{\"3\":[\"localhost\",\"app01.bjgt.nvwa.io\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,70,1,'2019-01-03 01:02:47','2019-01-23 22:50:51'),
	(8,8,2,1,3,'1,2,3','{\"3\":[\"localhost\",\"app01.bjgt.nvwa.io\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,80,1,'2019-01-03 01:02:47','2019-01-23 22:50:54'),
	(9,8,2,1,3,'1,2,3','{\"3\":[\"localhost\",\"app01.bjgt.nvwa.io\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,90,1,'2019-01-03 01:02:47','2019-01-23 22:50:58'),
	(10,8,2,1,3,'1,2,3','{\"3\":[\"localhost\",\"app01.bjgt.nvwa.io\"]}',1,'demo-01.6.master.9cf0947c.20190107010713.tar.gz',0,0,'','master','dbescd','',0,'',0,10,1,'2019-01-26 17:55:26','2019-01-26 17:55:29');

/*!40000 ALTER TABLE `deployment` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table env
# ------------------------------------------------------------

DROP TABLE IF EXISTS `env`;

CREATE TABLE `env` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建用户的 ID',
  `app_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '项目 ID',
  `name` varchar(128) NOT NULL COMMENT '环境名称，如：线上环境',
  `permit_branches` varchar(512) NOT NULL DEFAULT '' COMMENT '允许使用的分支包，* 表示所有',
  `is_auto_deploy` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否自动部署，1：自动；0：手动',
  `is_need_audit` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否需要审核，1：需要；0：不需要',
  `cmd_env` text NOT NULL COMMENT '环境差异化命令',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用的部署环境';

LOCK TABLES `env` WRITE;
/*!40000 ALTER TABLE `env` DISABLE KEYS */;

INSERT INTO `env` (`id`, `uid`, `app_id`, `name`, `permit_branches`, `is_auto_deploy`, `is_need_audit`, `cmd_env`, `enabled`, `ctime`, `utime`)
VALUES
	(1,1,1,'开发环境','*',0,0,'export ENV=dev',1,'2019-01-02 01:18:53','2019-01-24 02:30:24'),
	(2,1,2,'开发环境','*',0,1,'',1,'0000-00-00 00:00:00','2019-01-24 02:46:10'),
	(3,1,2,'测试环境','*',0,0,'',1,'0000-00-00 00:00:00','2019-01-02 22:14:13'),
	(4,1,2,'预发布环境','*',0,0,'',1,'0000-00-00 00:00:00','2019-01-02 22:14:13'),
	(5,1,2,'生成环境','*',0,0,'',1,'0000-00-00 00:00:00','2019-01-02 22:14:13'),
	(6,1,4,'开发环境','*',0,0,'',1,'2019-01-19 18:41:23','2019-01-19 18:41:23'),
	(7,1,4,'测试环境','*',0,0,'',1,'2019-01-19 18:41:23','2019-01-19 18:41:23'),
	(8,1,4,'预发布环境','*',0,0,'',1,'2019-01-19 18:41:23','2019-01-19 18:41:23'),
	(9,1,4,'生成环境','*',0,0,'',1,'2019-01-19 18:41:23','2019-01-19 18:41:23'),
	(10,1,5,'开发环境','*',0,0,'',1,'2019-01-19 18:44:36','2019-01-19 18:44:36'),
	(11,1,5,'测试环境','*',0,0,'',1,'2019-01-19 18:44:36','2019-01-19 18:44:36'),
	(12,1,5,'预发布环境','*',0,0,'',1,'2019-01-19 18:44:36','2019-01-19 18:44:36'),
	(13,1,5,'生成环境','*',0,0,'',1,'2019-01-19 18:44:36','2019-01-19 18:44:36'),
	(16,1,1,'测试环境','*',0,0,'export ENV=prod',1,'2019-01-20 23:25:41','2019-01-20 23:40:55');

/*!40000 ALTER TABLE `env` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table event
# ------------------------------------------------------------

DROP TABLE IF EXISTS `event`;

CREATE TABLE `event` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '项目 ID',
  `app_id` int(10) unsigned NOT NULL COMMENT '应用 ID',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户 ID',
  `content` text NOT NULL COMMENT '事件内容',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1：正常；2：删除',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `project_id` (`project_id`),
  KEY `app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='成员操作事件';



# Dump of table job
# ------------------------------------------------------------

DROP TABLE IF EXISTS `job`;

CREATE TABLE `job` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `deployment_id` int(11) unsigned NOT NULL COMMENT '部署单 ID',
  `app_id` int(21) unsigned NOT NULL DEFAULT '0' COMMENT '项目id',
  `env_id` int(11) unsigned NOT NULL DEFAULT '0',
  `cluster_id` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '分组 ID',
  `all_hosts` text NOT NULL COMMENT '创建该任务时，分组已有的机器',
  `deploy_hosts` text NOT NULL COMMENT '需要部署的主机',
  `exclude_hosts` text NOT NULL COMMENT '不用部署的主机列表，逗号分隔',
  `status` smallint(1) NOT NULL DEFAULT '0' COMMENT '状态,10：新创建，20: 开始部署, 30：部署中，40：成功，50：失败',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `deployment_id` (`deployment_id`),
  KEY `app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='部署单对应的要部署的任务';

LOCK TABLES `job` WRITE;
/*!40000 ALTER TABLE `job` DISABLE KEYS */;

INSERT INTO `job` (`id`, `deployment_id`, `app_id`, `env_id`, `cluster_id`, `all_hosts`, `deploy_hosts`, `exclude_hosts`, `status`, `enabled`, `ctime`, `utime`)
VALUES
	(1,1,2,3,3,'','','',10,1,'2019-01-03 00:59:58','2019-01-03 00:59:58'),
	(2,2,2,3,3,'47.97.217.198','47.97.217.198','',40,1,'2019-01-23 00:39:14','2019-01-23 01:26:56'),
	(3,10,2,3,3,'47.97.217.198','47.97.217.198','',20,1,'2019-01-23 00:39:14','2019-01-26 17:06:57'),
	(4,10,2,3,3,'47.97.217.198','47.97.217.198','',20,1,'2019-01-23 00:39:14','2019-01-23 22:52:54'),
	(5,10,2,3,3,'47.97.217.198','47.97.217.198','',30,1,'2019-01-23 00:39:14','2019-01-23 22:52:57'),
	(6,10,2,3,3,'47.97.217.198','47.97.217.198','',40,1,'2019-01-23 00:39:14','2019-01-23 22:52:42'),
	(7,10,2,3,3,'47.97.217.198','47.97.217.198','',50,1,'2019-01-23 00:39:14','2019-01-23 22:53:03'),
	(8,10,2,3,3,'47.97.217.198','47.97.217.198','',40,1,'2019-01-23 00:39:14','2019-01-23 22:52:44');

/*!40000 ALTER TABLE `job` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table job_step
# ------------------------------------------------------------

DROP TABLE IF EXISTS `job_step`;

CREATE TABLE `job_step` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `job_id` bigint(11) NOT NULL COMMENT '任务 id',
  `app_id` int(11) NOT NULL COMMENT '应用 id',
  `deployment_id` int(11) unsigned NOT NULL COMMENT '部署单 id',
  `cmd` text CHARACTER SET utf8 NOT NULL COMMENT '运行命令',
  `log` text CHARACTER SET utf8 COMMENT '执行日志',
  `consume` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '耗时，单位ms',
  `step` int(3) unsigned NOT NULL DEFAULT '10' COMMENT '任务当前执行到哪个步骤',
  `status` smallint(1) NOT NULL DEFAULT '1' COMMENT '当前步骤执行的状态 10：处理中，20：成功；30：失败',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `job_id` (`job_id`),
  KEY `app_id` (`app_id`),
  KEY `deployment_id` (`deployment_id`),
  KEY `jobid_step` (`job_id`,`step`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部署任务的记录';

LOCK TABLES `job_step` WRITE;
/*!40000 ALTER TABLE `job_step` DISABLE KEYS */;

INSERT INTO `job_step` (`id`, `job_id`, `app_id`, `deployment_id`, `cmd`, `log`, `consume`, `step`, `status`, `enabled`, `ctime`, `utime`)
VALUES
	(42,2,2,2,'ANSIBLE_SSH_ARGS=\'-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o CheckHostIP=false\' ansible all -m shell -a \'mkdir -p /data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713 warn=False\' -f 10 -i /tmp/nvwa/tmp_ansible_i_1548177957646204000 -T 10','47.97.217.198 | SUCCESS | rc=0 >>\n\n',7444,10,20,1,'2019-01-23 01:25:57','2019-01-23 01:25:57'),
	(43,2,2,2,'ANSIBLE_SSH_ARGS=\'-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o CheckHostIP=false\' ansible all -m copy -u root -a \'src=/data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713.tar.gz dest=/data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713.tar.gz warn=False\' -f 10 -i /tmp/nvwa/tmp_ansible_i_1548177965100661000 -T 10','47.97.217.198 | SUCCESS => {\n    \"changed\": false, \n    \"checksum\": \"effd45c25146d8775ff9e2eb7371694ede52d4c7\", \n    \"dest\": \"/data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713.tar.gz\", \n    \"gid\": 0, \n    \"group\": \"root\", \n    \"mode\": \"0644\", \n    \"owner\": \"root\", \n    \"path\": \"/data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713.tar.gz\", \n    \"size\": 139, \n    \"state\": \"file\", \n    \"uid\": 0\n}\n',9250,20,20,1,'2019-01-23 01:26:05','2019-01-23 01:26:05'),
	(44,2,2,2,'ANSIBLE_SSH_ARGS=\'-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o CheckHostIP=false\' ansible all -m shell -a \'cd /data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713 && tar --no-same-owner -pm -C /data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713 -xz -f /data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713.tar.gz warn=False\' -f 10 -i /tmp/nvwa/tmp_ansible_i_1548177974343498000 -T 10','47.97.217.198 | SUCCESS | rc=0 >>\n\n',6630,30,20,1,'2019-01-23 01:26:14','2019-01-23 01:26:14'),
	(45,2,2,2,'ANSIBLE_SSH_ARGS=\'-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o CheckHostIP=false\' ansible all -m shell -a \'cd /data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713; echo \'\\\'\'before deploy\'\\\'\' warn=False\' -f 10 -i /tmp/nvwa/tmp_ansible_i_1548177980978725000 -T 10','47.97.217.198 | SUCCESS | rc=0 >>\nbefore deploy\n',7410,40,20,1,'2019-01-23 01:26:20','2019-01-23 01:26:20'),
	(46,2,2,2,'ANSIBLE_SSH_ARGS=\'-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o CheckHostIP=false\' ansible all -m shell -a \'mkdir -p /data/nvwa/deploys && rm -rf /data/nvwa/deploys/demo-01.tmp && ln -sfn /data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713 /data/nvwa/deploys/demo-01.tmp && chown -h root /data/nvwa/deploys/demo-01.tmp && mv -fT /data/nvwa/deploys/demo-01.tmp /data/nvwa/deploys/demo-01 warn=False\' -f 10 -i /tmp/nvwa/tmp_ansible_i_1548177988390602000 -T 10','47.97.217.198 | SUCCESS | rc=0 >>\n\n',6212,50,20,1,'2019-01-23 01:26:28','2019-01-23 01:26:28'),
	(47,2,2,2,'ANSIBLE_SSH_ARGS=\'-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o CheckHostIP=false\' ansible all -m shell -a \'cd /data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713; echo \'\\\'\'before deploy\'\\\'\' warn=False\' -f 10 -i /tmp/nvwa/tmp_ansible_i_1548177994605467000 -T 10','47.97.217.198 | SUCCESS | rc=0 >>\nbefore deploy\n',7338,60,20,1,'2019-01-23 01:26:34','2019-01-23 01:26:34'),
	(48,2,2,2,'ANSIBLE_SSH_ARGS=\'-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o CheckHostIP=false\' ansible all -m shell -a \'cd /data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713; echo \'\\\'\'health check\'\\\'\' warn=False\' -f 10 -i /tmp/nvwa/tmp_ansible_i_1548178001945732000 -T 10','47.97.217.198 | SUCCESS | rc=0 >>\nhealth check\n',7504,70,20,1,'2019-01-23 01:26:41','2019-01-23 01:26:41'),
	(49,2,2,2,'ANSIBLE_SSH_ARGS=\'-o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -o CheckHostIP=false\' ansible all -m shell -a \'cd /data/nvwa/packages/demo-01/demo-01.6.master.9cf0947c.20190107010713; echo \'\\\'\'online\'\\\'\' warn=False\' -f 10 -i /tmp/nvwa/tmp_ansible_i_1548178009450299000 -T 10','47.97.217.198 | SUCCESS | rc=0 >>\nonline\n',7191,80,20,1,'2019-01-23 01:26:49','2019-01-23 01:26:49'),
	(50,2,2,2,'clean','',2,90,20,1,'2019-01-23 01:26:56','2019-01-23 01:26:56');

/*!40000 ALTER TABLE `job_step` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table member
# ------------------------------------------------------------

DROP TABLE IF EXISTS `member`;

CREATE TABLE `member` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `project_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '项目 ID',
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户 ID',
  `project_role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '用户角色 ID',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1：正常；2：删除',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `project_id` (`project_id`,`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='项目成员表';

LOCK TABLES `member` WRITE;
/*!40000 ALTER TABLE `member` DISABLE KEYS */;

INSERT INTO `member` (`id`, `project_id`, `uid`, `project_role_id`, `enabled`, `ctime`, `utime`)
VALUES
	(1,8,1,2,1,'2018-12-28 01:40:05','2019-01-27 11:01:46'),
	(2,9,1,2,1,'2019-01-01 21:20:59','2019-01-27 02:51:03');

/*!40000 ALTER TABLE `member` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table pkg
# ------------------------------------------------------------

DROP TABLE IF EXISTS `pkg`;

CREATE TABLE `pkg` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `app_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '应用 ID',
  `build_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '所属构建 ID',
  `name` varchar(1024) NOT NULL DEFAULT '' COMMENT '应用版本包名称',
  `branch` varchar(256) NOT NULL DEFAULT '' COMMENT '所属分支',
  `tag` varchar(256) NOT NULL DEFAULT '' COMMENT '所属 Tag',
  `commit_id` varchar(256) NOT NULL DEFAULT '' COMMENT 'commit_id',
  `storage_type` varchar(64) NOT NULL DEFAULT '' COMMENT '版本包存储类型',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1：正常；2：删除',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `app_id` (`app_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='应用版本包记录';

LOCK TABLES `pkg` WRITE;
/*!40000 ALTER TABLE `pkg` DISABLE KEYS */;

INSERT INTO `pkg` (`id`, `app_id`, `build_id`, `name`, `branch`, `tag`, `commit_id`, `storage_type`, `enabled`, `ctime`, `utime`)
VALUES
	(1,2,1,'demo-01.6.master.9cf0947c.20190106230648.tar.gz','master','','dbescd','oss',1,'2019-01-03 00:57:50','2019-01-18 00:11:21'),
	(2,1,6,'demo-01.6.master.9cf0947c.20190106230648.tar.gz','master','','dbescd','oss',1,'2019-01-06 23:06:48','2019-01-18 00:11:55'),
	(3,1,6,'demo-01.6.master.9cf0947c.20190106231751.tar.gz','master','','dbescd','oss',1,'2019-01-06 23:17:51','2019-01-18 00:11:56'),
	(4,1,6,'demo-01.6.master.9cf0947c.20190107004855.tar.gz','master','','dbescd','oss',1,'2019-01-07 00:48:55','2019-01-18 00:11:56'),
	(5,1,6,'demo-01.6.master.9cf0947c.20190107010713.tar.gz','master','','dbescd','oss',1,'2019-01-07 01:07:13','2019-01-18 00:11:58'),
	(8,1,1,'nvwa-demo.1.master.abcdefgh.20190112142046.tar.gz','master','','abcdefgh','oss',1,'2019-01-12 16:02:58','2019-01-18 00:11:43');

/*!40000 ALTER TABLE `pkg` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table project
# ------------------------------------------------------------

DROP TABLE IF EXISTS `project`;

CREATE TABLE `project` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '创建者 ID',
  `name` varchar(256) NOT NULL DEFAULT '' COMMENT '项目名称，如：女娲项目',
  `description` varchar(1024) NOT NULL DEFAULT '' COMMENT '项目介绍',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1：正常；2：删除',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `name` (`name`(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='项目表';

LOCK TABLES `project` WRITE;
/*!40000 ALTER TABLE `project` DISABLE KEYS */;

INSERT INTO `project` (`id`, `uid`, `name`, `description`, `enabled`, `ctime`, `utime`)
VALUES
	(8,1,'女娲部署系统','OK, 描述',1,'2018-12-28 01:40:05','2019-01-16 23:06:37'),
	(9,1,'容器化项目','OK, 描述',1,'2018-12-28 01:40:05','2019-01-16 23:06:37');

/*!40000 ALTER TABLE `project` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table project_role
# ------------------------------------------------------------

DROP TABLE IF EXISTS `project_role`;

CREATE TABLE `project_role` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(256) NOT NULL DEFAULT '' COMMENT '角色名称',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1：正常；2：删除',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `name` (`name`(191))
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='项目角色表';

LOCK TABLES `project_role` WRITE;
/*!40000 ALTER TABLE `project_role` DISABLE KEYS */;

INSERT INTO `project_role` (`id`, `name`, `enabled`, `ctime`, `utime`)
VALUES
	(2,'项目管理',1,'2018-12-30 14:30:18','2018-12-30 15:33:14'),
	(3,'开发',1,'2018-12-30 14:54:18','2018-12-30 14:54:18'),
	(4,'测试',1,'2018-12-30 14:54:22','2018-12-30 14:54:22');

/*!40000 ALTER TABLE `project_role` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table project_role_perm
# ------------------------------------------------------------

DROP TABLE IF EXISTS `project_role_perm`;

CREATE TABLE `project_role_perm` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `project_role_id` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '项目角色 ID',
  `perm` varchar(64) NOT NULL DEFAULT '' COMMENT '权限',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1：正常；2：删除',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `project_role_id_perm` (`project_role_id`,`perm`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='项目角色权限表';

LOCK TABLES `project_role_perm` WRITE;
/*!40000 ALTER TABLE `project_role_perm` DISABLE KEYS */;

INSERT INTO `project_role_perm` (`id`, `project_role_id`, `perm`, `enabled`, `ctime`, `utime`)
VALUES
	(22,2,'project.create',1,'2018-12-30 16:32:57','2019-01-27 02:16:52'),
	(23,2,'project.update',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(24,2,'member.add',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(25,2,'member.remove',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(26,2,'member.change.role',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(27,2,'app.create',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(28,2,'app.update',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(29,2,'app.delete',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(30,2,'env.create',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(31,2,'env.update',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(32,2,'env.delete',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(33,2,'env.audit',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(34,2,'cluster.create',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(35,2,'cluster.update',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(36,2,'cluster.delete',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(37,2,'deployment.create',1,'2018-12-30 16:32:57','2018-12-30 16:32:57'),
	(38,3,'project.create',1,'2018-12-30 16:33:07','2019-01-27 02:39:09'),
	(39,3,'project.update',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(40,3,'member.add',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(41,3,'member.remove',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(42,3,'member.change.role',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(43,3,'app.create',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(44,3,'app.update',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(45,3,'app.delete',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(46,3,'env.create',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(47,3,'env.update',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(48,3,'env.delete',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(49,3,'env.audit',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(50,3,'cluster.create',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(51,3,'cluster.update',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(52,3,'cluster.delete',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(53,3,'deployment.create',1,'2018-12-30 16:33:07','2018-12-30 16:33:07'),
	(54,4,'project.create',1,'2018-12-30 16:33:17','2019-01-27 02:17:00'),
	(55,4,'project.update',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(56,4,'member.add',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(57,4,'member.remove',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(58,4,'member.change.role',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(59,4,'app.create',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(60,4,'app.update',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(61,4,'app.delete',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(62,4,'env.create',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(63,4,'env.update',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(64,4,'env.delete',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(65,4,'env.audit',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(66,4,'cluster.create',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(67,4,'cluster.update',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(68,4,'cluster.delete',1,'2018-12-30 16:33:17','2018-12-30 16:33:17'),
	(69,4,'deployment.create',1,'2018-12-30 16:33:17','2018-12-30 16:33:17');

/*!40000 ALTER TABLE `project_role_perm` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table system
# ------------------------------------------------------------

DROP TABLE IF EXISTS `system`;

CREATE TABLE `system` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `version` varchar(128) NOT NULL COMMENT 'nvwa 的版本',
  `default_project_role_id` int(10) unsigned NOT NULL DEFAULT '1' COMMENT '创建项目的用户的角色 ID',
  `deploy_root_path` varchar(1024) NOT NULL DEFAULT '/data/nvwa/deploys' COMMENT '默认应用部署目录，应用具体部署路径={deploy_path}/{应用名}',
  `custom_deploy_path` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否允许应用自定义部署路径，1：允许；0：不允许',
  `deploy_user` varchar(64) NOT NULL DEFAULT 'nvwa' COMMENT '默认部署使用的用户',
  `custom_deploy_user` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否允许自定义部署的用户，1：允许；0：不允许；',
  `use_jenkins` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '是否使用 jenkins 进行构建，0：不使用；1：使用',
  `jenkins_url` varchar(256) NOT NULL COMMENT 'jenkins 地址， 如:http://localhost:8080',
  `jenkins_template` text NOT NULL COMMENT '创建 Jenkins 项目的模板',
  `jenkins_user` varchar(128) NOT NULL DEFAULT '' COMMENT '操作 jenkins 的用户',
  `jenkins_password` varchar(128) NOT NULL DEFAULT '' COMMENT '操作 jenkins 的用户密码',
  `pkg_limit` int(11) unsigned NOT NULL DEFAULT '15' COMMENT '保存应用最新版本包数量',
  `custom_pkg_root_path` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '是否允许应用自定义版本包根路径，1：允许；0：不允许',
  `pkg_storage_type` varchar(64) NOT NULL DEFAULT '1' COMMENT '版本包保存方式，local：本机；oss：aliyun oss; cos: tencent cloud cos; aws-s3: aws s3',
  `pkg_storage_config` text NOT NULL COMMENT '版本包保存方式对应的配置',
  `git_ci_auth_type` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '女娲拉去git仓库代码，对于 HTTP 的 URL 有两种方式：账号密码或Token。1：账号密码；2：Token。对于 SSH 的 URL，使用的是部署女娲的主机的 public key',
  `git_ci_user` varchar(256) NOT NULL DEFAULT '' COMMENT '（HTTP）Git  持续集成用户',
  `git_ci_password` varchar(256) NOT NULL DEFAULT '' COMMENT '（HTTP）Git 持续集成用户的密码',
  `git_ci_token` varchar(256) NOT NULL DEFAULT '' COMMENT '（HTTP）持续集成 Token，HTTP 认证方式有两种，一种账号密码；一种通过 Token。',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '1：正常；2：删除',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL ON UPDATE CURRENT_TIMESTAMP,
  `pkg_root_path` varchar(1024) NOT NULL DEFAULT '/data/nvwa/packages' COMMENT '默认版本包目录，具体版本包路径={pkg_path}/{应用名}/版本包.tar.gz',
  `repo_root_path` varchar(1024) NOT NULL DEFAULT '/data/nvwa/repos' COMMENT '本地仓库根路径',
  `build_root_path` varchar(1024) NOT NULL DEFAULT '/data/nvwa/builds' COMMENT '本地仓库构建的临时build 根路径',
  `notify_enable_types` varchar(256) NOT NULL DEFAULT '' COMMENT '所开启的通知类型，同时打开多个用逗号分隔，可选：email、wecat_work',
  `notify_config` text NOT NULL COMMENT '不同通知类型的配置',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='系统配置表';

LOCK TABLES `system` WRITE;
/*!40000 ALTER TABLE `system` DISABLE KEYS */;

INSERT INTO `system` (`id`, `version`, `default_project_role_id`, `deploy_root_path`, `custom_deploy_path`, `deploy_user`, `custom_deploy_user`, `use_jenkins`, `jenkins_url`, `jenkins_template`, `jenkins_user`, `jenkins_password`, `pkg_limit`, `custom_pkg_root_path`, `pkg_storage_type`, `pkg_storage_config`, `git_ci_auth_type`, `git_ci_user`, `git_ci_password`, `git_ci_token`, `enabled`, `ctime`, `utime`, `pkg_root_path`, `repo_root_path`, `build_root_path`, `notify_enable_types`, `notify_config`)
VALUES
	(1,'1.0.0',2,'/data/nvwa/deploys',0,'nvwa',0,1,'http://192.168.34.10:8080','<project>\n    <actions/>\n    <description>{{.App.Description}}</description>\n    <keepDependencies>false</keepDependencies>\n    <properties>\n        <jenkins.model.BuildDiscarderProperty>\n            <strategy class=\"hudson.tasks.LogRotator\">\n                <daysToKeep>3</daysToKeep>\n                <numToKeep>7</numToKeep>\n                <artifactDaysToKeep>-1</artifactDaysToKeep>\n                <artifactNumToKeep>-1</artifactNumToKeep>\n            </strategy>\n        </jenkins.model.BuildDiscarderProperty>\n        <hudson.model.ParametersDefinitionProperty>\n            <parameterDefinitions>\n                <net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition plugin=\"git-parameter@0.9.3\">\n                    <name>BUILD_BRANCH</name>\n                    <description/>\n                    <uuid>073cbea8-233b-4fc8-8b26-f9f79d7ebdf0</uuid>\n                    <type>PT_BRANCH</type>\n                    <branch/>\n                    <tagFilter>*</tagFilter>\n                    <branchFilter>.*</branchFilter>\n                    <sortMode>NONE</sortMode>\n                    <defaultValue>master</defaultValue>\n                    <selectedValue>NONE</selectedValue>\n                    <quickFilterEnabled>false</quickFilterEnabled>\n                    <listSize>5</listSize>\n                </net.uaznia.lukanus.hudson.plugins.gitparameter.GitParameterDefinition>\n                <hudson.model.StringParameterDefinition>\n                    <name>BUILD_ID</name>\n                    <description/>\n                    <defaultValue>0</defaultValue>\n                    <trim>false</trim>\n                </hudson.model.StringParameterDefinition>\n            </parameterDefinitions>\n        </hudson.model.ParametersDefinitionProperty>\n    </properties>\n    <scm class=\"hudson.plugins.git.GitSCM\" plugin=\"git@3.9.1\">\n        <configVersion>2</configVersion>\n        <userRemoteConfigs>\n            <hudson.plugins.git.UserRemoteConfig>\n                <url>\n                    {{.App.RepoUrl}}\n                </url>\n                <credentialsId>{{.JenkinsCredentialId}}</credentialsId>\n            </hudson.plugins.git.UserRemoteConfig>\n        </userRemoteConfigs>\n        <branches>\n            <hudson.plugins.git.BranchSpec>\n                <name>${BUILD_BRANCH}</name>\n            </hudson.plugins.git.BranchSpec>\n        </branches>\n        <doGenerateSubmoduleConfigurations>false</doGenerateSubmoduleConfigurations>\n        <submoduleCfg class=\"list\"/>\n        <extensions/>\n    </scm>\n    <canRoam>true</canRoam>\n    <disabled>false</disabled>\n    <blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>\n    <blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding>\n    <triggers/>\n    <concurrentBuild>false</concurrentBuild>\n    <builders>\n        <hudson.tasks.Shell>\n            <command>\n                export TIMESTAMP=`date \'+%Y%m%d%H%M%S\'`\n                echo PKG_NAME=${GITLAB_BRANCH}.${GIT_COMMIT:0:6}.${BUILD_NUMBER}.$TIMESTAMP > nvwa-env-inject-properties\n            </command>\n        </hudson.tasks.Shell>\n        <EnvInjectBuilder plugin=\"envinject@2.1.6\">\n            <info>\n                <propertiesFilePath>nvwa-env-inject-properties</propertiesFilePath>\n            </info>\n        </EnvInjectBuilder>\n        <hudson.tasks.Shell>\n            <command>\n                echo {{.App.RepoUrl}}\n                echo ${JOB_NAME}\n                # do some build jobs\n                {{.App.CmdBuild}}\n            </command>\n        </hudson.tasks.Shell>\n    </builders>\n    <publishers>\n        <org.jenkinsci.plugins.postbuildscript.PostBuildScript plugin=\"postbuildscript@2.7.0\">\n            <config>\n                <scriptFiles/>\n                <groovyScripts/>\n                <buildSteps>\n                    <org.jenkinsci.plugins.postbuildscript.model.PostBuildStep>\n                        <results>\n                            <string>SUCCESS</string>\n                        </results>\n                        <role>BOTH</role>\n                        <buildSteps>\n                            <hudson.tasks.Shell>\n                                <command>\n                                    echo \"hulk-container-agent buildJob update --subProjectName=${JOB_NAME} --buildJobId=${BUILD_ID} --status=0\"\n                                    echo \"hulk-container-agent buildJob notify --status=0 repoTag=${PKG_NAME}\"\n                                </command>\n                            </hudson.tasks.Shell>\n                        </buildSteps>\n                    </org.jenkinsci.plugins.postbuildscript.model.PostBuildStep>\n                    <org.jenkinsci.plugins.postbuildscript.model.PostBuildStep>\n                        <results>\n                            <string>FAILURE</string>\n                        </results>\n                        <role>BOTH</role>\n                        <buildSteps>\n                            <hudson.tasks.Shell>\n                                <command>\n                                    echo \"hulk-container-agent buildJob update --subProjectName=${JOB_NAME} --buildJobId=${BUILD_ID} --status=1\"\n                                    echo \"hulk-container-agent buildJob notify --status=1 repoTag=${PKG_NAME}\"\n                                </command>\n                            </hudson.tasks.Shell>\n                        </buildSteps>\n                    </org.jenkinsci.plugins.postbuildscript.model.PostBuildStep>\n                </buildSteps>\n                <markBuildUnstable>false</markBuildUnstable>\n            </config>\n        </org.jenkinsci.plugins.postbuildscript.PostBuildScript>\n    </publishers>\n    <buildWrappers/>\n</project>','admin','admin',15,1,'local','',0,'','','',1,'2018-12-28 01:33:29','2019-01-23 00:49:43','/data/nvwa/packages','/data/nvwa/repos','/data/nvwa/builds','','');

/*!40000 ALTER TABLE `system` ENABLE KEYS */;
UNLOCK TABLES;


# Dump of table user
# ------------------------------------------------------------

DROP TABLE IF EXISTS `user`;

CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(64) NOT NULL DEFAULT '' COMMENT '用户名',
  `email` varchar(128) NOT NULL DEFAULT '' COMMENT '邮箱',
  `display_name` varchar(64) NOT NULL DEFAULT '' COMMENT '姓名',
  `password` varchar(50) NOT NULL DEFAULT '' COMMENT '密码',
  `avatar` varchar(100) NOT NULL DEFAULT '' COMMENT '头像，默认：nvwa logo',
  `role` varchar(10) NOT NULL DEFAULT '' COMMENT '角色，普通用户：10；管理员：20',
  `enabled` tinyint(3) unsigned NOT NULL DEFAULT '1',
  `ctime` datetime NOT NULL DEFAULT '0000-00-00 00:00:00',
  `utime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_username` (`username`),
  KEY `idx_name` (`username`,`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;

INSERT INTO `user` (`id`, `username`, `email`, `display_name`, `password`, `avatar`, `role`, `enabled`, `ctime`, `utime`)
VALUES
	(1,'nvwa-io','hiko.qiu@qq.com','nvwa-io','d9244dfbd9919633056d14b520754e13','','user',1,'2018-12-27 02:15:31','2018-12-27 02:15:31');

/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;



/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
