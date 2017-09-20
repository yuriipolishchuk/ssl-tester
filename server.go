package main

import (
    "crypto/tls"
    "log"
    "net/http"
    "os"
)

func getparam(n int, fallback string) string {
    if len(os.Args) > 1 {
        if len(os.Args[n]) != 0 {
            return os.Args[n]
        }
    }
    return fallback
}

func main() {
    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
        w.Header().Add("Strict-Transport-Security", "max-age=63072000;")
        w.Write([]byte("<h1>Hello World!</h1>\n<h1>👋</h1>"))
    })
    cfg := &tls.Config{
        MinVersion:               tls.VersionTLS12,
        CurvePreferences:         []tls.CurveID{tls.CurveP521, tls.CurveP384, tls.CurveP256},
        PreferServerCipherSuites: true,
        CipherSuites: []uint16{
            tls.TLS_ECDHE_ECDSA_WITH_AES_128_CBC_SHA,
            tls.TLS_ECDHE_ECDSA_WITH_AES_256_CBC_SHA,
            tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
            tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
            // POLY1305 ciphers are not in Go 1.6 and 1.7
            //          tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
            //          tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
            tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
            tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
            tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
            tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
        },
    }
    srv := &http.Server{
        Addr:         "127.0.0.1:443",
        Handler:      mux,
        TLSConfig:    cfg,
        TLSNextProto: make(map[string]func(*http.Server, *tls.Conn, http.Handler), 0),
    }
    log.Fatal(srv.ListenAndServeTLS(getparam(1, "/etc/ssl-tester/tls.crt"), getparam(2, "/etc/ssl-tester/tls.key")))
}
