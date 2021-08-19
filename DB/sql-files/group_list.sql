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

-- 테이블 hoseo.group_list 구조 내보내기
CREATE TABLE IF NOT EXISTS `group_list` (
  `Group_code` varchar(45) NOT NULL COMMENT '그룹코드',
  `Group_name` varchar(45) DEFAULT NULL COMMENT '그룹이름',
  `address` varchar(45) DEFAULT NULL COMMENT '주소',
  `Group_log` datetime DEFAULT NULL COMMENT '그룹리스트 로그',
  PRIMARY KEY (`Group_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='그룹 리스트';

-- 테이블 데이터 hoseo.group_list:~10 rows (대략적) 내보내기
/*!40000 ALTER TABLE `group_list` DISABLE KEYS */;
INSERT IGNORE INTO `group_list` (`Group_code`, `Group_name`, `address`, `Group_log`) VALUES
	('041-31499-g1', 'g1', '경기도 화성시 17-1', NULL),
	('041-31499-g10', 'g10', '경기도 화성시 18-1', NULL),
	('041-31499-g2', 'g2', '경기도 화성시 17-2', NULL),
	('041-31499-g3', 'g3', '경기도 화성시 17-3', NULL),
	('041-31499-g4', 'g4', '경기도 화성시 17-4', NULL),
	('041-31499-g5', 'g5', '경기도 화성시 17-5', NULL),
	('041-31499-g6', 'g6', '경기도 화성시 17-6', NULL),
	('041-31499-g7', 'g7', '경기도 화성시 17-7', NULL),
	('041-31499-g8', 'g8', '경기도 화성시 17-8', NULL),
	('041-31499-g9', 'g9', '경기도 화성시 17-9', NULL);
/*!40000 ALTER TABLE `group_list` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
