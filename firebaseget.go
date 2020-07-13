
package main


import (
	"log"
	//"github.com/joho/godotenv"
	"smarttask/services"
	"fmt"
	 "google.golang.org/api/iterator"
)

func main() {
// connect to firebase
client, ctx := services.FirestoreClient()
defer client.Close()
//retrieve data
iter := client.Collection("users").Documents(ctx)
for {
        doc, err := iter.Next()
        if err == iterator.Done {
                break
        }
        if err != nil {
                log.Fatalf("Failed to iterate: %v", err)
		} 
		log.Println("Retrieving data...")
        fmt.Println(doc.Data())
}



}
  