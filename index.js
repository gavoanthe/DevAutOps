const { exec } = require('node:child_process')
const fs = require('fs')
const { stdout } = require('node:process')


const parametro = {
    'dist': 'maven:3.8.6-eclipse-temurin-11-alpine',
    'port': 8087,
    'command': 'apt update -y',
    'MAVEN_VERSION' : '3.8.6',
    'password_git' : 'd14g01010$$',
    'user_git' : 'gabriel.mocchetti',
    'repo_git' : 'gitlab.eldars.com.ar/devops/apis-remote-access/',
    'branch' : 'develop',
    'commands' : {
        'update' : 'RUN apk update -y',
        'wget' : 'RUN apk add wget -y',
        'git' : 'RUN apk add git -y'
    },
}

const linea =
`
FROM ${parametro.dist}
${parametro.commands.update}
${parametro.commands.wget}
${parametro.commands.git}
WORKDIR /
RUN mkdir app
WORKDIR app
RUN git clone -b ${parametro.branch} https://${parametro.user_git}:${parametro.password_git}@${parametro.repo_git}
RUN ls -lha
WORKDIR /app/apis-remote-access/
RUN mvn clean install -DskipTests
WORKDIR /app/apis-remote-access/target
RUN mv *.jar app.jar 
EXPOSE 8087
ENTRYPOINT ["java", "-jar", "app.jar"]
`


function creararchivo(){
    fs.writeFile("Dockerfile", linea, (err) => {
        if (err)
            console.log(err);
        else {
            console.log("File written successfully\n");
            console.log("The written has the following contents:");
            console.log(fs.readFileSync("Dockerfile", "utf8"));
        }
    })
}

creararchivo()





