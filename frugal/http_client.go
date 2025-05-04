/*CODE ORIGINAL	BY : APACHE*/
/*REMAKE 				  BY : DONI HARIANTO*/
/*LICENSE				   BY : APACHE 	Ft	 PT.Amin-Bot-Super*/


/*
DONI OLENG
*/

package superhyper

import (
	"bytes"
	"context"
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"github.com/valyala/fasthttp"
)

/*
DONI OLENG
*/

var DefaultHttpClient *http.Client = http.DefaultClient

type THttpClient struct {
	client             *http.Client
	response           *http.Response
	url                *url.URL
	urls               string
	requestBuffer      *bytes.Buffer
	header             http.Header
	nsecConnectTimeout int64
	nsecReadTimeout    int64
	results            string
	body               []byte
	moreCompact        bool
}

/*
DONI OLENG
*/

type THttpClientTransportFactory struct {
	options THttpClientOptions
	url     string
}

/*
DONI OLENG
*/

type THttpClientOptions struct {
	Client *http.Client
}

/*
DONI OLENG
*/

func FastModHttpClient(urlstr string, tr *http.Transport, headers http.Header) *THttpClient {
	parsedURL, _ := url.Parse(urlstr)
	return &THttpClient{client: &http.Client{Transport: tr}, url: parsedURL, urls: urlstr, requestBuffer: bytes.NewBuffer(make([]byte, 0, 1024)), header: headers}
}

func SuperModHttpClient(urlstr string, tr *http.Transport, headers http.Header) *THttpClient {
	parsedURL, _ := url.Parse(urlstr)
	return &THttpClient{client: &http.Client{Transport: tr}, url: parsedURL, urls: urlstr, requestBuffer: bytes.NewBuffer(make([]byte, 0, 1024)), header: headers}
}

func TurboModHttpClient(urlstr string, tr *http.Transport, headers http.Header) *THttpClient {
	parsedURL, _ := url.Parse(urlstr)
	return &THttpClient{client: &http.Client{Transport: tr}, url: parsedURL, urls: urlstr, requestBuffer: bytes.NewBuffer(make([]byte, 0, 1024)), header: headers}
}

/*
DONI OLENG
*/

func TalkProxyLOP(urlstr string, cl *http.Client) (TTransport, error) {
	return LOPProxyOptions(urlstr, THttpClientOptions{Client:cl})
}
func LOPProxyOptions(urlstr string, options THttpClientOptions) (TTransport, error) {
	parsedURL, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 0, 1024)
	client := options.Client
	if client == nil {
		client = DefaultHttpClient
	}
	httpHeader := map[string][]string{"Content-Type": {"application/x-thrift"}}
	return &THttpClient{client: client, url: parsedURL, urls: parsedURL.String(), requestBuffer: bytes.NewBuffer(buf), header: httpHeader}, nil
}

func TalkProxyMsg(urlstr string, cl *http.Client) (TTransport, error) {
	return MsgTalkProxyOptions(urlstr, THttpClientOptions{Client:cl})
}
func MsgTalkProxyOptions(urlstr string, options THttpClientOptions) (TTransport, error) {
	parsedURL, err := url.Parse(urlstr)
	if err != nil {
		return nil, err
	}
	buf := make([]byte, 0, 1024)
	client := options.Client
	if client == nil {
		client = DefaultHttpClient
	}
	httpHeader := map[string][]string{"Content-Type": {"application/x-thrift"}}
	return &THttpClient{client: client, url: parsedURL, urls: parsedURL.String(), requestBuffer: bytes.NewBuffer(buf), header: httpHeader}, nil
}

/*
DONI OLENG
*/

func NewFastHttpModHttpClient(urlstr string, tr *fasthttp.Client, headers http.Header) *THttpClient {
	parsedURL, _ := url.Parse(urlstr)
	return &THttpClient{client: &http.Client{}, url: parsedURL, urls: urlstr, requestBuffer: bytes.NewBuffer(make([]byte, 0, 1024)), header: headers}
}

/*
DONI OLENG
*/

func NewTHttpClientTransportFactory(url string) *THttpClientTransportFactory {
	return NewTHttpClientTransportFactoryWithOptions(url, THttpClientOptions{})
}

/*
DONI OLENG
*/

func NewTHttpClientTransportFactoryWithOptions(url string, options THttpClientOptions) *THttpClientTransportFactory {
	return &THttpClientTransportFactory{url: url, options: options}
}

/*
DONI OLENG
*/

/*
DONI OLENG
*/

func NewTHttpClientHeader(urlstr string, cl *http.Client, hed http.Header) TTransport {
	buf := make([]byte, 0, 512)
	return &THttpClient{client: cl, urls: urlstr, requestBuffer: bytes.NewBuffer(buf), header: hed}
}

