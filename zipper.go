package main;

import(
  "flag"
  "fmt"
  "zipper/google"
  "zipper/bitdo"
  "zipper/isgd"
  "zipper/vgd"
)

func main(){

  var s = flag.String("s", "google", "specify service to use")
  var u = flag.String("u", "http://xyz.com", "specify url to shorten")
  flag.Parse()

  switch *s {
    case "google":
      resp, err := google.Shorten(*u)
      if err != nil{
        panic(err)
      }
      fmt.Println(resp)
    case "bitdo":
      respb, errr := bitdo.Shorten(*u)
      if errr != nil{
        panic(errr)
      }
      fmt.Println(respb)
    case "isgd":
      isgdr, isgderr := isgd.Shorten(*u)
      if isgderr != nil{
        panic(isgderr)
      }
      fmt.Println(isgdr)
    case "vgd":
      vgdr, vgderr := vgd.Shorten(*u)
      if vgderr != nil{
        panic(vgderr)
      }
      fmt.Println(vgdr)
    default:
      resp, err := google.Shorten(*u)
      if err != nil{
        panic(err)
      }
      fmt.Println(resp)
  }
}