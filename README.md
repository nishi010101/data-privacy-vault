# data-privacy-vault

A basic data privacy vault written in Go. Uses redis behind the scenes to store values, and use aes encryption algorithm.

### **How to Run:**

```REDIS_HOST_URL=${REDIS_HOST_URL} DATA_ENCRYPTION_KEY=${DATA_ENCRYPTION_KEY} go run main.go```

### **Supported Commands**

* **Tokenize**

POST ```/tokenize```

Sample payload
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

* **Detokenize**

POST ```/detokenize```

Sample payload

```
{
  "Id": "req-12345",
  "Data": {
    "field1": "token1",
    "field2": "token2",
    "fieldn": "tokenn"
  }
}
```


Both the endpoints requires operation specific API keys. 
