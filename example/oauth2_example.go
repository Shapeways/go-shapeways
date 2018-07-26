package main

// import "net/http"
// import "io/ioutil"
import "fmt"
import "../shapeways"

func main() {
  client := shapeways.NewClient("YOUR CLIENT KEY", "YOUR CLIENT SECRET")
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
