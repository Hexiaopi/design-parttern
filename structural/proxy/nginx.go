package proxy

type server interface {
	handleRequest(string, string) (int, string)
}

type application struct {
}

func (a *application) handleRequest(url, method string) (int, string) {
	if url == "/app/status" && method == "GET" {
		return 200, "Ok"
	}
	if url == "/create/user" && method == "POST" {
		return 201, "User Create"
	}
	return 400, "Not Ok"
}

type nginx struct {
	application     *application
	maxAllowRequest int
	rateLimiter     map[string]int
}

func newNginxServer() *nginx {
	return &nginx{
		application:     &application{},
		maxAllowRequest: 2,
		rateLimiter:     make(map[string]int),
	}
}

func (n *nginx) checkRateLimiting(url string) bool {
	if n.rateLimiter[url] == 0 {
		n.rateLimiter[url] = 1
	}
	if n.rateLimiter[url] > n.maxAllowRequest {
		return false
	}
	n.rateLimiter[url] = n.rateLimiter[url] + 1
	return true
}

func (n *nginx) handleRequest(url, method string) (int, string) {
	allowed := n.checkRateLimiting(url)
	if !allowed {
		return 403, "Not Allowed"
	}
	return n.application.handleRequest(url, method)
}
