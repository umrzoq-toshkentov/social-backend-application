CREATE TABLE posts (
	id bigserial PRIMARY KEY,
	title text NOT NULL,
	content TEXT NOT NULL,
	user_id bigint NOT NULL,
	tags TEXT[] NOT NULL DEFAULT '{}',
	created_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW(),
	updated_at TIMESTAMP(0) WITH TIME ZONE NOT NULL DEFAULT NOW()
);
