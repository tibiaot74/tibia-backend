# tibia-backend

## Running Locally
You can run this API by using docker-compose. 

We expect the following environment variables:
- `DB_NAME`: Name of the database to connect
- `DB_HOST`: IP or Host of the database to connect
- `DB_PORT`: Port of the database to connect
- `DB_USER`: Username to connect to the database
- `DB_PASSWORD`: Password to connect to the database
- `JWT_KEY`: JWT key to be used

In order to locally set this variables you can run the following command:

```sh
export DB_NAME=ot
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_USER=tibia
export DB_PASSWORD=123123
export JWT_KEY=1029457820519
```

Then you can run the docker-compose:

```sh
docker-compose up -d
```

Make changes to your code and your the API:

```sh
go run main.go
```

If you want to delete everything you can run:

```sh
docker-compose down
```