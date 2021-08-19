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

-- 테이블 hoseo.kiosk_set 구조 내보내기
CREATE TABLE IF NOT EXISTS `kiosk_set` (
  `kiosk_SN` varchar(45) NOT NULL COMMENT '키오스크 S/N',
  `Group_code` varchar(45) DEFAULT NULL COMMENT '그룹코드',
  `detail_position` varchar(100) DEFAULT NULL COMMENT '상세위치',
  `building_name` varchar(45) DEFAULT NULL COMMENT '건물명',
  `latitude` float DEFAULT NULL COMMENT '위도',
  `longitude` float DEFAULT NULL COMMENT '경도',
  `kioskset_log` datetime DEFAULT NULL COMMENT '키오스크설치 로그',
  PRIMARY KEY (`kiosk_SN`),
  KEY `FK_Kiosk_Set_Group_code_Group_list_Group_code` (`Group_code`),
  CONSTRAINT `FK_Kiosk_Set_Group_code_Group_list_Group_code` FOREIGN KEY (`Group_code`) REFERENCES `group_list` (`Group_code`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `kiosk_set_FK` FOREIGN KEY (`kiosk_SN`) REFERENCES `kiosk_info` (`kiosk_SN`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='Kiosk 설정값';

-- 테이블 데이터 hoseo.kiosk_set:~11 rows (대략적) 내보내기
/*!40000 ALTER TABLE `kiosk_set` DISABLE KEYS */;
INSERT IGNORE INTO `kiosk_set` (`kiosk_SN`, `Group_code`, `detail_position`, `building_name`, `latitude`, `longitude`, `kioskset_log`) VALUES
	('KSN1111', '041-31499-g1', '충청남도 아산시 배방읍', '1공학관', 36.7364, 127.074, NULL),
	('KSN1112', '041-31499-g2', '충청남도 아산시 배방읍', '1공학관', 36.7365, 127.073, NULL),
	('KSN1113', '041-31499-g3', '충청남도 아산시 배방읍', '1공학관', 36.735, 127.074, NULL),
	('KSN1114', '041-31499-g4', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074, NULL),
	('KSN1115', '041-31499-g5', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074, NULL),
	('KSN1116', '041-31499-g6', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074, NULL),
	('KSN1117', '041-31499-g7', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074, NULL),
	('KSN1118', '041-31499-g8', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074, NULL),
	('KSN1119', '041-31499-g9', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074, NULL),
	('KSN1120', '041-31499-g10', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074, NULL),
	('ksn1121', '041-31499-g1', 'test', '1', 36.736, 127.074, '2021-08-17 16:32:47');
/*!40000 ALTER TABLE `kiosk_set` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
