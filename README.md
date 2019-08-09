# Faas demo project

This is a demo project to show how a faas gateway works in a docker-compose base faas cluster.

## Project structure and tools

Given the hint, the project follows the structure in the instruction and use fswatch together with govendor for development environment.

## Gateway

This is the main part of the project, it distributes the traffic from outside to different internal services according to the `sub-path` in the request url. As the time limitation, currently the service only support one end point each service, but it is possible to add support to be able to expose multiple end points.

### designs

Gateway service has two parts, one is mainly for maintaining the route table while the other part is for proxying the request to the correct service.

For the route table maintainer part, it runs as an independent go routine with a cron job to periodically query the docker cluster for service data and update its internal route table. The current update interval is set to 5 seconds.

The proxy part receives the request from outside, first parse the url and find the sub-path for the service, then get the service ip from the route table. After get everything just proxy the request to the correct service.

## Services

I initially created 3 services for this demo, but one is the ping service which I met huge privilege problem when running it in docker, and removed it from the repo in the late stage.

### Factorial

This is a service for calculating factorial result for the input number, only positive integer is allowed. Url pattern looks like below
`/mathoperations?number={positive integer}`

### Current weather

This service accept city name as a parameter and then send request to `openweather`(https://openweathermap.org/) public service to query the current weather for the input city and return it. Url pattern looks as following

`/currentweather?cityname={city name}`

## How to run and test it

As the project is mainly based on docker-compose, we will need that to run it with that.

### start

Go to the `modularfinancefaas` folder and run the following command
`docker compose up`

### test the function

After it is properly started, hit the following url in the browser for testing the functions

- http://localhost/gateway/mathoperations?number=1
- http://localhost/gateway/currentweather?cityname=london

## Improvements and future work

- Add support to serve multiple entry points for internal services
- Add a service to demo HTTP POST works(it works but all service use HTTP GET)
- Add more tests to it