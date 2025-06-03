# Rest With Go
I built this app to try my hand with the Go Programming Language, especially testing it with PostgreSQL and Docker.

For now, it's a CLI app that lets you perform crud actions on a database that launches on startup. In the near future, I'm adding a basic UI to display the user table.

## Startup Guide
1. Clone the app from Github.
2. Add your own docker secrets file (more info coming soon).
3. Start your docker container with `docker-compose up -d`.
4. Open your terminal and type `go run main.go`. 
5. Test the crud actions using CURL from your terminal. (examples coming soon)