# Worker Compose

Here's the documentation for the `worker-compose.yml` file:

### Docker Compose Configuration Documentation

The `worker-compose.yml` file defines the configuration for orchestrating Docker containers to set up the application environment.

#### Services

**db**

* **Image**: PostgreSQL version 13 is used as the database service.
* **Volumes**: The `./postgres_datas` directory is mounted to `/var/lib/postgresql/data/` within the container to persist data.
* **Restart Policy**: The container is set to restart always in case of failure.
* **Environment Variables**: Environment variables are loaded from the `.env` file to configure the PostgreSQL service.
* **Ports**: Port `5432` on the host machine is mapped to port `5432` on the container for database access.
* **Networks**: The service is connected to the `my-network` network for communication with other services.

#### Volumes

* **postgres\_datas**: An external volume is defined to persist PostgreSQL data outside of container instances.

#### Networks

* **my-network**: A custom network named `my-network` is defined for inter-container communication, allowing services to communicate with each other securely.

#### Conclusion

The `worker-compose.yml` file specifies the configuration for running the PostgreSQL database service required by the application. By defining services, volumes, and networks, developers can easily manage the application environment using Docker Compose, ensuring consistency and scalability across different deployment environments.
