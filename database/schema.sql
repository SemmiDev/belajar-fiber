CREATE TABLE `products` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `amount` int(11) DEFAULT NULL,
  `name` varchar(255) DEFAULT NULL,
  `price` bigint(20) DEFAULT NULL,
  `description` text,
  `category` text NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `id` (`id`),
  KEY `ix_length_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1;
