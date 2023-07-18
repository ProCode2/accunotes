# AccuNotes

A minimal REST API for a note taking app. 
```
POST /signup 
POST /login

GET /notes
POST /notes
DELETE /notes
```


### Details
- Get the published Docker Image [here](https://hub.docker.com/r/procode1/accunotes_api)
- Github Repository [here](https://github.com/procode2/accunotes)

### Run the Project
#### Run with docker compose
- Clone the repository (only the `docker-compose.yml` and `.env` file is really required)
```
git clone https://github.com/procode2/accunotes.git
```
- Create a `.env` file from `.env.sample` in the root of the project with your credentials
Note: These credentials are used for creating the db in docker.
- Run with docker compose
```
docker compose up
```
[Note: This should not happen, but incase `docker compose up` fails to connect the database, rerunning `docker compose up` generally fixes it]

#### Manually clone and run (required go, postgresql installed)
- Clone the repository
```
git clone https://github.com/procode2/accunotes.git
```
- Create a `.env` file from `.env.sample` in the root of the project with your credentials
- Run the project
```
go run cmd/main.go
```
