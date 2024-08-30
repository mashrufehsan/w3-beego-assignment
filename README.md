# Cat API Web Application

## Introduction
This project is a web application that displays cat images and information, utilizing [The Cat API](https://thecatapi.com). The backend is built with the Beego framework in Go, handling API requests, configuration management, and template rendering.

## Features

- **Favorite Cat Images:** Add or remove cat images from your favorites list, allowing you to keep track of your preferred images.
- **Voting System:** Upvote or downvote cat images to express your preferences.
- **Browse by Breed:** Explore cat images categorized by breeds, and view detailed information about each breed.
- **View Voted Images:** Access a personalized list of images you've upvoted or downvoted for easy reference.

## Installation ##

### Prerequisites ###
- Go

    >👉 https://go.dev/doc/install

- Beego framework
    ```bash
    go install github.com/beego/beego/v2@latest
    ```
- Bee command-line tool
    ```bash
    go install github.com/beego/bee/v2@latest
    ```
    Then  add bee binary to PATH environment variable in your ~/.bashrc or ~/.bash_profile file:
    ```bash
    export PATH=$PATH:<your_main_gopath>/bin
    ```

### Setup Instructions ###

1. **Clone and navigate to the the repository:**
    ```bash
    git clone https://github.com/mashrufehsan/w3-beego-assignment.git
    cd w3-beego-assignment
    ```
2. **Configure environment variables:**

    - Copy the `app.conf.sample` file in the `conf` folder to `app.conf`.
        ```bash
        cp conf/app.conf.sample conf/app.conf
        ```
    - Update `app.conf` with the Cat API key and other necessary configurations.
        ```ini
        appname   = Name of the application
        httpport  = Port to run the application
        runmode   = Mode to run the application (e.g., dev, prod)
        CatAPIKey = Your Cat API Key
        SubID = An ID to identify a user (e.g., mashruf, shanto)
        ```
3. **Install dependencies:**
    ```bash
    go mod tidy
    ```
4. **Start the Beego server:**
    ```
    bee run
    ```
5. **Access the application in your browser:**
    >👉 http://localhost:8080

## Notes

- **Configuration:** Ensure you replace placeholders in the `app.conf` file with actual values, such as your Cat API key.
