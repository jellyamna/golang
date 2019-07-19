package main

import (
	"database/sql"
	"flag"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// Define an application struct to hold the application-wide dependencies for the
// web application. For now we'll only include fields for the two custom loggers, but
// we'll add more to it as the build progresses.
type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	// Define a new command-line flag for the MySQL DSN string.
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "MySQL data source name")
	flag.Parse()
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	// To keep the main() function tidy I've put the code for creating a connection
	// pool into the separate openDB() function below. We pass openDB() the DSN
	// from the command-line flag.
	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	// We also defer a call to db.Close(), so that the connection pool is closed
	// before the main() function exits.
	defer db.Close()
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func maindb() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
func mainoutes() {

	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()

	// log new adalah log safe yang bisa di pakai oleh concurency program.
	//without needing to worry about race conditions.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/", app.home)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Initialize a new http.Server struct. We set the Addr and Handler fields so
	// that the server uses the same network address and routes as before, and set
	// the ErrorLog field so that the server now uses the custom errorLog logger in
	// the event of any problems.
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	// Call the ListenAndServe() method on our new http.Server struct.
	err := srv.ListenAndServe()
	errorLog.Fatal(err)

}

func mainbackup2() {
	addr := flag.String("addr", ":4000", "HTTP network address")
	flag.Parse()
	// Use log.New() to create a logger for writing information messages. This takes
	// three parameters: the destination to write the logs to (os.Stdout), a string
	// prefix for message (INFO followed by a tab), and flags to indicate what
	// additional information to include (local date and time). Note that the flags
	// are joined using the bitwise OR operator |.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	// Create a logger for writing error messages in the same way, but use stderr as
	// the destination and use the log.Lshortfile flag to include the relevant
	// file name and line number.
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	mux := http.NewServeMux()
	// mux.HandleFunc("/", home)
	// mux.HandleFunc("/snippet", showSnippet)
	// mux.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// Write messages using the two new loggers, instead of the standard logger.

	infoLog.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)

}
func mainbackup1() {

	// Define a new command-line flag with the name 'addr', a default value of ":4000"
	// and some short help text explaining what the flag controls. The value of the
	// flag will be stored in the addr variable at runtime.
	addr := flag.String("addr", ":4000", "HTTP network address")

	// Importantly, we use the flag.Parse() function to parse the command-line flag.
	// This reads in the command-line flag value and assigns it to the addr
	// variable. You need to call this *before* you use the addr variable
	// otherwise it will always contain the default value of ":4000". If any errors are
	// encountered during parsing the application will be terminated.
	flag.Parse()

	mux := http.NewServeMux()
	// mux.HandleFunc("/", home)
	// mux.HandleFunc("/snippet", showSnippet)
	// mux.HandleFunc("/snippet/create", createSnippet)
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	// The value returned from the flag.String() function is a pointer to the flag
	// value, not the value itself. So we need to dereference the pointer (i.e.
	// prefix it with the * symbol) before using it. Note that we're using the
	// log.Printf() function to interpolate the address with the log message.
	log.Printf("Starting server on %s", *addr) // Information message
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err) //The log.Fatal() function will also call os.Exit(1) after writing the message, causing the application to immediately exit.
	// Error message

	//untuk ekseusinya mengunakan go run ./cmd/web -addr=":9999"
	// -->2018/08/08 19:15:47 Starting server

	//go run ./cmd/web -help -->untuk melihat apa saja yang tersedia

}

//main_backup is backup func
func mainbackup() {
	mux := http.NewServeMux()
	// mux.HandleFunc("/", home)               // in adalah subtree path.
	// mux.HandleFunc("/snippet", showSnippet) // ini adalah fix path
	// mux.HandleFunc("/snippet/create", createSnippet)

	// Create a file server which serves files out of the "./ui/static" directory.
	// Note that the path given to the http.Dir function is relative to the project
	// directory root.
	fileServer := http.FileServer(http.Dir("./ui/static/"))

	// Use the mux.Handle() function to register the file server as the handler for
	// all URL paths that start with "/static/". For matching paths, we strip the
	// "/static" prefix before the request reaches the file server.
	//mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Println("Starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
