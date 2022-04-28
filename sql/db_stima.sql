-- MariaDB dump 10.19  Distrib 10.7.3-MariaDB, for osx10.17 (arm64)
--
-- Host: localhost    Database: test
-- ------------------------------------------------------
-- Server version	8.0.28

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
-- Current Database: `test`
--

CREATE DATABASE /*!32312 IF NOT EXISTS*/ `test` /*!40100 DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci */ /*!80016 DEFAULT ENCRYPTION='N' */;

USE `test`;

--
-- Table structure for table `diagnosis_penyakit`
--

DROP TABLE IF EXISTS `diagnosis_penyakit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `diagnosis_penyakit` (
  `tanggal` varchar(255) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `penyakit` varchar(255) NOT NULL,
  `diagnosis` varchar(255) DEFAULT NULL,
  `similarity` int DEFAULT NULL,
  PRIMARY KEY (`tanggal`,`nama`,`penyakit`)
) ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `diagnosis_penyakit`
--

LOCK TABLES `diagnosis_penyakit` WRITE;
/*!40000 ALTER TABLE `diagnosis_penyakit` DISABLE KEYS */;
INSERT INTO `diagnosis_penyakit` (`tanggal`, `nama`, `penyakit`, `diagnosis`, `similarity`) VALUES ('13 April 2022','Alan','Diabetes','True',50),
('13 April 2022','Bulan','Diabetes','True',50),
('13 April 2022','Fulan','Diabetes','True',50),
('26 April 2022','Gerald','HIV','True',50),
('26 April 2022','Jaya','HIV','True',50),
('27 April 2022','Aldwin','HIV','True',96),
('27 April 2022','Gerald','HIV','True',90),
('27 April 2022','Jaya','HIV','True',50);
/*!40000 ALTER TABLE `diagnosis_penyakit` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `penyakit`
--

DROP TABLE IF EXISTS `penyakit`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!40101 SET character_set_client = utf8 */;
CREATE TABLE `penyakit` (
  `nama` varchar(255) NOT NULL,
  `sequence_dna` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`nama`)
) ;
/*!40101 SET character_set_client = @saved_cs_client */;

--test
-- Dumping data for table `penyakit`
--

LOCK TABLES `penyakit` WRITE;
/*!40000 ALTER TABLE `penyakit` DISABLE KEYS */;
INSERT INTO `penyakit` (`nama`, `sequence_dna`) VALUES ('HIV','CCGGTGCCCGGAATCTAGATCTGTGGCGCC'),
('Penyakit-1','AAAABBBBCCCCDDDD'),
('Penyakit-2','ABCABABDFFFDDDD'),
('Penyakit-3','ABABABABABABABAB'),
('Penyakit-4','ABCDABCDABCDABCD');
/*!40000 ALTER TABLE `penyakit` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2022-04-28  0:17:52
