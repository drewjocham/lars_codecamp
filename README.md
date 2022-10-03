# lars_codecamp
When you first Checkout the project run `go mod init codecamp` then 
1) run `make mod-vendor`
2) run `make network`
3) run `make database`

## Create the database
The sql files are in LocalDepend but the command will create the Database
1) run `make database`
2) Access the database with the below credentials
    - database=demo
    - user=postgres
    - password=postgres
3) Run the sql files found in `localDepend`

