/*
 Navicat Premium Dump SQL

 Source Server         : mysql
 Source Server Type    : MySQL
 Source Server Version : 80100 (8.1.0)
 Source Host           : localhost:3306
 Source Schema         : data-space

 Target Server Type    : MySQL
 Target Server Version : 80100 (8.1.0)
 File Encoding         : 65001

 Date: 22/11/2024 15:52:32
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for sys_department
-- ----------------------------
DROP TABLE IF EXISTS `sys_department`;
CREATE TABLE `sys_department`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `state` tinyint NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_department
-- ----------------------------
INSERT INTO `sys_department` VALUES (1, '研发部', '', '2024-07-23 17:50:33', '2024-07-23 17:50:33', 1);
INSERT INTO `sys_department` VALUES (4, '市场部', '', '2024-07-23 17:50:47', '2024-07-23 17:50:47', 1);
INSERT INTO `sys_department` VALUES (5, '财务部', '', '2024-07-23 17:50:57', '2024-07-23 17:50:57', 1);
INSERT INTO `sys_department` VALUES (6, '人力资源部', '', '2024-07-23 17:51:06', '2024-09-18 17:23:02', 2);
INSERT INTO `sys_department` VALUES (7, '后勤部', '保安、卫生、食堂', '2024-07-23 22:21:51', '2024-07-23 22:22:39', 0);

-- ----------------------------
-- Table structure for sys_department_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_department_user`;
CREATE TABLE `sys_department_user`  (
  `department_id` bigint NOT NULL,
  `user_id` bigint NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`department_id`, `user_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_department_user
-- ----------------------------
INSERT INTO `sys_department_user` VALUES (1, 1, '2024-08-01 11:36:40');
INSERT INTO `sys_department_user` VALUES (1, 2, '2024-08-01 11:36:54');
INSERT INTO `sys_department_user` VALUES (1, 3, '2024-08-01 11:37:01');
INSERT INTO `sys_department_user` VALUES (1, 4, '2024-08-01 11:37:09');
INSERT INTO `sys_department_user` VALUES (1, 12, '2024-08-01 16:01:21');

-- ----------------------------
-- Table structure for sys_login_log
-- ----------------------------
DROP TABLE IF EXISTS `sys_login_log`;
CREATE TABLE `sys_login_log`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `user_name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户id',
  `remote_ip` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'ip地址',
  `user_agent` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户代理',
  `login_result` tinyint(1) NOT NULL DEFAULT 0 COMMENT '是否成功：1-成功，2-失败',
  `result_detail` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '登录失败详情',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 5 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_login_log
-- ----------------------------
INSERT INTO `sys_login_log` VALUES (1, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36', 2, 'xxx', '2024-09-30 11:24:03', '2024-10-14 17:30:35');
INSERT INTO `sys_login_log` VALUES (2, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36', 2, '嘻嘻嘻', '2024-10-14 14:35:45', '2024-10-14 17:30:39');
INSERT INTO `sys_login_log` VALUES (3, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36', 1, '', '2024-10-14 15:33:48', '2024-10-14 17:22:25');
INSERT INTO `sys_login_log` VALUES (4, 'admin', '127.0.0.1', 'Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/129.0.0.0 Safari/537.36', 1, '', '2024-10-14 16:46:35', '2024-10-14 17:22:25');

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu`;
CREATE TABLE `sys_menu`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `path` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `icon` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `component` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `order` int NOT NULL DEFAULT 0,
  `parent_id` bigint NOT NULL DEFAULT 0,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `state` tinyint NOT NULL DEFAULT 1,
  `hidden` tinyint NOT NULL DEFAULT 0,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 30 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
INSERT INTO `sys_menu` VALUES (1, 'system', '系统管理', '/system', 'Setting', '', 'Menu', 2, 0, '2024-06-04 16:49:56', '2024-06-04 16:49:56', 1, 0);
INSERT INTO `sys_menu` VALUES (2, 'department', '部门管理', '/system/department', 'School', '/system/department/index', 'Menu', 1, 1, '2024-06-04 17:11:47', '2024-06-04 17:11:47', 1, 0);
INSERT INTO `sys_menu` VALUES (3, 'user', '用户管理', '/system/user', 'User', '/system/user/index', 'Menu', 2, 1, '2024-06-04 17:13:04', '2024-06-04 17:13:04', 1, 0);
INSERT INTO `sys_menu` VALUES (4, 'role', '角色管理', '/system/role', 'UserFilled', '/system/role/index', 'Menu', 3, 1, '2024-06-04 17:13:18', '2024-06-04 17:13:18', 1, 0);
INSERT INTO `sys_menu` VALUES (5, 'menu', '菜单管理', '/system/menu', 'Menu', '/system/menu/index', 'Menu', 4, 1, '2024-06-04 17:14:24', '2024-06-04 17:14:24', 1, 0);
INSERT INTO `sys_menu` VALUES (6, 'link', '外链', '/link', 'Link', '', 'Menu', 10, 0, '2024-06-04 17:18:29', '2024-10-24 14:34:37', 1, 0);
INSERT INTO `sys_menu` VALUES (7, 'baidu', '百度', '/link/baidu', '', 'www.baidu.com', 'Link', 10, 6, '2024-06-04 17:23:29', '2024-06-04 17:23:29', 1, 0);
INSERT INTO `sys_menu` VALUES (8, 'home', '首页', '/home', 'HomeFilled', '/home/index', 'Menu', 1, 0, '2024-06-13 16:05:59', '2024-06-13 16:05:59', 1, 0);
INSERT INTO `sys_menu` VALUES (9, 'role-add', '添加角色', '', '', '', 'Button', 0, 4, '2024-06-17 12:29:34', '2024-06-17 12:29:34', 1, 0);
INSERT INTO `sys_menu` VALUES (10, 'role-update', '修改角色', '', '', '', 'Button', 0, 4, '2024-06-17 16:31:35', '2024-06-17 16:31:35', 1, 0);
INSERT INTO `sys_menu` VALUES (11, 'role-delete', '删除角色', '', '', '', 'Button', 0, 4, '2024-06-17 16:31:59', '2024-06-17 16:31:59', 1, 0);
INSERT INTO `sys_menu` VALUES (12, 'baidu', '百度', '/baidu', 'Bowl', 'https://www.baidu.com', 'Link', 3, 0, '2024-07-03 17:15:06', '2024-07-03 17:26:29', 0, 0);
INSERT INTO `sys_menu` VALUES (13, 'menu-add', '添加菜单', '', '', '', 'Button', 0, 5, '2024-07-03 17:28:56', '2024-07-03 17:44:39', 1, 0);
INSERT INTO `sys_menu` VALUES (14, 'menu-update', '修改菜单', '', '', '', 'Button', 0, 5, '2024-07-03 17:45:35', '2024-07-03 17:45:35', 1, 0);
INSERT INTO `sys_menu` VALUES (15, 'menu-delete', '删除菜单', '', '', '', 'Button', 0, 5, '2024-07-03 17:47:27', '2024-07-03 17:47:27', 1, 0);
INSERT INTO `sys_menu` VALUES (16, 'department-add', '添加部门', '', '', '', 'Button', 0, 2, '2024-07-23 20:50:29', '2024-07-23 20:50:29', 1, 0);
INSERT INTO `sys_menu` VALUES (17, 'department-update', '更新部门', '', '', '', 'Button', 0, 2, '2024-07-23 20:56:37', '2024-07-23 20:56:37', 1, 0);
INSERT INTO `sys_menu` VALUES (18, 'department-delete', '删除部门', '', '', '', 'Button', 0, 2, '2024-07-23 20:56:59', '2024-07-23 20:56:59', 1, 0);
INSERT INTO `sys_menu` VALUES (19, 'user-update', '更新用户', '', '', '', 'Button', 0, 3, '2024-08-01 14:57:33', '2024-08-01 14:57:33', 1, 0);
INSERT INTO `sys_menu` VALUES (20, 'user-delete', '删除用户', '', '', '', 'Button', 0, 3, '2024-08-01 14:57:57', '2024-08-01 14:57:57', 1, 0);
INSERT INTO `sys_menu` VALUES (21, 'user-add', '添加用户', '', '', '', 'Button', 0, 3, '2024-08-01 15:31:39', '2024-08-01 15:31:39', 1, 0);
INSERT INTO `sys_menu` VALUES (22, 'log', '日志管理', '/log', 'Menu', '', 'Menu', 3, 0, '2024-10-14 16:48:24', '2024-10-24 14:43:04', 1, 0);
INSERT INTO `sys_menu` VALUES (23, '登录日志', '登录日志', '/log/login', 'Location', '/log/login/index', 'Menu', 0, 22, '2024-10-14 16:49:39', '2024-10-14 16:49:39', 1, 0);

-- ----------------------------
-- Table structure for sys_menu_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_menu_role`;
CREATE TABLE `sys_menu_role`  (
  `menu_id` bigint NOT NULL DEFAULT 0,
  `role_id` bigint NOT NULL DEFAULT 0,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`menu_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_menu_role
-- ----------------------------
INSERT INTO `sys_menu_role` VALUES (1, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (1, 2, '2024-07-04 10:41:54');
INSERT INTO `sys_menu_role` VALUES (2, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (2, 2, '2024-07-04 10:41:54');
INSERT INTO `sys_menu_role` VALUES (3, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (3, 2, '2024-07-04 10:41:54');
INSERT INTO `sys_menu_role` VALUES (4, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (4, 2, '2024-07-04 10:41:54');
INSERT INTO `sys_menu_role` VALUES (5, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (5, 2, '2024-07-04 10:41:54');
INSERT INTO `sys_menu_role` VALUES (8, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (8, 2, '2024-07-04 10:41:54');
INSERT INTO `sys_menu_role` VALUES (9, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (10, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (11, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (13, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (14, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (15, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (16, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (17, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (18, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (19, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (20, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (21, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (22, 1, '2024-10-24 16:26:51');
INSERT INTO `sys_menu_role` VALUES (23, 1, '2024-10-24 16:26:51');

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_role`;
CREATE TABLE `sys_role`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `state` tinyint NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_role
-- ----------------------------
INSERT INTO `sys_role` VALUES (1, '管理员', '管理员拥有最多的权限', '2024-06-03 10:55:27', '2024-06-03 10:55:27', 1);
INSERT INTO `sys_role` VALUES (2, '测试', '测试用户', '2024-06-17 14:34:51', '2024-07-02 10:25:23', 2);

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS `sys_user`;
CREATE TABLE `sys_user`  (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `desc` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `password` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `avatar` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL DEFAULT '',
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `update_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `state` tinyint NOT NULL DEFAULT 1,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 13 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user
-- ----------------------------
INSERT INTO `sys_user` VALUES (1, 'admin', '管理员', '$2a$10$qTuvYPjR8os3HltI.cwn1O.pGF2CXtFdMvK2uWfu9pqr/cXaz5V/G', 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png', '2024-06-03 10:53:45', '2024-06-03 10:53:49', 1);
INSERT INTO `sys_user` VALUES (2, 'test', '测试', '$2a$10$qTuvYPjR8os3HltI.cwn1O.pGF2CXtFdMvK2uWfu9pqr/cXaz5V/G', 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png', '2024-06-03 10:53:51', '2024-06-03 10:53:54', 1);
INSERT INTO `sys_user` VALUES (3, 'admin_1', '', '$2a$10$qTuvYPjR8os3HltI.cwn1O.pGF2CXtFdMvK2uWfu9pqr/cXaz5V/G', 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png', '2024-06-03 10:53:56', '2024-08-16 15:53:40', 1);
INSERT INTO `sys_user` VALUES (4, 'admin_2', '', '$2a$10$qTuvYPjR8os3HltI.cwn1O.pGF2CXtFdMvK2uWfu9pqr/cXaz5V/G', 'https://www.bing.com/rp/ar_9isCNU2Q-VG1yEDDHnx8HAFQ.png', '2024-06-03 10:54:01', '2024-06-03 10:54:03', 1);
INSERT INTO `sys_user` VALUES (7, 'zhangsan', '张三', '', '', '2024-06-13 09:23:12', '2024-06-13 09:23:12', 0);
INSERT INTO `sys_user` VALUES (8, 'zhangsi', '张四', '', '', '2024-06-13 09:27:56', '2024-06-13 09:27:56', 0);
INSERT INTO `sys_user` VALUES (9, 'lisi', '李四', '$2a$10$YEEP562zMqKjN5lcPCzkBuoTF3riYqf74yoN9aUo/4/x9Uu3/ukK6', '', '2024-08-01 15:57:45', '2024-08-01 15:57:45', 1);
INSERT INTO `sys_user` VALUES (10, 'wangwu', '王五', '$2a$10$s1SnisHoydXRMy7V4Dpd7eV.M2noKs1dS5CQsucj7Fey6aU/GaBwS', '', '2024-08-01 15:58:28', '2024-08-01 15:58:28', 1);
INSERT INTO `sys_user` VALUES (11, 'admin_3', '', '$2a$10$JNz62W./5nIYWADWhNNiu.YCcbYviWHFivKchkfI/Has8r6.czfMC', '', '2024-08-01 15:59:29', '2024-08-01 15:59:29', 1);

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS `sys_user_role`;
CREATE TABLE `sys_user_role`  (
  `user_id` bigint NOT NULL,
  `role_id` bigint NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`user_id`, `role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
INSERT INTO `sys_user_role` VALUES (1, 1, '2024-06-03 15:36:09');

SET FOREIGN_KEY_CHECKS = 1;
