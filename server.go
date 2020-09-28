/*
 * Copyright 2020 Hayo van Loon
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */
package main

import (
	"context"
	"flag"
	"github.com/HayoVanLoon/go-commons/logjson"
	"golang.org/x/oauth2"
	"google.golang.org/api/idtoken"
	"google.golang.org/api/option"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

func chainDirector(proxy *httputil.ReverseProxy, host string) {
	std := proxy.Director
	proxy.Director = func(r *http.Request) {
		std(r)
		r.Host = host
	}
}

func createAuthProxy(ctx context.Context, t *url.URL, aud, creds string) (http.Handler, error) {
	var opts []idtoken.ClientOption
	if creds != "" {
		opts = append(opts, option.WithCredentialsFile(creds))
	}

	ts, err := idtoken.NewTokenSource(ctx, aud, opts...)
	if err != nil {
		return nil, err
	}

	proxy := httputil.NewSingleHostReverseProxy(t)
	chainDirector(proxy, t.Host)
	proxy.Transport = &oauth2.Transport{Source: ts}

	return proxy, nil
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	target := os.Getenv("TARGET")
	if target == "" {
		logjson.Critical("missing TARGET")
	}
	t, _ := url.Parse(target)

	audience := os.Getenv("AUDIENCE")
	if audience == "" {
		audience = target
	}

	creds := flag.String("creds", "", "")
	flag.Parse()

	ctx := context.Background()
	proxy, err := createAuthProxy(ctx, t, audience, *creds)
	if err != nil {
		logjson.Critical("could not create proxy: %s", err)
	}

	logjson.Info("Settings: TARGET: %s", target)
	logjson.Info("Settings: AUDIENCE: %s", audience)
	if *creds != "" {
		logjson.Info("Settings: credentials file: %s", *creds)
	}
	err = http.ListenAndServe(":"+port, proxy)
	if err != nil {
		logjson.Critical(err)
	}
}
