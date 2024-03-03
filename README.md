# data-privacy-vault

A basic data privacy vault written in Go. Uses redis behind the scenes to store values, and use aes encryption algorithm.

### **How to Run:**

```REDIS_HOST_URL=${REDIS_HOST_URL} DATA_ENCRYPTION_KEY=${DATA_ENCRYPTION_KEY} go run main.go```

### **Supported Commands**

* **Tokenize**

POST ```/tokenize```

Sample request payload
```
{
	"id": "req-12345",
	"data": {
		"field1": "value1",
		"field2": "value2",
		"fieldn": "valuen"
	}
}
```

Sample Respose payload
```
{
  "id": "req-12345",
  "data": {
    "field1": "token1",
    "field2": "token2",
    "fieldn": "tokenn"
  }
}
```

* **Detokenize**

POST ```/detokenize```

Sample payload

```
{
  "id": "req-12345",
  "data": {
    "field1": "token1",
    "field2": "token2",
    "fieldn": "tokenn"
  }
}
```

```
{
    "field1": {
        "found": true,
        "value": "value1"
    },
    "field2": {
        "found": true,
        "value": "value2"
    },
    "fieldn": {
        "found": true,
        "value": "valuen"
    }
}
```

Both the endpoints require operation specific API keys. 
