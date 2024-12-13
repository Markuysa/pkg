package errs

import (
	"reflect"
	"testing"
)

func TestError_Code(t *testing.T) {
	type fields struct {
		msg  string
		code ErrorCode
	}
	tests := []struct {
		name   string
		fields fields
		want   ErrorCode
	}{
		{
			name: "Test Error",
			fields: fields{
				msg:  "Test Error",
				code: 2,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Msg:  tt.fields.msg,
				Code: tt.fields.code,
			}
			if got := e.GetCode(); got != tt.want {
				t.Errorf("Code() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Error(t *testing.T) {
	type fields struct {
		msg  string
		code ErrorCode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Error",
			fields: fields{
				msg:  "Test Error",
				code: 2,
			},
			want: "{\"msg\":\"Test Error\",\"code\":2}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Msg:  tt.fields.msg,
				Code: tt.fields.code,
			}
			if got := e.Error(); got != tt.want {
				t.Errorf("Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestError_Message(t *testing.T) {
	type fields struct {
		msg  string
		code ErrorCode
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Test Error",
			fields: fields{
				msg:  "Test Error",
				code: 2,
			},
			want: "Test Error",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Error{
				Msg:  tt.fields.msg,
				Code: tt.fields.code,
			}
			if got := e.Message(); got != tt.want {
				t.Errorf("Message() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		msg  string
		code ErrorCode
	}
	tests := []struct {
		name string
		args args
		want *Error
	}{
		{
			name: "Test Error",
			args: args{
				msg:  "Test Error",
				code: 2,
			},
			want: &Error{
				Msg:  "Test Error",
				Code: 2,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.msg, tt.args.code); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
