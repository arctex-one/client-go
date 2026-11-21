// generate_public_api generates build-time Go bindings for the Arctex API.
package main

import (
	"os"
	"path/filepath"

	"github.com/arctex-one/api-protocol-go/wiregen"
)

func main() {
	if len(os.Args) != 3 {
		panic("usage: generate_public_api schema.lock.json output-directory")
	}
	document, err := loadSchema(os.Args[1])
	if err != nil {
		panic(err)
	}
	source, err := wiregen.Generate(document, wiregen.Options{
		ClientImports: []string{"github.com/arctex-one/client-go/transport"},
		ClientSource:  clientSource,
	})
	if err != nil {
		panic(err)
	}
	if err = os.MkdirAll(os.Args[2], 0755); err != nil {
		panic(err)
	}
	if err = os.WriteFile(filepath.Join(os.Args[2], "api.gen.go"), source, 0644); err != nil {
		panic(err)
	}
}

const clientSource = `
type Client struct {BaseURL string; HTTP *http.Client; Bearer string}
type ResponseError = transport.ResponseError
var ErrResponseTooLarge = transport.ErrResponseTooLarge
func(c *Client)call(ctx context.Context,method,path string,query url.Values,key string,body any,out any)error {
 var raw json.RawMessage
 if body!=nil {var err error;raw,err=json.Marshal(body);if err!=nil{return err}}
 client:=http.Client{}
 if c.HTTP!=nil {client=*c.HTTP}
 boundary,err:=transport.New(c.BaseURL,client)
 if err!=nil{return err}
 if len(query)>0 {path+="?"+query.Encode()}
 response,err:=boundary.Call(ctx,c.Bearer,method,path,key,raw)
 if err!=nil{return err}
 return json.Unmarshal(response,out)
}
`
