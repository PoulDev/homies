# homies
rommates app rest api

## Getting Started

### Prerequisites

- Docker
- Docker Compose

### Installation

#### 1. Clone the repo
```sh
git clone https://github.com/zibbadies/homies.git
```

#### 2. Change directory
```sh
cd homies
```

#### 3. Change the default environment variables
```sh
cp default.env .env
```

Now open the `.env` file with your favorite text editor and change the values to your own.

#### 4. Change the default docker-compose variables
Open the `docker-compose.yml` file with your favorite text editor and change the values to your own.
The variables to be changed are:

- `MYSQL_ROOT_PASSWORD`: The root password for the MySQL database
- `MYSQL_PASSWORD`: The password for the MySQL user, this must be the same as the `DB_PASSWORD` variable in the .env file

#### 5. Build the docker image
```sh
docker compose build
```

#### 6. Run the docker image
```sh
docker compose up -d
```

#### 7. Done!
The API is now running at `127.0.0.1:8080` 🥳

