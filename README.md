This repository is to showcase the projects I have done while learning GO. I decided to do three projects to obtain proficiency.

1 - Simple Webscraper: This webscraper uses Colly, a webscraping library, to scrape weather data from the HTML of a weather website. It then stores this data in a CSV. While I did not use any GO-specific tools to create this (just the library) it was essential in my understanding of the language's syntax.

2 - 2048: I created the 2048 game which can be played via the command line using W,A,S,D keys. This is a relatively simple game, but I found it useful because it taught me about Structs in GO and recieving inputs from the user via command line. 

3 - REST API: There are four parts to a REST API: POST, GET, PUT, DELETE. While this is not the most useful API, I created it to show that I learned how to implement all four. I also created a front end to access and interact with this via localhost. This is a fake banking API where the user can register an account with a username, password, and balance amount. The user can then login with those credentials to view the balance. This is the POST and GET portions of the API - the user can POST(create) a user and GET(retrieve) their data. The user can also modify their balance after creating the account which would be PUT(update), or delete thier account which would be DETELE. This project was important in me learning about html, http protocol, debugging in GO, and using the developer tools in a web browser. 

Instructions to run:
1 - Webscraper and 2048: Can be run simply from basic GO commands.
2 - REST API: Use go run main.go (wait to see "Server running on localhost:8001) while in the REST API directory, then in a different terminal, run "python -m http.server 8000" to run the front end. Then navigate to
"http://localhost:8000" to interact with the API.

Gitlab CI/CD practice:
![Latest Temperature](https://img.shields.io/endpoint?url=https://gitlab.com/zachgreening/golanglearnv1/-/jobs/artifacts/main/raw/temperature_badge.json?job=scrape)
