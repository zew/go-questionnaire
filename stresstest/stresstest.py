import threading
import time
import requests
import sys


# stackoverflow.com/questions/16181121/

# 'dev.zew.de' should point to your development server - i.e. localhost
urlBase = "https://dev.zew.de:8081/survey"
pathLogin = "/direct/"

urls = [
    # French
    urlBase + pathLogin + "22R38",
    urlBase + pathLogin + "22S39",
    urlBase + pathLogin + "22T3A",
    urlBase + pathLogin + "22U3B",
    urlBase + pathLogin + "22V3C",
    urlBase + pathLogin + "22W3D",
    urlBase + pathLogin + "22X3E",
    urlBase + pathLogin + "22Y3F",
    urlBase + pathLogin + "22Z3G",
    urlBase + pathLogin + "2323H",
    urlBase + pathLogin + "2333K",
    urlBase + pathLogin + "2343L",
    urlBase + pathLogin + "2353M",
    urlBase + pathLogin + "2363N",
    urlBase + pathLogin + "2373P",
    urlBase + pathLogin + "2383R",
    urlBase + pathLogin + "2393S",
    urlBase + pathLogin + "23A3T",
    urlBase + pathLogin + "23B3U",
    urlBase + pathLogin + "23C3V",
    urlBase + pathLogin + "23D3W",
    urlBase + pathLogin + "23E3X",
    urlBase + pathLogin + "23F3Y",
    urlBase + pathLogin + "23G3Z",
    urlBase + pathLogin + "23H42",
    urlBase + pathLogin + "23K43",
    urlBase + pathLogin + "23L44",
    urlBase + pathLogin + "23M45",
    urlBase + pathLogin + "23N46",
    urlBase + pathLogin + "23P47",
    urlBase + pathLogin + "23R48",
    urlBase + pathLogin + "23S49",
    urlBase + pathLogin + "23T4A",
    urlBase + pathLogin + "23U4B",
    urlBase + pathLogin + "23V4C",
    urlBase + pathLogin + "23W4D",
    # Germans
    urlBase + pathLogin + "2NANT",
    urlBase + pathLogin + "2NBNU",
    # Belgians - English
    urlBase + pathLogin + "3FHG2",
    urlBase + pathLogin + "3FKG3",
    # Spaniards
    urlBase + pathLogin + "3NGNZ",
    urlBase + pathLogin + "3NHP2",
    # Italians
    urlBase + pathLogin + "3YGYZ",
    urlBase + pathLogin + "3YHZ2",
    # Poles
    urlBase + pathLogin + "44U5B",
    urlBase + pathLogin + "44V5C",
]


def printResponseData(url, req, testString):
    print("\nHeaders %s" % req.headers)
    print("Url %s - Status Code %d" % (req.url, req.status_code))
    if testString not in req.text:
        # print(r.text)
        print("MISSING %s" % testString)
        sys.exit()
    else:
        print("FOUND %s" % testString)

    print("'%s' fetched in %ss" % (url, (time.time() - start)))






def walkThroughApplication(startURL):

    # print("Trying URL %s" % startURL)
    with requests.session() as s:
        
        # Navigate to page 1
        resp = s.get(startURL, params={}, allow_redirects=True)
        testString = "type='submit' name='submitBtn' value='1'"
        printResponseData(startURL, resp, testString)

        # Extract request token from response
        # Not the session key
        loc1 = 'type="hidden" name="token" value="'
        if loc1 not in resp.text:
            print("Request token at location '%s' not found" % loc1)
            return
        pos1 = resp.text.find(loc1)
        pos2 = resp.text.find("\"",len(loc1)+pos1+1)
        tkn = resp.text[pos1+len(loc1):pos2]

        print("Pos1 %d - pos2 %d. Request token is %s" % (pos1, pos2, tkn))
        print(resp.text[pos1+len(loc1)-2:pos2+2])


        # Navigate to page 2
        paramsPage2 = 'submitBtn=1&xx=yy'.split('&')
        pp2 = {}
        for item in paramsPage2:
            key, value = item.split('=')
            if value:
                pp2[key] = value
        pp2["token"] = tkn
        resp = s.post(urlBase, params=pp2, allow_redirects=True)
        testString = 'type="submit" name="submitBtn" value="next"'
        printResponseData(urlBase, resp, testString)

        # with open('x.htm', 'wb') as f:
        #     f.write(r.text.encode('utf8'))





start = time.time()

threads = [threading.Thread(target=walkThroughApplication, args=(url,)) for url in urls]
for thread in threads:
    thread.start()
for thread in threads:
    thread.join()

print("\nElapsed Time: %s" % (time.time() - start))




