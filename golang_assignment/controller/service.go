package controller

import (
	"encoding/json"
	"math"
	"net/http"
	"strconv"

	"github.com/go-gorm/gorm"
	"github.com/prabhat_mishra/golang_assignment/config/db"

	"github.com/gorilla/mux"
)

func dist(lat1 float32, long1 float32, lat2 float32, long2 float32) float32 {
	var p float32 = 22 / (7 * 180)
	a := 0.5 - math.Cos((lat2-lat1)*p)/2 + math.Cos(lat1*p)*math.Cos(lat2*p)*(1-math.Cos((long2-long1)*p))/2
	//a := 0.5 - cos((lat2-lat1)*p)/2 + cos(lat1*p)*cos(lat2*p)*(1-cos((lon2-lon1)*p))/2
	dist := 12742 * math.Asin(math.Sqrt(a)) // asin(sqrt(a))
	return dist
}

func getCab(w http.ResponseWriter, r *http.Request) {

	var nearCab []models.cab
	var customer models.rider
	var rideH models.ride
	var drH models.dr

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	strVar := params["id"]
	cust_id, _ := strconv.Atoi(strVar)
	var db1 *gorm.DB = db.Init()

	if err := db1.Where("cust_id = ?", cust_id).Find(&customer).Error; err != nil {
		http.Error(w, "Customer has not registered!", http.StatusBadRequest)
		return
	}

	// fmt.Println("Cab found")

	uD := dist()

	if err := db1.Where("SQRT( ABS(presentlatitude - ? )+ABS(presentlongitude - ? )) < ? AND status = ?", customer.Startlatitude, customer.Startlongitude, 2, true).Find(&nearCab).Error; err != nil {
		http.Error(w, "There are no vehicles around", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(nearCab)

}

func getHistory(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-type", "application/json")
	params := mux.Vars(r)
	strVar := params["id"]
	Cust_id, _ := strconv.Atoi(strVar)
	var ddr models.dr
	var riderH []models.rider
	var paym models.payment

	var db1 *gorm.DB = db.Init()

	if err := db1.Where("cust_id = ?", cust_id).Find(&customer).Error; err != nil {
		http.Error(w, "Customer has not registered!", http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(riderH)

}

func main()  {
	
}
