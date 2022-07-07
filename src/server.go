package main
import(
  "log"
  "net/http"
  "os"
)
type ShakeHandler struct{}
func (h *ShakeHandler)ServeHTTP(w http.ResponseWriter, req *http.Request) {
  defer req.Body.Close()
  w.Header().Set("Content-Type", "text/plain")
  w.WriteHeader(http.StatusOK)
  w.Write([]byte("OK"))
}
func main(){
  healthCheckPort:="6666"
  if os.Getenv("HEALTH_CHECK_PORT")!=""{
    healthCheckPort=os.Getenv("HEALTH_CHECK_PORT")
  }
  mux := http.NewServeMux()
  shakeHandler:=ShakeHandler{}
  mux.Handle("/",&shakeHandler)
  err := http.ListenAndServe(":"+healthCheckPort, mux)
  if err != nil {
      log.Fatal("ListenAndServe: ", err)
  }
}
