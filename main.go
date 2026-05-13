package main

import "fmt"

type Articulo struct {
	codigo int
	titulo string
	valor  float64
	stock  int
}

func NuevoArticulo(codigo int, titulo string, valor float64, stock int) Articulo {
	return Articulo{
		codigo: codigo,
		titulo: titulo,
		valor:  valor,
		stock:  stock,
	}
}
func (a Articulo) GetCodigo() int    { return a.codigo }
func (a Articulo) GetTitulo() string { return a.titulo }
func (a Articulo) GetValor() float64 { return a.valor }
func (a Articulo) GetStock() int     { return a.stock }

func (a *Articulo) ModificarValor(nuevoValor float64) {
	a.valor = nuevoValor
}

func (a *Articulo) ModificarStock(nuevoStock int) {
	a.stock = nuevoStock
}

func (a Articulo) VerInfo() {
	fmt.Printf("Codigo: %d | Titulo: %s | Valor: $%.2f | Stock: %d\n",
		a.codigo, a.titulo, a.valor, a.stock)
}

type Bodega struct {
	articulos map[int]Articulo
}

func NuevaBodega() Bodega {
	return Bodega{articulos: make(map[int]Articulo)}
}

func (b *Bodega) IngresarArticulo(a Articulo) {
	b.articulos[a.GetCodigo()] = a
	fmt.Printf("Articulo '%s' ingresado.\n", a.GetTitulo())
}

func (b *Bodega) RemoverArticulo(codigo int) {
	if _, existe := b.articulos[codigo]; existe {
		delete(b.articulos, codigo)
		fmt.Printf("Articulo con codigo %d removido.\n", codigo)
	} else {
		fmt.Printf("Articulo con codigo %d no encontrado.\n", codigo)
	}
}

func (b *Bodega) LocalizarArticulo(codigo int) (Articulo, bool) {
	a, existe := b.articulos[codigo]
	return a, existe
}

func (b *Bodega) MostrarArticulos() {
	if len(b.articulos) == 0 {
		fmt.Println("La bodega está vacía.")
		return
	}
	fmt.Println("******** BODEGA ********")
	for _, a := range b.articulos {
		a.VerInfo()
	}
	fmt.Println(" ")
}

func main() {
	bodega := NuevaBodega()

	cafe := NuevoArticulo(1, "Cafe", 2.30, 120)
	jabon := NuevoArticulo(2, "Jabon", 1.10, 60)
	atun := NuevoArticulo(3, "Atun", 1.80, 45)

	bodega.IngresarArticulo(cafe)
	bodega.IngresarArticulo(jabon)
	bodega.IngresarArticulo(atun)

	fmt.Println()
	bodega.MostrarArticulos()

	fmt.Println("\nActualizando Cafe")
	cafe.ModificarValor(2.70)
	cafe.ModificarStock(100)
	bodega.IngresarArticulo(cafe)
	cafe.VerInfo()

	fmt.Println("\nGetters de Cafe:")
	fmt.Printf("  GetCodigo  : %d\n", cafe.GetCodigo())
	fmt.Printf("  GetTitulo  : %s\n", cafe.GetTitulo())
	fmt.Printf("  GetValor   : $%.2f\n", cafe.GetValor())
	fmt.Printf("  GetStock   : %d\n", cafe.GetStock())

	fmt.Println("\nBuscando articulo 2")
	if a, existe := bodega.LocalizarArticulo(2); existe {
		a.VerInfo()
	} else {
		fmt.Println("No encontrado.")
	}

	fmt.Println("\nRemoviendo articulo 3")
	bodega.RemoverArticulo(3)
	fmt.Println()
	bodega.MostrarArticulos()
}
