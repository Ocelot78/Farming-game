package state

type GameState struct {
    Day     int
    Hour    int
    Money   int
    Fields  []Field  
    Season  string
    Weather Weather
    OwnedT  Tractor
    
}

type Field struct { 
    ID      int
    Size    int    
    Crop    string  
    Growth  int     
}

type Weather struct {
    Name  string
    Left  int   
}