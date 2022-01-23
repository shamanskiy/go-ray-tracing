# go-ray-tracing
This is a Go implementation of the [awesome ray tracing tutorial](http://in1weekend.blogspot.com/2016/01/ray-tracing-in-one-weekend.html) by Peter Shirley.

<img src="https://raw.githubusercontent.com/Shamanskiy/go-ray-tracing/media/images/megaScene1280x720.png" width="700">

## Running
Go is so great that you can download all the dependencies for this project, build and run an application with just one command:
```
go run apps/megaScene/megaScene.go
```
Run this command from the project root and it will put the `megaScene.png` image there. 
BEWARE: the mega scene is called so for a reason. With the current settings, it takes 2h30m to render on my laptop (2017 Macbook 12'').
If you want something to render faster, run
```
go run apps/threeSpheres/threeSpheres.go
```
You will get this image with 3 spheres (as promised!):

<img src="https://raw.githubusercontent.com/Shamanskiy/go-ray-tracing/media/images/threeSpheres640x360.png" width="350">

## Testing
Execute the following command from the project root to run the unit tests:
```
go test ./...
```
