package main

import "testing"

func TestRecoverPanic(t *testing.T) {
	app := newTestApplication(t)

	ts := newTestServer(t, app.routesTest())
	defer ts.Close()

}
