package gopdf

type imgInfo struct {
	w, h int
	//src              string
	formatName       string
	colspace         string
	bitsPerComponent string
	filter           string
	decodeParms      string
	trns             []byte
	smask            []byte
	smaskObjID       int
	pal              []byte
	deviceRGBObjID   int
	data             []byte
}
