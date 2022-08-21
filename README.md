# Tigerlily-bff

Tigerlily-bff by definition is the backend for the tigerlily-app (frontend)
All requests from the frontend goes through tigerlily-bff, which acts as a middleman proxy to all other services

## Prerequisites
> **IMPORTANT**: Make sure all project repositories are stored locally in a single folder under *$HOME* `~/tigercoders`. his is important for starting all services locally with docker compose later on

  **Go version 1.17 https://golang.org/dl/ (Go version 1.17.2 specifically)**
> Note: You can have multiple versions of the Go compiler installed locally and switch between them during development 

If you have other Go versions currently active, you'd have to either uninstall it and reinstall Go 1.17.2 or install the 1.17.2 Go compiler and and replace GOPATH env alias

## Setup

Tigerlily-bff is using GOMODULES for dependency management, so the source can be placed outside of ``GOPATH/src``

1. Clone the repository from Github
> **IMPORTANT**: Make sure all project repositories are stored locally in a single directory under *$HOME* `~/tigercoders`. his is important for starting all services locally with docker compose later on

> Note: Add your `ssh public key` to our github organisation before you begin these steps

In your local machine, run:
  - `mkdir ~/tigercoders`
  - `cd ~/tigercoders`
  - `git clone git@github.com:Tiger-Coders/<REPO-NAME>.git`
  > Note: Clone via ssh. Use HTTP if ssh doesn't work

2. Copy the service specific configuration into the specific project

We are following the configuration design from the `Twelve factor app` principles (https://12factor.net/)

All application has its own credentials and it lives outside the app in a config file stored in our secrets repository.
Our application injects those values into the app during build time and watches those configuration during runtime

> Note: We don't have a specific secret repository yet. Service Configurations are passed on manually currently between contributors. Just ask in slack

```touch ~/tigercoders/<SERVICE-REPO>/config.yml```

Then paste the configuration into `config.yml`

3. Start the service locally
> Note: There are 2 ways we can start all services locally for development (Docker, independent)

  ## a. Docker Compose ##

  When developing locally, we need more than one service/repo to start in order to test and verify that our changes are valid. Traditionally, we would start all required services manually in a seperate terminal. This requires you having multiple terminal windows and VSC open. This could introduce human errors and the need to restart those services upon each code change. 

  Our services requires a couple of external dependencies in order to run. (Postgres & redis at the moment). Installing these dependencies individually in each of our computers takes time and could also risk having them installed with a different breaking verison. If we were to have 10 external dependencies, imagine how much time would be spent debugging dependency errors ... Yikes! 😱😱😱

  To solve this, and to enable each new contributor to become as productive as soon as possible, we have compiled all services and dependencies into a container. Now all a new contributor has to do in order to start productively contributing is to `install docker` and run `docker compose up`. That simple!! No headaches installing external dependencies and running into errors (Ideally at least.. 😛😛😛) 

  > Note: To start all services locally with Docker compose, make sure you have *docker desktop* installed and that you are in *~/tigercoders/tigerlily-bff* and . This is where the docker-compose services map lives
    
    - Change Directory into `~/tigercoders/tigerlily-bff`
    
      `cd ~/tigercoders/tigerlily-bff`
    
    - Run `docker compose up`
  >**IMPORTANT**: Make sure all project repos live inside `~/tigercoders`
     
      `docker compose up`

  Go being a compiled language means that for each code change, we'd have to stop the running server and recompile our Go binary. This slows down productivity. 

  We are using `air` (https://github.com/cosmtrek/air) for Go live-reload to rebuild and compile upon detecting code changes. 
  
  *You don't have to recompile the Go binary to see your changes reflected during development* 😉

  For for our frontend application (*tigerlily-app*), we are already making use of *react-scripts start* to allow hot-reloads

  ## Without Docker ##
  > Note Currently our work is a little messy. In terms of configuration. So far only *tigerlily-inventory* & *tigerlily-bff* have been migrated to use the `Twelve Factor App` principles. To start *tigerlily-payment*, we have to run `go run main.go` passing in credentials as arguments. We have made it easier by creating a bash script to start each service, so just run that. 

  > Note: Once all services has been migrated to use the *Twelve Factor App* principle, we would use *Make* to automate configuration management

  - In each service, you'd find a `start_<SERVICE-NAME>.sh`. Just execute that file
    `bash start_<SERVICE-NAME>.sh`


# Architecture/System Diagrams #
![system design](https://user-images.githubusercontent.com/61228520/148627782-61206386-9490-4c89-a002-55a7651db1f7.png)

<h2>Flow of Real Time inventory updates:</h2>

![Tigerlily Project](https://user-images.githubusercontent.com/61228520/171106435-f03fa48b-18c4-4a79-98b9-8cacb5e184a5.png)

<h2>Flow of user sending a checkout request</h2>

![Tigerlily Project (3)](https://user-images.githubusercontent.com/61228520/172862392-11337251-5633-4f9b-9855-4aa5cc552f90.png)

<h2>User Flow Diagram</h2>

![Tigerlily Project (1)](https://user-images.githubusercontent.com/61228520/171108925-ee16476a-a3d5-4ac0-9a4a-278455a95f93.png)

