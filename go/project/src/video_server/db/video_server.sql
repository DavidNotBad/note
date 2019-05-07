/*
Navicat MySQL Data Transfer

Source Server         : mysql
Source Server Version : 50551
Source Host           : 127.0.0.1:3306
Source Database       : video_server

Target Server Type    : MYSQL
Target Server Version : 50551
File Encoding         : 65001

Date: 2018-08-11 15:55:22
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments` (
  `id` varchar(255) NOT NULL,
  `video_id` varchar(255) DEFAULT NULL,
  `author_id` int(11) DEFAULT NULL,
  `content` text,
  `time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES ('4dfedbee-9d60-4f7b-bea3-6f448b8a8dce', '12345', '1', '这是视屏不错，可以分享一下给我吗', null);
INSERT INTO `comments` VALUES ('590ee2b5-2b20-407c-bd04-bc373d7fe4d7', '12345', '1', '这是视屏不错，可以分享一下给我吗', null);
INSERT INTO `comments` VALUES ('ce095acf-fb24-41a5-8aac-9ea41bbbbe2a', '12345', '1', '这是视屏不错，可以分享一下给我吗', null);

-- ----------------------------
-- Table structure for sessions
-- ----------------------------
DROP TABLE IF EXISTS `sessions`;
CREATE TABLE `sessions` (
  `session_id``` varchar(255) NOT NULL,
  `TTL` tinytext,
  `login_name` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`session_id```)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of sessions
-- ----------------------------

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `login_name` varchar(255) NOT NULL,
  `pwd` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=22 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES ('10', 'avenssi', '123');
INSERT INTO `users` VALUES ('11', 'avenssi', '123');
INSERT INTO `users` VALUES ('12', 'avenssi', '123456');
INSERT INTO `users` VALUES ('13', 'avenssi', '1234561');
INSERT INTO `users` VALUES ('14', 'avenssi', '1234561');
INSERT INTO `users` VALUES ('15', 'avenssi', '1234561');
INSERT INTO `users` VALUES ('16', 'avenssi', '1234561');
INSERT INTO `users` VALUES ('17', 'avenssi', '1234561');
INSERT INTO `users` VALUES ('18', 'avenssi', '1234561');
INSERT INTO `users` VALUES ('19', 'avenssi', '1234561');
INSERT INTO `users` VALUES ('20', 'avenssi', '1234561');
INSERT INTO `users` VALUES ('21', 'avenssi', '1234561');

-- ----------------------------
-- Table structure for video_del_rec
-- ----------------------------
DROP TABLE IF EXISTS `video_del_rec`;
CREATE TABLE `video_del_rec` (
  `video_id` varchar(255) NOT NULL,
  PRIMARY KEY (`video_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of video_del_rec
-- ----------------------------

-- ----------------------------
-- Table structure for video_info
-- ----------------------------
DROP TABLE IF EXISTS `video_info`;
CREATE TABLE `video_info` (
  `id` varchar(255) NOT NULL,
  `author_id` int(11) DEFAULT NULL,
  `name` text,
  `display_ctime` text,
  `create_time` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of video_info
-- ----------------------------
INSERT INTO `video_info` VALUES ('1ee39758-83cb-4aa4-9f4c-49a079831139', '1', 'my-video', 'Aug 11 2018, 10:38:40', null);
