/*
 Navicat Premium Data Transfer

 Source Server         : 172.19.129.11
 Source Server Type    : MariaDB
 Source Server Version : 50560
 Source Host           : 172.19.129.11
 Source Database       : ping

 Target Server Type    : MariaDB
 Target Server Version : 50560
 File Encoding         : utf-8

 Date: 05/14/2019 11:18:29 AM
*/

SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
--  Table structure for `fw`
-- ----------------------------
DROP TABLE IF EXISTS `fw`;
CREATE TABLE `fw` (
  `id` int(255) NOT NULL AUTO_INCREMENT,
  `pro` varchar(15) DEFAULT '''NULL''',
  `eip` varchar(15) DEFAULT '''NULL''',
  `eport` varchar(15) DEFAULT '''NULL''',
  `iip` varchar(15) DEFAULT 'NULL',
  `iport` varchar(15) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=624 DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `host`
-- ----------------------------
DROP TABLE IF EXISTS `host`;
CREATE TABLE `host` (
  `host` char(100) NOT NULL,
  PRIMARY KEY (`host`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- ----------------------------
--  Table structure for `valu`
-- ----------------------------
DROP TABLE IF EXISTS `valu`;
CREATE TABLE `valu` (
  `src` char(100) DEFAULT NULL,
  `dst` char(100) DEFAULT NULL,
  `loss` char(100) DEFAULT NULL,
  `tss` int(100) NOT NULL,
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `rttmin` double(11,2) DEFAULT '0.00',
  `rttavg` double(11,2) DEFAULT '0.00',
  `rttmax` double(11,2) DEFAULT '0.00',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=178150 DEFAULT CHARSET=utf8;

SET FOREIGN_KEY_CHECKS = 1;
