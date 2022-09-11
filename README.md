# tibia-backend

## Dependencies
You must install the following to be able to run contribute with this repository:
- go
- docker
- docker-compose

## Running Locally
We expect the following environment variables:
- `DB_NAME`: Name of the database to connect
- `DB_HOST`: IP or Host of the database to connect
- `DB_PORT`: Port of the database to connect
- `DB_USER`: Username to connect to the database
- `DB_PASSWORD`: Password to connect to the database
- `JWT_KEY`: JWT key to be used


1) In order to locally set this variables you can run the following command:
```sh
export DB_NAME=ot
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_USER=tibia
export DB_PASSWORD=123123
export JWT_KEY="VHfw^**aBdPC)!zHAVev!%#dA@d@VDkX)KGnD+v!RvjH*5IvbBhUk3%kzte5jPIG"
```

2) Create the database and run the the API locally on port `3000`
```sh
make local-init
```

**PS**: Although you will be using port 3000 to test the API, it's actually using port 7474, but to speed up the development we are using a tool called `gin` that recompile our code everytime you save any go file and serve as a proxy.

3) Are you done? Destroy everything running:

```sh
make local-destroy
```
