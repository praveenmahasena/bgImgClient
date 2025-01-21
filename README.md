# BG_Img_Server

This repo has code for the server of Background_Image_changer
It's written using [Golang](https://go.dev/dl/).


### Prerequisites
- [server](https://github.com/praveenmahasena/bgImgServer) client repo
- [Golang](https://go.dev/dl/) programming language
- [feh](https://feh.finalrewind.org/)

### Installation

```
git clone git@github.com:praveenmahasena/bgImgClient.git
cd bgImgClient
make
```

This app connects to the server via TCP connection and the server returns a slice of byte and the bytes would be written into a file and it would be fed into (feh)[https://feh.finalrewind.org/] to render it as background image
