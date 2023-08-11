# smartshop-api-go-template

SmartShop template for api micro-services that are written in golang.
This is provided as an exemplar for consistency and as a quick start for new golang micro-services.

## Requirements

* Docker
* Go

## How to set up a new service

### 1) Use GitHub to copy smartshop-api-go-template to the new service repo:

* Open https://github.com/JSainsburyPLC and click on 'New'
* Click on 'No Template' and select smartshop-api-go-template
* Enter new service repo name 'Repository name'
* Fill in 'Description'
* Tick 'Private'
* Click 'Create repository'
* Add the GoCD Github user (js-smartshop-services) as a Collaborator to the repo with write access, a member of the ops
  team or your engineering manager needs to do this
* Ensure master is protected, only js-smartshop-services should be able to push to this branch, a member of the ops team
  or your engineering manager needs to do this

### 2) Clone new repo so that you can start working on it locally

### 3) Setup GoCD pipeline, dashboard and deployment config

https://sainsburys-confluence.valiantys.net/display/ISSA/Docker+Services

* GoCD
    * Create GoCD pipelines for the service, following the
      instructions [here](https://github.com/JSainsburyPLC/smartshop-services-gocd#adding-a-service)
    * Go to the GoCD pipelines and search for 'infra_ecr' and run 'infra_ecr.deploy.mgt' to pick up these changes

* Dashboard
    * Edit the .releasedash.yml in this repo and replace the text 'smartshop-api-go-template' with your new repo name
      using dashes

* ECS deployment config
    * Add an ECR repo for the service (this stores images for your service in AWS),
      Edit [smartshop-services-ops/blob/develop/ansible/roles/datastore/ecr/defaults/main.yml)](https://github.com/JSainsburyPLC/smartshop-services-ops/blob/develop/ansible/roles/datastore/ecr/defaults/main.yml)
      and add your github service repository (note: add it in alphabetical order!)
    * Create the ECS deployment config for the service, you will need to create an Ansible role, please use an existing
      one as a template for now using the examples listed below:
        * [Internal facing service, single Go container](https://github.com/JSainsburyPLC/smartshop-services-ops/tree/develop/ansible/roles/service/api_issa_datacash_fake_ecs)
        * [Public facing service, single Go container](https://github.com/JSainsburyPLC/smartshop-services-ops/tree/develop/ansible/roles/service/api_smartshop_toolbox_ecs)

### 4) Environment variables (with secrets)

Follow the
instructions [here](https://github.com/JSainsburyPLC/smartshop-services-tools#environment-variables-for-deployment-to-aws)

### 5) Update imports and tests to reflect your new repo name

* Replace the text 'api-go-template' with your new repo name using dashes
* Replace the text 'api_go_template' with your new repo name using underscores

### 6) Edit go.mod

Update 'module' with name of the new repo

Update 'go' with desired version

Update dependencies

### 7) Update New Relic app name

Change `NEW_RELIC_APP_NAME` and `NEW_RELIC_LABEL_ROLE` fields to each deployment/var field following the naming
pattern `smartshop-services-<env>-<service-name>`, e.g. `"smartshop-services-prd-api-identity-orchestrator"`.

### 8) Edit the Dockerfile

Set desired values for GO_VERSION and UBUNTU_VERSION.

The rest of this file uses the generic value of 'smartshop' inside the docker image for:

- working folder
- user
- group

The binary 'app' is installed in the root folder inside the docker image.

### 9) Register Error Code Prefix

Each SmartShop service requires a unique error code prefix. It must be registered in
the [SmartShop Backend repo](https://github.com/JSainsburyPLC/smartshop-backend)
Once you have registered a prefix you need to update the error codes constants that already exist to use that
prefix [HERE](errors/errors.go)

### 10) Check everything builds and tests run successfully

```BASH
$ make docker_deps
$ make deps
$ make docker_test
```

## Developing

#### First time setup after Git Clone

After cloning this repo from Github, run this command to setup all the dependencies, tools and mocks needed. Specific
commands for individual steps are listed below if needed.

```shell
make first_time_setup
```

#### Add the services tools submodule

```BASH
$ make docker_deps
```

Please see the README in the
[smartshop-services-tools](https://github.com/JSainsburyPLC/smartshop-services-tools/)
repo for more information about developing Docker based services.

#### Install Dependencies

```BASH
$ make deps
```

#### Build executable called 'smartshop-service', in the top level directory

```BASH
$ make build
```

### Develop locally

#### Build mocks

```BASH
$ make mocks
```

#### Run tests (excluding integration tests)

```BASH
$ make test
```

#### Run service locally and perform local testing

Start app in one window:

```BASH
$ make run
{"level":"info","msg":"app is running on IP 0.0.0.0, port 8080","time":"2020-02-19T14:02:38Z"}
```

In a separate window:

- Run integration tests (note: add smartshop-api-go-template-ecs.app.internal to /etc/hosts next to localhost )

```BASH
$ make ci_test
```

- Curl endpoints

```BASH
$ curl http://localhost:8080/healthcheck
```

### Running Docker locally

#### Run the service and any dependencies

```BASH
$ make docker_up
```

Note that when docker is up, you can go into a new docker container that is part of the docker network and then curl the
endpoints

```BASH
$ docker ps
$ docker exec -it <CONTAINER ID> /bin/bash
root@<CONTAINER ID>:/smartshop# curl http://api-go-template.app.internal:8080/healthcheck
```

#### Run the tests (including integration tests)

```BASH
$ make docker_test
```

#### Stop the service and any dependencies

```BASH
$ make docker_down
```

### Connecting to local postgres database in a Docker container

#### Installing psql

Ensure you have libpq installed and in your path

```BASH
$ brew install libpq
$ export PATH=$PATH:/usr/local/bin:/usr/local/opt/libpq/bin
```

#### Using psql

```BASH
$ psql -h localhost -p 5432 -U smartshop_api_go_template
```
