package permission

import "testing"

func TestMapHTTPMethodToOperation(t *testing.T) {
	tests := []struct {
		method string
		want   Operation
	}{
		{method: "GET", want: OperationRead},
		{method: "HEAD", want: OperationRead},
		{method: "OPTIONS", want: OperationRead},
		{method: "PROPFIND", want: OperationRead},
		{method: "PUT", want: OperationWrite},
		{method: "PATCH", want: OperationWrite},
		{method: "PROPPATCH", want: OperationWrite},
		{method: "POST", want: OperationCreate},
		{method: "MKCOL", want: OperationCreate},
		{method: "COPY", want: OperationWrite},
		{method: "MOVE", want: OperationWrite},
		{method: "DELETE", want: OperationDelete},
		{method: "UNKNOWN", want: OperationRead},
	}

	for _, tt := range tests {
		if got := MapHTTPMethodToOperation(tt.method); got != tt.want {
			t.Fatalf("MapHTTPMethodToOperation(%q) = %s, want %s", tt.method, got, tt.want)
		}
	}
}
