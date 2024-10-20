# Http Mirror

## Features

1. Proxy a request to another url
2. Read proxy config from json or yaml (toml ?)
3. Clone the request to 2 urls
4. Save the 2 reponses in a db
5. show the results in a html page
6. add a header to the cloned request X-Shadow-Request


- Enable and disable the db part (example we want to use dynatrace instead of saving the data locally)
- Add another header to both requests X-Shadow-Number: XSR-{uuid}

## Libraries

Cannot be 100% library-less because we need a database driver

1. go get github.com/mattn/go-sqlite3
2. go get github.com/jmoiron/sqlx
