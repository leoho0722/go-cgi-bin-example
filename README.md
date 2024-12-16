# Go cgi-bin Example

## Prerequisites

### Apache2

#### Install Apache2

```shell
sudo apt update
sudo apt-get install -y apache2
```

#### Check Apache2 status

```shell
sudo systemctl status apache2
```

#### Restart Apache2

Restart apache2 service when cgi-bin file copied to `/usr/lib/cgi-bin` and `/var/www/html`.

```shell
sudo systemctl restart apache2
```

### Go

#### Install Go

```shell
cd ~
wget -O go1.22.2.linux-amd64.tar.gz https://go.dev/dl/go1.22.2.linux-amd64.tar.gz
tar -C /usr/local -xzf go1.22.2.linux-amd64.tar.gz
export PATH=$PATH:/usr/local/go/bin
```

#### Check Go version

```shell
go version
```

## Build + Deploy

```shell
make up
```

## Build

### Complie Go source code

#### Automatic

```shell
make build
```

#### Manual

```shell
# Change current path to project directory path
cd ~/go-cgi-bin-test

# build path1
cd ./cgi-bin-path1
go build -o path1
cd ../

# build path2
cd ./cgi-bin-path2
go build -o path2
cd ../

# build path3
cd ./cgi-bin-path3
go build -o path3
cd ../

# build path4
cd ./cgi-bin-path4
go build -o path4
cd ../
```

## Deploy

### Copy file to `usr/lib/cgi-bin` and `/var/www/html`

#### Automatic

```shell
make deploy
```

#### Manual

```shell
# Change current path to project directory path
cd ~/go-cgi-bin-test

# copy path1
cd ./cgi-bin-path1
sudo cp ./path1 /usr/lib/cgi-bin/path1
sudo cp ./path1 /var/www/html/path1
cd ../

# copy path2
cd ./cgi-bin-path2
sudo cp ./path2 /usr/lib/cgi-bin/path2
sudo cp ./path2 /var/www/html/path2
cd ../

# copy path3
cd ./cgi-bin-path3
sudo cp ./path3 /usr/lib/cgi-bin/path3
sudo cp ./path3 /var/www/html/path3
cd ../

# copy path4
cd ./cgi-bin-path4
sudo cp ./path4 /usr/lib/cgi-bin/path4
sudo cp ./path4 /var/www/html/path4
cd ../

# restart apache2
```

## Test

### cgi-bin-path1

```shell
# input
curl -X POST -H 'Content-Type: application/json' -d '{"input":"path1"}' http://127.0.0.1/cgi-bin/path1
```

```text
HTTP/1.1 200 OK
Content-Type: application/json

{
    "output": "path1, path1"
}
```

### cgi-bin-path2

```shell
# input
curl -X POST -H 'Content-Type: application/json' -d '{"input":"path2"}' http://127.0.0.1/cgi-bin/path2
```

```text
HTTP/1.1 200 OK
Content-Type: application/json

{
    "output": "path2, path2"
}
```

### cgi-bin-path3

```shell
# input
curl -X POST -H 'Content-Type: application/json' -d '{"input":"path3"}' http://127.0.0.1/cgi-bin/path3
```

```text
HTTP/1.1 200 OK
Content-Type: application/json

{
    "output": "path3, path3"
}
```

### cgi-bin-path4

```shell
# input
curl -X POST -H 'Content-Type: application/json' -d '{"input":"path4"}' http://127.0.0.1/cgi-bin/path4
```

```text
HTTP/1.1 200 OK
Content-Type: application/json

{
    "output": "path4, path4"
}
```

## Reference

1. https://veitchkyrie.github.io/2019/10/24/Ubuntu%E9%85%8D%E7%BD%AEApache-CGI%E7%A8%8B%E5%BA%8F/
