// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"flag"
	"fmt"
	"log"
	//"runtime"

	"github.com/nlamirault/iris/logging"
)

var (
	debug       bool
	version     bool
	cadvisorURL string
	numstats    int
)

func init() {
	// runtime.GOMAXPROCS(4)
	flag.BoolVar(&version, "version", false, "print version and exit")
	flag.BoolVar(&debug, "d", false, "run in debug mode")
	flag.IntVar(&numstats, "numstats", 2, "number of stats to send")
	flag.StringVar(&cadvisorURL, "cadvisor-url", "http://localhost:8080", "URL of cAdvisor")
	flag.Parse()
}

func main() {
	if debug {
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
	if version {
		fmt.Printf("Iris v%s\n", Version)
		return
	}
	log.Printf("[INFO] [iris] CAdvisor monitoring")
	cadvisor, err := NewCAdvisor(cadvisorURL, numstats)
	if err != nil {
		log.Printf("[INFO] [iris] Failed to creates cAdvisor client : %v", err)
		return
	}
	monitor(cadvisor)
	log.Printf("[INFO] [iris] Stopping")
}
