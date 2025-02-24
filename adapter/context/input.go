// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package context

import (
	"github.com/tianmaotalk/beego/v2/server/web/context"
)

// BeegoInput operates the http request header, data, cookie and body.
// it also contains router params and current session.
type BeegoInput context.BeegoInput

// NewInput return BeegoInput generated by Context.
func NewInput() *BeegoInput {
	return (*BeegoInput)(context.NewInput())
}

// Reset init the BeegoInput
func (input *BeegoInput) Reset(ctx *Context) {
	(*context.BeegoInput)(input).Reset((*context.Context)(ctx))
}

// Protocol returns request protocol name, such as HTTP/1.1 .
func (input *BeegoInput) Protocol() string {
	return (*context.BeegoInput)(input).Protocol()
}

// URI returns full request url with query string, fragment.
func (input *BeegoInput) URI() string {
	return input.Context.Request.RequestURI
}

// URL returns request url path (without query string, fragment).
func (input *BeegoInput) URL() string {
	return (*context.BeegoInput)(input).URL()
}

// Site returns base site url as scheme://domain type.
func (input *BeegoInput) Site() string {
	return (*context.BeegoInput)(input).Site()
}

// Scheme returns request scheme as "http" or "https".
func (input *BeegoInput) Scheme() string {
	return (*context.BeegoInput)(input).Scheme()
}

// Domain returns host name.
// Alias of Host method.
func (input *BeegoInput) Domain() string {
	return (*context.BeegoInput)(input).Domain()
}

// Host returns host name.
// if no host info in request, return localhost.
func (input *BeegoInput) Host() string {
	return (*context.BeegoInput)(input).Host()
}

// Method returns http request method.
func (input *BeegoInput) Method() string {
	return (*context.BeegoInput)(input).Method()
}

// Is returns boolean of this request is on given method, such as Is("POST").
func (input *BeegoInput) Is(method string) bool {
	return (*context.BeegoInput)(input).Is(method)
}

// IsGet Is this a GET method request?
func (input *BeegoInput) IsGet() bool {
	return (*context.BeegoInput)(input).IsGet()
}

// IsPost Is this a POST method request?
func (input *BeegoInput) IsPost() bool {
	return (*context.BeegoInput)(input).IsPost()
}

// IsHead Is this a Head method request?
func (input *BeegoInput) IsHead() bool {
	return (*context.BeegoInput)(input).IsHead()
}

// IsOptions Is this an OPTIONS method request?
func (input *BeegoInput) IsOptions() bool {
	return (*context.BeegoInput)(input).IsOptions()
}

// IsPut Is this a PUT method request?
func (input *BeegoInput) IsPut() bool {
	return (*context.BeegoInput)(input).IsPut()
}

// IsDelete Is this a DELETE method request?
func (input *BeegoInput) IsDelete() bool {
	return (*context.BeegoInput)(input).IsDelete()
}

// IsPatch Is this a PATCH method request?
func (input *BeegoInput) IsPatch() bool {
	return (*context.BeegoInput)(input).IsPatch()
}

// IsAjax returns boolean of this request is generated by ajax.
func (input *BeegoInput) IsAjax() bool {
	return (*context.BeegoInput)(input).IsAjax()
}

// IsSecure returns boolean of this request is in https.
func (input *BeegoInput) IsSecure() bool {
	return (*context.BeegoInput)(input).IsSecure()
}

// IsWebsocket returns boolean of this request is in webSocket.
func (input *BeegoInput) IsWebsocket() bool {
	return (*context.BeegoInput)(input).IsWebsocket()
}

// IsUpload returns boolean of whether file uploads in this request or not..
func (input *BeegoInput) IsUpload() bool {
	return (*context.BeegoInput)(input).IsUpload()
}

// AcceptsHTML Checks if request accepts html response
func (input *BeegoInput) AcceptsHTML() bool {
	return (*context.BeegoInput)(input).AcceptsHTML()
}

// AcceptsXML Checks if request accepts xml response
func (input *BeegoInput) AcceptsXML() bool {
	return (*context.BeegoInput)(input).AcceptsXML()
}

// AcceptsJSON Checks if request accepts json response
func (input *BeegoInput) AcceptsJSON() bool {
	return (*context.BeegoInput)(input).AcceptsJSON()
}

// AcceptsYAML Checks if request accepts json response
func (input *BeegoInput) AcceptsYAML() bool {
	return (*context.BeegoInput)(input).AcceptsYAML()
}

