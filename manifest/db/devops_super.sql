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
                          `persistence_config` json DEFAULT NULL COMMENT '持久化配置',
                          `is_kaniko` bit(1) NOT NULL COMMENT '是否是 kaniko 客户端',
                          `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                          PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ci_env
-- ----------------------------
BEGIN;
INSERT INTO `ci_env` (`id`, `name`, `image`, `secret_name`, `persistence_config`, `is_kaniko`, `updated_at`) VALUES (1, 'Golang 1.19', 'registry.cn-shenzhen.aliyuncs.com/zze/devops-super-ci-client:202312121641', '', '[{\"pvcName\": \"devops-super-ci\", \"subPath\": \"golang\", \"mountPath\": \"/root/go\"}]', b'0', '2023-12-12 16:43:24');
INSERT INTO `ci_env` (`id`, `name`, `image`, `secret_name`, `persistence_config`, `is_kaniko`, `updated_at`) VALUES (2, 'Kaniko', 'registry.cn-shenzhen.aliyuncs.com/zze/gcriokaniko-executor:v1.19.0', '', '[{\"pvcName\": \"devops-super-ci\", \"subPath\": \"kaniko-cache\", \"mountPath\": \"/cache\"}]', b'1', '2023-12-14 16:52:07');
COMMIT;
-- ----------------------------
-- Table structure for ci_pipeline
-- ----------------------------
DROP TABLE IF EXISTS `ci_pipeline`;
CREATE TABLE `ci_pipeline` (
                               `id` int(11) NOT NULL AUTO_INCREMENT,
                               `name` varchar(64) NOT NULL COMMENT '名称',
                               `kubernetes_config_id` int(11) NOT NULL COMMENT '关联的 Kubernetes Config id',
                               `kubernetes_namespace` varchar(255) DEFAULT NULL COMMENT 'Pod 所在命名空间',
                               `parameterize` bit(1) NOT NULL COMMENT '是否是参数化构建',
                               `params` json DEFAULT NULL COMMENT '构建参数',
                               `config` json DEFAULT NULL COMMENT '配置',
                               `desc` varchar(256) DEFAULT NULL COMMENT '描述',
                               `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                               PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ci_pipeline
-- ----------------------------
BEGIN;
INSERT INTO `ci_pipeline` (`id`, `name`, `kubernetes_config_id`, `kubernetes_namespace`, `parameterize`, `params`, `config`, `desc`, `updated_at`) VALUES (1, 'test', 1, 'default', b'1', '[{\"name\": \"branch\", \"type\": 1, \"gitUrl\": \"http://gitlab.internal.azj/devops/devops-platform.git\", \"display\": \"分支\", \"secretId\": 3, \"gitPullData\": {}, \"shellExecData\": {}}]', '[{\"id\": 1, \"stages\": [{\"name\": \"拉取代码\", \"tasks\": [{\"type\": 1, \"gitPullData\": {\"branch\": \"{{ branch }}\", \"gitUrl\": \"http://gitlab.internal.azj/devops/devops-platform.git\", \"secretId\": 3}, \"shellExecData\": {}}]}, {\"name\": \"编译\", \"tasks\": [{\"type\": 2, \"gitPullData\": {}, \"shellExecData\": {\"content\": \"go env -w \'GOPROXY=https://goproxy.cn,direct\' && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app\", \"workDir\": \"devops-platform/microservice/app/api\"}}]}]}]', 'test', '2023-10-27 13:37:57');
INSERT INTO `ci_pipeline` (`id`, `name`, `kubernetes_config_id`, `kubernetes_namespace`, `parameterize`, `params`, `config`, `desc`, `updated_at`) VALUES (2, 'test2', 1, 'default', b'1', '[{\"name\": \"branch\", \"type\": 1, \"gitUrl\": \"http://gitlab.internal.azj/devops/devops-platform.git\", \"display\": \"分支\", \"secretId\": 3, \"gitPullData\": {}, \"shellExecData\": {}}]', '[{\"id\": 1, \"stages\": [{\"name\": \"拉取代码\", \"tasks\": [{\"type\": 1, \"gitPullData\": {\"branch\": \"{{ branch }}\", \"gitUrl\": \"http://gitlab.internal.azj/devops/devops-platform.git\", \"secretId\": 3}, \"shellExecData\": {}}]}, {\"name\": \"编译\", \"tasks\": [{\"type\": 2, \"gitPullData\": {}, \"shellExecData\": {\"content\": \"go env -w \'GOPROXY=https://goproxy.cn,direct\' && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app\", \"workDir\": \"devops-platform/microservice/app/api\"}}]}]}]', 'test', '2023-10-27 15:03:54');
COMMIT;

