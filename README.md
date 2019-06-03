# vehicle
Mango API: Vehicle

Vehicle is where you get the full information of a car.

## Run with Docker
* $ docker build -t avosa/vehicle:dev .
* $ docker rm VehicleDEV
* $ docker run -d -p 8098:8098 --network mango_net --name VehicleDEV avosa/vehicle:dev 
* $ docker logs VehicleDEV