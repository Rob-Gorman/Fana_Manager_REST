# Fana Manager

The Flag Manager consists of:
1. dashboard user interface that is used to create a set of attributes, audiences, and flags 
2. backend API server that processes the data and makes queries to a persistent PostgreSQL database.

# Usage
There are two options to get started in a self-hosted environment.

Deploy the entire Fana Platform stack using Docker Compose yaml file found [here](https://github.com/fana-io/fana-deploy).

Pull the Docker image from DockerHub and run the container in an existing Docker network with the prerequisite components:
```
$ docker pull fanaff/manager-static
```
Prerequisites include:

- PostgreSQL running on port `5432`
- Redis Cluster running on port `6379`
