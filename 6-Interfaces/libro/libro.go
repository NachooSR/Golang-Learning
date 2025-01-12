package libro

import "fmt"

//1) ***********ESTRUCTURAS***********

type Libro struct {
	Titulo string
	Autor  string
	Pages  int
}

type LibroTexto struct {
	Libro
	editorial string
	nivel     string
}


//2) ***********CONSTRUCTORES***********
func NuevoLibro(titulo, autor string, paginas int) *Libro {
	return &Libro{
		Titulo: titulo,
		Autor:  autor,
		Pages:  paginas,
	}
}

//Composicion (herencia)
func NuevoTextoLibro(titulo, autor, editorial, nivel string, paginas int) *LibroTexto {
	return &LibroTexto{
		Libro:     Libro{titulo, autor, paginas},
		editorial: editorial,
		nivel:     nivel,
	}
}




//3) ***********INTERFAZ***********
type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}



// 4) ***********METODOS***********
func (l Libro) PrintInfo() {

	fmt.Println("Libro:", l.Titulo)
	fmt.Println("Autor:", l.Autor)
	fmt.Println("Cant.Pags:", l.Pages)
}

func (l LibroTexto) PrintInfo(){
	fmt.Println("Libro:", l.Titulo)
	fmt.Println("Autor:", l.Autor)
	fmt.Println("Cant.Pags:", l.Pages)
	fmt.Println("Editorial: ",l.editorial)
	fmt.Println("Nivel: ",l.nivel)
}




// 5) ***********ACCESORS***********

func (l *Libro) SetTitulo(titulo string) {
	l.Titulo = titulo
}

func (l Libro) GetTitulo() string {
	return l.Titulo
}
/*Simular accesors ya que no existen las propiedades private
o public diferenciamos con mayuscula (Publico) o minuscula (Private)
*/



