package app

import "net/url"

type input struct {
	Link string
}

type output struct {
	Key string
}

type outErr struct {
	Error string
}

func getErr(e error) outErr {
	return outErr{e.Error()}
}

func getUrl(base string) (string, error) {
	u, _ := url.Parse(base)
	if !u.IsAbs() {
		u.Scheme = "http"
		base = u.String()
	}
	if _, err := url.Parse(base); err != nil {
		return "", err
	}
	return base, nil
}
