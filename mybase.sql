-- MariaDB dump 10.19  Distrib 10.5.11-MariaDB, for Linux (x86_64)
--
-- Host: localhost    Database: mybase
-- ------------------------------------------------------
-- Server version	10.5.11-MariaDB

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
-- Table structure for table `mybase`
--

DROP TABLE IF EXISTS `mybase`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `mybase` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `partnum` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `qty` int(11) NOT NULL,
  `price` float NOT NULL,
  `provider` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `name` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `remark` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `date` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  `color` varchar(30) COLLATE utf8mb4_unicode_ci NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=41 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mybase`
--

LOCK TABLES `mybase` WRITE;
/*!40000 ALTER TABLE `mybase` DISABLE KEYS */;
INSERT INTO `mybase` VALUES (13,'500558',1,621.35,'froza','ролик ремня','Vlad','02 July 2021','#ffffff'),(14,'xvr-81-38536-00',1,353.44,'mikado','сальник','Vlad','03 July 2021','#ffffff'),(15,'1131700710',1,189.69,'froza','ремкомплект ','Anton','09 July 2021','#ffffff'),(16,'444083	',2,1326,'автоскай','амортизатор','Anton','09 July 2021','#ffffff'),(17,'99xy19-h',1,888.93,'ivers','ремень ГРМ','Pavel','10 July 2021','#ffffff'),(18,'gmi1145a078',3,1433.15,'микадо','ролик','Pavel','10 July 2021','#ffffff'),(19,'jda3017',1,242.79,'склад','фильтр масляный','Vlad','10 July 2021','#ffffff'),(20,'c1012',1,562,'склад','ф масляный','Vlad','10 July 2021','#ffffff'),(21,'jdacx036c',1,225.31,'tiss','фильтр салона','Pavel','10 July 2021','#ffffff'),(22,'46950',2,1036.08,'froza','пружина','Anton','12 July 2021','#ffffff'),(23,'st-mb93-940-a2			',1,4035,'автотрейд','зеркало','pavel','12 July 2021','#ffffff'),(24,'st-mb93-940-a1			',1,4035,'автотрейд','зеркало','pavel','12 July 2021','#ffffff'),(25,'gt37304	',1,904,'автоскай','ролик','Pavel','12 July 2021','#ffffff'),(26,'ks-4288',1,3435,'mikado','сцепление','Pavel','12 July 2021','#ffffff'),(27,'1145a019',1,3067.14,'froza','ремень ГРМ','Pavel','13 July 2021','#ffffff'),(28,'1345a062',1,4670.28,'froza','ролик','Pavel','13 July 2021','#ffffff'),(29,'mn179270',1,1731.72,'froza','опора антенны','Pavel','13 July 2021','#ffffff'),(30,'mr958030',1,2540.82,'froza','ролик','Pavel','13 July 2021','#ffffff'),(31,'wpm068v',1,3577.72,'froza','помпа','Pavel','13 July 2021','#ffffff'),(32,'1230a186',1,853.74,'froza','фильтр масляный','Pavel','13 July 2021','#ffffff'),(33,'1770a053',1,2007.87,'froza','топливный фильтр','Pavel','13 July 2021','#ffffff'),(34,'5370b423',1,1708.53,'froza','Подкрылок','Pavel','13 July 2021','#ffffff'),(35,'brg422',1,1610.51,'froza','выжимной','Pavel','13 July 2021','#ffffff'),(36,'mbc645',1,4701.09,'froza','корзина сцепления','Pavel','13 July 2021','#ffffff'),(37,'mn157914',1,5014.93,'froza','фонарь','Pavel','13 July 2021','#ffffff'),(38,'8355a016',1,823.17,'froza','отражатель','Pavel','13 July 2021','#ffffff'),(39,'8723a010',1,1205.78,'froza','антенна','Pavel','13 July 2021','#ffffff'),(40,'444132',2,1512,'Шате-М','амортизатор','Anton Pavel','13 July 2021','#ffffff');
/*!40000 ALTER TABLE `mybase` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2021-07-13 15:25:05
