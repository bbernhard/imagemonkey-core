The following document briefly describes how to start up your own ImageMonkey instance in a Docker container. 

# Modes

The ImageMonkey docker container supports three different modes

* normal
* run tests: run all unit- and integration tests from the [tests folder](https://github.com/bbernhard/imagemonkey-core/tree/master/tests)
* run stresstest

# Normal Mode
## Scenario #1
This is the most common scenario and the easiest to set up. Choose this option, if your web browser and your docker container will run on the same machine. 

* install docker
* download ImageMonkey Dockerfile with `https://raw.githubusercontent.com/bbernhard/imagemonkey-core/develop/env/docker/Dockerfile`
* change into the directory where the Dockerfile resides and run `docker build -t imagemonkey-core .`
* start docker instance with `docker run --ulimit nofile=90000:90000 -p 8080:8080 -p 8081:8081 imagemonkey-core`

This will start a new ImageMonkey docker instance on your machine. After your docker instance is up and running, you will see the following screen: 

![alt text](https://raw.githubusercontent.com/bbernhard/imagemonkey-core/develop/env/docker/documentation/screenshots/docker_container_ready.png)

Now open your browser and navigate to `http://127.0.0.1:8080`

![alt text](https://raw.githubusercontent.com/bbernhard/imagemonkey-core/develop/env/docker/documentation/screenshots/imagemonkey_localhost.png)

You can now upload your first image to your own ImageMonkey instance.

## Scenario #2
As docker acquires a significant portion of your systems resources, one might want to run the docker instance on a different machine. 

Let's assume your workstation has the private IP `192.168.1.9`. As your workstation is quite old, you want to run the ImageMonkey docker container on a different machine (e.q Raspberry Pi) which is in the same subnet and has the IP `192.168.1.16`. 

* install docker
* download ImageMonkey Dockerfile with `https://raw.githubusercontent.com/bbernhard/imagemonkey-core/develop/env/docker/Dockerfile`
* change into the directory where the Dockerfile resides and run `docker build -t imagemonkey-core .`
* start docker instance with `docker run -e API_BASE_URL=http://192.168.1.16:8081 --ulimit nofile=90000:90000 -p 8080:8080 -p 8081:8081 imagemonkey-core`

The docker run command looks almost identical to the one in Scenario #1, except that we are setting the `API_BASE_URL` environmental variable inside the docker container to the host systems IP (i.e `192.168.1.16`). After your docker instance is up and running, you will see the following screen: 



Now open your browser and navigate to `http://192.168.1.16:8080`

## FAQ
**Detailed description of the docker run command** 

`-p 8080:8080 -p 8081:8081`

Both the ImageMonkey webservice and the ImageMonkey API listen on port 8080 resp. port 8081 inside the docker container.
In order to make ImageMonkey easily accessible on the host system, we are mapping the hosts ports 8080 and 8081 
to the corresponding ports inside the docker container. 

The docker port mapping is also helpful if you already have a service running on the hosts system, that is listening on port 8080 or 8081. In that case you would need to choose different host ports and map those to port 8080 and 8081 inside the docker container. e.q: the commandline options `-p 8082:8080 -p 8083:8081` map the host port 8082 to the docker container port 8080 and the host port 8083 to the docker container port 8081. 

`--ulimit nofile=90000:90000` This commandline option changes the number of available file descriptors within the docker container. Without that option `redis` will not be able to run inside the docker container. 

# Run Tests
`docker run imagemonkey-core --run-tests`

# Run stresstest

`docker run --mount type=bind,source=/home/imagemonkey/imagemonkey_04_11_2018.zip,target=/tmp/stresstest/imagemonkey_data.zip,readonly imagemonkey-core --run-stresstest`
