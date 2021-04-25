package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://apim.int.expedia.com/hotels/listings/?imageSizes=y&room1.adults=2&thumbnailImageSize=s&contentDetails=low&unit=mi&links=API&allRoomTypes=false&radius=25&availOnly=true&hotelPrograms=20&checkIn=2021-11-10&geoLocation=36.083332%2C-115.166664&checkOut=2021-11-13", nil)
	req.Header.Add("Accept", "application/vnd.exp-hotel.v3+json")
	req.Header.Add("Key", os.Getenv("EXPEDIA_KEY"))
	req.Header.Add("Partner-Transaction-ID", os.Getenv("EXPEDIA_PARTNER_ID"))
	req.SetBasicAuth(os.Getenv("EXPEDIA_KEY"), os.Getenv("EXPEDIA_PASSWORD"))
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error %s", err)
	} else {
		defer resp.Body.Close()
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error %s", err)
		}
		var result map[string]interface{}
		json.Unmarshal(body, &result)
		count := result["Count"].(float64)
		fmt.Printf("Count: %f.0", count)
	}
}
