# Websites Checker

This application will be check website available. Also, user can upload list of website in .CSV file. The Application developed by vite-vue, go

## 📁  File structure
    .
    ├── backend                   # Go REST API service folder
    ├── frontend                  # Vite-Vue front application folder
    ├── .gitignore                     
    ├── docker-compose.yml        # Compose file for startup all project
    ├── README.md                   
    └── site.csv                  # File for test upload

## Prerequisite
- Docker 🐳 

## 🚀  How to start application

For run all application go to the root folder and run follow this command.

```
$ docker-compose up -d
```


## 🎨  Frontend Start

If you want to run only frontend, you should install packge with `yarn install` before. You can use `Makefile` to run frontend follow this.

First step, build image file. This command will be create image tag name `myvuechecker`

```
$ make image
```

Second step, build container. This command will be create container startup with port `:8080`

```
$ make container
```

How to run frontend unit test.

```
$ make test
```

## ⚙  Backend Start

If you want to run only backend. You can use `Makefile` to run backend follow this.

First step, build image file. This command will be create image name `mygochecker`

```
$ make image
```

Second step, build container. This command will be create container startup with port `:3000`

```
$ make container
```

Other command, To run unit test you can use this command.

```
$ make test
```

To run benchmark test you can use this command.

```
$ make test-bench
```