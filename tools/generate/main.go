// generate_public_api produces the thin Go wire models/client from the maintained
// OpenAPI bundle. It uses only the standard library; normal builds run in Docker.
package main

import (
	"encoding/json"
	"fmt"
	"go/format"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

type schema struct {
	Ref        string            `json:"$ref"`
	Type       string            `json:"type"`
	Properties map[string]schema `json:"properties"`
	Required   []string          `json:"required"`
	Items      *schema           `json:"items"`
	OneOf      []schema          `json:"oneOf"`
	AllOf      []schema          `json:"allOf"`
}
type operation struct {
	ID          string `json:"operationId"`
	RequestBody struct {
		Content map[string]struct {
			Schema schema `json:"schema"`
		} `json:"content"`
	} `json:"requestBody"`
	Responses map[string]struct {
		Content map[string]struct {
			Schema schema `json:"schema"`
		} `json:"content"`
	} `json:"responses"`
}

func name(s string) string {
	var b strings.Builder
	for _, p := range strings.FieldsFunc(s, func(r rune) bool { return r == '_' || r == '-' }) {
		if p != "" {
			b.WriteString(strings.ToUpper(p[:1]) + p[1:])
		}
	}
	return b.String()
}
func kind(s schema) string {
	if s.Ref != "" {
		return strings.TrimPrefix(s.Ref, "#/components/schemas/")
	}
	for _, part := range s.AllOf {
		if part.Ref != "" {
			return kind(part)
		}
	}
	switch s.Type {
	case "string":
		return "string"
	case "integer":
		return "int64"
	case "boolean":
		return "bool"
	case "array":
		return "[]" + kind(*s.Items)
	case "object":
		return "map[string]json.RawMessage"
	}
	return "json.RawMessage"
}
func keys[T any](m map[string]T) []string {
	a := make([]string, 0, len(m))
	for k := range m {
		a = append(a, k)
	}
	sort.Strings(a)
	return a
}
func main() {
	if len(os.Args) != 3 {
		panic("usage: generate_public_api schema.lock.json output-directory")
	}
	raw, err := loadSchema(os.Args[1])
	if err != nil {
		panic(err)
	}
	var spec struct {
		Components struct {
			Schemas map[string]schema `json:"schemas"`
		} `json:"components"`
		Paths map[string]map[string]operation `json:"paths"`
	}
	if err = json.Unmarshal(raw, &spec); err != nil {
		panic(err)
	}
	var b strings.Builder
	b.WriteString("// Code generated from the maintained OpenAPI contract; DO NOT EDIT.\npackage generated\nimport(\"context\";\"encoding/json\";\"net/http\";\"net/url\";\"strings\";\"github.com/arctex-one/client-go/transport\")\n")
	for _, n := range keys(spec.Components.Schemas) {
		s := spec.Components.Schemas[n]
		if len(s.Properties) == 0 {
			fmt.Fprintf(&b, "type %s = %s\n", n, kind(s))
			continue
		}
		req := map[string]bool{}
		for _, x := range s.Required {
			req[x] = true
		}
		fmt.Fprintf(&b, "type %s struct {\n", n)
		for _, p := range keys(s.Properties) {
			typ := kind(s.Properties[p])
			if !req[p] {
				typ = "*" + typ
			}
			tag := p
			if !req[p] {
				tag += ",omitempty"
			}
			fmt.Fprintf(&b, "%s %s `json:%q`\n", name(p), typ, tag)
		}
		b.WriteString("AdditionalFields map[string]json.RawMessage `json:\"-\"`\n}\n")
		fmt.Fprintf(&b, "func(v *%s)UnmarshalJSON(raw []byte)error{type plain %s;var p plain;if err:=json.Unmarshal(raw,&p);err!=nil{return err};var extra map[string]json.RawMessage;if err:=json.Unmarshal(raw,&extra);err!=nil{return err};", n, n)
		for _, p := range keys(s.Properties) {
			fmt.Fprintf(&b, "delete(extra,%q);", p)
		}
		fmt.Fprintf(&b, "p.AdditionalFields=extra;*v=%s(p);return nil}\n", n)
		fmt.Fprintf(&b, "func(v %s)MarshalJSON()([]byte,error){type plain %s;raw,err:=json.Marshal(plain(v));if err!=nil{return nil,err};var out map[string]json.RawMessage;if err=json.Unmarshal(raw,&out);err!=nil{return nil,err};for k,value:=range v.AdditionalFields{switch k{", n, n)
		for _, p := range keys(s.Properties) {
			fmt.Fprintf(&b, "case %q:continue;", p)
		}
		b.WriteString("};if _,exists:=out[k];!exists{out[k]=value}};return json.Marshal(out)}\n")
	}
	b.WriteString(`
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
`)
	for _, path := range keys(spec.Paths) {
		for _, method := range keys(spec.Paths[path]) {
			op := spec.Paths[path][method]
			if op.ID == "" {
				continue
			}
			out := "json.RawMessage"
			for _, code := range []string{"200", "201", "202"} {
				if content, ok := op.Responses[code]; ok {
					out = kind(content.Content["application/json"].Schema)
					break
				}
			}
			input := ""
			if method == "post" {
				input = kind(op.RequestBody.Content["application/json"].Schema)
			}
			params := []string{}
			for _, part := range strings.Split(path, "/") {
				if strings.HasPrefix(part, "{") {
					params = append(params, strings.Trim(part, "{}"))
				}
			}
			fmt.Fprintf(&b, "func(c *Client)%s(ctx context.Context,", name(op.ID))
			for _, p := range params {
				fmt.Fprintf(&b, "%s string,", name(p))
			}
			b.WriteString("query url.Values,")
			if input != "" {
				fmt.Fprintf(&b, "key string,body %s,", input)
			}
			fmt.Fprintf(&b, ")(*%s,error){path:=%q;", out, path)
			for _, p := range params {
				fmt.Fprintf(&b, "path=strings.ReplaceAll(path,%q,url.PathEscape(%s));", "{"+p+"}", name(p))
			}
			body := "nil"
			if input != "" {
				body = "body"
			}
			key := `""`
			if input != "" {
				key = "key"
			}
			fmt.Fprintf(&b, "out:=new(%s);err:=c.call(ctx,%q,path,query,%s,%s,out);return out,err}\n", out, strings.ToUpper(method), key, body)
		}
	}
	source, err := format.Source([]byte(b.String()))
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
