# Overview
Pkge is a simple library for working with 
the api service [pkge doc](https://business.pkge.net/ru/docs). 
The library supports all methods that are available in the api.
#### Available methods is:
- Shipping service list
- Shipping service definition
- Adding a package 
- Updating package 
- Info about package
- Modifying package info
- Deleting package
- List of packages 
___ 
# Get started 
```go
package main

import (
	"log"
	"os"

	pkge "github.com/Apolisk/pkge/pkg"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	apiKey := os.Getenv("TOKEN")
	client := pkge.New(apiKey)
	data, err := client.ActivatedDeliveryServices()
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("Result: %v", data)
}

```
# Package status code 

| Value | Description|
|-------|---------------------|
| 0     | Package added to the system but not yet updated|
| 1     | Initial update of the package is in progress|
| 2     | The update of the package was completed but no information was received from the delivery services|
| 3     | Package on the way|
| 4     | Package delivered to delivery place|
| 5     | Package delivered to the recipient|
| 6     | Delivery failed. For example, failed delivery attempt|
| 7     | Package delivery failed. For example, when destroying a Package|
| 8     | The Package is ready to be sent. The information about the Package was received by the delivery service, but the Package has not yet been sent|
| 9     | End of the tracing path of a Package. For example, for international Packages that are tracked only when in the country of departure|
