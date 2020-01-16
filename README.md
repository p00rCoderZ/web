# web
repo for kup website

### Getting api git submodule
```
# make sure to update and initialize submodules
git submodule update --init --recursive
```

### Starting containers
Before creating containers make sure to create secrets.toml file at the root of project.

```
echo 'serial = "your-secret-key"' > secrets.toml
```

```
# To start up containers 
docker-compose up -d
# To attach to website container
docker attach website
```

### Once in website container
```
# get all go packages
go get -d ./...
# check if API works properly
curl api:8000
```

### Running the website
```
go run main.go
# then on your local machine in any browser
localhost:8000
```

### Helpful commands
```
# To rebuild containers
docker-compose up -d --build
# To rebuild specific container
docker-compose up -d --build <container-name>
# To rebuild images
docker-compose build --no-cache
# To force recreate containers
docker-compose up -d --force-recreate
```
