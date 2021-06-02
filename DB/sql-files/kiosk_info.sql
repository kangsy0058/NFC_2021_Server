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

-- 테이블 hoseo.kiosk_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `kiosk_info` (
  `kiosk_SN` varchar(45) NOT NULL,
  UNIQUE KEY `kiosk_SN` (`kiosk_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.kiosk_info:~0 rows (대략적) 내보내기
DELETE FROM `kiosk_info`;
/*!40000 ALTER TABLE `kiosk_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `kiosk_info` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
