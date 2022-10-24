package rest

import "testing"

func TestResult_Into(t *testing.T) {
	type fields struct {
		body       []byte
		err        error
		statusCode int
	}
	type args struct {
		val interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "is error",
			args: args{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := Result{
				body:       tt.fields.body,
				err:        tt.fields.err,
				statusCode: tt.fields.statusCode,
			}
			if err := r.Into(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Into() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
