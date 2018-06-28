# Weather Reporter

### Get Dependencies
**Run** `$ dep ensure`

### Setup
Set properties on > conf/app.conf
* appname = weather-reporter
* httpport = 8080
* runmode = dev
* endpoint = "http://api.openweathermap.org/data/2.5/weather?q=%s,%s&appid=%s"
* appid = "APP_ID"
* dbdsn = "DATA_SOURCE_NAME_/_CONNECTION_STRING

Notice the %s in the endpoint. We must keep them to inject the app id and query parameters formatting the endpoint string.

### DB
The app/service uses a mysql database. to create reports table (which is used to save weather reports temporally) run
```
CREATE TABLE `reports` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `location_name` varchar(100) DEFAULT NULL,
  `temperature` varchar(100) DEFAULT NULL,
  `wind` varchar(100) DEFAULT NULL,
  `pressure` varchar(100) DEFAULT NULL,
  `humidity` varchar(100) DEFAULT NULL,
  `sunrise` varchar(100) DEFAULT NULL,
  `sunset` varchar(100) DEFAULT NULL,
  `geo_coordinates` varchar(100) DEFAULT NULL,
  `requested_time` datetime DEFAULT NULL,
  `code_name` varchar(100) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=21 DEFAULT CHARSET=latin1;
``` 
### Run service
**Run** `$ bee run`