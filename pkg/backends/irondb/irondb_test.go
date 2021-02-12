/*
 * Copyright 2018 Comcast Cable Communications Management, LLC
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
 */

package irondb

import (
	"testing"

	"github.com/tricksterproxy/trickster/pkg/backends"
	"github.com/tricksterproxy/trickster/pkg/backends/irondb/model"
	bo "github.com/tricksterproxy/trickster/pkg/backends/options"
	cr "github.com/tricksterproxy/trickster/pkg/cache/registration"
	"github.com/tricksterproxy/trickster/pkg/config"
	tl "github.com/tricksterproxy/trickster/pkg/logging"
)

func TestIRONdbClientInterfacing(t *testing.T) {

	// this test ensures the client will properly conform to the
	// Client and TimeseriesBackend interfaces

	c, err := NewClient("test", nil, nil, nil, nil)
	if err != nil {
		t.Error(err)
	}
	var bo backends.Backend = c
	var to backends.TimeseriesBackend = c

	if bo.Name() != "test" {
		t.Errorf("expected %s got %s", "test", bo.Name())
	}

	if to.Name() != "test" {
		t.Errorf("expected %s got %s", "test", to.Name())
	}
}

var testModeler = model.NewModeler()

func TestNewClient(t *testing.T) {
	conf, _, err := config.Load("trickster", "test",
		[]string{"-origin-url", "http://example.com", "-provider", "TEST_CLIENT"})
	if err != nil {
		t.Fatalf("Could not load configuration: %s", err.Error())
	}

	caches := cr.LoadCachesFromConfig(conf, tl.ConsoleLogger("error"))
	defer cr.CloseCaches(caches)
	cache, ok := caches["default"]
	if !ok {
		t.Errorf("Could not find default configuration")
	}

	oc := &bo.Options{Provider: "TEST_CLIENT"}
	c, err := NewClient("default", oc, nil, cache, testModeler)
	if err != nil {
		t.Error(err)
	}
	if c.Name() != "default" {
		t.Errorf("expected %s got %s", "default", c.Name())
	}

	if c.Cache().Configuration().Provider != "memory" {
		t.Errorf("expected %s got %s", "memory", c.Cache().Configuration().Provider)
	}

	if c.Configuration().Provider != "TEST_CLIENT" {
		t.Errorf("expected %s got %s", "TEST_CLIENT", c.Configuration().Provider)
	}
}