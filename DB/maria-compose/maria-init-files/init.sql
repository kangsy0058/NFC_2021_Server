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

-- 테이블 hoseo.group_list 구조 내보내기
CREATE TABLE IF NOT EXISTS `group_list` (
  `Group_code` varchar(45) NOT NULL,
  `Group_name` varchar(45) NOT NULL,
  `address` varchar(45) NOT NULL,
  PRIMARY KEY (`Group_code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.group_list:~10 rows (대략적) 내보내기
DELETE FROM `group_list`;
/*!40000 ALTER TABLE `group_list` DISABLE KEYS */;
INSERT INTO `group_list` (`Group_code`, `Group_name`, `address`) VALUES
	('0001', '그룹1', '경기도 화성시 17-1'),
	('0002', '그룹2', '경기도 화성시 17-2'),
	('0003', '그룹3', '경기도 화성시 17-3'),
	('0004', '그룹4', '경기도 화성시 17-4'),
	('0005', '그룹5', '경기도 화성시 17-5'),
	('0006', '그룹6', '경기도 화성시 17-6'),
	('0007', '그룹7', '경기도 화성시 17-7'),
	('0008', '그룹8', '경기도 화성시 17-8'),
	('0009', '그룹9', '경기도 화성시 17-9'),
	('0010', '그룹10', '경기도 화성시 18-1');
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

-- 테이블 데이터 hoseo.group_map:~10 rows (대략적) 내보내기
DELETE FROM `group_map`;
/*!40000 ALTER TABLE `group_map` DISABLE KEYS */;
INSERT INTO `group_map` (`Group_code`, `UUID`) VALUES
	('0001', '이용자1'),
	('0002', '이용자2'),
	('0003', '이용자3'),
	('0004', '이용자4'),
	('0005', '이용자5'),
	('0006', '이용자6'),
	('0007', '이용자7'),
	('0008', '이용자8'),
	('0009', '이용자9'),
	('0010', '이용자10');
/*!40000 ALTER TABLE `group_map` ENABLE KEYS */;

-- 테이블 hoseo.kiosk_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `kiosk_info` (
  `kiosk_SN` varchar(45) NOT NULL,
  UNIQUE KEY `kiosk_SN` (`kiosk_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.kiosk_info:~0 rows (대략적) 내보내기
DELETE FROM `kiosk_info`;
/*!40000 ALTER TABLE `kiosk_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `kiosk_info` ENABLE KEYS */;

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

-- 테이블 hoseo.push_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `push_info` (
  `ch_id` int(11) NOT NULL AUTO_INCREMENT,
  `push` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`push`)),
  PRIMARY KEY (`ch_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.push_info:~0 rows (대략적) 내보내기
DELETE FROM `push_info`;
/*!40000 ALTER TABLE `push_info` DISABLE KEYS */;
/*!40000 ALTER TABLE `push_info` ENABLE KEYS */;

-- 테이블 hoseo.push_map 구조 내보내기
CREATE TABLE IF NOT EXISTS `push_map` (
  `ch_id` int(11) NOT NULL DEFAULT 0,
  `UUID` varchar(45) NOT NULL,
  KEY `FK1_chid` (`ch_id`),
  KEY `FK2_chUUID` (`UUID`),
  CONSTRAINT `FK1_chid` FOREIGN KEY (`ch_id`) REFERENCES `push_info` (`ch_id`),
  CONSTRAINT `FK2_chUUID` FOREIGN KEY (`UUID`) REFERENCES `user_info` (`UUID`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.push_map:~0 rows (대략적) 내보내기
DELETE FROM `push_map`;
/*!40000 ALTER TABLE `push_map` DISABLE KEYS */;
/*!40000 ALTER TABLE `push_map` ENABLE KEYS */;

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
  PRIMARY KEY (`UUID`) USING BTREE,
  UNIQUE KEY `wearable_SN` (`wearable_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.user_info:~10 rows (대략적) 내보내기
DELETE FROM `user_info`;
/*!40000 ALTER TABLE `user_info` DISABLE KEYS */;
INSERT INTO `user_info` (`UUID`, `email`, `displayname`, `token`, `PSN`, `PSN_img`, `ls_admin`, `wearable_SN`) VALUES
	('이용자1', 'uuid1@naver.com', '이용자1', 'aaaa.bbbb.cccc', '12가34나', 'C:Users사용자이름PicturesMyImg1.jpg', 0, 'wsn1111'),
	('이용자10', 'uuid5@gmail.com', '이용자10', 'aaca.bbbb.cbcc', '22가34나', 'C:Users사용자이름PicturesMyImg10.jpg', 0, 'wsn1120'),
	('이용자2', 'uuid2@naver.com', '이용자2', 'aaba.bbbb.cccc', '22가34나', 'C:Users사용자이름PicturesMyImg2.jpg', 0, 'wsn1112'),
	('이용자3', 'uuid3@naver.com', '이용자3', 'abba.bbbb.cccc', '32가34나', 'C:Users사용자이름PicturesMyImg3.jpg', 0, 'wsn1113'),
	('이용자4', 'uuid4@naver.com', '이용자4', 'aaaa.abbb.cccc', '42가34나', 'C:Users사용자이름PicturesMyImg4.jpg', 0, 'wsn1114'),
	('이용자5', 'uuid5@naver.com', '이용자5', 'aaaa.bbab.cccc', '52가34나', 'C:Users사용자이름PicturesMyImg5.jpg', 1, 'wsn1115'),
	('이용자6', 'uuid@gmail.com', '이용자6', 'aaaa.baba.cccc', '12가64나', 'C:Users사용자이름PicturesMyImg6.jpg', 1, 'wsn1116'),
	('이용자7', 'uuid2@gmail.com', '이용자7', 'aaaa.bbbb.accc', '12바36나', 'C:Users사용자이름PicturesMyImg7.jpg', 1, 'wsn1117'),
	('이용자8', 'uuid3@gmail.com', '이용자8', 'aaaa.bbbb.ccbc', '12가32나', 'C:Users사용자이름PicturesMyImg8.jpg', 0, 'wsn1118'),
	('이용자9', 'uuid4@gmail.com', '이용자9', 'aaba.bbbb.ccbc', '12가39나', 'C:Users사용자이름PicturesMyImg9.jpg', 1, 'wsn1119');
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
  CONSTRAINT `FK2_wearable_info_wearableSN` FOREIGN KEY (`wearable_SN`) REFERENCES `wearable_info` (`wearable_SN`),
  CONSTRAINT `FK_User_log_kiosk_SN_Kiosk_Set_kiosk_SN` FOREIGN KEY (`kiosk_SN`) REFERENCES `kiosk_set` (`kiosk_SN`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8 COMMENT='이용자 로그';

-- 테이블 데이터 hoseo.user_log:~10 rows (대략적) 내보내기
DELETE FROM `user_log`;
/*!40000 ALTER TABLE `user_log` DISABLE KEYS */;
INSERT INTO `user_log` (`IDX`, `wearable_SN`, `kiosk_SN`, `time`, `date`, `temp`, `Group_code`, `detail_position`, `building_name`, `latitude`, `longitude`) VALUES
	(1, 'WSN1111', 'KSN1111', '03:14:18', '2021-05-16', 36.5, '0001', '충청남도 아산시 배방읍', '1공학관', 36.7364, 127.074),
	(2, 'WSN1112', 'KSN1111', '03:14:18', '2021-05-16', 36.5, '0002', '충청남도 아산시 배방읍', '1공학관', 36.7365, 127.073),
	(3, 'WSN1113', 'KSN1111', '03:14:18', '2021-05-16', 36.5, '0003', '충청남도 아산시 배방읍', '1공학관', 36.735, 127.074),
	(4, 'WSN1114', 'KSN1113', '03:14:18', '2021-05-16', 37.5, '0004', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074),
	(5, 'WSN1115', 'KSN1113', '03:14:18', '2021-05-16', 36.5, '0005', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074),
	(6, 'WSN1116', 'KSN1115', '03:14:18', '2021-05-13', 36.5, '0006', '충청남도 아산시 배방읍', '2공학관', 36.7364, 127.074),
	(7, 'WSN1117', 'KSN1112', '03:14:18', '2021-05-15', 35.5, '0007', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074),
	(8, 'WSN1118', 'KSN1119', '03:14:18', '2021-05-14', 36, '0008', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074),
	(9, 'WSN1119', 'KSN1119', '03:14:18', '2021-05-12', 37, '0009', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074),
	(10, 'WSN1120', 'KSN1119', '03:14:18', '2021-05-13', 36.5, '0010', '충청남도 아산시 배방읍', '강석규교육관', 36.7364, 127.074);
/*!40000 ALTER TABLE `user_log` ENABLE KEYS */;

-- 테이블 hoseo.wearable_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `wearable_info` (
  `wearable_SN` varchar(45) NOT NULL,
  UNIQUE KEY `wearable_SN` (`wearable_SN`),
  CONSTRAINT `FK1_wearableSN` FOREIGN KEY (`wearable_SN`) REFERENCES `user_info` (`wearable_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 테이블 데이터 hoseo.wearable_info:~10 rows (대략적) 내보내기
DELETE FROM `wearable_info`;
/*!40000 ALTER TABLE `wearable_info` DISABLE KEYS */;
INSERT INTO `wearable_info` (`wearable_SN`) VALUES
	('wsn1111'),
	('wsn1112'),
	('wsn1113'),
	('wsn1114'),
	('wsn1115'),
	('wsn1116'),
	('wsn1117'),
	('wsn1118'),
	('wsn1119'),
	('wsn1120');
/*!40000 ALTER TABLE `wearable_info` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
