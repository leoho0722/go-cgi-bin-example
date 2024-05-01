# Go cgi-bin Test

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

#### Complie Go source code

```shell
# in go-cgi-bin-test directory...

# cgi-bin-path1
cd cgi-bin-path1
go build -o path1 .

# cgi-bin-path2
cd cgi-bin-path2
go build -o path2 .

# cgi-bin-path3
cd cgi-bin-path3
go build -o cgi-bin-path3

# cgi-bin-path4
cd cgi-bin-path4
go build -o path4 .
```

## Copy file to `usr/lib/cgi-bin` and `/var/www/html`

```shell
# in go-cgi-bin-test/cgi-bin-path1 directory...
sudo cp ./path1 /usr/lib/cgi-bin/path1
sudo cp ./path1 /var/www/html/path1

# in go-cgi-bin-test/cgi-bin-path2 directory...
sudo cp ./path2 /usr/lib/cgi-bin/path2
sudo cp ./path2 /var/www/html/path2

# in go-cgi-bin-test/cgi-bin-path3 directory...
sudo cp ./path3 /usr/lib/cgi-bin/path3
sudo cp ./path3 /var/www/html/path3

# in go-cgi-bin-test/cgi-bin-path4 directory...
sudo cp ./path4 /usr/lib/cgi-bin/path4
sudo cp ./path4 /var/www/html/path4
```

## Test

### cgi-bin-path1

```shell
# input
curl -X POST -H 'Content-Type: application/json' -d '{"title":"a"}' http://127.0.0.1/cgi-bin/path1
```

```text
Path1 Request body:  {a}
HTTP/1.1 200 OK
Content-Type: application/json
```

### cgi-bin-path2

```shell
# input
curl -X POST -H 'Content-Type: application/json' -d '{"title":"b"}' http://127.0.0.1/cgi-bin/path2
```

```text
Path1 Request body:  {b}
HTTP/1.1 200 OK
Content-Type: application/json
```

### cgi-bin-path3

```shell
# input
curl -X POST -H 'Content-Type: application/json' -d '{"title":"3"}' http://127.0.0.1/cgi-bin/path3
```

```text
Path1 Request body:  {3}
HTTP/1.1 200 OK
Content-Type: application/json
```

### cgi-bin-path4

```shell
# input
curl -X POST -H 'Content-Type: application/json' -d '{"title":"4"}' http://127.0.0.1/cgi-bin/path4
```

```text
Path1 Request body:  {4}
HTTP/1.1 200 OK
Content-Type: application/json
```