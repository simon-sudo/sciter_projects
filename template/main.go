/*

# linux

## prepare the icon

Create an app icon set of different size icons the compile the files to a single  `icns` file.
Use an app like `https://github.com/onmyway133/IconGenerator/` and generate an `icns` file from a `svg` file.

Alternativelly, download a free icon in all formats from
http://www.iconarchive.com/ or https://icon-icons.com/

### app searching: for desktop
Copy the desktop file: `~/.local/share/applications/cauta.desktop`.
Copy the icon file: `~/.local/share/icons/cauta.svg`.

### app running: for panel/taskbar/dock
Add the `<html lang="ro" window-icon="images/app.png">` to the `html` file.

## create the binary
Copy `sciter-sdk-master/bin.gtk/x64/packfolder` to project folder.

```
./packfolder ./res/app res.go -v resources -go
go build
./cauta
```

## terminal

```
./packfolder ./res/app res.go -v resources -go;go build;./cauta
```

## deploy
Copy binary to `~/bin`.
Copy `/res/linux/.local` to `~/.local`.


# macos

## prepare the icons file

### app searching: for desktop
Create an app bundle, with `appify` (see below)


### app running: for dock
Add the `<html lang="ro" window-icon="images/app.png">` to the `html` file.

## create the binary

```
go get github.com/machinebox/appify
./packfolder ./res/app res.go -v resources -go
go build
```

## add the icon file to binary

```
appify -name "cauta" -icon ./res/app/images/app.png cauta
```

Alternatively:

```
appify -name "cauta" cauta
```

edit `cauta.app/Contents/Info.plist`

```
<key>CFBundleIconFile</key>
<string>app</string>
```

then copy the icons file from `./res/mac/app.icns` to `cauta.app/Contents/Resources/app.icns`

## terminal

```
./packfolder ./res/app res.go -v resources -go;go build;appify -name "cauta" -icon ./res/app/images/app.png cauta;open cauta.app
```


# windows 10

## prepare the icon

### app searching: for desktop

```
go get github.com/akavel/rsrc
rsrc -ico ./res/win/app.ico -o win-res.syso -arch="amd64"
```

### app running: for panel/taskbar/dock
Add the `<html lang="ro" window-icon="images/app.png">` to the `html` file.

## create binary

For the win-res file to be automatically included, copy it to the project root path.

```
packfolder.exe ./res/app res.go -v resources -go
go build -ldflags -H=windowsgui
```

## terminal

```
packfolder.exe ./res/app res.go -v resources -go
rsrc -ico ./res/win/app.ico -o win-res.syso -arch="amd64"
go build -ldflags -H=windowsgui
cauta
```

*/

package main

import (
	"log"
	"syscall"

	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
)

func main() {
	w, err := window.New(sciter.DefaultWindowCreateFlag, sciter.DefaultRect)
	if err != nil {
		log.Fatal(err)
	}

	w.SetResourceArchive(resources)
	w.LoadFile("this://app/maxi.html")

	w.DefineFunction("goclose", goclose)

	w.Show()
	w.Run()
}

func goclose(vals ...*sciter.Value) *sciter.Value {
	syscall.Exit(0)
	return nil
}
