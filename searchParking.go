package main

import (
	"fmt"
    "encoding/json"
    "log"
    "net/http"
	"io/ioutil"
	"bytes"
    "github.com/gorilla/mux"
)

type ReserveParkingRequest struct {
	CardID string `json:"cardId"`
	Email string `json:"email"`
	Ends string `json:"ends"`
	FacilityID int `json:"facility_id"`
	ParkingExtension bool `json:"parkingExtension"`
	Price int `json:"price"`
	ReservationID string `json:"reservationId"`
	RuleGroupID int `json:"rule_group_id"`
	Starts string `json:"starts"`
	Timezone string `json:"timezone"`
	UUID string `json:"uuid"`

}

type ReserveParkingResponse struct {
	Error interface{} `json:"error"`
	Data []struct {
		ReservationID string `json:"reservationId"`
		FacilityID int `json:"facilityId"`
		RuleGroUpID int `json:"ruleGro upId"`
		ParkingStartTime int64 `json:"parkingStartTime"`
		ParkingEndTime int64 `json:"parkingEndTime"`
		ExtensionParkingEndTime int64 `json:"extensionParking EndTime"`
		Status string `json:"status"`
		Reason interface{} `json:"reason"`
		RentalID int `json:"rentalId"`
		Created int64 `json:"created"`
		ExtensionID string `json:"extensionId"`
		Price int `json:"price"`
		TotalPrice int `json:"totalPrice"`
		FacilityTitle string `json:"facilityTitle"`
		UUID string `json:"uuid"`
		ZoneID interface{} `json:"zoneId"`
		Zone interface{} `json:"zone"`
		SpaceID interface{} `json:"spaceId"`
		Space interface{} `json:"space"`
		Provider string `json:"provider"`
		Latitude interface{} `json:"latitude"`
		Longitude interface{} `json:"longitude"`
		TeRminal interface{} `json:"te rminal"`
		ReservationType string `json:"reservationType"`
		ParkingTransaction []struct {
			ParkingStartTime int64 `json:"parkingStartTime"`
			ParkingEndTime int64 `json:"parkingEndTime"`
			Status string `json:"status"`
			Reason string `json:"reason"`
			PaymentTransactionS []struct {
				Status string `json:"status"`
				Created int `json:"created"`
				ChargeID string `json:"chargeId"`
				StArtTime int `json:"st artTime"`
				EndTime int `json:"endTime"`
				Reason interface{} `json:"reason"`
				PaymentType interface{} `json:"paymentType"`
				RefundID interface{} `json:"refundId"`
				BalanceTransactionID string `json:" balanceTransactionId"`
				ModifiedDate interface{} `json:"modifiedDate"`
				Last4 string `json:"last4"`
				CaRdID string `json:"ca rdId"`
			} `json:"paymentTransaction s"`
		} `json:"parkingTransaction"`
		BarcodeContent string `json:"barcodeContent"`
		CancelURL string `json:"cancelUrl"`
		DisplayID int `json:"displayId"`
		EventID int `json:"eventId"`
		LicensePlate interface{} `json:"licensePlate"`
		ResErvationStatus string `json:"res ervationStatus"`
		ReservationURL string `json:"reservationUrl"`
		SeAtgeekID int `json:"se atgeekId"`
		GettingHere string `json:"gettingHere"`
		Duration string `json:"duration"`
		State string `json:"state"`
		City string `json:"city"`
		BleEncBytes interface{} `json:"bleEncBytes"`
		RedemptionInstruCtions []struct {
			Text string `json:"text"`
			ID int `json:"id"`
			Position int `json:"position"`
			IllustrationID string `json:"illustration_id"`
			IllustratIonVersion string `json:"illustrat ion_version,omitempty"`
			IllustrationVersion string `json:"illustration_version,omitempty"`
		} `json:"redemptionInstru ctions"`
		CustomerID interface{} `json:"customerId"`
		MeterTypeID interface{} `json:"meterTypeId"`
		Timezone string `json:"timezone"`
		ParkingExtension interface{} `json:"parkingExtension"`
		Count int `json:"count"`
	} `json:"data"`
}


func SearchParkingLots(w http.ResponseWriter, req *http.Request) {
	fmt.Println("GET Request")
	
}


func ReserveParkingLots(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Post Request")
    params := mux.Vars(req)
	fmt.Println("ss",params)
	fmt.Println("req",req)
	
	decoder := json.NewDecoder(req.Body)
    var t ReserveParkingRequest
    err := decoder.Decode(&t)
    if err != nil {
        panic(err)
    }
    fmt.Println(t.CardID)
	
	u := ReserveParkingRequest{CardID: t.CardID, Email: t.Email ,Ends: t.Ends,
	FacilityID : t.FacilityID,ParkingExtension : t.ParkingExtension,Price : t.Price,ReservationID : t.ReservationID,RuleGroupID : t.RuleGroupID, 
	Starts : t.Starts, Timezone : t.Timezone, UUID : t.UUID }
	
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(u)
	request, err := http.NewRequest("POST", "https://smh-parkingservice.azurewebsites.net/parking/reserveParkingLots", b)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("customToken","eyJhbGciOiJIUzUxMiJ9.eyJqdGkiOiI0ZjMyY2ViYy02OTAwLTRkNTItOTgxNy1hNDVkNGM2ZGYwNjkifQ.eVBlIgU_keSO1DglRC8biCFU5iPW-riu5QtfOeyb_NU5AIq9fh-STeAwBzIn61oASL4AWx7h3EnE-1XLLsfGIg")
	request.Header.Set("uuid","5b0be623eee0690d609994ff")
	client := &http.Client{}
    response, err := client.Do(request)
	
	if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
	    fmt.Println(string(data))
		fmt.Println(err)
		
		parkingPriceResponse := ReserveParkingResponse{}
	    err := json.Unmarshal(data, &parkingPriceResponse)
	    if err != nil {
        fmt.Println(err)
	  }
		fmt.Println("JSON RESPONSE" , json.NewEncoder(w).Encode(parkingPriceResponse))
    }
}


func main() {
    router := mux.NewRouter()
    router.HandleFunc("/parkingLots", SearchParkingLots).Methods("GET")
    router.HandleFunc("/reserveParkingLots", ReserveParkingLots).Methods("POST")
   // router.HandleFunc("/people/{id}", DeletePersonEndpoint).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":12346", router))
}