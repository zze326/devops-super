CREATE DATABASE `devops_super` default charset utf8mb4;
USE `devops_super`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ci_env
-- ----------------------------
DROP TABLE IF EXISTS `ci_env`;
CREATE TABLE `ci_env` (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `name` varchar(64) NOT NULL COMMENT '环境名称',
                          `image` varchar(256) NOT NULL COMMENT '镜像',
                          `secret_name` varchar(128) DEFAULT NULL COMMENT 'Kubernetes Secret 名称，拉取镜像使用',
                          `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ci_env
-- ----------------------------
BEGIN;
INSERT INTO `ci_env` (`id`, `name`, `image`, `secret_name`, `updated_at`) VALUES (1, 'Golang 1.19', 'test:12', 'azj', '2023-10-12 17:28:13');
COMMIT;

-- ----------------------------
-- Table structure for ci_pipeline
-- ----------------------------
DROP TABLE IF EXISTS `ci_pipeline`;
CREATE TABLE `ci_pipeline` (
                               `id` int(11) NOT NULL AUTO_INCREMENT,
                               `name` varchar(64) NOT NULL COMMENT '名称',
                               `kubernetes_config_id` int(11) NOT NULL COMMENT '关联的 Kubernetes Config id',
                               `config` json DEFAULT NULL COMMENT '配置',
                               `desc` varchar(256) DEFAULT NULL COMMENT '描述',
                               `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ci_pipeline
-- ----------------------------
BEGIN;
INSERT INTO `ci_pipeline` (`id`, `name`, `kubernetes_config_id`, `config`, `desc`, `updated_at`) VALUES (1, 'test', 1, '[{\"id\": 1, \"stages\": [{\"name\": \"拉取代码\", \"tasks\": [{\"type\": 1, \"gitPullData\": {\"branch\": \"master\", \"gitUrl\": \"http://192.168.1.195:8990/scm/ops/devops-platform-fe.git\", \"secretId\": 2}, \"shellExecData\": {}}]}, {\"name\": \"编译\", \"tasks\": [{\"type\": 2, \"gitPullData\": {}, \"shellExecData\": {\"content\": \"echo go build -o app\", \"workDir\": \"devops-platform-fe\"}}]}]}, {\"id\": 1, \"stages\": [{\"name\": \"上传镜像\", \"tasks\": [{\"type\": 2, \"gitPullData\": {}, \"shellExecData\": {\"content\": \"echo docker push\", \"workDir\": \"devops-platform-fe\"}}]}]}]', 'test', '2023-10-13 15:23:19');
COMMIT;

-- ----------------------------
-- Table structure for dept
-- ----------------------------
DROP TABLE IF EXISTS `dept`;
CREATE TABLE `dept` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `name` varchar(64) NOT NULL COMMENT '部门名称',
                        `rank` int(11) NOT NULL COMMENT '排序',
                        `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级部门 id',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of dept
-- ----------------------------
BEGIN;
INSERT INTO `dept` (`id`, `name`, `rank`, `parent_id`, `updated_at`) VALUES (1, '深圳总公司', 0, 0, '2023-09-22 13:43:26');
INSERT INTO `dept` (`id`, `name`, `rank`, `parent_id`, `updated_at`) VALUES (2, '研发部', 1, 1, '2023-09-22 13:52:25');
INSERT INTO `dept` (`id`, `name`, `rank`, `parent_id`, `updated_at`) VALUES (3, '开发部', 1, 2, '2023-09-22 13:52:36');
INSERT INTO `dept` (`id`, `name`, `rank`, `parent_id`, `updated_at`) VALUES (4, '运维部', 2, 2, '2023-09-22 13:52:41');
INSERT INTO `dept` (`id`, `name`, `rank`, `parent_id`, `updated_at`) VALUES (5, '运营部', 2, 1, '2023-09-22 13:52:57');
COMMIT;

-- ----------------------------
-- Table structure for host
-- ----------------------------
DROP TABLE IF EXISTS `host`;
CREATE TABLE `host` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `name` varchar(256) NOT NULL COMMENT '名称',
                        `host_addr` varchar(256) NOT NULL COMMENT '主机名或IP',
                        `port` bigint(20) NOT NULL COMMENT '端口',
                        `username` varchar(256) NOT NULL COMMENT '用户名',
                        `password` varchar(256) DEFAULT NULL COMMENT '密码',
                        `private_key` varchar(4096) DEFAULT NULL COMMENT '私钥',
                        `use_key` bit(1) DEFAULT NULL COMMENT '是否使用公钥连接',
                        `desc` varchar(256) DEFAULT NULL COMMENT '描述',
                        `save_session` bit(1) NOT NULL DEFAULT b'0' COMMENT '是否保存会话',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        `host_group_id` int(11) DEFAULT NULL COMMENT '主机组 id',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=8 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of host
-- ----------------------------
BEGIN;
INSERT INTO `host` (`id`, `name`, `host_addr`, `port`, `username`, `password`, `private_key`, `use_key`, `desc`, `save_session`, `updated_at`, `host_group_id`) VALUES (1, '测试机', '127.0.0.1', 22, 'zze', '128228', '-----BEGIN OPENSSH PRIVATE KEY-----\n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\n-----END OPENSSH PRIVATE KEY-----', b'0', 'test', b'0', '2023-09-22 17:12:41', NULL);
INSERT INTO `host` (`id`, `name`, `host_addr`, `port`, `username`, `password`, `private_key`, `use_key`, `desc`, `save_session`, `updated_at`, `host_group_id`) VALUES (2, '测试机', '127.0.0.1', 22, 'zze', '128228', '-----BEGIN OPENSSH PRIVATE KEY-----\n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\n-----END OPENSSH PRIVATE KEY-----', b'0', 'tesst', b'0', '2023-09-22 17:12:41', NULL);
INSERT INTO `host` (`id`, `name`, `host_addr`, `port`, `username`, `password`, `private_key`, `use_key`, `desc`, `save_session`, `updated_at`, `host_group_id`) VALUES (4, 'test2', '192.168.2.230', 22, 'root', '123456', '', b'0', '', b'1', '2023-10-10 17:57:03', 0);
INSERT INTO `host` (`id`, `name`, `host_addr`, `port`, `username`, `password`, `private_key`, `use_key`, `desc`, `save_session`, `updated_at`, `host_group_id`) VALUES (7, '测试新增', '192.168.3.33', 22, 'root', '123123', '', b'0', 'ss', b'0', '2023-10-10 17:52:33', 9);
COMMIT;

-- ----------------------------
-- Table structure for host_group
-- ----------------------------
DROP TABLE IF EXISTS `host_group`;
CREATE TABLE `host_group` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `name` varchar(64) NOT NULL COMMENT '主机组名称',
                              `rank` int(11) NOT NULL COMMENT '排序',
                              `parent_id` int(11) NOT NULL DEFAULT '0' COMMENT '上级主机组 id',
                              `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                              `role_ids` json DEFAULT NULL COMMENT '可访问的角色 id 列表',
                              `user_ids` json DEFAULT NULL COMMENT '可访问的用户 id 列表',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of host_group
-- ----------------------------
BEGIN;
INSERT INTO `host_group` (`id`, `name`, `rank`, `parent_id`, `updated_at`, `role_ids`, `user_ids`) VALUES (6, '阿里云', 1, 0, '2023-10-08 18:08:37', '[1]', '[1]');
INSERT INTO `host_group` (`id`, `name`, `rank`, `parent_id`, `updated_at`, `role_ids`, `user_ids`) VALUES (7, '腾讯云', 2, 0, '2023-09-27 13:47:26', '[2]', '[]');
INSERT INTO `host_group` (`id`, `name`, `rank`, `parent_id`, `updated_at`, `role_ids`, `user_ids`) VALUES (8, '华为云', 3, 0, '2023-09-22 18:33:46', NULL, NULL);
INSERT INTO `host_group` (`id`, `name`, `rank`, `parent_id`, `updated_at`, `role_ids`, `user_ids`) VALUES (9, '深圳', 1, 6, '2023-09-22 18:34:08', NULL, NULL);
INSERT INTO `host_group` (`id`, `name`, `rank`, `parent_id`, `updated_at`, `role_ids`, `user_ids`) VALUES (10, '张家口', 2, 6, '2023-10-08 15:39:53', '[2]', '[]');
COMMIT;

-- ----------------------------
-- Table structure for host_terminal_session
-- ----------------------------
DROP TABLE IF EXISTS `host_terminal_session`;
CREATE TABLE `host_terminal_session` (
                                         `id` int(11) NOT NULL AUTO_INCREMENT,
                                         `host_id` int(11) DEFAULT NULL COMMENT '主机 ID',
                                         `host_addr` varchar(256) DEFAULT NULL COMMENT '主机名或IP',
                                         `host_name` varchar(256) DEFAULT NULL COMMENT '主机名',
                                         `operator_id` int(11) DEFAULT NULL COMMENT '操作人 ID',
                                         `operator_name` varchar(256) DEFAULT NULL COMMENT '操作人用户名',
                                         `operator_real_name` varchar(256) DEFAULT NULL COMMENT '操作人真实姓名',
                                         `start_time` datetime DEFAULT NULL COMMENT '会话开始时间',
                                         `filepath` varchar(256) DEFAULT NULL COMMENT '会话文件路径',
                                         `updated_at` datetime DEFAULT NULL,
                                         PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=38 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of host_terminal_session
-- ----------------------------
BEGIN;
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (25, 4, '192.168.2.230', 'test2', 1, 'admin', '管理员', '2023-09-26 17:15:56', 'host-sessions/4/1695719755944915.sessionb', '2023-09-26 17:16:01');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (26, 4, '192.168.2.230', 'test2', 1, 'admin', '管理员', '2023-09-27 09:50:25', 'host-sessions/4/1695779424598082.sessionb', '2023-09-27 09:50:33');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (27, 4, '192.168.2.230', 'test2', 1, 'admin', '管理员', '2023-09-27 10:45:57', 'host-sessions/4/1695782756764543.sessionb', '2023-09-27 10:46:07');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (28, 4, '192.168.2.230', 'test2', 2, 'test', '测试用户', '2023-10-08 14:44:05', 'host-sessions/4/1696747444964082.sessionb', '2023-10-08 14:44:55');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (29, 4, '192.168.2.230', 'test2', 2, 'test', '测试用户', '2023-10-08 14:46:05', 'host-sessions/4/1696747564551481.sessionb', '2023-10-08 14:47:12');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (30, 4, '192.168.2.230', 'test2', 2, 'test', '测试用户', '2023-10-08 14:53:58', 'host-sessions/4/1696748037946844.sessionb', '2023-10-08 14:54:21');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (31, 4, '192.168.2.230', 'test2', 2, 'test', '测试用户', '2023-10-08 15:52:24', 'host-sessions/4/1696751543696470.sessionb', '2023-10-08 15:52:28');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (32, 4, '192.168.2.230', 'test2', 2, 'test', '测试用户', '2023-10-08 15:58:47', 'host-sessions/4/1696751926688589.sessionb', '2023-10-08 15:59:17');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (33, 4, '192.168.2.230', 'test2', 1, 'admin', '管理员', '2023-10-08 17:14:17', 'host-sessions/4/1696756456610441.sessionb', '2023-10-08 17:15:07');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (34, 4, '192.168.2.230', 'test2', 1, 'admin', '管理员', '2023-10-08 17:44:34', 'host-sessions/4/1696758273566098.sessionb', '2023-10-08 17:46:14');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (35, 4, '192.168.2.230', 'test2', 1, 'admin', '管理员', '2023-10-08 17:54:09', 'host-sessions/4/1696758849207336.sessionb', '2023-10-08 17:55:19');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (36, 4, '192.168.2.230', 'test2', 2, 'test', '测试用户', '2023-10-09 14:40:47', 'host-sessions/4/1696833646736335.sessionb', '2023-10-09 14:43:57');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (37, 4, '192.168.2.230', 'test2', 1, 'admin', '管理员', '2023-10-10 17:48:44', 'host-sessions/4/1696931323959342.sessionb', '2023-10-10 17:49:04');
COMMIT;

-- ----------------------------
-- Table structure for permission
-- ----------------------------
DROP TABLE IF EXISTS `permission`;
CREATE TABLE `permission` (
                              `id` int(11) NOT NULL AUTO_INCREMENT,
                              `title` varchar(64) NOT NULL COMMENT '标题',
                              `name` varchar(64) NOT NULL COMMENT '路由名称',
                              `type` tinyint(4) NOT NULL COMMENT '类型:1-目录,2-菜单,3-功能',
                              `f_route` varchar(64) DEFAULT NULL COMMENT '前端路由路径',
                              `b_routes` json DEFAULT NULL COMMENT '后端路由路径',
                              `redirect` varchar(64) DEFAULT NULL COMMENT '重定向路径',
                              `icon` varchar(32) DEFAULT NULL COMMENT '图标',
                              `rank` int(11) DEFAULT NULL COMMENT '排序',
                              `show_link` bit(1) NOT NULL COMMENT '是否在菜单中展示',
                              `show_parent` bit(1) NOT NULL COMMENT '是否展示父级菜单',
                              `keep_alive` bit(1) NOT NULL COMMENT '页面缓存',
                              `parent_id` int(11) NOT NULL COMMENT '父级权限 id',
                              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=68 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of permission
-- ----------------------------
BEGIN;
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (1, '系统管理', 'system-manage', 1, '/system', '[]', '/system/user', 'ep:setting', 3, b'1', b'0', b'0', 0);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (2, '权限管理', 'permission-manage', 2, '/system/permission', '[]', '', 'fa-solid:allergies', 3, b'1', b'1', b'1', 1);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (3, '用户管理', 'user-manage', 2, '/system/user', '[]', '', 'fa:address-card', 1, b'1', b'1', b'1', 1);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (4, '新增用户', 'user-add', 3, '/test4', '[\"post:/user\"]', '/test4', '', 2, b'1', b'0', b'0', 3);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (5, '角色管理', 'role-manage', 2, '/system/role', '[]', '', 'ep:avatar', 2, b'1', b'1', b'1', 1);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (6, '更新用户', 'user-upt', 3, '/test6', '[\"put:/user/:id\"]', '/test6', '', 3, b'1', b'0', b'0', 3);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (8, '更新用户密码', 'user-upt-password', 3, '', '[\"patch:/user/:id/password\"]', '', '', 5, b'0', b'0', b'0', 3);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (9, '启用、禁用用户', 'user-upt-enable', 3, '', '[\"patch:/user/:id/enabled\"]', '', '', 6, b'0', b'0', b'0', 3);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (10, '删除用户', 'user-del', 3, '', '[\"delete:/user/:id\"]', '', '', 4, b'0', b'0', b'0', 3);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (11, '查询用户', 'user-read', 3, '', '[\"get:/user/page-list\", \"get:/dept/list\"]', '', '', 1, b'0', b'0', b'0', 3);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (14, '查询权限', 'permission-read', 3, '', '[\"get:/permission/list\"]', '', '', 1, b'0', b'0', b'0', 2);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (15, '新增权限', 'permission-add', 3, '', '[\"post:/permission\"]', '', '', 2, b'0', b'0', b'0', 2);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (16, '更新权限', 'permission-upt', 3, '', '[\"put:/permission/:id\"]', '', '', 3, b'0', b'0', b'0', 2);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (17, '删除权限', 'permission-del', 3, '', '[\"delete:/permission/:id\"]', '', '', 4, b'0', b'0', b'0', 2);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (18, '更新权限是否展示在菜单', 'permission-upt-show-link', 3, '', '[\"patch:/permission/:id/show-link\"]', '', '', 5, b'0', b'0', b'0', 2);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (20, '查询角色', 'role-read', 3, '', '[\"get:/role/page-list\", \"get:/permission/list\"]', '', '', 1, b'0', b'0', b'0', 5);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (21, '新增角色', 'role-add', 3, '', '[\"post:/role\"]', '', '', 2, b'0', b'0', b'0', 5);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (22, '更新角色', 'role-upt', 3, '', '[\"put:/role:id\"]', '', '', 3, b'0', b'0', b'0', 5);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (23, '删除角色', 'role-del', 3, '', '[\"delete:/role/:id\"]', '', '', 4, b'0', b'0', b'0', 5);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (24, '保存角色权限信息', 'role-upt-permission', 3, '', '[\"patch:/role/:id/permission\"]', '', '', 5, b'0', b'0', b'0', 5);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (25, '系统必需', 'system-required', 3, '', '[\"get:/permission/route-list\"]', '', '', 0, b'0', b'0', b'0', 0);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (27, '部门管理', 'dept-manage', 2, '/system/dept', '[]', '', 'fa:group', 4, b'1', b'1', b'1', 1);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (28, '查询部门', 'dept-read', 3, '', '[\"get:/dept/list\"]', '', '', 1, b'0', b'0', b'0', 27);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (29, '新增部门', 'dept-add', 3, '', '[\"post:/dept\"]', '', '', 2, b'0', b'0', b'0', 27);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (30, '更新部门', 'dept-upt', 3, '', '[\"put:/dept/:id\"]', '', '', 3, b'0', b'0', b'0', 27);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (31, '删除部门', 'dept-del', 3, '', '[\"delete:/dept/:id\"]', '', '', 4, b'0', b'0', b'0', 27);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (32, '资源管理', 'resource-manage', 1, '/resource', '[]', '/resource/host', 'ep:box', 2, b'1', b'0', b'0', 0);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (33, '主机组管理', 'host-group-manage', 2, '/resource/host-group', '[]', '', 'fa:server', 2, b'1', b'1', b'0', 32);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (34, '主机管理', 'host-manage', 2, '/resource/host', '[]', '', 'fa:desktop', 1, b'1', b'1', b'0', 32);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (35, '查询主机组', 'host-group-read', 3, '', '[\"get:/host-group/list\"]', '', '', 1, b'0', b'0', b'0', 33);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (36, '新增主机组', 'host-group-add', 3, '', '[\"post:/host-group\"]', '', '', 2, b'0', b'0', b'0', 33);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (37, '更新主机组', 'host-group-upt', 3, '', '[\"put:/host-group/:id\"]', '', '', 3, b'0', b'0', b'0', 33);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (38, '删除主机组', 'host-group-del', 3, '', '[\"delete:/host-group/:id\"]', '', '', 4, b'0', b'0', b'0', 33);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (39, '查询主机', 'host-read', 3, '', '[\"get:/host/page-list\", \"get:/host-group/partial-list\"]', '', '', 1, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (40, '新增主机', 'host-add', 3, '', '[\"post:/host\"]', '', '', 2, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (41, '更新主机', 'host-upt', 3, '', '[\"put:/host/:id\"]', '', '', 3, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (42, '删除主机', 'host-del', 3, '', '[\"delete:/host/:id\"]', '', '', 4, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (43, '查询终端会话记录', 'host-terminal-session-read', 3, '', '[\"get:/host-terminal-session/page-list\"]', '', '', 10, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (44, '连接终端', 'host-terminal-connect', 3, '', '[\"get:/host/:id/ssh-ok\", \"get:/host/:id/terminal\", \"get:/host/authorized-list\", \"get:/host-group/list\", \"get:/host/:id/one\"]', '', '', 5, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (45, '连接终端文件管理器', 'host-terminal-sftp-read', 3, '', '[\"get:/host/:id/sftp-file-manager\"]', '', '', 6, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (46, '文件管理器上传文件', 'host-terminal-sftp-upload', 3, '', '[]', '', '', 7, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (47, '文件管理器下载文件', 'host-terminal-sftp-download', 3, '', '[\"get:/host/:id/download-file\"]', '', '', 8, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (48, '文件管理器删除文件', 'host-terminal-sftp-del', 3, '', '[]', '', '', 9, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (49, '回放终端会话记录', 'host-terminal-session-replay', 3, '', '[\"get:/host-terminal-session/:id/check-file\", \"get:/host-terminal-session/:id/replay\"]', '', '', 11, b'0', b'0', b'0', 34);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (50, '流水线', 'ci-pipeline-manage', 2, '/ci/ci-pipeline', '[]', '', 'fa-solid:grip-lines-vertical', 1, b'1', b'1', b'1', 56);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (51, '秘钥管理', 'secret-manage', 2, '/resource/secret', '[]', '', 'fa-solid:key', 3, b'1', b'0', b'0', 32);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (52, '查询秘钥', 'secret-read', 3, '', '[\"get:/secret/page-list\"]', '', '', 1, b'0', b'0', b'0', 51);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (53, '新增秘钥', 'secret-add', 3, '', '[\"post:/secret\"]', '', '', 2, b'0', b'0', b'0', 51);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (54, '更新秘钥', 'secret-upt', 3, '', '[\"put:/secret/:id\"]', '', '', 3, b'0', b'0', b'0', 51);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (55, '删除秘钥', 'secret-del', 3, '', '[\"delete:/secret/:id\"]', '', '', 4, b'0', b'0', b'0', 51);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (56, '持续集成', 'ci-manage', 1, '/ci', '[]', '/ci/ci-pipeline', 'fa-solid:smoking', 1, b'1', b'0', b'0', 0);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (57, '构建环境', 'ci-env-manage', 2, '/ci/ci-env', '[]', '', 'fa-solid:boxes', 1, b'1', b'1', b'1', 56);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (58, '查询构建环境', 'ci-env-read', 3, '', '[\"get:/ci-env/page-list\"]', '', '', 1, b'0', b'0', b'0', 57);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (59, '新增构建环境', 'ci-env-add', 3, '', '[\"post:/ci-env\"]', '', '', 2, b'0', b'0', b'0', 57);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (60, '更新构建环境', 'ci-env-upt', 3, '', '[\"put:/ci-env/:id\"]', '', '', 3, b'0', b'0', b'0', 57);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (61, '删除构建环境', 'ci-env-del', 3, '', '[\"delete:/ci-env/:id\"]', '', '', 4, b'0', b'0', b'0', 57);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (62, '查询流水线', 'ci-pipeline-read', 3, '', '[\"get:/ci-pipeline/page-list\", \"get:/kubernetes-config/partial-list\"]', '', '', 1, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (63, '新增流水线', 'ci-pipeline-add', 3, '', '[\"post:/ci-pipeline\"]', '', '', 2, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (64, '更新流水线', 'ci-pipeline-upt', 3, '', '[\"put:/ci-pipeline/:id\"]', '', '', 3, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (65, '删除流水线', 'ci-pipeline-del', 3, '', '[\"delete:/ci-pipeline/:id\"]', '', '', 4, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (66, '编排流水线', 'ci-pipeline-arrange', 3, '', '[\"get:/ci-pipeline/:id/config\", \"patch:/ci-pipeline/:id/config\", \"get:/secret/list\"]', '', '', 5, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (67, '运行流水线', 'ci-pipeline-run', 3, '', '[\"post:/ci-pipeline/:id/run\"]', '', '', 6, b'0', b'0', b'0', 50);
COMMIT;

-- ----------------------------
-- Table structure for role
-- ----------------------------
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `name` varchar(128) NOT NULL COMMENT '角色名称',
                        `code` varchar(30) NOT NULL COMMENT '角色代码',
                        `permission` json DEFAULT NULL COMMENT '关联权限',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of role
-- ----------------------------
BEGIN;
INSERT INTO `role` (`id`, `name`, `code`, `permission`, `updated_at`) VALUES (1, '管理员', 'admin', '[25, 56, 57, 32, 34, 39, 40, 41, 42, 44, 45, 46, 47, 48, 43, 49, 33, 35, 36, 37, 38, 51, 52, 53, 54, 55, 50, 1, 3, 11, 4, 6, 10, 8, 9, 5, 20, 21, 22, 23, 24, 2, 14, 15, 16, 17, 18, 27, 28, 29, 30, 31]', '2023-10-09 15:55:05');
INSERT INTO `role` (`id`, `name`, `code`, `permission`, `updated_at`) VALUES (2, '测试', 'test', '[25, 56, 50, 62, 63, 64, 65, 57, 58, 59, 60, 61]', '2023-10-09 18:18:59');
COMMIT;

-- ----------------------------
-- Table structure for secret
-- ----------------------------
DROP TABLE IF EXISTS `secret`;
CREATE TABLE `secret` (
                          `id` int(11) NOT NULL AUTO_INCREMENT,
                          `name` varchar(128) NOT NULL COMMENT '名称',
                          `type` tinyint(4) NOT NULL COMMENT '类型:1-git认证,2-Kubernetes config',
                          `content` json NOT NULL COMMENT '认证配置内容',
                          `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of secret
-- ----------------------------
BEGIN;
INSERT INTO `secret` (`id`, `name`, `type`, `content`, `updated_at`) VALUES (1, '线下集群', 2, '{\"text\": \"apiVersion: v1\\nclusters:\\n- cluster:\\n    certificate-authority-data: LS0tLS1CRUdJTiBDRVJUSxxMnM0Y3pFUE1BMEdBMVVFQ3hNR1UzbHpkR1Z0TVJNd0VRWURWUVFERXdwcmRXSmxjbTVsZEdWek1JSUJJakFOCkJna3Foa2lHOXcwQkFRRUZBQU9DQVE4QU1JSUJDZ0tDQVFFQXhNQTZWc3VNc1lXRDl2dVM5RmlKaU5CZGo3WEYKbm0rT0tlZWllNFc4aStCaFBYUHorWlBTcUptSkNXNHlPTjQ4MW02ZHRJU1hLRXhGek9pYWZGSzd3OTd2Y1pDeApQTUd5eVpqanQ0Q3lrbDg1V2tRZEhGa1dvcjBOMkVyWTZydGMxK3huSWJIUml2ZXBteVRydDJ5Y1hmWHB4cmxOCnlVcEQ3U2ErZmt1Nlp0a3RqNEZweVVjN0tvbCtodk13dmtNRExPdEtpWkt5WU95YzZac0FndFFwMGdGeXNaNEYKcTZxQTd3aUloWjdTYWRNTlJYbERmZmNaUW9sOFVNQTh3N01ub0VVT3d5ZUZnc1BBaE4vTkhnY29pSnd4cnFmcgp1K2U1VE9RWVJFRk5CUTVyTFFSL3FKYUt0MUVkV0hrZWkvRVJWbk5JZ3JFbVNzbWhtS3VDK01tUXd3SURBUUFCCm8wSXdRREFPQmdOVkhROEJBZjhFQkFNQ0FRWXdEd1lEVlIwVEFRSC9CQVV3QXdFQi96QWRCZ05WSFE0RUZnUVUKQVRQS0hnQXNMWTlKNXpqNzhXYm9lNmVjWUZvd0RRWUpLb1pJaHZjTkFRRUxCUUFEZ2dFQkFFZ3U3OWxsWnRLNgpWVTJpUzVTcWwreGFnN3J6NTJJWmhid2ZFc2N3K3doclBqTWxsdGQ2UTh6YUE2KzZRdFRLVWo0bWIwbVlONWhJCm9mdWhna2NaNnZ3NkxRVkNVZ25kZmFULzcrVEtFQXJkV2o2emRYanFveWRnenR5SDEyQXhtM05rcXI3N2E1dCsKUWMwV2ZULzltNHNMdXZCa0tNdWhqMkg3VUZzMXN2QVdkUmltRk1PeEhJUjZVQ0RMMkZ5M0htM2VPS1F1OXFoOAo0QnhKMTczVnZmZmpZUHp2UlhuR0NrWm52Mkt5MG5hWFRkRHoyK0FtU2NiNGs1ZUR3MUJ6M1J2czdxRitOTktqCndRaTJmTFpOREZxTGxOWXc3eDNxbWxJWGxNY01jdEUyUlM0SjRTbFVacy9RbHk2VWE4NkwrUFdwcUpLN1BsYzQKVSs1T00xWCtEbWs9Ci0tLS0tRU5EIENFUlRJRklDQVRFLS0tLS0K\\n    server: https://192.168.2.2:7443\\n  name: azj-cluster\\ncontexts:\\n- context:\\n    cluster: azj-cluster\\n    namespace: devops\\n    user: admin\\n  name: offline-cluster\\ncurrent-context: offline-cluster\\nkind: Config\\npreferences: {}\\nusers:\\n- name: admin\\n  user:\\n    client-certificate-data: LS0tLS1CRUdJTiBDRVJUSUZJQ0FURS0tLS0tCk1JSUQxekNDQXIrZ0F3SUJBZ0lVUW4wWXp6N2FRZkZxxkx0WXRmZVNGWUs0K3U0ZzFRNDRpbFNDcFJ1U3cvR0phenZ3Z0xIeGNtTWZEK004VmtaTnF1cUJUdm82CnZ3SURBUUFCbzM4d2ZUQU9CZ05WSFE4QkFmOEVCQU1DQmFBd0hRWURWUjBsQkJZd0ZBWUlLd1lCQlFVSEF3RUcKQ0NzR0FRVUZCd01DTUF3R0ExVWRFd0VCL3dRQ01BQXdIUVlEVlIwT0JCWUVGT3pwVStxaHUvSnVnUDZDZkFDTApnMW5VMVRXeE1COEdBMVVkSXdRWU1CYUFGQUV6eWg0QUxDMlBTZWM0Ky9GbTZIdW5uR0JhTUEwR0NTcUdTSWIzCkRRRUJDd1VBQTRJQkFRQ0lMRDJpNmNUQWp6Yk5pZW1hNDZhOHd3cWhacHRZdFNDMmdLTHJZRDJaQ2daYVVjU3oKdkg1ayt2SmtjQ29uWW91WkFpTVAyMGxDZElhaWhtYlREdnprcDBNaGZibUppcUlKR2dlcW0vVUsxRWRmNzUvbAo2SGx2MWsrMDdjZXVrcC9Hc2tRMkowRjFBR0ZJdFRncHFJQk5SNWJxR0FNeGcvS0FCdEt4ZUVPbm5EYU9pTGI0CkxiQ3RzWXdVQ09Zb1Z2Skt6V1F4RFVCcTRvcmxhNU81eHVPVWVOSWlHcTZiek1iV3lpbzNQcVJ2Y0JzMU1JaHYKdXBHZlBEQkErdFdrM0hDL0xETkEzdTlGNWpJa3c3ZkJGenlER3d3QUVxMmhLUUFncmlMY05GY2YxbzhPTHJZaQpzTU9YcUFucW1sbkxkWkJVc2ppdnRxZ3M1N2trYnR2R3ZOY3IKLS0tLS1FTkQgQ0VSVElGSUNBVEUtLS0tLQo=\\n    client-key-data: LS0tLS1CRUdJTiBSU0EgUFJJVkFURSBLRVktLS0tLQpNSUlFcFFJQkFBS0NBUUVBMkhiOHI5TEV2UUJoY0YxVHdUaTFsNzdvb3M3RTFVeWt1MkZic0VxaWJJSVhYNnczClBYUmdwWStvWTB5SUJ5ODdIV2FaY1luM1liQmp4aksvbmlPVkoyMHNTbDBuUFl2L2I0MDRrYnMvUktzL1BZRXAKxxYzZPS2VLajdncno0Qm5zVUd1NEpIaHJYSzBHMjk4NWhMeXcrSFlIbkkKYkpKc2laNnlyS3dNdklBeDZLQkdsR25pT1JlY1hXaXFrK1hHcEx4NldkRkpWOXhSTXoxVWhrblQ4TnVFeW1OSApHTlMwaE5LaDBTMHdnenh4YTB5Nmhoc2YxNW4zODBlUFkvRkI3WUd1WHIvb1lreFNFQ21LbGF4M3Ribk94dkF2CnJvN2M3Um5rbXFXTUZGTTlRV3dVVEpsckYvY0JKTkhPTkxBRm5VcWZabG9TS2NzY0xwWjVqUGZUWGdkWVRLWnkKcCtRxxApqa0VKWnpZcGM1QTBtU200b0hOMitSdFhSQ21vNlNvVkg3NjhnUkxuN1RZQ25FS3pvd09MdC81VFhCdnh6ZnpIClQyTlpPWlY2WG1EV09qbHBsTXQ4UVlwcFVXenBsZnp3UDV5bjh6RUNnWUVBMmVTVi91c0pibWVIQ3hkWkFNanIKOUZVUjh5OERmOStONE1xOFRYS1ZCVURjOUxOT0w0UVBndGYzTVVYNjJLQTQxb2xMaHVIS3pwMEVoZWVjMnxxFHcU5vTVF0MGZLMzJMUERXT1NHSzc0RWhkS0hGV1pqOTRoUTY4SzZJb1RWVmxDMGhuVApteTcvZW96RkkvYVFuVjlNQ3FSdlRrc1Myb2Z4OVA5a1dVWVZWQVlnYUNaeGR1RENyNm1VZmlRPQotLS0tLUVORCBSU0EgUFJJVkFURSBLRVktLS0tLQo=\"}', '2023-10-13 15:48:39');
INSERT INTO `secret` (`id`, `name`, `type`, `content`, `updated_at`) VALUES (2, '线下 BitBucket', 1, '{\"password\": \"123456\", \"username\": \"zhangzhongen\"}', '2023-10-13 15:48:15');
COMMIT;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
                        `id` int(11) NOT NULL AUTO_INCREMENT,
                        `username` varchar(128) NOT NULL COMMENT '用户名',
                        `password` varchar(128) NOT NULL COMMENT '密码',
                        `phone` varchar(128) DEFAULT NULL COMMENT '手机号码',
                        `email` varchar(128) DEFAULT NULL COMMENT '邮箱',
                        `real_name` varchar(128) DEFAULT NULL COMMENT '真实姓名',
                        `enabled` bit(1) NOT NULL DEFAULT b'1' COMMENT '是否启用状态',
                        `role_ids` json DEFAULT NULL COMMENT '角色 id',
                        `dept_id` int(11) NOT NULL COMMENT '所属部门 id',
                        `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                        PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- ----------------------------
-- Records of user
-- ----------------------------
BEGIN;
INSERT INTO `user` (`id`, `username`, `password`, `phone`, `email`, `real_name`, `enabled`, `role_ids`, `dept_id`, `updated_at`) VALUES (1, 'admin', '$2a$10$4R/ujw20O63gNwBTS0vJmOfAUukGT5pCMll0gsqy6IALPHZC7SDv2', '16666666666', '632404164@qq.com', '管理员', b'1', '[1]', 1, '2023-09-22 16:45:33');
INSERT INTO `user` (`id`, `username`, `password`, `phone`, `email`, `real_name`, `enabled`, `role_ids`, `dept_id`, `updated_at`) VALUES (2, 'test', '$2a$10$aEX83iCGh/JrxiTImN0PE.0bK/dLE1lFFeZ4ssHdK4/rrXCqMgRHe', '16666666666', '632404164@qq.com', '测试用户', b'1', '[2]', 4, '2023-10-08 15:53:38');
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
