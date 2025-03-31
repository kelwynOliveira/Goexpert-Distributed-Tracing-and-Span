# Distributed Tracing and Span

Goexpert postgraduation project

## Lab project Description

> **Objective**: To develop a Go system that receives a postcode, identifies the city and returns the current climate (temperature in degrees celsius, fahrenheit and kelvin) along with the city. This system should implement OTEL (Open Telemetry) and Zipkin.
>
> Based on the known scenario "Temperature system by postcode" called Service B, a new project will be included, called Service A.
>
> ### Requirements - Service A (responsible for input):
>
> - The system must receive an 8-digit input via POST, using the schema: { "cep": "29902555" }
> - The system must validate that the input is valid (contains 8 digits) and is a STRING.
>   - If it is valid, it will be forwarded to Service B via HTTP
>   - If it is not valid, it must return:
>     - HTTP code: 422
>     - Message: invalid zipcode
>
> ### Requirements - Service B (responsible for orchestration):
>
> - The system must receive a valid 8-digit postcode
> - The system must search the postcode and find the name of the location, then return the temperatures and format them in: Celsius, Fahrenheit, Kelvin along with the location name.
>   The system should respond appropriately in the following scenarios:
>
>   - On success:
>     - HTTP code: 200
>     - Response Body: { "city: "São Paulo", "temp_C": 28.5, "temp_F": 28.5, "temp_K": 28.5 }
>   - In case of failure, if the postcode is not valid (with correct format):
>     - HTTP code: 422
>     - Message: invalid zipcode
>   - In case of failure, if the postcode is not found: - HTTP code: 404 - Message: can not find zipcode
>     After implementing the services, add the OTEL + Zipkin implementation:
>
> - Implement distributed tracing between Service A - Service B
> - Use span to measure the response time of the postcode search and temperature search services
>
> ### Tips:
>
> - Use the viaCEP API (or similar) to find the location you want to check the temperature: https://viacep.com.br/
> - Use the WeatherAPI API (or similar) to look up the desired temperatures: https://www.weatherapi.com/
> - To convert from Celsius to Fahrenheit, use the following formula: F = C \* 1.8 + 32
> - To convert from Celsius to Kelvin, use the following formula: K = C + 273
>   - Where F = Fahrenheit
>   - Where C = Celsius
>   - Where K = Kelvin
> - For questions about implementing OTEL, click here.
> - For the implementation of spans, you can click here
> - You will need to use an OTEL collector service
> - For more information about Zipkin, you can click here
>
> ### Delivery:
>
> - The complete source code of the implementation.
> - Documentation explaining how to run the project in a dev environment.
> - Use docker/docker-compose so we can test your application.

## How to run

Set your WeatherAPI APIKey in the WEATHER_API_KEY variable in the `.env` file.
Set the CEP in the Makefile

After updating the `.env` file run the command `make up` or `docker compose up -d` at the main folder.

Run `make run`or

- `curl -X POST -d '{"cep": "<CEP>"}' http://localhost:8080`
- `curl http://localhost:8081/<CEP>`

- Zipkin: `http://localhost:9411`
