package models

import (
	"gorm.io/gorm"
)

type Empresa struct {
	gorm.Model

	Id_empresa         uint    `json:"id_empresa"`
	Emp_ruc            string  `json:"emp_ruc"`
	Emp_nombre_empresa string  `json:"emp_nombre_empresa"`
	Emp_matriz         string  `json:"emp_matriz"`
	Emp_sucursal       string  `json:"emp_sucursal"`
	Emp_pais           string  `json:"emp_pais"`
	Emp_provincia      string  `json:"emp_provincia"`
	Emp_ciudad         string  `json:"emp_ciudad"`
	Emp_telefono       string  `json:"emp_telefono"`
	Emp_email          string  `json:"emp_email"`
	Nro_empleados      int     `json:"nro_empleados"`
	Ingresos_anuales   float64 `json:"ingresos_anuales"`
}
