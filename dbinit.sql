create database maldb;
use maldb;

CREATE TABLE `malware` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `hostname` varchar(128) DEFAULT NULL,
  `type` char(32) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

insert into maldb.malware values 
(NULL,'test1.com','malware'),
(NULL,'test2.com','phishing'),
(NULL,'test3.com','malware'),
(NULL,'test4.com','spyware'),
(NULL,'test5.com','adware');
