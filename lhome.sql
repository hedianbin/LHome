-- MySQL dump 10.13  Distrib 5.7.17, for Win64 (x86_64)
--
-- Host: localhost    Database: lovehome
-- ------------------------------------------------------
-- Server version	5.7.17-log

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `area`
--

DROP TABLE IF EXISTS `area`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `area` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=16 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `area`
--

LOCK TABLES `area` WRITE;
/*!40000 ALTER TABLE `area` DISABLE KEYS */;
INSERT INTO `area` VALUES (1,'东城区'),(2,'西城区'),(3,'朝阳区'),(4,'海淀区'),(5,'昌平区'),(6,'丰台区'),(7,'房山区'),(8,'通州区'),(9,'顺义区'),(10,'大兴区'),(11,'怀柔区'),(12,'平谷区'),(13,'密云区'),(14,'延庆区'),(15,'石景山区');
/*!40000 ALTER TABLE `area` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `facility`
--

DROP TABLE IF EXISTS `facility`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `facility` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=26 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `facility`
--

LOCK TABLES `facility` WRITE;
/*!40000 ALTER TABLE `facility` DISABLE KEYS */;
INSERT INTO `facility` VALUES (1,'无线网络'),(2,'热水淋浴'),(3,'空调'),(4,'暖气'),(5,'允许吸烟'),(6,'饮水设备'),(7,'牙具'),(8,'香皂'),(9,'拖鞋'),(10,'手纸'),(11,'毛巾'),(12,'沐浴露、洗发露'),(13,'冰箱'),(14,'洗衣机'),(15,'电梯'),(16,'允许做饭'),(17,'允许带宠物'),(18,'允许聚会'),(19,'门禁系统'),(20,'停车位'),(21,'有线网络'),(22,'电视'),(23,'浴缸'),(24,'吃鸡'),(25,'打台球');
/*!40000 ALTER TABLE `facility` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `facility_houses`
--

DROP TABLE IF EXISTS `facility_houses`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `facility_houses` (
  `id` bigint(20) NOT NULL AUTO_INCREMENT,
  `facility_id` int(11) NOT NULL,
  `house_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=101 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `facility_houses`
--

LOCK TABLES `facility_houses` WRITE;
/*!40000 ALTER TABLE `facility_houses` DISABLE KEYS */;
INSERT INTO `facility_houses` VALUES (1,4,1),(2,5,1),(3,6,1),(4,7,1),(5,11,1),(6,12,1),(7,1,2),(8,5,2),(9,8,2),(10,11,2),(11,12,2),(12,18,2),(13,4,3),(14,6,3),(15,7,3),(16,8,3),(17,11,3),(18,14,3),(19,17,3),(20,6,4),(21,12,4),(22,17,4),(23,5,5),(24,8,5),(25,9,5),(26,14,5),(27,16,5),(28,1,6),(29,5,6),(30,8,6),(31,9,6),(32,14,6),(33,16,6),(34,20,6),(35,1,7),(36,5,7),(37,8,7),(38,12,7),(39,16,7),(40,4,8),(41,7,8),(42,10,8),(43,11,8),(44,15,8),(45,5,9),(46,6,9),(47,12,9),(48,13,9),(49,16,9),(50,17,9),(51,3,10),(52,4,10),(53,6,10),(54,9,10),(55,10,10),(56,13,10),(57,15,10),(58,18,10),(59,1,11),(60,8,11),(61,11,11),(62,12,11),(63,16,11),(64,18,11),(65,19,11),(66,1,12),(67,7,12),(68,8,12),(69,1,13),(70,5,13),(71,6,13),(72,12,13),(73,16,13),(74,5,14),(75,7,14),(76,13,14),(77,18,14),(78,3,15),(79,7,15),(80,10,15),(81,16,15),(82,2,16),(83,6,16),(84,9,16),(85,13,16),(86,14,16),(87,15,16),(88,9,17),(89,17,17),(90,3,18),(91,7,18),(92,8,18),(93,9,18),(94,12,18),(95,16,18),(96,17,18),(97,10,19),(98,11,19),(99,14,19),(100,17,19);
/*!40000 ALTER TABLE `facility_houses` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `house`
--

DROP TABLE IF EXISTS `house`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `house` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `area_id` int(11) NOT NULL,
  `title` varchar(64) NOT NULL DEFAULT '',
  `price` int(11) NOT NULL DEFAULT '0',
  `address` varchar(512) NOT NULL DEFAULT '',
  `room_count` int(11) NOT NULL DEFAULT '1',
  `acreage` int(11) NOT NULL DEFAULT '0',
  `unit` varchar(32) NOT NULL DEFAULT '',
  `capacity` int(11) NOT NULL DEFAULT '1',
  `beds` varchar(64) NOT NULL DEFAULT '',
  `deposit` int(11) NOT NULL DEFAULT '0',
  `min_days` int(11) NOT NULL DEFAULT '1',
  `max_days` int(11) NOT NULL DEFAULT '0',
  `order_count` int(11) NOT NULL DEFAULT '0',
  `index_image_url` varchar(256) NOT NULL DEFAULT '',
  `ctime` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=20 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `house`
--

LOCK TABLES `house` WRITE;
/*!40000 ALTER TABLE `house` DISABLE KEYS */;
INSERT INTO `house` VALUES (1,19,4,'111111',11100,'1111',111,11,'111',111,'111',11100,111,111,1,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-23 08:58:44'),(2,19,5,'222222',22200,'222',22,22,'22',22,'22',2200,22,22,0,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-23 13:00:23'),(3,19,8,'333333',33300,'333',333,333,'333',333,'333',33300,333,333,0,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-23 13:12:54'),(4,19,9,'444444',44400,'444',444,444,'444',444,'444',44400,444,444,0,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-23 13:25:49'),(5,19,3,'55555',66600,'666',666,666,'666',666,'666',66600,666,66,1,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-23 13:31:44'),(6,19,1,'666666',111100,'1111',1111,1111,'11111',111,'1111',11100,11,11,0,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-23 14:05:16'),(7,19,3,'777777',77700,'77777',77,77,'77',777,'77',7700,77,77,0,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-23 14:24:38'),(8,19,4,'888888',11100,'111',11,11,'11',111,'111',1100,1,11,0,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-25 00:07:10'),(9,19,2,'999999',666600,'6666',6666,6666,'6666',666,'66',66600,6,66,0,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-25 01:00:59'),(10,19,5,'aaaaaa',88800,'888',88,88,'88',88,'88',8800,88,88,0,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg','2018-06-24 17:01:53'),(11,19,3,'bbbbbb',99900,'999',99,99,'99',99,'99',9900,99,99,0,'group1\\M00/00/00/CgCX3FswP_SAFLMKAAFqdROOF2Y301.jpg','2018-06-23 17:04:59'),(12,20,1,'cccccc',22200,'2222',2,22,'22',22,'22',2200,22,22,0,'group1\\M00/00/00/CgCX3FswrQWAQi16AAFqdROOF2Y425.jpg','2018-06-24 16:50:21'),(13,19,3,'cccccc',1111100,'111',11,11,'11',11,'11',1100,11,11,0,'group1\\M00/00/00/CgCX3FsyI2-ASZmvAAGvjbpp7u4031.jpg','2018-06-26 03:27:50'),(14,19,2,'7777',777700,'7777',7777,7777,'7777',7777,'7777',777700,7777,7777,0,'','2018-07-01 12:52:24'),(15,19,2,'8888',888800,'8888',8888,8888,'8888',8888,'8888',888800,8888,8888,0,'','2018-07-01 13:03:45'),(16,19,1,'9999',999900,'9999',999,99,'99',99,'99',9900,99,99,0,'','2018-07-01 13:05:00'),(17,19,2,'3333',333300,'33',3,3,'333',3,'3',300,3,3,0,'','2018-07-01 21:15:58'),(18,19,3,'345345',5500,'55',55,55,'55',55,'55',5500,55,55,0,'group1/M00/00/00/CgCXrFtLcwWAakx-AAFqdROOF2Y274.jpg','2018-07-16 00:14:44'),(19,19,3,'人脸摇摇徐 ',5500,'55',55,55,'55',55,'55',5500,55,55,0,'group1/M00/00/00/CgCXrFtLc0iAaUG-AAFqdROOF2Y919.jpg','2018-07-16 00:16:03');
/*!40000 ALTER TABLE `house` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `house_image`
--

DROP TABLE IF EXISTS `house_image`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `house_image` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `url` varchar(256) NOT NULL DEFAULT '',
  `house_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `house_image`
--

LOCK TABLES `house_image` WRITE;
/*!40000 ALTER TABLE `house_image` DISABLE KEYS */;
INSERT INTO `house_image` VALUES (1,'group1\\M00/00/00/CgCX3FswPzmAWKKQAAFqdROOF2Y013.jpg',10),(2,'group1\\M00/00/00/CgCX3FswP_SAFLMKAAFqdROOF2Y301.jpg',11),(3,'group1\\M00/00/00/CgCX3FswP_iAOIRLAAEJTLCfWmk305.jpg',11),(4,'group1\\M00/00/00/CgCX3FswP_yAW7ckAAGvjbpp7u4713.jpg',11),(5,'group1\\M00/00/00/CgCX3FswQACALpvLAAGK0WoGz_Y594.jpg',11),(6,'group1\\M00/00/00/CgCX3FswrQWAQi16AAFqdROOF2Y425.jpg',12),(7,'group1\\M00/00/00/CgCX3FswrQmAN_nVAAGvjbpp7u4299.jpg',12),(8,'group1\\M00/00/00/CgCX3FsyI2-ASZmvAAGvjbpp7u4031.jpg',13),(9,'group1/M00/00/00/CgCXrFtLcwWAakx-AAFqdROOF2Y274.jpg',18),(10,'group1/M00/00/00/CgCXrFtLc0iAaUG-AAFqdROOF2Y919.jpg',19);
/*!40000 ALTER TABLE `house_image` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `order_house`
--

DROP TABLE IF EXISTS `order_house`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `order_house` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `house_id` int(11) NOT NULL,
  `begin_date` datetime NOT NULL,
  `end_date` datetime NOT NULL,
  `days` int(11) NOT NULL DEFAULT '0',
  `house_price` int(11) NOT NULL DEFAULT '0',
  `amount` int(11) NOT NULL DEFAULT '0',
  `status` varchar(255) NOT NULL DEFAULT 'WAIT_ACCEPT',
  `comment` varchar(512) NOT NULL DEFAULT '',
  `ctime` datetime NOT NULL,
  `begin_data` datetime NOT NULL,
  `end_data` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `order_house`
--

LOCK TABLES `order_house` WRITE;
/*!40000 ALTER TABLE `order_house` DISABLE KEYS */;
INSERT INTO `order_house` VALUES (1,1,1,'2018-06-26 16:00:00','2018-06-28 16:00:00',3,11100,33300,'WAIT_COMMENT','','2018-06-25 05:45:33','0000-00-00 00:00:00','0000-00-00 00:00:00'),(2,1,1,'2018-06-26 16:00:00','2018-06-26 16:00:00',1,11100,11100,'WAIT_COMMENT','','2018-06-25 05:47:49','0000-00-00 00:00:00','0000-00-00 00:00:00'),(3,1,1,'2018-06-28 16:00:00','2018-06-28 16:00:00',1,11100,11100,'REJECTED','不接，哈哈','2018-06-25 16:50:09','0000-00-00 00:00:00','0000-00-00 00:00:00'),(4,1,1,'2018-06-28 16:00:00','2018-06-29 16:00:00',2,11100,22200,'COMPLETE','太脏','2018-06-25 17:43:21','0000-00-00 00:00:00','0000-00-00 00:00:00'),(5,1,3,'2018-06-29 16:00:00','2018-06-29 16:00:00',1,33300,33300,'REJECTED','无房','2018-06-25 22:07:05','0000-00-00 00:00:00','0000-00-00 00:00:00'),(6,1,5,'2018-06-27 16:00:00','2018-06-27 16:00:00',1,66600,66600,'COMPLETE','还不错','2018-06-25 22:07:50','0000-00-00 00:00:00','0000-00-00 00:00:00'),(7,1,12,'2018-06-28 00:00:00','2018-06-28 00:00:00',1,22200,22200,'WAIT_ACCEPT','','2018-06-26 11:32:04','0000-00-00 00:00:00','0000-00-00 00:00:00'),(8,1,1,'2018-06-29 00:00:00','2018-06-29 00:00:00',1,11100,11100,'WAIT_ACCEPT','','2018-06-26 11:34:28','0000-00-00 00:00:00','0000-00-00 00:00:00'),(9,1,1,'2018-07-21 08:00:00','2018-07-21 08:00:00',1,11100,11100,'WAIT_ACCEPT','','2018-07-04 17:10:26','0000-00-00 00:00:00','0000-00-00 00:00:00');
/*!40000 ALTER TABLE `order_house` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL DEFAULT '',
  `password_hash` varchar(128) NOT NULL DEFAULT '',
  `mobile` varchar(11) NOT NULL DEFAULT '',
  `real_name` varchar(32) NOT NULL DEFAULT '',
  `id_card` varchar(20) NOT NULL DEFAULT '',
  `avatar_url` varchar(256) NOT NULL DEFAULT '',
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
INSERT INTO `user` VALUES (1,'何殿斌','111','111','何殿斌','222424112345675433','group1\\M00/00/00/CgCX3FstB1iAU3OcAAFqdROOF2Y384.jpg'),(2,'4444','4444','4444','何殿斌','222424112345675433',''),(3,'666','666','666','ttgttwe','52356345433456','group1\\M00/00/00/CgCX3FstGw2AKrEIAAGvjbpp7u4107.jpg'),(5,'12345678910','123','12345678910','','',''),(6,'7777','7777','7777','','',''),(7,'8888','888','8888','','','group1\\M00/00/00/CgCX3FssmyeACiLZAAGK0WoGz_Y904.jpg'),(8,'1111','1111','1111','','','group1\\M00/00/00/CgCX3FssmIyAA11dAAFqdROOF2Y380.jpg'),(9,'001','001','001','','',''),(10,'002','002','002','','',''),(11,'003','003','003','','',''),(12,'22222222222','222','22222222222','','','group1\\M00/00/00/CgCX3FssZweAZJd0AAFqdROOF2Y96..jpg'),(15,'0002','333','0002','','',''),(16,'0018','333','0018','','',''),(17,'1212','1212','1212','','','group1/M00/00/00/CgCX3Fsssv-ARkblAAGK0WoGz_Y128.jpg'),(18,'2323','2323','2323','','','group1/M00/00/00/CgCX3Fsstg6AHf2vAAFqdROOF2Y913.jpg'),(19,'xiaonu','222','222','fdgf','4343262345236','group1/M00/00/00/CgCXrFtL8VuAIlKBAAFqdROOF2Y613.jpg'),(20,'333','333','333','333','333','group1\\M00/00/00/CgCX3FswpG6AI-flAAFqdROOF2Y534.jpg');
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2018-07-16  9:36:15
