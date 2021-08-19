-- --------------------------------------------------------
-- 호스트:                          210.119.104.207
-- 서버 버전:                        10.6.2-MariaDB-1:10.6.2+maria~focal - mariadb.org binary distribution
-- 서버 OS:                        debian-linux-gnu
-- HeidiSQL 버전:                  11.1.0.6116
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- hoseo 데이터베이스 구조 내보내기
CREATE DATABASE IF NOT EXISTS `hoseo` /*!40100 DEFAULT CHARACTER SET utf8mb4 */;
USE `hoseo`;

-- 테이블 hoseo.wearable_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `wearable_info` (
  `wearable_SN` varchar(45) NOT NULL,
  `wearable_log` datetime DEFAULT NULL,
  PRIMARY KEY (`wearable_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='웨어러블 정보';

-- 테이블 데이터 hoseo.wearable_info:~12 rows (대략적) 내보내기
/*!40000 ALTER TABLE `wearable_info` DISABLE KEYS */;
INSERT IGNORE INTO `wearable_info` (`wearable_SN`, `wearable_log`) VALUES
	('wsn1111', NULL),
	('wsn1112', NULL),
	('wsn1113', NULL),
	('wsn1114', NULL),
	('wsn1115', NULL),
	('wsn1116', NULL),
	('wsn1117', NULL),
	('wsn1118', NULL),
	('wsn1119', NULL),
	('wsn1120', NULL),
	('wsn1121', NULL),
	('wsn1122', NULL);
/*!40000 ALTER TABLE `wearable_info` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
