CREATE TABLE users(
	id SERIAL PRIMARY KEY NOT NULL,
	name TEXT NOT NULL,
	type TEXT NOT NULL);
CREATE TABLE user_relas(
	id INT NOT NULL,
	user_id INT NOT NULL,
	state TEXT NOT NULL,
	type TEXT NOT NULL,
	PRIMARY KEY (id, user_id),
	FOREIGN KEY (id) REFERENCES users(id),
	FOREIGN KEY (user_id) REFERENCES users(id));
CREATE INDEX UserTypeIndex ON users (type);
CREATE INDEX UserRelaTypeIndex ON user_relas (type);