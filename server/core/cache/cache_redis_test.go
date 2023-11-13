package cache

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestCacheRedis_Decr(t *testing.T) {

	rds := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	ctx := context.Background()

	type fields struct {
		redis *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		{
			name:    "decr",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "name1"},
			want:    -1,
			wantErr: true,
		},
		{
			name:    "decr",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "name1"},
			want:    -2,
			wantErr: true,
		},
		{
			name:    "decr",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "name1"},
			want:    -3,
			wantErr: true,
		},
	}

	init := &CacheRedis{
		rds: rds,
	}
	init.Delete(context.Background(), "name1")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheRedis{
				rds: tt.fields.redis,
			}
			got, err := c.Decr(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Decr() got = %v, want %v", got, tt.want)
			}

		})
	}
}

func TestCacheRedis_Delete(t *testing.T) {
	type fields struct {
		redis *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheRedis{
				rds: tt.fields.redis,
			}
			if err := c.Delete(tt.args.ctx, tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCacheRedis_Exist(t *testing.T) {
	type fields struct {
		redis *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheRedis{
				rds: tt.fields.redis,
			}
			got, err := c.Exist(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exist() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exist() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCacheRedis_Expire(t *testing.T) {
	type fields struct {
		redis *redis.Client
	}
	type args struct {
		ctx        context.Context
		key        string
		expiration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    bool
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheRedis{
				rds: tt.fields.redis,
			}
			got, err := c.Expire(tt.args.ctx, tt.args.key, tt.args.expiration)
			if (err != nil) != tt.wantErr {
				t.Errorf("Expire() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Expire() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCacheRedis_Get(t *testing.T) {
	type fields struct {
		redis *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheRedis{
				rds: tt.fields.redis,
			}
			got, err := c.Get(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCacheRedis_Incr(t *testing.T) {
	type fields struct {
		redis *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheRedis{
				rds: tt.fields.redis,
			}
			got, err := c.Incr(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Incr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Incr() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCacheRedis_Set(t *testing.T) {
	type fields struct {
		redis *redis.Client
	}
	type args struct {
		ctx   context.Context
		key   string
		value string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheRedis{
				rds: tt.fields.redis,
			}
			if err := c.Set(tt.args.ctx, tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCacheRedis_SetEx(t *testing.T) {
	type fields struct {
		redis *redis.Client
	}
	type args struct {
		ctx        context.Context
		key        string
		value      string
		expiration time.Duration
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheRedis{
				rds: tt.fields.redis,
			}
			if err := c.SetEx(tt.args.ctx, tt.args.key, tt.args.value, tt.args.expiration); (err != nil) != tt.wantErr {
				t.Errorf("SetEx() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCacheRedis_Ttl(t *testing.T) {
	type fields struct {
		redis *redis.Client
	}
	type args struct {
		ctx context.Context
		key string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &CacheRedis{
				rds: tt.fields.redis,
			}
			got, err := c.Ttl(tt.args.ctx, tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("Ttl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Ttl() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewCacheRedis(t *testing.T) {
	type args struct {
		client *redis.Client
	}
	tests := []struct {
		name string
		args args
		want *CacheRedis
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewCacheRedis(tt.args.client); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewCacheRedis() = %v, want %v", got, tt.want)
			}
		})
	}
}
