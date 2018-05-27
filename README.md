# anaco-go-backend

# Setup
- Install go and set the $GOPATH https://golang.org/doc/install
- go get -u github.com/kardianos/govendor
- govendor sync
- govendor add +external

# To run
- go run main.go

## Planned features:
https://docs.google.com/document/d/1YWNL2iQOYzkZVnDpC5yVHeYbIxBTosDT1ThOAKoMJi0/edit?usp=sharing

##API server details:
- Used go and gin
- API URL - http://206.189.147.1:8080/
- Used token authentication
- API doc - https://docs.google.com/spreadsheets/d/1TyvrXztTUfDRLHvE9G6_heaTBV2c1xGh7NN7yZjSgOU/edit#gid=0
- Repo - https://github.com/ranjur/anaco-go-backend

## Main features
- As a user I can register into the system using sayonetech email address
- As a user I can login to the system using email and password
- As a logged in user I can view list of users who is already into the system
- As a logged in user I can change my profile pic
- As logged in user I can view the detail page of a user with image, name and anonymous comments for that user
- As a logged in user I can comment anonymously into a user's profile
- As a user I can like a comment of a user

## API's
- Register API
- Login API
- Users listing
- User detail
- User comment listing
- Create comment
- Comment like
- Comment dislike
- Update profile(pic)
