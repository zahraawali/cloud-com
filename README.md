# Install and run the project

## Run PostgreSQL and SMTP test server

> Both PostgreSQL and SMTP test server are inside docker container.

Run `docker compose up -d` that will start `PostgreSQL` on port 5432 and `SMTP test server` on port `2525` as well as a dashboard on port `3000`.

## Download and run `ActiveMQ`

1. Download `ActiveMQ` for your desired operating system from `https://activemq.apache.org/activemq-5018001-release`.
2. Extract the downloaded file.
3. Open powershell inside the extracted directory, and run: `./bin/activemq start`.

## Run `auth` and `reporter` services.

1. To run `auth` service, by executing `go run cmd/main.go`
2. Then, run `reporter` service, by executing `go run main.go serve`

## Run `core` service.

1. `cd core`
2. `./mvnw clean install`
3. `java -jar target/fum-cloud-notification-core-0.0.1-SNAPSHOT.jar`