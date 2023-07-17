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

### Installation Steps
- Clone the repository
```
git clone https://github.com/procode2/accunotes.git
```
- Create a `.env` file from `.env.sample` in the root of the project with your credentials
Note: These credentials are used for creating the db in docker.

- Run with docker compose
```
docker compose up
```
