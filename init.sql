SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_api
-- ----------------------------
DROP TABLE IF EXISTS `admin_api`;
CREATE TABLE `admin_api`  (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '接口名称',
  `group` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分组',
  `path` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '接口路径',
  `method` char(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '请求方式',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_path_method`(`path`, `method`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台api表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_api
-- ----------------------------
INSERT INTO `admin_api` VALUES ('00abe856-398c-47a6-a42a-760d346eed0a', '修改用户角色', '用户管理', '/admin/user/roleUpdate', 'PUT', '2022-08-17 11:15:43', '2022-08-17 17:43:12');
INSERT INTO `admin_api` VALUES ('049046b7-437a-4cec-863e-4486e247374e', '编辑', '配置管理', '/admin/config/update', 'PUT', '2022-08-22 14:49:38', '2022-08-22 14:49:38');
INSERT INTO `admin_api` VALUES ('079c39f6-0f43-4d44-90cb-a3e93ffadbab', '排序', '菜单管理', '/admin/menu/sort', 'PUT', '2022-08-17 17:49:57', '2022-08-17 17:49:57');
INSERT INTO `admin_api` VALUES ('0e96c3f6-55e1-46b8-832a-cd25d8f6f804', '删除', '菜单管理', '/admin/menu/delete', 'DELETE', '2022-08-17 17:49:44', '2022-08-17 17:49:44');
INSERT INTO `admin_api` VALUES ('26ef8e36-583f-43e4-a6f7-cc08be5947e2', '编辑', 'api管理', '/admin/api/update', 'PUT', '2022-08-17 17:44:31', '2022-08-17 17:44:31');
INSERT INTO `admin_api` VALUES ('2d8b5915-a067-491f-a3b8-7454edee6055', '详情', '后台基础', '/admin/user/info', 'GET', '2022-08-17 11:11:55', '2022-08-18 14:06:58');
INSERT INTO `admin_api` VALUES ('33462df2-88c8-4877-8d3a-7a997920f9a4', '创建', '用户管理', '/admin/user/create', 'POST', '2022-08-17 11:14:17', '2022-08-17 14:02:56');
INSERT INTO `admin_api` VALUES ('46b79bf3-eac3-4b2c-ac07-10e7831a5cbe', '获取所有api分组', 'api管理', '/admin/api/group', 'GET', '2022-08-17 17:45:17', '2022-08-17 17:45:30');
INSERT INTO `admin_api` VALUES ('5e435d97-5717-4b7c-b46c-23d2c373833b', '获取角色权限', '角色管理', '/admin/role/auth', 'GET', '2022-08-17 17:42:33', '2022-08-17 17:45:58');
INSERT INTO `admin_api` VALUES ('61a1ef46-8a5e-4cc1-b35d-a46777a58198', '创建', '菜单管理', '/admin/menu/create', 'POST', '2022-08-17 17:47:09', '2022-08-17 17:47:09');
INSERT INTO `admin_api` VALUES ('665a9069-f9e6-4b9f-b36e-7ae022fc3e8d', '编辑', '菜单管理', '/admin/menu/update', 'PUT', '2022-08-17 17:49:07', '2022-08-17 17:49:07');
INSERT INTO `admin_api` VALUES ('6a48389c-5c17-434a-beb4-ba08e7f738ab', '删除', '用户管理', '/admin/user/delete', 'DELETE', '2022-08-17 11:14:49', '2022-08-17 14:03:04');
INSERT INTO `admin_api` VALUES ('70478682-6ae7-49b1-a641-1f2e04296cf3', '列表', '角色管理', '/admin/role/list', 'GET', '2022-08-17 17:38:45', '2022-08-17 17:38:45');
INSERT INTO `admin_api` VALUES ('78348ad4-ec3b-45db-9087-6c3ac01584c1', '列表', 'api管理', '/admin/api/list', 'GET', '2022-08-17 17:43:40', '2022-08-17 17:43:40');
INSERT INTO `admin_api` VALUES ('a12827f6-a7e2-4650-b131-a776a009c4a1', '获取所有角色', '用户管理', '/admin/user/role', 'GET', '2022-08-17 11:16:17', '2022-08-17 17:45:36');
INSERT INTO `admin_api` VALUES ('a876b7b7-feaa-4be6-a4bf-30bd43d275a6', '编辑', '角色管理', '/admin/role/update', 'PUT', '2022-08-17 17:42:05', '2022-08-17 17:42:05');
INSERT INTO `admin_api` VALUES ('b95a2f1b-1d1a-4426-a7da-08d9f529e950', '修改密码', '用户管理', '/admin/user/passwordUpdate', 'PUT', '2022-08-17 11:16:45', '2022-08-17 14:03:35');
INSERT INTO `admin_api` VALUES ('c5730910-c962-4d72-92b5-d2182e7f4014', '删除', '角色管理', '/admin/role/delete', 'DELETE', '2022-08-17 17:42:18', '2022-08-17 17:42:18');
INSERT INTO `admin_api` VALUES ('cbfa623c-1e33-4df5-9fdf-be58200bb516', '创建', 'api管理', '/admin/api/create', 'POST', '2022-08-17 17:44:16', '2022-08-17 17:44:16');
INSERT INTO `admin_api` VALUES ('cd6b1856-41dc-400e-b184-7af3bb7ebc9c', '创建', '角色管理', '/admin/role/create', 'POST', '2022-08-17 17:39:07', '2022-08-17 17:39:07');
INSERT INTO `admin_api` VALUES ('cf15d782-7cc2-4a84-8898-a5d59ec46a54', '获取自身权限菜单', '后台基础', '/admin/user/menu', 'GET', '2022-08-17 11:12:42', '2022-08-17 17:45:47');
INSERT INTO `admin_api` VALUES ('cf6c3be0-2c62-4853-8815-e1f6af2ab037', '详情', '配置管理', '/admin/config/list', 'GET', '2022-08-22 14:49:02', '2022-08-22 14:49:02');
INSERT INTO `admin_api` VALUES ('d6d90a00-f4fd-4883-bd14-65d0267bd5ea', '列表', '菜单管理', '/admin/menu/list', 'GET', '2022-08-17 17:46:31', '2022-08-17 17:46:31');
INSERT INTO `admin_api` VALUES ('e7c19820-f676-44a3-ba80-29232158ee29', '编辑', '用户管理', '/admin/user/update', 'PUT', '2022-08-17 11:14:32', '2022-08-17 14:03:48');
INSERT INTO `admin_api` VALUES ('ea34edd6-c9fe-456e-a08c-8c66fa5ff1b6', '列表', '用户管理', '/admin/user/list', 'GET', '2022-08-17 11:13:22', '2022-08-17 14:03:24');
INSERT INTO `admin_api` VALUES ('f76cdc76-2cf3-48cf-9783-b76907e360c8', '修改角色权限', '角色管理', '/admin/role/authUpdate', 'PUT', '2022-08-17 17:42:55', '2022-08-17 17:42:55');
INSERT INTO `admin_api` VALUES ('f7ed923d-5d2e-4dfd-aa97-d2cfae0497ad', '删除', 'api管理', '/admin/api/delete', 'DELETE', '2022-08-17 17:44:48', '2022-08-17 17:44:48');
INSERT INTO `admin_api` VALUES ('f7fb8570-9543-4aa2-9760-9993ea1e541a', '修改自身信息', '后台基础', '/admin/user/info/update', 'PUT', '2022-08-17 11:17:14', '2022-08-18 14:07:04');

-- ----------------------------
-- Table structure for admin_config
-- ----------------------------
DROP TABLE IF EXISTS `admin_config`;
CREATE TABLE `admin_config`  (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `group` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '分组',
  `key` char(50) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '配置key',
  `value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '配置value',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_group_key`(`group`, `key`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台配置表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_config
-- ----------------------------

-- ----------------------------
-- Table structure for admin_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_menu`;
CREATE TABLE `admin_menu`  (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `parent_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '父菜单id',
  `path` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由path',
  `name` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '路由name',
  `component` char(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '文件路径',
  `icon` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单图标',
  `title` char(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '菜单名称',
  `sort` smallint(5) UNSIGNED NOT NULL DEFAULT 0 COMMENT '排序 越小排序越靠前',
  `is_hidden` tinyint(1) UNSIGNED NOT NULL DEFAULT 0 COMMENT '是否隐藏菜单 0显示 1隐藏 ',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_sort`(`sort`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台菜单表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_menu
-- ----------------------------
INSERT INTO `admin_menu` VALUES ('1d555d33-f209-426f-b609-e01d1acd2438', '6bd826cf-4998-442e-aca0-f6093b6bdcd1', 'user', 'systemUser', 'view/admin/system/user/index.vue', 'user', '用户管理', 3, 0, '2022-07-22 14:46:05', '2022-08-18 15:32:30');
INSERT INTO `admin_menu` VALUES ('2bc8b4dc-2250-42d4-a833-c3876cba5ec2', '6bd826cf-4998-442e-aca0-f6093b6bdcd1', 'role', 'systemRole', 'view/admin/system/role/index.vue', 's-custom', '角色管理', 4, 0, '2022-07-22 14:46:05', '2022-08-12 10:31:15');
INSERT INTO `admin_menu` VALUES ('444ef6e7-42e4-4bf4-b96e-96e38c5f02c0', '6bd826cf-4998-442e-aca0-f6093b6bdcd1', 'menu', 'systemMenu', 'view/admin/system/menu/index.vue', 'menu', '菜单管理', 5, 0, '2022-07-22 14:46:05', '2022-08-18 15:32:46');
INSERT INTO `admin_menu` VALUES ('6bd826cf-4998-442e-aca0-f6093b6bdcd1', '', 'system', 'system', 'view/admin/system/index.vue', 'setting', '系统管理', 2, 0, '2022-07-22 14:46:05', '2022-08-18 15:32:15');
INSERT INTO `admin_menu` VALUES ('6e2be136-848c-43a7-a170-ed15b237dbd3', '', 'home', 'home', 'view/admin/home/index.vue', 's-home', '首页', 1, 0, '2022-07-22 14:46:05', '2022-08-12 10:58:09');
INSERT INTO `admin_menu` VALUES ('8b906223-7b82-430e-ba92-1e1300990a3c', '6bd826cf-4998-442e-aca0-f6093b6bdcd1', 'api', 'systemApi', 'view/admin/system/api/index.vue', 's-operation', 'api管理', 6, 0, '2022-07-22 14:46:05', '2022-08-18 15:35:24');
INSERT INTO `admin_menu` VALUES ('9489fa67-47cc-49ef-a225-a8cd7642f356', '6bd826cf-4998-442e-aca0-f6093b6bdcd1', 'person', 'systemPerson', 'view/admin/system/person/index.vue', '', '个人信息', 7, 1, '2022-07-22 14:46:05', '2022-08-18 15:44:00');
INSERT INTO `admin_menu` VALUES ('9e8ae06c-e26f-4229-b58e-6c0534be4b1e', '6bd826cf-4998-442e-aca0-f6093b6bdcd1', 'config', 'systemConfig', 'view/admin/system/config/index.vue', 'set-up', '配置管理', 8, 0, '2022-08-22 14:19:54', '2022-08-22 14:19:54');

-- ----------------------------
-- Table structure for admin_role
-- ----------------------------
DROP TABLE IF EXISTS `admin_role`;
CREATE TABLE `admin_role`  (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色名称',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台角色表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role
-- ----------------------------
INSERT INTO `admin_role` VALUES ('63ac2ba6-5b84-499a-aae0-a19971045498', '超级管理员', '2022-08-04 16:27:58', '2022-08-05 14:22:10');

-- ----------------------------
-- Table structure for admin_role_api
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_api`;
CREATE TABLE `admin_role_api`  (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `admin_role_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色id',
  `admin_api_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'api id',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_admin_role_id`(`admin_role_id`) USING BTREE,
  INDEX `idx_admin_api_id`(`admin_api_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台角色与api关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role_api
-- ----------------------------
INSERT INTO `admin_role_api` VALUES ('01da9079-26b7-470f-a892-43503de6fd2c', '63ac2ba6-5b84-499a-aae0-a19971045498', '26ef8e36-583f-43e4-a6f7-cc08be5947e2', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('0d4825a9-6e61-4bfb-8b6b-93dca06bb366', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '5e435d97-5717-4b7c-b46c-23d2c373833b', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('0ff290d0-edf5-4382-b893-3d96caf9be84', '63ac2ba6-5b84-499a-aae0-a19971045498', 'cf6c3be0-2c62-4853-8815-e1f6af2ab037', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('11bcc529-8d79-447e-969b-6bef349373fe', '63ac2ba6-5b84-499a-aae0-a19971045498', '00abe856-398c-47a6-a42a-760d346eed0a', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('1255c472-0404-4b0c-98fc-11f01a4f3b04', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'd6d90a00-f4fd-4883-bd14-65d0267bd5ea', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('12c39d7e-2745-4eb3-85a4-ca9ffdc0ab80', '63ac2ba6-5b84-499a-aae0-a19971045498', 'a876b7b7-feaa-4be6-a4bf-30bd43d275a6', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('14670315-29ae-410c-894a-24188ee484e7', '63ac2ba6-5b84-499a-aae0-a19971045498', '079c39f6-0f43-4d44-90cb-a3e93ffadbab', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('151479a4-99b1-4432-9647-e16d1ffcabe7', '63ac2ba6-5b84-499a-aae0-a19971045498', '5e435d97-5717-4b7c-b46c-23d2c373833b', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('17911b78-3348-4f59-87a0-747f399864ed', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '079c39f6-0f43-4d44-90cb-a3e93ffadbab', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('1c06c16b-6bc7-4b2a-9051-981cd65a2cb1', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '70478682-6ae7-49b1-a641-1f2e04296cf3', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('1dcde0d3-d1e8-4d02-86c4-923bbb2b11ce', '63ac2ba6-5b84-499a-aae0-a19971045498', '665a9069-f9e6-4b9f-b36e-7ae022fc3e8d', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('1de6f8b0-0a8b-4a70-9d4b-976c42543cbf', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '78348ad4-ec3b-45db-9087-6c3ac01584c1', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('26561192-4568-4092-976c-f5af911c5d0d', '63ac2ba6-5b84-499a-aae0-a19971045498', 'cf15d782-7cc2-4a84-8898-a5d59ec46a54', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('2dd68249-cc4e-4e0c-8f28-8814eb87be6a', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'b95a2f1b-1d1a-4426-a7da-08d9f529e950', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('3443cfa5-dd5c-460f-92f4-fd2d7e2f2a5f', '63ac2ba6-5b84-499a-aae0-a19971045498', '0e96c3f6-55e1-46b8-832a-cd25d8f6f804', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('369de122-5f01-44a3-87ef-e2dae5c0373c', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'cf15d782-7cc2-4a84-8898-a5d59ec46a54', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('46000353-3dac-409d-bb2a-cbb375ffeddb', '63ac2ba6-5b84-499a-aae0-a19971045498', 'f7fb8570-9543-4aa2-9760-9993ea1e541a', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('4a3a59e7-7ceb-4fd0-8b9e-1cc8bd2d9a1f', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'a876b7b7-feaa-4be6-a4bf-30bd43d275a6', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('4ebeebe0-76e8-4845-ba96-eafa2ba8a553', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'f7fb8570-9543-4aa2-9760-9993ea1e541a', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('52160a6b-3782-41ec-90e5-d60742701e72', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '46b79bf3-eac3-4b2c-ac07-10e7831a5cbe', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('61f3843d-df62-4f9d-ba80-5109938c9167', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '26ef8e36-583f-43e4-a6f7-cc08be5947e2', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('66578d54-fbdf-41d3-b365-4dd495d17b56', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'a12827f6-a7e2-4650-b131-a776a009c4a1', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('6a0e6a04-8a19-4170-833d-a716e8558081', '63ac2ba6-5b84-499a-aae0-a19971045498', '049046b7-437a-4cec-863e-4486e247374e', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('6f9a722d-d7f3-4c71-a0a3-bc54f39878ed', '63ac2ba6-5b84-499a-aae0-a19971045498', 'b95a2f1b-1d1a-4426-a7da-08d9f529e950', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('74e5d485-c081-4e9a-b254-a3f585418dc4', '63ac2ba6-5b84-499a-aae0-a19971045498', '46b79bf3-eac3-4b2c-ac07-10e7831a5cbe', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('771b9cc3-7ecd-4973-b486-3b61bc1ffb69', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'cbfa623c-1e33-4df5-9fdf-be58200bb516', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('7932c040-81ab-40ee-a6e3-c200faffb04f', '63ac2ba6-5b84-499a-aae0-a19971045498', 'cd6b1856-41dc-400e-b184-7af3bb7ebc9c', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('80a73ce2-20ce-47c7-a34a-9ff281a13b79', '63ac2ba6-5b84-499a-aae0-a19971045498', 'e7c19820-f676-44a3-ba80-29232158ee29', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('8701c3ae-e9c9-43b6-92eb-9319d9380276', '63ac2ba6-5b84-499a-aae0-a19971045498', '61a1ef46-8a5e-4cc1-b35d-a46777a58198', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('8bd94e19-5117-4b13-ab08-a680e4f8408d', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '33462df2-88c8-4877-8d3a-7a997920f9a4', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('8e033f63-1692-4e38-8d05-3732e0041a61', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '665a9069-f9e6-4b9f-b36e-7ae022fc3e8d', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('93513fc6-667a-4184-ac07-2d2692a86b0a', '63ac2ba6-5b84-499a-aae0-a19971045498', '78348ad4-ec3b-45db-9087-6c3ac01584c1', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('94161698-e481-4c43-b3bf-968cc5db6827', '63ac2ba6-5b84-499a-aae0-a19971045498', 'd6d90a00-f4fd-4883-bd14-65d0267bd5ea', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('9524dae2-5bd3-41d4-a796-43a178d0d7c6', '63ac2ba6-5b84-499a-aae0-a19971045498', '2d8b5915-a067-491f-a3b8-7454edee6055', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('9f783e35-4984-4cf1-b30b-f5c95eccdb89', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '00abe856-398c-47a6-a42a-760d346eed0a', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('a18f133a-383d-4758-b628-23427a69a169', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'e7c19820-f676-44a3-ba80-29232158ee29', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('a8393036-4687-4e82-86f3-f2535f2631bb', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'cd6b1856-41dc-400e-b184-7af3bb7ebc9c', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('ab01ef74-12b9-4807-b811-8a7d4557141b', '63ac2ba6-5b84-499a-aae0-a19971045498', '70478682-6ae7-49b1-a641-1f2e04296cf3', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('b151fc7d-207d-4ae5-a8db-ff4c10713b1e', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'f7ed923d-5d2e-4dfd-aa97-d2cfae0497ad', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('b3432b52-1a1d-41cf-a9d2-0e805484a328', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'f76cdc76-2cf3-48cf-9783-b76907e360c8', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('bb2a7d39-2314-4ad7-bbf8-f26c16a994d5', '63ac2ba6-5b84-499a-aae0-a19971045498', 'cbfa623c-1e33-4df5-9fdf-be58200bb516', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('bfc638e6-1828-4307-a3d0-087e470bf81f', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'ea34edd6-c9fe-456e-a08c-8c66fa5ff1b6', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('c1cd7231-2207-4ece-a740-0ab84506c83f', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '2d8b5915-a067-491f-a3b8-7454edee6055', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('c5c6e61c-1298-431e-bf81-01275f9c7c54', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '0e96c3f6-55e1-46b8-832a-cd25d8f6f804', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('cc626a5d-33d7-488b-a6dd-1ea24876a0a1', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', 'c5730910-c962-4d72-92b5-d2182e7f4014', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('cdaf603b-c4fe-4198-84dc-c97e6bb5eb0d', '63ac2ba6-5b84-499a-aae0-a19971045498', 'a12827f6-a7e2-4650-b131-a776a009c4a1', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('d15eb138-b2aa-48c1-bebb-7308bb39b0dc', '63ac2ba6-5b84-499a-aae0-a19971045498', 'ea34edd6-c9fe-456e-a08c-8c66fa5ff1b6', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('d37ba993-4702-4484-a7ae-5244fc4a0976', '63ac2ba6-5b84-499a-aae0-a19971045498', '6a48389c-5c17-434a-beb4-ba08e7f738ab', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('d7976804-3b9a-43bd-8197-ee1aff3a0fec', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '6a48389c-5c17-434a-beb4-ba08e7f738ab', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('db4ff7a1-456f-474b-9656-d0b29b68faa8', '63ac2ba6-5b84-499a-aae0-a19971045498', 'f76cdc76-2cf3-48cf-9783-b76907e360c8', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('de7be626-c5ab-4ad0-a9d8-e55ffebc97a7', '63ac2ba6-5b84-499a-aae0-a19971045498', 'f7ed923d-5d2e-4dfd-aa97-d2cfae0497ad', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('df96d9a2-d38e-4bbc-8377-f806faa39578', '63ac2ba6-5b84-499a-aae0-a19971045498', 'c5730910-c962-4d72-92b5-d2182e7f4014', '2022-08-22 14:50:30', '2022-08-22 14:50:30');
INSERT INTO `admin_role_api` VALUES ('e1464f4a-8e3d-4690-a4a1-63cec1c9e63f', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '61a1ef46-8a5e-4cc1-b35d-a46777a58198', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_api` VALUES ('e4fe515a-fd8e-4eaa-9700-43b37f8cb7c4', '63ac2ba6-5b84-499a-aae0-a19971045498', '33462df2-88c8-4877-8d3a-7a997920f9a4', '2022-08-22 14:50:30', '2022-08-22 14:50:30');

-- ----------------------------
-- Table structure for admin_role_menu
-- ----------------------------
DROP TABLE IF EXISTS `admin_role_menu`;
CREATE TABLE `admin_role_menu`  (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `admin_role_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '管理员id',
  `admin_menu_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '菜单id',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `idx_admin_role_id`(`admin_role_id`) USING BTREE,
  INDEX `idx_admin_menu_id`(`admin_menu_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台角色与菜单关联表' ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_role_menu
-- ----------------------------
INSERT INTO `admin_role_menu` VALUES ('0a298db0-172d-4dcd-89d6-688ad6a839fa', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '8b906223-7b82-430e-ba92-1e1300990a3c', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_menu` VALUES ('12d5135a-0075-4bbb-96da-7535ab54f9f8', '63ac2ba6-5b84-499a-aae0-a19971045498', '2bc8b4dc-2250-42d4-a833-c3876cba5ec2', '2022-08-22 14:50:31', '2022-08-22 14:50:31');
INSERT INTO `admin_role_menu` VALUES ('16474571-af56-46f4-a856-2b58ca01e172', '63ac2ba6-5b84-499a-aae0-a19971045498', '444ef6e7-42e4-4bf4-b96e-96e38c5f02c0', '2022-08-22 14:50:31', '2022-08-22 14:50:31');
INSERT INTO `admin_role_menu` VALUES ('1e811f90-530f-41d1-9bed-d4ccd4171ad2', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '444ef6e7-42e4-4bf4-b96e-96e38c5f02c0', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_menu` VALUES ('36aaf346-b07b-4d20-b8b5-86a499bc1ed8', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '6e2be136-848c-43a7-a170-ed15b237dbd3', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_menu` VALUES ('36b96ab9-c759-42a7-8d8e-80234fb5e9b6', '63ac2ba6-5b84-499a-aae0-a19971045498', '9489fa67-47cc-49ef-a225-a8cd7642f356', '2022-08-22 14:50:31', '2022-08-22 14:50:31');
INSERT INTO `admin_role_menu` VALUES ('4c0d4aaf-68ef-480d-b007-b35e4bf98c02', '63ac2ba6-5b84-499a-aae0-a19971045498', '8b906223-7b82-430e-ba92-1e1300990a3c', '2022-08-22 14:50:31', '2022-08-22 14:50:31');
INSERT INTO `admin_role_menu` VALUES ('9b0ce165-990d-46fc-96d4-47cf4a77fe73', '63ac2ba6-5b84-499a-aae0-a19971045498', '1d555d33-f209-426f-b609-e01d1acd2438', '2022-08-22 14:50:31', '2022-08-22 14:50:31');
INSERT INTO `admin_role_menu` VALUES ('b122f531-2c1e-40d7-bc13-3d447fa881de', '63ac2ba6-5b84-499a-aae0-a19971045498', '9e8ae06c-e26f-4229-b58e-6c0534be4b1e', '2022-08-22 14:50:31', '2022-08-22 14:50:31');
INSERT INTO `admin_role_menu` VALUES ('cab38f89-6f66-4b2c-98ba-4c7139dea0ff', '63ac2ba6-5b84-499a-aae0-a19971045498', '6bd826cf-4998-442e-aca0-f6093b6bdcd1', '2022-08-22 14:50:31', '2022-08-22 14:50:31');
INSERT INTO `admin_role_menu` VALUES ('d315e44b-37d5-472b-93d1-60e4d7d6679c', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '9489fa67-47cc-49ef-a225-a8cd7642f356', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_menu` VALUES ('d7ccedf3-4c5b-4421-a6f6-8c9cb4dcd86b', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '6bd826cf-4998-442e-aca0-f6093b6bdcd1', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_menu` VALUES ('e3d46048-b03b-481b-bc90-91f20f052623', '63ac2ba6-5b84-499a-aae0-a19971045498', '6e2be136-848c-43a7-a170-ed15b237dbd3', '2022-08-22 14:50:31', '2022-08-22 14:50:31');
INSERT INTO `admin_role_menu` VALUES ('fcac754e-7899-419d-95aa-9b9f100354ac', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '1d555d33-f209-426f-b609-e01d1acd2438', '2022-08-18 09:14:25', '2022-08-18 09:14:25');
INSERT INTO `admin_role_menu` VALUES ('fd8fac03-26be-4c1c-9e77-b57c10a93098', 'ecfa23d7-54b6-4ff9-bca1-376ed41c241e', '2bc8b4dc-2250-42d4-a833-c3876cba5ec2', '2022-08-18 09:14:25', '2022-08-18 09:14:25');

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL,
  `username` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  `nickname` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '昵称',
  `admin_role_id` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT '角色id',
  `password` char(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'hash后的密码',
  `salt` char(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '' COMMENT 'hash密码的时候加入的salt',
  `last_login_time` datetime NULL DEFAULT NULL COMMENT '最后登陆时间',
  `create_time` datetime NULL DEFAULT NULL COMMENT '创建时间',
  `update_time` datetime NULL DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `uk_user_username`(`username`) USING BTREE,
  INDEX `idx_admin_role_id`(`admin_role_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '后台用户表' ROW_FORMAT = DYNAMIC;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES ('ab99cf87-62d4-46e5-9851-b084d168e702', 'admin', '傍晚升起的太阳', '63ac2ba6-5b84-499a-aae0-a19971045498', '8308aa8b2198a8d330c21f9bce0ca8cf', '8887c049-b931-4cf9-9c94-ea5db288bacb', '2022-08-23 09:01:04', '2022-05-16 17:05:46', '2022-08-18 14:11:09');

SET FOREIGN_KEY_CHECKS = 1;
