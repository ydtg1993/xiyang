CREATE TABLE `source_seed` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `source_url` varchar(255) NOT NULL,
  `title` varchar(255) NOT NULL,
  `cover` varchar(255) NOT NULL,
  `big_cover` varchar(255) NOT NULL DEFAULT '',
  `description` varchar(500) NOT NULL DEFAULT '',
  `publish_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `type` varchar(100) NOT NULL DEFAULT '',
  `tag` json NOT NULL,
  `origin` varchar(100) NOT NULL DEFAULT '',
  `content` text NOT NULL,
  `links` text NOT NULL,
  `raw_content` text NOT NULL,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `source_url_UNIQUE` (`source_url`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资源'