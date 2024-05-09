package main

import (
	"fmt"
	"log"
	testapi "test-api/test-api"
	"time"
)

// https://docs.greynoise.io/reference/get_v3-community-ip

// Us1HQNYLlep4ibU1rPIvIEbp6j781UyCASaLpQFYRJ4eVdyZ7lveOwch3Mt55vdW
func main() {
	myClient, err := testapi.NewClient(time.Second * 10)
	if err != nil {
		log.Fatal(err.Error())
	}

	posts, err := myClient.GetPosts()

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, post := range posts.Items {
		fmt.Println(post.Info())
	}

	profile_link := "https://www.linkedin.com/in/ajjames"
	profile_recomendation, err := myClient.GetRecomendations(profile_link)

	if err != nil {
		log.Fatal(err.Error())
	}

	res, err := profile_recomendation.Info()

	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(res)
}
