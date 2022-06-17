CREATE TABLE covid_observations(
	the_serial_id SERIAL PRIMARY KEY,
	the_country_name VARCHAR(100) NOT NULL, 
        the_state_name VARCHAR(100) NOT NULL, 

	confirmed_no INT  NOT NULL DEFAULT 0,

	deaths_no INT NOT NULL DEFAULT 0,

	recovered_no INT NOT NULL DEFAULT 0,
	observation_date TIMESTAMP NOT NULL,

	last_update TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
) 



INSERT INTO covid_observations (the_country_name, the_state_name,confirmed_no,deaths_no,recovered_no,observation_date)
VALUES('Mainland China','Anhui',102,10,2,timestamp '2015-01-10 17:00:00' );