-- MySQL dump 10.16  Distrib 10.1.23-MariaDB, for debian-linux-gnueabihf (armv7l)
--
-- Host: localhost    Database: furnir
-- ------------------------------------------------------
-- Server version	10.1.23-MariaDB-9+deb9u1

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `essences`
--

DROP TABLE IF EXISTS `essences`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `essences` (
  `essences_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `code` varchar(20) NOT NULL,
  PRIMARY KEY (`essences_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `essences`
--

LOCK TABLES `essences` WRITE;
/*!40000 ALTER TABLE `essences` DISABLE KEYS */;
/*!40000 ALTER TABLE `essences` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `movements`
--

DROP TABLE IF EXISTS `movements`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `movements` (
  `movements_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `source_pallets_id` int(10) unsigned NOT NULL,
  `dest_pallets_id` int(10) unsigned NOT NULL,
  `executed` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `pieces_id` int(10) unsigned NOT NULL,
  `existed` smallint(6) DEFAULT NULL,
  PRIMARY KEY (`movements_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `movements`
--

LOCK TABLES `movements` WRITE;
/*!40000 ALTER TABLE `movements` DISABLE KEYS */;
/*!40000 ALTER TABLE `movements` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pallets`
--

DROP TABLE IF EXISTS `pallets`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pallets` (
  `pallets_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`pallets_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

alter table pallets add column essences_id integer unsigned;
alter table pallets add foreign key (essences_id)  references essences(essences_id);

--
-- Dumping data for table `pallets`
--

LOCK TABLES `pallets` WRITE;
/*!40000 ALTER TABLE `pallets` DISABLE KEYS */;
/*!40000 ALTER TABLE `pallets` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `pieces`
--

DROP TABLE IF EXISTS `pieces`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `pieces` (
  `pieces_id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `pallets_id` int(10) unsigned NOT NULL,
  `created` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `barcode` varchar(45) NOT NULL,
  `essences_id` int(10) unsigned NOT NULL,
  `length` int(11) DEFAULT NULL,
  `width` int(11) DEFAULT NULL,
  `foi` int(11) DEFAULT NULL,
  PRIMARY KEY (`pieces_id`),
  KEY `idx_pallets_id` (`pallets_id`),
  KEY `essences_id` (`essences_id`),
  CONSTRAINT `pieces_ibfk_1` FOREIGN KEY (`pallets_id`) REFERENCES `pallets` (`pallets_id`),
  CONSTRAINT `pieces_ibfk_2` FOREIGN KEY (`essences_id`) REFERENCES `essences` (`essences_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `pieces`
--

LOCK TABLES `pieces` WRITE;
/*!40000 ALTER TABLE `pieces` DISABLE KEYS */;
/*!40000 ALTER TABLE `pieces` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2017-09-24 12:29:14
