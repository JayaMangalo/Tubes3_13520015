-- --------------------------------------------------------
-- Host:                         127.0.0.1
-- Server version:               10.7.3-MariaDB - mariadb.org binary distribution
-- Server OS:                    Win64
-- HeidiSQL Version:             11.3.0.6295
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- Dumping database structure for db_stima
CREATE DATABASE IF NOT EXISTS `db_stima` /*!40100 DEFAULT CHARACTER SET latin1 */;
USE `db_stima`;

-- Dumping structure for table db_stima.diagnosis_penyakit
CREATE TABLE IF NOT EXISTS `diagnosis_penyakit` (
  `tanggal` varchar(255) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `penyakit` varchar(255) NOT NULL,
  `diagnosis` varchar(255) DEFAULT NULL,
  `similarity` int(11) DEFAULT NULL,
  PRIMARY KEY (`tanggal`,`nama`,`penyakit`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table db_stima.diagnosis_penyakit: ~9 rows (approximately)
/*!40000 ALTER TABLE `diagnosis_penyakit` DISABLE KEYS */;
INSERT INTO `diagnosis_penyakit` (`tanggal`, `nama`, `penyakit`, `diagnosis`, `similarity`) VALUES
	('13 April 2022', 'Alan', 'Diabetes', 'True', 90),
	('13 April 2022', 'Bulan', 'Diabetes', 'True', 90),
	('13 April 2022', 'Fulan', 'Diabetes', 'True', 90),
	('26 April 2022', 'Gerald', 'HIV', 'True', 90),
	('26 April 2022', 'Jaya', 'HIV', 'True', 90),
	('27 April 2022', 'Jaya', 'HIV', 'True', 90),
	('29 April 2022', 'Aldwin', 'Flu', 'False', 26),
	('29 April 2022', 'Gerald', 'Flu', 'True', 100),
	('29 April 2022', 'Jaya', 'Covid', 'False', 23);
/*!40000 ALTER TABLE `diagnosis_penyakit` ENABLE KEYS */;

-- Dumping structure for table db_stima.penyakit
CREATE TABLE IF NOT EXISTS `penyakit` (
  `nama` varchar(255) NOT NULL,
  `sequence_dna` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`nama`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

-- Dumping data for table db_stima.penyakit: ~5 rows (approximately)
/*!40000 ALTER TABLE `penyakit` DISABLE KEYS */;
INSERT INTO `penyakit` (`nama`, `sequence_dna`) VALUES
	('Covid', 'TCGTGGTAGCGGGGCGTACACGTCCCTCCT'),
	('Flu', 'AAGGTCAACCCTGCACCGGAACC'),
	('HIV', 'CCGGTGCCCGGAATCTAGATCTGTGGCGCC'),
	('Polio', 'CGAGAATGGAGTTGGAGAACCGATGTAGAA'),
	('Sinusitis', 'TTTATCCCCGGATTCCAGCGCTGGGATAAG');
/*!40000 ALTER TABLE `penyakit` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
