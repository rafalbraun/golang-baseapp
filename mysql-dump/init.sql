-- TODO - log warnings and errors

create table if not exists database_sessions (
	session_id INT unsigned auto_increment primary key,
	session_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

create table if not exists database_logs (
	correlation_id INT unsigned auto_increment primary key,
	log_timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
)

insert into database_sessions () values ();

-- TODO implement db logging


