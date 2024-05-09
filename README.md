# test-public-api

## Preview info 
Api for testing https://rapidapi.com/freshdata-freshdata-default/api/fresh-linkedin-profile-data/


## strucure 

responses.go - structure and metods
client.go - Client for request and json parse

## Usage

init

```go
myClient, err := testapi.NewClient(time.Second * 10)
```

get posts

```go
posts, err := myClient.GetPosts()
```

get recomendation for api
```go
profile_recomendation, err := myClient.GetRecomendations(profile_link)
```

## auth

Auth work by header. Headers setUp in CustomTransport RoundTrip