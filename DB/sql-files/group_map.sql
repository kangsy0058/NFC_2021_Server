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

-- 테이블 hoseo.group_map 구조 내보내기
CREATE TABLE IF NOT EXISTS `group_map` (
  `Group_code` varchar(45) NOT NULL,
  `UUID` varchar(45) NOT NULL,
  KEY `FK2_UUID` (`UUID`),
  KEY `FK1_Group_code` (`Group_code`),
  CONSTRAINT `FK1_Group_code` FOREIGN KEY (`Group_code`) REFERENCES `group_list` (`Group_code`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FK2_UUID` FOREIGN KEY (`UUID`) REFERENCES `user_info` (`UUID`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='단체 정보';

-- 테이블 데이터 hoseo.group_map:~10 rows (대략적) 내보내기
/*!40000 ALTER TABLE `group_map` DISABLE KEYS */;
INSERT IGNORE INTO `group_map` (`Group_code`, `UUID`) VALUES
	('041-31499-g1', '이용자1'),
	('041-31499-g1', '이용자2'),
	('041-31499-g3', '이용자3'),
	('041-31499-g4', '이용자4'),
	('041-31499-g5', '이용자5'),
	('041-31499-g6', '이용자6'),
	('041-31499-g7', '이용자7'),
	('041-31499-g8', '이용자8'),
	('041-31499-g9', '이용자9'),
	('041-31499-g10', '이용자10');
/*!40000 ALTER TABLE `group_map` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
