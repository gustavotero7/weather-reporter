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

### DB migration
The app/service uses a mysql database. run
```
$ bee migrate -driver=mysql -conn=root:password@/weatherreporter?parseTime=true
``` 
To create necessary db tables

### Run service
**Run** `$ bee run`