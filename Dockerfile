FROM maven:3.8.6-eclipse-temurin-11-alpine
RUN apk update -y
RUN apk add wget -y
RUN apk add git -y
WORKDIR /
RUN mkdir app
WORKDIR app
RUN git clone -b develop https://gabriel.mocchetti:d14g01010$$@gitlab.eldars.com.ar/devops/apis-remote-access/
RUN ls -lha
WORKDIR /app/apis-remote-access/
RUN mvn clean install -DskipTests
WORKDIR /app/apis-remote-access/target
RUN mv *.jar app.jar 
EXPOSE 8087
ENTRYPOINT ["java", "-jar", "app.jar"]
