package main

type CollectRequest struct {
	Star int `form:"star" validation:"gte=1,lte=5" doc:"formData"`
}
