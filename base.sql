-- MySQL dump 10.13  Distrib 5.5.40, for Win64 (x86)
--
-- Host: localhost    Database: temp
-- ------------------------------------------------------
-- Server version	5.5.40

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
-- Table structure for table `comment`
--

DROP TABLE IF EXISTS `comment`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `comment` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT '评论用户id',
  `video_id` int(10) unsigned NOT NULL COMMENT '被评论视频id',
  `content` varchar(255) NOT NULL COMMENT '评论内容',
  `create_date` varchar(5) NOT NULL COMMENT '评论发布日期，格式 mm-dd',
  PRIMARY KEY (`id`),
  KEY `comment_user_id` (`user_id`),
  KEY `comment_video_id` (`video_id`),
  CONSTRAINT `comment_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `comment_video_id` FOREIGN KEY (`video_id`) REFERENCES `video` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `comment`
--

LOCK TABLES `comment` WRITE;
/*!40000 ALTER TABLE `comment` DISABLE KEYS */;
/*!40000 ALTER TABLE `comment` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `message`
--

DROP TABLE IF EXISTS `message`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `message` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `to_user_id` int(10) unsigned NOT NULL COMMENT '该消息接收者的id',
  `from_user_id` int(10) unsigned NOT NULL COMMENT '该消息发送者的id',
  `content` varchar(255) NOT NULL COMMENT '消息内容',
  `create_time` datetime NOT NULL COMMENT '消息创建时间',
  PRIMARY KEY (`id`),
  KEY `message_to_user_id` (`to_user_id`),
  KEY `message_from_user_id` (`from_user_id`),
  CONSTRAINT `message_to_user_id` FOREIGN KEY (`to_user_id`) REFERENCES `user` (`id`),
  CONSTRAINT `message_from_user_id` FOREIGN KEY (`from_user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `message`
--

LOCK TABLES `message` WRITE;
/*!40000 ALTER TABLE `message` DISABLE KEYS */;
/*!40000 ALTER TABLE `message` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `user`
--

DROP TABLE IF EXISTS `user`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `user` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(32) NOT NULL COMMENT '用户名',
  `password` varchar(32) NOT NULL COMMENT '密码',
  `follow_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '关注',
  `follower_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '粉丝',
  `avatar` varchar(1000) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '头像',
  `background-image` varchar(1000) CHARACTER SET utf8 COLLATE utf8_bin DEFAULT NULL COMMENT '主页背景',
  `signature` varchar(255) DEFAULT NULL COMMENT '个人简介',
  `total_favorited` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '获赞数',
  `work_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '作品数',
  `favorite_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `id` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `user`
--

LOCK TABLES `user` WRITE;
/*!40000 ALTER TABLE `user` DISABLE KEYS */;
/*!40000 ALTER TABLE `user` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `video`
--

DROP TABLE IF EXISTS `video`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `video` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` int(10) unsigned NOT NULL COMMENT '发布用户',
  `play_url` varchar(255) NOT NULL COMMENT '播放url',
  `cover_url` varchar(255) NOT NULL COMMENT '封面url',
  `favorite_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点赞数',
  `comment_count` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '评论数',
  `title` varchar(255) NOT NULL COMMENT '标题',
  `create_time` datetime NOT NULL COMMENT '上传时间',
  PRIMARY KEY (`id`),
  KEY `video_user_id` (`user_id`),
  CONSTRAINT `video_user_id` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `video`
--

LOCK TABLES `video` WRITE;
/*!40000 ALTER TABLE `video` DISABLE KEYS */;
/*!40000 ALTER TABLE `video` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2023-08-05 20:47:46
