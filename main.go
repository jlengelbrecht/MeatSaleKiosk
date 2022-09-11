package main

import "fmt"

func main() {

	var TotG = 50
	var TotB = 50
	const TB = "Totes of Beef"
	const TG = "Totes of Grind"
	// variables to define ammount of product avalielbe. Constant values to define product.
	// will need to work with Receiving to get actual values for what is always avalielbe during meat sale.

	fmt.Println("Hello, Welcome to the Cargill meat sales program")
	fmt.Println("Select a product you would like to buy")
	fmt.Println()
	fmt.Printf("We currently have %v %v and %v %v\n", TotG, TG, TotB, TB)
	// Custom welcome message to the user. informes of what products are avalielbe and how much.

	var FirstName string
	var LastName string
	var EmployeeID string
	var CustomerType string
	var ProductSelected string
	// veriables set to define First & Last Name, Employee ID, customer type, and what product they wish to buy

	fmt.Println("Please enter if you are a EMPLOYEE or a CONTRACTOR:")
	fmt.Println("Indique si es EMPLEADO o CONTRATISTA:")
	fmt.Scan(&CustomerType)
	// asks user if they are a employee or contractor. scan command checks memory for response before proceeding

	fmt.Println("Please enter your first name:")
	fmt.Println("Por favor, introduzca su nombre de pila:")
	fmt.Scan(&FirstName)
	// asks user if their first name. scan command checks memory for response before proceeding

	fmt.Println("Please enter your last name:")
	fmt.Println("Por favor ingrese su apellido:")
	fmt.Scan(&LastName)
	// asks user if their last name. scan command checks memory for response before proceeding

	fmt.Println("Please enter your employee number - Enter N/A if you are a not a cargill employee:")
	fmt.Println("Por favor ingrese su número de empleado - Ingrese N/A si no es un empleado de Cargill:")
	fmt.Scan(&EmployeeID)
	// asks user for their Employee ID. N/A is used to identify a non employees. scan command checks memory for response before proceeding

	fmt.Println("Please enter if you wish to buy BEEF or GRIND:")
	fmt.Println("Por favor ingrese si desea comprar CARNE O MOLIENDA:")
	fmt.Scan(&ProductSelected)
	// asks user for their product selection. scan command checks memory for response before proceeding

	fmt.Printf("Thanks %v %v, one tote of %v has been ordered! Please provide your Employee ID to BSS on the day of the sale\n", FirstName, LastName, ProductSelected)
	fmt.Printf("Gracias %v %v, ¡se ha pedido una bolsa de %v! Proporcione su identificación de empleado a BSS el día de la venta\n", FirstName, LastName, ProductSelected)
	// gives a custom thank you message to show the order was processed.

}
