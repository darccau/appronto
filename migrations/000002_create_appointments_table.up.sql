CREATE TABLE IF NOT EXISTS appointments (
	id bigserial PRIMARY KEY,
	/* patient_id INTEGER REFERENCES users(user_id), */
	/* doctor_id INTEGER REFERENCES users(user_id), */
	date_time TIMESTAMP WITH TIME ZONE NOT NULL,
	reason text,
	notes text,
	create_at timestamp(0) with time zone NULL NULL DEFAULT NOW(),
	version INTEGER NOT NULL DEFAULT 1
);

