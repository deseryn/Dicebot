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

INSERT INTO attributes VALUES 
	("Mut", "Mut steht für den Wagemut, die Willenskraft und Entschlossenheit des Charakters."),
	("Klugheit", "Klugheit steht für die abstrakte Intelligenz, das logische Denken und gutes Gedächtnis eines Lebewesens. "),
	("Intuition", "Intuition ist ein regeltechnisches Maß der Fähigkeit, spontan aus dem Bauch heraus richtig zu entscheiden und der Fähigkeit, sich auf seine Umgebung einzustellen."),
	("Gewandtheit", "Gewandtheit steht für körperliche Beweglichkeit und das Gelingen schneller Reaktionen eines Charakters."),
	("Charisma", "Charisma steht für die Wirkung und den Eindruck, den der Charakter auf andere Personen hat und macht."),
	("Fingerfertigkeit", "Fingerfertigkeit ist das regeltechnische Maß der Geschicklichkeit von Spielerhelden. Sie wird benötigt, um Aufgaben zu bewältigen, welche viel mit den Händen zu tun haben (die sog. Hand-Augen-Koordination)."),
	("Konstitution", "Konstitution ist ein regeltechnisches Maß der körperlichen Widerstandskraft und der Durchhaltefähigkeit eines Charakters."),
	("Körperkraft", "Körperkraft steht für die körperliche Muskelkraft eines Charakters.");
