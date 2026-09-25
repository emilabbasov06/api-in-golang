package middlewares

import "net/http"

func SecurityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("X-DNS-Prefetch-Control", "off")                                           // disables DNS Prefetching
		w.Header().Set("X-Frame-Options", "DENY")                                                 // prevents the <iframe> action
		w.Header().Set("X-XSS-Protection", "1;mode-block")                                        // enables Cross-Site Scripting Filter
		w.Header().Set("X-Content-Type-Options", "nosniff")                                       // prevents browsers from Mime sniffing the response
		w.Header().Set("Strict-Transport-Security", "max-age=63072000;includeSubDomains;preload") // it enforces HTTPS for specified max-age in seconds
		w.Header().Set("Content-Security-Policy", "default-src 'self'")                           // it controls which resources can be loaded from the same origin
		w.Header().Set("Referrer-Policy", "no-referrer")                                          // how much referrer information should be included inside the request
		w.Header().Set("X-Powered-By", "rubyonrails")                                             // this provides wrong information about your api infrastructure (like I used golang for this api but I say it uses ruby on rails which provides wrong information for someone with wrong intentions)

		w.Header().Set("Server", "")
		w.Header().Set("X-Permitted-Cross-Domain-Policies", "none")
		w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, max-age=0")
		w.Header().Set("Cross-Origin-Resource-Polics", "same-origin")
		w.Header().Set("Cross-Origin-Opener-Policy", "same-origin")
		w.Header().Set("Cross-Origin-Embedder-Policy", "require-corp")
		w.Header().Set("Permission-Policy", "geolocation=(self), microphone=()")

		next.ServeHTTP(w, r)
	})
}
