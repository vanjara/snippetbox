package main

import (
	"net/http"
	"testing"
)

func TestPing(t *testing.T) {
	// rr := httptest.NewRecorder()

	// r, err := http.NewRequest(http.MethodGet, "/", nil)

	// if err != nil {
	// 	t.Fatal(err)
	// }

	// ping(rr, r)

	// rs := rr.Result()

	app := newTestApplication(t)
	ts := newTestServer(t, app.routes())

	defer ts.Close()

	code, _, body := ts.get(t, "/ping")

	if code != http.StatusOK {
		t.Errorf("want %d; got %d", http.StatusOK, code)
	}

	if string(body) != "OK" {
		t.Errorf("want body to equal %q", "OK")
	}
	// app := &application{
	// 	errorLog: log.New(io.Discard, "", 0),
	// 	infoLog:  log.New(io.Discard, "", 0),
	// }

	// ts := httptest.NewTLSServer(app.routes())
	// defer ts.Close()

	// rs, err := ts.Client().Get(ts.URL + "/ping")
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// if rs.StatusCode != http.StatusOK {
	// 	t.Errorf("want %d; got %d", http.StatusOK, rs.StatusCode)
	// }

	// defer rs.Body.Close()

	// body, err := io.ReadAll(rs.Body)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if string(body) != "OK" {
	// 	t.Errorf("want body equal %q", "OK")
	// }
}
