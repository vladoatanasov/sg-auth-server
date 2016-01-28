package rest

//Auth body payload for authenticating to sync gateway
// type Auth struct {
// 	Name string `json:"name"`
// 	TTL  int    `json:"ttl"`
// }
//
// //Credentials used for authenticating the write user
// type Credentials struct {
// 	Username string `json:"username"`
// 	Password string `json:"password"`
// }
//
// var globalHTTP = &http.Client{}
//
// func authenticate(r *http.Request) string {
// 	document, err := ioutil.ReadAll(r.Body)
// 	logError(err)
//
// 	credentials := &Credentials{}
// 	err = json.Unmarshal(document, credentials)
// 	logError(err)
//
// 	auth := &Auth{}
// 	auth.Name = credentials.Username
// 	auth.TTL = config.TTL
//
// 	json, err := json.Marshal(auth)
// 	logError(err)
//
// 	request, err := http.NewRequest("POST", getSyncSessionEndpoint(), bytes.NewReader([]byte(json)))
// 	logError(err)
//
// 	logRequest(request)
//
// 	response, err := globalHTTP.Do(request)
// 	logError(err)
//
// 	defer response.Body.Close()
//
// 	document, err = ioutil.ReadAll(response.Body)
// 	logError(err)
//
// 	return string(document)
// }
//
// func logRequest(request *http.Request) {
// 	log, _ := httputil.DumpRequest(request, true)
// 	logg.LogTo(logTag, "%s", log)
// }
//
// func getSyncSessionEndpoint() string {
// 	return config.SyncEndpoint + "/" + config.Bucket + "/_session"
// }
//
// func logError(err error) {
// 	if err != nil {
// 		logg.LogPanic("%v", err)
// 	}
// }
