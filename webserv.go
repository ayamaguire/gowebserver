package main 

import ("fmt"
"net"
"net/http"
"io/ioutil"
"crypto/sha512"
"encoding/base64"
"time"
"strings"
)

/* Written January 2016 by Aya Maguire
*/


var listener net.Listener

/* I recogize that using this as a global variable is... nonideal.
Without writing a new HandleFunc which would let me pass the listener,
I wasn't sure how to pass it through to Shutdown.
I think I could find a more graceful solution with a little more time.
*/

func main() {

	/* Listen on port 8080 and hash some passwords!
	*/

	LaunchServ()

}

func Parse(s string) (string, int) {

	/* I'm assuming that the text you pass in is either "Shutdown"
	or "password=YourPassword". Anything else is an error.
	*/

	if strings.Index(s, "password=") == 0 {

		parsed := strings.Split(s, "=")

		return parsed[1], 0

	} else if s == "Shutdown" {

		return "", 1

	} else {

		panic("Invalid string received: " + s)
	}

}

func LaunchServ(){

	l, err := net.Listen("tcp", ":8080")
	if err != nil {
    	fmt.Println("%s", err)
	}
	listener = l

	http.HandleFunc("/", PwEncryptor)
	http.Serve(listener, nil)

}

func PwEncryptor(w http.ResponseWriter, r *http.Request){

	rawtext, err := ioutil.ReadAll(r.Body);
	
	if err != nil {
		fmt.Fprintf(w, "%s", err)
	}

	stringtext := string(rawtext[:])

	var pw string
	var signal int
	pw, signal = Parse(stringtext)

	if signal == 1 {
		/* This is a bit of a cheap way to ensure we don't close on connections
		which are already in flight.
		*/
		time.Sleep(5000 * time.Millisecond)
		Shutdown(listener)

	} else {

	b64pw := EncryptString(pw)

	time.Sleep(5000 * time.Millisecond)

	fmt.Println(pw, "Hashed to:", b64pw)
	fmt.Fprintf(w, "Hashing this password: " + pw + "\n")

	}

}

func Shutdown(l net.Listener){
	fmt.Println("Received shutdown request... shutting down.")

	l.Close()
}

func EncryptString(s string) string {

	encryptor := sha512.New()
	encryptor.Write([]byte(s))
	encryptedtext := encryptor.Sum(nil)
	b64text := base64.StdEncoding.EncodeToString([]byte(encryptedtext))

	return b64text

}

