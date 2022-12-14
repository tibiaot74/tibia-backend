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
export JWT_KEY="VHfw3zAaBdPC)zHAVevFadAad2VDkX)KGnD+vLRvjH*5IvbBhUk3Fkzte5jPIG"
```

2) Create the database and run the the API locally on port `3000`
```sh
make local-init
```

**PS**: Although you will be using port 3000 to test the API, it's actually using port 7474, but to speed up the development we are using a tool called `gin` that recompile our code everytime you save any go file and serve as a proxy.

3) Create new features or fix bugs
Now you are free to develop with a simple local setup!

4) Run isolated tests
```sh
make  test
```

5) Are you done? Destroy everything running:
```sh
make local-destroy
make test-destroy
```

## Running locally (frontend)
If you are part of the frontend and wants to run this project locally you can simply run:

First set the environment variables:
```sh
export DB_NAME=ot
export DB_HOST=127.0.0.1
export DB_PORT=3306
export DB_USER=tibia
export DB_PASSWORD=123123
export JWT_KEY="VHfw3zAaBdPC)zHAVevFadAad2VDkX)KGnD+vLRvjH*5IvbBhUk3Fkzte5jPIG"
```

Then, run docker compose:
```sh
docker-compose -f docker-compose-local.yaml up --build
```

This will run the server on port 7474.

To shut down all the infrastructure just run:
```sh
docker-compose -f docker-compose-local.yaml down
```

## Documentation
This API server swagger on the following path: `swagger/index.html` if you want to see the webpage or `/swagger/doc.json` if you want the openapi specification JSON.
