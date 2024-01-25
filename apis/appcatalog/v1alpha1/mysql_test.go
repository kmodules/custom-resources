/*
Copyright AppsCode Inc. and Contributors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

import "testing"

func TestCanonicalMySQLDSN(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "already dsn",
			args:    args{dsn: "tcp(mysql-demo.mysql.svc)/"},
			want:    "tcp(mysql-demo.mysql.svc)/",
			wantErr: false,
		},
		{
			name:    "already dsn",
			args:    args{dsn: "tcp(mysql-demo.mysql.svc)/dbname"},
			want:    "tcp(mysql-demo.mysql.svc)/dbname",
			wantErr: false,
		},
		{
			name:    "already dsn",
			args:    args{dsn: "tcp(mysql-demo.mysql.svc)/dbname?param=value"},
			want:    "tcp(mysql-demo.mysql.svc)/dbname?param=value",
			wantErr: false,
		},
		{
			name:    "already dsn",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value"},
			want:    "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value",
			wantErr: false,
		},
		{
			name:    "already dsn",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value&param2=value2"},
			want:    "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value&param2=value2",
			wantErr: false,
		},

		{
			name:    "no port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc"},
			want:    "tcp(mysql-demo.mysql.svc)/",
			wantErr: false,
		},
		{
			name:    "no port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc/dbname"},
			want:    "tcp(mysql-demo.mysql.svc)/dbname",
			wantErr: false,
		},
		{
			name:    "no port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc/dbname?param=value"},
			want:    "tcp(mysql-demo.mysql.svc)/dbname?param=value",
			wantErr: false,
		},
		{
			name:    "no port",
			args:    args{dsn: "tcp://username:password@mysql-demo.mysql.svc/dbname?param=value"},
			want:    "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value",
			wantErr: false,
		},
		{
			name:    "no port",
			args:    args{dsn: "tcp://username:password@mysql-demo.mysql.svc/dbname?param=value&param2=value2"},
			want:    "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value&param2=value2",
			wantErr: false,
		},

		{
			name:    "default port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc:3306"},
			want:    "tcp(mysql-demo.mysql.svc:3306)/",
			wantErr: false,
		},
		{
			name:    "default port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc:3306/dbname"},
			want:    "tcp(mysql-demo.mysql.svc:3306)/dbname",
			wantErr: false,
		},
		{
			name:    "default port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc:3306/dbname?param=value"},
			want:    "tcp(mysql-demo.mysql.svc:3306)/dbname?param=value",
			wantErr: false,
		},
		{
			name:    "default port",
			args:    args{dsn: "tcp://username:password@mysql-demo.mysql.svc:3306/dbname?param=value"},
			want:    "username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value",
			wantErr: false,
		},
		{
			name:    "default port",
			args:    args{dsn: "tcp://username:password@mysql-demo.mysql.svc:3306/dbname?param=value&param2=value2"},
			want:    "username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value&param2=value2",
			wantErr: false,
		},

		{
			name:    "custom port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc:1234"},
			want:    "tcp(mysql-demo.mysql.svc:1234)/",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc:1234/dbname"},
			want:    "tcp(mysql-demo.mysql.svc:1234)/dbname",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "tcp://mysql-demo.mysql.svc:1234/dbname?param=value"},
			want:    "tcp(mysql-demo.mysql.svc:1234)/dbname?param=value",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "tcp://username:password@mysql-demo.mysql.svc:1234/dbname?param=value"},
			want:    "username:password@tcp(mysql-demo.mysql.svc:1234)/dbname?param=value",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "tcp://username:password@mysql-demo.mysql.svc:1234/dbname?param=value&param2=value2"},
			want:    "username:password@tcp(mysql-demo.mysql.svc:1234)/dbname?param=value&param2=value2",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CanonicalMySQLDSN(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("CanonicalMySQLDSN() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("CanonicalMySQLDSN() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMySQLHost(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "missing port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value"},
			want:    "mysql-demo.mysql.svc:3306",
			wantErr: false,
		},
		{
			name:    "default port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value"},
			want:    "mysql-demo.mysql.svc:3306",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:1234)/dbname?param=value"},
			want:    "mysql-demo.mysql.svc:1234",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMySQLHost(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMySQLHost() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseMySQLHost() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMySQLHostname(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "missing port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value"},
			want:    "mysql-demo.mysql.svc",
			wantErr: false,
		},
		{
			name:    "default port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value"},
			want:    "mysql-demo.mysql.svc",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:1234)/dbname?param=value"},
			want:    "mysql-demo.mysql.svc",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMySQLHostname(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMySQLHostname() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseMySQLHostname() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseMySQLPort(t *testing.T) {
	type args struct {
		dsn string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "missing port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc)/dbname?param=value"},
			want:    "3306",
			wantErr: false,
		},
		{
			name:    "default port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:3306)/dbname?param=value"},
			want:    "3306",
			wantErr: false,
		},
		{
			name:    "custom port",
			args:    args{dsn: "username:password@tcp(mysql-demo.mysql.svc:1234)/dbname?param=value"},
			want:    "1234",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseMySQLPort(tt.args.dsn)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseMySQLPort() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseMySQLPort() got = %v, want %v", got, tt.want)
			}
		})
	}
}
