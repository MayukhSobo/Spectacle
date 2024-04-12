package app

var _banner = `
.d8888b.                             888                      888          
d88P  Y88b                            888                      888          
Y88b.                                 888                      888          
 "Y888b.   88888b.   .d88b.   .d8888b 888888  8888b.   .d8888b 888  .d88b.  
    "Y88b. 888 "88b d8P  Y8b d88P"    888        "88b d88P"    888 d8P  Y8b 
      "888 888  888 88888888 888      888    .d888888 888      888 88888888 
Y88b  d88P 888 d88P Y8b.     Y88b.    Y88b.  888  888 Y88b.    888 Y8b.     
 "Y8888P"  88888P"   "Y8888   "Y8888P  "Y888 "Y888888  "Y8888P 888  "Y8888  
           888                                                              
           888                                                              
           888                                                              
` 

func Banner() string {
	return _banner
}