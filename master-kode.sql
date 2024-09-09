-- MySQL dump 10.13  Distrib 8.0.33, for Win64 (x86_64)
--
-- Host: 127.0.0.1    Database: db_wkm
-- ------------------------------------------------------
-- Server version	8.0.33

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!50503 SET NAMES utf8 */;
/*!40103 SET @OLD_TIME_ZONE=@@TIME_ZONE */;
/*!40103 SET TIME_ZONE='+00:00' */;
/*!40014 SET @OLD_UNIQUE_CHECKS=@@UNIQUE_CHECKS, UNIQUE_CHECKS=0 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;

--
-- Table structure for table `mst_agama`
--

DROP TABLE IF EXISTS `mst_agama`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mst_agama` (
  `kd_agama` varchar(2) NOT NULL,
  `agama` varchar(10) DEFAULT NULL,
  PRIMARY KEY (`kd_agama`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_agama`
--

LOCK TABLES `mst_agama` WRITE;
/*!40000 ALTER TABLE `mst_agama` DISABLE KEYS */;
INSERT INTO `mst_agama` VALUES ('1','Islam'),('2','Kristen'),('3','Katolik'),('4','Hindu'),('5','Budha'),('6','Lainnya'),('7','Konghucu');
/*!40000 ALTER TABLE `mst_agama` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mst_jns_beli`
--

DROP TABLE IF EXISTS `mst_jns_beli`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mst_jns_beli` (
  `kd_beli` varchar(10) NOT NULL,
  `nm_beli` varchar(50) NOT NULL,
  PRIMARY KEY (`kd_beli`)
) ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_jns_beli`
--

LOCK TABLES `mst_jns_beli` WRITE;
/*!40000 ALTER TABLE `mst_jns_beli` DISABLE KEYS */;
INSERT INTO `mst_jns_beli` VALUES ('C','COLLECTIVE'),('G','GROUP CUSTOMER'),('I','INDIVIDUAL'),('J','JOIN PROMO');
/*!40000 ALTER TABLE `mst_jns_beli` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mst_jns_jual`
--

DROP TABLE IF EXISTS `mst_jns_jual`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mst_jns_jual` (
  `kd_jual` varchar(10) NOT NULL,
  `nm_jual` varchar(50) NOT NULL,
  PRIMARY KEY (`kd_jual`)
) ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_jns_jual`
--

LOCK TABLES `mst_jns_jual` WRITE;
/*!40000 ALTER TABLE `mst_jns_jual` DISABLE KEYS */;
INSERT INTO `mst_jns_jual` VALUES ('1','Tunai'),('2','Kredit');
/*!40000 ALTER TABLE `mst_jns_jual` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mst_keluar_bln`
--

DROP TABLE IF EXISTS `mst_keluar_bln`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mst_keluar_bln` (
  `id` int unsigned NOT NULL,
  `keluar_bln2` varchar(2) NOT NULL,
  `nm_keluar_bln2` varchar(45) NOT NULL,
  `ket_keluar_bln2` varchar(45) NOT NULL,
  `sts_keluar_bln2` varchar(1) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_keluar_bln`
--

LOCK TABLES `mst_keluar_bln` WRITE;
/*!40000 ALTER TABLE `mst_keluar_bln` DISABLE KEYS */;
INSERT INTO `mst_keluar_bln` VALUES (0,'','','','1'),(1,'01','<700rb','','0'),(2,'02','1jt-1,5jt','','0'),(3,'03','1,5jt-2jt','','0'),(4,'04','2jt-3jt','','0'),(5,'05','3jt-4jt','','0'),(6,'06','>4jt','','0'),(7,'E','<RP 1.010.000','','0'),(8,'D','Rp 1.010.001 - Rp. 1.550.000','','0'),(9,'C2','Rp 1.550.001 - Rp. 2.2260.000','','0'),(10,'C1','Rp 2.260.001 - Rp. 3.220.000','','0'),(11,'B','Rp 3.220.001 - Rp. 5.170.001','','0'),(12,'A2','Rp 5.170.001 - Rp. 8.700.000','','0'),(13,'A1','>Rp 8.700.000','','0'),(14,'1','<RP 900.000','','1'),(15,'2','RP 900.001 - RP 1.250.000','','1'),(16,'3','RP 1.250.001 - RP 1.750.000','','1'),(17,'4','RP 1.750.001 - RP 2.500.000','','1'),(18,'5','RP 2.500.001 - RP 4.000.000','','1'),(19,'6','RP 4.000.001 - RP 6.000.000','','1'),(20,'7','>RP 6.000.000','','1');
/*!40000 ALTER TABLE `mst_keluar_bln` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mst_leasing`
--

DROP TABLE IF EXISTS `mst_leasing`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mst_leasing` (
  `no_leas2` varchar(2) NOT NULL,
  `nm_leasing` varchar(50) DEFAULT NULL,
  `nm_wil` varchar(20) DEFAULT NULL,
  `alamat_leasing` varchar(100) DEFAULT NULL,
  `kodepos` varchar(5) DEFAULT NULL,
  `email` varchar(100) DEFAULT NULL,
  `telp` varchar(30) DEFAULT NULL,
  `fax` varchar(30) DEFAULT NULL,
  `owner` varchar(30) DEFAULT NULL,
  `kd_leas` varchar(2) DEFAULT NULL,
  PRIMARY KEY (`no_leas2`) USING BTREE,
  KEY `index_2` (`no_leas2`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=latin1;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_leasing`
--

LOCK TABLES `mst_leasing` WRITE;
/*!40000 ALTER TABLE `mst_leasing` DISABLE KEYS */;
INSERT INTO `mst_leasing` VALUES ('','','','','','','','','','56'),('01','MD - DLR FINCOY','','','','','','','','99'),('02','MD - DLR FUND','','','','','','','','99'),('03','NSS - NUSANTARA SURYA SAKTI','','','','','','','','16'),('04','FIF - FEDERAL INTERNATIONAL FINANCE','','','','','','','','11'),('05','WOM - WAHANA OTTOMITRA MULTIARTHA','','','','','','','','15'),('06','ADIRA - ADIRA DINAMIKA MULTI FINANCE','','','','','','','','12'),('07','SOF (OTO) - SUMMIT OTO FINANCE','','','','','','','','14'),('08','SAF - SASANA ARTHA FINANCE','','','','','','','','99'),('09','TUNAS','','','','','','','','99'),('10','PARA','','','','','','','','99'),('11','HD - RADANA','','','','','','','','99'),('12','ITC - INTERNUSA TRIBUANA CITRA    ','','','','','','','','99'),('13','CMD - CAPELA DINAMIKA NUSANTARA','','','','','','','','99'),('14','STAR','','','','','','','','99'),('15','MCF - MEGA CENTRAL FINANCE PT','','','','','','','','13'),('16','CSF - CENTRAL SANTOSA FINANCE','','','','','','','','99'),('17','OTHERS','','','','','','','','51'),('18','NUSA SURYA CIPTADANA PT','','','','','','','','99'),('19','MTF - MANDIRI TUNAS FINANCE PT','','','','','','','','99'),('20','MEGA FINANCE PT','','','','','','','','18'),('21','INDOMOBIL FINANCE INDONESIA PT','','','','','','','','99'),('22','AL IJARAH INDONESIA FINANCE PT','','','','','','','','99'),('23','CIMB NIAGA AUTO FINANCE','','','','','','','','99'),('24','BHAKTI FINANCE PT','','','','','','','','99'),('25','BNI MULTI FINANCE','','','','','','','','99'),('26','BPR - KARYAJATNIKA SADAYA PT','','','','','','','','99'),('27','BPR - UNIVERSAL SENTOSA PT','','','','','','','','99'),('28','MULTINDO AUTO FINANCE PT','','','','','','','','99'),('29','SINAR MITRA SEPADAN FINANCE PT','','','','','','','','99'),('30','MANDALA MULTI FINANCE TBK PT','','','','','','','','99'),('31','FINANCIA MULTI FINANCE PT','','','','','','','','99'),('32','BPR - BANK PERKREDITAN RAKYAT SE','','','','','','','','99'),('33','BETA INTI MULTI FINANCE PT','','','','','','','','99'),('34','SEJAHTERA PERTAMA MULTIFINANCE PT','','','','','','','','99'),('35','SUNINDO PARAMA FINANCE','','','','','','','','99'),('36','FIF SYARIAH','','','','','','','','99'),('37','WOM SYARIAH','','','','','','','','99'),('38','MANDIRI UTAMA FINANCE','','','','','','','','17'),('39','TRANSPACIFIC FINANCE PT','','','','','','','','99'),('N','N','','','','','','','','55'),('X1','BELUM DIVERIFIKASI','','','','','','','','53'),('X2','KONSUMEN TIDAK BERSEDIA','','','','','','','','54'),('XX','TIDAK TAHU','','','','','','','','52');
/*!40000 ALTER TABLE `mst_leasing` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mst_pendidikan`
--

DROP TABLE IF EXISTS `mst_pendidikan`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mst_pendidikan` (
  `kd_pendidikan` varchar(10) NOT NULL,
  `nm_pendidikan` varchar(100) NOT NULL,
  PRIMARY KEY (`kd_pendidikan`)
) ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_pendidikan`
--

LOCK TABLES `mst_pendidikan` WRITE;
/*!40000 ALTER TABLE `mst_pendidikan` DISABLE KEYS */;
INSERT INTO `mst_pendidikan` VALUES ('1','TIDAK TAMAT SD'),('2','SD'),('3','SLTP/SMP'),('4','SLTA/SMA'),('5','AKADEMI/DIPLOMA'),('6','SARJANA'),('7','PASCA SARJANA');
/*!40000 ALTER TABLE `mst_pendidikan` ENABLE KEYS */;
UNLOCK TABLES;

