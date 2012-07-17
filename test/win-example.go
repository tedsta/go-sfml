package main

import (	
	"sfml/gfx"
	"sfml/sys"
	. "sfml/win"
	"sfml/win"
	"fmt"
	"log"
)

func Debug(x interface{}){	fmt.Printf("%#+v\n", x) }

func AwesomeKeyHandler(ke KeyEvent) {
	switch ke.Type {
	case EvtKeyPressed:
		log.Println("Key Pressed: ", ke)
	case EvtKeyReleased:
		log.Println("Key Released: ", ke)
	}
}

func main() {
	c := sys.NewClock()
	vm := NewVideoMode(512,512,24)

	w, err := NewWindow(
		vm,		
		"HelloWorld",
		StyleDefaultStyle,
		ContextSettings{})
	if err != nil {
		log.Fatal(err)
	}

	img, err := gfx.ImageFromFile("./gopher.png")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(img.Getsize())

	w.SetFramerateLimit(60);

	t := c.GetElapsedTime()		

    for w.IsOpen() {
		t = c.GetElapsedTime()			
		
		e := w.PollEvent()
		switch e.(type) {
		case KeyEvent:
			AwesomeKeyHandler(e.(KeyEvent))
		case MouseMoveEvent:		
			me := e.(MouseMoveEvent)
			log.Printf("MouseMove <%d, %d>\n", me.X(), me.Y())
		case nil:
			log.Println("LOOOOOOOOOL")
		}
	}
	log.Println(t.AsSeconds())
}



// 	// fnt := gfx.FontCreateFromFile("./Inconsolata.otf", 96)
// 	// Debug(fnt)

// 	// txt := gfx.StringCreate()
// 	// txt.SetText("Hello Go")
// 	// txt.SetFont(fnt)
// 	// Debug(txt)

// 	img := gfx.ImageCreateFromFile("../test/gopher.png")
// 	gopher := gfx.CreateSprite()
// 	gopher.SetImage(img)
// 	gopher.SetX(200)
// 	gopher.SetY(200)
	
// 	seagreen := gfx.FromRGB(244,23,34)
// 	// //seagreen := gfx.ColorFromRGBA(244,23,34,34)
// 	// //seagreen = gfx.ColorFromRGB(0,34,23)
// 	// Debug(seagreen)

// 	//frame := 0 
// 	//evt := gfx.Event{}
	

// 	// app.SetFramerateLimit(60)

// 	// var tick float32 = 0
// 	// for app.IsOpened() {				
// 	// 	//tick := clock.GetTime()
// 	// 	//txt.SetText(fmt.Sprintf("%v fps", app.GetFrameTime()))
// 	// 	tick += .001;
// 	// 	gopher.SetRotation(tick*100)
// 	// 	scale := float32(1 + (math.Sin(float64(tick)/10)))
// 	// 	gopher.SetScaleX(scale)
// 	// 	gopher.SetScaleY(scale)

// 	// 	app.DrawSprite(gopher)
// 	// 	app.DrawString(txt)
		
// 	// 	seagreen.R += 1
// 	// 	seagreen.G -= 1
// 	// 	app.Display()
// 	// 	app.Clear(seagreen)
// 	// }
// }

