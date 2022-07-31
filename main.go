package main

import (
    "fmt"
    "io"
    "os"
    "io/ioutil"
    "net/http"
)

func IthPowerOfTwo(i int) int {
    return 1 << i
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
    // Parse our multipart form, 10 << 20 specifies a maximum
    // upload of 10 MB files.
    r.ParseMultipartForm(10 << 20)

    // FormFile returns the first file for the given key.
    // it also returns the FileHeader so we can get the Filename,
    // the Header and the size of the file
    file, handler, err := r.FormFile("file-upload")

    if err != nil {
        fmt.Println("Error Retrieving the File")
        fmt.Println(err)
        return
    }

    defer file.Close()

    fmt.Printf("Uploaded File: %+v\n", handler.Filename)
    fmt.Printf("File Size: %+v\n", handler.Size)
    fmt.Printf("MIME Header: %+v\n", handler.Header)

    // Create a temporary file within our temp-images directory that follows
    // a particular naming pattern
    tempFile, err := ioutil.TempFile("temp-images", "upload-*.png")
    if err != nil {
        fmt.Println(err)
    }

    defer tempFile.Close()

    // read all of the contents of the uploaded file into a byte array
    fileBytes, err := ioutil.ReadAll(file)
    if err != nil {
        fmt.Println(err)
    }

    // write this byte array to our temporary file
    tempFile.Write(fileBytes)

    // return that we have successfully uploaded our file!
    fmt.Fprintf(w, "Successfully Uploaded File\n")
}

func setupRoutes() {
    http.HandleFunc("/upload", uploadFile)

    fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)

    fmt.Println("Listening on port 5000")
    http.ListenAndServe(":5000", nil)
}

func main() {
    setupRoutes()
}

func basic() {
    buffer_size := IthPowerOfTwo(24)

    file, err := os.Open("example.mov")

    if err != nil {
        fmt.Println(err)
        return
    }

    defer file.Close()

    i := 1

    for {
        buffer := make([]byte, buffer_size)
        bytes_read, err := file.Read(buffer)

        if err != nil {
            if err != io.EOF {
                fmt.Println(err)
            }

            break
        }

        fmt.Println(buffer, bytes_read)
        //   fmt.Println(buffer[:bytesread])
        i += 1
    }

    fmt.Println("chunks:", i)
    fmt.Println("buffer size:", buffer_size)
}
