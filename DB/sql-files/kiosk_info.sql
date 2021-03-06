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

-- 테이블 hoseo.kiosk_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `kiosk_info` (
  `kiosk_SN` varchar(45) NOT NULL COMMENT '키오스크 시리얼넘버',
  `kiosk_log` datetime DEFAULT NULL COMMENT '키오스크 로그',
  UNIQUE KEY `kiosk_SN` (`kiosk_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='Kiosk 정보';

-- 테이블 데이터 hoseo.kiosk_info:~11 rows (대략적) 내보내기
/*!40000 ALTER TABLE `kiosk_info` DISABLE KEYS */;
INSERT IGNORE INTO `kiosk_info` (`kiosk_SN`, `kiosk_log`) VALUES
	('KSN1111', NULL),
	('KSN1112', NULL),
	('KSN1113', NULL),
	('KSN1114', NULL),
	('KSN1115', NULL),
	('KSN1116', NULL),
	('KSN1117', NULL),
	('KSN1118', NULL),
	('KSN1119', NULL),
	('KSN1120', NULL),
	('ksn1121', '2021-08-17 05:00:00');
/*!40000 ALTER TABLE `kiosk_info` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
