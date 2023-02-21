package controller

import (
	"fmt"
	"github.com/MarioUTN/api_rest_go/database"
	"github.com/MarioUTN/api_rest_go/models"
)

func Create_Company(e models.Empresa) error {
	db, err := database.DatabaseConnection()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	query := "INSERT INTO company (emp_ruc, emp_nombre_empresa, emp_matriz, emp_sucursal, emp_pais, emp_provincia, emp_ciudad, emp_telefono, emp_email, nro_empleados, ingresos_anuales) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)"
	err = db.Exec(query, e.Emp_ruc, e.Emp_nombre_empresa, e.Emp_matriz, e.Emp_sucursal, e.Emp_pais, e.Emp_provincia, e.Emp_ciudad, e.Emp_telefono, e.Emp_email, e.Nro_empleados, e.Ingresos_anuales).Error
	return err
}

func Get_CompanyById(id int) (models.Empresa, error) {
	db, err := database.DatabaseConnection()
	var e models.Empresa
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	query := "SELECT id_empresa, emp_ruc, emp_nombre_empresa, emp_matriz, emp_sucursal, emp_pais, emp_provincia, emp_ciudad, emp_telefono, emp_email, nro_empleados, ingresos_anuales FROM company WHERE id_empresa = $1"
	row := db.Raw(query, id).Row()
	err = row.Scan(&e.Id_empresa, &e.Emp_ruc, &e.Emp_nombre_empresa, &e.Emp_matriz, &e.Emp_sucursal, &e.Emp_pais, &e.Emp_provincia, &e.Emp_ciudad, &e.Emp_telefono, &e.Emp_email, &e.Nro_empleados, &e.Ingresos_anuales)
	if err != nil {
		return e, err
	}
	return e, nil
}

func Get_AllCompanys() ([]models.Empresa, error) {
	db, err := database.DatabaseConnection()
	companys := []models.Empresa{}
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	query := "SELECT id_empresa, emp_ruc, emp_nombre_empresa, emp_matriz, emp_sucursal, emp_pais, emp_provincia, emp_ciudad, emp_telefono, emp_email, nro_empleados, ingresos_anuales FROM company"
	rows, err := db.Raw(query).Rows()
	if err != nil {
		return companys, err
	}
	for rows.Next() {
		var e models.Empresa
		err = rows.Scan(&e.Id_empresa, &e.Emp_ruc, &e.Emp_nombre_empresa, &e.Emp_matriz, &e.Emp_sucursal, &e.Emp_pais, &e.Emp_provincia, &e.Emp_ciudad, &e.Emp_telefono, &e.Emp_email, &e.Nro_empleados, &e.Ingresos_anuales)
		if err != nil {
			return companys, err
		}
		// and append it to the array
		companys = append(companys, e)
	}
	return companys, nil
}

func Update_Company(id int, e models.Empresa) error {
	db, err := database.DatabaseConnection()
	c, err := Get_CompanyById(id)
	if err != nil || c.Id_empresa == 0 {
		fmt.Println("Error: ", err.Error())
	} else {
		query := "UPDATE company SET emp_ruc = $1, emp_nombre_empresa = $2, emp_matriz = $3, emp_sucursal = $4, emp_pais = $5, emp_provincia = $6, emp_ciudad = $7, emp_telefono = $8, emp_email = $9, nro_empleados = $10, ingresos_anuales = $11 WHERE id_empresa = $12"
		err = db.Exec(query, e.Emp_ruc, e.Emp_nombre_empresa, e.Emp_matriz, e.Emp_sucursal, e.Emp_pais, e.Emp_provincia, e.Emp_ciudad, e.Emp_telefono, e.Emp_email, e.Nro_empleados, e.Ingresos_anuales, id).Error
	}
	return err
}

func Delete_Company(id int) error {
	db, err := database.DatabaseConnection()
	c, err := Get_CompanyById(id)
	if err != nil {
		fmt.Println("Error, Company not found!: ", err.Error())
	} else {
		query := "DELETE FROM company WHERE id_empresa = $1"
		err = db.Exec(query, c.Id_empresa).Error
	}
	return err
}
