### Code level Design Patterns: 

* **SRP (Single Responsibility Principle)-** Each class or an entity should have a single responsibility and it should not deflect from the behaviour.

* **Open Closed Principle-** Entities should be open for extension but not for modification.

    ```go
    type Connector interface {
        Initialize() (*Connector, error)
        Connect()
    }

    type MongoConnector struct {
        // Data members
    }

    type PostgresConnector struct {
        // Data members
    }

    func(mc *MongoConnector) Initialize() (*MongoConnector, error) {
        return nil, nil
    }

    func (pc *PostgresConnector) Initialize() (*PostgresConnector, error) {
        return nil, nil
    }

    func main() {
    }
    ```
* **Liskov Substitution Principle-** Subtypes must be suitable for their base types without breaking behaviour. (Use the same example as above)

* **Interface Segregarion Principle-** Clients should not be dependent on interfaces which they don't use. *Pro tip* - 

* **Dependency Inversion Principle-***


### Service level Design Pattern:


### Designing Passwordless Authentication:
```go
type DeviceType string

const MAC DeviceType = "mac"
const ANDROID DeviceType = "android"
const IOS DeviceType = "ios"
const WEB DeviceType = "web"

type User struct {
	Id     uint   `json:"id"`
	Email  string `json:"email"`
	Active bool   `json:"active"`
}

type Device struct {
	Id         uint        `json:"id"`
	DeviceType *DeviceType `json:"deviceType"`
}

type DeviceMetadata struct {
	DeviceId   uint   `json:"deviceId"`
	PrivateKey string `json:"privateKey"`
	PublicKey  string `json:"publicKey"`
}
```

- To set up the Passwordless authentication - Client and Server will generate random private keys and based on the scaler multiplication of the curve you generate the public key.

- Exchange the Public key of the client to the server and the server to the client, compute the secret based on the opposite.

- If the secret matches, register the device for passwordless authentication. Store the private key in a secure storage for the devices.


