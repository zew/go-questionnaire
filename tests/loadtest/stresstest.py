import threading
import time
import requests
import sys

print("Set is_production to True")
time.sleep(3)


# stackoverflow.com/questions/16181121/

# 'dev.zew.de' should point to your development server - i.e. localhost
urlBase = "https://dev.zew.de:8081/survey"
pathLogin = "/d/"

urls = [
    urlBase + pathLogin + "RZWPB7",
    urlBase + pathLogin + "7P8GA7",
    urlBase + pathLogin + "7EWV87",
    urlBase + pathLogin + "74D457",
    urlBase + pathLogin + "7K9MX7",
    urlBase + pathLogin + "RW3NDR",
    urlBase + pathLogin + "RDYKB7",
    urlBase + pathLogin + "RYNZM7",
    urlBase + pathLogin + "RXB58R",
    urlBase + pathLogin + "R9XZW7",
    urlBase + pathLogin + "RB6M6R",
    urlBase + pathLogin + "R3GYWR",
    urlBase + pathLogin + "75ZYPR",
    urlBase + pathLogin + "76PY67",
    urlBase + pathLogin + "7GZ99R",
    urlBase + pathLogin + "RAXMZ7",
    urlBase + pathLogin + "RLBPER",
    urlBase + pathLogin + "7N84Z7",
    urlBase + pathLogin + "7M8E3R",
    urlBase + pathLogin + "784YX7",
    urlBase + pathLogin + "RV43YR",
    urlBase + pathLogin + "7ZW5BR",
    urlBase + pathLogin + "RP8WAR",
    urlBase + pathLogin + "REWA8R",
    urlBase + pathLogin + "R4DG5R",
    urlBase + pathLogin + "RK9ZXR",
    urlBase + pathLogin + "7W39D7",
    urlBase + pathLogin + "7DYMBR",
    urlBase + pathLogin + "7YN8MR",
    urlBase + pathLogin + "7XBE87",
    urlBase + pathLogin + "79X3WR",
    urlBase + pathLogin + "7B6P67",
    urlBase + pathLogin + "73G9W7",
    urlBase + pathLogin + "R5Z9P7",
    urlBase + pathLogin + "R6PL6R",
    urlBase + pathLogin + "RGZD97",
    urlBase + pathLogin + "7AXZZR",
    urlBase + pathLogin + "7LBLE7",
    urlBase + pathLogin + "RN8VZR",
    urlBase + pathLogin + "RM8L37",
    urlBase + pathLogin + "R849XR",
    urlBase + pathLogin + "7V4ZY7",
    urlBase + pathLogin + "RZWKB7",
    urlBase + pathLogin + "7P85A7",
    urlBase + pathLogin + "7EWK87",
    urlBase + pathLogin + "74DL57",
    urlBase + pathLogin + "7K9LX7",
    urlBase + pathLogin + "RW3PDR",
    urlBase + pathLogin + "RDY4B7",
    urlBase + pathLogin + "RYN4M7",
    urlBase + pathLogin + "RXBD8R",
    urlBase + pathLogin + "R9XBW7",
    urlBase + pathLogin + "RB6Z6R",
    urlBase + pathLogin + "R3G6WR",
    urlBase + pathLogin + "75ZLPR",
    urlBase + pathLogin + "76PB67",
    urlBase + pathLogin + "7GZW9R",
    urlBase + pathLogin + "RAXNZ7",
    urlBase + pathLogin + "RLBEER",
    urlBase + pathLogin + "7N8KZ7",
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




