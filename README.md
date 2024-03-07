# Earthly Elixirs API

This is the API for Earthly Elixirs, an e-commerce store powered by Gin/Gonic and integrated with the Stripe API for payment processing.

## Features

- Create, read, update, and delete products
- Handle user authentication and authorization
- Process payments securely using the Stripe API

## Requirements

- Go (version 1.16 or higher)
- [Gin](https://github.com/gin-gonic/gin) (v1.7.4 or higher)
- [Stripe-Go](https://github.com/stripe/stripe-go) (v76 or higher)

## Installation

1. Clone this repository:

    ```bash
    git clone https://github.com/yourusername/earthly-elixirs-api.git
    ```

2. Navigate to the project directory:

    ```bash
    cd earthly-elixirs-api
    ```

3. Set up your environment variables:

    - `KEY`: Your Stripe API secret key

4. Install dependencies:

    ```bash
    go mod download
    ```

5. Build and run the server:

    ```bash
    go build
    ./earthly-elixirs-api
    ```

The server should now be running on `localhost:8080`.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the [MIT License](LICENSE).

---

Feel free to customize this template to fit the specific details of your project and its setup. Good luck with your project!
