package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Server struct {
	client *mongo.Client
}

func NewServer(c *mongo.Client) *Server {
	return &Server{
		client: c,
	}
}

func (s *Server) Start() error {
	http.HandleFunc("/facts", s.handleGetAllFacts)
	return http.ListenAndServe(":8080", nil)
}

func (s *Server) handleGetAllFacts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

	// Handle pre-flight OPTIONS request
	if r.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
		return
	}
	coll := s.client.Database("catfacts").Collection("facts")
	query := bson.M{}
	cursor, err := coll.Find(context.TODO(), query)
	if err != nil {
		log.Fatal(err)
	}
	var facts []bson.M
	if err := cursor.All(context.TODO(), &facts); err != nil {
		log.Fatal(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(facts)

}

type catFactWorker struct {
	client *mongo.Client
}

func NewCatFactWoker(c *mongo.Client) *catFactWorker {
	return &catFactWorker{
		client: c,
	}
}

func (cfw *catFactWorker) start() error {
	coll := cfw.client.Database("catfacts").Collection("facts")
	ticker := time.NewTicker(15 * time.Second)

	for {
		resp, err := http.Get("https://catfact.ninja/fact")
		if err != nil {
			return err
		}

		var catFact bson.M
		if err := json.NewDecoder(resp.Body).Decode(&catFact); err != nil {
			return err
		}
		_, err = coll.InsertOne(context.TODO(), catFact)
		if err != nil {
			return err
		}
		<-ticker.C
	}
}

func main() {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	worker := NewCatFactWoker(client)
	go worker.start()

	server := NewServer(client)
	server.Start()
}
