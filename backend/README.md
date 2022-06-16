
# pAPI API Application

This application connects to various public API's that are available around the internet and adds them to a database for easy access. This means that the data from various API's are all available from a single location in a standard format.

## Endpoints

- Get all categories
- Get single category
- Get all remote API's in database
- Get single remote API in database
- Update single response from a single API in database

## Technologies

- [MongoDB](https://www.mongodb.com/) as the database. Easy and scalable document database solution.
- [Go](https://go.dev/) as the core back-end language. Fast, modern, and concurrent programming language.
- [Gin](https://gin-gonic.com/) as the Go framework. Fully functional web framework that makes routing a breeze.
- [Redis](https://redis.io/) as an in-memory data store. Leveraged by the Tollbooth middleware to provide a simple rate limiter.
- [Swagger](https://swagger.io/) used for API documentation and testing, specifically the implementation provided by [Swaggo](https://github.com/swaggo/swag). 
- [Postman](https://www.postman.com/) was used for debugging and testing during development.

## Security

- Only allow one request every 10 seconds from a single IP address to a single endpoint to prevent spam
- Require admin login details to access the database, either from the Mongo web interface, or the supplied admin application.

## Architecture

- 3-tier structure (controller, service, data) to provide code maintainability and scalability, by making code more manageable due to separation of concerns.
- Using a document database because it suits the needs of our application - flexible input without near-perfect error checking.
- Built using Go modules, with multiple packages that provide the logic for different parts of the application. This makes the code more manageable, and follows Go's recommended best practices.

## Deployment

- App hosted on Heroku free tier. This is perfect for proof of concept/MVP, and is easily scalable if the demand for the application grows.
- Database hosted in MongoDB Atlas. Similar to Heroku, this is perfect for proof of concept/MVP, and is easily scalable if the demand for the application grows.

## Testing

Due to the scope of the application being quite small, white-box tests have been temporarily shelved. It was deemed more important that the black-box tests pass, as that means that the front-end application will be receiving the correct response to any requests which will ensure proper functionality.

For black-box testing, a combination of Postman and Swagger was used. Swagger tests utilise the OpenAPI standard to check whether a HTTP code and body response matches what is expected. For Postman, custom requests can be made to check a response, with detailed information such as Headers and Authorization information. Custom tests can be written in Javascript if required. Visit our [application](https://papi-project.herokuapp.com) to see our Swagger UI at work. Alternatively, download `pAPI.postman_collection.js` from our [docs](docs) and import it into your own local Postman application to easily view the endpoints there instead.

## Future Considerations

Following are some possible enhancements that can be made to the application in future releases.

- Static IP address for application | Currently, since the app is hosted on a free tier in Heroku, the IP address is constantly changing as the application is spun up and taken down on various cloud servers. This means that database access has to allow access from anywhere, which is a security concern. In future, if the app is hosted at a static location, we could set that as the only location that can query the database.
- Code snippets | Currently we only display the unformatted output from various public API's. To provide a more positive user experience, snippets for how this API can be used should be provided in the database, for various programming languages.
- Documentation | Currently, the app is lacking sufficient code annotations. This should be remedied so that the application remains maintainable even if another programmer was to take over the development of the application. It would also make it easier to develop white-box tests in the future if the application grows in scale.
