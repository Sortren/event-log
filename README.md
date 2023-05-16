# Event Logger

#### Service that allows to send and store variety of events and retrieve them from the database layer.

### Technologies used:

<img align = "left" alt = "Go" width = "40px" src = "https://user-images.githubusercontent.com/79079000/177060907-76fcb9d0-2853-4c16-8f79-d6a29ea4689b.png" />
<img align = "left" alt = "Fiber" width = "40px" src = "https://user-images.githubusercontent.com/79079000/177061065-d67891e6-3087-4d83-b69a-ba9b9de53cec.png" />
<img align = "left" alt = "Gorm" width = "40px" src = "https://user-images.githubusercontent.com/79079000/177060995-f8247638-445d-4464-a312-5fa946e1d084.png" />
<img align = "left" alt = "Postgresql" width = "40px" src = "https://user-images.githubusercontent.com/79079000/177060938-76d1029d-aabc-459c-9538-4b2d14da3980.png" />
<img align = "left" alt = "Swagger" width = "40px" src = "https://user-images.githubusercontent.com/79079000/177061111-e305e35c-592d-4583-9e2a-09c1bb18d953.png" />
<img align = "left" alt = "http" width = "40px" src = "https://user-images.githubusercontent.com/79079000/177061166-a1277575-f625-4a38-90a3-86c1b654d651.png" />
<img align = "left" alt = "http" width = "40px" src = "https://user-images.githubusercontent.com/79079000/177061179-50a3adad-742b-48ba-b5da-b9a42f2f867d.png" />
<img align = "left" alt = "http" width = "40px" src = "https://user-images.githubusercontent.com/79079000/177060964-de848787-0509-4ed8-869b-7dcc14c8e0a0.png" />

<br />

---

### How to build the service

download deps to the local cache
```
go mod download
```

run docker-compose with a build switch

```
docker-compose up -d --build
```

run migrations

```
goose postgres "user=sortren password=sortren123 dbname=main sslmode=disable" up
```
(same credentials as in .env file)

restart event-log docker container

---

### General service information

The main endpoints are:
```
[GET]/api/v1/events?type=login&start=2022-07-03T21:00:00&end=2022-07-03T22:00:00&limit=10&offset=5
```
Which allows the user to retrieve the list of events that match the particular queryparams. More details are in the section "Detailed API Documentation" 

<br />

```
[POST]/api/v1/events
```
```json
{
  "type": "login",
  "description": "User has logged in to the auth service"
}
```

Which allows the user to send and store the event in the database via HTTP REST request


---

### Detailed API Documentation

I have used the Swagger (OpenAPI), to get auto-generated docs of controllers/endpoints <br />

To display the docs, after running up the server, get to the endpoint:

```
/api/v1/docs
```

You will see smth like this:
![image](https://user-images.githubusercontent.com/79079000/177061516-5a07ed29-5d4d-40f0-ae05-63bf659179d2.png) <br />

You can easily check what is the expected request body or what are the expected arguments of the particular endpoint <br />
![image](https://user-images.githubusercontent.com/79079000/177061563-a82ccb08-a007-4f1b-9714-34c0f6b51c92.png)
![image](https://user-images.githubusercontent.com/79079000/177061567-5f90c25a-f964-4032-b8ba-a224aa8152ec.png)

---

### Tests

If you want to test the REST API Layer, I have added the postman collection that includes every defined endpoint with sample data.

---
