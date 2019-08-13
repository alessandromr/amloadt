package ualyzer

import (
	"errors"
	"net/url"
	"regexp"
)

//URL struct is used to manage and extract information from an url
type URL struct {
	Raw    string //Raw Url (lightsail.aws.amazon.com/new/server)
	FQDN   string //FD Fully Qualified Domain Name (lightsail.aws.amazon.com)
	TLD    string //TLD Top Level Domain (com)
	HN     string //HostName (lightsail) ???
	Domain string //Domain (amazon.com) ???
}

//Parse function
func Parse(raw string) (*URL, error) {
	u := URL{Raw: raw}
	if !u.validate() {
		return nil, errors.New("Error: URL not valid")
	}
	err := u.doParse()
	if err != nil {
		return nil, errors.New("Error parsing URL")
	}
	return &u, nil
}

func (u *URL) doParse() error {
	u.getDomain()
	u.getFQDN()
	return nil
}

func (u *URL) validate() bool {
	//Check Using Standard Lib
	_, err := url.ParseRequestURI(u.Raw)
	if err != nil {
		return false
	}
	//validare url with custmo regex
	main := regexp.MustCompile("^(?:http(s)?:\\/\\/)?(?:[a-zA-Z0-9](?:-)?[a-zA-Z0-9_.]*(?:-)?[a-zA-Z0-9][\\.\\-])+(?:[a-zA-Z]{2,10})(?:\\/(?:[\\w\\/\\[\\]\\(\\)\\-._~:?#@!$&'`*+§|µ@€~^,%<>;=]+)?)?$")
	if !main.MatchString(u.Raw) {
		return false
	}
	return true
}

func (u *URL) getFQDN() error {
	re := regexp.MustCompile("http[s]{0,1}:\\/\\/([a-zA-Z0-9\\.\\-]+)\\/(.*)")
	match := re.FindStringSubmatch(u.Raw)
	if len(match) > 1 {
		u.FQDN = match[1]
		return nil
	}
	return errors.New("URL not valid")
}

func (u *URL) getDomain() error {
	re := regexp.MustCompile("([a-zA-Z0-9\\-\\.]*\\.)*([a-zA-Z0-9\\-]+\\.{1}[a-z]{2,6}){1}(.*)$")
	match := re.FindStringSubmatch(u.Raw)
	if len(match) > 2 {
		u.FQDN = match[2]
		return nil
	}
	return errors.New("URL not valid")
}
