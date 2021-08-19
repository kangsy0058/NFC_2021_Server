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

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
