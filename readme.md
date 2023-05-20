# Basic usage

I have a ServerConfig struct that contains the database config
you can use your own struct or just pass the config directly

```go

	db, err = godab.Connect(&godab.DatabaseConfig{
		Hostname:  "localhost",
		Port:      3306,
		Driver:    "mysql",
		Name:      "database",
		Username:  "root",
		Password:  "root",
		ParseTime: true,
        })
	if err != nil {
		log.Fatalln("error while setup database", err)
	}

```

If you want to setup a Database from 0 you can use the OpenAndCreate function

i have a database.sql file in the same directory as my main.go
this file contains the sql code to create the database
if the database already exists, it will be ignored

```go

	// Setup database
	db, err = godab.OpenAndCreate(&godab.DatabaseConfig{
		Hostname:  "localhost",
		Port:      3306,
		Driver:    "mysql",
		Name:      "database",
		Username:  "root",
		Password:  "root",
		ParseTime: true,
        }, "../database.sql")
	if err != nil {
		log.Fatalln("error while setup database", err)
	}
```


After the connection is established, you can use the db variable to execute queries.

```go
    
    rows, err := godab.Query("SELECT * FROM `table`")
    if err != nil {
		log.Println(err)
	}

    row, err := godab.QueryRow("SELECT * FROM `table` LIMIT 1")
    if err != nil {
		log.Println(err)
	}

    result, err := godab.Exec("USE `database`")
    if err != nil {
		log.Println(err)
	}

```

OR 

```go

    rows, err := db.Query("SELECT * FROM `table`")
    if err != nil {
		log.Println(err)
	}

    row, err := db.QueryRow("SELECT * FROM `table` LIMIT 1")
    if err != nil {
		log.Println(err)
	}

    result, err := db.Exec("USE `database`")
    if err != nil {
		log.Println(err)
	}
```
