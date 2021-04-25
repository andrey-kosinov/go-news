CREATE TABLE users (
	id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	login VARCHAR(255) NOT NULL,
	password VARCHAR(255) NOT NULL,
	PRIMARY KEY (id),
	KEY (login)
);

INSERT INTO users (login, password) VALUES ("admin", MD5('admin'));

CREATE TABLE news (
	id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	title VARCHAR(255) NOT NULL,
	body Text NOT NULL,
	PRIMARY KEY (id)
);

INSERT INTO news (title, body) VALUES
	("Тектонический корунд: основные моменты","Фирн сдвигает стеклянный криптархей. Геосинклиналь, а также комплексы фораминифер, известные из валунных суглинков роговской серии, окислена. Активная тектоническая зона быстроспредингового хребта магнитуда землетрясения поступает в днепровский блеск."),
	("Почему нерезко проникновение глубинных магм?","Оттаивание пород занимает батолит. Застываение лавы, а также в преимущественно песчаных и песчано-глинистых отложениях верхней и средней юры, пододвигается под кристаллический комплекс. Присоединение органического вещества изменяет излом, что связано с мощностью вскрыши и полезного ископаемого."),
	("Кремнистый кварцит: гипотеза и теории","Магматическая дифференциация жестко пододвигается под голоцен. Эоловое засоление обедняет авгит. Кама анизотропно покрывает пирокластический отрог, что связано с мощностью вскрыши и полезного ископаемого. Глубина очага землетрясения, основываясь большей частью на сейсмических данных, причленяет к себе боксит.");

CREATE TABLE sessions (
	id INT(10) UNSIGNED NOT NULL AUTO_INCREMENT,
	name VARCHAR(255) NOT NULL,
	user_id INT(10) UNSIGNED NOT NULL,
	time_to_kill INT(10) UNSIGNED NOT NULL,
	PRIMARY KEY (id),
	KEY (name),
	KEY (time_to_kill)
);
