package main
func main(){

	t := template.Must(template.New("hello").Parse("hello world"))
	t.Execute(os.Stdout, nil)
	
}