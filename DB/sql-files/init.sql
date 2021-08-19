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

-- 테이블 hoseo.FCM 구조 내보내기
CREATE TABLE IF NOT EXISTS `FCM` (
  `token` varchar(45) DEFAULT NULL,
  `ch_id` int(11) DEFAULT NULL,
  KEY `FCM_FK` (`token`),
  KEY `FCM_FK_1` (`ch_id`),
  CONSTRAINT `FCM_FK` FOREIGN KEY (`token`) REFERENCES `user_info` (`token`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `FCM_FK_1` FOREIGN KEY (`ch_id`) REFERENCES `push_info` (`ch_id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='FCM';

-- 테이블 데이터 hoseo.FCM:~0 rows (대략적) 내보내기
/*!40000 ALTER TABLE `FCM` DISABLE KEYS */;
/*!40000 ALTER TABLE `FCM` ENABLE KEYS */;

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

-- 테이블 hoseo.kiosk_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `kiosk_info` (
  `kiosk_SN` varchar(45) NOT NULL COMMENT '키오스크 시리얼넘버',
  `kiosk_log` datetime DEFAULT NULL COMMENT '키오스크 로그',
  UNIQUE KEY `kiosk_SN` (`kiosk_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='Kiosk 정보';

-- 테이블 데이터 hoseo.kiosk_info:~11 rows (대략적) 내보내기
/*!40000 ALTER TABLE `kiosk_info` DISABLE KEYS */;
INSERT IGNORE INTO `kiosk_info` (`kiosk_SN`, `kiosk_log`) VALUES
	('KSN1111', NULL),
	('KSN1112', NULL),
	('KSN1113', NULL),
	('KSN1114', NULL),
	('KSN1115', NULL),
	('KSN1116', NULL),
	('KSN1117', NULL),
	('KSN1118', NULL),
	('KSN1119', NULL),
	('KSN1120', NULL),
	('ksn1121', '2021-08-17 05:00:00');
/*!40000 ALTER TABLE `kiosk_info` ENABLE KEYS */;

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

-- 테이블 hoseo.push_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `push_info` (
  `ch_id` int(11) NOT NULL AUTO_INCREMENT,
  `push` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL CHECK (json_valid(`push`)),
  PRIMARY KEY (`ch_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb3 COMMENT='푸쉬';

-- 테이블 데이터 hoseo.push_info:~1 rows (대략적) 내보내기
/*!40000 ALTER TABLE `push_info` DISABLE KEYS */;
INSERT IGNORE INTO `push_info` (`ch_id`, `push`) VALUES
	(1, '{"token":"aaa.bbb.ccc"}');
/*!40000 ALTER TABLE `push_info` ENABLE KEYS */;

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

-- 테이블 hoseo.user_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `user_info` (
  `UUID` varchar(45) NOT NULL COMMENT '이용자',
  `email` varchar(45) DEFAULT NULL COMMENT '이메일',
  `displayname` varchar(45) DEFAULT NULL COMMENT '사용자이름표시',
  `token` varchar(45) DEFAULT NULL COMMENT '토큰',
  `PSN` varchar(45) DEFAULT NULL COMMENT '개인안심번호',
  `PSN_img` varchar(100) DEFAULT NULL COMMENT '개인안심번호 이미지',
  `Is_admin` tinyint(4) DEFAULT NULL COMMENT '관리자판별',
  `wearable_SN` varchar(45) DEFAULT NULL COMMENT '웨어러블S/N',
  `Group_code` varchar(45) DEFAULT NULL,
  PRIMARY KEY (`UUID`) USING BTREE,
  UNIQUE KEY `user_info_UN` (`token`),
  UNIQUE KEY `user_info_UN2` (`wearable_SN`),
  KEY `user_info_FK` (`Group_code`),
  CONSTRAINT `user_info_FK` FOREIGN KEY (`Group_code`) REFERENCES `group_list` (`Group_code`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `user_info_FK_1` FOREIGN KEY (`wearable_SN`) REFERENCES `wearable_info` (`wearable_SN`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='이용자 정보';

-- 테이블 데이터 hoseo.user_info:~11 rows (대략적) 내보내기
/*!40000 ALTER TABLE `user_info` DISABLE KEYS */;
INSERT IGNORE INTO `user_info` (`UUID`, `email`, `displayname`, `token`, `PSN`, `PSN_img`, `Is_admin`, `wearable_SN`, `Group_code`) VALUES
	('user11', NULL, NULL, NULL, NULL, NULL, NULL, 'wsn1121', NULL),
	('이용자1', 'uuid1@naver.com', '이용자1', 'aaaa.bbbb.cccc', '12가34나', 'C:Users사용자이름PicturesMyImg1.jpg', 1, 'wsn1111', '041-31499-g1'),
	('이용자10', 'uuid5@gmail.com', '이용자10', 'aaca.bbbb.cbcc', '22가34나', 'C:Users사용자이름PicturesMyImg10.jpg', 1, 'wsn1120', '041-31499-g10'),
	('이용자2', 'uuid2@naver.com', '이용자2', 'aaba.bbbb.cccc', '22가34나', 'C:Users사용자이름PicturesMyImg2.jpg', 1, 'wsn1112', '041-31499-g10'),
	('이용자3', 'uuid3@naver.com', '이용자3', 'abba.bbbb.cccc', '32가34나', 'C:Users사용자이름PicturesMyImg3.jpg', 1, 'wsn1113', '041-31499-g3'),
	('이용자4', 'uuid4@naver.com', '이용자4', 'aaaa.abbb.cccc', '42가34나', 'C:Users사용자이름PicturesMyImg4.jpg', 1, 'wsn1114', '041-31499-g4'),
	('이용자5', 'uuid5@naver.com', '이용자5', 'aaaa.bbab.cccc', '52가34나', 'C:Users사용자이름PicturesMyImg5.jpg', 1, 'wsn1115', '041-31499-g5'),
	('이용자6', 'uuid@gmail.com', '이용자6', 'aaaa.baba.cccc', '12가64나', 'C:Users사용자이름PicturesMyImg6.jpg', 1, 'wsn1116', '041-31499-g5'),
	('이용자7', 'uuid2@gmail.com', '이용자7', 'aaaa.bbbb.accc', '12바36나', 'C:Users사용자이름PicturesMyImg7.jpg', 1, 'wsn1117', '041-31499-g6'),
	('이용자8', 'uuid3@gmail.com', '이용자8', 'aaaa.bbbb.ccbc', '12가32나', 'C:Users사용자이름PicturesMyImg8.jpg', 1, 'wsn1118', '041-31499-g7'),
	('이용자9', 'uuid4@gmail.com', '이용자9', 'aaba.bbbb.ccbc', '12가39나', 'C:Users사용자이름PicturesMyImg9.jpg', 1, 'wsn1119', '041-31499-g8');
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

-- 테이블 hoseo.wearable_info 구조 내보내기
CREATE TABLE IF NOT EXISTS `wearable_info` (
  `wearable_SN` varchar(45) NOT NULL,
  `wearable_log` datetime DEFAULT NULL,
  PRIMARY KEY (`wearable_SN`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb3 COMMENT='웨어러블 정보';

-- 테이블 데이터 hoseo.wearable_info:~12 rows (대략적) 내보내기
/*!40000 ALTER TABLE `wearable_info` DISABLE KEYS */;
INSERT IGNORE INTO `wearable_info` (`wearable_SN`, `wearable_log`) VALUES
	('wsn1111', NULL),
	('wsn1112', NULL),
	('wsn1113', NULL),
	('wsn1114', NULL),
	('wsn1115', NULL),
	('wsn1116', NULL),
	('wsn1117', NULL),
	('wsn1118', NULL),
	('wsn1119', NULL),
	('wsn1120', NULL),
	('wsn1121', NULL),
	('wsn1122', NULL);
/*!40000 ALTER TABLE `wearable_info` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
