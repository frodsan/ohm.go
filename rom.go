package rom

import "github.com/garyburd/redigo/redis"

type Connection struct {
	Pool *redis.Pool
}

func Connect(url string, options ...func(*redis.Pool)) Connection {
	pool := redis.Pool{
		Dial: func() (redis.Conn, error) {
			return redis.DialURL(url)
		},
	}

	for _, option := range options {
		option(&pool)
	}

	conn := Connection{
		Pool: &pool,
	}

	return conn
}

func (c Connection) Flush() error {
	conn := c.Pool.Get()

	defer conn.Close()

	_, err := conn.Do("FLUSHDB")

	return err
}
