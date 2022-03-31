CREATE TABLE IF NOT EXISTS `heroes` (
	`name` VARCHAR(50) NOT NULL,
	`max_hp` INT,
	`current_hp` INT,
	`race` VARCHAR(30),
	`gender` VARCHAR(30),
	`class` VARCHAR(30),
	`weight` FLOAT,
	`height` FLOAT,
	`money` FLOAT,
	`talents` INT
	PRIMARY KEY (`name`)
);

CREATE TABLE IF NOT EXISTS `attributes` (
	`name` VARCHAR(50) NOT NULL,
	`description` text
	PRIMARY KEY (`name`)
);

CREATE TABLE IF NOT EXISTS `talents` (
	`name` VARCHAR(50) NOT NULL,
    `attr1` VARCHAR(50) NOT NULL,
	`attr2` VARCHAR(50) NOT NULL,
	`attr3` VARCHAR(50) NOT NULL,
	`description` text
	PRIMARY KEY (`name`)
);

CREATE TABLE `hero_attributes` (
	`hero` VARCHAR(50) NOT NULL,
	`attribute` VARCHAR(50) NOT NULL,
	`value` INT NOT NULL,
	PRIMARY KEY (`hero`,`attribute`)
);

CREATE TABLE `hero_talents` (
	`hero` VARCHAR(50) NOT NULL,
	`talent` VARCHAR(50) NOT NULL,
	`value` INT NOT NULL,
	PRIMARY KEY (`hero`,`talent`)
);