-- ----------------------------
-- Table structure for ci_pipeline_run
-- ----------------------------
DROP TABLE IF EXISTS `ci_pipeline_run`;
CREATE TABLE `ci_pipeline_run` (
                                   `id` int(11) NOT NULL AUTO_INCREMENT,
                                   `pipeline_id` int(11) NOT NULL COMMENT '流水线 id',
                                   `pod_name` varchar(128) NOT NULL COMMENT 'Pod 名称',
                                   `namespace` varchar(128) NOT NULL COMMENT '名称空间',
                                   `status` tinyint(4) NOT NULL COMMENT '状态:0-运行中,1:成功,2:失败,3:取消',
                                   `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
                                   `created_at` datetime NOT NULL COMMENT '创建时间',
                                   PRIMARY KEY (`id`),
                                   KEY `idx_status` (`status`)
) ENGINE=InnoDB AUTO_INCREMENT=135 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of ci_pipeline_run
-- ----------------------------
BEGIN;
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (1, 1, 'ci-pod-1-20231016170017', 'default', 1, '2023-10-16 17:00:44', '2023-10-16 17:00:18');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (2, 1, 'ci-pod-1-20231016170107', 'default', 2, '2023-10-16 17:01:34', '2023-10-16 17:01:08');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (3, 1, 'ci-pod-1-20231016171232', 'default', 2, '2023-10-16 17:12:58', '2023-10-16 17:12:33');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (4, 1, 'ci-pod-1-20231016171329', 'default', 1, '2023-10-16 17:13:56', '2023-10-16 17:13:29');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (5, 1, 'ci-test-1-20231017111948', 'default', 1, '2023-10-17 11:20:16', '2023-10-17 11:19:48');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (6, 1, 'ci-test-1-20231017120158', 'default', 1, '2023-10-17 12:02:26', '2023-10-17 12:01:59');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (7, 1, 'ci-test-1-20231017120242', 'default', 1, '2023-10-17 12:03:09', '2023-10-17 12:02:42');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (8, 1, 'ci-test-1-20231017123201', 'default', 1, '2023-10-17 12:32:28', '2023-10-17 12:32:02');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (9, 1, 'ci-test-1-20231017123358', 'default', 1, '2023-10-17 12:34:25', '2023-10-17 12:33:58');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (10, 1, 'ci-test-1-20231017123432', 'default', 1, '2023-10-17 12:35:00', '2023-10-17 12:34:33');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (11, 1, 'ci-test-1-20231017123651', 'default', 1, '2023-10-17 12:37:18', '2023-10-17 12:36:51');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (12, 1, 'ci-test-1-20231017124016', 'default', 1, '2023-10-17 12:40:44', '2023-10-17 12:40:16');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (13, 1, 'ci-test-1-20231017150455', 'default', 1, '2023-10-17 15:05:22', '2023-10-17 15:04:55');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (14, 1, 'ci-test-1-20231017150817', 'default', 1, '2023-10-17 15:08:45', '2023-10-17 15:08:18');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (15, 1, 'ci-test-1-20231017151155', 'default', 1, '2023-10-17 15:12:23', '2023-10-17 15:11:56');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (16, 1, 'ci-test-1-20231017151356', 'default', 1, '2023-10-17 15:14:28', '2023-10-17 15:13:57');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (17, 1, 'ci-test-1-20231017151446', 'default', 1, '2023-10-17 15:15:14', '2023-10-17 15:14:47');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (18, 1, 'ci-test-1-20231017154433', 'default', 1, '2023-10-17 15:45:04', '2023-10-17 15:44:33');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (19, 1, 'ci-test-1-20231017154537', 'default', 1, '2023-10-17 15:46:05', '2023-10-17 15:45:37');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (20, 1, 'ci-test-1-20231017161536', 'default', 1, '2023-10-17 16:16:05', '2023-10-17 16:15:36');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (21, 1, 'ci-test-1-20231017163049', 'default', 1, '2023-10-17 16:31:15', '2023-10-17 16:30:49');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (22, 1, 'ci-test-1-20231017163604', 'default', 1, '2023-10-17 16:36:35', '2023-10-17 16:36:04');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (23, 1, 'ci-test-1-20231017163712', 'default', 1, '2023-10-17 16:37:42', '2023-10-17 16:37:12');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (24, 1, 'ci-test-1-20231017163822', 'default', 1, '2023-10-17 16:38:53', '2023-10-17 16:38:23');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (25, 1, 'ci-test-1-20231017163939', 'default', 1, '2023-10-17 16:40:06', '2023-10-17 16:39:40');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (26, 1, 'ci-test-1-20231017164112', 'default', 1, '2023-10-17 16:41:42', '2023-10-17 16:41:13');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (27, 1, 'ci-test-1-20231017164341', 'default', 1, '2023-10-17 16:44:12', '2023-10-17 16:43:41');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (28, 1, 'ci-test-1-20231017164541', 'default', 1, '2023-10-17 16:46:08', '2023-10-17 16:45:41');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (29, 1, 'ci-test-1-20231017165755', 'default', 1, '2023-10-17 16:58:28', '2023-10-17 16:57:56');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (30, 1, 'ci-test-1-20231017170034', 'default', 1, '2023-10-17 17:01:01', '2023-10-17 17:00:34');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (31, 1, 'ci-test-1-20231017170246', 'default', 1, '2023-10-17 17:03:36', '2023-10-17 17:02:46');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (32, 1, 'ci-test-1-20231017170808', 'default', 1, '2023-10-17 17:08:35', '2023-10-17 17:08:09');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (33, 1, 'ci-test-1-20231017171023', 'default', 1, '2023-10-17 17:10:53', '2023-10-17 17:10:23');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (34, 1, 'ci-test-1-20231017171110', 'default', 1, '2023-10-17 17:11:32', '2023-10-17 17:11:10');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (35, 1, 'ci-test-1-20231017171321', 'default', 2, '2023-10-17 17:13:37', '2023-10-17 17:13:22');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (36, 1, 'ci-test-1-20231017171646', 'default', 2, '2023-10-17 17:16:59', '2023-10-17 17:16:47');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (37, 1, 'ci-test-1-20231017171927', 'default', 2, '2023-10-17 17:19:40', '2023-10-17 17:19:28');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (38, 1, 'ci-test-1-20231017172137', 'default', 2, '2023-10-17 17:21:51', '2023-10-17 17:21:37');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (39, 1, 'ci-test-1-20231017175236', 'default', 2, '2023-10-17 17:52:50', '2023-10-17 17:52:37');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (40, 1, 'ci-test-1-20231017175420', 'default', 2, '2023-10-17 17:54:33', '2023-10-17 17:54:21');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (41, 1, 'ci-test-1-20231017180013', 'default', 2, '2023-10-17 18:00:30', '2023-10-17 18:00:14');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (42, 1, 'ci-test-1-20231017180627', 'default', 2, '2023-10-17 18:06:42', '2023-10-17 18:06:28');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (43, 1, 'ci-test-1-20231017180653', 'default', 1, '2023-10-17 18:07:14', '2023-10-17 18:06:54');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (44, 1, 'ci-test-1-20231017181432', 'default', 1, '2023-10-17 18:14:57', '2023-10-17 18:14:33');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (45, 1, 'ci-test-1-20231017182718', 'default', 1, '2023-10-17 18:27:38', '2023-10-17 18:27:18');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (46, 1, 'ci-test-1-20231017182858', 'default', 1, '2023-10-17 18:29:18', '2023-10-17 18:28:59');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (47, 1, 'ci-test-1-20231017182908', 'default', 1, '2023-10-17 18:29:32', '2023-10-17 18:29:09');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (48, 1, 'ci-test-1-20231018092945', 'default', 1, '2023-10-18 09:30:15', '2023-10-18 09:29:45');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (49, 1, 'ci-test-1-20231018093234', 'default', 1, '2023-10-18 09:32:56', '2023-10-18 09:32:35');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (50, 1, 'ci-test-1-20231018093321', 'default', 1, '2023-10-18 09:33:42', '2023-10-18 09:33:22');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (51, 1, 'ci-test-1-20231018093434', 'default', 1, '2023-10-18 09:34:54', '2023-10-18 09:34:34');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (52, 1, 'ci-test-1-20231018093451', 'default', 1, '2023-10-18 09:35:12', '2023-10-18 09:34:52');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (53, 1, 'ci-test-1-20231018094115', 'default', 1, '2023-10-18 09:41:35', '2023-10-18 09:41:16');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (54, 1, 'ci-test-1-20231018094155', 'default', 1, '2023-10-18 09:42:16', '2023-10-18 09:41:56');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (55, 1, 'ci-test-1-20231018095107', 'default', 1, '2023-10-18 09:51:31', '2023-10-18 09:51:08');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (56, 1, 'ci-test-1-20231018095259', 'default', 1, '2023-10-18 09:53:20', '2023-10-18 09:53:00');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (57, 1, 'ci-test-1-20231018095707', 'default', 1, '2023-10-18 09:57:27', '2023-10-18 09:57:08');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (58, 1, 'ci-test-1-20231018100203', 'default', 1, '2023-10-18 10:02:27', '2023-10-18 10:02:04');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (59, 1, 'ci-test-1-20231018101716', 'default', 1, '2023-10-18 10:17:42', '2023-10-18 10:17:16');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (60, 1, 'ci-test-1-20231018112529', 'default', 2, '2023-10-18 11:25:36', '2023-10-18 11:25:29');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (61, 1, 'ci-test-1-20231018112707', 'default', 1, '2023-10-18 11:27:29', '2023-10-18 11:27:08');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (62, 1, 'ci-test-1-20231018113148', 'default', 1, '2023-10-18 11:32:08', '2023-10-18 11:31:49');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (63, 1, 'ci-test-1-20231018113226', 'default', 2, '2023-10-18 11:32:31', '2023-10-18 11:32:26');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (64, 1, 'ci-test-1-20231018113451', 'default', 2, '2023-10-18 11:34:56', '2023-10-18 11:34:51');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (65, 1, 'ci-test-1-20231018113500', 'default', 2, '2023-10-18 11:35:04', '2023-10-18 11:35:00');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (66, 1, 'ci-test-1-20231018113531', 'default', 2, '2023-10-18 11:35:37', '2023-10-18 11:35:32');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (67, 1, 'ci-test-1-20231018113608', 'default', 2, '2023-10-18 11:36:13', '2023-10-18 11:36:09');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (68, 1, 'ci-test-1-20231018113626', 'default', 2, '2023-10-18 11:36:30', '2023-10-18 11:36:26');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (69, 1, 'ci-test-1-20231018113644', 'default', 2, '2023-10-18 11:36:49', '2023-10-18 11:36:45');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (70, 1, 'ci-test-1-20231018113659', 'default', 2, '2023-10-18 11:37:04', '2023-10-18 11:36:59');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (71, 1, 'ci-test-1-20231018113754', 'default', 2, '2023-10-18 11:37:59', '2023-10-18 11:37:54');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (72, 1, 'ci-test-1-20231018113821', 'default', 2, '2023-10-18 11:38:26', '2023-10-18 11:38:22');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (73, 1, 'ci-test-1-20231018113855', 'default', 2, '2023-10-18 11:39:00', '2023-10-18 11:38:56');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (74, 1, 'ci-test-1-20231018113910', 'default', 2, '2023-10-18 11:39:15', '2023-10-18 11:39:11');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (75, 1, 'ci-test-1-20231018114015', 'default', 2, '2023-10-18 11:40:20', '2023-10-18 11:40:16');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (76, 1, 'ci-test-1-20231018114044', 'default', 2, '2023-10-18 11:40:49', '2023-10-18 11:40:45');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (77, 1, 'ci-test-1-20231018114255', 'default', 2, '2023-10-18 11:42:59', '2023-10-18 11:42:56');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (78, 1, 'ci-test-1-20231018114344', 'default', 1, '2023-10-18 11:44:05', '2023-10-18 11:43:45');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (79, 1, 'ci-test-1-20231018114415', 'default', 2, '2023-10-18 11:44:28', '2023-10-18 11:44:15');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (80, 1, 'ci-test-1-20231018114456', 'default', 2, '2023-10-18 11:45:08', '2023-10-18 11:44:56');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (81, 1, 'ci-test-1-20231018114611', 'default', 2, '2023-10-18 11:46:23', '2023-10-18 11:46:12');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (82, 1, 'ci-test-1-20231018114659', 'default', 2, '2023-10-18 11:47:13', '2023-10-18 11:47:00');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (83, 1, 'ci-test-1-20231018114854', 'default', 2, '2023-10-18 11:49:06', '2023-10-18 11:48:54');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (84, 1, 'ci-test-1-20231018115101', 'default', 2, '2023-10-18 11:51:14', '2023-10-18 11:51:02');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (85, 1, 'ci-test-1-20231018115238', 'default', 2, '2023-10-18 11:52:54', '2023-10-18 11:52:40');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (86, 1, 'ci-test-1-20231018115310', 'default', 1, '2023-10-18 11:53:39', '2023-10-18 11:53:10');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (87, 1, 'ci-test-1-20231018115505', 'default', 1, '2023-10-18 11:55:24', '2023-10-18 11:55:06');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (88, 1, 'ci-test-1-20231018150628', 'default', 2, '2023-10-18 15:06:35', '2023-10-18 15:06:29');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (89, 1, 'ci-test-1-20231018150902', 'default', 2, '2023-10-18 15:13:48', '2023-10-18 15:09:03');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (90, 1, 'ci-test-1-20231018151409', 'default', 2, '2023-10-18 15:18:48', '2023-10-18 15:14:09');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (91, 1, 'ci-test-1-20231018152016', 'default', 2, '2023-10-18 15:46:55', '2023-10-18 15:20:17');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (92, 1, 'ci-test-1-20231018154801', 'default', 2, '2023-10-18 16:15:20', '2023-10-18 15:48:02');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (93, 1, 'ci-test-1-20231018164345', 'default', 0, '2023-10-18 16:43:45', '2023-10-18 16:43:45');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (94, 1, 'ci-test-1-20231018164445', 'default', 1, '2023-10-18 16:47:12', '2023-10-18 16:44:45');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (95, 1, 'ci-test-1-20231018165322', 'default', 1, '2023-10-18 16:56:43', '2023-10-18 16:53:22');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (96, 1, 'ci-test-1-20231018165938', 'default', 1, '2023-10-18 17:02:49', '2023-10-18 16:59:39');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (97, 1, 'ci-test-1-20231018170400', 'default', 2, '2023-10-18 17:08:20', '2023-10-18 17:04:00');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (98, 1, 'ci-test-1-20231018170520', 'default', 2, '2023-10-18 17:05:47', '2023-10-18 17:05:21');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (99, 1, 'ci-test-1-20231018170702', 'default', 2, '2023-10-18 17:07:07', '2023-10-18 17:07:02');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (100, 1, 'ci-test-1-20231018170723', 'default', 1, '2023-10-18 17:07:29', '2023-10-18 17:07:24');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (101, 1, 'ci-test-1-20231018170802', 'default', 1, '2023-10-18 17:08:10', '2023-10-18 17:08:03');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (102, 1, 'ci-test-1-20231018170958', 'default', 1, '2023-10-18 17:10:05', '2023-10-18 17:09:59');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (103, 1, 'ci-test-1-20231018172817', 'default', 1, '2023-10-18 17:28:25', '2023-10-18 17:28:18');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (104, 1, 'ci-test-1-20231018172957', 'default', 1, '2023-10-18 17:30:04', '2023-10-18 17:29:57');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (105, 1, 'ci-test-1-20231018173041', 'default', 1, '2023-10-18 17:30:48', '2023-10-18 17:30:42');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (106, 1, 'ci-test-1-20231018173157', 'default', 2, '2023-10-18 17:32:03', '2023-10-18 17:31:57');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (107, 1, 'ci-test-1-20231018173215', 'default', 1, '2023-10-18 17:36:59', '2023-10-18 17:32:16');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (108, 1, 'ci-test-1-20231018174130', 'default', 1, '2023-10-18 17:48:06', '2023-10-18 17:41:31');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (109, 1, 'ci-test-1-20231018181249', 'default', 1, '2023-10-18 18:23:55', '2023-10-18 18:12:50');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (110, 1, 'ci-test-1-20231019163554', 'default', 1, '2023-10-19 16:42:34', '2023-10-19 16:35:54');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (111, 1, 'ci-test-1-20231020134146', 'default', 1, '2023-10-20 13:44:22', '2023-10-20 13:41:47');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (112, 1, 'ci-test-1-20231020174919', 'default', 3, '2023-10-26 10:46:14', '2023-10-20 17:49:20');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (113, 1, 'ci-test-1-20231020175422', 'default', 3, '2023-10-26 10:46:13', '2023-10-20 17:54:27');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (114, 1, 'ci-test-1-20231020175616', 'default', 1, '2023-10-20 18:04:41', '2023-10-20 17:56:25');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (115, 1, 'ci-test-1-20231020180151', 'default', 2, '2023-10-20 18:03:26', '2023-10-20 18:01:51');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (116, 1, 'ci-test-1-20231020180915', 'default', 2, '2023-10-20 18:09:22', '2023-10-20 18:09:15');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (117, 1, 'ci-test-1-20231020180941', 'default', 2, '2023-10-20 18:09:46', '2023-10-20 18:09:41');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (118, 1, 'ci-test-1-20231020181109', 'default', 2, '2023-10-20 18:11:15', '2023-10-20 18:11:10');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (119, 1, 'ci-test-1-20231020181304', 'default', 2, '2023-10-20 18:13:10', '2023-10-20 18:13:05');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (120, 1, 'ci-test-1-20231020181756', 'default', 3, '2023-10-26 10:42:15', '2023-10-20 18:17:57');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (121, 1, 'ci-test-1-20231026095035', 'default', 1, '2023-10-26 09:57:59', '2023-10-26 09:50:36');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (122, 1, 'ci-test-1-20231026104634', 'default', 2, '2023-10-26 10:46:45', '2023-10-26 10:46:35');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (123, 1, 'ci-test-1-20231026104956', 'default', 3, '2023-10-26 10:49:58', '2023-10-26 10:49:57');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (124, 1, 'ci-test-1-20231026181028', 'default', 1, '2023-10-26 18:18:01', '2023-10-26 18:10:28');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (125, 1, 'ci-test-1-20231027095222', 'default', 1, '2023-10-27 09:57:31', '2023-10-27 09:52:22');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (126, 1, 'ci-test-1-20231027110632', 'default', 3, '2023-10-27 11:07:10', '2023-10-27 11:06:33');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (127, 1, 'ci-test-1-20231027111748', 'default', 3, '2023-10-27 11:29:14', '2023-10-27 11:17:48');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (128, 1, 'ci-test-1-20231027112957', 'default', 3, '2023-10-27 11:31:29', '2023-10-27 11:29:58');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (129, 1, 'ci-test-1-20231027113135', 'default', 1, '2023-10-27 11:37:38', '2023-10-27 11:31:36');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (130, 1, 'ci-test-1-20231027114509', 'default', 2, '2023-10-27 11:49:39', '2023-10-27 11:45:10');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (131, 1, 'ci-test-1-20231027115018', 'default', 2, '2023-10-27 11:56:50', '2023-10-27 11:50:19');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (132, 1, 'ci-test-1-20231027115444', 'default', 1, '2023-10-27 12:01:56', '2023-10-27 11:54:44');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (133, 1, 'ci-test-1-20231027144002', 'default', 2, '2023-10-27 14:45:36', '2023-10-27 14:40:02');
INSERT INTO `ci_pipeline_run` (`id`, `pipeline_id`, `pod_name`, `namespace`, `status`, `updated_at`, `created_at`) VALUES (134, 2, 'ci-test2-2-20231027150838', 'default', 0, '2023-10-27 15:08:39', '2023-10-27 15:08:39');
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
INSERT INTO `host` (`id`, `name`, `host_addr`, `port`, `username`, `password`, `private_key`, `use_key`, `desc`, `save_session`, `updated_at`, `host_group_id`) VALUES (1, '测试机', '127.0.0.1', 22, 'zze', '123456', '', b'0', 'test', b'0', '2023-10-18 10:24:59', NULL);
INSERT INTO `host` (`id`, `name`, `host_addr`, `port`, `username`, `password`, `private_key`, `use_key`, `desc`, `save_session`, `updated_at`, `host_group_id`) VALUES (2, '测试机', '127.0.0.1', 22, 'zze', '123456', '', b'0', 'tesst', b'0', '2023-10-18 10:25:00', NULL);
INSERT INTO `host` (`id`, `name`, `host_addr`, `port`, `username`, `password`, `private_key`, `use_key`, `desc`, `save_session`, `updated_at`, `host_group_id`) VALUES (4, 'test2', '192.168.2.238', 22, 'root', '123456', '', b'0', '', b'1', '2023-10-18 14:02:37', 0);
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
) ENGINE=InnoDB AUTO_INCREMENT=40 DEFAULT CHARSET=utf8mb4;

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
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (38, 4, '192.168.2.230', 'test2', 1, 'admin', '管理员', '2023-10-17 14:24:41', 'host-sessions/4/1697523880882536.sessionb', '2023-10-17 14:24:53');
INSERT INTO `host_terminal_session` (`id`, `host_id`, `host_addr`, `host_name`, `operator_id`, `operator_name`, `operator_real_name`, `start_time`, `filepath`, `updated_at`) VALUES (39, 4, '192.168.2.238', 'test2', 1, 'admin', '管理员', '2023-10-18 10:01:50', 'host-sessions/4/1697594510454217.sessionb', '2023-10-18 10:02:10');
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
) ENGINE=InnoDB AUTO_INCREMENT=71 DEFAULT CHARSET=utf8mb4;

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
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (62, '查询流水线', 'ci-pipeline-read', 3, '', '[\"get:/ci-pipeline/page-list\", \"get:/kubernetes-config/partial-list\", \"get:/ci-pipeline-run/page-list\"]', '', '', 1, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (63, '新增流水线', 'ci-pipeline-add', 3, '', '[\"post:/ci-pipeline\"]', '', '', 2, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (64, '更新流水线', 'ci-pipeline-upt', 3, '', '[\"put:/ci-pipeline/:id\"]', '', '', 3, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (65, '删除流水线', 'ci-pipeline-del', 3, '', '[\"delete:/ci-pipeline/:id\"]', '', '', 4, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (66, '编排流水线', 'ci-pipeline-arrange', 3, '', '[\"get:/ci-pipeline/:id/config\", \"patch:/ci-pipeline/:id/config\", \"get:/secret/list\", \"get:/common/git-branch-list\"]', '', '', 5, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (67, '运行流水线', 'ci-pipeline-run', 3, '', '[\"post:/ci-pipeline/:id/run\", \"get:/common/git-branch-list\"]', '', '', 6, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (68, '查看流水线日志', 'ci-pipeline-run-log', 3, '', '[\"get:/ci-pipeline-run/:id/log\"]', '', '', 8, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (69, '取消流水线运行', 'ci-pipeline-run-cancel', 3, '', '[\"delete:/ci-pipeline-run/:id/cancel\"]', '', '', 9, b'0', b'0', b'0', 50);
INSERT INTO `permission` (`id`, `title`, `name`, `type`, `f_route`, `b_routes`, `redirect`, `icon`, `rank`, `show_link`, `show_parent`, `keep_alive`, `parent_id`) VALUES (70, '克隆流水线', 'ci-pipeline-clone', 3, '', '[\"post:/ci-pipeline/:id/clone\"]', '', '', 7, b'0', b'0', b'0', 50);
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
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Records of secret
-- ----------------------------
BEGIN;
INSERT INTO `secret` (`id`, `name`, `type`, `content`, `updated_at`) VALUES (1, '线下集群', 2, '{\"text\": \"apiVersion: v1\\nclusters:\\n- cluster:\\n    certificate-authority-data: x\\n    server: https://192.168.2.2:7443\\n  name: azj-cluster\\ncontexts:\\n- context:\\n    cluster: azj-cluster\\n    namespace: default\\n    user: admin\\n  name: offline-cluster\\ncurrent-context: offline-cluster\\nkind: Config\\npreferences: {}\\nusers:\\n- name: admin\\n  user:\\n    client-certificate-data: x=\\n    client-key-data: x=\"}', '2023-10-20 14:16:45');
INSERT INTO `secret` (`id`, `name`, `type`, `content`, `updated_at`) VALUES (2, '线下 BitBucket', 1, '{\"password\": \"xxx\", \"username\": \"zhangzhongen1\"}', '2023-10-27 11:41:44');
INSERT INTO `secret` (`id`, `name`, `type`, `content`, `updated_at`) VALUES (3, '线下 Gitlab', 1, '{\"password\": \"f/FofvWYwI/x+/1eZXHAWto=\", \"username\": \"root\"}', '2023-10-26 18:08:25');
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
