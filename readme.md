# Explicación de Keywords en Go

En el lenguaje de programación Go, hay tres keywords importantes que controlan el flujo de ejecución:

## 1. defer

La keyword `defer` se utiliza para posponer la ejecución de una función hasta que la función padre que la contiene haya terminado. Esto es útil para asegurarse de que ciertas acciones se realicen al final de una función, como cerrar archivos o liberar recursos.

## 2. continue

La keyword `continue` se utiliza en bucles para saltar a la siguiente iteración sin ejecutar el resto del código dentro del bucle en esa iteración en particular. Es útil cuando deseas evitar que cierto código se ejecute bajo ciertas condiciones dentro de un bucle.

## 3. break

La keyword `break` se utiliza para salir inmediatamente de un bucle. Cuando se encuentra la declaración `break`, el bucle en el que se encuentra se termina abruptamente, y la ejecución continúa con el código que sigue después del bucle. Es útil cuando necesitas salir de un bucle antes de que se complete todas las iteraciones, generalmente en respuesta a ciertas condiciones.
