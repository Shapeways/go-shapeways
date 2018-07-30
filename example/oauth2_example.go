package main
import "fmt"
import "../shapeways_oauth2/"

func main() {
  client := shapeways_oauth2.NewClient("CLIENT_KEY", "CLIENT_SECRET")

  // Authenticate client
  status, err := client.Authenticate()
  if err != nil {
    fmt.Println(status)
    panic(err)
  }

  // Get list of materials
  client.GetMaterials()

  // Get Single Material
  material, err := client.GetMaterial(6)
  if err != nil {
    panic(err)
  }
  fmt.Println(material.Title)

  // Upload Model
  client.UploadModel("/path/to/cube.stl")
}
