package state

type GameState struct {
    Day     int
    Money   int
    Plots   []Plot  
    Season  string
    Weather Weather
}

type Plot struct { 
    ID      int
    Size    int    
    Crop    string  
    Growth  int     
}

type Weather struct {
    Name  string
    Left  int   
}