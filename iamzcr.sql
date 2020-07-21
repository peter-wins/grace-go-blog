/*
Navicat MySQL Data Transfer

Source Server         : oa
Source Server Version : 50529
Source Host           : localhost:2002
Source Database       : iamzcr

Target Server Type    : MYSQL
Target Server Version : 50529
File Encoding         : 65001

Date: 2020-07-21 22:01:45
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for zcr_article
-- ----------------------------
DROP TABLE IF EXISTS `zcr_article`;
CREATE TABLE `zcr_article` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(16) DEFAULT NULL,
  `title` varchar(255) NOT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `sorts` int(11) DEFAULT NULL,
  `content` text NOT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=54 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for zcr_article_category
-- ----------------------------
DROP TABLE IF EXISTS `zcr_article_category`;
CREATE TABLE `zcr_article_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `category_id` int(11) DEFAULT NULL,
  `article_id` int(11) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=53 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for zcr_article_ip
-- ----------------------------
DROP TABLE IF EXISTS `zcr_article_ip`;
CREATE TABLE `zcr_article_ip` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `article_id` int(11) DEFAULT NULL,
  `ip` varchar(255) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for zcr_article_tags
-- ----------------------------
DROP TABLE IF EXISTS `zcr_article_tags`;
CREATE TABLE `zcr_article_tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `tag_id` int(11) DEFAULT NULL,
  `article_id` int(11) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=49 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for zcr_attach
-- ----------------------------
DROP TABLE IF EXISTS `zcr_attach`;
CREATE TABLE `zcr_attach` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `path` varchar(255) DEFAULT NULL,
  `create_time` int(11) DEFAULT NULL,
  `update_time` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for zcr_category
-- ----------------------------
DROP TABLE IF EXISTS `zcr_category`;
CREATE TABLE `zcr_category` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `remark` varchar(16) DEFAULT NULL,
  `name` varchar(16) DEFAULT NULL,
  `ctype` tinyint(4) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=13 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for zcr_log
-- ----------------------------
DROP TABLE IF EXISTS `zcr_log`;
CREATE TABLE `zcr_log` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(16) DEFAULT NULL,
  `content` text,
  `created_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for zcr_tags
-- ----------------------------
DROP TABLE IF EXISTS `zcr_tags`;
CREATE TABLE `zcr_tags` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) DEFAULT NULL,
  `status` tinyint(4) DEFAULT NULL,
  `created_at` int(11) DEFAULT NULL,
  `updated_at` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Table structure for zcr_user
-- ----------------------------
DROP TABLE IF EXISTS `zcr_user`;
CREATE TABLE `zcr_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `username` varchar(16) NOT NULL,
  `salt` varchar(32) NOT NULL,
  `password` varchar(32) NOT NULL,
  `logintime` int(11) DEFAULT NULL,
  `loginip` varchar(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;
SET FOREIGN_KEY_CHECKS=1;
