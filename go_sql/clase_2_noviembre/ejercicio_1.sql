--  ¿A qué se denomina JOIN en una base de datos y para qué se utiliza?
--      Se utiliza para obtener datos de varias tablas relacionadas entre sí. Consiste en combinar datos de una tabla con datos de la otra tabla, a partir de una o varias condiciones en común.
--  Explicar dos tipos de JOIN.
--      Inner Join se utiliza para traer los datos relacionados de dos o más tablas.
--      Left Join se utiliza para traer los datos de la tabla izquierda más los relacionados de la tabla derecha.
--  ¿Para qué se utiliza el GROUP BY?
--      Agrupa los resultados según las columnas indicadas.
--      Genera un solo registro por cada grupo de filas que compartan las columnas indicadas.
--      Reduce la cantidad de filas de la consulta.
--      Se suele utilizar en conjunto con funciones de agregación, para obtener datos resumidos y agrupados por las columnas que se necesiten.
--  ¿Para qué se utiliza el HAVING?
--      La cláusula HAVING se utiliza para incluir condiciones con algunas funciones SQL.
--      Afecta a los resultados traidos por Group By.

-- Escribir una consulta genérica para cada uno de los siguientes diagramas:

describe series;
select
    m.title titulo,
    g.name genero
from movies m
left join genres g on m.genre_id = g.id;

select
    m.title titulo,
    g.name genero
from movies m
inner join genres g on m.genre_id = g.id;

-- Se propone realizar las siguientes consultas a la base de datos movies_db.sql trabajada en la primera clase.

--  Mostrar el título y el nombre del género de todas las series.
select
    s.title,
    g.name
from series s
join genres g on g.id = s.genre_id;
--  Mostrar el título de los episodios, el nombre y apellido de los actores que trabajan en cada uno de ellos.
describe actor_episode;
select
    e.title,
    a.first_name,
    a.last_name
from actors a
join actor_episode ae on a.id = ae.actor_id
join episodes e on ae.episode_id = e.id;
--  Mostrar el título de todas las series y el total de temporadas que tiene cada una de ellas.
select
   s.title,
   COUNT(1) as 'Total Temporadas'
from series s
join seasons s2 on s.id = s2.serie_id
group by title;
--  Mostrar el nombre de todos los géneros y la cantidad total de películas por cada uno, siempre que sea mayor o igual a 3.
select distinct
   g.name as "nombre",
   count(1) as "total"
from movies m
join genres g on m.genre_id = g.id
group by  g.name having count(1) > 3;
--  Mostrar sólo el nombre y apellido de los actores que trabajan en todas las películas de la guerra de las galaxias y que estos no se repitan.
select
    nombre,
    apellido
from
(select
   a.first_name as "nombre",
   a.last_name as "apellido"
from actors a
join actor_movie am on a.id = am.actor_id
join movies m on am.movie_id = m.id
group by a.first_name, a.last_name, m.title
having m.title like 'La Guerra de las galaxias:%') as x
group by nombre, apellido
having count(1) = 2