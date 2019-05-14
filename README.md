Pingmesh：用于数据中心网络延迟测量和分析的大规模系统
====

---------------------------后端配置（客户端和服务端时间必须准确）---------------------------<br>
一、服务端（Centos7）<br>
1.配置步骤<br>
yum install -y golang<br>

sudo yum install mariadb-servergo

sudo systemctl start mariadb<br>
sudo systemctl enable mariadb

sudo systemctl status mariadb


mysql_secure_installation //初始化mariadb paswd:123456<br>
sudo systemctl stop firewalld.service<br>
setenforce 0<br>

2.倒入数据库表：
新建一个库为ping，然后倒入如下数据<br>
------------分隔符------------<br>
/*<br>
 Navicat Premium Data Transfer<br>

 Source Server         : localhost<br>
 Source Server Type    : MySQL<br>
 Source Server Version : 50505<br>
 Source Host           : localhost<br>
 Source Database       : ping<br>

 Target Server Type    : MySQL<br>
 Target Server Version : 50505<br>
 File Encoding         : utf-8<br>

 Date: 04/28/2019 17:56:44 PM<br>
*/<br>

SET NAMES utf8;<br>
SET FOREIGN_KEY_CHECKS = 0;<br>

-- ----------------------------<br>
--  Table structure for `fw`<br>
-- ----------------------------<br>
DROP TABLE IF EXISTS `fw`;<br>
CREATE TABLE `fw` (<br>
  `id` int(255) NOT NULL AUTO_INCREMENT,<br>
  `pro` varchar(15) DEFAULT '''NULL''',<br>
  `eip` varchar(15) DEFAULT '''NULL''',<br>
  `eport` varchar(15) DEFAULT '''NULL''',<br>
  `iip` varchar(15) DEFAULT 'NULL',<br>
  `iport` varchar(15) DEFAULT NULL,<br>
  PRIMARY KEY (`id`)<br>
) ENGINE=InnoDB AUTO_INCREMENT=624 DEFAULT CHARSET=utf8;<br>

-- ----------------------------<br>
--  Table structure for `host`<br>
-- ----------------------------<br>
DROP TABLE IF EXISTS `host`;<br>
CREATE TABLE `host` (<br>
  `host` char(100) NOT NULL,<br>
  PRIMARY KEY (`host`) USING BTREE<br>
) ENGINE=InnoDB DEFAULT CHARSET=utf8;<br>

-- ----------------------------<br>
--  Table structure for `valu`<br>
-- ----------------------------<br>
DROP TABLE IF EXISTS `valu`;<br>
CREATE TABLE `valu` (<br>
  `src` char(100) DEFAULT NULL,<br>
  `dst` char(100) DEFAULT NULL,<br>
  `loss` char(100) DEFAULT NULL,<br>
  `tss` int(100) NOT NULL,<br>
  `id` int(11) NOT NULL AUTO_INCREMENT,<br>
  `rttmin` double(11,2) DEFAULT 0.00,<br>
  `rttavg` double(11,2) DEFAULT 0.00,<br>
  `rttmax` double(11,2) DEFAULT 0.00,<br>
  PRIMARY KEY (`id`)<br>
) ENGINE=InnoDB AUTO_INCREMENT=5108 DEFAULT CHARSET=utf8;<br>

SET FOREIGN_KEY_CHECKS = 1;<br>
------------分隔符------------<br>

3.编译（pingmesh-s-v1.1-GetResult.go/pingmesh-s-v1.1-GetHostIp.go）<br>
go build pingmesh-s-v1.1-GetResult.go<br>
go build pingmesh-s-v1.1-GetHostIp.go<br>

4.运行<br>
nohup ./pingmesh-s-v1.1-GetResult > output.log 2>&1 &<br>
nohup ./pingmesh-s-v1.1-GetHostIp > output.log 2>&1 &<br>

二、客户端<br>
1.基础安装<br>
yum install epel-release -y<br>
yum install fping -y<br>

2.编译后端（pingmesh-c-v1.1.go）<br>
go build pingmesh-c-v1.1.go<br>

3.运行<br>
nohup ./pingmesh-c-v1.1 > output.log 2>&1 &<br>

---------------------------前端配置---------------------------<br>
修改参数pingmesh.py<br>
conn = pymysql.connect(<br>
    host='172.19.129.11',<br>
    user='root',<br>
    password='123456',<br>
    db='ping',<br>
    charset='utf8'<br>
)<br>
部分<br>


1.运行项目<br>
python pingmesh.py<br>

2.结果显示<br>
访问http://127.0.0.1:9000<br>
