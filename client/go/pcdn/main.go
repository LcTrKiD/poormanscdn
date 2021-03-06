/*
 * Copyright (c) 2017 Salle, Alexandre <atsalle@inf.ufrgs.br>
 * Author: Salle, Alexandre <atsalle@inf.ufrgs.br>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of
 * this software and associated documentation files (the "Software"), to deal in
 * the Software without restriction, including without limitation the rights to
 * use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of
 * the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS
 * FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER
 * IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN
 * CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/alexandres/poormanscdn/client"
)

var path, cdnUrl, method, refererHost, userHost, secret string
var modified, expires int64

func init() {
	flag.StringVar(&cdnUrl, "cdnurl", "", "cdnurl")
	flag.StringVar(&method, "method", "GET", "method")
	flag.StringVar(&refererHost, "refererhost", "", "restrict referer host")
	flag.StringVar(&userHost, "userhost", "", "restrict user host")
	flag.Int64Var(&modified, "modified", 0, "modified")
	flag.Int64Var(&expires, "expires", 0, "expires")
	flag.StringVar(&path, "path", "", "path")
	flag.StringVar(&secret, "secret", "", "secret")
}

func main() {
	flag.Parse()
	if cdnUrl == "" || secret == "" {
		log.Fatal("cdnurl and secret are mandatory")
	}
	lastModifiedAt := time.Unix(modified, 0)
	expiresAt := time.Unix(expires, 0)
	p := client.SigParams{
		Method:      method,
		Path:        client.TrimPath(path),
		UserHost:    userHost,
		RefererHost: refererHost,
		Modified:    lastModifiedAt,
		Expires:     expiresAt,
	}
	url, err := client.GetSignedUrl(secret, cdnUrl, p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(url)
}
