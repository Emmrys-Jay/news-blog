# Docker and Postgres container details

<h3>Download docker image using:</h3>

```bash
$ docker pull postgres
```
<h3>Run docker image using command in the form:</h3>

```bash
docker run --name <container_name> -e <environment_variable> -p <host_ports:container_ports> -d <image>:<tag> <br>
```

### Example:
```bash
$ docker run --name postgres14 -p 8080:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres
```

<h3>Connect to container and run a specific command inside:</h3>

```bash
$ docker exec -it <container_name_or_id> <command_and_args>
```

### Example:
```bash
$ docker exec -it postgres14 psql -U root 
```

<h3>Check logs of all commands that has been run on a container:</h3>

```bash
$ docker logs postgres14
```

### Within Shell:
<h3>Connect to postgres container and enter the shell:</h3>

```bash
$ docker exec -it postgres14 /bin/sh
```

<h3>Create DB from container shell:</h3>

```bash
$ createdb --username=root --owner=root simple_bank
```

<h3>Drop DB from the container shell:</h3>

```bash
$ dropdb simple_bank
```
<ul>
<li> simple_bank is the name of the database
</ul>