-- --------------------------------------------------------
-- 호스트:                          127.0.0.1
-- 서버 버전:                        10.6.0-MariaDB - mariadb.org binary distribution
-- 서버 OS:                        Win64
-- HeidiSQL 버전:                  11.2.0.6213
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- hoseo 데이터베이스 구조 내보내기
CREATE DATABASE IF NOT EXISTS `hoseo` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `hoseo`;

-- 테이블 hoseo.wearable_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `wearable_info` (
  `wearable_SN` varchar(45) NOT NULL,
  UNIQUE KEY `wearable_SN` (`wearable_SN`),
  CONSTRAINT `FK1_wearableSN` FOREIGN KEY (`wearable_SN`) REFERENCES `user_info` (`wearable_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.wearable_info:~10 rows (대략적) 내보내기
DELETE FROM `wearable_info`;
/*!40000 ALTER TABLE `wearable_info` DISABLE KEYS */;
INSERT INTO `wearable_info` (`wearable_SN`) VALUES
	('wsn1111'),
	('wsn1112'),
	('wsn1113'),
	('wsn1114'),
	('wsn1115'),
	('wsn1116'),
	('wsn1117'),
	('wsn1118'),
	('wsn1119'),
	('wsn1120');
/*!40000 ALTER TABLE `wearable_info` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
