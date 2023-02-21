create database app_golang;
create table punto(
    id_punto serial not null,
    coord_x integer not null,
    coord_y integer not null,
    constraint pk_punto primary key(id_punto)
);
create table circunferencia(
    id_circunferencia serial not null,
    centro integer not null,
    radio numeric(6, 2) not null,
    constraint pk_circunferencia primary key(id_circunferencia),
    constraint fk_punto_circ foreign key(centro) references punto(id_punto)
);
create table parabola(
    id_parabola serial not null,
    vertice integer not null,
    p numeric(6, 2) not null,
    eje_focal varchar(1) not null,
    constraint pk_parabola primary key(id_parabola),
    constraint fk_punto_parab foreign key(vertice) references punto(id_punto)
);
create table elipse(
    id_elipse serial not null,
    centro integer not null,
    a numeric(6, 2) not null,
    b numeric(6, 2) not null,
    eje_focal varchar(1) not null,
    constraint pk_elipse primary key(id_elipse),
    constraint fk_punto_elipse foreign key(centro) references punto(id_punto)
);
CREATE TABLE empresa(
    id_empresa SERIAL NOT NULL,
    emp_ruc VARCHAR(13) NOT NULL,
    emp_nombre_empresa VARCHAR(100) NOT NULL,
    emp_matriz VARCHAR(100) NOT NULL,
    emp_sucursal VARCHAR(100),
    emp_pais VARCHAR(50) NOT NULL,
    emp_provincia VARCHAR(50) NOT NULL,
    emp_ciudad VARCHAR(50) NOT NULL,
    emp_telefono VARCHAR(15) NOT NULL,
    emp_email VARCHAR(60) NOT NULL UNIQUE,
    nro_empleados INTEGER NOT NULL,
    ingresos_anuales NUMERIC(10, 2),
    CONSTRAINT pk_empresa PRIMARY KEY(id_empresa)
);
select json_build_object(
        'id_punto',
        p.id_punto,
        'coord_x',
        p.coord_x,
        'coord_y',
        p.coord_y
    )
from punto p
insert into punto(coord_x, coord_y)
values(-1, 3);
insert into punto(coord_x, coord_y)
values(-4, -1);
insert into punto(coord_x, coord_y)
values(2, 5);
insert into punto(coord_x, coord_y)
values(2, -3);
insert into recta(punto_a, punto_b)
values(1, 4);
insert into recta(punto_a, punto_b)
values(2, 3);
insert into recta(punto_a, punto_b)
values(3, 4);
insert into recta(punto_a, punto_b)
values(5, 2);
insert into recta(punto_a, punto_b)
values(4, 5);