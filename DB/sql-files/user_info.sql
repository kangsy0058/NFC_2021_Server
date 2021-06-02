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

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IFNULL(@OLD_FOREIGN_KEY_CHECKS, 1) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40111 SET SQL_NOTES=IFNULL(@OLD_SQL_NOTES, 1) */;
