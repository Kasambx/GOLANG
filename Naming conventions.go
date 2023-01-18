pacakge main
import (
  "fnt"
)
const a init = 27
func main(){
        const a int = 14
        var b init = 30  
  `     fmt.Printf("%V,%T\n", a, a) 
        fmt.Printf("%V,%T\n",a+b, a+b)
        cost x = 42 
        var y init16 = 27
        fmt.Printf("%V,%T\n",x+y, X+y)



/// Enumerators 
const (
       f = iota
       g
       h)
const (
        f2 = iota 

       fmt.Printf("%V\n",f)
       fmt.Printf("%V\n",g)
       fmt.Printf("%V\n",h)
       fnt.Printf("%V\n",f2)
       
const (
      errorSpecialist = iota 
      catSpecialist
      dogSpecialist 
      snakeSpecialist
)

func main() { 
         var specialistType int = dogSPecialist
         fmt.Printf("%v\n",specialistType == dogSpecialist)
         
         
         
         }
cont (
        _ = iota // ignore first value by assigning to blank identifier
        KB = 1 << (10 * iota )
        MB
        GB
        TB
        PB
        EB
        ZB
        YB
        )
func main() {
        fileSize:= 4000000000.
        fmt.Printf("%.2fGB",filesize/GB)
   }
   
   
const (
        isAdmin = 1 << iota 
        is Headquaters
        canSeeFinancials
        
        canSeeAfrica
        canSeeAsia
        canSeeEurope
        canSeeNorthAmerica
        canSeeSouthAmerica
)

func main () {
        var roles byte = isAdmin | canSeeFinancials | canSeeEurope
        fmt.Printf("%b\n",roles )
        fmt.Printf("is Admin? %v",isAdmin & roles == isAdmin)
         fmt.Printf("is HQ? %v",isHeadquaters & roles == isHeadquaters)
}


}

