package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"testing"
	"time"
)

var (
	rds = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})
	ctx = context.Background()
	cr  = &CacheRedis{rds: rds}
)

func TestCacheRedis_Decr(t *testing.T) {
	cr.Set(context.Background(), "decr_name", "0")
	cr.Delete(context.Background(), "decr_not_found")
	cr.Set(ctx, "str", "str1")

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
			args:    args{ctx: ctx, key: "decr_name"},
			want:    -1,
			wantErr: false,
		},
		{
			name:    "decr",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "decr_name"},
			want:    -2,
			wantErr: false,
		},
		{
			name:    "empty",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "decr_not_found"},
			want:    -1,
			wantErr: false,
		},
		{
			name:    "string",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "str"},
			want:    0,
			wantErr: true,
		},
	}

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
	cr.Set(ctx, "exits_name", "zhangsan")

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
		{
			name:    "exits",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "exits_name"},
			wantErr: false,
		},
		{
			name:    "not_exits",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "not_exits_name1"},
			wantErr: false,
		},
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
	cr.Set(ctx, "exist_str", "hello")
	cr.Delete(ctx, "not_exist_str")

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
		{
			name:    "exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "exist_str"},
			want:    true,
			wantErr: false,
		},
		{
			name:    "exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "not_exist_str"},
			want:    false,
			wantErr: false,
		},
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
	cr.Set(ctx, "exist_expire", "world")
	cr.Delete(ctx, "not_exist_expire")

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
		{
			name:    "exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "exist_expire", expiration: time.Second * 10},
			want:    true,
			wantErr: false,
		},
		{
			name:    "not_exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "not_exist_expire", expiration: time.Second * 10},
			want:    false,
			wantErr: false,
		},
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
	cr.Set(ctx, "exist_get", "world")
	cr.Delete(ctx, "not_exist_get")

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
		{
			name:    "exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "exist_get"},
			want:    "world",
			wantErr: false,
		},
		{
			name:    "not_exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "not_exist_get"},
			want:    "",
			wantErr: true,
		},
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
	cr.Set(context.Background(), "incr_name", "0")
	cr.Delete(context.Background(), "incr_not_found")
	cr.Set(ctx, "str", "str1")

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
			name:    "incr",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "incr_name"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "incr",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "incr_name"},
			want:    2,
			wantErr: false,
		},
		{
			name:    "empty",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "incr_not_found"},
			want:    1,
			wantErr: false,
		},
		{
			name:    "string",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "str"},
			want:    0,
			wantErr: true,
		},
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
	cr.Set(ctx, "exist_set", "world")
	cr.Delete(ctx, "not_exist_set")

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
		{
			name:    "exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "exist_set", value: "world"},
			wantErr: false,
		},
		{
			name:    "not_exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "not_exist_set", value: "world"},
			wantErr: false,
		},
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
	cr.Set(ctx, "exist_setex", "world")
	cr.Delete(ctx, "not_exist_setex")

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
		{
			name:    "exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "exist_setex", value: "world", expiration: time.Second * 20},
			wantErr: false,
		},
		{
			name:    "not_exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "not_exist_setex", value: "world", expiration: time.Second * 20},
			wantErr: false,
		},
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
	cr.Set(ctx, "exist_durable_ttl", "world")
	cr.SetEx(ctx, "exist_ttl", "world", time.Second*20)
	cr.Delete(ctx, "not_exist_ttl")

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
		{
			name:    "exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "exist_ttl"},
			want:    20,
			wantErr: false,
		},
		{
			name:    "not_exist",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "not_exist_ttl"},
			want:    -2,
			wantErr: false,
		},
		{
			name:    "exist_durable",
			fields:  fields{redis: rds},
			args:    args{ctx: ctx, key: "exist_durable_ttl"},
			want:    -1,
			wantErr: false,
		},
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
