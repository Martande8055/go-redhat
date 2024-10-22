<div align="center">
  
![Create Go App](https://github.com/create-go-app/cli/assets/11155743/95024afc-5e3b-4d6f-8c9c-5daaa51d080d)
  
![Go version](https://img.shields.io/github/go-mod/go-version/Martande8055/go-redhat)
![License](https://img.shields.io/github/license/Martande8055/go-redhat)
![issues](https://img.shields.io/github/issues/Martande8055/go-redhat)<br/>
![Github Stars](https://img.shields.io/github/stars/Martande8055/go-redhat)
![Github Forks](https://img.shields.io/github/forks/Martande8055/go-redhat)

# Average Calculation API

A simple HTTP server for calculating averages, sums, and counts of numbers (even, odd, or all) with a RESTful API.

</div>

## ⚡️ Quick Start

First, ensure you have [Go](https://go.dev/doc/install) installed. Version `1.22.7` or higher is required.

### Installation

Use the following command to clone and install the project:

```bash
git clone https://github.com/Martande8055/go-redhat.git
cd go-redhat/
go mod tidy
go build -o main main.go
```

To run the server:
```
./main
```
The server will start on port 3100.

Usage
You can send POST requests to the following endpoints to calculate averages:

- For Even Numbers: `/even`
  
- For Odd Numbers: `/odd`
  
- For All Numbers: `/evenodd`

Example Request
Here's how to make a request using curl:

```
curl -X POST http://localhost:3100/even -H "Content-Type: application/json" -d '{"op": "avg", "nums": [2, 4, 6, 8]}'
```
## ⚙️ API Endpoints

### `POST /even`

Calculate the average, sum, or count of even numbers.

You can Specify different Operations as follows:

- `avg` for average

- `sum` for sum

- `count` for count

Json Input Example:

```
{
  "op": "avg",
  "nums": [2, 4, 6, 8]
}
```

## 📝 Testing

Run the unit tests with:
```
go test ./...
```

## 🚚 Docker Support

To run the application in a Docker container:

1. Build the image:
```
podman build -t go_average_app .
```
2. Run the container:
```
podman run -p 3100:3100 go_average_app
```

## ⭐️ Project Assistance

If you find this project helpful, please give it a star on GitHub!

## ⚠️ License

The Average Calculation API is free and open-source software licensed under the MIT License.

