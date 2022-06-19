CREATE TABLE covid_observations(
id uuid DEFAULT uuid_generate_v4 () PRIMARY KEY,
sno INT ,
countryregion VARCHAR(100) NOT NULL,
provincestate VARCHAR(100) NOT NULL,
confirmed INT  NOT NULL DEFAULT 0,
deaths INT NOT NULL DEFAULT 0,
recovered INT NOT NULL DEFAULT 0,
observationdate TIMESTAMP NOT NULL,
lastupdate TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
UNIQUE(sno),
UNIQUE(countryregion,provincestate,confirmed,deaths,recovered)

)

INSERT INTO covid_observations (sno,   observationdate,  provincestate,  countryRegion,   lastupdate, confirmed, deaths, recovered )  VALUES  (1,  timestamp '01/22/2020',  'Anhui',  'Mainland China',  timestamp '1/22/2020 17:00', 1.0, 0.0, 0.0 ) ,  
(2,  timestamp '01/22/2020',  'Beijing',  'Mainland China',  timestamp '1/22/2020 17:00', 14.0, 0.0, 0.0 ) ,  
(3,  timestamp '01/22/2020',  'Chongqing',  'Mainland China',  timestamp '1/22/2020 17:00', 6.0, 0.0, 0.0 ) ,  
(6,  timestamp '01/22/2020',  'Guangdong',  'Mainland China',  timestamp '1/22/2020 17:00', 26.0, 0.0, 0.0 ) ,  
(7,  timestamp '01/22/2020',  'Guangxi',  'Mainland China',  timestamp '1/22/2020 17:00', 2.0, 0.0, 0.0 ) ,  
(9,  timestamp '01/22/2020',  'Hainan',  'Mainland China',  timestamp '1/22/2020 17:00', 4.0, 0.0, 0.0 ) 