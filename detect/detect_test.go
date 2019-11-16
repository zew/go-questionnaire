package detect

import (
	"net/http"
	"testing"
)

func TestDetect(t *testing.T) {

	type tst struct {
		Label, UA string
		Mobile    bool
	}

	// Taken from deviceatlas.com/blog/mobile-browser-user-agent-strings

	tsts := []tst{
		{
			"Safari for iOS",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_1 like Mac OS X) AppleWebKit/603.1.30 (KHTML, like Gecko) Version/10.0 Mobile/14E304 Safari/602.1",
			true,
		}, {
			"Chrome Mobile",
			"Mozilla/5.0 (Linux; Android 7.0; SM-G930V Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/59.0.3071.125 Mobile Safari/537.36",
			true,
		}, {}, {
			// https://developer.chrome.com/multidevice/user-agent
			"Chrome iOS Generic",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1",
			true,
		}, {
			// https://developer.chrome.com/multidevice/user-agent
			"Chrome Android Generic",
			"Mozilla/5.0 (Linux; <Android Version>; <Build Tag etc.>) AppleWebKit/<WebKit Rev> (KHTML, like Gecko) Chrome/<Chrome Rev> Mobile Safari/<WebKit Rev>",
			true,
		}, {
			"Firefox for Android",
			"Mozilla/5.0 (Android 7.0; Mobile; rv:54.0) Gecko/54.0 Firefox/54.0",
			true,
		}, {
			"Firefox for iOS",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3_2 like Mac OS X) AppleWebKit/603.2.4 (KHTML, like Gecko) FxiOS/7.5b3349 Mobile/14F89 Safari/603.2.4",
			true,
		}, {
			"Samsung Browser",
			"Mozilla/5.0 (Linux; Android 7.0; SAMSUNG SM-G955U Build/NRD90M) AppleWebKit/537.36 (KHTML, like Gecko) SamsungBrowser/5.4 Chrome/51.0.2704.106 Mobile Safari/537.36",
			true,
		}, {
			"Android Browser",
			"Mozilla/5.0 (Linux; U; Android 4.4.2; en-us; SCH-I535 Build/KOT49H) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Mobile Safari/534.30",
			true,
		},
	}

	for i, tst := range tsts {
		r, _ := http.NewRequest("GET", "", nil)
		r.Header.Add("User-Agent", tst.UA)
		mob := IsMobile(r)
		if mob != tst.Mobile {
			short := tst.UA
			if len(short) > 32 {
				short = short[0:32]
			}
			t.Errorf("%2d %-20v must %v; %v", i, tst.Label, tst.Mobile, short)
		}
	}

}
