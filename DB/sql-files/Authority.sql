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

-- 테이블 hoseo.Authority 구조 내보내기
CREATE TABLE IF NOT EXISTS `Authority` (
  `Au_idx` int(11) NOT NULL AUTO_INCREMENT COMMENT '권한 목록',
  `UUID` varchar(45) DEFAULT NULL COMMENT '이용자',
  `Group_code` varchar(45) DEFAULT NULL COMMENT '그룹코드',
  `status` varchar(45) DEFAULT NULL COMMENT '승인 상태',
  `Au_log` datetime DEFAULT NULL COMMENT '권한 로그',
  PRIMARY KEY (`Au_idx`),
  KEY `Authority_FK` (`UUID`),
  KEY `Authority_FK_1` (`Group_code`),
  CONSTRAINT `Authority_FK` FOREIGN KEY (`UUID`) REFERENCES `user_info` (`UUID`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `Authority_FK_1` FOREIGN KEY (`Group_code`) REFERENCES `group_list` (`Group_code`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb3 COMMENT='권한';

-- 테이블 데이터 hoseo.Authority:~4 rows (대략적) 내보내기
/*!40000 ALTER TABLE `Authority` DISABLE KEYS */;
INSERT IGNORE INTO `Authority` (`Au_idx`, `UUID`, `Group_code`, `status`, `Au_log`) VALUES
	(1, '이용자4', '041-31499-g3', '승인', NULL),
	(2, '이용자7', '041-31499-g3', '승인 대기', NULL),
	(3, '이용자8', '041-31499-g7', '승인', NULL),
	(4, '이용자9', '041-31499-g4', NULL, NULL);
/*!40000 ALTER TABLE `Authority` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
