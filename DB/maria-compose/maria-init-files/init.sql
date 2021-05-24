-- --------------------------------------------------------
-- 호스트:                          127.0.0.1
-- 서버 버전:                        10.6.0-MariaDB - mariadb.org binary distribution
-- 서버 OS:                        Win64
-- HeidiSQL 버전:                  11.1.0.6116
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

-- 테이블 hoseo.group_list 구조 내보내기
CREATE TABLE IF NOT EXISTS `group_list` (
  `Group_code` varchar(45) NOT NULL,
  `Group_name` varchar(45) NOT NULL,
  `address` varchar(45) NOT NULL,
  PRIMARY KEY (`Group_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.group_list:~0 rows (대략적) 내보내기
DELETE FROM `group_list`;
/*!40000 ALTER TABLE `group_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_list` ENABLE KEYS */;

-- 테이블 hoseo.group_map 구조 내보내기
CREATE TABLE IF NOT EXISTS `group_map` (
  `Group_code` varchar(45) NOT NULL,
  `UUID` varchar(45) NOT NULL,
  KEY `FK2_UUID` (`UUID`),
  KEY `FK1_Group_code` (`Group_code`),
  CONSTRAINT `FK1_Group_code` FOREIGN KEY (`Group_code`) REFERENCES `group_list` (`Group_code`),
  CONSTRAINT `FK2_UUID` FOREIGN KEY (`UUID`) REFERENCES `user_info` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.group_map:~0 rows (대략적) 내보내기
DELETE FROM `group_map`;
/*!40000 ALTER TABLE `group_map` DISABLE KEYS */;
/*!40000 ALTER TABLE `group_map` ENABLE KEYS */;

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
  CONSTRAINT `FK_Kiosk_Set_Group_code_Group_list_Group_code` FOREIGN KEY (`Group_code`) REFERENCES `group_list` (`Group_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='Kiosk 설정값';

-- 테이블 데이터 hoseo.kiosk_set:~0 rows (대략적) 내보내기
DELETE FROM `kiosk_set`;
/*!40000 ALTER TABLE `kiosk_set` DISABLE KEYS */;
/*!40000 ALTER TABLE `kiosk_set` ENABLE KEYS */;

-- 테이블 hoseo.user_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `user_info` (
  `UUID` varchar(45) NOT NULL,
  `email` varchar(45) NOT NULL,
  `displayname` varchar(45) NOT NULL,
  `token` varchar(45) NOT NULL,
  `PSN` varchar(45) NOT NULL,
  `PSN_img` varchar(100) NOT NULL,
  `ls_admin` tinyint(1) NOT NULL DEFAULT 0,
  `wearable_SN` varchar(45) NOT NULL,
  PRIMARY KEY (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.user_info:~0 rows (대략적) 내보내기
DELETE FROM `user_info`;
/*!40000 ALTER TABLE `user_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_info` ENABLE KEYS */;

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
  CONSTRAINT `FK_User_log_wearable_SN_Wearable_info_wearable_SN` FOREIGN KEY (`wearable_SN`) REFERENCES `wearable_info` (`wearable_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='이용자 로그';

-- 테이블 데이터 hoseo.user_log:~0 rows (대략적) 내보내기
DELETE FROM `user_log`;
/*!40000 ALTER TABLE `user_log` DISABLE KEYS */;
/*!40000 ALTER TABLE `user_log` ENABLE KEYS */;

-- 테이블 hoseo.wearable_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `wearable_info` (
  `wearable_SN` varchar(45) NOT NULL,
  `UUID` varchar(45) NOT NULL,
  PRIMARY KEY (`wearable_SN`),
  KEY `FK1_UUID` (`UUID`),
  CONSTRAINT `FK1_UUID` FOREIGN KEY (`UUID`) REFERENCES `user_info` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.wearable_info:~0 rows (대략적) 내보내기
DELETE FROM `wearable_info`;
/*!40000 ALTER TABLE `wearable_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `wearable_info` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