--
-- Table structure for table `mst_pengeluaran`
--

DROP TABLE IF EXISTS `mst_pengeluaran`;
/*!40101 SET @saved_cs_client     = @@character_set_client */;
/*!50503 SET character_set_client = utf8mb4 */;
CREATE TABLE `mst_pengeluaran` (
  `kd_pengeluaran` varchar(10) NOT NULL,
  `pengeluaran` varchar(100) NOT NULL,
  PRIMARY KEY (`kd_pengeluaran`)
) ;
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `mst_pengeluaran`
--

LOCK TABLES `mst_pengeluaran` WRITE;
/*!40000 ALTER TABLE `mst_pengeluaran` DISABLE KEYS */;
INSERT INTO `mst_pengeluaran` VALUES ('1','< Rp 900,000'),('2','Rp 900,001 - Rp 1,250,000'),('3','Rp 1,250,001 - Rp 1,750,000'),('4','Rp 1,750,001 - Rp 2,500,000'),('5','Rp 2,500,001 - Rp 4,000,000'),('6','Rp 4,000,001 - Rp 6,000,000'),('7','> Rp 6,000,000');
/*!40000 ALTER TABLE `mst_pengeluaran` ENABLE KEYS */;
UNLOCK TABLES;
/*!40103 SET TIME_ZONE=@OLD_TIME_ZONE */;

/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40014 SET UNIQUE_CHECKS=@OLD_UNIQUE_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;

-- Dump completed on 2024-09-09 15:50:50
