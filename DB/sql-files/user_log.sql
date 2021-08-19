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

-- 테이블 hoseo.user_log 구조 내보내기
CREATE TABLE IF NOT EXISTS `user_log` (
  `IDX` int(11) NOT NULL AUTO_INCREMENT COMMENT '순번',
  `wearable_SN` varchar(45) DEFAULT NULL COMMENT '웨어러블 S/N',
  `kiosk_SN` varchar(45) DEFAULT NULL COMMENT '키오스크 S/N',
  `time` time DEFAULT NULL COMMENT '시간',
  `date` date DEFAULT NULL COMMENT '날짜',
  `temp` float DEFAULT NULL COMMENT '체온',
  `Group_code` varchar(45) DEFAULT NULL COMMENT '그룹코드',
  `detail_position` varchar(100) DEFAULT NULL COMMENT '상세위치',
  `building_name` varchar(45) DEFAULT NULL COMMENT '건물명',
  `latitude` float DEFAULT NULL COMMENT '위도',
  `longitude` float DEFAULT NULL COMMENT '경도',
  PRIMARY KEY (`IDX`),
  KEY `FK_User_log_kiosk_SN_Kiosk_Set_kiosk_SN` (`kiosk_SN`),
  KEY `FK_User_log_wearable_SN_Wearable_info_wearable_SN` (`wearable_SN`),
  CONSTRAINT `FK_User_log_kiosk_SN_Kiosk_Set_kiosk_SN` FOREIGN KEY (`kiosk_SN`) REFERENCES `kiosk_set` (`kiosk_SN`),
  CONSTRAINT `user_log_FK` FOREIGN KEY (`wearable_SN`) REFERENCES `user_info` (`wearable_SN`)
) ENGINE=InnoDB AUTO_INCREMENT=12 DEFAULT CHARSET=utf8mb3 COMMENT='이용자 로그';

-- 테이블 데이터 hoseo.user_log:~10 rows (대략적) 내보내기
/*!40000 ALTER TABLE `user_log` DISABLE KEYS */;
INSERT IGNORE INTO `user_log` (`IDX`, `wearable_SN`, `kiosk_SN`, `time`, `date`, `temp`, `Group_code`, `detail_position`, `building_name`, `latitude`, `longitude`) VALUES
	(1, 'WSN1111', 'KSN1111', '03:14:18', '2021-05-16', 36.5, '041-31499-g1', '충청남도 아산시 배방읍', '1공학관', 36.7364, 127.074),
	(2, 'WSN1112', 'KSN1111', '03:14:18', '2021-05-16', 36.5, '041-31499-g2', '충청남도 아산시 배방읍', '1공학관', 36.7365, 127.073),
	(3, 'WSN1113', 'KSN1111', '03:14:18', '2021-05-16', 36.5, '041-31499-g3', '충청남도 아산시 배방읍', '1공학관', 36.735, 127.074),
	(4, 'WSN1114', 'KSN1113', '03:14:18', '2021-05-16', 37.5, '041-31499-g4', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074),
	(5, 'WSN1115', 'KSN1113', '03:14:18', '2021-05-16', 36.5, '041-31499-g5', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074),
	(6, 'WSN1116', 'KSN1115', '03:14:18', '2021-05-13', 36.5, '041-31499-g6', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074),
	(7, 'WSN1117', 'KSN1112', '03:14:18', '2021-05-15', 35.5, '041-31499-g7', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074),
	(8, 'WSN1118', 'KSN1119', '03:14:18', '2021-05-14', 36, '041-31499-g8', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074),
	(9, 'WSN1119', 'KSN1119', '03:14:18', '2021-05-12', 37, '041-31499-g9', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074),
	(10, 'WSN1120', 'KSN1119', '03:14:18', '2021-05-13', 36.5, '041-31499-g10', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074);
/*!40000 ALTER TABLE `user_log` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
