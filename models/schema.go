package models

func schemas() []string {

	return []string{
		`
		CREATE TABLE IF NOT EXISTS users (
			id bigint NOT NULL AUTO_INCREMENT,
			username varchar(16),
			email varchar(151),
			password varchar(255),
			user_type varchar(255),
			created_at bigint,
			updated_at bigint,
			deleted_at bigint,
			PRIMARY KEY (id)
		);`,
	}
}
