/*
 Navicat Premium Data Transfer

 Source Server         : 1.9
 Source Server Type    : MySQL
 Source Server Version : 50547
 Source Host           : 192.168.1.9:3306
 Source Schema         : ak_test

 Target Server Type    : MySQL
 Target Server Version : 50547
 File Encoding         : 65001

 Date: 08/08/2019 19:08:47
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for book_marks
-- ----------------------------
DROP TABLE IF EXISTS `book_marks`;
CREATE TABLE `book_marks`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT 'id',
  `user_id` int(10) NOT NULL DEFAULT 0 COMMENT '用户id(为0代表是官方标引)',
  `book_id` int(10) NOT NULL COMMENT '书籍序号',
  `chapter_id` int(10) NOT NULL COMMENT '篇章序号',
  `mark_type` enum('notes','bookmark','indexing') CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT 'indexing' COMMENT '注释的类型(笔记、书签、标引)',
  `chapter_name` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL DEFAULT 'null' COMMENT '章节名称',
  `chapter_value` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '书签内容(选中的内容)',
  `comm_text` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT '批注',
  `range` text CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NULL COMMENT 'range前台json',
  `create_time` datetime NULL DEFAULT NULL COMMENT '注释时间',
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '书签,笔记,标引' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for books
-- ----------------------------
DROP TABLE IF EXISTS `books`;
CREATE TABLE `books`  (
  `id` int(10) UNSIGNED NOT NULL DEFAULT 0 AUTO_INCREMENT,
  `book_name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `book_author` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `book_img` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `book_desc` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '简介',
  `book_tag` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标签',
  `state` tinyint(1) UNSIGNED NOT NULL DEFAULT 1 COMMENT '0完本,1连载',
  `hot` int(6) UNSIGNED NOT NULL DEFAULT 0,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '书籍' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for bookshelf
-- ----------------------------
DROP TABLE IF EXISTS `bookshelf`;
CREATE TABLE `bookshelf`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_id` int(10) UNSIGNED NOT NULL DEFAULT '' COMMENT '用户ID',
  `book_id` int(10) UNSIGNED NOT NULL,
  `create_time` datetime NOT NULL COMMENT '创建时间',
  `update_time` datetime NOT NULL COMMENT '更新时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_unicode_ci COMMENT = '书房收藏' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for chapters
-- ----------------------------
DROP TABLE IF EXISTS `chapters`;
CREATE TABLE `chapters`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `book_id` int(10) UNSIGNED NOT NULL,
  `chapter_name` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '章节' ROW_FORMAT = Compact;

-- ----------------------------
-- Table structure for fm_tags
-- ----------------------------
DROP TABLE IF EXISTS `fm_tags`;
CREATE TABLE `fm_tags`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(50) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT '',
  `type` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL DEFAULT 'default',
  `hot` int(6) UNSIGNED NOT NULL DEFAULT 0,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `name`(`name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 32 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '标签' ROW_FORMAT = Compact;

-- ----------------------------
-- Records of fm_tags
-- ----------------------------
INSERT INTO `fm_tags` VALUES (1, '玄幻', 'category', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (2, '修真', 'category', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (3, '都市', 'category', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (4, '穿越', 'category', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (5, '网游', 'category', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (6, '科幻', 'category', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (7, '言情', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (8, '推理', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (9, '悬疑', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (10, '爱情传奇', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (11, '历史', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (12, '军事', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (17, '现代修真', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (18, '军事历史', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (19, '架空历史', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (20, '魔幻', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (21, 'YY', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (22, '耽美', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (23, '黑道', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (24, '同人', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (25, '无限流', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (26, '系统', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (27, '兑换', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (28, '灵异', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (29, '恐怖惊悚', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (30, '古典仙侠', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');
INSERT INTO `fm_tags` VALUES (31, '官场', 'default', 0, '2019-08-08 10:10:13', '2019-08-08 10:10:13');

-- ----------------------------
-- Table structure for fm_users
-- ----------------------------
DROP TABLE IF EXISTS `fm_users`;
CREATE TABLE `fm_users`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `user_nick` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `user_pass` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `user_img` varchar(150) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `login_time` datetime NULL DEFAULT NULL,
  `create_time` datetime NULL DEFAULT NULL,
  `update_time` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `user_name`(`user_name`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci COMMENT = '用户' ROW_FORMAT = Compact;

SET FOREIGN_KEY_CHECKS = 1;
