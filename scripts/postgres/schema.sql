CREATE TABLE adtable(
	id serial PRIMARY KEY,
	ad_name character varying(200) NOT NULL,
	ad_description character varying(1000) NOT NULL,
	photo_links text ARRAY[3] NOT NULL,
	price decimal NOT NULL,
	date timestamp with time zone NOT NULL
);