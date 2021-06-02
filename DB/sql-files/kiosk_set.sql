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

-- 테이블 hoseo.kiosk_set 구조 내보내기
CREATE TABLE IF NOT EXISTS `kiosk_set` (
  `kiosk_SN` varchar(45) NOT NULL COMMENT '키오스크 S/N',
  `Group_code` varchar(45) DEFAULT NULL COMMENT '그룹코드',
  `detail_position` varchar(100) DEFAULT NULL COMMENT '상세위치',
  `building_name` varchar(45) DEFAULT NULL COMMENT '건물명',
  `latitude` float DEFAULT NULL COMMENT '위도',
  `longitude` float DEFAULT NULL COMMENT '경도',
  PRIMARY KEY (`kiosk_SN`),
  KEY `FK_Kiosk_Set_Group_code_Group_list_Group_code` (`Group_code`),
  CONSTRAINT `FK2_kioskSN` FOREIGN KEY (`kiosk_SN`) REFERENCES `kiosk_info` (`kiosk_SN`),
  CONSTRAINT `FK_Kiosk_Set_Group_code_Group_list_Group_code` FOREIGN KEY (`Group_code`) REFERENCES `group_list` (`Group_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Kiosk 설정값';

-- 테이블 데이터 hoseo.kiosk_set:~10 rows (대략적) 내보내기
DELETE FROM `kiosk_set`;
/*!40000 ALTER TABLE `kiosk_set` DISABLE KEYS */;
INSERT INTO `kiosk_set` (`kiosk_SN`, `Group_code`, `detail_position`, `building_name`, `latitude`, `longitude`) VALUES
	('KSN1111', '0001', '충청남도 아산시 배방읍', '1공학관', 36.7364, 127.074),
	('KSN1112', '0002', '충청남도 아산시 배방읍', '1공학관', 36.7365, 127.073),
	('KSN1113', '0003', '충청남도 아산시 배방읍', '1공학관', 36.735, 127.074),
	('KSN1114', '0004', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074),
	('KSN1115', '0005', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074),
	('KSN1116', '0006', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074),
	('KSN1117', '0007', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074),
	('KSN1118', '0008', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074),
	('KSN1119', '0009', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074),
	('KSN1120', '0010', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074);
/*!40000 ALTER TABLE `kiosk_set` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
