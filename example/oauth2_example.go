package main

// import "net/http"
// import "io/ioutil"
import "fmt"
import "../shapeways"

func main() {
  client := shapeways.NewClient("ENPbso1i8tyUWX2tmMhHK5pm8pcHN9cauV96CVALkaIQUa3qQr", "j7aEPlq83n9UtU0t3u1dg6Gk1Bw1CwU4TsTosKTjTSwl3GVQ0y")
  // client := shapeways.NewClient("iOUMAH66yc67yPOIUWn9vLFTtFeVyjAIXclUbq2vNCcYf1Cbnf", "O7O5VPXeIQStWjX6OmeGlrVnCkpzUbuAY4jnoYLZGPjYQaLQr0")
  status, err := client.Authenticate()
  if err != nil {
    fmt.Println(status)
    panic(err)
  }
  // material, err := client.GetMaterial(6)
  // if err != nil {
  //   panic(err)
  // }
  // fmt.Println(material.Title)
  // client.GetMaterials()
  client.UploadModel("/home/matt/code/project_euler/random_crap/cube.stl")
}
