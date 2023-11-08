// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package config

import (
	"time"
)

// Web containts the configuration information for building for web and creating
// the HTML page that loads a Go wasm app and its resources.
type Web struct { //gti:add

	// Port is the port to serve the page at when using the serve command.
	Port string `def:"8080"`

	// RandomVersion is whether to automatically add a random string to the
	// end of the version string for the app when building for web. This is
	// necessary in order for changes made during local development to show up,
	// but should not be enabled in release builds to prevent constant inaccurate
	// update messages. It is enabled by default in the serve command and disabled
	// by default otherwise.
	RandomVersion bool

	// Gzip is whether to gzip the app.wasm file that is built in the build command
	// and serve it as a gzip-encoded file in the run command.
	Gzip bool

	// A placeholder background color for the application page to display before
	// its stylesheets are loaded.
	//
	// DEFAULT: #2d2c2c.
	BackgroundColor string `def:"#2d2c2c"`

	// The theme color for the application. This affects how the OS displays the
	// app (e.g., PWA title bar or Android's task switcher).
	//
	// DEFAULT: #2d2c2c.
	ThemeColor string `def:"#2d2c2c"`

	// The text displayed while loading a page. Load progress can be inserted by
	// including "{progress}" in the loading label.
	//
	// DEFAULT: "{progress}%".
	LoadingLabel string

	// The page language.
	//
	// DEFAULT: en.
	Lang string `def:"en"`

	// The page title.
	Title string

	// The page description.
	Description string

	// The page authors.
	Author string

	// The page keywords.
	Keywords []string

	// The path of the default image that is used by social networks when
	// linking the app.
	Image string

	// The interval between each app auto-update while running in a web browser.
	// Zero or negative values deactivates the auto-update mechanism.
	//
	// Default is 10 seconds.
	AutoUpdateInterval time.Duration `def:"10000000000"`

	// The environment variables that are passed to the progressive web app.
	//
	// Reserved keys:
	// - GOAPP_VERSION
	// - GOAPP_GOAPP_STATIC_RESOURCES_URL
	Env map[string]string

	// The HTTP header to retrieve the WebAssembly file content length.
	//
	// Content length finding falls back to the Content-Length HTTP header when
	// no content length is found with the defined header.
	WasmContentLengthHeader string

	// The template used to generate app-worker.js. The template follows the
	// text/template package model.
	//
	// By default set to DefaultAppWorkerJS, changing the template have very
	// high chances to mess up go-app usage. Any issue related to a custom app
	// worker template is not supported and will be closed.
	ServiceWorkerTemplate string
}
