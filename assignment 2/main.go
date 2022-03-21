package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type Order struct {
	gorm.Model
	CustomerName string
	OrderedAt    time.Time
	Items        []Item `gorm:"foreignkey:OrderID"`
}

type Item struct {
	gorm.Model
	ItemCode    string `gorm:"unique"`
	Description string
	Quantity    int
	OrderID     int
}

type TmpOrder struct {
	CustomerName string
	OrderedAt    time.Time
	Items        []struct {
		LineItemId  int    `json:"lineItemId"`
		ItemCode    string `json:"itemCode"`
		Description string `json:"Description"`
		Quantity    int    `json:"quantity"`
		OrderID     int    `json:"orderID"`
	}
}

var db *gorm.DB
var err error

func main() {

	dialect := "postgres"
	host := "localhost"
	dbPort := "5432"
	user := "postgres"
	password := "postgres"
	dbName := "orders_by"

	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	db, err = gorm.Open(dialect, dbURI)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Printf("Listening in port:%v\n", dbPort)
	}

	defer db.Close()

	db.AutoMigrate(&Order{})
	db.AutoMigrate(&Item{})

	router := mux.NewRouter()

	router.HandleFunc("/orders", getOrders).Methods("GET")
	router.HandleFunc("/orders/{orderID}", getOrder).Methods("GET")

	router.HandleFunc("/orders", createOrder).Methods("POST")
	router.HandleFunc("/orders/{orderID}", updateOrder).Methods("PUT")
	router.HandleFunc("/orders/{orderID}", deleteOrder).Methods("DELETE")
	http.ListenAndServe(":8080", router)

}

func getOrders(w http.ResponseWriter, r *http.Request) {
	var orders []Order

	db.Find(&orders)
	db.Preload("Items").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func getOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var order Order
	var items []Item

	getOrder := db.First(&order, params["orderID"])
	db.Model(&order).Related(&items)

	order.Items = items

	if getOrder.Error != nil {
		json.NewEncoder(w).Encode("data yang dicari tidak ada")
	} else {
		json.NewEncoder(w).Encode(order)
	}
}

func createOrder(w http.ResponseWriter, r *http.Request) {
	var order Order
	json.NewDecoder(r.Body).Decode(&order)

	createdOrder := db.Create(&order)
	err = createdOrder.Error
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode("data berhasil disimpan")
	}
}

func deleteOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var order Order
	var items Item

	deletedOrders := db.First(&order, params["orderID"]).Delete(&order)
	deletedItems := db.Where("order_id = ?", params["orderID"]).Delete(&items)

	if deletedOrders.Error != nil && deletedItems.RowsAffected < 1 {
		json.NewEncoder(w).Encode("Data gagal dihapus")
	} else {
		json.NewEncoder(w).Encode("Data berhasil dihapus")
	}
}

func updateOrder(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var newOrder TmpOrder
	var updatedOrder Order

	b, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	json.NewDecoder(r.Body).Decode(&newOrder)

	if err := json.Unmarshal([]byte(string(b)), &newOrder); err != nil {
		log.Fatalln(err)
		return
	}

	orderExist := db.First(&updatedOrder, params["orderID"])
	if orderExist.Error != nil {
		json.NewEncoder(w).Encode("tidak ada hasil")
	} else {
		updatedOrder.CustomerName = newOrder.CustomerName
		updatedOrder.OrderedAt = newOrder.OrderedAt
		savedOrder := db.Save(&updatedOrder)
		if savedOrder.Error != nil {
			json.NewEncoder(w).Encode("gagal menyimpan perubahan order")
		} else {
			err := false
			for _, item := range newOrder.Items {
				var updatedItem Item
				db.First(&updatedItem, item.LineItemId)
				updatedItem.ItemCode = item.ItemCode
				updatedItem.Description = item.Description
				updatedItem.Quantity = item.Quantity
				updatedItem.OrderID = int(updatedOrder.ID)
				savedItem := db.Save(&updatedItem)
				if savedItem.Error != nil {
					json.NewEncoder(w).Encode(savedItem.Error)
					err = true
				}
			}
			if !err {
				json.NewEncoder(w).Encode("data berhasil diperbarui")
			}
		}
	}
}
