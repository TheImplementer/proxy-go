package transformers

import (
    "image"
    "image/png"
    "net/http"
    "fmt"
    "bytes"
    "io/ioutil"
)

const PngContentType = "image/png"

func FlipPng(response *http.Response) (*http.Response) {
    if contentType := response.Header.Get("Content-Type"); contentType != PngContentType {
        fmt.Printf("unsupported response: %v\n", contentType)
        return nil
    }

    if img, err := png.Decode(response.Body); err != nil {
        fmt.Printf("can't decode image contained in response\n")
        return response
    } else {
        return createResponseForImage(response, img)
    }
}

func createResponseForImage(response *http.Response, img image.Image) (*http.Response) {
    transformed := flipImage(img)

    return &http.Response{
        Status: response.Status,
        StatusCode: response.StatusCode,
        Proto: response.Proto,
        ProtoMajor: response.ProtoMajor,
        ProtoMinor: response.ProtoMinor,
        Header: response.Header,
        Body: ioutil.NopCloser(transformed),
        ContentLength: int64(transformed.Len()),
        Request: response.Request,
    }
}

func flipImage(srcImage image.Image) (*bytes.Buffer) {
    imgBuffer := &bytes.Buffer{}
    png.Encode(imgBuffer, srcImage)
    return imgBuffer
}
