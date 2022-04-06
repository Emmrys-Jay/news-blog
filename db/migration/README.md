# Database Details

Generate the SQL code using <a href="dbdiagram.io">dbdiagram.io</a> <br>

<p>DB migration files are used to easily change database qualities as the software requirements change.</p> 

The db migration CLI application used to create the migration files is <a href="https://github.com/golang-migrate">here</a>

```bash
$ migrate create -ext sql -dir db/migration -seq init_schema
```

The bash command above creates two db migration files, the UP and DOWN migration scripts. <br>
<ul>
    <li> The UP migration file is used to create the database schema.
    <li> The DOWN migration file is used to drop the database schema.
</ul>

<h3>Run migration command to create schema: <h3>

```bash
$ migrate -path db/migration -database "postgresql://root:password@localhost:8080/simple_bank?sslmode=disable" -verbose up
```