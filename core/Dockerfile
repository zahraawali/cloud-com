FROM openjdk:17-bullseye

WORKDIR /app/activemq

COPY ../activemq/* .

RUN ./bin/activemq start

WORKDIR /app

COPY target/*.jar .

EXPOSE 8080

CMD ["java", "-jar", "fum-cloud-notification-core-0.0.1-SNAPSHOT.jar"]