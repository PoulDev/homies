# roommates-api
rommates app rest api

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

1. Clone the repo
```sh
git clone https://github.com/PoulDev/roommates-api.git
```

2. Change directory
```sh
cd roommates-api
```

3. Change the default environment variables in the `.env` file
```sh
cp default.env .env
```

Now open the `.env` file with your favorite text editor and change the values to your own.

4. Change the default environment variables in the `docker-compose.yml` file
Open the `docker-compose.yml` file with your favorite text editor and change the values to your own.
The variables to be changed are:

- `MYSQL_ROOT_PASSWORD`: The root password for the MySQL database
- `MYSQL_PASSWORD`: The password for the MySQL user, this must be the same as the `DB_PASSWORD` variable in the .env file

5. Build the docker image
```sh
docker-compose build
```

6. Run the docker image
```sh
docker-compose up -d
```

7. The API is now running on `127.0.0.1:8080` 🥳

