CREATE TABLE IF NOT EXISTS `blog_auth` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(50) UNIQUE,
    `password` varchar(50) DEFAULT '',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `blog_tag` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(100) UNIQUE,
    `created_on` datetime,
    `created_by` varchar(50),
    `modified_on` datetime,
    `modified_by` varchar(50),
    `deleted_on` datetime,
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '0: disable, 1: enable',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `blog_post` (
    `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
    `title` varchar(100) NOT NULL,
    `description` varchar(255),
    `link` varchar(100) UNIQUE NOT NULL,
    `content` text,
    `created_on` datetime,
    `created_by` varchar(50),
    `modified_on` datetime,
    `modified_by` varchar(50),
    `deleted_on` datetime,
    `state` tinyint(3) unsigned DEFAULT '1' COMMENT '0: disable, 1: enable',
    PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;

CREATE TABLE IF NOT EXISTS `post_tag` (
    `tag_id` int(10) unsigned,
    `post_id` int(10) unsigned,
    PRIMARY KEY (`tag_id`, `post_id`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8;
