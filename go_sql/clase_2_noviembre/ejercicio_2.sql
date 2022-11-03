-- Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores.
select
    e.nombre,
    d.nombre_depto,
    d.localidad
from EMPLEADO e
left join DEPARTAMENTO d
on e.depto_nro = d.depto_nro;

-- Visualizar los departamentos con más de cinco empleados.
select
    d.nombre_depto,
    count(e.cod_emp)
from EMPLEADO e
left join DEPARTAMENTO d
on e.depto_nro = d.depto_nro
group by d.depto_nro
having count(e.cod_emp) > 5;

-- Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’.
select
    e.nombre,
    e.salario,
    d.nombre_depto
from EMPLEADO e
left join DEPARTAMENTO d
on e.depto_nro = d.depto_nro
where e.puesto = (select puesto from EMPLEADO where cod_emp = 'E-0006');

-- Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre.
select
    e.cod_emp,
    e.nombre,
    e.apellido,
    e.puesto,
    e.fecha_alta,
    e.salario
from EMPLEADO e
where e.depto_nro = 'D-000-3'
order by e.nombre;

-- Mostrar el nombre del empleado que tiene el salario más bajo.
select
    nombre
from EMPLEADO
order by salario asc
limit 1;

-- Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’.
select
    nombre,
    salario
from EMPLEADO
where depto_nro = 'D-000-4'
order by salario desc
limit 1