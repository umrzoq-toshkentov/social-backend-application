CREATE TABLE IF NOT EXISTS comments (
	id bigserial PRIMARY KEY,
	post_id bigserial NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
	user_id bigserial NOT NULL REFERENCES users(id) ON DELETE CASCADE,
	content TEXT NOT NULL,
	created_at TIMESTAMP(0) with time zone NOT NULL DEFAULT NOW()
);
