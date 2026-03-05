package database
import(
	"fmt"
	"log"
	"os"
	"context"
	"time"
	"github.com/joho/godotenv"                  // .env dosyasındaki gizli şifreleri okuyan kütüphane.
	"go.mongodb.org/mongo-driver/mongo"         // MongoDB ile konuşan ana sürücü (Driver).
	"go.mongodb.org/mongo-driver/mongo/options" // Bağlantı ayarlarını yapılandıran paket.
	
)

func Connect() *mongo.Client {
	err:=godotenv.Load(".env")
	if err !=nil {
		log.Print("error loading .env file")
	}
	MongoDb := os.Getenv("MONGO_URI")
	if MongoDb =="" {
		log.Fatal("MONGO_URI not set!")
	}
	fmt.Print("MongoDb URL:",MongoDb)
	ctx,cancel:=context.WithTimeout(context.Background(),10*time.Second)
	defer cancel()
	clientOptions:=options.Client().ApplyURI(MongoDb)
	client,err:=mongo.Connect(ctx,clientOptions)
	if err!=nil {
		return nil
	}
	return client
}

func OpenCollection (collectionName string, client *mongo.Client)*mongo.Collection{
	err :=godotenv.Load(".env")
	if err !=nil{
		log.Print("error loading .env file")
	}
	databaseName:=os.Getenv("DATABASE_ENV")
	fmt.Print("DATABASE_NAME:",databaseName)
	collection:=client.Database(databaseName).Collection(collectionName)
	if collection == nil {
		return nil
	}
	return collection

}
