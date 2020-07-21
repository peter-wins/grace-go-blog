/*
Navicat MySQL Data Transfer

Source Server         : oa
Source Server Version : 50529
Source Host           : localhost:2002
Source Database       : iamzcr

Target Server Type    : MYSQL
Target Server Version : 50529
File Encoding         : 65001

Date: 2020-07-21 22:02:20
*/

SET FOREIGN_KEY_CHECKS=0;

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

-- ----------------------------
-- Records of zcr_user
-- ----------------------------
INSERT INTO `zcr_user` VALUES ('1', 'zcr', '394868456436dbe743e4380554c0493a', 'bba20d36c512b2745e67bb056d70d50b', '1595340128', '192.168.93.1');
SET FOREIGN_KEY_CHECKS=1;
