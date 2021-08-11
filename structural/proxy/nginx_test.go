package proxy

import "testing"

func Test_nginx_handleRequest(t *testing.T) {
	n := newNginxServer()
	type args struct {
		url    string
		method string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 string
	}{
		{name: "app第一次", args: args{url: "/app/status", method: "GET"}, want: 200, want1: "Ok"},
		{name: "app第二次", args: args{url: "/app/status", method: "GET"}, want: 200, want1: "Ok"},
		{name: "app限流", args: args{url: "/app/status", method: "GET"}, want: 403, want1: "Not Allowed"},
		{name: "user正常", args: args{url: "/create/user", method: "POST"}, want: 201, want1: "User Create"},
		{name: "user方法异常", args: args{url: "/create/user", method: "GET"}, want: 400, want1: "Not Ok"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := n.handleRequest(tt.args.url, tt.args.method)
			if got != tt.want {
				t.Errorf("handleRequest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("handleRequest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
