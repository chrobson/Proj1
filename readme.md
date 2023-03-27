# Cat Facts API

This is a simple Cat Facts API written in Go. It fetches cat facts from an external API, stores them in a MongoDB database, and serves them through an HTTP server.

## Features

- Fetches cat facts from https://catfact.ninja/fact
- Stores cat facts in a MongoDB database
- Serves all stored cat facts through an HTTP endpoint

## Installation

1. Make sure you have Go installed on your system. You can download it from the official website: https://golang.org/dl/

2. Install the required dependencies:

```bash
go get -u go.mongodb.org/mongo-driver/mongo
go get -u go.mongodb.org/mongo-driver/mongo/options
go get -u go.mongodb.org/mongo-driver/bson
```

3. Clone this repository:

```bash
git clone https://github.com/chrobson/catfacts-api.git
```

4. Change the current directory to the project directory:

```bash
cd catfacts-api
```

5. Run the project:

```bash
npm start
```

## Usage

1. Ensure you have 
 - MongoDB installed and running on your local machine (or docker image). 
If not, you can follow the instructions here to install it: https://docs.mongodb.com/manual/installation/
 - Same for React, First, make sure you have Node.js and npm installed on your machine. You can download Node.js from https://nodejs.org/.
    - Create a new React application using the following command:
    ```npx create-react-app catfacts-frontend```
    - Change to the new directory ```cd catfacts-frontend```
    - Install Axios to handle HTTP requests: ```npm install axios```


2. Start the Cat Facts API:
```
go run main.go
./catfacts-api
npm start
```

3. The API will start fetching cat facts every 15 seconds, store them in the MongoDB database and display on React Frontend page. The HTTP API server will listen on port `8080`, frontend will be on port `3000`.

4. To get all cat facts from the API, make a GET request to the `/facts` endpoint:

```bash
curl http://localhost:8080/facts
```

5. The API will return a JSON array of cat facts.

## API Endpoints

- `GET /facts`: Retrieve all stored cat facts.

## Configuration

You can customize the MongoDB connection string by editing the following line in the `main` function:

```go
client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
```
Replace "mongodb://localhost:27017" with your desired MongoDB connection string.

## TODO

1. Improve build (to build npm into build and serve it with go)
2. Improve frontend by adding features such as automatic refresh, pagination, or styling.