package controller
import(
	"html/template"
	"net/http"
"fmt"
)
type Sample struct{
	Name        string
	Location string
	Gender     string
}
var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("src/views/*.html"))
}

func Index(w http.ResponseWriter, r *http.Request) {
        
    
		sam := Sample{
			
			Name:       "SS",
			Location: "HYD",
			Gender:       "Male",
			
		}
			res := []Sample{}
		res = append(res,sam)
		fmt.Println(res)
		tpl.ExecuteTemplate(w, "index", res)
}