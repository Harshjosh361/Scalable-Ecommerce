
# Scalable Microservice eCommerce API

## Table of Contents
- [Project Overview](#project-overview)
- [Technologies Used](#technologies-used)
- [Project Setup](#project-setup)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [Contributing](#contributing)
- [License](#license)

## Project Overview
This eCommerce API project provides a scalable backend for managing an eCommerce platform, including user management, product catalog, orders, and payments. It utilizes a microservices architecture with an NGINX API Gateway for request routing and is containerized using Docker for easy deployment. The project also integrates Keploy for testing API functionalities, ensuring reliable responses.
- **Microservices**: Each major feature (user management, products, orders, etc.) is a separate service.
- **API Gateway**: An NGINX-based gateway to route requests to the respective microservices.
- **Database**: MongoDB is used for different microservices.
- **Testing**: Integrated with Keploy to capture and test API calls for accuracy.

## Technologies Used
- **Programming Languages**: Go
- **Database**: MongoDB
- **Containers**: Docker
- **API Gateway**: NGINX
- **Testing**: Keploy for API testing
- **Other Tools**: Twilio for notifications, Stripe for Payment

## Project Setup

1. **Clone the repository**:
    ```bash
    git clone https://github.com/Harshjosh361/Scalable-Ecommerce.git
    cd Scalable-Ecommerce
    ```

2. **Environment Variables**: 
   Create a `.env` file for each microservice and add necessary environment variables like database URIs, Twilio credentials, etc.

3. **Docker Setup**:
    To run individual services, navigate to the respective service folder and run:
   ```bash
   cd user-service
   docker compose up
   ```
   Or for another service:
   ```bash
   cd product-service
   docker compose up
   ```

   To run all services together, navigate to the root:
   ```bash
   docker compose up
   ```
   
## Usage

- **Starting the Application**: Once all services and the API Gateway are up, the application will be accessible via the configured API Gateway port.
- **Connecting to Databases**: MongoDB database is connected to their respective services via environment variables.

endpoints?

## Testing
This project uses Keploy for API testing. To run tests:
1. Ensure Keploy is configured in each service’s `keploy.yaml` file.
2. recording test case
   
   ```bash
   keploy record -c "docker compose up" --container-name "my-app-container"
   ```
3. testing
   
   ```bash
   keploy test -c "docker compose up" --container-name "my-app-container"
   ```
   Replace my-app-container with the original container name

## Contributing
Contributions are welcome! Please follow these steps:
1. Fork the repository.
2. Create a new branch for your feature (`git checkout -b feature-branch`).
3. Commit your changes (`git commit -m "Add feature"`).
4. Push to your fork (`git push origin feature-branch`).
5. Open a pull request.

## License
This project is licensed under the terms of the [MIT License](LICENSE).