/*
DONI OLENG
*/

//func NewTHttpClient(urlstr string, tr *http.Transport) (TTransport, error) {
	//return NewTHttpClientWithOptions(urlstr, tr)
//}

/*
DONI OLENG
*/

func (p *THttpClient) SetMoreCompact(value bool) {
	p.moreCompact = value
}
func (p *THttpClient) GetBody() []byte {
	return p.body
}
func (p *THttpClient) GetTPCopy() *THttpClient {
	var a = p
	return a
}
func (p *THttpClient) SetHeader(key string, value string) {
	p.header.Add(key, value)
}
func (p *THttpClient) GetHeader(key string) string {
	return p.header.Get(key)
}
func (p *THttpClient) DelHeader(key string) {
	p.header.Del(key)
}
func (p *THttpClient) Open() error {
	return nil
}
func (p *THttpClient) IsOpen() bool {
	return p.response != nil || p.requestBuffer != nil
}
func (p *THttpClient) closeResponse() error {
	var err error
	if p.response != nil && p.response.Body != nil {
		io.Copy(ioutil.Discard, p.response.Body)
		err = p.response.Body.Close()
	}
	p.response = nil
	return err
}
func (p *THttpClient) Close() error {
	if p.requestBuffer != nil {
		p.requestBuffer.Reset()
		p.requestBuffer = nil
	}
	return p.closeResponse()
}
func (p *THttpClient) Read(buf []byte) (int, error) {
	if p.response == nil {
		return 0, NewTTransportException(NOT_OPEN, "Response buffer is empty, no request.")
	}
	n, err := p.response.Body.Read(buf)
	if n > 0 && (err == nil || errors.Is(err, io.EOF)) {
		return n, nil
	}
	return n, NewTTransportExceptionFromError(err)
}
func (p *THttpClient) ReadByte() (c byte, err error) {
	if p.response == nil {
		return 0, NewTTransportException(NOT_OPEN, "Response buffer is empty, no request.")
	}
	return readByte(p.response.Body)
}
func (p *THttpClient) Write(buf []byte) (int, error) {
	if p.requestBuffer == nil {
		return 0, NewTTransportException(NOT_OPEN, "Request buffer is nil, connection may have been closed.")
	}
	return p.requestBuffer.Write(buf)
}
func (p *THttpClient) WriteByte(c byte) error {
	if p.requestBuffer == nil {
		return NewTTransportException(NOT_OPEN, "Request buffer is nil, connection may have been closed.")
	}
	return p.requestBuffer.WriteByte(c)
}
func (p *THttpClient) WriteString(s string) (n int, err error) {
	if p.requestBuffer == nil {
		return 0, NewTTransportException(NOT_OPEN, "Request buffer is nil, connection may have been closed.")
	}
	return p.requestBuffer.WriteString(s)
}
func (p *THttpClient) FlushMod(ctx context.Context) ([]byte, error) {
	req, _ := http.NewRequest("POST", p.urls, p.requestBuffer)
	req.Header = p.header
	req = req.WithContext(ctx)
	response, err := p.client.Do(req)
	if err != nil {
		return []byte{}, err
	}
	res, _ := ioutil.ReadAll(response.Body)
	return res, err
}
func (p *THttpClient) Flush(ctx context.Context) error {
	req, err := http.NewRequest("POST", p.urls, p.requestBuffer)
	if err != nil {
		return NewTTransportExceptionFromError(err)
	}
	req.Header = p.header
	response, err := p.client.Do(req)
	if err != nil {
		return NewTTransportExceptionFromError(err)
	}
	if response.StatusCode != http.StatusOK {
		p.response = response
		p.closeResponse()
		return NewTTransportException(UNKNOWN_TRANSPORT_EXCEPTION, "HTTP Response code: "+strconv.Itoa(response.StatusCode))
	}
	p.response = response
	if p.moreCompact {
		p.body, _ = ioutil.ReadAll(response.Body)
	}
	return nil
}
func (p *THttpClient) RemainingBytes() (num_bytes uint64) {
	len := p.response.ContentLength
	if len >= 0 {
		return uint64(len)
	}
	const maxSize = ^uint64(0)
	return maxSize
}

/*
DONI OLENG
*/

func NewTHttpPostClientTransportFactory(url string) *THttpClientTransportFactory {
	return NewTHttpClientTransportFactoryWithOptions(url, THttpClientOptions{})
}

/*
DONI OLENG
*/

func NewTHttpPostClientTransportFactoryWithOptions(url string, options THttpClientOptions) *THttpClientTransportFactory {
	return NewTHttpClientTransportFactoryWithOptions(url, options)
}

/*
DONI OLENG
*/