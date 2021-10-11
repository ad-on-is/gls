package iconizer

type Icon_Info struct {
	i string
	c [3]uint8 // represents the color in rgb (default 0,0,0 is black)
}

func (i *Icon_Info) GetGlyph() string {
	return i.i
}

func (i *Icon_Info) GetColor() [3]uint8 {
	return i.c
}

// default icons in case nothing can be found
var Icon_Def = map[string]*Icon_Info{
	"dir":        {i: "\uf74a", c: [3]uint8{224, 177, 77}},
	"diropen":    {i: "\ufc6e", c: [3]uint8{224, 177, 77}},
	"hiddendir":  {i: "\uf755", c: [3]uint8{224, 177, 77}},
	"exe":        {i: "\uf713", c: [3]uint8{76, 175, 80}},
	"file":       {i: "\uf723", c: [3]uint8{65, 129, 190}},
	"hiddenfile": {i: "\ufb12", c: [3]uint8{65, 129, 190}},
}
