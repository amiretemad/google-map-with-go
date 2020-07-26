# Google map APIs as microservice
Useful google map APIs as a microservice (Distance calculator and ...)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes. See deployment for notes on how to deploy the project on a live system.

### Samples 

Calculating distance between 2 location you can send latitude and longitude of points to  
http://localhost:8083/distance service.

**Request**
```
{
  "origin": {
    "lat": 29.5926,
    "lng": 52.5836
  },
  "destination": {
    "lat": 41.0082,
    "lng": 28.9784
  }
}
```

**Response**
```
{
    "origin_addresses": "Fars Province, Shiraz, Azadegan, Iran",
    "destination_addresses": "Cankurtaran, Alemdar Cd., 34110 Fatih/İstanbul, Turkey",
    "distance_km": "3,203 km",
    "distance_meters": 3202789,
    "success": true
}
```
