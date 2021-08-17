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

-- 테이블 hoseo.push_map 구조 내보내기
CREATE TABLE IF NOT EXISTS `push_map` (
  `ch_id` int(11) NOT NULL DEFAULT 0,
  `UUID` varchar(45) NOT NULL,
  KEY `FK1_chid` (`ch_id`),
  KEY `FK2_chUUID` (`UUID`),
  CONSTRAINT `FK1_chid` FOREIGN KEY (`ch_id`) REFERENCES `push_info` (`ch_id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK2_chUUID` FOREIGN KEY (`UUID`) REFERENCES `user_info` (`UUID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='푸쉬 매핑';

-- 테이블 데이터 hoseo.push_map:~0 rows (대략적) 내보내기
/*!40000 ALTER TABLE `push_map` DISABLE KEYS */;
/*!40000 ALTER TABLE `push_map` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
