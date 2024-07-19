# Usage
First, define the required environment variables (PORT, POSTGRES_USER, POSTGRES_PASSWORD, and POSTGRES_HOSTNAME):
```
cd server
nvim .env
```
Terminal session 1
```
cd server
go run .
```
Terminal session 2
```
cd server
docker compose up
```
Terminal session 3
```
cd client
npm i
npm run dev
```
Once done, CTRL+C in all terminal sessions and run ```docker compose down``` in session 2
