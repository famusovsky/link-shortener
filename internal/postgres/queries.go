package postgres

const (
	createLinks = `CREATE TABLE IF NOT EXISTS links (
		id SERIAL PRIMARY KEY,
		link TEXT
	);`
	dropLinks = `DROP TABLE IF EXISTS links;`
	getLink   = `SELECT link FROM links WHERE id = $1;`
	getLinkId = `SELECT id FROM links WHERE link = $1;`
	addLink   = `INSERT INTO links (link) VALUES ($1) RETURNING id;`
)
