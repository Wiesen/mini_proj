/*
 Navicat MySQL Data Transfer

 Source Server         : 192.168.56.101
 Source Server Type    : MySQL
 Source Server Version : 50722
 Source Host           : 192.168.56.101:3306
 Source Schema         : livingdb

 Target Server Type    : MySQL
 Target Server Version : 50722
 File Encoding         : 65001

 Date: 02/05/2018 23:00:55
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comment
-- ----------------------------
DROP TABLE IF EXISTS `comment`;
CREATE TABLE `comment`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `emotion_id` int(11) UNSIGNED NOT NULL COMMENT '心情ID',
  `content` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '评论内容',
  `poster` int(11) UNSIGNED NOT NULL COMMENT '发布人id',
  `create_time` datetime(0) NOT NULL COMMENT '时间',
  `rspto` int(11) NULL DEFAULT NULL COMMENT '被回复人id',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `comment_eid_ref`(`emotion_id`) USING BTREE,
  INDEX `comment_poster_ref`(`poster`) USING BTREE,
  CONSTRAINT `comment_eid_ref` FOREIGN KEY (`emotion_id`) REFERENCES `emotion` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `comment_poster_ref` FOREIGN KEY (`poster`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for emotion
-- ----------------------------
DROP TABLE IF EXISTS `emotion`;
CREATE TABLE `emotion`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `content` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '心情内容',
  `label_id` tinyint(4) UNSIGNED NOT NULL COMMENT '心情标签ID，需存在标签表中',
  `strong` tinyint(4) NOT NULL COMMENT '强度',
  `create_time` datetime(0) NOT NULL COMMENT '创建时间',
  `visiable` tinyint(4) NOT NULL COMMENT '1. 个人可见；2. 社区可见',
  `poster` int(11) UNSIGNED NOT NULL COMMENT '发布人id',
  `comment_cnt` int(11) UNSIGNED NOT NULL,
  `like_cnt` int(11) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `emotion_lid_ref`(`label_id`) USING BTREE,
  INDEX `emotion_poster_ref`(`poster`) USING BTREE,
  CONSTRAINT `emotion_lid_ref` FOREIGN KEY (`label_id`) REFERENCES `label` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `emotion_poster_ref` FOREIGN KEY (`poster`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 4 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for label
-- ----------------------------
DROP TABLE IF EXISTS `label`;
CREATE TABLE `label`  (
  `id` tinyint(4) UNSIGNED NOT NULL AUTO_INCREMENT,
  `label_name` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '标签名',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for like
-- ----------------------------
DROP TABLE IF EXISTS `like`;
CREATE TABLE `like`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `emotion_id` int(11) UNSIGNED NOT NULL COMMENT '心情id',
  `poster` int(11) UNSIGNED NOT NULL COMMENT '发布人id',
  `create_time` datetime(0) NOT NULL COMMENT '时间',
  PRIMARY KEY (`id`) USING BTREE,
  INDEX `like_eid_ref`(`emotion_id`) USING BTREE,
  INDEX `like_poster_ref`(`poster`) USING BTREE,
  CONSTRAINT `like_eid_ref` FOREIGN KEY (`emotion_id`) REFERENCES `emotion` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `like_poster_ref` FOREIGN KEY (`poster`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `qq_number` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL UNIQUE COMMENT 'qq号',
  `phone_number` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL UNIQUE COMMENT '手机号',
  `password` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL,
  `nickname` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL UNIQUE COMMENT '昵称',
  `token` varchar(64) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '登录token',
  `avatar` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '头像',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE INDEX `phone_number_uq_index`(`phone_number`) USING BTREE,
  UNIQUE INDEX `nickname_uq_index`(`nickname`) USING BTREE,
  UNIQUE INDEX `qq_number_uq_index`(`qq_number`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `message`;
CREATE TABLE `message`  (
  `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT,
  `emotion_id` int(11) UNSIGNED NOT NULL COMMENT '心情id',
  `poster` int(11) UNSIGNED NOT NULL COMMENT '发布人id',
  `owner` int(11) UNSIGNED NOT NULL COMMENT 'owner id',
  `create_time` datetime(0) NOT NULL COMMENT '时间',
  `content` varchar(256) CHARACTER SET utf8 COLLATE utf8_general_ci NOT NULL COMMENT '评论内容',
  `type_id` int(11) UNSIGNED NOT NULL,
  PRIMARY KEY (`id`) USING BTREE,
  CONSTRAINT `message_eid_ref` FOREIGN KEY (`emotion_id`) REFERENCES `emotion` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `message_poster_ref` FOREIGN KEY (`poster`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `message_owner_ref` FOREIGN KEY (`owner`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE = InnoDB AUTO_INCREMENT = 3 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;


SET FOREIGN_KEY_CHECKS = 1;
