package main

func Authorization(token string) Decorator {
	return Header("Authorization", token)
}

func Header(name, value string) Decorator {
	return func(c Client) Client {
		return ClientFunc(func(r string) (string, error) {
			r = name + ":" + value + "\n" + r
			return c.Do(r)
		})
	}
}

func Decorate(c Client, ds ...Decorator) Client {
	decorated := c
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}
