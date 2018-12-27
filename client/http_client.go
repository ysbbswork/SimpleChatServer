package main
//
//func FormServer(w http.ResponseWriter, request *http.Request) {
//	w.Header().Set("Content-Type", "text/html")
//	strSrc,err :=os.Getwd()
//	fileSrc :=strSrc+"\\index.html"
//	t, err := template.ParseFiles(fileSrc)
//	if err != nil {
//		fmt.Println("parse file err:", err)
//		return
//	}
//
//	switch request.Method {
//	case "GET":{
//		p := []string{"nihao","lalala","end","sss"}
//		if err := t.Execute(w, p); err != nil {
//			fmt.Println("There was an error:", err.Error())
//		}
//	}
//
//	case "POST":{
//		input:= request.FormValue("in")
//		fmt.Println(input)
//	}
//
//	}
//
//
//}
//
//func main() {
//	http.HandleFunc("/chatroom", logPanics(FormServer))
//	if err := http.ListenAnd Serve(":8888", nil); err != nil {
//	}
//}
//
//func logPanics(handle http.HandlerFunc) http.HandlerFunc {
//	return func(writer http.ResponseWriter, request *http.Request) {
//		defer func() {
//			if x := recover(); x != nil {
//				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
//			}
//		}()
//		handle(writer, request)
//	}
//}
//