// IP returns request client ip.
// if in proxy, return first proxy id.
// if error, return RemoteAddr.
func (input *BeegoInput) IP() string {
	return (*context.BeegoInput)(input).IP()
}

// Proxy returns proxy client ips slice.
func (input *BeegoInput) Proxy() []string {
	return (*context.BeegoInput)(input).Proxy()
}

// Referer returns http referer header.
func (input *BeegoInput) Referer() string {
	return (*context.BeegoInput)(input).Referer()
}

// Refer returns http referer header.
func (input *BeegoInput) Refer() string {
	return (*context.BeegoInput)(input).Refer()
}

// SubDomains returns sub domain string.
// if aa.bb.domain.com, returns aa.bb .
func (input *BeegoInput) SubDomains() string {
	return (*context.BeegoInput)(input).SubDomains()
}

// Port returns request client port.
// when error or empty, return 80.
func (input *BeegoInput) Port() int {
	return (*context.BeegoInput)(input).Port()
}

// UserAgent returns request client user agent string.
func (input *BeegoInput) UserAgent() string {
	return (*context.BeegoInput)(input).UserAgent()
}

// ParamsLen return the length of the params
func (input *BeegoInput) ParamsLen() int {
	return (*context.BeegoInput)(input).ParamsLen()
}

// Param returns router param by a given key.
func (input *BeegoInput) Param(key string) string {
	return (*context.BeegoInput)(input).Param(key)
}

// Params returns the map[key]value.
func (input *BeegoInput) Params() map[string]string {
	return (*context.BeegoInput)(input).Params()
}

// SetParam will set the param with key and value
func (input *BeegoInput) SetParam(key, val string) {
	(*context.BeegoInput)(input).SetParam(key, val)
}

// ResetParams clears any of the input's Params
// This function is used to clear parameters so they may be reset between filter
// passes.
func (input *BeegoInput) ResetParams() {
	(*context.BeegoInput)(input).ResetParams()
}

// Query returns input data item string by a given string.
func (input *BeegoInput) Query(key string) string {
	return (*context.BeegoInput)(input).Query(key)
}

// Header returns request header item string by a given string.
// if non-existed, return empty string.
func (input *BeegoInput) Header(key string) string {
	return (*context.BeegoInput)(input).Header(key)
}

// Cookie returns request cookie item string by a given key.
// if non-existed, return empty string.
func (input *BeegoInput) Cookie(key string) string {
	return (*context.BeegoInput)(input).Cookie(key)
}

// Session returns current session item value by a given key.
// if non-existed, return nil.
func (input *BeegoInput) Session(key interface{}) interface{} {
	return (*context.BeegoInput)(input).Session(key)
}

// CopyBody returns the raw request body data as bytes.
func (input *BeegoInput) CopyBody(MaxMemory int64) []byte {
	return (*context.BeegoInput)(input).CopyBody(MaxMemory)
}

// Data return the implicit data in the input
func (input *BeegoInput) Data() map[interface{}]interface{} {
	return (*context.BeegoInput)(input).Data()
}

// GetData returns the stored data in this context.
func (input *BeegoInput) GetData(key interface{}) interface{} {
	return (*context.BeegoInput)(input).GetData(key)
}

// SetData stores data with given key in this context.
// This data are only available in this context.
func (input *BeegoInput) SetData(key, val interface{}) {
	(*context.BeegoInput)(input).SetData(key, val)
}

// ParseFormOrMulitForm parseForm or parseMultiForm based on Content-type
func (input *BeegoInput) ParseFormOrMulitForm(maxMemory int64) error {
	return (*context.BeegoInput)(input).ParseFormOrMultiForm(maxMemory)
}

// Bind data from request.Form[key] to dest
// like /?id=123&isok=true&ft=1.2&ol[0]=1&ol[1]=2&ul[]=str&ul[]=array&user.Name=astaxie
// var id int  beegoInput.Bind(&id, "id")  id ==123
// var isok bool  beegoInput.Bind(&isok, "isok")  isok ==true
// var ft float64  beegoInput.Bind(&ft, "ft")  ft ==1.2
// ol := make([]int, 0, 2)  beegoInput.Bind(&ol, "ol")  ol ==[1 2]
// ul := make([]string, 0, 2)  beegoInput.Bind(&ul, "ul")  ul ==[str array]
// user struct{Name}  beegoInput.Bind(&user, "user")  user == {Name:"astaxie"}
func (input *BeegoInput) Bind(dest interface{}, key string) error {
	return (*context.BeegoInput)(input).Bind(dest, key)
}
